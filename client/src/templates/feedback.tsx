namespace feedback {
  function renderFeedback(model: Feedback): JSX.Element {
    const profile = system.cache.getProfile();
    const ret = <div id={"feedback-" + model.id} class="feedback-detail uk-border-rounded section" onclick={"events.openModal('feedback', '" + model.id + "');"}>
      <a class={profile.linkColor + "-fg section-link"}>{system.getMemberName(model.authorID)}</a>
      <div class="feedback-content">loading...</div>
    </div>;

    if(model.html.length > 0) {
      util.setHTML(util.req(".feedback-content", ret), model.html).style.display = "block";
    }

    return ret;
  }

  export function renderFeedbackArray(f: feedback.Feedback[]): JSX.Element {
    if (f.length === 0) {
      return <div>
        <button class="uk-button uk-button-default" onclick="events.openModal('add-feedback');" type="button">Add Feedback</button>
      </div>;
    } else {
      const cats = getFeedbackCategories(f, retro.cache.detail?.categories || []);
      const profile = system.cache.getProfile();
      return <div class="uk-grid-small uk-grid-match uk-child-width-expand@m uk-grid-divider" uk-grid="">
        {cats.map(cat => <div class="feedback-list uk-transition-toggle">
          <div class="feedback-category-header">
            <span class="right">
              <a class={profile.linkColor + "-fg uk-icon-button uk-transition-fade"} data-uk-icon="plus" onclick={"events.openModal('add-feedback', '" + cat.category + "');"} title="edit session"></a>
            </span>

            <span class="feedback-category-title" onclick={"events.openModal('add-feedback', '" + cat.category + "');"}>{cat.category}</span>

          </div>
          <div>
            {cat.feedback.map(fb => <div>{renderFeedback(fb)}</div>)}
          </div>
        </div>)}
      </div>;
    }
  }
}
