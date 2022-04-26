package utils

import (
	"fmt"
	"testing"
)

func Test_Get(t *testing.T) {
	header := Header{"x-token": "U9TEbO5hyC7aFnpMDwudkSrxv8ZIt2eV"}
	url := "http://cmdb-dev.sheincorp.cn/api/v2/cmdb/resource_server/"
	params := Params{"no_page": "true", "is_user": "false", "name": "taskpower"}
	resp, err := Get(url, header, params)
	if err != nil {
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}
func Test_Put(t *testing.T) {
	header := Header{"x-token": "U9TEbO5hyC7aFnpMDwudkSrxv8ZIt2eV"}
	url := "http://cmdb-dev.sheincorp.cn/api/v2/cmdb/server_up/73556284-dc09-41ce-9aea-807acc8b45b7/"
	resp, err := Put(url, header, Jsons(map[string]interface{}{"available_state": "up"}))
	if err != nil {
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}
func Test_Delete(t *testing.T) {
	header := Header{"x-token": "U9TEbO5hyC7aFnpMDwudkSrxv8ZIt2eV"}
	url := "http://cmdb-dev.sheincorp.cn/api/v2/cmdb/service/1F0E8C65165E460B87415886D70E76C3/"
	resp, err := PostJson(url, header)
	if err != nil {
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}
func Test_PostJson(t *testing.T) {
	header := Header{"x-token": "U9TEbO5hyC7aFnpMDwudkSrxv8ZIt2eV"}
	url := "http://cmdb-dev.sheincorp.cn/api/v1/task_schedule/task/sync/"
	jsons := Jsons(map[string]interface{}{
		"script_id": 15,
		"ip_list":   []string{"10.123.12.7"},
		"task_name": "task_scheduling_stress_test",
		"timeout":   8,
	})
	resp, err := PostJson(url, header, jsons)
	if err != nil {
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text())
}
func Test_PostForm(t *testing.T) {}
