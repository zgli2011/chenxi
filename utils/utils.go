package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func StringListContain(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func IntListContain(items []int, item int) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func DelListItem(items []string, item string) []string {
	for index, i := range items {
		if i == item {
			items = append(items[:index], items[index+1:]...)
		}
	}
	return items
}

func ArrayDuplication(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	return arr[:j]
}

func CheckFile(path string) bool {
	return CheckDirOrFileExist(path)
}

func CheckDir(path string) bool {
	return CheckDirOrFileExist(path)
}

// 判断目录/文件是否存在
func CheckDirOrFileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return os.IsExist(err)
	}
	return true
}

// 判断目录是否存在， 不存在则创建目录
func CheckDirAndCreate(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return true
		} else {
			return false
		}
	}
}

func GetLocalIP4() []string {
	var ipList []string
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, address := range addrList {
		if ip, ok := address.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				ipList = append(ipList, ip.IP.String())
			}

		}
	}
	return ipList
}

func Interface2String(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
