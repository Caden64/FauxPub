package main

type Object struct {
	Context      `json:"@context"`
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
	Context   `json:"@context"`
	Type      ActivityStreamType `json:"type"`
	Href      `json:"href"`
	HrefLang  `json:"hrefLang"`
	MediaType `json:"mediaType"`
	Name      `json:"name"`
	Height    `json:"height"`
	Width     `json:"width"`
	Preview   `json:"preview"`
	Rel       `json:"rel"`
}

type Activity struct {
	Object     `json:"object"`
	Actor      `json:"actor"`
	Target     `json:"target"`
	Result     `json:"result"`
	Origin     `json:"origin"`
	Instrument `json:"instrument"`
}

type IntransitiveActivity struct {
	Activity `json:"activity"`
}

type Collection struct {
	Object     `json:"object"`
	TotalItems `json:"totalItems"`
	Current    `json:"current"`
	First      `json:"first"`
	Last       `json:"last"`
	Items      `json:"items"`
}

type OrderedCollection struct {
	Collection `json:"collection"`
}

type CollectionPage struct {
	Collection `json:"collection"`
	PartOf     `json:"partOf"`
	Next       `json:"next"`
	Prev       `json:"prev"`
}

type OrderedCollectionPage struct {
	CollectionPage `json:"collectionPage"`
	StartIndex     `json:"startIndex"`
}
