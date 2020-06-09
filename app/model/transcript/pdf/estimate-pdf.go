package pdf

import (
	"emperror.dev/errors"
	"fmt"
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/util"
	"strings"
)

func renderEstimate(rsp transcript.EstimateResponse, m pdfgen.Maroto) (string, error) {
	hr(m)
	caption(rsp.Session.Title, m)
	detailRow(util.Title(util.KeyOwner), rsp.Members.GetName(rsp.Session.Owner), m)
	detailRow(util.PluralTitle(util.KeyChoice), strings.Join(rsp.Session.Choices, ", "), m)
	if rsp.Team != nil {
		detailRow(util.SvcTeam.Title, rsp.Team.Title, m)
	}
	if rsp.Sprint != nil {
		detailRow(util.SvcSprint.Title, rsp.Sprint.Title, m)
	}
	detailRow(util.Title(util.KeyCreated), util.ToDateString(&rsp.Session.Created), m)

	var err error
	_, err = renderPermissionList(rsp.Permissions, m)
	if err != nil {
		return "", err
	}
	_, err = renderMemberList(rsp.Members, m)
	if err != nil {
		return "", err
	}
	err = renderStoryList(rsp.Stories, rsp.Votes, rsp.Members, rsp.Comments.ForType("story"), m)
	if err != nil {
		return "", err
	}
	_, err = renderCommentList(rsp.Comments.ForType(""), rsp.Members, m, true)
	if err != nil {
		return "", err
	}

	return rsp.Session.Slug, nil
}

func renderEstimateList(sessions estimate.Sessions, members member.Entries, m pdfgen.Maroto) (string, error) {
	if len(sessions) > 0 {
		hr(m)
		caption(util.SvcEstimate.PluralTitle, m)
		cols := []string{util.Title(util.KeyOwner), util.Title(util.KeyTitle), util.Title(util.KeyCreated)}
		var data [][]string
		for _, s := range sessions {
			data = append(data, []string{members.GetName(s.Owner), s.Title, util.ToDateString(&s.Created)})
		}
		table(cols, data, []uint{3, 6, 3}, m)
	}
	return "", nil
}

func renderStoryList(stories estimate.Stories, votes estimate.Votes, members member.Entries, commments comment.Comments, m pdfgen.Maroto) error {
	if len(stories) > 0 {
		for _, story := range stories {
			err := renderStory(story, votes, members, commments, m)
			if err != nil {
				return errors.Wrap(err, "")
			}
		}
	}
	return nil
}

func renderStory(story *estimate.Story, votes estimate.Votes, members member.Entries, commments comment.Comments, m pdfgen.Maroto) error {
	hr(m)
	tr(func() {
		th(story.ID.String(), 11, m)
		td(story.FinalVote, 1, m)
	}, 12, m)
	tr(func() {
		td(members.GetName(story.UserID), 6, m)
		td(story.Status.Key, 3, m)
		td(util.ToDateString(&story.Created), 3, m)
	}, 8, m)
	storyVotes := estimate.VotesForStory(votes, story.ID)
	if len(storyVotes) > 0 {
		tr(func() {
			th(util.PluralTitle(util.KeyVote), 12, m)
		}, 8, m)
		var msg []string
		for _, v := range storyVotes {
			msg = append(msg, members.GetName(story.UserID)+": "+v.Choice)
		}
		tr(func() {
			td(strings.Join(msg, ", "), 12, m)
		}, 8, m)
		tr(func() {
			td("Count", 2, m)
			td("Min", 2, m)
			td("Max", 2, m)
			td("Mean", 2, m)
			td("Mode", 2, m)
			td("Median", 2, m)
		}, 8, m)
		tr(func() {
			res := estimate.CalculateVoteResult(storyVotes)
			td(fmt.Sprint(res.Count), 2, m)
			td(res.Min, 2, m)
			td(res.Max, 2, m)
			td(res.Mean, 2, m)
			td(res.Mode, 2, m)
			td(res.Median, 2, m)
		}, 8, m)
	}
	storyComments := commments.ForID(story.ID)
	if len(storyComments) > 0 {
		_, err := renderCommentList(storyComments, members, m, false)
		if err != nil {
			return err
		}
	}
	return nil
}
