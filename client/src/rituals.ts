interface Message {
  svc: string;
  cmd: string;
  param: any;
}

interface Profile {
  userID: string;
  name: string;
  role: string;
  theme: string;
  navColor: string;
  linkColor: string;
  locale: string;
}

interface Session {
  id: string;
  slug: string;
  title: string;
  owner: string;
  status: { key: string };
  created: string;
}

interface SessionJoined {
  profile: Profile;
  session: Session;
  members: Member[];
  online: string[];
}

function onSocketMessage(msg: Message) {
  console.log("message received");
  console.log(msg);
  switch (msg.svc) {
    case services.system:
      onSystemMessage(msg.cmd, msg.param);
      break;
    case services.estimate:
      onEstimateMessage(msg.cmd, msg.param);
      break;
    default:
      console.warn("unhandled message for service [" + msg.svc + "]");
  }
}

function setDetail(session: Session) {
  systemCache.session = session;
  $id("model-title").innerText = session.title;
  $id<HTMLInputElement>("model-title-input").value = session.title;
  let items = $("#navbar .uk-navbar-item");
  if (items.length > 0) {
    items[items.length - 1].innerText = session.title;
  }

  UIkit.modal("#modal-session").hide();
}

function onError(err: string) {
  console.warn(err);
  const idx = err.lastIndexOf(":");
  if (idx > -1) {
    err = err.substr(idx + 1);
  }
  UIkit.notification(err, { status: "danger", pos: "top-right" });
}

function onSystemMessage(cmd: string, param: any) {
  switch(cmd) {
  case serverCmd.error:
    onError("server error: " + param);
    break;
  case serverCmd.memberUpdate:
    onMemberUpdate(param as Member);
    break;
  case serverCmd.onlineUpdate:
    onOnlineUpdate(param as OnlineUpdate);
    break;
  default:
    console.warn("unhandled system message for command [" + cmd + "]");
  }
}
function onSessionJoin(param: SessionJoined) {
  console.log("joined");

  systemCache.session = param.session;
  systemCache.profile = param.profile;

  systemCache.members = param.members;
  systemCache.online = param.online;

  setMembers();
}
