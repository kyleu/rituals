namespace member {
  let me: member.Member | undefined;

  export function isSelf(x: Member) {
    return x.userID === system.cache.getProfile().userID;
  }

  export function selfCanEdit() {
    return me !== undefined && canEdit(me);
  }

  export function updateSelf(self: member.Member | undefined) {
    if (self) {
      me = self
      dom.setContent("#self-picture", setPicture(self.picture));
      dom.setText("#member-self .member-name", self.name);
      dom.setValue("#self-name-input", self.name);
      dom.setText("#member-self .member-role", self.role);
      const e = canEdit(self);
      dom.setDisplay("#history-container", e);
      dom.setDisplay("#session-edit-section", e);
      dom.setDisplay("#session-view-section", !e);
    }
  }

  export function onSubmitSelf() {
    const name = dom.req<HTMLInputElement>("#self-name-input").value;
    const choice = dom.req<HTMLInputElement>("#self-name-choice-global").checked ? "global" : "local";
    const picture = dom.req<HTMLInputElement>("#self-picture-input").value;
    const msg = {svc: services.system.key, cmd: command.client.updateProfile, param: {name, choice, picture}};
    socket.send(msg);
  }

  function canEdit(m: member.Member) {
    return m.role == "owner";
  }
}
