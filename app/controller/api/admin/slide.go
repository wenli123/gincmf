/**
* @Title  轮播图资源控制器
* @Description  轮播图增删改查接口
* @Author  Return  20200519 10：46
* @Update
**/
package admin

import (
	"fmt"
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf"
	"github.com/gincmf/cmf/controller"
	"strconv"
)

// CarouselController 轮播图资源控制器，定义了轮播图增删改查接口
type SlideController struct {
	rc controller.RestController
}

func (i *SlideController) Get(c *gin.Context) {

	slide := []model.Slide{}
	notFount := cmf.Db.Find(&slide).RecordNotFound()
	if notFount {
		i.rc.Error(c,"该内容不存在！")
	}

	i.rc.Success(c, "carousel操作成功Get", slide)
}

func (i *SlideController) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	fmt.Println("id",rewrite.Id)

	i.rc.Success(c, "操作成功show", rewrite.Id)
}

func (i *SlideController) Edit(c *gin.Context) {
	i.rc.Success(c, "操作成功Edit", nil)
}

//Store	新增顶级轮播提
func (i *SlideController) Store(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		i.rc.Error(c,"幻灯片名称不能为空！")
		return
	}
	remark := c.PostForm("remark")
	status := c.DefaultPostForm("status","1")

	s, _ := strconv.Atoi(status)

	slide := &model.Slide{Name: name,Remark: remark,Status: s}
	fmt.Println("slide：",slide)
	cmf.Db.Create(slide)
	result := cmf.Db.NewRecord(slide)

	fmt.Println("result：",result)

	if result {
		i.rc.Error(c, "添加轮播图失败！", nil)
		return
	}

	//model.Slide{}
	i.rc.Success(c, "添加轮播图成功", nil)
}

func (i *SlideController) Delete(c *gin.Context) {
	i.rc.Success(c, "操作成功Delete", nil)
}