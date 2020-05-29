namespace feedback {
  export interface Feedback {
    readonly id: string;
    readonly idx: number;
    readonly authorID: string;
    readonly category: string;
    readonly content: string;
    readonly html: string;
    readonly created: string;
  }

  export interface FeedbackCategory {
    readonly category: string;
    readonly feedback: Feedback[];
  }

  export function setFeedback(feedback: feedback.Feedback[]) {
    retro.cache.feedback = feedback;
    dom.setContent("#feedback-detail", renderFeedbackArray(feedback));
    UIkit.modal("#modal-add-feedback").hide();
  }

  export function onSubmitFeedback() {
    const category = dom.req<HTMLInputElement>("#feedback-category").value;
    const content = dom.req<HTMLInputElement>("#feedback-content").value;
    const msg = {svc: services.retro.key, cmd: command.client.addFeedback, param: {category, content}};
    socket.send(msg);
    return false;
  }

  export function onEditFeedback() {
    const id = retro.cache.activeFeedback;
    const category = dom.req<HTMLInputElement>("#feedback-edit-category").value;
    const content = dom.req<HTMLInputElement>("#feedback-edit-content").value;
    const msg = {svc: services.retro.key, cmd: command.client.updateFeedback, param: {id, category, content}};
    socket.send(msg);
    return false;
  }

  export function onRemoveFeedback() {
    const id = retro.cache.activeFeedback;
    if (id) {
      UIkit.modal.confirm("Delete this feedback?").then(function () {
        const msg = {svc: services.retro.key, cmd: command.client.removeFeedback, param: id};
        socket.send(msg);
        UIkit.modal("#modal-feedback").hide();
      });
    }
    return false;
  }

  export function getActiveFeedback() {
    if (!retro.cache.activeFeedback) {
      return undefined;
    }
    const curr = retro.cache.feedback.filter(x => x.id === retro.cache.activeFeedback).shift();
    if (!curr) {
      console.warn(`cannot load active Feedback [${retro.cache.activeFeedback}]`);
    }
    return curr;
  }

  export function viewActiveFeedback() {
    const profile = system.cache.getProfile();

    const fb = getActiveFeedback();
    if (!fb) {
      console.warn("no active feedback");
      return;
    }

    const same = fb.authorID === profile.userID;

    dom.setText("#feedback-title", `${fb.category} / ${system.getMemberName(fb.authorID)}`);
    dom.setSelectOption("#feedback-edit-category", same ? fb.category : undefined);
    contents.onContentDisplay("feedback", same, fb.content, fb.html);
  }

  export function onFeedbackUpdate(r: feedback.Feedback) {
    const x = preUpdate(r.id);
    x.push(r);
    postUpdate(x, r.id);
  }

  export function onFeedbackRemoved(id: string) {
    const x = preUpdate(id);
    postUpdate(x, id);
    UIkit.notification("feedback has been deleted", {status: "success", pos: "top-right"});
  }

  function preUpdate(id: string) {
    return retro.cache.feedback.filter((p) => p.id !== id);
  }

  function postUpdate(x: feedback.Feedback[], id: string) {
    feedback.setFeedback(x);
    if (id === retro.cache.activeFeedback) {
      UIkit.modal("#modal-feedback").hide();
    }
  }

  export function getFeedbackCategories(feedback: feedback.Feedback[], categories: string[]) {
    function toCollection(c: string): FeedbackCategory {
      const reports = feedback.filter(r => r.category === c).sort((l, r) => (l.created > r.created ? -1 : 1));
      return {category: c, feedback: reports};
    }

    const ret = categories.map(toCollection);
    const extras = feedback.filter(r => !categories.find(x => x === r.category));
    if (extras.length > 0) {
      ret.push({category: "unknown", feedback: extras});
    }
    return ret;
  }
}
