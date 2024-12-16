package rules

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/utils"
)

// Image checks if the value is an uploaded image file
// Params: [] (no additional parameters required)
func Image(value interface{}, params []string, c *gin.Context) bool {
	// Ensure the value is of type *multipart.FileHeader
	file, ok := value.(*multipart.FileHeader)
	if !ok {
		return false
	}

	err := utils.ValidateImage(c, file)
	return err == nil
}
