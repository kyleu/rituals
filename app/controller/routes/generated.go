package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cestimate"
	"github.com/kyleu/rituals/app/controller/cestimate/cstory"
	"github.com/kyleu/rituals/app/controller/cretro"
	"github.com/kyleu/rituals/app/controller/csprint"
	"github.com/kyleu/rituals/app/controller/cstandup"
	"github.com/kyleu/rituals/app/controller/cteam"
)

const routeNew, routeRandom, routeEdit, routeDelete = "/_new", "/_random", "/edit", "/delete"

func generatedRoutes(r *mux.Router) {
	generatedRoutesUser(r, "/admin/db/user")
	generatedRoutesTeamPermission(r, "/admin/db/team/permission")
	generatedRoutesTeamMember(r, "/admin/db/team/member")
	generatedRoutesTeamHistory(r, "/admin/db/team/history")
	generatedRoutesTeam(r, "/admin/db/team")
	generatedRoutesStandupPermission(r, "/admin/db/standup/permission")
	generatedRoutesStandupMember(r, "/admin/db/standup/member")
	generatedRoutesStandupHistory(r, "/admin/db/standup/history")
	generatedRoutesReport(r, "/admin/db/standup/report")
	generatedRoutesStandup(r, "/admin/db/standup")
	generatedRoutesSprintPermission(r, "/admin/db/sprint/permission")
	generatedRoutesSprintMember(r, "/admin/db/sprint/member")
	generatedRoutesSprintHistory(r, "/admin/db/sprint/history")
	generatedRoutesSprint(r, "/admin/db/sprint")
	generatedRoutesRetroPermission(r, "/admin/db/retro/permission")
	generatedRoutesRetroMember(r, "/admin/db/retro/member")
	generatedRoutesRetroHistory(r, "/admin/db/retro/history")
	generatedRoutesFeedback(r, "/admin/db/retro/feedback")
	generatedRoutesRetro(r, "/admin/db/retro")
	generatedRoutesVote(r, "/admin/db/estimate/story/vote")
	generatedRoutesStory(r, "/admin/db/estimate/story")
	generatedRoutesEstimatePermission(r, "/admin/db/estimate/permission")
	generatedRoutesEstimateMember(r, "/admin/db/estimate/member")
	generatedRoutesEstimateHistory(r, "/admin/db/estimate/history")
	generatedRoutesEstimate(r, "/admin/db/estimate")
	generatedRoutesEmail(r, "/admin/db/email")
	generatedRoutesComment(r, "/admin/db/comment")
	generatedRoutesAction(r, "/admin/db/action")
}

func generatedRoutesUser(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.UserList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.UserCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.UserCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.UserRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.UserDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.UserEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.UserEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.UserDelete)
}

func generatedRoutesTeamPermission(r *mux.Router, prefix string) {
	const pkn = "/{teamID}/{key}/{value}"
	makeRoute(r, http.MethodGet, prefix, cteam.TeamPermissionList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cteam.TeamPermissionCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cteam.TeamPermissionCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cteam.TeamPermissionRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cteam.TeamPermissionDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cteam.TeamPermissionEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cteam.TeamPermissionEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cteam.TeamPermissionDelete)
}

func generatedRoutesTeamMember(r *mux.Router, prefix string) {
	const pkn = "/{teamID}/{userID}"
	makeRoute(r, http.MethodGet, prefix, cteam.TeamMemberList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cteam.TeamMemberCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cteam.TeamMemberCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cteam.TeamMemberRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cteam.TeamMemberDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cteam.TeamMemberEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cteam.TeamMemberEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cteam.TeamMemberDelete)
}

func generatedRoutesTeamHistory(r *mux.Router, prefix string) {
	const pkn = "/{slug}"
	makeRoute(r, http.MethodGet, prefix, cteam.TeamHistoryList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cteam.TeamHistoryCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cteam.TeamHistoryCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cteam.TeamHistoryRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cteam.TeamHistoryDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cteam.TeamHistoryEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cteam.TeamHistoryEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cteam.TeamHistoryDelete)
}

func generatedRoutesTeam(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.TeamList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.TeamCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.TeamCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.TeamRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.TeamDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.TeamEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.TeamEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.TeamDelete)
}

func generatedRoutesStandupPermission(r *mux.Router, prefix string) {
	const pkn = "/{standupID}/{key}/{value}"
	makeRoute(r, http.MethodGet, prefix, cstandup.StandupPermissionList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cstandup.StandupPermissionCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cstandup.StandupPermissionCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cstandup.StandupPermissionRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cstandup.StandupPermissionDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cstandup.StandupPermissionEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cstandup.StandupPermissionEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cstandup.StandupPermissionDelete)
}

func generatedRoutesStandupMember(r *mux.Router, prefix string) {
	const pkn = "/{standupID}/{userID}"
	makeRoute(r, http.MethodGet, prefix, cstandup.StandupMemberList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cstandup.StandupMemberCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cstandup.StandupMemberCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cstandup.StandupMemberRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cstandup.StandupMemberDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cstandup.StandupMemberEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cstandup.StandupMemberEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cstandup.StandupMemberDelete)
}

func generatedRoutesStandupHistory(r *mux.Router, prefix string) {
	const pkn = "/{slug}"
	makeRoute(r, http.MethodGet, prefix, cstandup.StandupHistoryList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cstandup.StandupHistoryCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cstandup.StandupHistoryCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cstandup.StandupHistoryRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cstandup.StandupHistoryDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cstandup.StandupHistoryEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cstandup.StandupHistoryEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cstandup.StandupHistoryDelete)
}

func generatedRoutesReport(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, cstandup.ReportList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cstandup.ReportCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cstandup.ReportCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cstandup.ReportRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cstandup.ReportDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cstandup.ReportEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cstandup.ReportEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cstandup.ReportDelete)
}

func generatedRoutesStandup(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.StandupList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.StandupCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.StandupCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.StandupRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.StandupDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.StandupEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.StandupEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.StandupDelete)
}

func generatedRoutesSprintPermission(r *mux.Router, prefix string) {
	const pkn = "/{sprintID}/{key}/{value}"
	makeRoute(r, http.MethodGet, prefix, csprint.SprintPermissionList)
	makeRoute(r, http.MethodGet, prefix+routeNew, csprint.SprintPermissionCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, csprint.SprintPermissionCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, csprint.SprintPermissionRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, csprint.SprintPermissionDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, csprint.SprintPermissionEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, csprint.SprintPermissionEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, csprint.SprintPermissionDelete)
}

func generatedRoutesSprintMember(r *mux.Router, prefix string) {
	const pkn = "/{sprintID}/{userID}"
	makeRoute(r, http.MethodGet, prefix, csprint.SprintMemberList)
	makeRoute(r, http.MethodGet, prefix+routeNew, csprint.SprintMemberCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, csprint.SprintMemberCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, csprint.SprintMemberRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, csprint.SprintMemberDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, csprint.SprintMemberEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, csprint.SprintMemberEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, csprint.SprintMemberDelete)
}

func generatedRoutesSprintHistory(r *mux.Router, prefix string) {
	const pkn = "/{slug}"
	makeRoute(r, http.MethodGet, prefix, csprint.SprintHistoryList)
	makeRoute(r, http.MethodGet, prefix+routeNew, csprint.SprintHistoryCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, csprint.SprintHistoryCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, csprint.SprintHistoryRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, csprint.SprintHistoryDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, csprint.SprintHistoryEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, csprint.SprintHistoryEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, csprint.SprintHistoryDelete)
}

func generatedRoutesSprint(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.SprintList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.SprintCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.SprintCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.SprintRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.SprintDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.SprintEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.SprintEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.SprintDelete)
}

func generatedRoutesRetroPermission(r *mux.Router, prefix string) {
	const pkn = "/{retroID}/{key}/{value}"
	makeRoute(r, http.MethodGet, prefix, cretro.RetroPermissionList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cretro.RetroPermissionCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cretro.RetroPermissionCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cretro.RetroPermissionRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cretro.RetroPermissionDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cretro.RetroPermissionEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cretro.RetroPermissionEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cretro.RetroPermissionDelete)
}

func generatedRoutesRetroMember(r *mux.Router, prefix string) {
	const pkn = "/{retroID}/{userID}"
	makeRoute(r, http.MethodGet, prefix, cretro.RetroMemberList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cretro.RetroMemberCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cretro.RetroMemberCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cretro.RetroMemberRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cretro.RetroMemberDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cretro.RetroMemberEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cretro.RetroMemberEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cretro.RetroMemberDelete)
}

func generatedRoutesRetroHistory(r *mux.Router, prefix string) {
	const pkn = "/{slug}"
	makeRoute(r, http.MethodGet, prefix, cretro.RetroHistoryList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cretro.RetroHistoryCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cretro.RetroHistoryCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cretro.RetroHistoryRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cretro.RetroHistoryDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cretro.RetroHistoryEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cretro.RetroHistoryEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cretro.RetroHistoryDelete)
}

func generatedRoutesFeedback(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, cretro.FeedbackList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cretro.FeedbackCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cretro.FeedbackCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cretro.FeedbackRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cretro.FeedbackDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cretro.FeedbackEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cretro.FeedbackEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cretro.FeedbackDelete)
}

func generatedRoutesRetro(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.RetroList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.RetroCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.RetroCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.RetroRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.RetroDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.RetroEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.RetroEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.RetroDelete)
}

func generatedRoutesVote(r *mux.Router, prefix string) {
	const pkn = "/{storyID}/{userID}"
	makeRoute(r, http.MethodGet, prefix, cstory.VoteList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cstory.VoteCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cstory.VoteCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cstory.VoteRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cstory.VoteDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cstory.VoteEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cstory.VoteEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cstory.VoteDelete)
}

func generatedRoutesStory(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, cestimate.StoryList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cestimate.StoryCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cestimate.StoryCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cestimate.StoryRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cestimate.StoryDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cestimate.StoryEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cestimate.StoryEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cestimate.StoryDelete)
}

func generatedRoutesEstimatePermission(r *mux.Router, prefix string) {
	const pkn = "/{estimateID}/{key}/{value}"
	makeRoute(r, http.MethodGet, prefix, cestimate.EstimatePermissionList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cestimate.EstimatePermissionCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cestimate.EstimatePermissionCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cestimate.EstimatePermissionRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cestimate.EstimatePermissionDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cestimate.EstimatePermissionEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cestimate.EstimatePermissionEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cestimate.EstimatePermissionDelete)
}

func generatedRoutesEstimateMember(r *mux.Router, prefix string) {
	const pkn = "/{estimateID}/{userID}"
	makeRoute(r, http.MethodGet, prefix, cestimate.EstimateMemberList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cestimate.EstimateMemberCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cestimate.EstimateMemberCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cestimate.EstimateMemberRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cestimate.EstimateMemberDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cestimate.EstimateMemberEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cestimate.EstimateMemberEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cestimate.EstimateMemberDelete)
}

func generatedRoutesEstimateHistory(r *mux.Router, prefix string) {
	const pkn = "/{slug}"
	makeRoute(r, http.MethodGet, prefix, cestimate.EstimateHistoryList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cestimate.EstimateHistoryCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cestimate.EstimateHistoryCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cestimate.EstimateHistoryRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cestimate.EstimateHistoryDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cestimate.EstimateHistoryEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cestimate.EstimateHistoryEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cestimate.EstimateHistoryDelete)
}

func generatedRoutesEstimate(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.EstimateList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.EstimateCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.EstimateCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.EstimateRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.EstimateDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.EstimateEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.EstimateEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.EstimateDelete)
}

func generatedRoutesEmail(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.EmailList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.EmailCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.EmailCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.EmailRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.EmailDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.EmailEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.EmailEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.EmailDelete)
}

func generatedRoutesComment(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.CommentList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.CommentCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.CommentCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.CommentRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.CommentDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.CommentEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.CommentEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.CommentDelete)
}

func generatedRoutesAction(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.ActionList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.ActionCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.ActionCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.ActionRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.ActionDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.ActionEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.ActionEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.ActionDelete)
}
