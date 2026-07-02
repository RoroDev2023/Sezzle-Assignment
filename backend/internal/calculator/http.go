package calculator

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Server struct{}

type calculateRequest struct {
	Operation Operation `json:"operation"`
	A         *float64  `json:"a"`
	B         *float64  `json:"b"`
}

type calculateResponse struct {
	Operation Operation `json:"operation"`
	Result    float64   `json:"result"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /api/operations", s.operations)
	mux.HandleFunc("POST /api/calculate", s.calculate)

	return withCORS(mux)
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) operations(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string][]Operation{"operations": Operations()})
}

func (s *Server) calculate(w http.ResponseWriter, r *http.Request) {
	var request calculateRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "request body must be valid JSON")
		return
	}

	if request.A == nil {
		writeError(w, http.StatusBadRequest, "field a is required")
		return
	}

	if RequiresSecondOperand(request.Operation) && request.B == nil {
		writeError(w, http.StatusBadRequest, "field b is required for this operation")
		return
	}

	result, err := Calculate(Calculation{
		Operation: request.Operation,
		A:         *request.A,
		B:         request.B,
	})
	if err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, ErrInvalidOperation) {
			status = http.StatusNotFound
		}
		writeError(w, status, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, calculateResponse{
		Operation: request.Operation,
		Result:    result,
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

