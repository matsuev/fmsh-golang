package csvstorage

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// CsvStorage ...
type CsvStorage struct {
	headers []string
	data    [][]string
	pos     int
	tmp     []string
}

// Create function
func Create(path string, sep string) (*CsvStorage, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	if !scanner.Scan() {
		return nil, errors.New("empty file")
	}

	result := &CsvStorage{
		pos: 0,
	}

	headLine := strings.TrimPrefix(scanner.Text(), "\ufeff")
	result.headers = strings.Split(headLine, sep)

	for scanner.Scan() {
		items := strings.Split(scanner.Text(), sep)
		result.data = append(result.data, items)
	}

	return result, nil
}

// GetLineArray function
func (s *CsvStorage) GetLineArray(n int) ([]string, error) {
	if n < 0 || n >= len(s.data) {
		return nil, errors.New("index out of range")
	}
	return s.data[n], nil
}

// GetLineMap function
func (s *CsvStorage) GetLineMap(n int) (map[string]string, error) {
	if n < 0 || n >= len(s.data) {
		return nil, errors.New("index out of range")
	}

	result := make(map[string]string)

	for index, value := range s.data[n] {
		result[s.headers[index]] = value
	}

	return result, nil
}

// GetHeaders function
func (s *CsvStorage) GetHeaders() []string {
	return s.headers
}

// Len function
func (s *CsvStorage) Len() int {
	return len(s.data)
}

// Scan function
func (s *CsvStorage) Scan() bool {
	if s.pos >= s.Len() {
		return false
	}

	s.tmp = s.data[s.pos]
	s.pos++

	return true
}

// LineArray function
func (s *CsvStorage) LineArray() []string {
	return s.tmp
}

// LineMap function
func (s *CsvStorage) LineMap() map[string]string {
	result := make(map[string]string)

	for index, value := range s.tmp {
		result[s.headers[index]] = value
	}

	return result
}
