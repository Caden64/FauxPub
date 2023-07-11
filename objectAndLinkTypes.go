package main

type Relationship struct {
	Object
	Subject      string
	Relationship string
	Object2      string
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
	Accuracy string
	Altitude string
}