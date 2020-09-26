package model

import "github.com/go-programming-tour-book/blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

// swagger使用的结构体
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
