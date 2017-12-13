package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/jon20/Weather-Command/get_wes"
)

var (
	flag_a   bool = true
	flag_l   bool = true
	flag_env bool = true
)

func init() {
	flag.BoolVar(&flag_a, "a", false, "getwes")
	flag.BoolVar(&flag_l, "l", false, "getwes_l")
	flag.BoolVar(&flag_env, "env", false, "show env")
	flag.Parse()
}

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

	if flag_a != false {
		fmt.Println(get_wes.Getwes_a())
	} else if flag_l != false {
		fmt.Println(get_wes.Getwes_l())
	} else if flag_env != false {
		fmt.Println(os.Getenv("otenki"))
	} else {
		get_wes.Getwes_d()
	}
}
