package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	for {
		resp, err := http.Get("http://app2-svc.default.svc.cluster.local/api")
		if err != nil {
			log.Fatalln(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		sb := string(body)
		log.Print(sb)

		time.Sleep(time.Second)

	}

}
