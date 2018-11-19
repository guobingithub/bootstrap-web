/*
 * @Author: bingo
 * @Date: 2018-11-16 10:55:02
 */
package main

import "bootstrap-web/utils"
import "github.com/gin-gonic/gin"
import "html/template"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		//创建一个分页器，一万条数据，每页30条
		pagination := utils.NewPagination(c.Request, 10000, 30)
		//传到模板中需要转换成template.HTML类型，否则html代码会被转义
		c.HTML(200,"C:/My/gopath/src/bootstrap-web/views/index.tmpl",gin.H{"pages":template.HTML(pagination.Pages())})
	})
	r.Run(":8011")
}

