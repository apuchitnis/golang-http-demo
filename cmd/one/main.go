package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func performHttpRequests() {
	// Get and print the Brexit date.
	response, _ := http.Get("http://localhost:1234/brexitDate")
	body, _ := ioutil.ReadAll(response.Body)
	println("brexit date is:", string(body))

	// Update the Brexit date.
	request, _ := http.NewRequest(http.MethodPut,
			"http://localhost:1234/brexitDate",
			strings.NewReader("31st October"))
	_, _ = http.DefaultClient.Do(request)
	println("PUT succeeded")
}

func main() {
	performHttpRequests()
}
