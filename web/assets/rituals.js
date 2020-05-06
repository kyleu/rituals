"use strict";
function onEstimateMessage(cmd, param) {
    switch (cmd) {
        case "detail":
            setEstimateDetail(param);
            break;
        case "polls":
            setPolls(param);
            break;
        case "votes":
            setVotes(param);
            break;
        default:
            console.warn("unhandled command [" + cmd + "] for estimate");
    }
}
function setEstimateDetail(param) {
    $id("model-choices-input").value = param.choices.join(", ");
    setDetail(param);
}
function setPolls(polls) {
    var detail = $id("poll-detail");
    detail.innerHTML = "";
    detail.appendChild(renderPolls(polls));
    UIkit.modal("#modal-poll").hide();
}
function setVotes(votes) {
    var detail = $id("vote-detail");
    detail.innerHTML = "";
    detail.appendChild(renderVotes(votes));
}
function onSubmitEstimateSession() {
    var title = $req("#model-title-input").value;
    var choices = $req("#model-choices-input").value;
    var msg = {
        svc: "estimate",
        cmd: "session-save",
        param: {
            title: title,
            choices: choices
        }
    };
    send(msg);
}
function onSubmitPoll() {
    var title = $req("#poll-title-input").value;
    var msg = {
        svc: "estimate",
        cmd: "new-poll-save",
        param: {
            title: title
        }
    };
    send(msg);
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
        case "poll":
            var pollInput_1 = $req("#poll-title-input");
            pollInput_1.value = "";
            delay(function () { return pollInput_1.focus(); });
            break;
        default:
            console.debug("unhandled modal [" + key + "]");
    }
}
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
            if (child.nodeType == null) {
                child = document.createTextNode(child.toString());
            }
            e.appendChild(child);
        }
    }
    return e;
}
var currentMembers = [];
var currentOnline = [];
var activeMember = null;
function setMembers(members) {
    currentMembers = members;
    UIkit.modal("#modal-self").hide();
    function isSelf(x) {
        if (activeProfile == null) {
            return false;
        }
        return x.userID == activeProfile.userID;
    }
    var self = members.filter(isSelf);
    if (self.length == 1) {
        $req("#member-self .member-name").innerText = self[0].name;
        $req("#self-name-input").value = self[0].name;
        $req("#member-self .member-role").innerText = self[0].role.key;
    }
    else if (self.length == 0) {
        console.warn("self not found among members");
    }
    else {
        console.warn("multiple self entries found among members");
    }
    var others = members.filter(function (x) { return !isSelf(x); });
    var detail = $id("member-detail");
    detail.innerHTML = "";
    detail.appendChild(renderMembers(others));
    renderOnline();
}
function setOnline(users) {
    currentOnline = users;
    renderOnline();
}
function renderOnline() {
    for (var _i = 0, currentMembers_1 = currentMembers; _i < currentMembers_1.length; _i++) {
        var member = currentMembers_1[_i];
        var els = $("#online-status-" + member.userID);
        if (els.length == 1) {
            if (currentOnline.indexOf(member.userID) == -1) {
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
        svc: "system",
        cmd: "member-name-save",
        param: {
            name: name,
            choice: choice
        }
    };
    send(msg);
}
function viewActiveMember() {
    if (activeMember == null) {
        console.warn("no active member");
        return;
    }
    var curr = currentMembers.filter(function (x) { return x.userID == activeMember; });
    if (curr.length != 1) {
        console.log("cannot load member [" + activeMember + "]");
        return;
    }
    var member = curr[0];
    $req("#member-modal-name").innerText = member.name;
    $req("#member-modal-role").innerText = member.role.key;
}
var socket;
var debug = true;
function onSocketMessage(msg) {
    console.log("message received");
    console.log(msg);
    switch (msg.svc) {
        case "system":
            onSystemMessage(msg.cmd, msg.param);
            break;
        case "estimate":
            onEstimateMessage(msg.cmd, msg.param);
            break;
        default:
            console.warn("unhandled message for service [" + msg.svc + "]");
    }
}
function setDetail(param) {
    $id("model-title").innerText = param.title;
    $id("model-title-input").value = param.title;
    UIkit.modal("#modal-session").hide();
}
var activeProfile = null;
function setProfile(profile) {
    activeProfile = profile;
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
        case "profile":
            setProfile(param);
            break;
        case "online":
            setOnline(param);
            break;
        case "members":
            setMembers(param);
            break;
        case "error":
            onError("server error: " + param);
            break;
        default:
            console.warn("unhandled system message for command [" + cmd + "]");
    }
}
function socketUrl() {
    var l = document.location;
    var protocol = "ws";
    if (l.protocol === "https:") {
        protocol = "wss";
    }
    return protocol + "://" + l.host + "/s";
}
var currentService = "";
var currentId = "";
var connectTime = 0;
function socketConnect(svc, id) {
    currentService = svc;
    currentId = id;
    connectTime = Date.now();
    socket = new WebSocket(socketUrl());
    socket.onopen = function () {
        console.debug("socket connected");
        var msg = { "svc": svc, "cmd": "connect", "param": id };
        send(msg);
    };
    socket.onmessage = function (event) {
        var msg = JSON.parse(event.data);
        onSocketMessage(msg);
    };
    socket.onerror = function (event) {
        onError("socket error: " + event.type);
    };
    socket.onclose = function (event) {
        onSocketClose();
    };
}
function send(msg) {
    console.log("sending message");
    console.log(msg);
    socket.send(JSON.stringify(msg));
}
function onSocketClose() {
    if (!appUnloading) {
        var delta = Date.now() - connectTime;
        if (delta < 2000) {
            console.warn("socket closed immediately, reconnecting in 4 seconds");
            setTimeout(function () { socketConnect(currentService, currentId); }, 4000);
        }
        else {
            console.warn("socket closed, reconnecting in a second");
            setTimeout(function () { socketConnect(currentService, currentId); }, 1000);
        }
    }
}
function $(selector, context) {
    return UIkit.util.$$(selector, context);
}
function $req(selector) {
    var res = $(selector);
    if (res.length == 0) {
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
var appInitialized = false;
var appUnloading = false;
function init(svc, id) {
    appInitialized = true;
    window.onbeforeunload = function () {
        appUnloading = true;
    };
    socketConnect(svc, id);
}
function debugPoll(poll) {
    return JSX("div", null,
        JSX("hr", null),
        JSX("div", null,
            "id: ",
            poll.id),
        JSX("div", null,
            "idx: ",
            poll.idx),
        JSX("div", null,
            "author: ",
            poll.author),
        JSX("div", null,
            "title: ",
            poll.title),
        JSX("div", null,
            "status: ",
            poll.status.key),
        JSX("div", null,
            "finalVote: ",
            poll.finalVote),
        JSX("pre", null, JSON.stringify(poll, null, 2)));
}
function renderPoll(poll) {
    var profile = activeProfile;
    if (profile == null) {
        return JSX("li", null, "error");
    }
    else {
        return JSX("li", null,
            JSX("a", { "class": profile.linkColor + "-fg", href: "" }, poll.title));
    }
}
function renderPolls(polls) {
    return JSX("ul", { "class": "uk-list uk-list-divider" }, polls.map(function (p) { return renderPoll(p); }));
}
function debugVote(vote) {
    return JSX("pre", null, JSON.stringify(vote, null, 2));
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
    var profile = activeProfile;
    if (profile == null) {
        return JSX("div", { "class": "uk-margin-bottom" }, "error");
    }
    else {
        var b = Math.random() >= 0.5;
        return JSX("li", null,
            JSX("div", { title: "user is offline", id: "online-status-" + member.userID, "class": "right uk-article-meta online-indicator" }, "offline"),
            JSX("a", { "class": profile.linkColor + "-fg", href: "", onclick: "activeMember = '" + member.userID + "';", "data-uk-toggle": "target: #modal-member" }, member.name));
    }
}
function renderMembers(members) {
    return JSX("ul", { "class": "uk-list uk-list-divider" }, members.map(function (m) { return renderMember(m); }));
}
