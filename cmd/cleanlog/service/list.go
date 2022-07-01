package service

import (
	"encoding/json"
	"fmt"
	"github.com/Ubbo-Sathla/tools/cmd/cleanlog/consul"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListServiceConfig(c *gin.Context) {
	k := "Gemini/ServiceConfig/"
	pairs, _, err := consul.ConsulListKey(k)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "请求Consul数据异常",
		})
		return
	}
	ss := make([]*Service, 0)
	for _, pair := range pairs {
		//fmt.Printf("%#v\n", pair)
		var kc Config
		err = json.Unmarshal(pair.Value, &kc)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Service数据解析异常",
			})
			return
		}
		ss = append(ss, &Service{
			Name:   pair.Key[len(k):],
			Config: kc,
		})

	}
	fmt.Println(ss)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "",
		"data":    ss,
	})

}
