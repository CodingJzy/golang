package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/query_map_param", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("users")
		c.JSON(200, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
	// query_map_param?ids[id1]=1&ids[id2]=2
	// post_data: users[name] jiang_wei users[age] 23
	/*

		{
		    "ids": {
		        "id1": "1",
		        "id2": "2"
		    },
		    "names": {
		        "age": "23",
		        "name": "jiang_wei"
		    }
		}
	*/

	r.Run()
}
