package pokeapi

type LocationsList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"` // if null in case page1 returns nul, pointer so we make sure we are actually dynamically keeping track of current page .
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
