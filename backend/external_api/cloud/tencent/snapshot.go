package tencent

type ISnapshot interface {
	List()
	Create()
	Delete()
}
