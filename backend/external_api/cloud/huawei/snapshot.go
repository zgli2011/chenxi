package huawei

type ISnapshot interface {
	List()
	Create()
	Delete()
}
