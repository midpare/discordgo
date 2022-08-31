package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Method string

const (
	Method_Put    Method = "PUT"
	Method_Get    Method = "GET"
	Method_Delete Method = "DELETE"
	Method_Post   Method = "POST"
)

const baseUrl = "https://discord.com/api/v10"

func Post(endpoint string, body interface{}) []byte {
	bBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("error marshaling body \n%s", err)
	}
	return request(Method_Post, endpoint, bytes.NewBuffer(bBytes))
}

func Get(endpoint string) []byte {
	return request(Method_Get, endpoint, nil)
}

func Delete(endpoint string) []byte {
	return request(Method_Delete, endpoint, nil)
}

func Put(endpoint string, body interface{}) []byte {
	bBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("error marshaling body \n%s", err)
	}
	return request(Method_Put, endpoint, bytes.NewBuffer(bBytes))
}

func request(method Method, endpoint string, body io.Reader) []byte {
	req, err := http.NewRequest(string(method), baseUrl+endpoint, body)

	if err != nil {
		log.Fatalf("error sending request \n%s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", os.Getenv("DISCORD_TOKEN")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error receiving response %s", err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("error reading response %s", err)
	}

	return bytes
}
