namespace feedback {
  export interface Feedback {
    readonly id: string;
    readonly idx: number;
    readonly userID: string;
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
    comment.setCounts();
    modal.hide("add-feedback");
  }

  export function onSubmitFeedback() {
    const category = dom.req<HTMLInputElement>("#feedback-category").value;
    const content = dom.req<HTMLInputElement>("#feedback-content").value;
    const msg = { svc: services.retro.key, cmd: command.client.addFeedback, param: { category, content } };
    socket.send(msg);
    return false;
  }

  export function onEditFeedback() {
    const id = retro.cache.activeFeedback;
    const category = dom.req<HTMLInputElement>("#feedback-edit-category").value;
    const content = dom.req<HTMLInputElement>("#feedback-edit-content").value;
    const msg = { svc: services.retro.key, cmd: command.client.updateFeedback, param: { id, category, content } };
    socket.send(msg);
    return false;
  }

  export function onRemoveFeedback() {
    const id = retro.cache.activeFeedback;
    if (id) {
      notify.confirm("Delete this feedback?", function () {
        socket.send({ svc: services.retro.key, cmd: command.client.removeFeedback, param: id });
        modal.hide("feedback");
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

  export function viewActiveFeedback(id?: string) {
    const profile = system.cache.getProfile();

    if (id) {
      retro.cache.activeFeedback = id;
    }
    const fb = getActiveFeedback();
    if (!fb) {
      console.warn("no active feedback");
      return;
    }

    const same = fb.userID === profile.userID;

    dom.setText("#feedback-title", `${fb.category} / ${member.getMember(fb.userID)?.name}`);
    dom.setSelectOption("#feedback-edit-category", same ? fb.category : undefined);
    contents.onContentDisplay("feedback", same, fb.content, fb.html);
    comment.setActive("feedback", fb.id);
    comment.setCounts();
  }

  export function onFeedbackUpdate(r: feedback.Feedback) {
    const x = preUpdate(r.id);
    x.push(r);
    postUpdate(x, r.id);
  }

  export function onFeedbackRemoved(id: string) {
    const x = preUpdate(id);
    postUpdate(x, id);
    notify.notify("feedback has been deleted", true);
  }

  function preUpdate(id: string) {
    return retro.cache.feedback.filter(p => p.id !== id);
  }

  function postUpdate(x: feedback.Feedback[], id: string) {
    feedback.setFeedback(x);
    if (id === retro.cache.activeFeedback) {
      modal.hide("feedback");
    }
  }

  export function getFeedbackCategories(feedback: readonly feedback.Feedback[], categories: readonly string[]) {
    function toCollection(c: string): FeedbackCategory {
      const reports = feedback.filter(r => r.category === c).sort((l, r) => (l.created > r.created ? -1 : 1));
      return { category: c, feedback: reports };
    }

    const ret = categories.map(toCollection);
    const extras = feedback.filter(r => !categories.find(x => x === r.category));
    if (extras.length > 0) {
      ret.push({ category: "unknown", feedback: extras });
    }
    return ret;
  }

  export function viewAddFeedback(p: string | undefined) {
    dom.setSelectOption("#feedback-category", p);
    const feedbackContent = dom.setValue("#feedback-content", "");
    dom.wireTextarea(feedbackContent as HTMLTextAreaElement);
    setTimeout(() => feedbackContent.focus(), 250);
  }

  export function viewFeedback(p: string | undefined) {
    feedback.viewActiveFeedback(p);
    const feedbackEditContent = dom.req("#feedback-edit-content");
    setTimeout(() => {
      dom.wireTextarea(feedbackEditContent as HTMLTextAreaElement);
      feedbackEditContent.focus();
    }, 250);
  }
}
