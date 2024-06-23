package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// Register HTTP handlers
	http.HandleFunc("/echo", handleMatrixOperation(echoMatrix))
	http.HandleFunc("/invert", handleMatrixOperation(invertMatrix))
	http.HandleFunc("/flatten", handleMatrixOperation(flattenMatrix))
	http.HandleFunc("/sum", handleMatrixOperation(sumMatrix))
	http.HandleFunc("/multiply", handleMatrixOperation(multiplyMatrix))

	// Start the server
	port := ":8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Processes the matrix from the request and calls the operation function.
func handleMatrixOperation(operation func(matrix [][]int) interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		matrix, err := parseCSV(file)
		if err != nil {
			http.Error(w, "Invalid Character in CSV "+err.Error(), http.StatusBadRequest)
			return
		}

		result := operation(matrix)
		fmt.Fprintf(w, "%v\n", result)
	}
}

// parseCSV parses the CSV file into a matrix of integers.
func parseCSV(file multipart.File) ([][]int, error) {
	reader := csv.NewReader(file)
	var matrix [][]int
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		var nums []int
		for _, cell := range row {
			num, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		matrix = append(matrix, nums)
	}
	return matrix, nil
}

// Matrix operation functions

func echoMatrix(matrix [][]int) interface{} {
	return matrix
}

func invertMatrix(matrix [][]int) interface{} {
	n := len(matrix)
	inverted := make([][]int, n)
	for i := 0; i < n; i++ {
		inverted[i] = make([]int, n)
		for j := 0; j < n; j++ {
			inverted[i][j] = matrix[j][i]
		}
	}
	return inverted
}

func flattenMatrix(matrix [][]int) interface{} {
	var nums []string
	for _, row := range matrix {
		for _, num := range row {
			nums = append(nums, strconv.Itoa(num))
		}
	}
	return strings.Join(nums, ",")
}

func sumMatrix(matrix [][]int) interface{} {
	sum := 0
	for _, row := range matrix {
		for _, num := range row {
			sum += num
		}
	}
	return sum
}

func multiplyMatrix(matrix [][]int) interface{} {
	product := 1
	for _, row := range matrix {
		for _, num := range row {
			product *= num
		}
	}
	return product
}
