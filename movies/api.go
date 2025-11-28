package movies

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func CallApi(title string, bearer string) []byte {

	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?query=%s&include_adult=true&language=en-US&page=1", title)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", bearer)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
