package MicroMDMAPI

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sethgrid/pester"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *pester.Client
	log        *log.Logger
}

func New(baseURL string, apiKey string, loggerFlags int) *Client {
	client := new(Client)

	client.baseURL = baseURL
	client.apiKey = apiKey
	client.log = log.New(os.Stderr, "[MicroMDM Client] ", loggerFlags)

	httpClient := pester.New()
	client.httpClient = httpClient
	httpClient.Timeout = time.Duration(30 * time.Second)
	httpClient.Concurrency = 10
	httpClient.MaxRetries = 5
	httpClient.Backoff = pester.ExponentialJitterBackoff
	httpClient.Transport = &http.Transport{
		MaxConnsPerHost:     50,
		MaxIdleConnsPerHost: 51,
	}

	return client
}

func (client *Client) ExecuteMDMCommand(command CommandRequest) (*CommandResponse, error) {
	data, err := client.doAPIRequest("v1/commands", "POST", command)
	if err != nil {
		return nil, err
	}

	response := new(CommandResponse)
	err = json.Unmarshal(data, &response)
	return response, err
}

func (client *Client) doAPIRequest(path string, method string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		println(string(bodyData))
		bodyReader = bytes.NewBuffer(bodyData)
	}

	req, _ := http.NewRequest("GET", client.baseURL+"/"+path, bodyReader)
	req.Header.Add("User-Agent", "Peach-Panic")
	req.SetBasicAuth("micromdm", client.apiKey)
	req.Method = method

	resp, err := client.httpClient.Do(req)
	if err != nil {
		log.Println(err)
		client.log.Println("Request error: " + err.Error())
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = errors.New(fmt.Sprintf("[MicroMDM Client] Server returned failing status code %s (%v)", resp.Status, resp.StatusCode))
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return responseBody, nil
}
