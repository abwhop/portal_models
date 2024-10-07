package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentAPI struct {
	Id         int       `json:"id"`
	Text       string    `json:"text"`
	SourceId   int       `json:"source_id"`
	DateCreate int64     `json:"date_create"`
	Author     *UserAPI  `json:"author"`
	Likes      *LikesAPI `json:"likes"`
}

type CommentDB struct {
	Id          int      `json:"id"`
	Text        string   `json:"text"`
	SourceId    int      `json:"source_id"`
	DateCreated int64    `json:"dateCreated"`
	Author      *UserDB  `json:"author"`
	Likes       *LikesDB `json:"likes"`
}

type CommentDetail struct {
	Id         int       `json:"id"`
	Text       string    `json:"text"`
	SourceId   int       `json:"source_id"`
	DateCreate int64     `json:"date_create"`
	Author     *UserAPI  `json:"author"`
	IsLiked    bool      `json:"is_liked"`
	Likes      *LikesAPI `json:"likes"`
}

type ListOfCommentDB []*CommentDB

func (u ListOfCommentDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if user, err := json.Marshal(u); err == nil {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{string(user)},
		}
	} else {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{nil},
		}
	}
}

func (u ListOfCommentDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}
func (u *ListOfCommentDB) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), &u)
}
func (u CommentDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if user, err := json.Marshal(u); err == nil {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{string(user)},
		}
	} else {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{nil},
		}
	}
}

func (u CommentDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}
func (u *CommentDB) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), &u)
}
