package instance

import (
	"fmt"
	"github.com/Ubbo-Sathla/tools/cmd/cleanlog/consul"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteInstance(c *gin.Context) {
	var inst Instance
	var err error
	c.BindJSON(&inst)

	for _, ip := range inst.IP {

		err = consul.ConsulDelKV(fmt.Sprintf("Gemini/Service/%s/%s", inst.Service, ip))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "服务删除主机失败",
			})
			return
		}
		err = consul.ConsulDelKV(fmt.Sprintf("Gemini/Instance/%s/%s", ip, inst.Service))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "主机删除配置失败",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "主机删除成功",
	})
}
