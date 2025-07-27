package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{Port: port}
}

type Response struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

func (s *Server) Start() {
	http.HandleFunc("/add", s.handleAdd)
	http.HandleFunc("/sub", s.handleSub)

	err := http.ListenAndServe(":"+s.Port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func (s *Server) handleAdd(w http.ResponseWriter, r *http.Request) {
	s.handleOperation(w, r, "add")
}

func (s *Server) handleSub(w http.ResponseWriter, r *http.Request) {
	s.handleOperation(w, r, "sub")
}

func (s *Server) handleOperation(w http.ResponseWriter, r *http.Request, op string) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, Response{"", "Invalid HTTP method"})
		return
	}

	numbersParam := r.URL.Query().Get("numbers")
	if numbersParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, Response{"", "'numbers' parameter missing"})
		return
	}

	numsStr := strings.Split(numbersParam, ",")
	if len(numsStr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, Response{"", "'numbers' parameter missing"})
		return
	}

	numbers := make([]int64, 0, len(numsStr))
	for _, numStr := range numsStr {
		numStr = strings.TrimSpace(numStr)
		if numStr == "" {
			continue
		}
		num, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeJSON(w, Response{"", "Invalid number: " + numStr})
			return
		}
		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, Response{"", "'numbers' parameter missing"})
		return
	}

	var result int64
	var overflow bool

	switch op {
	case "add":
		result, overflow = safeAdd(numbers)
	case "sub":
		result, overflow = safeSub(numbers)
	default:
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, Response{"", "Invalid operation"})
		return
	}

	if overflow {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, Response{"", "Overflow"})
		return
	}

	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("The result of your query is: %d", result)
	writeJSON(w, Response{msg, ""})
}

func safeAdd(numbers []int64) (int64, bool) {
	var sum int64 = 0
	for _, n := range numbers {
		if willAddOverflow(sum, n) {
			return 0, true
		}
		sum += n
	}
	return sum, false
}

func safeSub(numbers []int64) (int64, bool) {
	if len(numbers) == 0 {
		return 0, false
	}
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if willSubOverflow(result, numbers[i]) {
			return 0, true
		}
		result -= numbers[i]
	}
	return result, false
}

func willAddOverflow(a, b int64) bool {
	if b > 0 && a > math.MaxInt64-b {
		return true
	}
	if b < 0 && a < math.MinInt64-b {
		return true
	}
	return false
}

func willSubOverflow(a, b int64) bool {
	if b == math.MinInt64 {
		return true
	}
	negB := -b
	return willAddOverflow(a, negB)
}

func writeJSON(w http.ResponseWriter, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	s := NewServer("8000")
	s.Start()
}
