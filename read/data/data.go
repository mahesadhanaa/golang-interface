package data

import (
	"fmt"

	readinterface "github.com/mahesadhanaa/golang-interface/read"
)

type ReadData struct{}

func (r *ReadData) Read() (string, error) {
	return "mahesadhana is a ", nil
}

func NewReadData() readinterface.ReadInterface {
	return &ReadData{}
}

func Add(r readinterface.ReadInterface) {
	write, _ := r.Read()
	fmt.Printf("%sSite Reliability Engineer\n", write)
}
