// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func httpRequest() {
	resp, _ := http.Get("http://localhost:1234/brexitDate")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	println("brexit date is: " + string(body))

	request, _ := http.NewRequest(http.MethodPut, "http://localhost:1234/brexitDate", strings.NewReader("31st October"))
	resp, _ = http.DefaultClient.Do(request)
	println("PUT succeeded")
}

func main() {
	httpRequest()
}
