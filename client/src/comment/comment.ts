namespace comment {
  export interface Comment {
    readonly id: string;
    readonly targetType: string;
    readonly targetID: string;
    readonly userID: string;
    readonly content: string;
    readonly html: string;
    readonly created: string;
  }

  let activeComments: comment.Comment[] = [];
  let activeType: string | undefined
  let activeID: string | undefined

  export function applyComments(comments: comment.Comment[]) {
    activeComments = comments;
  }

  export function setActive(t: string, id: string | undefined) {
    activeType = t;
    activeID = id;
  }

  export function show(t: string) {
    modal.open("comment", t);
  }

  export function add(t: string) {
    const textarea = dom.req<HTMLTextAreaElement>(`#comment-add-content-${t}`);
    const v = textarea.value;
    textarea.value = "";
    const param = {targetType: activeType, targetID: activeID, content: v};
    socket.send({svc: services.system.key, cmd: command.client.addComment, param: param});
  }

  export function onCommentUpdate(u: Comment) {
    activeComments.push(u);
    setCounts();
    load();
  }

  function find(t: string | undefined, id: string | undefined) {
    if ((!t) || t === "modal") {
      t = activeType;
      if(!id) {
        id = activeID;
      }
    }
    if (t === "root") {
      t = "";
    }
    if(id) {
      return activeComments.filter(x => x.targetType === t && x.targetID == id);
    }
    return activeComments.filter(x => x.targetType === t);
  }

  export function load(t?: string, id?: string) {
    if ((!t) || t === "modal") {
      t = activeType;
      if(!id) {
        id = activeID;
      }
    }
    if (!t) {
      console.warn(`invalid comment type [${t}]`);
      return;
    }

    activeType = t;
    activeID = id;
    const comments = find(t, id);
    if (t !== "root") {
      t = "modal";
    }
    const el = dom.req(`#drop-comment-${t} .uk-comment-list`);
    dom.setContent(el, renderComments(comments, system.cache.getProfile()));
    el.scrollTop = el.scrollHeight;
  }

  export function setCounts() {
    const containers = dom.els(`.comment-count-container`);
    let matchedModal = false
    const modalCount = dom.opt(`#comment-count-modal`);
    containers.forEach(cc => {
      const t = cc.dataset["commentType"];
      const id = cc.dataset["commentId"];
      if(!t) {
        throw `invalid comment type [${t}] with id [${id}]`;
      }
      let comments = find(t, id);
      setCount(t, comments, cc);
      if(activeType === t) {
        if (modalCount) {
          setCount(t, comments, modalCount);
          matchedModal = true;
        }
      }
    });
    if (!matchedModal && modalCount) {
      setCount("modal", [], modalCount, true);
    }
  }

  export function closeModal() {
    if (activeType === "story") {
      modal.openSoon("story");
    } else if (activeType === "report") {
      modal.openSoon("report");
    } else if (activeType === "feedback") {
      modal.openSoon("feedback");
    }
    activeType = undefined;
  }

  export function remove(id: string) {
    if (confirm("remove this comment?")) {
      socket.send({svc: services.system.key, cmd: command.client.removeComment, param: id});
    }
    return false;
  }

  function setCount(t: string, comments: ReadonlyArray<Comment>, cc: HTMLElement, force?: boolean) {
    dom.req(".text", cc).innerText = comments.length.toString();
    if (t !== "root" && t !== "modal" && t !== "") {
      dom.setDisplay(cc, (comments.length !== 0) || force === true);
    }
  }

  export function onCommentRemoved(id: string) {
    activeComments = activeComments.filter((c) => c.id !== id);
    setCounts();
    load();
  }
}
