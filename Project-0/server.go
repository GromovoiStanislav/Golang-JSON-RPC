package main

import (
	"encoding/json"
	"net/http"
)

func handleJSONRPCRequest(w http.ResponseWriter, r *http.Request) {
	
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	method, ok := request["method"].(string)
	if !ok {
		http.Error(w, "Method not provided", http.StatusBadRequest)
		return
	}

	var result interface{}
	switch method {
	case "add":
		params, ok := request["params"].([]interface{})
		if !ok {
			http.Error(w, "Invalid parameters", http.StatusBadRequest)
			return
		}

		sum := 0.0
		for _, param := range params {
			if num, ok := param.(float64); ok {
				sum += num
			} else {
				http.Error(w, "Invalid parameter type", http.StatusBadRequest)
				return
			}
		}

		result = sum
	default:
		http.Error(w, "Method not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"jsonrpc": "2.0",
		"result":  result,
		"id":      request["id"],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/jsonrpc", handleJSONRPCRequest)
	http.ListenAndServe(":3000", nil)
}
