package query

import "github.com/kyleu/rituals.dev/app/util"

var (
	allowedActionSortColumns     = []string{util.KeyID, util.KeySvc, "modelID", "authorID", "act", "content", "note", util.KeyCreated}
	allowedAdminSortColumns      = []string{util.KeyName, "recent", "count"}
	allowedAuthSortColumns       = []string{util.KeyID, util.WithID(util.KeyUser), "provider", "providerID", "expires", util.KeyName, "email", "picture", util.KeyCreated}
	allowedEstimateSortColumns   = []string{util.KeyID, util.KeySlug, util.KeyTitle, util.WithID(util.SvcSprint.Key), util.KeyOwner, util.KeyStatus, "choices", "options", util.KeyCreated}
	allowedInvitationSortColumns = []string{util.KeyKey, "k", "v", "src", "tgt", "note", util.KeyStatus, "redeemed", util.KeyCreated}
	allowedMemberSortColumns     = []string{util.WithID(util.KeyUser), util.KeyName, util.KeyRole, util.KeyCreated}
	allowedPermissionSortColumns = []string{"k", "v", "access", util.KeyCreated}
	allowedReportSortColumns     = []string{util.KeyID, "standupID", "d", "authorID", "content", "html", util.KeyCreated}
	allowedRetroSortColumns      = []string{util.KeyID, util.KeySlug, util.KeyTitle, util.WithID(util.SvcSprint.Key), util.KeyOwner, util.KeyStatus, "categories", util.KeyCreated}
	allowedSocketSortColumns     = []string{util.KeySvc, "cmd", "param"}
	allowedSprintSortColumns     = []string{util.KeyID, util.KeySlug, util.KeyTitle, util.KeyOwner, "startDate", "endDate", util.KeyCreated}
	allowedStandupSortColumns    = []string{util.KeyID, util.KeySlug, util.KeyTitle, util.WithID(util.SvcSprint.Key), util.KeyOwner, util.KeyStatus, util.KeyCreated}
	allowedStorySortColumns      = []string{util.KeyID, "estimateID", util.KeyIdx, "authorID", util.KeyTitle, util.KeyStatus, "finalVote", util.KeyCreated}
	allowedTeamSortColumns       = []string{util.KeyID, util.KeySlug, util.KeyTitle, util.KeyOwner, util.KeyCreated}
	allowedUserSortColumns       = []string{util.KeyID, util.KeyName, util.KeyRole, util.KeyTheme, "navColor", "linkColor", "picture", "locale", util.KeyCreated}
	allowedVoteSortColumns       = []string{"storyID", util.WithID(util.KeyUser), "choice", "updated", util.KeyCreated}
)

var allowedColumns = map[string][]string{
	util.KeyAction:       allowedActionSortColumns,
	util.KeyAdmin:        allowedAdminSortColumns,
	util.KeyAuth:         allowedAuthSortColumns,
	util.SvcEstimate.Key: allowedEstimateSortColumns,
	util.KeyInvitation:   allowedInvitationSortColumns,
	util.KeyMember:       allowedMemberSortColumns,
	util.KeyPermission:   allowedPermissionSortColumns,
	util.KeyReport:       allowedReportSortColumns,
	util.SvcRetro.Key:    allowedRetroSortColumns,
	util.KeySocket:       allowedSocketSortColumns,
	util.SvcSprint.Key:   allowedSprintSortColumns,
	util.SvcStandup.Key:  allowedStandupSortColumns,
	util.KeyStory:        allowedStorySortColumns,
	util.KeyUser:         allowedUserSortColumns,
	util.SvcTeam.Key:     allowedTeamSortColumns,
	util.KeyVote:         allowedVoteSortColumns,
}
