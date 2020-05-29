namespace action {
  export interface Action {
    readonly id: string;
    readonly svc: string;
    readonly modelID: string;
    readonly authorID: string;
    readonly act: string;
    readonly content: any;
    readonly note: string;
    readonly created: string;
  }

  export function loadActions() {
    const msg = {svc: services.system.key, cmd: command.client.getActions, param: null};
    socket.send(msg);
  }

  export function viewActions(actions: action.Action[]) {
    dom.setContent("#action-list", renderActions(actions));
  }
}
