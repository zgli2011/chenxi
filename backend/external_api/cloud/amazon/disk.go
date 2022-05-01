package amazon

type IDisk interface {
	List()
	Create()
	Delete()
	Update()
}
