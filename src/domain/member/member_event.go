package domain

type (
	MemberCreated struct {
		Id string
	}
	MemberStateChanged struct {
		Id    string
		State MemberState
	}
)
