package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"./get_wes"
)

func main() {
	values := url.Values{}
	//Get ID env
	otenki := os.Getenv("otenki")

	if otenki == "" {
		fmt.Println("Please set env ")
		os.Exit(1)
	}

	//Add env
	values.Add("city", otenki)
	resp, err := http.Get("http://weather.livedoor.com/forecast/webservice/json/v1" + "?" + values.Encode())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	get_wes.Getwes(resp)
	msg := flag.Bool("a", false, "sdf")
	flag.Parse()
	if *msg == true {
		fmt.Println(get_wes.Getwes_t())
	}
}
