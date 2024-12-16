package rules

import (
	"mime/multipart"

	"github.com/weiloon1234/gokit/utils"

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

	return utils.ValidateFile(c, file) == nil
}
