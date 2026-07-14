package main

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Location struct {
	ID     int              `json:"id"`
	Name   string           `json:"name"`
	Region NamedAPIResource `json:"region"`

	Names []struct {
		Name     string           `json:"name"`
		Language NamedAPIResource `json:"language"`
	} `json:"names"`

	GameIndices []struct {
		GameIndex  int              `json:"game_index"`
		Generation NamedAPIResource `json:"generation"`
	} `json:"game_indices"`

	Areas []NamedAPIResource `json:"areas"`
}