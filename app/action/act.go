package action

type Act string

const (
	ActError  Act = "error"
	ActCreate Act = "create"

	ActUpdate      Act = "update"
	ActPermissions Act = "permissions"
	ActComment     Act = "comment"

	ActMemberSelf   Act = "self"
	ActMemberAdd    Act = "member-add"
	ActMemberUpdate Act = "member-update"
	ActMemberRemove Act = "member-remove"

	ActChildAdd    Act = "child-add"
	ActChildUpdate Act = "child-update"
	ActChildRemove Act = "child-remove"
	ActChildStatus Act = "child-status"

	ActVote Act = "vote"
)
