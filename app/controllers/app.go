package controllers

import (
	"MyTest/app/models"
	"github.com/revel/revel"
	"time"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	println("index")
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	//读取所有的博客文章
	blogs := dao.FindBlogs()
	now := time.Now().Add(-1 * time.Hour)
	recentCnt := dao.FindBlogsByDate(now)
	return c.Render(blogs, recentCnt)
}

func (c App) WBlog() revel.Result {
	return c.Render()
}

func (c App) BlogInfor(id string, rcnt int) revel.Result {
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	blog := dao.FindBlogById(id)
	if blog.ReadCnt == rcnt {
		blog.ReadCnt = rcnt + 1
		dao.UpdateBlogById(id, blog)
	}
	comments := dao.FindCommentsByBlogId(blog.Id)
	if len(comments) == 0 && blog.CommentCnt != 0 {
		blog.CommentCnt = 0
		dao.UpdateBlogById(id, blog)
	} else if len(comments) != blog.CommentCnt {
		blog.CommentCnt = len(comments)
		dao.UpdateBlogById(id, blog)
	}
	return c.Render(blog, rcnt, comments)
}

func (c App) Message() revel.Result {
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	//dao := models.NewDao(c.MongoSession)
	messages := dao.FindAllMessages()
	return c.Render(messages)
	return c.Render()
}

func (c App) History() revel.Result {
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	dao.CreateAllHistory()
	histories := dao.FindHistory()
	for i, _ := range histories {
		histories[i].Blogs = dao.FindBlogsByYear(histories[i].Year)
	}
	return c.Render(histories)
}
