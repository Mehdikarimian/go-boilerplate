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

func (ctr *BaseController) ArticlesRoutes(r *gin.Engine) {
	api := r.Group("/articles")
	{
		api.GET("/", func(ctx *gin.Context) {
			results, err := models.ArticlesModel().GetAllArticles()

			if err != nil {
				ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, results)
		})

		api.GET("/:id", func(ctx *gin.Context) {
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
		})

		api.POST("/", func(ctx *gin.Context) {
			body := utils.GetBody[models.CreateArticleForm](ctx)

			results, err := models.ArticlesModel().CreateArticle(body)

			if err != nil {
				ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, results)
		})

		api.PUT("/:id", func(ctx *gin.Context) {
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
		})

		api.DELETE("/:id", func(ctx *gin.Context) {
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
		})
	}
}
