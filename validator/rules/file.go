package rules

import (
	"github.com/weiloon1234/gokit/utils"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

// File checks if the value is an uploaded file
// Params: [] (no additional parameters required)
func File(value interface{}, params []string, c *gin.Context) bool {
	// Ensure the value is of type *multipart.FileHeader
	file, ok := value.(*multipart.FileHeader)
	if !ok {
		return false
	}

	err := utils.ValidateFile(c, file)
	if err != nil {
		return false
	}

	return true
}
