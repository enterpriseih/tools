package instance

import (
	"fmt"
	"github.com/Ubbo-Sathla/tools/cmd/cleanlog/consul"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Instances struct {
	Name string
}

func ListInstance(c *gin.Context) {
	fmt.Println(c.Param("name"), c.Query("name"))
	k := fmt.Sprintf("Gemini/Service/%s/", c.Query("name"))
	pairs, _, err := consul.ConsulListKey(k)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "请求Consul数据异常",
		})
		return
	}
	ss := make([]*Instances, 0)
	for _, pair := range pairs {
		//fmt.Printf("%#v\n", pair)

		ss = append(ss, &Instances{
			Name: pair.Key[len(k):],
		})

	}
	fmt.Println(ss)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "",
		"data":    ss,
	})

}
