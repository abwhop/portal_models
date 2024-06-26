package models

type BlogGQLRespond struct {
	Data struct {
		Blogs []*BlogAPI `json:"blogs"`
	} `json:"data"`
}

type BlogAPI struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	DateCreate  int        `json:"date_create"`
	Author      *UserAPI   `json:"author"`
	Subscribers []*UserAPI `json:"subscribers"`
}

type BlogDB struct {
	BitrixId         int            `gorm:"column:id;primaryKey" json:"id"`
	Name             string         `gorm:"column:name;" json:"name"`
	Description      string         `gorm:"column:description;" json:"description"`
	DateCreated      int            `gorm:"column:dateCreated;" json:"dateCreated"`
	Author           *UserDB        `gorm:"column:author;" json:"author"`
	Subscribers      *ListOfUsersDB `gorm:"column:subscribers;" json:"subscribers"`
	SubscribersCount int            `gorm:"column:subscribersCount;" json:"subscribersCount"`
	IsFavorite       bool           `gorm:"column:isFavorite;" json:"isFavorite"`
	IsSubscription   bool           `gorm:"column:isSubscription;" json:"isSubscription"`
	IsSubscriber     bool           `gorm:"column:isSubscriber;" json:"isSubscriber"`
	CountCommunities int            `gorm:"column:countCommunities;" json:"countCommunities"`
}

func (BlogDB) TableName() string {
	return "public.bitrix_blogs"
}
