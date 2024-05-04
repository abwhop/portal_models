package models

import "github.com/lib/pq"

type NewsGQLRespond struct {
	Data struct {
		News []*NewsAPI `json:"news"`
	} `json:"data"`
}

type NewsAPI struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	LogId       int      `json:"log_id"`
	CanComment  bool     `json:"can_comment"`
	Text        string   `json:"text"`
	CreateDate  int      `json:"create_date"`
	PublishDate int      `json:"publish_date"`
	Published   bool     `json:"published"`
	Rights      []string `json:"rights"`
	Img         string   `json:"img"`
	PreviewText string   `json:"preview_text"`
	//XmlId       string        `json:"xml_id"`
	SliderFile string        `json:"slider_file"`
	Rubric     *RubricAPI    `json:"rubric"`
	Author     *UserAPI      `json:"author"`
	Likes      *LikesAPI     `json:"likes"`
	Views      *ViewsAPI     `json:"views"`
	Comments   []*CommentAPI `json:"comments"`
	RepostBlog *PostAPI      `json:"repost_blog"`
	Files      []*FileAPI    `json:"files"`
	VoteNum    interface{}   `json:"vote_num"`
	Tags       []*Tag        `json:"tags"`
}

type NewsBrief struct {
	Id                 int                  `gorm:"column:id;primaryKey" json:"id"`
	Type               string               `gorm:"column:type;primaryKey" json:"type"`
	PublishDate        int                  `gorm:"column:publish_date" json:"publish_date"`
	CreateDate         int                  `gorm:"column:date_created" json:"date_created"`
	Title              string               `gorm:"column:title" json:"title"`
	LogId              int                  `gorm:"column:log_id" json:"log_id"`
	CanComment         bool                 `gorm:"column:can_comment" json:"can_comment"`
	Descriptions       *ListOfDescriptionDB `gorm:"column:descriptions" json:"descriptions"`
	Description        string               `gorm:"column:description" json:"description"`
	Published          bool                 `gorm:"column:published" json:"published"`
	Rights             pq.StringArray       `gorm:"column:rights;type:varchar[]" json:"rights"`
	ImageUrl           string               `gorm:"column:image_url" json:"image_url"`
	PreviewDescription string               `gorm:"column:preview_description" json:"previewDescription"`
	//XmlId              string               `gorm:"column:xml_id" json:"xml_id"`
	Rubric          *RubricDB        `gorm:"column:rubric" json:"rubric"`
	Author          *UserDB          `gorm:"column:author" json:"author"`
	Likes           *LikesDB         `gorm:"column:likes" json:"likes"`
	Views           *ViewsDB         `gorm:"column:views" json:"views"`
	ReposBlogPostId int              `gorm:"column:repost_blog_post_id" json:"repost_blog_post_id"`
	Comments        *ListOfCommentDB `gorm:"column:comments" json:"comments"`
	Files           *ListOfFileDB    `gorm:"column:files" json:"files"`
	VoteNum         pq.Int64Array    `gorm:"column:vote_num;type:int[]" json:"vote_num"`
	FormId          pq.Int64Array    `gorm:"column:form_id;type:int[]" json:"form_id"`
	Tags            pq.StringArray   `gorm:"column:tags;type:varchar[]" json:"tags"`
}

func (NewsBrief) TableName() string {
	return "public.bitrix_news"
}

type NewsDetail struct {
	Id            int       `gorm:"column:id;primaryKey" json:"id"`
	Type          string    `gorm:"column:type;primaryKey" json:"type"`
	Title         string    `gorm:"column:title" json:"title"`
	ImageUrl      string    `gorm:"column:imageUrl" json:"imageUrl"`
	DateCreated   int64     `gorm:"column:dateCreated" json:"dateCreated"`
	DateUpdated   int64     `gorm:"column:dateUpdated" json:"dateUpdated"`
	Rubric        *RubricDB `gorm:"column:rubric" json:"rubric"`
	LogId         int       `gorm:"column:logId" json:"logId"`
	CountLikes    int       `gorm:"column:countLikes" json:"countLikes"`
	CountViews    int       `gorm:"column:countViews" json:"countViews"`
	CanComment    bool      `gorm:"column:canComment" json:"canComment"`
	CountComments int       `gorm:"column:countComments" json:"countComments"`
	IsView        bool      `gorm:"column:isView" json:"isView"`
	IsFavorite    bool      `gorm:"column:isFavorite" json:"isFavorite"`
	IsRubric      bool      `gorm:"column:isRubric" json:"isRubric"`
	//Data          interface{}          `gorm:"column:data" json:"data"`
	//Votes         interface{}          `gorm:"column:votes" json:"votes"`
	//Forms         interface{}          `gorm:"column:forms" json:"forms"`
	Descriptions  *ListOfDescriptionDB `gorm:"column:descriptions" json:"descriptions"`
	Comments      *ListOfCommentDB     `gorm:"column:comments" json:"comments"`
	Files         *ListOfFileDB        `gorm:"column:files" json:"files"`
	CalendarEvent string               `gorm:"column:calendarEvent" json:"calendarEvent"`
	Description   string               `gorm:"column:description" json:"description"`
	IsLiked       bool                 `gorm:"column:isLiked" json:"isLiked"`
}

func (NewsDetail) TableName() string {
	return "public.bitrix_news"
}
