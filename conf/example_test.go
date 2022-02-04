package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/leong-y/go-gin-example/models"
	"github.com/leong-y/go-gin-example/pkg/gredis"
	"github.com/leong-y/go-gin-example/pkg/setting"
	"github.com/leong-y/go-gin-example/service/article_service"
)

func TestSetup(t *testing.T) {
	setting.Setup()
	models.Setup()
	gredis.Setup()
	articleService := article_service.Article{ID: 1}
	exists, err := articleService.ExistByID()
	if err != nil {
		log.Fatalln(err)
		return
	}
	if !exists {
		log.Fatalln("没有这个ID")
		return
	}
	article, err := articleService.Get()
	if err != nil {
		log.Fatalln("文章获取失败")
		return
	}
	fmt.Printf("article: %v\n", article)
}
