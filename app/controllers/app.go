package controllers

import (
	"github.com/revel/revel"
	"MyTest/app/models"
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
	return c.Render(blogs)
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
		dao.UpdateBlogById(id ,blog)
	} else if len(comments) != blog.CommentCnt {
		blog.CommentCnt = len(comments)
		dao.UpdateBlogById(id, blog)
	}
	return c.Render(blog, rcnt, comments)
}





















