package model

import (
	"database/sql"

	"github.com/vanng822/go-solr/solr"
)

const (
	FeedTableName      = "twitter_feed"
	FeedCollectionName = "twitter_feed"
)

type Feed struct {
	Id             int          `field:"id"`
	FeedId         string       `field:"feed_id" json:"id"`
	PublisedDate   sql.NullTime `field:"published_date" json:"postedTime"`
	Link           string       `field:"link" json:"link"`
	Content        string       `field:"content" json:"body"`
	AuthorName     string       `field:"author_name" json:"actor.displayName"`
	AuthorLink     string       `field:"author_link" json:"actor.link"`
	AuthorAvatar   string       `field:"author_avatar" json:"actor.image"`
	AuthorUsername string       `field:"author_username" json:"actor.prefferedUsername"`
	Location       string       `field:"location" json:"actor.location.displayName"`
	CoordLat       float64      `field:"coord_lat" json:"coord_lat"`
	CoordLon       float64      `field:"coord_lon" json:"coord_lon"`
	Following      int          `field:"following" json:"actor.friendsCount"`
	Followers      int          `field:"followers" json:"actor.followersCount"`
	Language       string       `field:"language" json:"twitter_lang"`
	NumReplies     int          `field:"num_replies" json:"num_replies"`
	NumRts         int          `field:"num_rts" json:"retweet_count"`
	PublishedTs    int          `field:"published_ts" json:"in_reply_to"`
	StatusType     int          `field:"status_type" json:"status_type"`
}

// to Args slice interface for insert feed
func (feed Feed) ToInsertArgs() []interface{} {
	var arg []interface{}
	arg = append(arg, feed.FeedId)
	arg = append(arg, feed.PublisedDate)
	arg = append(arg, feed.Link)
	arg = append(arg, feed.Content)
	arg = append(arg, feed.AuthorName)
	arg = append(arg, feed.AuthorLink)
	arg = append(arg, feed.AuthorAvatar)
	arg = append(arg, feed.AuthorUsername)
	arg = append(arg, feed.Location)
	arg = append(arg, feed.CoordLat)
	arg = append(arg, feed.CoordLon)
	arg = append(arg, feed.Following)
	arg = append(arg, feed.Followers)
	arg = append(arg, feed.Language)
	arg = append(arg, feed.NumReplies)
	arg = append(arg, feed.NumRts)
	arg = append(arg, feed.PublishedTs)
	arg = append(arg, feed.StatusType)
	return arg
}

// to Docu slice interface for insert log
func (feed Feed) ToInsertDocument(docs []solr.Document) []solr.Document {

	docs = append(docs, solr.Document{
		"feed_id_i":         feed.FeedId,
		"published_date_dt": feed.PublisedDate,
		"link_s":            feed.Link,
		"content_t":         feed.Content,
		"author_name_s":     feed.AuthorName,
		"author_link_s":     feed.AuthorLink,
		"author_avatar_s":   feed.AuthorAvatar,
		"author_username_s": feed.AuthorUsername,
		"location_s":        feed.Location,
		"coord_lat_f":       feed.CoordLat,
		"coord_lon_f":       feed.CoordLon,
		"following_l":       feed.Following,
		"followers_l":       feed.Followers,
		"language_s":        feed.Language,
		"num_replies_i":     feed.NumReplies,
		"num_rts":           feed.NumRts,
		"published_ts_l":    feed.PublishedTs,
		"status_type_i":     feed.StatusType,
	})
	return docs
}
