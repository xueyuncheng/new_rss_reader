package model

type Chart struct {
	Name  string  `json:"name"`
	Items []*Item `json:"items"`
}

type Item struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
