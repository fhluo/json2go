package examples

import _ "embed"

var (
	//go:embed "data/JSON Example.json"
	jsonExample string
	//go:embed "data/Get the latest release.json"
	getTheLatestRelease string

	Examples = []Example{
		{
			Title:   "JSON Example",
			Content: jsonExample,
		},
		{
			Title:   "Github API Response",
			Content: getTheLatestRelease,
		},
	}
)

type Example struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
