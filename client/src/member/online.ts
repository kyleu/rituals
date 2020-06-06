namespace member {
  export interface OnlineUpdate {
    readonly userID: string;
    readonly connected: boolean;
  }

  let online: string[] = [];

  export function onOnlineUpdate(update: OnlineUpdate) {
    if (update.connected) {
      if (!online.find(x => x === update.userID)) {
        online.push(update.userID);
      }
    } else {
      online = online.filter(x => x !== update.userID);
    }
    renderOnline();
  }

  export function applyOnline(o: string[]) {
    online = o
  }

  export function renderOnline() {
    for (const member of getMembers()) {
      const el = dom.opt(`#member-${member.userID} .online-indicator`);
      if (el) {
        if (!online.find(x => x === member.userID)) {
          el.classList.add("offline");
        } else {
          el.classList.remove("offline");
        }
      }
    }
  }

  function canEdit(m: member.Member) {
    return m.role == "owner";
  }
}
