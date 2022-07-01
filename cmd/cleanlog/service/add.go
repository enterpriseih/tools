package service

import (
	"encoding/json"
	"fmt"
	"github.com/Ubbo-Sathla/tools/cmd/cleanlog/consul"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	Name   string
	Config Config
}

type Config struct {
	LogDirs    []string `yaml:"LogDirs"`
	RegLogDirs []string `yaml:"RegLogDirs"`
	RegFiles   []string `yaml:"RegFiles"`
	ModTime    int64    `yaml:"ModTime"`
	FileSize   int64    `yaml:"FileSize"`
	Delete     bool     `yaml:"Delete"`
}

func AddServiceConfig(c *gin.Context) {

	var s Service
	var sk string
	err := c.BindJSON(&s)
	fmt.Println(s)
	if err != nil {
		fmt.Println(err)
	}
	if s.Config.FileSize == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "请求参数异常",
		})
		return
	}
	if s.Name != "" {
		sk = fmt.Sprintf("Gemini/ServiceConfig/%s", s.Name)
	}

	data, err := json.Marshal(s.Config)
	if err != nil {
		fmt.Println(err)
	} else {
		pair, _, err := consul.ConsulGetKV(sk)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "请求Consul数据异常",
			})
			return
		}

		if pair == nil {
			err = consul.ConsulPutKV(sk, string(data))
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "请求Consul数据异常",
				})
				return
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "请不要重复添加",
			})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "添加成功",
	})
}
