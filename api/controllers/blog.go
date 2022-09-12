package controllers

import (
	"blog/api/services"
	"blog/models"
	"blog/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//PostController -> PostController
type PostController struct {
	services services.PostService
}

//NewPostController : NewPostController
func NewPostController(s services.PostService) PostController {
	return PostController{
		services: s,
	}
}

// GetPosts : GetPosts controller
func (p PostController) GetPosts(ctx *gin.Context) {
	var posts models.Post

	keyword := ctx.Query("keyword")

	data, total, err := p.services.FindAll(posts, keyword)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Post result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddPost : AddPost controller
func (p *PostController) AddPost(ctx *gin.Context) {
	var post models.Post
	ctx.ShouldBindJSON(&post)

	if post.Title == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if post.Body == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}
	err := p.services.Save(post)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create post")
		return
	}
	utils.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Post")
}

//GetPost : get post by id
func (p *PostController) GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var post models.Post
	post.ID = id
	foundPost, err := p.services.Find(post)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Error Finding Post")
		return
	}
    response := foundPost.ResponseMap()

	c.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Result set of Post",
		Data:    &response})

}

//DeletePost : Deletes Post
func (p *PostController) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.services.Delete(id)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Post")
		return
	}
	response := &utils.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

//UpdatePost : get update by id
func (p PostController) UpdatePost(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var post models.Post
	post.ID = id

	postRecord, err := p.services.Find(post)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Post with given id not found")
		return
	}
	ctx.ShouldBindJSON(&postRecord)

	if postRecord.Title == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if postRecord.Body == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	if err := p.services.Update(postRecord); err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Post")
		return
	}
	response := postRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Successfully Updated Post",
		Data:    response,
	})
}
