package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moshushu/gin-blog/pkg/e"
	"github.com/moshushu/gin-blog/pkg/logging"
	"github.com/moshushu/gin-blog/pkg/upload"
)

// @Summary 图片上传
// @Description 图片上传
// @Produce json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	// c.Request.FormFile：获取上传的图片（返回提供的表单键的第一个文件）
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName

		// CheckImageExt、CheckImageSize检查图片大小，检查图片后缀
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {

			// CheckImage：检查上传图片所需（权限、文件夹）
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL

				// SaveUploadedFile：保存图片
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
