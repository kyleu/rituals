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

    for (const el of dom.els(".story-status-body")) {
      setActive(el, status);
    }
    for (const el of dom.els(".story-status-actions")) {
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
    dom.setText("#story-status", txt);

    vote.viewVotes();
  }

  export function requestStoryStatus(status: string) {
    const story = getActiveStory();
    if (!story) {
      return;
    }
    const param = {storyID: story.id, status};
    socket.send({svc: services.estimate.key, cmd: command.client.setStoryStatus, param: param});
  }

  export function setStoryStatus(storyID: string, status: string, currStory: story.Story | undefined, calcTotal: boolean) {
    if (currStory && currStory!.status === "complete") {
      if (currStory!.finalVote.length > 0) {
        status = currStory!.finalVote;
      }
    }
    dom.setContent(`#story-${storyID} .story-status`, renderStatus(status));
    if (calcTotal) {
      showTotalIfNeeded();
    }
  }

  export function onStoryStatusChange(u: estimate.StoryStatusChange) {
    let currStory: Story | undefined = undefined;
    estimate.cache.stories.forEach(s => {
      if (s.id === u.storyID) {
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
