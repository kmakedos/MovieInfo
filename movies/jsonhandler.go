package movies

import "encoding/json"

func JsonToMovie(jsonData []byte) (*Movie, error) {
	movie := NewMovie()
	err := json.Unmarshal(jsonData, movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
