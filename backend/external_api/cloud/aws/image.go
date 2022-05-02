package aws

type IImage interface {
	List()
	Create()
	Delete()
}
