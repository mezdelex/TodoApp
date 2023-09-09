package interfaces

type Getter interface {
	GetAll() []interface{}
	Get() interface{}
}
