package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for {
		resp, err := http.Get("http://app2-svc.default.svc.cluster.local:4040/api")
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

		sb = fmt.Sprintln(sb)
		log.Print(sb)
		_, err2 := f.WriteString(sb)

		if err2 != nil {
			log.Fatal(err2)
		}

		time.Sleep(time.Second)

	}

}
