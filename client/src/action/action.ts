namespace action {
  export interface Action {
    readonly id: string;
    readonly svc: string;
    readonly modelID: string;
    readonly userID: string;
    readonly act: string;
    readonly content: any;
    readonly note: string;
    readonly created: string;
  }

  export function loadActions() {
    socket.send({svc: services.system.key, cmd: command.client.getActions, param: null});
  }

  export function viewActions(actions: action.Action[]) {
    dom.setContent("#action-list", renderActions(actions));
  }
}
