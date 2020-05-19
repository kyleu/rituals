package query

import "github.com/kyleu/rituals.dev/app/util"

var allowedActionSortColumns = []string{"id", "svc", "model_id", "author_id", "act", "content", "note", "occurred"}
var allowedAuthSortColumns = []string{"id", "user_id", "provider", "provider_id", "expires", "name", "email", "picture", "created"}
var allowedEstimateSortColumns = []string{"id", "slug", "title", "sprint_id", "owner", "status", "choices", "options", "created"}
var allowedInvitationSortColumns = []string{"key", "k", "v", "src", "tgt", "note", "status", "redeemed", "created"}
var allowedMemberSortColumns = []string{"user_id", "name", "role", "created"}
var allowedReportSortColumns = []string{"id", "standup_id", "d", "author_id", "content", "html", "created"}
var allowedRetroSortColumns = []string{"id", "slug", "title", "sprint_id", "owner", "status", "categories", "created"}
var allowedSocketSortColumns = []string{"svc", "cmd", "param"}
var allowedSprintSortColumns = []string{"id", "slug", "title", "owner", "start_date", "end_date", "created"}
var allowedStandupSortColumns = []string{"id", "slug", "title", "sprint_id", "owner", "status", "created"}
var allowedStorySortColumns = []string{"id", "estimate_id", "idx", "author_id", "title", "status", "final_vote", "created"}
var allowedTeamSortColumns = []string{"id", "slug", "title", "owner", "created"}
var allowedUserSortColumns = []string{"id", "name", "role", "theme", "nav_color", "link_color", "picture", "locale", "created"}
var allowedVoteSortColumns = []string{"story_id", "user_id", "choice", "updated", "created"}

var allowedColumns = map[string][]string{
	util.KeyAction:       allowedActionSortColumns,
	util.KeyAuth:         allowedAuthSortColumns,
	util.SvcEstimate.Key: allowedEstimateSortColumns,
	util.KeyInvitation:   allowedInvitationSortColumns,
	util.KeyMember:       allowedMemberSortColumns,
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
