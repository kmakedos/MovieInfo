package movies

type Movie struct {
	Adult        bool   `json:"adult"`
	BackdropPath string `json:"backdrop_path"`
	//GenreIds         []string `json:"genre_ids"`
	Id               string  `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

func NewMovie() *Movie {
	return &Movie{}
}
