namespace update {
  export interface Update {
    id: string;
    d: string;
    author: string;
    content: string;
    created: string;
  }

  export function onSubmitUpdate() {
    const d = util.req<HTMLInputElement>("#standup-update-date").value;
    const content = util.req<HTMLInputElement>("#standup-update-input").value;
    const msg = {
      svc: services.standup,
      cmd: command.client.addUpdate,
      param: {d: d, title: content},
    };
    socket.send(msg);
    return false;
  }

  export function viewActiveUpdate() {
    console.log("viewActiveUpdate");
  }

  export function setUpdates(updates: update.Update[]) {
    standup.cache.updates = updates;
    util.setContent("#update-detail", renderUpdates(updates));
    UIkit.modal("#modal-add-update").hide();
  }
}
