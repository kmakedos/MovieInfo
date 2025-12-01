package main

import (
	"MovieInfo/movies"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: compare <titles_file> <movies_file>")
	}
	titlesFilename := os.Args[1]
	moviesFilename := os.Args[2]
	mv, err := movies.ParseFile(moviesFilename, "\n")
	if err != nil {
		log.Fatal(err)
	}
	titles, err := movies.ParseFile(titlesFilename, "\n")
	if err != nil {
		log.Fatal(err)

	}
	for _, mvline := range mv {
		fmt.Println(mvline)
	}
	for _, titleline := range titles {
		fmt.Println(titleline)
	}
}
