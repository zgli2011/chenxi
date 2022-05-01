package amazon

type ISnapshot interface {
	List()
	Create()
	Delete()
}
