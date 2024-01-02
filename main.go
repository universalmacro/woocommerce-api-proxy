package main

import (
	"fmt"
	"io"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/woocommerce-api/:notifyUrl/:uri", func(c *gin.Context) {
		notifyUrl := c.Param("notifyUrl")
		uri := c.Param("uri")
		client := &http.Client{}
		url := fmt.Sprintf("https://%s/?wc-api=%s", notifyUrl, uri)
		b, _ := io.ReadAll(c.Request.Body)
		myReader := strings.NewReader(string(b))
		req, _ := http.NewRequest(http.MethodPost, url, myReader)
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode == 404 {
			c.String(404, "error")
			return
		} else {
			c.String(200, "success")
		}
	})
	router.Run(":10000")
}
