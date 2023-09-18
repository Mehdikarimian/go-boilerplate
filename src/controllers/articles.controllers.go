package controllers

import (
	"encoding/json"
	"go-boilerplate/src/common"
	"go-boilerplate/src/models"
	"go-boilerplate/src/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ArticlesController(r *gin.Engine) *BaseController {
	ctr := &BaseController{
		ControllerConstructor: &common.ControllerConstructor{
			// Collection: core.InitMongoDB().Collection("testArticlesCollection"),
		},
	}

	ctr.ArticlesRoutes(r)

	return ctr
}

// @BasePath /articles
func (ctr *BaseController) ArticlesRoutes(r *gin.Engine) {
	api := r.Group("/articles")
	{
		api.GET("/", func(ctx *gin.Context) {
			getArticles(ctx, ctr)
		})
		api.GET("/:id", func(ctx *gin.Context) {
			getArticle(ctx, ctr)
		})
		api.POST("/", func(ctx *gin.Context) {
			createArticle(ctx, ctr)
		})
		api.PUT("/:id", func(ctx *gin.Context) {
			updateArticle(ctx, ctr)
		})
		api.DELETE("/:id", func(ctx *gin.Context) {
			deleteArticle(ctx, ctr)
		})
	}
}

// @Summary      Get All Articles
// @Tags         articles
// @Produce      json
// @Success      200  {object}  models.Article
// @Router       /articles [get]
func getArticles(ctx *gin.Context, ctr *BaseController) {
	results, err := models.ArticlesModel().GetAllArticles()

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}

// @Summary      Find a article
// @Description  Returns the the article with ID.
// @Tags         articles
// @Produce      json
// @Success      200  {object}  models.Article
// @Router       /articles/{id} [get]
// @Param        id    path    int  false  "id"  Format(id)
func getArticle(ctx *gin.Context, ctr *BaseController) {
	id := ctx.Param("id")
	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	res := models.CacheModel().GetCache("article_" + strconv.Itoa(int(getID)))
	if len(res) > 0 {
		var article models.Article
		err := json.Unmarshal([]byte(res), &article)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid Cache"})
		}

		ctx.JSON(http.StatusOK, article)
		return
	}

	results, err := models.ArticlesModel().GetOneArticle(getID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Article Not Found!"})
		return
	}

	bs, _ := json.Marshal(results)

	models.CacheModel().SetCache("article_"+strconv.Itoa(int(getID)), bs, 0)

	ctx.JSON(http.StatusOK, results)
}

// @Summary      Create An Article
// @Description  Create An Article.
// @Tags         articles
// @Produce      json
// @Success      200  {object}  models.Article
// @Router       /articles [post]
// @Param request body models.CreateArticleForm true "body"
func createArticle(ctx *gin.Context, ctr *BaseController) {
	body := utils.GetBody[models.CreateArticleForm](ctx)

	results, err := models.ArticlesModel().CreateArticle(body)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}

// @Summary      Update An Article
// @Description  Update An Article.
// @Tags         articles
// @Produce      json
// @Success      200  {object}  models.Article
// @Router       /articles/{id} [put]
// @Param request body models.CreateArticleForm true "body"
// @Param        id    path    int  false  "id"  Format(id)
func updateArticle(ctx *gin.Context, ctr *BaseController) {
	id := ctx.Param("id")
	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}
	body := utils.GetBody[models.UpdateArticleForm](ctx)

	err = models.ArticlesModel().UpdateArticle(getID, body)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, "Article Updated")
}

// @Summary      Delete An Article
// @Description  Delete An Article.
// @Tags         articles
// @Produce      json
// @Success      200  {object}  models.Article
// @Router       /articles/{id} [delete]
// @Param        id    path    int  false  "id"  Format(id)
func deleteArticle(ctx *gin.Context, ctr *BaseController) {
	id := ctx.Param("id")
	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	err = models.ArticlesModel().DeleteArticle(getID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, "Article Deleted")
}
