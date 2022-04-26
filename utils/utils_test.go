package utils

import (
	"fmt"
	"runtime"
	"testing"
)

func Test_ArrayDuplication(t *testing.T) {
	a := []string{"1", "2", "3", "4", "1", "3"}
	fmt.Println(a)
	b := ArrayDuplication(a)
	fmt.Println(b)
}

func Test_AES(t *testing.T) {
	fmt.Println(AesEncrypt("LWGqfCStcVD8oJhN", "VfIXn1uyGonrTaof1l1!BbZfANg&FJpJ"))
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// func test1(ctx context.Context) {
// 	traceId, ok := ctx.Value("traceID").(string)
// 	if ok {
// 		fmt.Printf("process over. traceID=%s\n", traceId)
// 	} else {
// 		fmt.Printf("process over. no traceID\n")
// 	}
// }

func Test_test(t *testing.T) {

	// context1 := context.Background()
	// context1 = context.WithValue(context1, "traceID", "1")
	// test1(context1)

	// context2 := context.Background()
	// context2 = context.WithValue(context2, "traceID", "2")
	// test1(context2)

	// interval := 1
	// time.Sleep(time.Duration(interval) * time.Second)

	// url := fmt.Sprintf("/api/v1/task_scheduling/task/result/%s/", "dddd")
	// fmt.Println(url)

	// p := "{\"test1\":{\"test2\":{\"test3\":1}}}"
	// paramName := "test1::test2::test3"
	// for _, param := range strings.Split(paramName, "::") {
	// 	fmt.Println(param)
	// 	var data map[string]interface{}
	// 	if err := json.Unmarshal([]byte(p), &data); err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	fmt.Println(data)
	// 	p = Interface2String(data[param])
	// }
	// fmt.Println(p)

	// reg := regexp.MustCompile(`^{{.*}}$`)
	// buf := "12 {{1.1.1.1-->2-->test}} {{1-->test}} "
	// if reg == nil {
	// 	fmt.Println("MustCompile err")
	// 	return
	// }
	// reg2 := regexp.MustCompile(`^{{\d+-->.+}}$`)
	// reg3 := regexp.MustCompile(`^{{.+-->\d+-->.+}}$`)
	// for index, param := range strings.Split(buf, " ") {
	// 	if param == "" {
	// 		continue
	// 	}
	// 	fmt.Println(index)
	// 	fmt.Println(reg2.FindString(param))
	// 	fmt.Println(reg3.FindString(param))
	// 	fmt.Println(index)
	// 	fmt.Println(index, param)
	// 	result := reg.FindString(param)
	// 	if result == "" {
	// 		continue
	// 	}
	// 	reg1 := regexp.MustCompile("{|}")
	// 	fmt.Println(strings.Split(reg1.ReplaceAllString(result, ""), "-->"))

	// }

	// fmt.Printf("Task Suspend by %s at %s", "aaa", time.Now())
	// s := map[int]string{10: "10", 9: "9", 8: "8"}
	// keys := []int{}
	// for index, step := range s {
	// 	fmt.Println(index, step)
	// 	keys = append(keys, index)
	// }
	// sort.Ints(keys)
	// fmt.Println(keys)
	// for index, value := range keys {
	// 	fmt.Println(index, s[value])
	// }

}
