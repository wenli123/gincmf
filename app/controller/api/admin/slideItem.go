/**
* @Title  轮播图子项资源控制器
* @Description  轮播图子项增删改查接口
* @Author  Return  20200519 10：46
* @Update
**/
package admin

import (
	"fmt"
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
	"strconv"
	"time"
)

// CarouselController 轮播图资源控制器，定义了轮播图增删改查接口
type SlideItemController struct {
	rc controller.RestController
}

func (rest *SlideItemController) Get(c *gin.Context) {
	var slideItem []model.SlideItem
	slideId := c.Query("slide_id")

	if slideId == "" {
		rest.rc.Error(c, "轮播图id不能为空！", nil)
		return
	}

	query := "slide_id = ? AND status = ?"
	queryArgs := []interface{}{slideId, "1"}
	title := c.Query("title")
	desc := c.Query("desc")

	if title != "" {
		query += " AND name like ?"
		queryArgs = append(queryArgs, "%"+title+"%")
	}

	if desc != "" {
		query += " AND description like ?"
		queryArgs = append(queryArgs, "%"+desc+"%")
	}

	current := c.DefaultQuery("current", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	intCurrent, _ := strconv.Atoi(current)
	intPageSize, _ := strconv.Atoi(pageSize)

	if intCurrent <= 0 {
		rest.rc.Error(c, "当前页码需大于0！", nil)
		return
	}

	if intPageSize <= 0 {
		rest.rc.Error(c, "每页数需大于0！", nil)
		return
	}

	total := 0
	cmf.Db.Limit(pageSize).Where(query, queryArgs...).Find(&slideItem).Count(&total)
	notFount := cmf.Db.Where(query, queryArgs...).Limit(intPageSize).Offset((intCurrent - 1) * intPageSize).Find(&slideItem).RecordNotFound()

	if notFount {
		rest.rc.Error(c, "该内容不存在！", nil)
		return
	}

	paginationData := &model.Paginate{Data: &slideItem, Current: current, PageSize: pageSize, Total: total}
	rest.rc.Success(c, "轮播图获取成功！", paginationData)
}

func (rest *SlideItemController) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	slide := model.Slide{}
	notFount := cmf.Db.Where("delete_at = ?", "0").First(&slide, rewrite.Id).RecordNotFound()
	if notFount {
		rest.rc.Error(c, "该内容不存在！", nil)
		return
	}
	rest.rc.Success(c, "获取轮播图成功！", slide)
}

//Store	新增顶级轮播提
func (rest *SlideItemController) Store(c *gin.Context) {

	// 接受slide_id
	slideId := c.PostForm("slide_id")
	if slideId == "" {
		rest.rc.Error(c, "该轮播项不存在！", nil)
		return
	}
	slideIdInt, _ := strconv.Atoi(slideId)

	title := c.PostForm("title")
	if title == "" {
		rest.rc.Error(c, "标题不能为空！", nil)
		return
	}

	aTarget := c.PostForm("target")
	url := c.PostForm("url")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	image := c.PostForm("image")
	listOrder := c.PostForm("list_order")
	listOrderFloat, _ := strconv.ParseFloat(listOrder, 64)

	slide := &model.SlideItem{
		SlideId:     slideIdInt,
		Title:       title,
		Target:      aTarget,
		Url:         url,
		Description: desc,
		Content:     content,
		Image:       image,
		ListOrder:   listOrderFloat,
	}

	slideData := model.SlideItem{}
	notFount := cmf.Db.Where("title = ? ", "0").First(&slideData).RecordNotFound()
	if !notFount {
		if slideData.Status == 0 {
			cmf.Db.Model(&model.SlideItem{}).Where("id = ? ", slideData.Id).Update(slide)
			rest.rc.Success(c, "添加轮播图成功！", nil)
			return
		} else {
			rest.rc.Error(c, "该轮播图已存在！", nil)
			return
		}
	}

	cmf.Db.Create(slide)
	result := cmf.Db.NewRecord(slide)
	fmt.Println("result：", result)
	if result {
		rest.rc.Error(c, "添加轮播图失败！", nil)
		return
	}

	//model.Slide{}
	rest.rc.Success(c, "添加轮播图成功！", nil)
}

func (rest *SlideItemController) Edit(c *gin.Context) {
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
		rest.rc.Error(c, "幻灯片名称不能为空！", nil)
		return
	}

	remark := c.PostForm("remark")
	status := c.DefaultPostForm("status", "1")

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

func (rest *SlideItemController) Delete(c *gin.Context) {

	var rewrite struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	slide := model.Slide{}

	notFount := cmf.Db.First(&slide, rewrite.Id).RecordNotFound()
	if notFount {
		rest.rc.Error(c, "该内容不存在！", nil)
		return
	}

	slide.DeleteAt = int(time.Now().Unix())

	if err := cmf.Db.Save(slide).Error; err != nil {
		rest.rc.Error(c, "删除轮播图失败！", nil)
		return
	}

	rest.rc.Success(c, "删除轮播图成功！", slide)
}
