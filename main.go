package main

import (
	data "github.com/mahesadhanaa/golang-interface/read"
)

func main() {
	reader := data.NewReadData()
	data.Add(reader)
}
