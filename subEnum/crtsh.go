package subsEnum

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CrtShResult struct {
	Name string `json:"name_value"`
}

// GetCrtshSubs ...
func GetCrtshSubs(domain string) ([]string, error) {
	var results []CrtShResult

	client := &http.Client{}
	client.Timeout = time.Second * 15

	fullUrl := "https://crt.sh?q=%25." + domain + "&output=json"

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		log.Fatalf("http.NewRequest() failed with '%s'\n", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do() failed with '%s'\n", err)
	}

	defer resp.Body.Close()
	output := make([]string, 0)

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &results); err != nil {
		return []string{}, err
	}

	for _, res := range results {
		output = append(output, res.Name)
	}
	return output, nil
}
