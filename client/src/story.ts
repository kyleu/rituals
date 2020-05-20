namespace story {
  export interface Story {
    id: string;
    idx: number;
    authorID: string;
    title: string;
    status: string;
    finalVote: string;
    created: string;
  }

  export function setStories(stories: Story[]) {
    estimate.cache.stories = stories;
    dom.setContent("#story-detail", renderStories(stories));
    stories.forEach(s => setStoryStatus(s.id, s.status, s, false));
    showTotalIfNeeded();
    UIkit.modal("#modal-add-story").hide();
  }

  export function onSubmitStory() {
    const title = dom.req<HTMLInputElement>("#story-title-input").value;
    const msg = {svc: services.estimate.key, cmd: command.client.addStory, param: {title: title}};
    socket.send(msg);
    return false;
  }

  export function beginEditStory() {
    const s = getActiveStory()!;
    const newTitle = prompt("Edit your story", s.title)
    if(newTitle !== null && newTitle !== s.title) {
      const msg = {svc: services.estimate.key, cmd: command.client.updateStory, param: { id: s.id, title: newTitle }};
      socket.send(msg);
    }
    return false;
  }

  export function onRemoveStory() {
    const id = estimate.cache.activeStory;
    if(id && confirm("Delete this story?")) {
      const msg = {svc: services.estimate.key, cmd: command.client.removeStory, param: id};
      socket.send(msg);
      UIkit.modal("#modal-story").hide();
    }
    return false;
  }

  export function getActiveStory() {
    if (estimate.cache.activeStory === undefined) {
      return undefined;
    }
    const curr = estimate.cache.stories.filter(x => x.id === estimate.cache.activeStory);
    if (curr.length !== 1) {
      console.warn(`cannot load active story [${estimate.cache.activeStory}]`);
      return undefined;
    }
    return curr[0];
  }

  export function viewActiveStory() {
    const s = getActiveStory();
    if (s === undefined) {
      return;
    }
    dom.setText("#story-title", s.title);
    viewStoryStatus(s.status);
  }

  export function showTotalIfNeeded() {
    let stories = estimate.cache.stories;
    let strings = stories.filter(s => s.status === "complete").map(s => s.finalVote).filter(c => c.length > 0);
    let floats = strings.map(c => parseFloat(c)).filter(f => !isNaN(f));
    let sum = 0;
    floats.forEach(f => sum += f);
    let curr = dom.opt("#story-total");
    let panel = dom.req("#story-list");
    if (curr !== null) {
      panel.removeChild(curr);
    }
    if(sum > 0) {
      panel.appendChild(renderTotal(sum));
    }
  }
}
