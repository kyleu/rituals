namespace story {
  export interface Story {
    readonly id: string;
    readonly idx: number;
    readonly userID: string;
    readonly title: string;
    status: string;
    finalVote: string;
    readonly created: string;
  }

  export function setStories(stories: Story[]) {
    estimate.cache.stories = stories;
    dom.setContent("#story-detail", renderStories(stories));
    stories.forEach(s => setStoryStatus(s.id, s.status, s, false));
    showTotalIfNeeded();
    comment.setCounts();
    modal.hide("add-story");
  }

  export function onSubmitStory() {
    const title = dom.req<HTMLInputElement>("#story-title-input").value;
    const msg = { svc: services.estimate.key, cmd: command.client.addStory, param: { title } };
    socket.send(msg);
    return false;
  }

  export function beginEditStory() {
    const s = getActiveStory()!;
    const title = prompt("Edit your story", s.title);
    if (title === null) {
      return false;
    }
    if (title && title !== s.title) {
      const msg = { svc: services.estimate.key, cmd: command.client.updateStory, param: { storyID: s.id, title } };
      socket.send(msg);
    }
    return false;
  }

  export function onRemoveStory() {
    const id = estimate.cache.activeStory;
    if (id) {
      notify.confirm("Delete this story?", function () {
        const msg = { svc: services.estimate.key, cmd: command.client.removeStory, param: id };
        socket.send(msg);
        modal.hide("story");
      });
    }
    return false;
  }

  export function getActiveStory() {
    if (!estimate.cache.activeStory) {
      return undefined;
    }
    const curr = estimate.cache.stories.filter(x => x.id === estimate.cache.activeStory).shift();
    if (!curr) {
      console.warn(`cannot load active story [${estimate.cache.activeStory}]`);
    }
    return curr;
  }

  export function viewActiveStory(id?: string) {
    if (id) {
      estimate.cache.activeStory = id;
    }

    const s = getActiveStory();
    if (!s) {
      return;
    }
    dom.setText("#story-title", s.title);
    viewStoryStatus(s.status);
    comment.setActive("story", s.id);
    comment.setCounts();
  }

  export function showTotalIfNeeded() {
    const stories = estimate.cache.stories;
    const strings = stories.filter(s => s.status === "complete").map(s => s.finalVote).filter(c => c.length > 0);
    const floats = strings.map(c => parseFloat(c)).filter(f => !isNaN(f));
    let sum = 0;
    floats.forEach(f => (sum += f));
    const curr = dom.opt("#story-total");
    const panel = dom.req("#story-list");
    if (curr !== undefined) {
      panel.removeChild(curr);
    }
    if (sum > 0) {
      panel.appendChild(renderTotal(sum));
    }
  }
}
