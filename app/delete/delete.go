package delete

type Deleter interface {
	Locate() error
	Delete() error
}
