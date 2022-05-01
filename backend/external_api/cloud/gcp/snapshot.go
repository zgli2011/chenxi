package gcp

type ISnapshot interface {
	List()
	Create()
	Delete()
}
