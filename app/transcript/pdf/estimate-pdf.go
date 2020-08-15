package pdf

import (
	"fmt"
	npnpdf "github.com/kyleu/npn/npnexport/pdf"
	"strings"

	"github.com/kyleu/npn/npncore"

	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderEstimate(rsp transcript.EstimateResponse, m pdfgen.Maroto) string {
	npnpdf.HR(m)
	npnpdf.Caption(rsp.Session.Title, m)
	npnpdf.DetailRow(npncore.Title(npncore.KeyOwner), rsp.Members.GetName(rsp.Session.Owner), m)
	npnpdf.DetailRow(npncore.PluralTitle(npncore.KeyChoice), strings.Join(rsp.Session.Choices, ", "), m)
	if rsp.Team != nil {
		npnpdf.DetailRow(util.SvcTeam.Title, rsp.Team.Title, m)
	}
	if rsp.Sprint != nil {
		npnpdf.DetailRow(util.SvcSprint.Title, rsp.Sprint.Title, m)
	}
	npnpdf.DetailRow(npncore.Title(npncore.KeyCreated), npncore.ToDateString(&rsp.Session.Created), m)

	renderPermissionList(rsp.Permissions, m)
	renderMemberList(rsp.Members, m)
	renderStoryList(rsp.Stories, rsp.Votes, rsp.Members, rsp.Comments.ForType("story"), m)
	renderCommentList(rsp.Comments.ForType(""), rsp.Members, m, true)

	return rsp.Session.Slug
}

func renderEstimateList(sessions estimate.Sessions, members member.Entries, m pdfgen.Maroto) {
	if len(sessions) > 0 {
		npnpdf.HR(m)
		npnpdf.Caption(util.SvcEstimate.PluralTitle, m)
		cols := []string{npncore.Title(npncore.KeyOwner), npncore.Title(npncore.KeyTitle), npncore.Title(npncore.KeyCreated)}
		var data [][]string
		for _, s := range sessions {
			data = append(data, []string{members.GetName(s.Owner), s.Title, npncore.ToDateString(&s.Created)})
		}
		npnpdf.Table(cols, data, []uint{3, 6, 3}, m)
	}
}

func renderStoryList(stories estimate.Stories, votes estimate.Votes, members member.Entries, comments comment.Comments, m pdfgen.Maroto) {
	if len(stories) > 0 {
		for _, story := range stories {
			renderStory(story, votes, members, comments, m)
		}
	}
}

func renderStory(story *estimate.Story, votes estimate.Votes, members member.Entries, comments comment.Comments, m pdfgen.Maroto) {
	npnpdf.HR(m)
	npnpdf.TR(func() {
		npnpdf.TH(story.ID.String(), 11, m)
		npnpdf.TD(story.FinalVote, 1, m)
	}, 12, m)
	npnpdf.TR(func() {
		npnpdf.TD(members.GetName(story.UserID), 6, m)
		npnpdf.TD(story.Status.Key, 3, m)
		npnpdf.TD(npncore.ToDateString(&story.Created), 3, m)
	}, 8, m)
	storyVotes := estimate.VotesForStory(votes, story.ID)
	if len(storyVotes) > 0 {
		npnpdf.TR(func() {
			npnpdf.TH(npncore.PluralTitle(util.KeyVote), 12, m)
		}, 8, m)
		var msg []string
		for _, v := range storyVotes {
			msg = append(msg, members.GetName(story.UserID)+": "+v.Choice)
		}
		npnpdf.TR(func() {
			npnpdf.TD(strings.Join(msg, ", "), 12, m)
		}, 8, m)
		npnpdf.TR(func() {
			npnpdf.TD("Count", 2, m)
			npnpdf.TD("Min", 2, m)
			npnpdf.TD("Max", 2, m)
			npnpdf.TD("Mean", 2, m)
			npnpdf.TD("Mode", 2, m)
			npnpdf.TD("Median", 2, m)
		}, 8, m)
		npnpdf.TR(func() {
			res := estimate.CalculateVoteResult(storyVotes)
			npnpdf.TD(fmt.Sprint(res.Count), 2, m)
			npnpdf.TD(res.Min, 2, m)
			npnpdf.TD(res.Max, 2, m)
			npnpdf.TD(res.Mean, 2, m)
			npnpdf.TD(res.Mode, 2, m)
			npnpdf.TD(res.Median, 2, m)
		}, 8, m)
	}
	storyComments := comments.ForID(story.ID)
	if len(storyComments) > 0 {
		renderCommentList(storyComments, members, m, false)
	}
}
