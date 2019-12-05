package search

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json: "link"`
	Type string `json: "type"`
}
