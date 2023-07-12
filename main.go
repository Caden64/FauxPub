package main

import "fmt"

func main() {
	q := Question{
		IntransitiveActivity: IntransitiveActivity{
			Activity: Activity{
				Object: Object{
					Context: Context{
						Context: "https://www.w3.org/ns/activitystreams",
					},
					Type: ActivityStreamType{
						ActivityStreamType: "Question",
					},
					Id: Id{
						Id: "http://polls.example.org/question/1",
					},
					Name: Name{
						Name: "A question about robots",
					},
					Content: Content{
						Context: "I'd like to build a robot to feed my cat. Which platform is best?",
					},
					Replies: Replies{
						Replies: Replies{
							Replies: Collection[Object]{
								Object: Object{
									Type: ActivityStreamType{"Collection"},
								},
								TotalItems: TotalItems{3},
								Items: []Items{{Items: Object{
									AttributedTo: AttributedTo{AttributedTo: "http://sally.example.org"},
									InReplyTo:    InReplyTo{InReplyTo: "http://polls.example.org/question/1"},
									Name:         Name{"arduino"},
								}}, {Items: Object{
									AttributedTo: AttributedTo{AttributedTo: "http://joe.example.org"},
									InReplyTo:    InReplyTo{InReplyTo: "http://polls.example.org/question/1"},
									Name:         Name{"arduino"},
								}},
									{Items: Object{
										AttributedTo: AttributedTo{AttributedTo: "http://john.example.org"},
										InReplyTo:    InReplyTo{InReplyTo: "http://polls.example.org/question/1"},
										Name:         Name{"rasberry pi"},
									}},
								},
							},
						},
					},
				},
				Result: Result{Object{
					Type:    ActivityStreamType{"Note"},
					Content: Content{"Users are favorting &amp;quot;arduino&amp;quot; by a 33% margin"},
				}},
			},
		},
		OneOf: OneOf{OneOf: []string{"arduino", "rasberry pi"}},
	}
	fmt.Println(q)
}
