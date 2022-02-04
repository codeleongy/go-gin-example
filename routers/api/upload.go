package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/leong-y/go-gin-example/pkg/app"
	"github.com/leong-y/go-gin-example/pkg/e"
	"github.com/leong-y/go-gin-example/pkg/logging"
	"github.com/leong-y/go-gin-example/pkg/upload"
)

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	// 将gin上下文容器赋值给统一返回值规范的类
	appG := app.Gin{C: c}

	// 从请求中获取文本获取 file,image
	/*	file用来实现一些io操作，比如读文件
		type File interface {
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Closer
		}
	*/
	/*	FileHeader也就是我们的image，用于获取文件的一些参数
		type FileHeader struct {
			Filename string
			Header   textproto.MIMEHeader
			Size     int64
			content []byte
			tmpfile string
		}
	*/
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		// 封装和规范了返回值 Response(httpCode, errCode int, data interface{})
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// image.Filename
	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
