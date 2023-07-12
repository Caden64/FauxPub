package main

type Base interface {
	Object | Link
}
type Object struct {
	Context      `json:"@context,omitempty"`
	Type         ActivityStreamType `json:"type,omitempty"`
	Id           `json:"id,omitempty"`
	Name         `json:"name,omitempty"`
	Attachment   `json:"attachment,omitempty"`
	AttributedTo `json:"attributedTo,omitempty"`
	Audience     `json:"audience,omitempty"`
	Content      `json:"content,omitempty"`
	EndTime      `json:"endTime,omitempty"`
	Generator    `json:"generator,omitempty"`
	Icon         `json:"icon,omitempty"`
	Image        ImageType `json:"image,omitempty"`
	InReplyTo    `json:"inReplyTo,omitempty"`
	Location     `json:"location,omitempty"`
	Preview      `json:"preview,omitempty"`
	Published    `json:"published,omitempty"`
	Replies      `json:"replies,omitempty"`
	StartTime    `json:"startTime,omitempty"`
	Summary      `json:"summary,omitempty"`
	Tag          `json:"tag,omitempty"`
	Updated      `json:"updated,omitempty"`
	URL          `json:"URL,omitempty"`
	To           `json:"to,omitempty"`
	BTo          `json:"BTo,omitempty"`
	CC           `json:"CC,omitempty"`
	BCC          `json:"BCC,omitempty"`
	MediaType    `json:"mediaType,omitempty"`
	Duration     `json:"duration,omitempty"`
}

type Link struct {
	Context   `json:"@context,omitempty"`
	Type      ActivityStreamType `json:"type,omitempty"`
	Href      `json:"href,omitempty"`
	HrefLang  `json:"hrefLang,omitempty"`
	MediaType `json:"mediaType,omitempty"`
	Name      `json:"name,omitempty"`
	Height    `json:"height,omitempty"`
	Width     `json:"width,omitempty"`
	Preview   `json:"preview,omitempty"`
	Rel       `json:"rel,omitempty"`
}

type Activity struct {
	Object     `json:"object,omitempty"`
	Actor      `json:"actor,omitempty"`
	Target     `json:"target,omitempty"`
	Result     `json:"result,omitempty"`
	Origin     `json:"origin,omitempty"`
	Instrument `json:"instrument,omitempty"`
}

type IntransitiveActivity struct {
	Activity `json:"activity,omitempty"`
}

type Collection[b Base] struct {
	PossibleLink b `json:"object,omitempty"`
	Object
	TotalItems `json:"totalItems,omitempty"`
	Current    `json:"current,omitempty"`
	First      `json:"first,omitempty"`
	Last       `json:"last,omitempty"`
	Items      []Items `json:"items,omitempty"`
}

type OrderedCollection[b Base] struct {
	Collection[b] `json:"collection,omitempty"`
}

type CollectionPage[b Base] struct {
	Collection[b] `json:"collection,omitempty"`
	PartOf        `json:"partOf,omitempty"`
	Next          `json:"next,omitempty"`
	Prev          `json:"prev,omitempty"`
}

type OrderedCollectionPage[b Base] struct {
	CollectionPage[b] `json:"collectionPage,omitempty"`
	StartIndex        `json:"startIndex,omitempty"`
}
