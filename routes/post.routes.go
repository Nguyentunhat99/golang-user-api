package routes

import (
	"github.com/example/Nhat-golang-test/controllers"
	"github.com/gin-gonic/gin"

	"github.com/example/Nhat-golang-test/services"
)

type PostRouteController struct {
	PostController controllers.PostController
}

func NewPostRouteController(PostController controllers.PostController) PostRouteController {
	return PostRouteController{PostController}
}

func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup, postService services.PostService) {
	router := rg.Group("/posts")
	router.POST("/create", pc.PostController.CreatePost)
	router.PATCH("/update/:postId", pc.PostController.UpdatePost)
	router.DELETE("/delete/:postId", pc.PostController.DeletePost)
	router.GET("/", pc.PostController.FindPosts)
	router.GET("/:postId", pc.PostController.FindPostById)
}
