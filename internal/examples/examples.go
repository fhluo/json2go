package examples

import _ "embed"

var (
	//go:embed "data/JSON Example.json"
	jsonExample string
	//go:embed "data/Get the latest release.json"
	getTheLatestRelease string
	//go:embed "data/package.json"
	packageJSON string

	Examples = []Example{
		{
			Title:   "JSON Example",
			Content: jsonExample,
		},
		{
			Title:   "Github API Response",
			Content: getTheLatestRelease,
		},
		{
			Title:   "package.json",
			Content: packageJSON,
		},
	}
)

type Example struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
