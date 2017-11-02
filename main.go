package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const hookURL = "http://uchat-sandbox.corp.uber.internal/hooks/61nuj4jhifdixp1wpwaan6y34h"

func init() {
	fmt.Println("Program Initializing")
}
func main() {
	fmt.Println("Neethi Learning GO ")
	myLogic()

}

/*func learnGO(name string) (reply string) {
	fmt.Println("Learning go, this is an awesome function ")
	return "Totally"
}*/

func myLogic() (string, int, error) {

	dataMap := map[string]string{
		"channel":  "super-test",
		"text":     "Danger!!",
		"username": "Neethi",
	}

	data, err := json.Marshal(dataMap)
	if err != nil {
		return "", 1, fmt.Errorf("failed to marshal to json - %v", err.Error())

	}

	rdr := bytes.NewReader(data)
	mimeType := "text/plain"
	response, err := http.Post(hookURL, mimeType, rdr)
	if err != nil {
		return "", 0, fmt.Errorf("failed to marshal to json - %v", err.Error())
	}

	fmt.Println("the response is ", response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", 0, fmt.Errorf("could not read response body - %v", err.Error())
	}
	fmt.Println("Response Code : ", response.StatusCode)
	fmt.Println("Response:", body)
	return "neethi", response.StatusCode, err

}
