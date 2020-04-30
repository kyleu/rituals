"use strict";
function onEstimateMessage(cmd, param) {
    switch (cmd) {
        case "detail":
            break;
        case "members":
            break;
        case "polls":
            break;
        default:
            console.warn("Unhandled command [" + cmd + "] for estimate");
    }
}
var socket;
var debug = true;
var $ = UIkit.util.$$;
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
function sandbox() {
    send({ svc: "estimate", cmd: "sandbox", param: null });
}
