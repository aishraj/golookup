package main

type searchResult interface {
	PackagePath() string
	Info() string
}

type goDocResult struct {
	Path     string `json:"path"`
	Synopsis string `json:"synopsis"`
}

type goDocResponse struct {
	Results []goDocResult `json:"results"`
}

type goSearchResponse struct {
	Query string           `json:"query"`
	Hits  []goSearchResult `json:"hits"`
}

type goSearchResult struct {
	Name        string `json:"name"`
	Package     string `json:"package"`
	Author      string `json:"author"`
	Synopsis    string `json:"synopsis"`
	Description string `json:"description"`
	ProjectURL  string `json:"projecturl"`
}

func (result goDocResult) PackagePath() string {
	return result.Path
}

func (result goDocResult) Info() string {
	return result.Synopsis
}

func (result goSearchResult) PackagePath() string {
	return result.ProjectURL
}

func (result goSearchResult) Info() string {
	return result.Synopsis
}
