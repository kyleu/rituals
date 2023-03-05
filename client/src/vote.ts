import {els, req} from "./dom";
import {svgRef} from "./util";

export type Vote = {
  storyID: string;
  userID: string;
  choice: string;
}

export function onVote(v: Vote) {
  els("#modal-story-" + v.storyID + " .story-members .member").forEach((m) => {
    if (m.dataset.member == v.userID) {
      req(".choice", m).innerHTML = svgRef("check", 18, "");
    }
  });
  console.log("VOTE", v);
}

