import {els, req} from "./dom";
import {svgRef} from "./util";
import {memberList} from "./member";
import {snippetVote} from "./votes";
import {send} from "./app";

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

function wireFinalVote(f: HTMLFormElement) {
  f.onsubmit = () => {
    const storyID = req<HTMLInputElement>("input[name=\"storyID\"]", f).value;
    const typ = req<HTMLInputElement>("input[name=\"typ\"]", f).value;
    const value = req<HTMLInputElement>("input[name=\"value\"]", f).value;

    const storyEl = req("#modal-story-" + storyID);
    req(".final-vote .value", storyEl).innerText = value;

    send("vote", {storyID, typ, value});
    return false;
  };
}

export function applyCalcs(storyID: string, votes: Vote[], results: VoteResults, finalVote: string) {
  const items = memberList().map((member) => {
    const v = votes.find((vote) => {
      return vote.userID === member.id;
    });
    return snippetVote(storyID, member, v);
  });

  const storyEl = req("#modal-story-" + storyID);

  req(" .vote-results", storyEl).replaceChildren(...items);
  const fv = req(".final-vote", storyEl);
  req(".value", fv).innerText = finalVote;
  req(".value", fv).style.display = "inline-block";
  req(".message", fv).style.display = "none";
  req(".description", fv).style.display = "block";

  const calcEl = req(".vote-calculations", storyEl);
  req(".calc-counted .value", calcEl).innerText = results.count + "/" + results.floats.length;
  req(".calc-range .value", calcEl).innerText = results.min + "-" + results.max;
  req(".calc-mean .value", calcEl).innerText = results.mean.toString(10);
  req<HTMLInputElement>(".calc-mean input[name=\"value\"]", calcEl).value = results.mean.toString(10);
  req(".calc-median .value", calcEl).innerText = results.median.toString(10);
  req<HTMLInputElement>(".calc-median input[name=\"value\"]", calcEl).value = results.median.toString(10);
  req(".calc-mode .value", calcEl).innerText = results.modeString === "" ? "0" : results.modeString;
  req<HTMLInputElement>(".calc-mode input[name=\"value\"]", calcEl).value = results.modeString === "" ? "0" : results.modeString;

  els<HTMLFormElement>(".final-vote-form", storyEl).forEach((frm) => {
    wireFinalVote(frm);
  });
}
