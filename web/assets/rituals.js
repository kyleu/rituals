"use strict";
function onEstimateMessage(cmd, param) {
    switch (cmd) {
        case "detail":
            setEstimateDetail(param);
            break;
        case "members":
            setMembers(param);
            break;
        case "polls":
            setPolls(param);
            break;
        case "votes":
            setVotes(param);
            break;
        default:
            console.warn("Unhandled command [" + cmd + "] for estimate");
    }
}
function setEstimateDetail(param) {
    setDetail(param);
}
function setPolls(polls) {
    var detail = $id("poll-detail");
    detail.innerHTML = "";
    detail.appendChild(renderPolls(polls));
}
function setVotes(votes) {
    var detail = $id("vote-detail");
    detail.innerHTML = "";
    detail.appendChild(renderVotes(votes));
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
function setMembers(members) {
    var detail = $id("member-detail");
    detail.innerHTML = "";
    detail.appendChild(renderMembers(members));
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
            console.warn("Unhandled message for service [" + msg.svc + "]");
    }
}
function setDetail(param) {
    $id("model-title").innerText = param.title;
}
var activeProfile = null;
function setProfile(profile) {
    activeProfile = profile;
}
function onError(err) {
    console.error("server error: " + err);
}
function onSystemMessage(cmd, param) {
    switch (cmd) {
        case "profile":
            setProfile(param);
            break;
        case "error":
            onError(param);
            break;
        default:
            console.warn("Unhandled system message for command [" + cmd + "]");
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
        var msg = { "svc": svc, "cmd": "connect", "param": id };
        send(msg);
    };
    socket.onmessage = function (event) {
        var msg = JSON.parse(event.data);
        onSocketMessage(msg);
    };
    socket.onerror = function (event) {
        onSocketError(event.type);
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
function onSocketError(err) {
    console.error("socket error: " + err);
}
function onSocketClose() {
    var delta = Date.now() - connectTime;
    if (delta < 2000) {
        console.warn("socket closed immediately, reconnecting in 10 seconds");
        setTimeout(function () { socketConnect(currentService, currentId); }, 10000);
    }
    else {
        console.warn("socket closed, reconnecting in 2 seconds");
        setTimeout(function () { socketConnect(currentService, currentId); }, 2000);
    }
}
function renderMember(member) {
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
function renderMembers(members) {
    return JSX("div", null, members.map(function (m) { return renderMember(m); }));
}
function renderPoll(poll) {
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
function renderPolls(polls) {
    return JSX("div", null, polls.map(function (p) { return renderPoll(p); }));
}
function renderVote(vote) {
    return JSX("pre", null, JSON.stringify(vote, null, 2));
}
function renderVotes(votes) {
    return JSX("div", null, votes.map(function (v) { return renderVote(v); }));
}
var $ = UIkit.util.$$;
function $id(id) {
    if (id.length > 0 && !(id[0] === '#')) {
        id = "#" + id;
    }
    return $(id)[0];
}
