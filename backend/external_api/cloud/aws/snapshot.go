package aws

type ISnapshot interface {
	List()
	Create()
	Delete()
}
