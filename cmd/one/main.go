// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package one

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func httpRequest() error {
	resp, err := http.Get("http://localhost:1234/brexitDate")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	println("brexit date is: " + string(body))

	request, err := http.NewRequest(http.MethodPut, "http://localhost:1234/brexitDate", strings.NewReader("29rd April"))
	if err != nil {
		return err
	}
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	println("put succeeded")
	return nil
}

func main() {
	if err := httpRequest(); err != nil {
		log.Fatal(err)
	}
}
