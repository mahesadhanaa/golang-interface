package main

import (
	data "github.com/mahesadhanaa/golang-interface/read/data"
)

func main() {
	reader := data.NewReadData()
	data.Add(reader)
}
