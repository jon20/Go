package get_wes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

	//res := weather.(map[string]interface{})["description"]
	//fmt.Println((weather.(map[string]interface{})["location"]).(map[string]interface{})["prefecture"])
	//fmt.Println(res.(map[string]interface{})["text"])
}

func Getwes_t() string {
	wes := weather.(map[string]interface{})["description"].(map[string]interface{})["text"]
	res := wes.(string)
	return (res)
}

func Getwes_l() string {
	wes := (weather.(map[string]interface{})["title"]).(string)

	//res := wes.(string)
	return (wes)
}
