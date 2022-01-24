package controller

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/InOuttt/snr/insert/repository"
	"github.com/InOuttt/snr/insert/service"
)

type FeedController interface {
	Handle(dirPath string)
}

type feedController struct {
	service service.FeedService
	repo    repository.FeedRepository
}

func NewFeedController(svc service.FeedService, rp repository.FeedRepository) FeedController {
	return feedController{
		service: svc,
		repo:    rp,
	}
}

func (fc feedController) Handle(dirPath string) {
	ctx := context.Background()

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan string, 3)
	for _, f := range files {
		if !f.IsDir() {
			go fc.ReadLineByLine(ctx, dirPath+"/"+f.Name(), c)
		}
	}

	for running := true; running; {
		var s string
		select {
		case s, running = <-c:
			fmt.Println(s)
		}
	}
	defer close(c)

}

func (fc feedController) ReadLineByLine(ctx context.Context, filepath string, out chan string) {

	file, err := os.Open(filepath)
	if err != nil {
		print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("inserting %v", scanner.Text())
		fc.service.CreateFeed(ctx, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		print(err)
	}
}
