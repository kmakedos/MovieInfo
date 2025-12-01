package movies

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func ParseFile(filename string, separator string) ([]string, error) {
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
		result := strings.Split(line, separator)
		results = append(results, result[len(result)-1])
	}
	return results, nil
}

func CharDistance(a, b uint8) int {
	diff := (a - b)
	if diff > 26 {
		diff = 26
	}
	return int(math.Abs(float64(diff)))
}

func CalcDistanceWord(needle, haystack string) int {
	distance := 0
	needle = strings.TrimSpace(needle)
	needle = strings.ToLower(needle)
	haystack = strings.TrimSpace(haystack)
	haystack = strings.ToLower(haystack)
	min := len(haystack)
	if len(needle) < len(haystack) {
		min = len(needle)
	}
	for i := 0; i < min; i++ {
		distance += CharDistance(haystack[i], needle[i])
	}
	return distance
}

func CalcDistancePhrase(needles, haystacks []string) int {
	min := len(haystacks)
	if len(needles) < min {
		min = len(needles)
	}
	for i := 0; i < min; i++ {
		fmt.Printf("%s %s\n", haystacks[i], needles[i])
	}
	return 0
}
