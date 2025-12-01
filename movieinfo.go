package main

import (
	"MovieInfo/movies"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type V struct {
	Results []movies.Movie `json:"results"`
}

func main() {
	if len(os.Args) < 4 {
		log.Println("Not enough arguments")
		log.Fatal("Usage: movieinfo <filename> <good|bad|not_found> <threshold>")
	}
	filename := os.Args[1]
	option := os.Args[2]
	threshold, err := strconv.ParseFloat(os.Args[3], 32)
	if err != nil {
		log.Println("Wrong threshold,setting default to 5")
		threshold = 5.0
	}
	found := make(map[string]movies.Movie)
	var goodRatings []movies.Movie
	var badRatings []movies.Movie
	var notFound []string
	var v V

	lines, err := movies.ParseFile(filename, "/")
	if err != nil {
		log.Fatal(err)
	}
	bearer := "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIzNzg2OWM3N2JkNDhhOGVmN2QwYjRiNTliYmI1YWI5OCIsIm5iZiI6MTc2Mzg4MzMxNS40NzgwMDAyLCJzdWIiOiI2OTIyYjkzM2U2YjNjYmQxMjIwYmEzZDciLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.OdtY11VY_j9yM5EtgJ9mlDShqDmao94M2WaEJps9ryw"
	for _, title := range lines {
		resp := movies.CallApi(title, bearer)
		err := json.Unmarshal(resp, &v)
		if err != nil {
			log.Printf("Error unmarshaling: %s %v\n", title, err)
		}
		if len(v.Results) > 0 {
			var maxPopularity float32
			var selectedMovie movies.Movie
			for _, result := range v.Results {
				if result.Popularity > maxPopularity {
					maxPopularity = result.Popularity
					selectedMovie = result
				}
			}
			selectedMovie.FilePath = title
			found[selectedMovie.Title] = selectedMovie
		} else {
			notFound = append(notFound, title)
		}
	}

	for _, mv := range found {
		if mv.VoteAverage > threshold {
			goodRatings = append(goodRatings, mv)

		} else {
			badRatings = append(badRatings, mv)
		}
	}
	switch option {
	case "good":
		for _, mv := range goodRatings {
			fmt.Println(mv.String())
		}
		fmt.Printf("Total good movies: %d\n", len(goodRatings))
	case "bad":
		for _, mv := range badRatings {
			fmt.Println(mv.String())
		}
		fmt.Printf("Total bad movies: %d\n", len(badRatings))
	case "not_found":
		for _, title := range notFound {
			fmt.Println(title)
		}
	}
	fmt.Println("Total", len(lines), "Movies checked")
	fmt.Printf("Found %d movies\n", len(found))
	fmt.Printf("Not found %d movies\n", len(notFound))

}
