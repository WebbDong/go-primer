package main

import (
	"fmt"
	"reflect"
	"time"
)

var startWeiboId = 1

// 用户
type User struct {
	name     string
	age      int
	birthday time.Time
}

// 博主接口
type BloggerInterface interface {
	// 关注
	Follow(fan FansInterface)
	// 取消关注
	UnFollow(fan FansInterface)
	// 发布微博
	PostWeibo(topic string, content string)
	// 添加评论
	AddComment(weiboId int, newComment Comment)
	// 打印所有微博
	PrintAllWeibo()
	// 打印所有粉丝
	PrintAllFans()
	// 打印所有评论
	PrintAllComments()
}

// 博主
type Blogger struct {
	User
	microblogs []*Weibo
	// 评论，key为微博id，value为评论切片
	commentsMap map[int][]*Comment
	fans        []FansInterface
}

func (b *Blogger) PostWeibo(topic string, content string) {
	newBlog := new(Weibo)
	newBlog.id = startWeiboId
	newBlog.topic = topic
	newBlog.content = content
	newBlog.types = 1
	newBlog.postMan = b.name
	newBlog.postTime = time.Now()
	b.microblogs = append(b.microblogs, newBlog)
	startWeiboId++
}

func (b *Blogger) Follow(fan FansInterface) {
	b.fans = append(b.fans, fan)
}

func (b *Blogger) UnFollow(fan FansInterface) {
	for i, length := 0, len(b.fans); i < length; i++ {
		f := b.fans[i]
		if f.GetName() == fan.GetName() {
			b.fans = append(b.fans[:i:i], b.fans[i+1:]...)
			break
		}
	}
}

func (b *Blogger) AddComment(weiboId int, newComment Comment) {
	if reflect.ValueOf(b.commentsMap).IsNil() {
		b.commentsMap = map[int][]*Comment{}
	}
	comments := b.commentsMap[weiboId]
	comments = append(comments, &newComment)
	b.commentsMap[weiboId] = comments
}

func (b *Blogger) PrintAllWeibo() {
	for _, v := range b.microblogs {
		fmt.Println(*v)
	}
}

func (b *Blogger) PrintAllFans() {
	for _, v := range b.fans {
		fmt.Println(v.GetName())
	}
}

func (b *Blogger) PrintAllComments() {
	fmt.Println(b.commentsMap)
	for _, v := range b.commentsMap {
		for _, v := range v {
			fmt.Println(v.content)
		}
	}
}

// 粉丝接口
type FansInterface interface {
	// 发布评论
	PostComment(weiboId int, bloggerInter BloggerInterface, commentContent string)
	GetName() string
}

// 粉丝
type Fans struct {
	User
}

func (f *Fans) GetName() string {
	return f.name
}

func (a *Fans) PostComment(weiboId int, bloggerInter BloggerInterface, commentContent string) {
	newComment := Comment{id: 1, commentUser: a.name, content: commentContent, commentTime: time.Now()}
	bloggerInter.AddComment(weiboId, newComment)
}

// 黑粉
type AntiFans struct {
	Fans
}

// 红粉
type Follower struct {
	Fans
}

// 微博
type Weibo struct {
	id       int
	topic    string
	content  string
	postTime time.Time
	types    int
	postMan  string
}

// 评论
type Comment struct {
	id int
	// 评论人
	commentUser string
	content     string
	commentTime time.Time
	// 评论给谁
	toUser string
}

func main() {
	var b BloggerInterface = &Blogger{User: User{"科比", 38, time.Now()}}
	b.PostWeibo("火箭总冠军", "哈登62分")
	b.PostWeibo("勇士总冠军", "库里52分")
	b.PrintAllWeibo()

	var f1 FansInterface = &Follower{Fans{User{"dwb", 28, time.Now()}}}
	var f2 FansInterface = &AntiFans{Fans{User{"gcc", 30, time.Now()}}}

	b.Follow(f1)
	b.Follow(f2)
	b.PrintAllFans()
	fmt.Println("----------------")

	var f3 = &AntiFans{Fans{User{"fer", 35, time.Now()}}}
	b.UnFollow(f1)
	b.Follow(f3)
	b.PrintAllFans()
	fmt.Println("----------------")

	f1.PostComment(1, b, "差评")
	f1.PostComment(1, b, "好评")
	b.PrintAllComments()
}
