package routers

import (
	"net/http"

	_ "github.com/leong-y/go-gin-example/docs" // 这里需要引入本地已生成文档
	"github.com/leong-y/go-gin-example/middleware/jwt"

	"github.com/gin-gonic/gin"
	"github.com/leong-y/go-gin-example/pkg/export"
	"github.com/leong-y/go-gin-example/pkg/setting"
	"github.com/leong-y/go-gin-example/pkg/upload"
	"github.com/leong-y/go-gin-example/routers/api"
	"github.com/leong-y/go-gin-example/routers/api/version1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	// Create an instance of Engine, by using New() or Default()
	r := gin.New()
	// *gin.Engine.Use(中间件名称) 代表全局使用该中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置运行模式即测试，开发之类的提示
	gin.SetMode(setting.ServerSetting.RunMode)

	// 用于测试的路由分组
	test := r.Group("/test")
	{
		test.GET("/string", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello LEONG")
		})
	}

	// 新增获取token的方法
	r.GET("/auth", api.GetAuth)

	// 新增添加上传图片功能
	r.POST("/upload", api.UploadImage)

	// 新增图片查看功能
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// 下载excel表格
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))

	// 设置一个路由分组
	apiv1 := r.Group("/api/v1")
	// 引用JWT中间件（自定义中间件）
	apiv1.Use(jwt.JWT())
	{
		// 路由.Get(url,响应方法)
		// 获取标签列表
		apiv1.GET("/tags", version1.GetTags)
		// 新建标签
		apiv1.POST("/tags", version1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tags/:id", version1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", version1.DeleteTag)
		// 导出标签
		apiv1.POST("/tags/export", version1.ExportTag)
		// 导出标签
		apiv1.POST("/tags/import", version1.ImportTag)
		// 获取文章列表
		apiv1.GET("/articles", version1.GetArticles)
		// 获取指定文章
		apiv1.GET("/articles/:id", version1.GetArticle)
		// 新建文章
		apiv1.POST("/articles", version1.AddArticle)
		// 更新指定文章
		apiv1.PUT("/articles/:id", version1.EditArticle)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", version1.DeleteArticle)
		// 生成二维码
		apiv1.POST("/articles/poster/generate", version1.GenerateArticlePoster)
	}
	return r
}
