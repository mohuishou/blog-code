package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	_, err := client.Get("http://localhost:8080")
	fmt.Println(err)
}
