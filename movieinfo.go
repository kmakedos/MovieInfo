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

func main() {
	bearer := "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIzNzg2OWM3N2JkNDhhOGVmN2QwYjRiNTliYmI1YWI5OCIsIm5iZiI6MTc2Mzg4MzMxNS40NzgwMDAyLCJzdWIiOiI2OTIyYjkzM2U2YjNjYmQxMjIwYmEzZDciLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.OdtY11VY_j9yM5EtgJ9mlDShqDmao94M2WaEJps9ryw"
	title := "King Kong"
	resp := movies.CallApi(title, bearer)
	var v V
	err := json.Unmarshal(resp, &v)
	if err != nil {
		log.Printf("Error unmarshaling: %v\n", err)
	}
	for _, mv := range v.Results {
		fmt.Printf("%s %s\n", mv.Title, mv.ReleaseDate)
		fmt.Println(mv.Overview)
		fmt.Println("================================")
	}
}
