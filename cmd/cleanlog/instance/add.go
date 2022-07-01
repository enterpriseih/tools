package instance

import (
	"fmt"
	"github.com/Ubbo-Sathla/tools/cmd/cleanlog/consul"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type Instance struct {
	IP      []string
	Service string
}

func AddInstance(c *gin.Context) {
	var inst Instance
	var err error
	c.BindJSON(&inst)

	for _, ip := range inst.IP {
		address := net.ParseIP(ip)
		if address == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "IP异常",
			})
			return
		}
	}

	pair, _, err := consul.ConsulGetKV(fmt.Sprintf("Gemini/ServiceConfig/%s", inst.Service))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "请求Consul数据异常",
		})
		return
	}

	if pair == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Service不存在",
		})
		return
	}

	//fmt.Printf("%#v\n", pair)
	for _, ip := range inst.IP {

		err = consul.ConsulPutKV(fmt.Sprintf("Gemini/Service/%s/%s", inst.Service, ip), "")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "服务添加主机失败",
			})
			return
		}
		err = consul.ConsulPutKV(fmt.Sprintf("Gemini/Instance/%s/%s", ip, inst.Service), "")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "主机添加配置失败",
			})
			return
		}
	}
}
