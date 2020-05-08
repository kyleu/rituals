"use strict";
var EstimateCache = /** @class */ (function () {
    function EstimateCache() {
        this.stories = [];
        this.votes = [];
    }
    EstimateCache.prototype.activeVotes = function () {
        var _this = this;
        if (this.activeStory === undefined) {
            console.log("!!!");
            return [];
        }
        return this.votes.filter(function (x) { return x.storyID == _this.activeStory; });
    };
    return EstimateCache;
}());
var estimateCache = new EstimateCache();
function onEstimateMessage(cmd, param) {
    switch (cmd) {
        case serverCmd.sessionJoined:
            var sj = param;
            onSessionJoin(sj);
            setEstimateDetail(sj.session);
            setStories(sj.stories);
            setVotes(sj.votes);
            break;
        case serverCmd.sessionUpdate:
            setEstimateDetail(param);
            break;
        case serverCmd.storyUpdate:
            onStoryUpdate(param);
            break;
        case serverCmd.storyStatusChange:
            onStoryStatusChange(param);
            break;
        case serverCmd.voteUpdate:
            onVoteUpdate(param);
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
function getStory(id) {
    return estimateCache.stories.filter(function (x) { return x.id === id; }).pop();
}
function onStoryUpdate(story) {
    var x = estimateCache.stories;
    x = x.filter(function (p) { return p.id !== story.id; });
    x.push(story);
    x = x.sort(function (l, r) { return (l.idx > r.idx ? 1 : -1); });
    setStories(x);
}
function delay(f) {
    setTimeout(f, 250);
}
function openModal(key) {
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
        case "add-story":
            var storyInput_1 = $req("#story-title-input");
            storyInput_1.value = "";
            delay(function () { return storyInput_1.focus(); });
            break;
        case "story":
            viewActiveStory();
            break;
        default:
            console.debug("unhandled modal [" + key + "]");
    }
    UIkit.modal("#modal-" + key).show();
}
var debug = true;
var appUnloading = false;
var SystemCache = /** @class */ (function () {
    function SystemCache() {
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
    addStory: "add-story",
    updateStory: "update-story",
    setStoryStatus: "set-story-status",
    submitVote: "submit-vote"
};
var serverCmd = {
    error: "error",
    pong: "pong",
    sessionJoined: "session-joined",
    sessionUpdate: "session-update",
    memberUpdate: "member-update",
    onlineUpdate: "online-update",
    storyUpdate: "story-update",
    storyStatusChange: "story-status-change",
    voteUpdate: "vote-update"
};
function JSX(tag, attrs) {
    var e = document.createElement(tag);
    for (var name_1 in attrs) {
        if (name_1 && attrs.hasOwnProperty(name_1)) {
            var v = attrs[name_1];
            if (v === true) {
                e.setAttribute(name_1, name_1);
            }
            else if (v !== false && v !== null && v !== undefined) {
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
    if (systemCache.profile === undefined) {
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
    x = x.filter(function (m) { return m.userID !== member.userID; });
    x.push(member);
    x = x.sort(function (l, r) { return (l.name > r.name) ? 1 : -1; });
    systemCache.members = x;
    setMembers();
    if (estimateCache.activeStory) {
        viewActiveVotes();
    }
}
function onOnlineUpdate(update) {
    if (update.connected) {
        if (systemCache.online.indexOf(update.userID) === -1) {
            systemCache.online.push(update.userID);
        }
    }
    else {
        systemCache.online = systemCache.online.filter(function (x) { return x !== update.userID; });
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
    if (systemCache.activeMember === undefined) {
        console.warn("no active member");
        return undefined;
    }
    var curr = systemCache.members.filter(function (x) { return x.userID === systemCache.activeMember; });
    if (curr.length !== 1) {
        console.log("cannot load active member [" + systemCache.activeMember + "]");
        return undefined;
    }
    return curr[0];
}
function viewActiveMember() {
    var member = getActiveMember();
    if (member === undefined) {
        return;
    }
    $req("#member-modal-name").innerText = member.name;
    $req("#member-modal-role").innerText = member.role.key;
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
    UIkit.notification(err, { status: "danger", pos: "top-right" });
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
        var msg = { svc: svc, cmd: clientCmd.connect, param: id };
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
        setTimeout(function () {
            socketConnect(systemCache.currentService, systemCache.currentID);
        }, 4000);
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
function setStories(stories) {
    estimateCache.stories = stories;
    var detail = $id("story-detail");
    detail.innerHTML = "";
    detail.appendChild(renderStories(stories));
    UIkit.modal("#modal-add-story").hide();
}
function setVotes(votes) {
    estimateCache.votes = votes;
    if (estimateCache.activeStory) {
        viewActiveVotes();
    }
}
function onSubmitStory() {
    var title = $req("#story-title-input").value;
    var msg = {
        svc: services.estimate,
        cmd: clientCmd.addStory,
        param: { title: title }
    };
    send(msg);
}
function getActiveStory() {
    if (estimateCache.activeStory === undefined) {
        console.warn("no active story");
        return undefined;
    }
    var curr = estimateCache.stories.filter(function (x) { return x.id === estimateCache.activeStory; });
    if (curr.length !== 1) {
        console.warn("cannot load active story [" + estimateCache.activeStory + "]");
        return undefined;
    }
    return curr[0];
}
function viewActiveStory() {
    var story = getActiveStory();
    if (story === undefined) {
        console.log("no active story");
        return;
    }
    $req("#story-title").innerText = story.title;
    viewStoryStatus(story.status.key);
}
function viewStoryStatus(status) {
    switch (status) {
        case "pending":
            break;
        case "active":
            viewActiveVotes();
            break;
        case "complete":
            break;
    }
    for (var _i = 0, _a = $(".story-status-section"); _i < _a.length; _i++) {
        var el = _a[_i];
        var s = el.id.substr(el.id.lastIndexOf("-") + 1);
        if (s === status) {
            el.classList.add("active");
        }
        else {
            el.classList.remove("active");
        }
    }
}
function onVoteUpdate(vote) {
    var x = estimateCache.votes;
    x = x.filter(function (v) { return v.userID != vote.userID || v.storyID != vote.storyID; });
    x.push(vote);
    estimateCache.votes = x;
    if (vote.storyID === estimateCache.activeStory) {
        viewActiveVotes();
    }
}
function viewActiveVotes() {
    var votes = estimateCache.activeVotes();
    var activeVote = votes.filter(function (v) { return v.userID === systemCache.profile.userID; }).pop();
    var m = $id("story-vote-members");
    m.innerHTML = "";
    m.appendChild(renderVoteMembers(systemCache.members, votes));
    var c = $id("story-vote-choices");
    c.innerHTML = "";
    c.appendChild(renderVoteChoices(estimateCache.detail.choices, activeVote === null || activeVote === void 0 ? void 0 : activeVote.choice));
}
function requestStoryStatus(s) {
    var story = getActiveStory();
    if (story === undefined) {
        console.log("no active story");
        return;
    }
    var msg = {
        svc: services.estimate,
        cmd: clientCmd.setStoryStatus,
        param: { storyID: story.id, status: s }
    };
    send(msg);
}
function onStoryStatusChange(u) {
    $req("#story-" + u.storyID + " .story-status").innerText = u.status.key;
    viewStoryStatus(u.status.key);
}
function onSubmitVote(choice) {
    var msg = {
        svc: services.estimate,
        cmd: clientCmd.submitVote,
        param: { storyID: estimateCache.activeStory, choice: choice }
    };
    send(msg);
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
    if (id.length > 0 && !(id[0] === "#")) {
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
function renderStory(story) {
    var profile = systemCache.profile;
    if (profile === undefined) {
        return JSX("li", null, "profile error");
    }
    else {
        return JSX("li", { id: "story-" + story.id },
            JSX("div", { "class": "right uk-article-meta story-status" }, story.status.key),
            JSX("a", { "class": profile.linkColor + "-fg", href: "", onclick: "estimateCache.activeStory = '" + story.id + "';openModal('story');return false;" }, story.title));
    }
}
function renderStories(stories) {
    if (stories.length === 0) {
        return JSX("div", null,
            JSX("button", { "class": "uk-button uk-button-default", onclick: "openModal('add-story');", type: "button" }, "Add Story"));
    }
    else {
        return JSX("ul", { "class": "uk-list uk-list-divider" }, stories.map(function (s) { return renderStory(s); }));
    }
}
function renderVoteMember(member, hasVote) {
    return JSX("div", { "class": "vote-member", title: member.name + " has " + (hasVote ? "voted" : "not voted") },
        JSX("div", null,
            JSX("span", { "data-uk-icon": "icon: " + (hasVote ? "check" : "minus") + "; ratio: 1.6" })),
        member.name);
}
function renderVoteMembers(members, votes) {
    return JSX("div", { "class": "uk-flex uk-flex-wrap uk-flex-around" }, members.map(function (m) { return renderVoteMember(m, votes.filter(function (v) { return v.userID == m.userID; }).length > 0); }));
}
function renderVoteChoices(choices, choice) {
    return JSX("div", { "class": "uk-flex uk-flex-wrap uk-flex-center" }, choices.map(function (c) { return JSX("div", { "class": "vote-choice uk-border-circle uk-box-shadow-hover-medium" + (c === choice ? " active" : ""), onclick: "onSubmitVote('" + c + "');" }, c); }));
}
function renderMember(member) {
    var profile = systemCache.profile;
    if (profile === undefined) {
        return JSX("div", { "class": "uk-margin-bottom" }, "error");
    }
    else {
        var b = Math.random() >= 0.5;
        return JSX("div", null,
            JSX("div", { title: "user is offline", "class": "right uk-article-meta online-indicator" }, "offline"),
            JSX("a", { "class": profile.linkColor + "-fg", href: "", onclick: "systemCache.activeMember = '" + member.userID + "';openModal('member');return false;" }, member.name));
    }
}
function renderMembers(members) {
    if (members.length === 0) {
        return JSX("div", null,
            JSX("button", { "class": "uk-button uk-button-default", onclick: "openModal('invite');", type: "button" }, "Invite Members"));
    }
    else {
        return JSX("ul", { "class": "uk-list uk-list-divider" }, members.map(function (m) { return JSX("li", { id: "member-" + m.userID }, renderMember(m)); }));
    }
}
