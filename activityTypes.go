package main

type ActivityTypes interface {
	Accept | TentativeAccept | Add | Arrive | Create | Delete | Follow |
		Ignore | Join | Leave | Like | Offer | Invite | Reject | TentativeReject |
		Remove | Update | View | Listen | Read | Move | Travel | Announce | Dislike |
		Question | Flag | Block
}
type Accept struct {
	Activity
}

type TentativeAccept struct {
	Activity
}

type Add struct {
	Activity
}

type Arrive struct {
	Activity
}

type Create struct {
	Activity
}

type Delete struct {
	Activity
}

type Follow struct {
	Activity
}

type Ignore struct {
	Activity
}

type Join struct {
	Activity
}

type Leave struct {
	Activity
}

type Like struct {
	Activity
}

type Offer struct {
	Activity
}

type Invite struct {
	Activity
}

type Reject struct {
	Activity
}

type TentativeReject struct {
	Activity
}

type Remove struct {
	Activity
}

type Update struct {
	Activity
}

type View struct {
	Activity
}

type Listen struct {
	Activity
}

type Read struct {
	Activity
}

type Move struct {
	Activity
}

type Travel struct {
	IntransitiveActivity
}

type Announce struct {
	Activity
}

type Block struct {
	Ignore
}

type Flag struct {
	Activity
}

type Dislike struct {
	Activity
}

type Question struct {
	IntransitiveActivity
	OneOf  string
	AnyOf  string
	Closed string
}
