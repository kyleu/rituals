namespace modal {
  let openEvents: Map<string, Function> = new Map<string, Function>();
  let closeEvents: Map<string, Function> = new Map<string, Function>();

  let activeParam: string | undefined;

  export function register(key: string, o?: (param?: string) => void, c?: (param?: string) => void) {
    if (!o) {
      o = () => {}
    }
    openEvents.set(key, o);
    if (c) {
      closeEvents.set(key, c);
    }
  }

  export function wire() {
    UIkit.util.on(".modal", "show", onModalOpen);
    UIkit.util.on(".modal", "hide", onModalHide);

    register("welcome");

    // session
    register("session", session.onModalOpen);
    register("action", action.loadActions);
    register("comment", comment.load, comment.closeModal);

    // member
    register("self", member.viewSelf);
    register("invitation");
    register("member", member.viewActiveMember);

    // estimate
    register("add-story", story.viewAddStory);
    register("story", story.viewActiveStory);

    // standup
    register("add-report", report.viewAddReport);
    register("report", report.viewReport);

    // retro
    register("add-feedback", feedback.viewAddFeedback);
    register("feedback", feedback.viewFeedback);
  }

  export function open(key: string, param?: string) {
    activeParam = param;
    const m = UIkit.modal(`#modal-${key}`);
    if (!m) {
      console.warn(`no modal available with key [${key}]`);
    }
    m.show();
    return false;
  }

  export function openSoon(key: string) {
    setTimeout(() => open(key), 0);
  }

  export function hide(key: string) {
    const m = UIkit.modal(`#modal-${key}`);
    const el: HTMLElement = m.$el;
    if(el.classList.contains("uk-open")) {
      m.hide();
    }
  }

  function onModalOpen(e: Event) {
    if(!e.target) {
      return;
    }
    const el = e.target as HTMLElement;
    const key = el.id.substr("modal-".length);
    const f = openEvents.get(key);
    if (f) {
      f(activeParam);
    } else {
      console.warn(`no modal open handler registered for [${key}]`);
    }
    activeParam = undefined;
  }

  function onModalHide(e: Event) {
    if(!e.target) {
      return;
    }
    const el = e.target as HTMLElement;
    if(el.classList.contains("uk-open")) {
      const key = el.id.substr("modal-".length);
      const f = closeEvents.get(key);
      if (f) {
        f(activeParam);
      }
      activeParam = undefined;
    }
  }
}
