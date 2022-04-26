package middleware

import (
	"chenxi/initialize"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"

	"chenxi/utils"

	"github.com/gin-gonic/gin"
)

type APIInfo struct {
	ApiVersion  string `json:"api_version"`
	Project     string `json:"project"`
	Application string `json:"application"`
	Uri         string `json:"uri"`
	Method      string `json:"method"`
	Service     string `json:"service"`
}

// 将所有接口注册到网关
func RegistGateway(route_list gin.RoutesInfo) gin.HandlerFunc {
	data := []APIInfo{}
	for _, route := range route_list {
		s := strings.Split(route.Path, "/")
		if s[1] == "api" {
			s[1] = "api/" + s[2]
			s = append(s[:2], s[3:]...)
		}
		for index, a := range s {
			if strings.HasPrefix(a, ":") {
				s[index] = "<PK>"
			}
		}
		r := APIInfo{s[1], s[2], s[3], route.Path, route.Method, ""}
		data = append(data, r)
	}
	resp, err := utils.PostJson(initialize.Config.System.APIRegister, utils.Jsons(data))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Text())
	return func(c *gin.Context) {
		c.Next()
	}
}

// 校验请求源是否在白名单中
func AccessIPWhitelist(c *gin.Context) {
	request_ip, _ := c.RemoteIP()
	ip := request_ip.String()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	if ok := utils.StringListContain(initialize.Config.System.AccessIPWhiteList, ip); !ok {
		panic("request source not permission")
	}
	c.Next()
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()

			// recover错误，转string
			errorToString := func(r interface{}) string {
				switch v := r.(type) {
				case error:
					return v.Error()
				default:
					return r.(string)
				}
			}

			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 1,
				"msg":  errorToString(r),
				"data": nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}
