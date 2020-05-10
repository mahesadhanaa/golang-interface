package interface

type ReadInterface interface {
	Read() (string, error)
}