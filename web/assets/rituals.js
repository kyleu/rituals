"use strict";
var socket;
var debug = true;
function onMessage(msg) {
    console.log("message received", msg);
}
function socketUrl() {
    var l = document.location;
    var protocol = "ws";
    if (l.protocol == "https") {
        protocol = "wss";
    }
    return protocol + "://" + l.host + "/s";
}
function connect(k, v) {
    socket = new WebSocket(socketUrl());
    socket.onopen = function () {
        var msg = { "t": "connect", "k": k, "v": v };
        send(msg);
    };
    socket.onmessage = function (event) {
        var msg = JSON.parse(event.data);
        onMessage(msg);
    };
}
function send(msg) {
    console.log("sending message", msg);
    socket.send(JSON.stringify(msg));
}
function sandbox() {
    send({ t: "sandbox" });
}
