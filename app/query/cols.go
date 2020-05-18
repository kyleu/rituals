package query

var allowedActionSortColumns = []string{"id", "svc", "model_id", "author_id", "act", "content", "note", "occurred"}
var allowedAuthSortColumns = []string{"id", "user_id", "provider", "provider_id", "expires", "name", "email", "picture", "created"}
var allowedEstimateSortColumns = []string{"id", "slug", "title", "sprint_id", "owner", "status", "choices", "options", "created"}
var allowedInvitationSortColumns = []string{"key", "k", "v", "src", "tgt", "note", "status", "redeemed", "created"}
var allowedMemberSortColumns = []string{"user_id", "name", "role", "created"}
var allowedReportSortColumns = []string{"id", "standup_id", "d", "author_id", "content", "html", "created"}
var allowedRetroSortColumns = []string{"id", "slug", "title", "sprint_id", "owner", "status", "categories", "created"}
var allowedSocketSortColumns = []string{"svc", "cmd", "param"}
var allowedSprintSortColumns = []string{"id", "slug", "title", "owner", "end_date", "created"}
var allowedStandupSortColumns = []string{"id", "slug", "title", "sprint_id", "owner", "status", "created"}
var allowedStorySortColumns = []string{"id", "estimate_id", "idx", "author_id", "title", "status", "final_vote", "created"}
var allowedUserSortColumns = []string{"id", "name", "role", "theme", "nav_color", "link_color", "picture", "locale", "created"}
var allowedVoteSortColumns = []string{"story_id", "user_id", "choice", "updated", "created"}

var allowedColumns = map[string][]string{
	"action":   allowedActionSortColumns,
	"auth":     allowedAuthSortColumns,
	"estimate": allowedEstimateSortColumns,
	"invite":   allowedInvitationSortColumns,
	"member":   allowedMemberSortColumns,
	"report":   allowedReportSortColumns,
	"retro":    allowedRetroSortColumns,
	"socket":   allowedSocketSortColumns,
	"sprint":   allowedSprintSortColumns,
	"standup":  allowedStandupSortColumns,
	"story":    allowedStorySortColumns,
	"user":     allowedUserSortColumns,
	"vote":     allowedVoteSortColumns,
}
