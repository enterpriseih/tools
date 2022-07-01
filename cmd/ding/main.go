package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/nxsre/godingtalk"
	"github.com/tidwall/pretty"
	"log"
	"os"
	"strings"
)

var ding *godingtalk.DingTalkClient

func main() {
	ding = godingtalk.NewDingTalkClient("dingmt4vjb7lxsimvolp", "kHxHzI5zwgdvULqteCvWlz4I7eDpzaJYgtSP6ekZ4ZmEwreYELhEDqlljYO5F3kW")
	ding.RefreshAccessToken()
	uid, err := ding.UseridByMobile(os.Args[1]) // Input: 18345463525
	if err != nil {
		log.Fatalln(err)
	}
	uinfo, err := ding.UserInfoByID(uid)

	if err != nil {
		log.Fatalln(err)
	}

	dst := struct {
		Name       string
		Mobile     string
		Email      string
		Userid     string
		Department []int
		Depts      []string
	}{}

	copier.Copy(&dst, &uinfo)

	// 查找用户归属组（直到顶层架构）
	for _, deptId := range dst.Department {
		var res []string
		res = walk(deptId, &res)
		dst.Depts = append(dst.Depts, strings.Join(reverse(res), "-"))
	}

	bs, err := json.Marshal(&dst)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(pretty.Pretty(bs)))

	//dlist,err:=ding.DepartmentList(1,true)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("%+v",dlist)
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func walk(deptId int, res *[]string) []string {
	dept, err := ding.DepartmentDetail(deptId)
	if err != nil {
		log.Fatalln(err)
	}
	*res = append(*res, fmt.Sprintf("%s(%d)", dept.Name, dept.Id))
	if deptId != 1 {
		walk(dept.ParentId, res)
	}
	return *res
}
