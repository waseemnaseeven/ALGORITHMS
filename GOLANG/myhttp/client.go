package main

import (
	"net/http"
	"fmt"
)

func Client() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status : ", resp.Status)

}