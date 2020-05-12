namespace story {
  function renderStory(story: story.Story): JSX.Element {
    const profile = system.cache.getProfile();
    return <li id={"story-" + story.id} class="section" onclick={"events.openModal('story', '" + story.id + "');"}>
      <div class="right uk-article-meta story-status">{story.status}</div>
      <div class={profile.linkColor + "-fg section-link"}>{story.title}</div>
    </li>;
  }

  export function renderStories(stories: story.Story[]): JSX.Element {
    if (stories.length === 0) {
      return <div>
        <button class="uk-button uk-button-default" onclick="events.openModal('add-story');" type="button">Add Story</button>
      </div>;
    } else {
      return <ul id="story-list" class="uk-list uk-list-divider">
        {stories.map(s => renderStory(s))}
      </ul>;
    }
  }

  export function renderStatus(status: string): JSX.Element {
    switch (status) {
      case "pending":
        return <span>{status}</span>;
      case "active":
        return <span>{status}</span>;
      default:
        return <span class="vote-badge uk-border-rounded">{status}</span>;
    }
  }

  export function renderTotal(sum: number) {
    return <li id="story-total"><div class="right uk-article-meta"><span class="vote-badge uk-border-rounded">{sum}</span></div> Total</li>;
  }
}
