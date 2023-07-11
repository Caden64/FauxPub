package main

import "fmt"

func main() {
	x := OrderedCollectionPage{
		CollectionPage: CollectionPage{
			Collection: Collection{
				Object: Object{
					Context:      "",
					Type:         "",
					Id:           "",
					Name:         "",
					Attachment:   "",
					AttributedTo: "",
					Audience:     "",
					Content:      "",
					EndTime:      "",
					Generator:    "",
					Icon:         "",
					Image:        "",
					InReplyTo:    "",
					Location:     "",
					Preview:      "",
					Published:    "",
					Replies:      "",
					StartTime:    "",
					Summary:      "",
					Tag:          "",
					Updated:      "",
					URL:          "",
					To:           "",
					BTo:          "",
					CC:           "",
					BCC:          "",
					MediaType:    "",
					Duration:     "",
				},
				TotalItems: "",
				Current:    "",
				First:      "",
				Last:       "",
				Items:      "",
			},
			PartOf: "",
			Next:   "",
			Prev:   "",
		},
		StartIndex: "",
	}
	fmt.Println(x.TotalItems)
}
