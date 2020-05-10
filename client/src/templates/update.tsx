namespace update {
  function renderUpdate(model: Update): JSX.Element {
    let profile = system.cache.profile;
    if (profile === undefined) {
      return <div class="uk-margin-bottom">error</div>;
    } else {
      return <div class="section" onclick={"events.openModal('update', '" + model + "');"}>
        <div class={profile.linkColor + "-fg section-link"}>{model.id}</div>
      </div>;
    }
  }

  export function renderUpdates(updates: Update[]): JSX.Element {
    if (updates.length === 0) {
      return <div>
        <button class="uk-button uk-button-default" onclick="events.openModal('add-stuff');" type="button">Add Update</button>
      </div>;
    } else {
      return <ul class="uk-list uk-list-divider">
        {updates.map(update => <li id={"update-" + update.id}>
          {renderUpdate(update)}
        </li>)}
      </ul>;
    }
  }
}
