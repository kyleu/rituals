namespace feedback {
  export interface Feedback {
    id: string;
    idx: number;
    authorID: string;
    category: string;
    content: string;
    html: string;
    created: string;
  }

  export interface FeedbackCategory {
    category: string;
    feedback: Feedback[];
  }

  export function setFeedback(feedback: feedback.Feedback[]) {
    retro.cache.feedback = feedback;
    util.setContent("#feedback-detail", renderFeedbackArray(feedback));
    UIkit.modal("#modal-add-feedback").hide();
  }

  export function onSubmitFeedback() {
    const category = util.req<HTMLInputElement>("#retro-feedback-category").value;
    const content = util.req<HTMLInputElement>("#retro-feedback-content").value;
    const msg = {
      svc: services.retro,
      cmd: command.client.addFeedback,
      param: {category: category, content: content},
    };
    socket.send(msg);
    return false;
  }

  export function onEditFeedback() {
    const id = retro.cache.activeFeedback;
    const category = util.req<HTMLInputElement>("#retro-feedback-edit-category").value;
    const content = util.req<HTMLInputElement>("#retro-feedback-edit-content").value;
    const msg = {
      svc: services.retro,
      cmd: command.client.editFeedback,
      param: {id: id, category: category, content: content},
    };
    socket.send(msg);
    return false;
  }

  export function getActiveFeedback() {
    if (retro.cache.activeFeedback === undefined) {
      return undefined;
    }
    const curr = retro.cache.feedback.filter(x => x.id === retro.cache.activeFeedback);
    if (curr.length !== 1) {
      console.warn("cannot load active Feedback [" + retro.cache.activeFeedback + "]");
      return undefined;
    }
    return curr[0];
  }

  export function viewActiveFeedback() {
    const profile = system.cache.getProfile();
    const fb = getActiveFeedback();
    if (fb === undefined) {
      console.warn("no active feedback");
      return;
    }

    util.setText("#feedback-title", fb.category + " / " + member.getMemberName(fb.authorID));
    const contentEdit = util.req("#modal-feedback .content-edit");
    const contentEditCategory = util.req<HTMLSelectElement>("#retro-feedback-edit-category", contentEdit);
    const contentEditTextarea = util.req<HTMLTextAreaElement>("#retro-feedback-edit-content", contentEdit);
    const contentView = util.req("#modal-feedback .content-view");
    const buttonsEdit = util.req("#modal-feedback .buttons-edit");
    const buttonsView = util.req("#modal-feedback .buttons-view");

    if(fb.authorID === profile.userID) {
      contentEdit.style.display = "block";
      util.setSelectOption(contentEditCategory, fb.category);
      util.setValue(contentEditTextarea, fb.content);
      util.wireTextarea(contentEditTextarea);
      contentView.style.display = "none";
      util.setHTML(contentView, "");
      buttonsEdit.style.display = "block";
      buttonsView.style.display = "none";
    } else {
      contentEdit.style.display = "none";
      util.setSelectOption(contentEditCategory, undefined);
      util.setValue(contentEditTextarea, "");
      contentView.style.display = "block";
      util.setHTML(contentView, fb.html);
      buttonsEdit.style.display = "none";
      buttonsView.style.display = "block";
    }
  }

  export function onFeedbackUpdate(r: feedback.Feedback) {
    let x = retro.cache.feedback;

    x = x.filter((p) => p.id !== r.id);
    x.push(r);

    feedback.setFeedback(x);

    if(r.id === retro.cache.activeFeedback) {
      UIkit.modal("#modal-feedback").hide();
    }
  }

  export function getFeedbackCategories(feedback: feedback.Feedback[], categories: string[]) {
    function toCollection(c: string): FeedbackCategory {
      const reports = feedback.filter(r => r.category === c).sort((l, r) => (l.created > r.created ? -1 : 1));
      return {category: c, feedback: reports};
    }

    let ret = categories.map(toCollection);
    const extras = feedback.filter(r => categories.indexOf(r.category) == -1);
    if (extras.length > 0) {
      ret.push({category: "unknown", feedback: extras})
    }
    return ret;
  }
}
