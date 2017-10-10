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
	var def bool = true
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
	msg_l := flag.Bool("l", false, "location")

	flag.Parse()
	Check_flag(*msg, &def, get_wes.Getwes_a)
	Check_flag(*msg_l, &def, get_wes.Getwes_l)
	if def == true {
		get_wes.Getwes_d()
	}
}

func Check_flag(check bool, def *bool, getwes func() string) {

	if check == true {
		fmt.Println(getwes())
		if *def == true {
			*def = false
		}
	}
}
