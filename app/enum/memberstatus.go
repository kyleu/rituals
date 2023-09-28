// Package enum - Content managed by Project Forge, see [projectforge.md] for details.
package enum

type MemberStatus string

const (
	MemberStatusOwner    MemberStatus = "owner"
	MemberStatusMember   MemberStatus = "member"
	MemberStatusObserver MemberStatus = "observer"
)
