package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiURL = "http://localhost:5000"
)

// SendSMSRest Contact REST API
func SendSMSRest(phone string, message string) {
	url := apiURL + "/sms"
	fmt.Println("URL:>", url)

	values := map[string]string{"message": message, "phone": phone}
	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("X-Custom-Header", "GoSMS")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
