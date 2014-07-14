package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	b := strings.NewReader(`{"Name":"httpclient"}`)
	bodyType:="application/json"
	reqest, err:= http.Post( "http://localhost:9092/",bodyType, b)
	if err == nil {
		body, _ := ioutil.ReadAll(reqest.Body)
		fmt.Printf("Body: %s\n", body)
		reqest.Body.Close()
	} else {
		fmt.Printf("Error: %v\n", err)
	}

	/*response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body);
		fmt.Println(bodystr)
	}*/
}

