package query

import "github.com/kyleu/rituals.dev/app/util"

var (
	allowedActionSortColumns     = []string{util.KeyID, util.KeySvc, util.WithID(util.KeyModel), util.WithID(util.KeyUser), "act", util.KeyContent, util.KeyNote, util.KeyCreated}
	allowedAdminSortColumns      = []string{util.KeyName, "recent", "count"}
	allowedAuthSortColumns       = []string{util.KeyID, util.WithID(util.KeyUser), util.KeyProvider, util.WithID(util.KeyProvider), "expires", util.KeyName, util.KeyEmail, "picture", util.KeyCreated}
	allowedCommentSortColumns    = []string{util.KeyID, util.KeySvc, util.WithID(util.KeyModel), util.WithID(util.KeyUser), util.KeyContent, util.KeyHTML, util.KeyCreated}
	allowedEstimateSortColumns   = []string{util.KeyID, util.KeySlug, util.KeyTitle, util.WithID(util.SvcSprint.Key), util.KeyOwner, util.KeyStatus, util.Plural(util.KeyChoice), "options", util.KeyCreated}
	allowedFeedbackSortColumns   = []string{util.KeyID, util.WithID(util.SvcRetro.Key), util.KeyIdx, util.WithID(util.KeyUser), util.KeyCategory, util.KeyContent, util.KeyHTML, util.KeyCreated}
	allowedHistorySortColumns    = []string{util.KeySlug, util.WithID(util.KeyModel), "modelName", util.KeyCreated}
	allowedMemberSortColumns     = []string{util.WithID(util.KeyUser), util.KeyName, util.KeyRole, util.KeyCreated}
	allowedPermissionSortColumns = []string{"k", "v", "access", util.KeyCreated}
	allowedReportSortColumns     = []string{util.KeyID, util.WithID(util.SvcStandup.Key), "d", util.WithID(util.KeyUser), util.KeyContent, util.KeyHTML, util.KeyCreated}
	allowedRetroSortColumns      = []string{util.KeyID, util.KeySlug, util.KeyTitle, util.WithID(util.SvcSprint.Key), util.KeyOwner, util.KeyStatus, util.Plural(util.KeyCategory), util.KeyCreated}
	allowedSocketSortColumns     = []string{util.KeySvc, "cmd", "param"}
	allowedSprintSortColumns     = []string{util.KeyID, util.KeySlug, util.KeyTitle, util.KeyOwner, "startDate", "endDate", util.KeyCreated}
	allowedStandupSortColumns    = []string{util.KeyID, util.KeySlug, util.KeyTitle, util.WithID(util.SvcSprint.Key), util.KeyOwner, util.KeyStatus, util.KeyCreated}
	allowedStorySortColumns      = []string{util.KeyID, util.WithID(util.SvcEstimate.Key), util.KeyIdx, util.WithID(util.KeyUser), util.KeyTitle, util.KeyStatus, "finalVote", util.KeyCreated}
	allowedTeamSortColumns       = []string{util.KeyID, util.KeySlug, util.KeyTitle, util.KeyOwner, util.KeyCreated}
	allowedUserSortColumns       = []string{util.KeyID, util.KeyName, util.KeyRole, util.KeyTheme, "navColor", "linkColor", "picture", "locale", util.KeyCreated}
	allowedVoteSortColumns       = []string{"storyID", util.WithID(util.KeyUser), util.KeyChoice, "updated", util.KeyCreated}
)

var allowedColumns = map[string][]string{
	util.KeyAction:       allowedActionSortColumns,
	util.KeyAdmin:        allowedAdminSortColumns,
	util.KeyAuth:         allowedAuthSortColumns,
	util.KeyComment:      allowedCommentSortColumns,
	util.SvcEstimate.Key: allowedEstimateSortColumns,
	util.KeyFeedback:     allowedFeedbackSortColumns,
	util.KeyHistory:      allowedHistorySortColumns,
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
