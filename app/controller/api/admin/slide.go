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
	"time"
)

// CarouselController 轮播图资源控制器，定义了轮播图增删改查接口
type SlideController struct {
	rc controller.RestController
}

func (rest *SlideController) Get(c *gin.Context) {
	var slides []model.Slide

	query := "delete_at = ?"
	queryArgs := []interface{}{"0"}
	name := c.Query("name")

	if name != "" {
		query += " AND name like ?"
		queryArgs = append(queryArgs,"%"+name+"%")
	}

	current := c.DefaultQuery("current","1")
	pageSize := c.DefaultQuery("pageSize","10")

	intCurrent ,_ := strconv.Atoi(current)
	intPageSize,_ := strconv.Atoi(pageSize)

	if intCurrent <= 0 {
		rest.rc.Error(c,"当前页码需大于0！",nil)
	}

	if intPageSize <= 0 {
		rest.rc.Error(c,"每页数需大于0！",nil)
	}

	total := 0
	cmf.Db.Limit(pageSize).Where(query, queryArgs...).Find(&slides).Count(&total)
	notFount := cmf.Db.Where(query, queryArgs...).Limit(intPageSize).Offset((intCurrent - 1) * intPageSize).Find(&slides).RecordNotFound()

	if notFount {
		rest.rc.Error(c,"该内容不存在！",nil)
	}

	paginationData := &model.Paginate{Data: &slides, Current: current, PageSize: pageSize, Total: total}
	rest.rc.Success(c, "获取轮播图成功！", paginationData)
}

func (rest *SlideController) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	slide := model.Slide{}
	notFount := cmf.Db.Where("delete_at = ?", "0").First(&slide,rewrite.Id).RecordNotFound()
	if notFount {
		rest.rc.Error(c,"该内容不存在！",nil)
		return
	}
	rest.rc.Success(c, "获取轮播图成功！", slide)
}

//Store	新增顶级轮播提
func (rest *SlideController) Store(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		rest.rc.Error(c,"幻灯片名称不能为空！",nil)
		return
	}
	remark := c.PostForm("remark")
	status := c.DefaultPostForm("status","1")

	slide := &model.Slide{Name: name,Remark: remark,Status: status}

	notFount := cmf.Db.Where("name = ? AND delete_at = ?",name,"0").First(&slide).RecordNotFound()
	if !notFount {
		rest.rc.Error(c,"该内容已存在！",nil)
		return
	}
	cmf.Db.Create(slide)
	result := cmf.Db.NewRecord(slide)

	fmt.Println("result：",result)

	if result {
		rest.rc.Error(c, "添加轮播图失败！", nil)
		return
	}

	//model.Slide{}
	rest.rc.Success(c, "添加轮播图成功！", nil)
}

func (rest *SlideController) Edit(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	slide := model.Slide{}

	name := c.PostForm("name")
	if name == "" {
		rest.rc.Error(c,"幻灯片名称不能为空！",nil)
		return
	}

	remark := c.PostForm("remark")
	status := c.DefaultPostForm("status","1")

	slide.Id = rewrite.Id
	slide.Name = name
	slide.Remark = remark
	slide.Status = status

	if err := cmf.Db.Save(slide).Error; err != nil {
		rest.rc.Error(c, "修改轮播图失败！", nil)
		return
	}

	rest.rc.Success(c, "修改轮播图成功！", slide)
}



func (rest *SlideController) Delete(c *gin.Context) {

	var rewrite struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	slide := model.Slide{}

	notFount := cmf.Db.First(&slide,rewrite.Id).RecordNotFound()
	if notFount {
		rest.rc.Error(c,"该内容不存在！",nil)
		return
	}

	slide.DeleteAt = int(time.Now().Unix())

	if err := cmf.Db.Save(slide).Error; err != nil {
		rest.rc.Error(c, "删除轮播图失败！", nil)
		return
	}

	rest.rc.Success(c, "删除轮播图成功！", slide)
}