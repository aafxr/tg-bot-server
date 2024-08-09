package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

func UploadFile(s *apiserver.Server) func(*gin.Context) {
	return func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, types.Response{
				Ok:      false,
				Message: err.Error() + " 20",
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("./assets/%s", file.Filename)
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, dst)

		c.JSON(http.StatusOK, types.Response{
			Ok:      true,
			Message: fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	}
}
