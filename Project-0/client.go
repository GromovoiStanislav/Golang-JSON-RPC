// client.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func jsonRPCRequest(method string, params []interface{}) (interface{}, error) {
	request := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1, // Уникальный идентификатор запроса
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:3000/jsonrpc", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	result, ok := response["result"]
	if !ok {
		return nil, fmt.Errorf("Error: %v", response["error"])
	}

	return result, nil
}

func main() {
	result, err := jsonRPCRequest("add", []interface{}{2, 3, 5})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", result)
	}
}
