namespace story {
  export interface Story {
    id: string;
    idx: number;
    author: string;
    title: string;
    status: {
      key: string;
    };
    finalVote: string;
    created: string;
  }

  export function setStories(stories: Story[]) {
    estimate.cache.stories = stories;
    util.setContent("#story-detail", renderStories(stories));
    stories.forEach(s => setStoryStatus(s.id, s.status.key, s, false));
    showTotalIfNeeded();
    UIkit.modal("#modal-add-story").hide();
  }

  export function onSubmitStory() {
    const title = util.req<HTMLInputElement>("#story-title-input").value;
    const msg = {
      svc: services.estimate,
      cmd: command.client.addStory,
      param: {title: title},
    };
    socket.send(msg);
    return false;
  }

  export function getActiveStory() {
    if (estimate.cache.activeStory === undefined) {
      return undefined;
    }
    const curr = estimate.cache.stories.filter(x => x.id === estimate.cache.activeStory);
    if (curr.length !== 1) {
      console.warn("cannot load active story [" + estimate.cache.activeStory + "]");
      return undefined;
    }
    return curr[0];
  }

  export function viewActiveStory() {
    const s = getActiveStory();
    if (s === undefined) {
      console.log("no active story");
      return;
    }
    util.setText("#story-title", s.title);
    viewStoryStatus(s.status.key);
  }

  function viewStoryStatus(status: string) {
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
      console.log("no active story");
      return;
    }
    const msg = {
      svc: services.estimate,
      cmd: command.client.setStoryStatus,
      param: {storyID: story.id, status: s},
    };
    socket.send(msg);
  }

  function setStoryStatus(storyID: string, status: string, currStory: story.Story | null, calcTotal: boolean) {
    if (currStory !== null && currStory!.status.key == "complete") {
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

    setStoryStatus(u.storyID, u.status.key, currStory, true);
    if(u.storyID === estimate.cache.activeStory) {
      viewStoryStatus(u.status.key);
    }
  }

  function showTotalIfNeeded() {
    let stories = estimate.cache.stories;
    let strings = stories.filter(s => s.status.key === "complete").map(s => s.finalVote).filter(c => c.length > 0);
    let floats = strings.map(c => parseFloat(c)).filter(f => !isNaN(f));
    let sum = 0;
    floats.forEach(f => sum += f);
    let curr = util.opt("#story-total");
    let panel = util.req("#story-list");
    if (curr !== null) {
      panel.removeChild(curr);
    }
    if(sum > 0) {
      panel.appendChild(renderTotal(sum));
    }
  }
}
