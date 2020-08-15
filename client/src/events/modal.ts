namespace modal {
  let activeParam: string | undefined;

  export function wire() {
    dom.els(".modal").forEach(el => {
      el.addEventListener("show", onModalOpen);
      el.addEventListener("hide", onModalHide);
    });

    events.register("welcome");

    // session
    events.register("session", session.onModalOpen);
    events.register("action", action.loadActions);

    // member
    events.register("self", member.viewSelf);
    events.register("invitation", member.loadQR);
    events.register("member", member.viewActiveMember);

    // estimate
    events.register("add-story", story.viewAddStory);
    events.register("story", story.viewActiveStory);

    // standup
    events.register("add-report", report.viewAddReport);
    events.register("report", report.viewReport);

    // retro
    events.register("add-feedback", feedback.viewAddFeedback);
    events.register("feedback", feedback.viewFeedback);
  }

  export function open(key: string, param?: string) {
    activeParam = param;
    const m = notify.modal(`#modal-${key}`);
    m.show();
    return false;
  }

  export function openSoon(key: string) {
    setTimeout(() => open(key), 0);
  }

  export function hide(key: string) {
    const m = notify.modal(`#modal-${key}`);
    const el: HTMLElement = m.$el;
    if (el.classList.contains("uk-open")) {
      m.hide();
    }
  }

  function onModalOpen(e: Event) {
    if (!e.target) {
      return;
    }
    const el = e.target as HTMLElement;
    if (el.id.indexOf("modal") !== 0) {
      return;
    }
    const key = el.id.substr("modal-".length);
    const f = events.getOpenEvent(key);
    if (f) {
      f(activeParam);
    } else {
      console.warn(`no modal open handler registered for [${key}]`);
    }
    activeParam = undefined;
  }

  function onModalHide(e: Event) {
    if (!e.target) {
      return;
    }
    const el = e.target as HTMLElement;
    if (el.classList.contains("uk-open")) {
      const key = el.id.substr("modal-".length);
      const f = events.getCloseEvent(key);
      if (f) {
        f(activeParam);
      }
      activeParam = undefined;
    }
  }
}
