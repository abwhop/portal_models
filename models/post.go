package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostGQLRespond struct {
	Data struct {
		BlogPosts []*PostAPI `json:"blog_posts"`
	} `json:"data"`
}

type PostAPI struct {
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	VoteNum     interface{}   `json:"vote_num"`
	BlogId      int           `json:"blog_id"`
	Text        string        `json:"text"`
	IsDraft     bool          `json:"is_draft"`
	CreateDate  int64         `json:"create_date"`
	PublishDate int64         `json:"publish_date"`
	Img         string        `json:"img"`
	Rights      []string      `json:"rights"`
	PostRights  []string      `json:"post_rights"`
	Files       []*FileAPI    `json:"files"`
	Author      *UserAPI      `json:"author"`
	Likes       *LikesAPI     `json:"likes"`
	Views       *ViewsAPI     `json:"views"`
	Comments    []*CommentAPI `json:"comments"`
	RepostBlog  *PostAPI      `json:"repost_blog"`
	RepostNews  *NewsAPI      `json:"repost_news"`
	Vote        interface{}   `json:"vote"`
}

/*type PostDB struct {
	BitrixId           int                  `json:"bitrix_id"`
	Type               string               `json:"type"`
	Title              string               `json:"title"`
	Description        string               `json:"description"`
	ImageUrl           string               `json:"image_url"`
	AuthorId           int                  `json:"author_id"`
	Comments           *ListOfCommentDB     `json:"comments"`
	Descriptions       *ListOfDescriptionDB `json:"descriptions"`
	PreviewDescription string               `json:"preview_description"`
	Likes              *LikesDB             `json:"likes"`
	DateUpdated        int                  `json:"date_updated"`
	RubricId           int                  `json:"rubric_id"`
	Published          bool                 `json:"published"`
	Rights             []string             `json:"rights"`
	PublishDate        int                  `json:"publish_date"`
	Views              *ViewsDB             `json:"views"`
	RepostBlogPostId   int                  `json:"repost_blog_post_id"`
	LogId              int                  `json:"log_id"`
	IsDeleted          bool                 `json:"is_deleted"`
	CanComment         bool                 `json:"can_comment"`
	Files              []*FileDB            `json:"files"`
	VoteId             int                  `json:"vote_id"`
	FormId             int                  `json:"form_id"`
}*/

type PostDB struct {
	Id               int                  `gorm:"column:id;primaryKey" json:"id"`
	Type             string               `gorm:"column:type;" json:"type"`
	Title            string               `gorm:"column:title;" json:"title"`
	Text             string               `gorm:"column:text;" json:"text"`
	CreatedDate      int64                `gorm:"column:createDate;" json:"createDate"`
	PublishDate      int64                `gorm:"column:publishDate;" json:"publishDate"`
	Img              string               `gorm:"column:img;" json:"img"`
	Rights           pq.StringArray       `gorm:"column:rights;type:varchar[]" json:"rights"`
	Files            *ListOfFileDB        `gorm:"column:files;" json:"files"`
	Author           *UserDB              `gorm:"column:author;" json:"author"`
	Likes            *LikesDB             `gorm:"column:likes;" json:"likes"`
	Views            *ViewsDB             `gorm:"column:views;" json:"views"`
	Comments         *ListOfCommentDB     `gorm:"column:comments;" json:"comments"`
	Descriptions     *ListOfDescriptionDB `gorm:"column:descriptions;" json:"descriptions"`
	BlogId           int                  `gorm:"column:blogId;" json:"blogId"`
	RepostBlogPostId int                  `gorm:"column:repostBlogPostId;" json:"repostBlogPostId"`
	RepostNewsId     int                  `gorm:"column:repostNewsId;" json:"repostNewsId"`
	IsDraft          bool                 `gorm:"column:isDraft;" json:"isDraft"`
	LastUpdateDate   int64                `gorm:"column:lastUpdateDate;" json:"lastUpdateDate"`
	VoteNum          pq.Int64Array        `gorm:"column:voteNum;type:int[]" json:"voteNum"`
	IsDeleted        bool                 `gorm:"column:isDeleted;" json:"isDeleted"`
	PostRights       pq.StringArray       `gorm:"column:postRights;type:varchar[]" json:"postRights"`
	FormId           pq.Int64Array        `gorm:"column:formId;type:int[]" json:"formId"`
	CommentsCount    int                  `gorm:"column:commentsCount;" json:"commentsCount"`
	Data             *PostOrNews          `gorm:"column:data;" json:"data"`
}

func (PostDB) TableName() string {
	return "public.bitrix_posts"
}

type RepostedPostDB struct {
	Id               int              `gorm:"column:id;primaryKey" json:"id"`
	Type             string           `gorm:"column:type;" json:"type"`
	Title            string           `gorm:"column:title;" json:"title"`
	Text             string           `gorm:"column:text;" json:"text"`
	CreatedDate      int64            `gorm:"column:createDate;" json:"createDate"`
	PublishDate      int64            `gorm:"column:publishDate;" json:"publishDate"`
	Img              string           `gorm:"column:img;" json:"img"`
	Rights           []string         `gorm:"column:rights;type:varchar[]" json:"rights"`
	Files            []*FileDB        `gorm:"column:files;" json:"files"`
	Author           *UserDB          `gorm:"column:author;" json:"author"`
	Likes            *LikesDB         `gorm:"column:likes;" json:"likes"`
	Views            *ViewsDB         `gorm:"column:views;" json:"views"`
	Comments         []*CommentDB     `gorm:"column:comments;" json:"comments"`
	Descriptions     []*DescriptionDB `gorm:"column:descriptions;" json:"descriptions"`
	BlogId           int              `gorm:"column:blogId;" json:"blogId"`
	RepostBlogPostId int              `gorm:"column:repostBlogPostId;" json:"repostBlogPostId"`
	RepostNewsId     int              `gorm:"column:repostNewsId;" json:"repostNewsId"`
	IsDraft          bool             `gorm:"column:isDraft;" json:"isDraft"`
	LastUpdateDate   int64            `gorm:"column:lastUpdateDate;" json:"lastUpdateDate"`
	VoteNum          []int            `gorm:"column:voteNum;type:int[]" json:"voteNum"`
	IsDeleted        bool             `gorm:"column:isDeleted;" json:"isDeleted"`
	PostRights       []string         `gorm:"column:postRights;type:varchar[]" json:"postRights"`
	FormId           []int            `gorm:"column:formId;type:int[]" json:"formId"`
	CommentsCount    int              `gorm:"column:commentsCount;" json:"commentsCount"`
	Data             *PostOrNews      `gorm:"column:data;" json:"data"`
}

func (u RepostedPostDB) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	if item, err := json.Marshal(u); err == nil {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{string(item)},
		}
	} else {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{nil},
		}
	}
}

func (u RepostedPostDB) Value() (driver.Value, error) {
	var item []byte
	var err error
	if item, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(item), nil
}
func (u RepostedPostDB) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), &u)
}

type Post struct {
	Id            int              `gorm:"column:id;primaryKey" json:"id"`
	Title         string           `gorm:"column:title;" json:"title"`
	ImageUrl      string           `gorm:"column:imageUrl;" json:"imageUrl"`
	CreatedDate   int64            `gorm:"column:dateCreated;" json:"dateCreated"`
	Files         *ListOfFileDB    `gorm:"column:files;" json:"files"`
	Author        *UserDB          `gorm:"column:author;" json:"author"`
	Description   string           `gorm:"column:description;" json:"description"`
	CountLikes    int              `gorm:"column:countLikes;" json:"countLikes"`
	CountViews    int              `gorm:"column:countViews;" json:"countViews"`
	IsDraft       bool             `gorm:"column:isDraft;" json:"isDraft"`
	CountComments int              `gorm:"column:countComments;" json:"countComments"`
	IsView        bool             `gorm:"column:isView;" json:"isView"`
	IsFavorite    bool             `gorm:"column:isFavorite;" json:"isFavorite"`
	IsLiked       bool             `gorm:"column:isLiked;" json:"isLiked"`
	Type          string           `gorm:"column:type;" json:"type"`
	Data          *PostOrNews      `gorm:"column:data;" json:"data"`
	Comments      *ListOfCommentDB `gorm:"column:comments;" json:"comments"`
}

func (Post) TableName() string {
	return "public.bitrix_posts"
}

type PostOrNews struct {
	Type          string        `gorm:"column:type;" json:"type"`
	Id            int           `gorm:"column:id;primaryKey" json:"id"`
	IsFavorite    bool          `gorm:"column:isFavorite;" json:"isFavorite"`
	IsRubric      bool          `gorm:"column:isRubric" json:"isRubric"`
	Title         string        `gorm:"column:title;" json:"title"`
	Description   string        `gorm:"column:description;" json:"description"`
	ImageUrl      string        `gorm:"column:imageUrl;" json:"imageUrl"`
	CanComment    bool          `gorm:"column:canComment" json:"canComment"`
	IsLiked       bool          `gorm:"column:isLiked;" json:"isLiked"`
	CanLikes      bool          `gorm:"column:canLikes" json:"canLikes"`
	CountComments int           `gorm:"column:countComments;" json:"countComments"`
	CountViews    int           `gorm:"column:countViews;" json:"countViews"`
	CountLikes    int           `gorm:"column:countLikes;" json:"countLikes"`
	DateCreated   int64         `gorm:"column:dateCreated;" json:"dateCreated"`
	DateUpdated   int           `gorm:"column:date_updated" json:"dateUpdated"`
	Files         *ListOfFileDB `gorm:"column:files;" json:"files"`
	Author        *UserDB       `gorm:"column:author;" json:"author"`
	IsDraft       bool          `gorm:"column:isDraft;" json:"isDraft"`
	IsView        bool          `gorm:"column:isView;" json:"isView"`

	Data     *PostOrNews      `gorm:"column:data;" json:"data"`
	Comments *ListOfCommentDB `gorm:"column:comments;" json:"comments"`
}

func (u PostOrNews) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

// Scan implements the sql.Scanner interface
func (u *PostOrNews) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), &u)
}
