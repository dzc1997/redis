package main

import (
	"context"
	"gitee.com/wedone/redis_course/redis/article/model"
	"gitee.com/wedone/redis_course/redis/redisConn"
	"testing"
)

var ctx = context.Background()

func Test(t *testing.T) {
	conn := redisConn.ConnectRedis()
	client := model.NewArticleRepo(conn)

	articleId := client.PostArticle(ctx, "username", "A title", "http://www.google.com")
	t.Log("We posted a new article with id:", articleId)
	assertStringResult(t, "1", articleId)
	//从指定hash中获取文章的信息
	r := client.Conn.HGetAll(ctx, "article:"+articleId).Val()
	t.Log("\nIts HASH looks like:", r)
	assertTrue(t, len(r) != 0)
	//
	client.ArticleVote(ctx, "article:"+articleId, "other_user")
	v, _ := client.Conn.HGet(ctx, "article:"+articleId, "votes").Int()
	t.Log("\nWe voted for the article,it now has votes:", v)
	assertTrue(t, v >= 1)
	t.Log("\nThe currently highest-scoring articles are:")
	//获取
	articles := client.GetArticles(ctx, 1, "")
	assertTrue(t, len(articles) >= 1)
	for k, v := range articles {
		t.Log(k, v)
	}
	//将文章分组
	client.AddRemoveGroups(ctx, articleId, []string{"new-group"}, []string{})
	articles = client.GetGroupArticles(ctx, "new-group", "score:", 1)
	t.Log("\nWe added the article to a new group, other article include:")
	assertTrue(t, len(articles) >= 1)
	for k, v := range articles {
		t.Log(k, v)
	}
	defer client.Reset(ctx)
}

func assertStringResult(t *testing.T, want, get string) {
	t.Helper()
	if want != get {
		t.Errorf("want get %v,actual get %v\n", want, get)
	}
}

func assertTrue(t *testing.T, v bool) {
	t.Helper()
	if v != true {
		t.Error("assert true but get a false value")
	}
}
