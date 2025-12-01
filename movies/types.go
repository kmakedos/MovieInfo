package movies

import (
	"bytes"
	"html/template"
	"log"
)

type Movie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	Id               string  `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	FilePath         string
}

func NewMovie() *Movie {
	return &Movie{}
}

type MovieFilter struct {
	Popularity  float64
	VoteAverage float64
}

func (m *Movie) String() string {
	const response = `
Title: {{.Title}}
Release: {{.ReleaseDate}}
Overview: {{.Overview}}
Popularity: {{.Popularity}}
VoteAverage: {{.VoteAverage}}
FilePath: {{.FilePath}}
`
	t := template.Must(template.New("movie").Parse(response))
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, m); err != nil {
		log.Println(err)
	}
	return tpl.String()
}
