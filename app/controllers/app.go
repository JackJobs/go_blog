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
