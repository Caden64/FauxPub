package main

type Object struct {
	Context      string `json:"@context"`
	Type         string `json:"type"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	Attachment   string `json:"attachment"`
	AttributedTo string `json:"attributedTo"`
	Audience     string `json:"audience"`
	Content      string `json:"content"`
	EndTime      string `json:"endTime"`
	Generator    string `json:"generator"`
	Icon         string `json:"icon"`
	Image        string `json:"image"`
	InReplyTo    string `json:"inReplyTo"`
	Location     string `json:"location"`
	Preview      string `json:"preview"`
	Published    string `json:"published"`
	Replies      string `json:"replies"`
	StartTime    string `json:"startTime"`
	Summary      string `json:"summary"`
	Tag          string `json:"tag"`
	Updated      string `json:"updated"`
	URL          string `json:"URL"`
	To           string `json:"to"`
	BTo          string `json:"BTo"`
	CC           string `json:"CC"`
	BCC          string `json:"BCC"`
	MediaType    string `json:"mediaType"`
	Duration     string `json:"duration"`
}

type Link struct {
	Context   string
	Type      string
	Href      string
	Hreflang  string
	MediaType string
	Name      string
	Height    string
	Width     string
	Preview   string
	Rel       string
}

type Activity struct {
	Object
	Actor      string
	Target     string
	Result     string
	Origin     string
	Instrument string
}

type IntransitiveActivity struct {
	Activity
}

type Collection struct {
	Object
	TotalItems string
	Current    string
	First      string
	Last       string
	Items      string
}

type OrderedCollection struct {
	Collection
}

type CollectionPage struct {
	Collection
	PartOf string
	Next   string
	Prev   string
}

type OrderedCollectionPage struct {
	CollectionPage
	StartIndex string
}
