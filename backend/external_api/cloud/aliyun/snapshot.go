package aliyun

type ISnapshot interface {
	List()
	Create()
	Delete()
}
