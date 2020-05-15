namespace action {
  export interface Action {
    id: string;
    svc: string;
    modelID: string;
    authorID: string;
    act: string;
    content: any;
    note: string;
    occurred: string;
  }

  export function loadActions() {
    const msg = {
      svc: services.system,
      cmd: command.client.getActions,
      param: null
    };
    socket.send(msg);
  }

  export function viewActions(actions: action.Action[]) {
    const el = util.req("#action-list");
    el.innerHTML = "";
    el.appendChild(renderActions(actions))
  }
}
