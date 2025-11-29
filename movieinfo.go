package main

import (
	"MovieInfo/movies"
	"encoding/json"
	"fmt"
	"log"
)

type V struct {
	Results []movies.Movie `json:"results"`
}

//	func PrintMovieInfo() {
//		if len(v.Results) > 0 {
//			mv := v.Results[0]
//			fmt.Printf("%s %s\n", mv.Title, mv.ReleaseDate)
//			fmt.Println(mv.Overview)
//			fmt.Println(mv.Popularity, mv.VoteAverage)
//			fmt.Println("================================")
//		}
//	}
func main() {
	var v V
	var found []movies.Movie
	var notFound []string
	filename := "data/movies-02.txt"
	lines, err := movies.ParseFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	bearer := "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIzNzg2OWM3N2JkNDhhOGVmN2QwYjRiNTliYmI1YWI5OCIsIm5iZiI6MTc2Mzg4MzMxNS40NzgwMDAyLCJzdWIiOiI2OTIyYjkzM2U2YjNjYmQxMjIwYmEzZDciLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.OdtY11VY_j9yM5EtgJ9mlDShqDmao94M2WaEJps9ryw"
	for _, title := range lines {
		resp := movies.CallApi(title, bearer)
		err := json.Unmarshal(resp, &v)
		if err != nil {
			log.Printf("Error unmarshaling: %v\n", err)
		}
		if len(v.Results) > 0 {
			found = append(found, v.Results[0])
		} else {
			notFound = append(notFound, title)
		}
	}
	fmt.Println("Found", len(found), "Movies")
	fmt.Println("Not Found", len(notFound), "Movies")
	for _, title := range notFound {
		fmt.Println(title)
	}
}
