package main

type Object struct {
	Context      string             `json:"@context"`
	Type         ActivityStreamType `json:"type"`
	Id           `json:"id"`
	Name         `json:"name"`
	Attachment   `json:"attachment"`
	AttributedTo `json:"attributedTo"`
	Audience     `json:"audience"`
	Content      `json:"content"`
	EndTime      `json:"endTime"`
	Generator    `json:"generator"`
	Icon         `json:"icon"`
	Image        ImageType `json:"image"`
	InReplyTo    `json:"inReplyTo"`
	Location     `json:"location"`
	Preview      `json:"preview"`
	Published    `json:"published"`
	Replies      `json:"replies"`
	StartTime    `json:"startTime"`
	Summary      `json:"summary"`
	Tag          `json:"tag"`
	Updated      `json:"updated"`
	URL          `json:"URL"`
	To           `json:"to"`
	BTo          `json:"BTo"`
	CC           `json:"CC"`
	BCC          `json:"BCC"`
	MediaType    `json:"mediaType"`
	Duration     `json:"duration"`
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
