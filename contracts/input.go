package contracts

type Input interface {
	Validate() error
}
