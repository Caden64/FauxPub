package main

type ObjectAndLinkTypes interface {
	Relationship | Article | Document | Audio | Image | Video | Note | Page |
		Event | Place | Mention | Profile | Tombstone
}
type Relationship struct {
	Object
	Subject
	Relationship  RelationshipType
	RelatedObject Object
}

type Article struct {
	Object
}

type Document struct {
	Object
}

type Audio struct {
	Document
}

type Image struct {
	Document
}

type Video struct {
	Document
}

type Note struct {
	Object
}

type Page struct {
	Object
}

type Event struct {
	Object
}

type Place struct {
	Object
	Accuracy
	Altitude
	Latitude
	Longitude
	Radius
	Units
}

type Mention struct {
	Link
}

type Profile struct {
	Object
}

type Tombstone struct {
	Object
	FormerType
	Deleted
}
