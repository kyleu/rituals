import {els, req} from "./dom";
import {svgRef} from "./util";
import {memberList} from "./member";
import {snippetVote} from "./votes";

export type Vote = {
  storyID: string;
  userID: string;
  choice: string;
}

export type VoteResults = {
  floats: number[];
  count: number;
  min: number;
  max: number;
  range: number;
  sum: number;
  mean: number;
  median: number;
  mode: number[];
  modeString: string;
}

export function onVote(v: Vote) {
  els("#modal-story-" + v.storyID + " .story-members .member").forEach((m) => {
    if (m.dataset.member === v.userID) {
      req(".choice", m).innerHTML = svgRef("check", 18, "");
    }
  });
}

export function applyCalcs(storyID: string, votes: Vote[], results: VoteResults) {
  const items = memberList().map((member) => {
    const v = votes.find((vote) => {
      return vote.userID === member.id;
    });
    return snippetVote(member, v);
  });

  const storyEl = req("#modal-story-" + storyID);
  req(" .vote-results", storyEl).replaceChildren(...items);

  const calcEl = req(".vote-calculations", storyEl);
  req(".calc-counted .value", calcEl).innerText = results.count + "/" + results.floats.length;
  req(".calc-range .value", calcEl).innerText = results.min + "-" + results.max;
  req(".calc-mean .value", calcEl).innerText = results.mean.toString(10);
  req(".calc-median .value", calcEl).innerText = results.median.toString(10);
  req(".calc-mode .value", calcEl).innerText = results.modeString;
}

