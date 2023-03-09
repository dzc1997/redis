package model

import (
	"context"
	"gitee.com/wedone/redis_course/redis/article/common"
	"github.com/go-redis/redis/v9"
	"log"
	"strconv"
	"strings"
	"time"
)

type Article interface {
	ArticleVote(string, string)
	PostArticle(string, string, string) string
	GetArticles(int64, string) []map[string]string
	AddRemoveGroups(string, []string, []string)
	GetGroupArticles(string, string, int64) []map[string]string
	Reset()
}

type ArticleRepo struct {
	Conn *redis.Client
}

func NewArticleRepo(conn *redis.Client) *ArticleRepo {
	return &ArticleRepo{Conn: conn}
}

func (r *ArticleRepo) ArticleVote(ctx context.Context, article, user string) {
	cutoff := time.Now().Unix() - common.OneWeekInSeconds
	//如果发布时间超过了一个礼拜直接返回不做任何操作
	if r.Conn.ZScore(ctx, "time", article).Val() < float64(cutoff) {
		return
	}
	//获取id
	articleId := strings.Split(article, ":")[1]
	//添加一个作者
	if r.Conn.SAdd(ctx, "voted:"+articleId, user).Val() != 0 {
		//对有序集合score中的value的score增加一个增量(value为article)
		r.Conn.ZIncrBy(ctx, "score", common.VoteScore, article)
		//为键名为article的hash中key为votes的value增加1
		r.Conn.HIncrByFloat(ctx, article, "votes", 1)
	}
}

// PostArticle 发布并获取文章
func (r *ArticleRepo) PostArticle(ctx context.Context, user, title, link string) string {
	//使用string记录id号(键名为article:,键值为文章数量,即目前新加入文章的id号)
	articleId := strconv.Itoa(int(r.Conn.Incr(ctx, "article:").Val()))
	voted := "voted:" + articleId
	//将文章发布者的ID放入set中,每个文章有一个set(键名为voted:id,value为发布者名称)
	r.Conn.SAdd(ctx, voted, user)
	r.Conn.Expire(ctx, voted, common.OneWeekInSeconds*time.Second)
	now := time.Now().Unix()
	article := "article:" + articleId
	//使用hash记录文章信息,每个文章有一个hash(键名为article:id,key为文章属性,value为具体值)
	r.Conn.HMSet(ctx, article, map[string]interface{}{
		"title":  title,
		"link":   link,
		"poster": user,
		"time":   now,
		"votes":  1,
	})
	//使用zset记录文章评分(键名为time,value为article:id,score当前时间戳+VoteScore)
	r.Conn.ZAdd(ctx, "score:", redis.Z{Score: float64(now + common.VoteScore), Member: article})
	//使用zset记录文章时间戳(键名为time,value为article:id,score为当前时间戳)
	r.Conn.ZAdd(ctx, "time", redis.Z{Score: float64(now), Member: article})
	//返回文章id
	return articleId
}

// GetArticles 获取文章 获取最新发布的一页文章 25篇
func (r *ArticleRepo) GetArticles(ctx context.Context, page int64, order string) []map[string]string {
	//order为排列顺序
	if order == "" {
		order = "score:"
	}
	//指定一页为25篇文章,start为每页的起始索引
	start := (page - 1) * common.ArticlesPerPage // 获取文章的起始索引
	//end为每页的的结束索引
	end := start + common.ArticlesPerPage - 1 // 获取文章的结束索引
	//从键名为scores:的zset获取start到end文章的id,并且按score:从大到小排列,即最新文章在前面
	ids := r.Conn.ZRevRange(ctx, order, start, end).Val() // 获取多个文章ID ID格式为article:id
	var articles []map[string]string                      // 根据文章的ID获取文章的详细信息
	for _, id := range ids {
		articleData := r.Conn.HGetAll(ctx, id).Val()
		articleData["id"] = id
		articles = append(articles, articleData)
	}
	return articles
}

// AddRemoveGroups 对文章进行分组
func (r *ArticleRepo) AddRemoveGroups(ctx context.Context, articleId string, toAdd, toRemove []string) {
	article := "article:" + articleId // 构建存储文章信息的value
	for _, group := range toAdd {     // 将文章添加到它所属的群组里面
		//将article添加到键名为group:group的集合中
		r.Conn.SAdd(ctx, "group:"+group, article)
	}
	//将article从键名为group:group的集合中移除
	for _, group := range toRemove { // 从群组里面移除文章
		r.Conn.SRem(ctx, "group:"+group, article)
	}
}

// GetGroupArticles 从群组里面获取整页文章
// 从群组group获取某页文章,按照order排序
func (r *ArticleRepo) GetGroupArticles(ctx context.Context, group, order string, page int64) []map[string]string {
	//order为排列顺序
	if order == "" {
		order = "score:"
	}
	//score:new-group
	key := order + group                    // 为每个群组的每种排列顺序都创建一个键
	if r.Conn.Exists(ctx, key).Val() == 0 { // 检查是否有已缓存的排序结果，如果没有的话就现在进行排序
		//计算order与group:group的交集,如果res小于0说明没有交集,无法进行排序并获取文章
		//Aggregate:"MAX"
		/*
			SUM：将所有集合中某一个元素的score值之和作为结果集中该成员的score值
			MIN：将所有集合中某一个元素的score值中最小值作为结果集中该成员的score值
			MAX：将所有集合中某一个元素score值中最大值作为结果集中该成员的score值
		*/
		res := r.Conn.ZInterStore(ctx, key, &redis.ZStore{Aggregate: "MAX", Keys: []string{"group:" + group, order}}).Val()
		if res <= 0 {
			log.Println("ZInterStore")
		}
	}
	r.Conn.Expire(ctx, key, 60*time.Second)
	return r.GetArticles(ctx, page, key)
}

// Reset 清空
func (r *ArticleRepo) Reset(ctx context.Context) {
	r.Conn.FlushDB(ctx)
}
