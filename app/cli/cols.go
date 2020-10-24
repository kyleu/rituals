package cli

import (
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	allowedActionSortColumns     = []string{npncore.KeyID, npncore.KeySvc, npncore.WithID(npncore.KeyModel), npncore.WithID(npncore.KeyUser), "act", npncore.KeyContent, npncore.KeyNote, npncore.KeyCreated}
	allowedAdminSortColumns      = []string{npncore.KeyName, "recent", "count"}
	allowedAuthSortColumns       = []string{npncore.KeyID, npncore.WithID(npncore.KeyUser), npncore.KeyProvider, npncore.WithID(npncore.KeyProvider), "expires", npncore.KeyName, npncore.KeyEmail, "picture", npncore.KeyCreated}
	allowedCommentSortColumns    = []string{npncore.KeyID, npncore.KeySvc, npncore.WithID(npncore.KeyModel), npncore.WithID(npncore.KeyUser), npncore.KeyContent, npncore.KeyHTML, npncore.KeyCreated}
	allowedEmailSortColumns      = []string{npncore.KeyID, "recipients", "subject", "data", "plain", "html", npncore.WithID(npncore.KeyUser), npncore.KeyStatus, npncore.KeyCreated}
	allowedEstimateSortColumns   = []string{npncore.KeyID, npncore.KeySlug, npncore.KeyTitle, npncore.WithID(util.SvcSprint.Key), npncore.KeyOwner, npncore.KeyStatus, npncore.Plural(npncore.KeyChoice), "options", npncore.KeyCreated}
	allowedFeedbackSortColumns   = []string{npncore.KeyID, npncore.WithID(util.SvcRetro.Key), npncore.KeyIdx, npncore.WithID(npncore.KeyUser), npncore.KeyCategory, npncore.KeyContent, npncore.KeyHTML, npncore.KeyCreated}
	allowedHistorySortColumns    = []string{npncore.KeySlug, npncore.WithID(npncore.KeyModel), "modelName", npncore.KeyCreated}
	allowedMemberSortColumns     = []string{npncore.WithID(npncore.KeyUser), npncore.KeyName, npncore.KeyRole, npncore.KeyCreated}
	allowedMigrationSortColumns  = []string{npncore.KeyIdx, npncore.KeyTitle, "src", npncore.KeyCreated}
	allowedPermissionSortColumns = []string{"k", "v", "access", npncore.KeyCreated}
	allowedReportSortColumns     = []string{npncore.KeyID, npncore.WithID(util.SvcStandup.Key), "d", npncore.WithID(npncore.KeyUser), npncore.KeyContent, npncore.KeyHTML, npncore.KeyCreated}
	allowedRetroSortColumns      = []string{npncore.KeyID, npncore.KeySlug, npncore.KeyTitle, npncore.WithID(util.SvcSprint.Key), npncore.KeyOwner, npncore.KeyStatus, npncore.Plural(npncore.KeyCategory), npncore.KeyCreated}
	allowedSocketSortColumns     = []string{npncore.KeySvc, "cmd", "param"}
	allowedSprintSortColumns     = []string{npncore.KeyID, npncore.KeySlug, npncore.KeyTitle, npncore.KeyOwner, "startDate", "endDate", npncore.KeyCreated}
	allowedStandupSortColumns    = []string{npncore.KeyID, npncore.KeySlug, npncore.KeyTitle, npncore.WithID(util.SvcSprint.Key), npncore.KeyOwner, npncore.KeyStatus, npncore.KeyCreated}
	allowedStorySortColumns      = []string{npncore.KeyID, npncore.WithID(util.SvcEstimate.Key), npncore.KeyIdx, npncore.WithID(npncore.KeyUser), npncore.KeyTitle, npncore.KeyStatus, "finalVote", npncore.KeyCreated}
	allowedTeamSortColumns       = []string{npncore.KeyID, npncore.KeySlug, npncore.KeyTitle, npncore.KeyOwner, npncore.KeyCreated}
	allowedUserSortColumns       = []string{npncore.KeyID, npncore.KeyName, npncore.KeyRole, npncore.KeyTheme, "navColor", "linkColor", "picture", "locale", npncore.KeyCreated}
	allowedVoteSortColumns       = []string{"storyID", npncore.WithID(npncore.KeyUser), npncore.KeyChoice, "updated", npncore.KeyCreated}
)

func InitCols() {
	// TODO
	AllowedColumns := make(map[string][]string)

	AllowedColumns[npncore.KeyAction] = allowedActionSortColumns
	AllowedColumns[npncore.KeyAdmin] = allowedAdminSortColumns
	AllowedColumns[npncore.KeyAuth] = allowedAuthSortColumns
	AllowedColumns[npncore.KeyComment] = allowedCommentSortColumns
	AllowedColumns[npncore.KeyEmail] = allowedEmailSortColumns
	AllowedColumns[util.SvcEstimate.Key] = allowedEstimateSortColumns
	AllowedColumns[util.KeyFeedback] = allowedFeedbackSortColumns
	AllowedColumns[npncore.KeyHistory] = allowedHistorySortColumns
	AllowedColumns[npncore.KeyMember] = allowedMemberSortColumns
	AllowedColumns[npncore.KeyMigration] = allowedMigrationSortColumns
	AllowedColumns[npncore.KeyPermission] = allowedPermissionSortColumns
	AllowedColumns[npncore.KeyReport] = allowedReportSortColumns
	AllowedColumns[util.SvcRetro.Key] = allowedRetroSortColumns
	AllowedColumns[npncore.KeySocket] = allowedSocketSortColumns
	AllowedColumns[util.SvcSprint.Key] = allowedSprintSortColumns
	AllowedColumns[util.SvcStandup.Key] = allowedStandupSortColumns
	AllowedColumns[util.KeyStory] = allowedStorySortColumns
	AllowedColumns[npncore.KeyUser] = allowedUserSortColumns
	AllowedColumns[util.SvcTeam.Key] = allowedTeamSortColumns
	AllowedColumns[util.KeyVote] = allowedVoteSortColumns
}
