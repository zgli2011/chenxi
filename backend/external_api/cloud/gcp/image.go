package gcp

type IImage interface {
	List()
	Create()
	Delete()
}
