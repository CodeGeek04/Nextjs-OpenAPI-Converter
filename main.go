package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriteToFile(data []byte, filename string) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

type PostmanInfo struct {
	PostmanId  string `json:"_postman_id,omitempty"`
	Name       string `json:"name"`
	Schema     string `json:"schema,omitempty"`
	ExporterId string `json:"_exporter_id,omitempty"`
}

type PostmanCollection struct {
	Info PostmanInfo `json:"info"`
	Item []Item      `json:"item"`
}

type Item struct {
	Name     string   `json:"name,omitempty"`
	Request  *Request `json:"request,omitempty"`
	Response []string `json:"response,omitempty"`
	Item     []Item   `json:"item,omitempty"` // For folders containing more items
}

type Request struct {
	Method string   `json:"method"`
	Header []Header `json:"header"`
	Body   Body     `json:"body"`
	URL    URL      `json:"url"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Body struct {
	Mode    string      `json:"mode"`
	Raw     string      `json:"raw"`
	Options BodyOptions `json:"options"`
}

type BodyOptions struct {
	Raw RawOptions `json:"raw"`
}

type RawOptions struct {
	Language string `json:"language"`
}

type URL struct {
	Raw      string   `json:"raw"`
	Protocol string   `json:"protocol"`
	Host     []string `json:"host"`
	Port     string   `json:"port"`
	Path     []string `json:"path"`
	Query    []Query  `json:"query"`
}

type Query struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	header := Header{
		Key:   "Authrization",
		Value: "Bearer fsejkfnke",
		Type:  "text",
	}

	bodyRawOptions := RawOptions{
		Language: "json",
	}

	bodyOptions := BodyOptions{
		Raw: bodyRawOptions,
	}

	body := Body{
		Mode:    "raw",
		Raw:     "{\n  \"recordType\": \"ns\",\n  \"value\": \"ns1.example.com\"\n}",
		Options: bodyOptions,
	}

	url := URL{
		Raw:      "http://localhost:3000/dns/shivammittal.in",
		Protocol: "http",
		Host:     []string{"localhost"},
		Port:     "3000",
		Path:     []string{"dns", "shivammitta.in"},
	}

	req := Request{
		Method: "POST",
		Header: []Header{header},
		Body:   body,
		URL:    url,
	}

	requestItem := Item{
		Name:     "new request",
		Request:  &req,
		Response: []string{},
	}

	subfolderRequestItem := Item{
		Name:     "subfolder request",
		Request:  &req,
		Response: []string{},
	}

	subfolderItem := Item{
		Name: "sub folder 1",
		Item: []Item{subfolderRequestItem},
	}

	postmanInfo := PostmanInfo{
		Name:   "Nextjs to API Collection",
		Schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
	}

	postmanCollection := PostmanCollection{
		Info: postmanInfo,
		Item: []Item{requestItem, subfolderItem},
	}

	indentedMarshal, err := json.MarshalIndent(postmanCollection, "", "    ")
	checkError(err)

	err = WriteToFile(indentedMarshal, "postmanSchema.json")
	checkError(err)
}
