package models

import "github.com/lib/pq"

type FeedDB struct {
	Type          string          `gorm:"column:type;primaryKey" json:"type"`
	Id            int             `gorm:"column:id;primaryKey" json:"id"`
	IsFavorite    bool            `gorm:"column:isFavorite" json:"isFavorite"`
	IsRubric      bool            `gorm:"column:isRubric" json:"isRubric"`
	Title         string          `gorm:"column:title" json:"title"`
	Description   string          `gorm:"column:description" json:"description"`
	ImageUrl      string          `gorm:"column:imageUrl" json:"imageUrl"`
	CanComment    bool            `gorm:"column:canComment" json:"canComment"`
	IsLiked       bool            `gorm:"column:isLiked" json:"isLiked"`
	CanLikes      bool            `gorm:"column:canLikes" json:"canLikes"`
	CountComments int             `gorm:"column:countComments" json:"countComments"`
	CountViews    int             `gorm:"column:countViews" json:"countViews"`
	CountLikes    int             `gorm:"column:countLikes" json:"countLikes"`
	DateCreated   int             `gorm:"column:date_created" json:"dateCreated"`
	DateUpdated   int             `gorm:"column:date_updated" json:"dateUpdated"`
	LogId         int             `gorm:"column:logId" json:"logId"`
	Rubric        *RubricDB       `gorm:"column:rubric" json:"rubric"`
	Data          *RepostedPostDB `gorm:"column:data"  json:"data"`
	Tags          pq.StringArray  `gorm:"column:tags;type:varchar[]" json:"tags"`
}

func (FeedDB) TableName() string {
	return "public.bitrix_feed"
}
