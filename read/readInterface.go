package readinterface

type ReadInterface interface {
	Read() (string, error)
}
