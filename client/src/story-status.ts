namespace story {
  export function viewStoryStatus(status: string) {
    function setActive(el: HTMLElement, status: string) {
      const s = el.id.substr(el.id.lastIndexOf("-") + 1);
      if (s === status) {
        el.classList.add("active");
      } else {
        el.classList.remove("active");
      }
    }

    for (let el of util.els(".story-status-body")) {
      setActive(el, status);
    }
    for (let el of util.els(".story-status-actions")) {
      setActive(el, status);
    }

    let txt = "";
    switch(status) {
      case "pending":
        txt = "Story";
        break;
      case "active":
        txt = "Voting";
        break;
      case "complete":
        txt = "Results";
        break;
    }
    util.setText("#story-status", txt);

    vote.viewVotes();
  }

  export function requestStoryStatus(s: string) {
    const story = getActiveStory();
    if (story === undefined) {
      console.warn("no active story");
      return;
    }
    const msg = {
      svc: services.estimate,
      cmd: command.client.setStoryStatus,
      param: {storyID: story.id, status: s},
    };
    socket.send(msg);
  }

  export function setStoryStatus(storyID: string, status: string, currStory: story.Story | null, calcTotal: boolean) {
    if (currStory !== null && currStory!.status == "complete") {
      if (currStory!.finalVote.length > 0) {
        status = currStory!.finalVote;
      }
    }
    util.setContent("#story-" + storyID + " .story-status", renderStatus(status));
    if (calcTotal) {
      showTotalIfNeeded();
    }
  }

  export function onStoryStatusChange(u: estimate.StoryStatusChange) {
    let currStory: Story | null = null;
    estimate.cache.stories.forEach(s => {
      if (s.id == u.storyID) {
        currStory = s;
        s.finalVote = u.finalVote;
        s.status = u.status;
      }
    });

    setStoryStatus(u.storyID, u.status, currStory, true);
    if(u.storyID === estimate.cache.activeStory) {
      viewStoryStatus(u.status);
    }
  }
}
