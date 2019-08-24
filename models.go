package main

//Data type to parse POST request to json
type Data struct {
	Val string `json:"val"`
}

//ResponseData struct parsed to json
type ResponseData struct {
	ShortURL    string `json:"shortURL"`
	OriginalURL string `json:"originalURL"`
}

//URLList List of URL's
type URLList map[string]string
