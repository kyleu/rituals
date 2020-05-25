package query

import "github.com/kyleu/rituals.dev/app/util"

var (
	allowedActionSortColumns     = []string{util.KeyID, util.KeySvc, "modelID", "authorID", "act", "content", "note", util.KeyCreated}
	allowedAuthSortColumns       = []string{util.KeyID, "userID", "provider", "providerID", "expires", "name", "email", "picture", util.KeyCreated}
	allowedEstimateSortColumns   = []string{util.KeyID, "slug", "title", "sprintID", "owner", util.KeyStatus, "choices", "options", util.KeyCreated}
	allowedInvitationSortColumns = []string{util.KeyKey, "k", "v", "src", "tgt", "note", util.KeyStatus, "redeemed", util.KeyCreated}
	allowedMemberSortColumns     = []string{"userID", "name", "role", util.KeyCreated}
	allowedPermissionSortColumns = []string{"k", "v", "access", util.KeyCreated}
	allowedReportSortColumns     = []string{util.KeyID, "standupID", "d", "authorID", "content", "html", util.KeyCreated}
	allowedRetroSortColumns      = []string{util.KeyID, "slug", "title", "sprintID", "owner", util.KeyStatus, "categories", util.KeyCreated}
	allowedSocketSortColumns     = []string{util.KeySvc, "cmd", "param"}
	allowedSprintSortColumns     = []string{util.KeyID, "slug", "title", "owner", "startDate", "endDate", util.KeyCreated}
	allowedStandupSortColumns    = []string{util.KeyID, "slug", "title", "sprintID", "owner", util.KeyStatus, util.KeyCreated}
	allowedStorySortColumns      = []string{util.KeyID, "estimateID", "idx", "authorID", "title", util.KeyStatus, "finalVote", util.KeyCreated}
	allowedTeamSortColumns       = []string{util.KeyID, "slug", "title", "owner", util.KeyCreated}
	allowedUserSortColumns       = []string{util.KeyID, "name", "role", util.KeyTheme, "navColor", "linkColor", "picture", "locale", util.KeyCreated}
	allowedVoteSortColumns       = []string{"storyID", "userID", "choice", "updated", util.KeyCreated}
)

var allowedColumns = map[string][]string{
	util.KeyAction:       allowedActionSortColumns,
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
