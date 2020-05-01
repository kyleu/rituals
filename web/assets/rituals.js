"use strict";
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
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
    for (var _i = 0, polls_1 = polls; _i < polls_1.length; _i++) {
        var poll = polls_1[_i];
        console.log(poll);
        detail.appendChild(renderPoll(poll));
    }
}
function setVotes(votes) {
    var detail = $id("vote-detail");
    detail.innerHTML = "";
    for (var _i = 0, votes_1 = votes; _i < votes_1.length; _i++) {
        var vote = votes_1[_i];
        console.log(vote);
        detail.appendChild(renderVote(vote));
    }
}
function nonNull(val, fallback) {
    return Boolean(val) ? val : fallback;
}
function JSXparseChildren(children) {
    return children.map(function (child) {
        if (typeof child === 'string') {
            return document.createTextNode(child);
        }
        return child;
    });
}
function JSXparseNode(element, properties, children) {
    var el = document.createElement(element);
    Object.keys(nonNull(properties, {})).forEach(function (key) {
        el[key] = properties[key];
    });
    JSXparseChildren(children).forEach(function (child) {
        el.appendChild(child);
    });
    return el;
}
function JSX(element, properties) {
    var children = [];
    for (var _i = 2; _i < arguments.length; _i++) {
        children[_i - 2] = arguments[_i];
    }
    if (typeof element === 'function') {
        return element(__assign({}, nonNull(properties, {}), { children: children }));
    }
    return JSXparseNode(element, properties, children);
}
function setMembers(members) {
    var detail = $id("member-detail");
    detail.innerHTML = "";
    for (var _i = 0, members_1 = members; _i < members_1.length; _i++) {
        var member = members_1[_i];
        console.log(member);
        detail.appendChild(renderMember(member));
    }
}
var socket;
var debug = true;
function onMessage(msg) {
    console.log("message received");
    console.log(msg);
    switch (msg.svc) {
        case "estimate":
            onEstimateMessage(msg.cmd, msg.param);
            break;
        default:
            console.warn("Unhandled message for service [" + msg.svc + "]");
    }
}
function setDetail(param) {
    $id("model-title").innerText = param.title + "!!!!";
}
function sandbox() {
    send({ svc: "estimate", cmd: "sandbox", param: null });
}
function socketUrl() {
    var l = document.location;
    var protocol = "ws";
    if (l.protocol === "https:") {
        protocol = "wss";
    }
    return protocol + "://" + l.host + "/s";
}
function connect(svc, value) {
    socket = new WebSocket(socketUrl());
    socket.onopen = function () {
        var msg = { "svc": svc, "cmd": "connect", "param": value };
        send(msg);
    };
    socket.onmessage = function (event) {
        var msg = JSON.parse(event.data);
        onMessage(msg);
    };
}
function send(msg) {
    console.log("sending message");
    console.log(msg);
    socket.send(JSON.stringify(msg));
}
function renderMember(member) {
    return JSX("pre", null, JSON.stringify(member, null, 2));
}
function renderPoll(poll) {
    return JSX("pre", null, JSON.stringify(poll, null, 2));
}
function renderVote(vote) {
    return JSX("pre", null, JSON.stringify(vote, null, 2));
}
var $ = UIkit.util.$$;
function $id(id) {
    if (id.length > 0 && !(id[0] === '#')) {
        id = "#" + id;
    }
    return $(id)[0];
}
