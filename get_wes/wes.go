package get_wes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	scan "github.com/mattn/go-scan"
)

var weather interface{}

func Getwes(response *http.Response) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Read failure")
		return
	}
	//fmt.Println(string(body))
	json.Unmarshal(body, &weather)
	if weather == nil {
		fmt.Println("Read failure")
		os.Exit(1)
	}

}

func Getwes_d() {
	var day_td string
	var day_tm string
	var td string
	var tm string
	scan.ScanTree(weather, "/forecasts[0]/telop", &td)
	scan.ScanTree(weather, "/forecasts[1]/telop", &tm)
	scan.ScanTree(weather, "/forecasts[0]/date", &day_td)
	scan.ScanTree(weather, "/forecasts[1]/date", &day_tm)
	fmt.Println(day_td+" "+td,
		"\n"+day_tm+" "+tm)
}

func Getwes_a() string {
	wes := weather.(map[string]interface{})["description"].(map[string]interface{})["text"]
	res := wes.(string)
	return (res)
}

func Getwes_l() string {
	wes := (weather.(map[string]interface{})["title"]).(string)
	return (wes)
}
