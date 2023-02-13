package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"text/template"
)

// GetIndex
// @Tags 首頁
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	index, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	err = index.Execute(c.Writer, "")
	if err != nil {
		fmt.Println(err)
	}
	//c.JSON(200, gin.H{
	//	"status": "success",
	//})
}
