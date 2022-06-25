package subEnum

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CertSpotterEntry struct {
	DNSNAMES []string `json:"dns_names"`
}

// GetCertSpotterSubs ...
func GetCertSpotterSubs(domain string) ([]string, error) {
	fullUrl := "https://api.certspotter.com/v1/issuances?domain=" + domain + "&expand=dns_names&include_subdomains=true"

	client := &http.Client{}
	client.Timeout = time.Second * 15

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

	entries := make([]CertSpotterEntry, 0)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&entries)

	for i := 0; i < len(entries); i++ {
		for j := 0; j < len(entries[i].DNSNAMES); j++ {
			output = append(output, entries[i].DNSNAMES[j])
		}
	}

	return output, nil
}
