package routers

import (
	"net/http"

	"github.com/moshushu/gin-blog/docs"
	"github.com/moshushu/gin-blog/middleware/jwt"
	"github.com/moshushu/gin-blog/pkg/setting"
	"github.com/moshushu/gin-blog/pkg/upload"
	"github.com/moshushu/gin-blog/routers/api"
	v1 "github.com/moshushu/gin-blog/routers/api/v1"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 实现 http.FileServer
	// 用于读取的上传文件
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/auth", api.GetAuth)

	// 针对 swagger 新增初始化动作和对应的路由规则
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 图片上传
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		// 新建文章
		apiv1.POST("/articles", v1.AddArticle)
		// 更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
