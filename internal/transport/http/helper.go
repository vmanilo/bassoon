package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getRequestData(r *http.Request, val any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}

	err = json.Unmarshal(body, val)
	if err != nil {
		return fmt.Errorf("failed to unmarshal request body: %w", err)
	}

	return nil
}

func sendResponse(writer http.ResponseWriter, val any) error {
	writer.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("failed to marshal port response: %w", err)
	}

	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("failed to write response data: %w", err)
	}

	return nil
}

type Error struct {
	Error string `json:"error"`
}

func httpError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := sendResponse(w, &Error{Error: err}); err != nil {
		log.Println("failed to send response:", err)
	}
}
