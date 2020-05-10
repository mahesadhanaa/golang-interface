package data

import "fmt"

type ReadInterface interface {
	Read() (string, error)
}

type ReadData struct{}

func (r *ReadData) Read() (string, error) {
	return "mahesadhana is a ", nil
}

func NewReadData() ReadInterface {
	return &ReadData{}
}

func Add(r ReadInterface) {
	write, _ := r.Read()
	fmt.Printf("%sSite Reliability Engineer\n", write)
}
