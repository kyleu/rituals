"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.socketInit = void 0;
// Content managed by Project Forge, see [projectforge.md] for details.
function socketInit() {
    window.rituals.Socket = Socket;
}
exports.socketInit = socketInit;
let appUnloading = false;
class Socket {
    constructor(debug, open, recv, err, url) {
        this.debug = debug;
        this.open = open;
        this.recv = recv;
        this.err = err;
        this.url = socketUrl(url);
        this.connected = false;
        this.pauseSeconds = 1;
        this.pendingMessages = [];
        this.connect();
    }
    connect() {
        this.connectTime = Date.now();
        this.sock = new WebSocket(socketUrl(this.url));
        const s = this;
        this.sock.onopen = () => {
            s.connected = true;
            s.pendingMessages.forEach(s.send);
            s.pendingMessages = [];
            if (s.debug) {
                console.log("WebSocket connected");
            }
            s.open("todo");
        };
        this.sock.onmessage = (event) => {
            const msg = JSON.parse(event.data);
            if (s.debug) {
                console.debug("in", msg);
            }
            s.recv(msg);
        };
        this.sock.onerror = (event) => () => {
            s.err("socket", event.type);
        };
        this.sock.onclose = () => {
            s.connected = false;
            const elapsed = s.connectTime ? Date.now() - s.connectTime : 0;
            if (0 < elapsed && elapsed < 2000) {
                s.pauseSeconds = s.pauseSeconds * 2;
                if (s.debug) {
                    console.debug(`socket closed immediately, reconnecting in ${s.pauseSeconds} seconds`);
                }
                setTimeout(() => {
                    s.connect();
                }, s.pauseSeconds * 1000);
            }
            else {
                console.debug("socket closed after [" + elapsed + "ms]");
                s.connect();
            }
        };
    }
    disconnect() {
    }
    send(msg) {
        if (this.debug) {
            console.debug("out", msg);
        }
        if (!this.sock) {
            throw "not initialized";
        }
        if (this.connected) {
            const m = JSON.stringify(msg, null, 2);
            this.sock.send(m);
        }
        else {
            this.pendingMessages.push(msg);
        }
    }
}
function socketUrl(u) {
    if (!u) {
        u = "/connect";
    }
    if (u.indexOf("ws") == 0) {
        return u;
    }
    const l = document.location;
    let protocol = "ws";
    if (l.protocol === "https:") {
        protocol = "wss";
    }
    if (u.indexOf("/") != 0) {
        u = "/" + u;
    }
    return protocol + `://${l.host}${u}`;
}
