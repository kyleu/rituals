"use strict";
var EstimateCache = /** @class */ (function () {
    function EstimateCache() {
        this.polls = [];
        this.votes = [];
    }
    return EstimateCache;
}());
var estimateCache = new EstimateCache();
function onEstimateMessage(cmd, param) {
    switch (cmd) {
        case serverCmd.sessionJoined:
            onSessionJoin(param);
            setEstimateDetail(param.session);
            setPolls(param.polls);
            setVotes(param.votes);
            break;
        case serverCmd.sessionUpdate:
            setEstimateDetail(param.session);
            break;
        case serverCmd.pollUpdate:
            onPollUpdate(param);
            break;
        default:
            console.warn("unhandled command [" + cmd + "] for estimate");
    }
}
function setEstimateDetail(detail) {
    estimateCache.detail = detail;
    $id("model-choices-input").value = detail.choices.join(", ");
    setDetail(detail);
}
function onSubmitEstimateSession() {
    var title = $req("#model-title-input").value;
    var choices = $req("#model-choices-input").value;
    var msg = {
        svc: services.estimate,
        cmd: clientCmd.updateSession,
        param: {
            title: title,
            choices: choices
        }
    };
    send(msg);
}
function onPollUpdate(poll) {
    var x = estimateCache.polls;
    x = x.filter(function (p) { return p.id != poll.id; });
    x.push(poll);
    x = x.sort(function (l, r) { return (l.idx > r.idx) ? 1 : -1; });
    setPolls(x);
}
function delay(f) {
    setTimeout(f, 250);
}
function modalShow(key) {
    switch (key) {
        case "self":
            var selfInput_1 = $req("#self-name-input");
            selfInput_1.value = $req("#member-self .member-name").innerText;
            delay(function () { return selfInput_1.focus(); });
            break;
        case "session":
            var sessionInput_1 = $req("#model-title-input");
            sessionInput_1.value = $req("#model-title").innerText;
            delay(function () { return sessionInput_1.focus(); });
            break;
        case "invite":
            break;
        case "member":
            viewActiveMember();
            break;
        case "add-poll":
            var pollInput_1 = $req("#poll-title-input");
            pollInput_1.value = "";
            delay(function () { return pollInput_1.focus(); });
            break;
        case "poll":
            viewActivePoll();
            break;
        default:
            console.debug("unhandled modal [" + key + "]");
    }
}
var debug = true;
var appUnloading = false;
var SystemCache = /** @class */ (function () {
    function SystemCache() {
        this.profile = null;
        this.currentService = "";
        this.currentID = "";
        this.connectTime = 0;
        this.members = [];
        this.online = [];
    }
    return SystemCache;
}());
var systemCache = new SystemCache();
var services = {
    system: "system",
    estimate: "estimate",
    standup: "standup",
    retro: "retro"
};
var clientCmd = {
    error: "error",
    ping: "ping",
    connect: "connect",
    updateProfile: "update-profile",
    updateSession: "update-session",
    addPoll: "add-poll",
    updatePoll: "update-poll",
    setPollStatus: "set-poll-status",
    submitVote: "submit-vote"
};
var serverCmd = {
    error: "error",
    pong: "pong",
    sessionJoined: "session-joined",
    sessionUpdate: "session-update",
    memberUpdate: "member-update",
    onlineUpdate: "online-update",
    pollUpdate: "poll-update",
    voteUpdate: "vote-update"
};
function JSX(tag, attrs, children) {
    var e = document.createElement(tag);
    for (var name_1 in attrs) {
        if (name_1 && attrs.hasOwnProperty(name_1)) {
            var v = attrs[name_1];
            if (v === true) {
                e.setAttribute(name_1, name_1);
            }
            else if (v !== false && v != null) {
                e.setAttribute(name_1, v.toString());
            }
        }
    }
    for (var i = 2; i < arguments.length; i++) {
        var child = arguments[i];
        if (Array.isArray(child)) {
            child.forEach(function (c) {
                e.appendChild(c);
            });
        }
        else {
            if (child.nodeType === null || child.nodeType === undefined) {
                child = document.createTextNode(child.toString());
            }
            e.appendChild(child);
        }
    }
    return e;
}
function isSelf(x) {
    if (systemCache.profile === null) {
        return false;
    }
    return x.userID === systemCache.profile.userID;
}
function setMembers() {
    var self = systemCache.members.filter(isSelf);
    if (self.length === 1) {
        $req("#member-self .member-name").innerText = self[0].name;
        $req("#self-name-input").value = self[0].name;
        $req("#member-self .member-role").innerText = self[0].role.key;
    }
    else if (self.length === 0) {
        console.warn("self not found among members");
    }
    else {
        console.warn("multiple self entries found among members");
    }
    var others = systemCache.members.filter(function (x) { return !isSelf(x); });
    var detail = $id("member-detail");
    detail.innerHTML = "";
    detail.appendChild(renderMembers(others));
    renderOnline();
}
function onMemberUpdate(member) {
    if (isSelf(member)) {
        UIkit.modal("#modal-self").hide();
    }
    var x = systemCache.members;
    x = x.filter(function (m) { return m.userID != member.userID; });
    x.push(member);
    x = x.sort(function (l, r) { return (l.name > r.name) ? 1 : -1; });
    systemCache.members = x;
    setMembers();
}
function onOnlineUpdate(update) {
    if (update.connected) {
        if (systemCache.online.indexOf(update.userID) == -1) {
            systemCache.online.push(update.userID);
        }
    }
    else {
        systemCache.online = systemCache.online.filter(function (x) { return x != update.userID; });
    }
    renderOnline();
}
function renderOnline() {
    for (var _i = 0, _a = systemCache.members; _i < _a.length; _i++) {
        var member = _a[_i];
        var els = $("#member-" + member.userID + " .online-indicator");
        if (els.length === 1) {
            if (systemCache.online.indexOf(member.userID) === -1) {
                els[0].classList.add("offline");
            }
            else {
                els[0].classList.remove("offline");
            }
        }
    }
}
function onSubmitSelf() {
    var name = $req("#self-name-input").value;
    var choice = $req("#self-name-choice-global").checked ? "global" : "local";
    var msg = {
        svc: services.system,
        cmd: clientCmd.updateProfile,
        param: {
            name: name,
            choice: choice
        }
    };
    send(msg);
}
function getActiveMember() {
    if (systemCache.activeMember === null) {
        console.warn("no active member");
        return null;
    }
    var curr = systemCache.members.filter(function (x) { return x.userID === systemCache.activeMember; });
    if (curr.length !== 1) {
        console.log("cannot load active member [" + systemCache.activeMember + "]");
        return null;
    }
    return curr[0];
}
function viewActiveMember() {
    var member = getActiveMember();
    if (member === null) {
        return;
    }
    $req("#member-modal-name").innerText = member.name;
    $req("#member-modal-role").innerText = member.role.key;
}
function setPolls(polls) {
    estimateCache.polls = polls;
    var detail = $id("poll-detail");
    detail.innerHTML = "";
    detail.appendChild(renderPolls(polls));
    UIkit.modal("#modal-add-poll").hide();
}
function setVotes(votes) {
    console.log("todo: votes");
}
function onSubmitPoll() {
    var title = $req("#poll-title-input").value;
    var msg = {
        svc: services.estimate,
        cmd: clientCmd.addPoll,
        param: {
            title: title
        }
    };
    send(msg);
}
function getActivePoll() {
    if (estimateCache.activePoll === null) {
        console.warn("no active poll");
        return null;
    }
    var curr = estimateCache.polls.filter(function (x) { return x.id === estimateCache.activePoll; });
    if (curr.length !== 1) {
        console.log("cannot load active poll [" + estimateCache.activePoll + "]");
        return null;
    }
    return curr[0];
}
function viewActivePoll() {
    var poll = getActivePoll();
    if (poll === null) {
        console.log("no active poll");
        return;
    }
    $req("#poll-title").innerText = poll.title;
}
function onSocketMessage(msg) {
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
function setDetail(session) {
    systemCache.session = session;
    $id("model-title").innerText = session.title;
    $id("model-title-input").value = session.title;
    var items = $("#navbar .uk-navbar-item");
    if (items.length > 0) {
        items[items.length - 1].innerText = session.title;
    }
    UIkit.modal("#modal-session").hide();
}
function onError(err) {
    console.warn(err);
    var idx = err.lastIndexOf(":");
    if (idx > -1) {
        err = err.substr(idx + 1);
    }
    UIkit.notification(err, { status: 'danger', pos: 'top-right' });
}
function onSystemMessage(cmd, param) {
    switch (cmd) {
        case serverCmd.error:
            onError("server error: " + param);
            break;
        case serverCmd.memberUpdate:
            onMemberUpdate(param);
            break;
        case serverCmd.onlineUpdate:
            onOnlineUpdate(param);
            break;
        default:
            console.warn("unhandled system message for command [" + cmd + "]");
    }
}
function onSessionJoin(param) {
    console.log("joined");
    systemCache.session = param.session;
    systemCache.profile = param.profile;
    systemCache.members = param.members;
    systemCache.online = param.online;
    setMembers();
}
var socket;
function socketUrl() {
    var l = document.location;
    var protocol = "ws";
    if (l.protocol === "https:") {
        protocol = "wss";
    }
    return protocol + "://" + l.host + "/s";
}
function socketConnect(svc, id) {
    systemCache.currentService = svc;
    systemCache.currentID = id;
    systemCache.connectTime = Date.now();
    socket = new WebSocket(socketUrl());
    socket.onopen = function () {
        console.debug("socket connected");
        var msg = { "svc": svc, "cmd": clientCmd.connect, "param": id };
        send(msg);
    };
    socket.onmessage = function (event) {
        var msg = JSON.parse(event.data);
        onSocketMessage(msg);
    };
    socket.onerror = function (event) {
        onError("socket error: " + event.type);
    };
    socket.onclose = function () {
        onSocketClose();
    };
}
function send(msg) {
    console.log("sending message");
    console.log(msg);
    socket.send(JSON.stringify(msg));
}
function onSocketClose() {
    function disconnect(seconds) {
        if (seconds === 1) {
            console.warn("socket closed, reconnecting in a second");
        }
        else {
            console.warn("socket closed, reconnecting in " + seconds + " seconds");
        }
        setTimeout(function () { socketConnect(systemCache.currentService, systemCache.currentID); }, 4000);
    }
    if (!appUnloading) {
        var delta = Date.now() - systemCache.connectTime;
        if (delta < 2000) {
            disconnect(5);
        }
        else {
            disconnect(1);
        }
    }
}
function $(selector, context) {
    return UIkit.util.$$(selector, context);
}
function $req(selector) {
    var res = $(selector);
    if (res.length === 0) {
        console.error("no element found for selector [" + selector + "]");
    }
    return res[0];
}
function $id(id) {
    if (id.length > 0 && !(id[0] === '#')) {
        id = "#" + id;
    }
    return $req(id);
}
function init(svc, id) {
    window.onbeforeunload = function () {
        appUnloading = true;
    };
    socketConnect(svc, id);
}
function renderPoll(poll) {
    var profile = systemCache.profile;
    if (profile == null) {
        return JSX("li", null, "profile error");
    }
    else {
        return JSX("li", null,
            JSX("div", { id: "poll-status-" + poll.id, "class": "right uk-article-meta poll-status" }, poll.status.key),
            JSX("a", { "class": profile.linkColor + "-fg", href: "", onclick: "estimateCache.activePoll = '" + poll.id + "';", "data-uk-toggle": "target: #modal-poll" }, poll.title));
    }
}
function renderPolls(polls) {
    return JSX("ul", { "class": "uk-list uk-list-divider" }, polls.map(function (p) { return renderPoll(p); }));
}
function renderVote(vote) {
    return JSX("li", null,
        vote.userID,
        ": ",
        vote.choice);
}
function renderVotes(votes) {
    return JSX("ul", { "class": "uk-list uk-list-divider" }, votes.map(function (v) { return renderVote(v); }));
}
function debugMember(member) {
    return JSX("div", null,
        JSX("hr", null),
        JSX("div", null,
            "user: ",
            member.userID),
        JSX("div", null,
            "name: ",
            member.name),
        JSX("div", null,
            "role: ",
            member.role.key),
        JSX("div", null,
            "created: ",
            member.created),
        JSX("pre", null, JSON.stringify(member, null, 2)));
}
function renderMember(member) {
    var profile = systemCache.profile;
    if (profile == null) {
        return JSX("div", { "class": "uk-margin-bottom" }, "error");
    }
    else {
        var b = Math.random() >= 0.5;
        return JSX("div", null,
            JSX("div", { title: "user is offline", "class": "right uk-article-meta online-indicator" }, "offline"),
            JSX("a", { "class": profile.linkColor + "-fg", href: "", onclick: "systemCache.activeMember = '" + member.userID + "';", "data-uk-toggle": "target: #modal-member" }, member.name));
    }
}
function renderMembers(members) {
    return JSX("ul", { "class": "uk-list uk-list-divider" }, members.map(function (m) { return JSX("li", { id: "member-" + m.userID }, renderMember(m)); }));
}
