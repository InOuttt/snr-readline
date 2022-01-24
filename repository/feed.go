package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/InOuttt/snr/insert/config"
	"github.com/InOuttt/snr/insert/model"
	"github.com/vanng822/go-solr/solr"
)

// DB: MySQL
const (
	queryInsert = `INSERT INTO ` + model.FeedTableName + `
		(feed_id, published_date, link, content, author_name, author_link, author_avatar, 
			author_username, location, coord_lat, coord_lon, following, followers, language,
			num_replies, num_rts, published_ts, status_type) VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
)

type FeedRepository interface {
	CreateFeed(ctx context.Context, feed model.Feed) error
}

// Implementation of Repository
type feedRepository struct {
	my *config.MysqlSession
	so *solr.SolrInterface
	// next .ChainRepository
}

// Constructor of vaRepository
func NewFeedRepository(db *config.MysqlSession, solr *solr.SolrInterface) FeedRepository {
	return feedRepository{
		my: db,
		so: solr,
	}
}

func (r feedRepository) insertRecord(tx *sql.Tx, param []interface{}) error {
	stmt, err := tx.Prepare(queryInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(param...); err != nil {
		return err
	}
	return nil
}

func (r feedRepository) insertLog(doc []solr.Document) error {
	r.so.SetCore(model.FeedCollectionName)
	params := url.Values{}

	if _, err := r.so.Add(doc, len(doc), &params); err != nil {
		return err
	}

	return nil
}

// Create feed
func (r feedRepository) CreateFeed(ctx context.Context, feed model.Feed) error {
	tx, err := r.my.Db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	fmt.Printf("insert data %v", feed.FeedId)
	if err := r.insertRecord(tx, feed.ToInsertArgs()); err != nil {
		fmt.Printf("insert feed failed! err: %v", err)
		return err
	}

	var docs []solr.Document
	if err := r.insertLog(feed.ToInsertDocument(docs)); err != nil {
		fmt.Printf("insert log failed! err: %v", err)
		return err
	}

	return tx.Commit()
}
