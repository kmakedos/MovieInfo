package movies

import (
	"bufio"
	"os"
	"strings"
)

func ParseFile(filename string) ([]string, error) {
	var results []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		result := strings.Split(line, "/")
		results = append(results, result[len(result)-1])
	}
	return results, nil
}
