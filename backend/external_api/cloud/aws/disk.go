package aws

type IDisk interface {
	List()
	Create()
	Delete()
	Update()
}
