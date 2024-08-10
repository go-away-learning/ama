package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.HelloHandler)

	return mux
}

func (s *Server) HelloHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World!"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		panic(fmt.Sprintf("error handling JSON marshal. Err: %v", err))
	}

	if _, err := w.Write(jsonResp); err != nil {
		fmt.Printf("error writing to HTTP. Err: %v\n", err)
	}
}
