namespace story {
  function renderStory(story: story.Story) {
    const profile = system.cache.getProfile();
    return (
      <li id={`story-${story.id}`} class="section" onclick={`modal.open('story', '${story.id}');`}>
        <div class="right uk-article-meta story-status">{story.status}</div>
        <div class={`${profile.linkColor}-fg section-link left`}>{story.title}</div>
        <div class="left" style="margin-top: 4px;">
          {comment.renderCount("story", story.id)}
        </div>
      </li>
    );
  }

  export function renderStories(stories: ReadonlyArray<story.Story>) {
    if (stories.length === 0) {
      return (
        <div id="story-list">
          <button class="uk-button uk-button-default" onclick="modal.open('add-story');" type="button">
            Add Story
          </button>
        </div>
      );
    } else {
      return (
        <table class="uk-table uk-table-divider uk-table-small">
          <ul id="story-list" class="uk-list uk-list-divider">
            {stories.map(s => renderStory(s))}
          </ul>
        </table>
      );
    }
  }

  export function renderStatus(status: string) {
    switch (status) {
      case "pending":
        return <span>{status}</span>;
      case "active":
        return <span>{status}</span>;
      default:
        return <span class="vote-badge">{status}</span>;
    }
  }

  export function renderTotal(sum: number) {
    return (
      <li id="story-total">
        <div class="right uk-article-meta">
          <span class="vote-badge">{sum}</span>
        </div>{" "}
        Total
      </li>
    );
  }

  export function viewAddStory() {
    const storyInput = dom.setValue("#story-title-input", "");
    setTimeout(() => storyInput.focus(), 250);
  }
}
