package controllers

import (
	"go-boilerplate/src/common"
	"go-boilerplate/src/models"
	"go-boilerplate/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductsController(r *gin.Engine) *BaseController {
	ctr := &BaseController{
		ControllerConstructor: &common.ControllerConstructor{
			// Collection: core.InitMongoDB().Collection("ProductsCollection"),
		},
	}

	ctr.ProductsRoutes(r)

	return ctr
}

// @BasePath /products
func (ctr *BaseController) ProductsRoutes(r *gin.Engine) {
	api := r.Group("/products")
	{
		api.GET("/", func(ctx *gin.Context) {
			getProducts(ctx, ctr)
		})
		api.GET("/:id", func(ctx *gin.Context) {
			getProduct(ctx, ctr)
		})
		api.POST("/", func(ctx *gin.Context) {
			createProduct(ctx, ctr)
		})
		api.PUT("/:id", func(ctx *gin.Context) {
			updateProduct(ctx, ctr)
		})
		api.DELETE("/:id", func(ctx *gin.Context) {
			deleteProduct(ctx, ctr)
		})
	}
}

// @Summary      Get All Products
// @Tags         products
// @Produce      json
// @Success      200  {object}  models.Product
// @Router       /products [get]
// @Param        limit    query    int  false  "limit"  Format(limit)
// @Param        skip    query    int  false  "skip"  Format(skip)
// @Param        search    query    string  false  "search"  Format(search)
func getProducts(ctx *gin.Context, ctr *BaseController) {
	limit := utils.GetQueryInt(ctx, "limit", 20)
	skip := utils.GetQueryInt(ctx, "skip", 0)
	search := utils.GetQueryString(ctx, "search", "")

	results, err := models.ProductsModel().GetAllProducts(int64(limit), int64(skip), search)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}

// @Summary      Find a product
// @Description  Returns the the product with ID.
// @Tags         products
// @Produce      json
// @Success      200  {object}  models.Product
// @Router       /products/{id} [get]
// @Param        id    path    string  false  "id"  Format(id)
func getProduct(ctx *gin.Context, ctr *BaseController) {
	id := ctx.Param("id")

	results, err := models.ProductsModel().GetOneProduct(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product Not Found!"})
		return
	}

	ctx.JSON(http.StatusOK, results)
}

// @Summary      Create An Product
// @Description  Create An Product.
// @Tags         products
// @Produce      json
// @Success      200  {object}  models.Product
// @Router       /products [post]
// @Param request body models.CreateProductForm true "body"
func createProduct(ctx *gin.Context, ctr *BaseController) {
	body := utils.GetBody[models.CreateProductForm](ctx)

	results, err := models.ProductsModel().CreateProduct(body)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}

// @Summary      Update An Product
// @Description  Update An Product.
// @Tags         products
// @Produce      json
// @Success      200  {object}  models.Product
// @Router       /products/{id} [put]
// @Param request body models.CreateProductForm true "body"
// @Param        id    path    string  false  "id"  Format(id)
func updateProduct(ctx *gin.Context, ctr *BaseController) {
	id := ctx.Param("id")

	body := utils.GetBody[models.UpdateProductForm](ctx)

	result, err := models.ProductsModel().UpdateProduct(id, body)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// @Summary      Delete An Product
// @Description  Delete An Product.
// @Tags         products
// @Produce      json
// @Success      200  {object}  models.Product
// @Router       /products/{id} [delete]
// @Param        id    path    string  false  "id"  Format(id)
func deleteProduct(ctx *gin.Context, ctr *BaseController) {
	id := ctx.Param("id")

	result, err := models.ProductsModel().DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
