"use strict";
var estimate;
(function (estimate) {
    var Cache = /** @class */ (function () {
        function Cache() {
            this.stories = [];
            this.votes = [];
        }
        Cache.prototype.activeVotes = function () {
            var _this = this;
            if (this.activeStory === undefined) {
                console.log("!!!");
                return [];
            }
            return this.votes.filter(function (x) { return x.storyID == _this.activeStory; });
        };
        return Cache;
    }());
    estimate.cache = new Cache();
    function onEstimateMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.estimate, param);
                break;
            case command.server.sessionJoined:
                var sj = param;
                rituals.onSessionJoin(sj);
                setEstimateDetail(sj.session);
                story.setStories(sj.stories);
                story.setVotes(sj.votes);
                break;
            case command.server.sessionUpdate:
                setEstimateDetail(param);
                break;
            case command.server.storyUpdate:
                onStoryUpdate(param);
                break;
            case command.server.storyStatusChange:
                story.onStoryStatusChange(param);
                break;
            case command.server.voteUpdate:
                story.onVoteUpdate(param);
                break;
            default:
                console.warn("unhandled command [" + cmd + "] for estimate");
        }
    }
    estimate.onEstimateMessage = onEstimateMessage;
    function setEstimateDetail(detail) {
        estimate.cache.detail = detail;
        util.req("#model-choices-input").value = detail.choices.join(", ");
        rituals.setDetail(detail);
    }
    function onSubmitEstimateSession() {
        var title = util.req("#model-title-input").value;
        var choices = util.req("#model-choices-input").value;
        var msg = {
            svc: services.estimate,
            cmd: command.client.updateSession,
            param: {
                title: title,
                choices: choices
            }
        };
        socket.send(msg);
    }
    estimate.onSubmitEstimateSession = onSubmitEstimateSession;
    function onStoryUpdate(s) {
        var x = estimate.cache.stories;
        x = x.filter(function (p) { return p.id !== s.id; });
        x.push(s);
        x = x.sort(function (l, r) { return (l.idx > r.idx ? 1 : -1); });
        story.setStories(x);
    }
})(estimate || (estimate = {}));
var events;
(function (events) {
    function delay(f) {
        setTimeout(f, 250);
    }
    function openModal(key, id) {
        switch (key) {
            case "self":
                var selfInput_1 = util.req("#self-name-input");
                selfInput_1.value = util.req("#member-self .member-name").innerText;
                delay(function () { return selfInput_1.focus(); });
                break;
            case "session":
                var sessionInput_1 = util.req("#model-title-input");
                sessionInput_1.value = util.req("#model-title").innerText;
                delay(function () { return sessionInput_1.focus(); });
                break;
            case "invite":
                break;
            case "member":
                system.cache.activeMember = id;
                member.viewActiveMember();
                break;
            case "add-story":
                var storyInput_1 = util.req("#story-title-input");
                storyInput_1.value = "";
                delay(function () { return storyInput_1.focus(); });
                break;
            case "story":
                estimate.cache.activeStory = id;
                story.viewActiveStory();
                break;
            default:
                console.debug("unhandled modal [" + key + "]");
        }
        UIkit.modal("#modal-" + key).show();
        return false;
    }
    events.openModal = openModal;
})(events || (events = {}));
var system;
(function (system) {
    var Cache = /** @class */ (function () {
        function Cache() {
            this.currentService = "";
            this.currentID = "";
            this.connectTime = 0;
            this.members = [];
            this.online = [];
        }
        return Cache;
    }());
    system.cache = new Cache();
})(system || (system = {}));
var services;
(function (services) {
    services.system = "system";
    services.estimate = "estimate";
    services.standup = "standup";
    services.retro = "retro";
})(services || (services = {}));
var command;
(function (command) {
    command.client = {
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
    command.server = {
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
})(command || (command = {}));
// noinspection JSUnusedGlobalSymbols
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
var member;
(function (member_1) {
    function isSelf(x) {
        if (system.cache.profile === undefined) {
            return false;
        }
        return x.userID === system.cache.profile.userID;
    }
    function setMembers() {
        var self = system.cache.members.filter(isSelf);
        if (self.length === 1) {
            util.req("#member-self .member-name").innerText = self[0].name;
            util.req("#self-name-input").value = self[0].name;
            util.req("#member-self .member-role").innerText = self[0].role.key;
        }
        else if (self.length === 0) {
            console.warn("self not found among members");
        }
        else {
            console.warn("multiple self entries found among members");
        }
        var others = system.cache.members.filter(function (x) { return !isSelf(x); });
        var detail = util.req("#member-detail");
        detail.innerHTML = "";
        detail.appendChild(member_1.renderMembers(others));
        renderOnline();
    }
    member_1.setMembers = setMembers;
    function onMemberUpdate(member) {
        if (isSelf(member)) {
            UIkit.modal("#modal-self").hide();
        }
        var x = system.cache.members;
        x = x.filter(function (m) { return m.userID !== member.userID; });
        x.push(member);
        x = x.sort(function (l, r) { return (l.name > r.name) ? 1 : -1; });
        system.cache.members = x;
        setMembers();
        if (estimate.cache.activeStory) {
            story.viewVotes();
        }
    }
    member_1.onMemberUpdate = onMemberUpdate;
    function onOnlineUpdate(update) {
        if (update.connected) {
            if (system.cache.online.indexOf(update.userID) === -1) {
                system.cache.online.push(update.userID);
            }
        }
        else {
            system.cache.online = system.cache.online.filter(function (x) { return x !== update.userID; });
        }
        renderOnline();
    }
    member_1.onOnlineUpdate = onOnlineUpdate;
    function renderOnline() {
        for (var _i = 0, _a = system.cache.members; _i < _a.length; _i++) {
            var member_2 = _a[_i];
            var els = util.els("#member-" + member_2.userID + " .online-indicator");
            if (els.length === 1) {
                if (system.cache.online.indexOf(member_2.userID) === -1) {
                    els[0].classList.add("offline");
                }
                else {
                    els[0].classList.remove("offline");
                }
            }
        }
    }
    function onSubmitSelf() {
        var name = util.req("#self-name-input").value;
        var choice = util.req("#self-name-choice-global").checked ? "global" : "local";
        var msg = {
            svc: services.system,
            cmd: command.client.updateProfile,
            param: {
                name: name,
                choice: choice
            }
        };
        socket.send(msg);
    }
    member_1.onSubmitSelf = onSubmitSelf;
    function getActiveMember() {
        if (system.cache.activeMember === undefined) {
            console.warn("no active member");
            return undefined;
        }
        var curr = system.cache.members.filter(function (x) { return x.userID === system.cache.activeMember; });
        if (curr.length !== 1) {
            console.log("cannot load active member [" + system.cache.activeMember + "]");
            return undefined;
        }
        return curr[0];
    }
    function viewActiveMember() {
        var member = getActiveMember();
        if (member === undefined) {
            return;
        }
        util.req("#member-modal-name").innerText = member.name;
        util.req("#member-modal-role").innerText = member.role.key;
    }
    member_1.viewActiveMember = viewActiveMember;
})(member || (member = {}));
var retro;
(function (retro) {
    var Cache = /** @class */ (function () {
        function Cache() {
        }
        return Cache;
    }());
    var cache = new Cache();
    function onRetroMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.retro, param);
                break;
            case command.server.sessionJoined:
                var sj = param;
                rituals.onSessionJoin(sj);
                setRetroDetail(sj.session);
                break;
            case command.server.sessionUpdate:
                setRetroDetail(param);
                break;
            default:
                console.warn("unhandled command [" + cmd + "] for retro");
        }
    }
    retro.onRetroMessage = onRetroMessage;
    function setRetroDetail(detail) {
        cache.detail = detail;
        rituals.setDetail(detail);
    }
    function onSubmitRetroSession() {
        var title = util.req("#model-title-input").value;
        var msg = {
            svc: services.retro,
            cmd: command.client.updateSession,
            param: {
                title: title
            }
        };
        socket.send(msg);
    }
    retro.onSubmitRetroSession = onSubmitRetroSession;
})(retro || (retro = {}));
var rituals;
(function (rituals) {
    function onSocketMessage(msg) {
        console.log("message received");
        console.log(msg);
        switch (msg.svc) {
            case services.system:
                onSystemMessage(msg.cmd, msg.param);
                break;
            case services.estimate:
                estimate.onEstimateMessage(msg.cmd, msg.param);
                break;
            case services.standup:
                standup.onStandupMessage(msg.cmd, msg.param);
                break;
            case services.retro:
                retro.onRetroMessage(msg.cmd, msg.param);
                break;
            default:
                console.warn("unhandled message for service [" + msg.svc + "]");
        }
    }
    rituals.onSocketMessage = onSocketMessage;
    function setDetail(session) {
        system.cache.session = session;
        util.req("#model-title").innerText = session.title;
        util.req("#model-title-input").value = session.title;
        var items = util.els("#navbar .uk-navbar-item");
        if (items.length > 0) {
            items[items.length - 1].innerText = session.title;
        }
        UIkit.modal("#modal-session").hide();
    }
    rituals.setDetail = setDetail;
    function onError(svc, err) {
        console.warn(svc + ": " + err);
        var idx = err.lastIndexOf(":");
        if (idx > -1) {
            err = err.substr(idx + 1);
        }
        UIkit.notification(svc + " error: " + err, { status: "danger", pos: "top-right" });
    }
    rituals.onError = onError;
    function onSystemMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                onError("system", param);
                break;
            case command.server.memberUpdate:
                member.onMemberUpdate(param);
                break;
            case command.server.onlineUpdate:
                member.onOnlineUpdate(param);
                break;
            default:
                console.warn("unhandled system message for command [" + cmd + "]");
        }
    }
    function onSessionJoin(param) {
        system.cache.session = param.session;
        system.cache.profile = param.profile;
        system.cache.members = param.members;
        system.cache.online = param.online;
        member.setMembers();
    }
    rituals.onSessionJoin = onSessionJoin;
    function init(svc, id) {
        window.onbeforeunload = function () {
            socket.setAppUnloading();
        };
        socket.socketConnect(svc, id);
    }
    rituals.init = init;
})(rituals || (rituals = {}));
var socket;
(function (socket_1) {
    var socket;
    var appUnloading = false;
    function socketUrl() {
        var l = document.location;
        var protocol = "ws";
        if (l.protocol === "https:") {
            protocol = "wss";
        }
        return protocol + "://" + l.host + "/s";
    }
    function setAppUnloading() {
        appUnloading = true;
    }
    socket_1.setAppUnloading = setAppUnloading;
    function socketConnect(svc, id) {
        system.cache.currentService = svc;
        system.cache.currentID = id;
        system.cache.connectTime = Date.now();
        socket = new WebSocket(socketUrl());
        socket.onopen = function () {
            console.debug("socket connected");
            var msg = { svc: svc, cmd: command.client.connect, param: id };
            send(msg);
        };
        socket.onmessage = function (event) {
            var msg = JSON.parse(event.data);
            rituals.onSocketMessage(msg);
        };
        socket.onerror = function (event) {
            rituals.onError("socket", event.type);
        };
        socket.onclose = function () {
            onSocketClose();
        };
    }
    socket_1.socketConnect = socketConnect;
    function send(msg) {
        console.log("sending message");
        console.log(msg);
        socket.send(JSON.stringify(msg));
    }
    socket_1.send = send;
    function onSocketClose() {
        function disconnect(seconds) {
            if (seconds === 1) {
                console.warn("socket closed, reconnecting in a second");
            }
            else {
                console.warn("socket closed, reconnecting in " + seconds + " seconds");
            }
            setTimeout(function () {
                socketConnect(system.cache.currentService, system.cache.currentID);
            }, seconds * 1000);
        }
        if (!appUnloading) {
            var delta = Date.now() - system.cache.connectTime;
            if (delta < 2000) {
                disconnect(6);
            }
            else {
                disconnect(1);
            }
        }
    }
})(socket || (socket = {}));
var standup;
(function (standup) {
    var Cache = /** @class */ (function () {
        function Cache() {
        }
        return Cache;
    }());
    var cache = new Cache();
    function onStandupMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.standup, param);
                break;
            case command.server.sessionJoined:
                var sj = param;
                rituals.onSessionJoin(sj);
                setStandupDetail(sj.session);
                break;
            case command.server.sessionUpdate:
                setStandupDetail(param);
                break;
            default:
                console.warn("unhandled command [" + cmd + "] for standup");
        }
    }
    standup.onStandupMessage = onStandupMessage;
    function setStandupDetail(detail) {
        cache.detail = detail;
        rituals.setDetail(detail);
    }
    function onSubmitStandupSession() {
        var title = util.req("#model-title-input").value;
        var msg = {
            svc: services.standup,
            cmd: command.client.updateSession,
            param: {
                title: title
            }
        };
        socket.send(msg);
    }
    standup.onSubmitStandupSession = onSubmitStandupSession;
})(standup || (standup = {}));
var story;
(function (story_1) {
    function setStories(stories) {
        estimate.cache.stories = stories;
        var detail = util.req("#story-detail");
        detail.innerHTML = "";
        detail.appendChild(story_1.renderStories(stories));
        UIkit.modal("#modal-add-story").hide();
    }
    story_1.setStories = setStories;
    function setVotes(votes) {
        estimate.cache.votes = votes;
        if (estimate.cache.activeStory) {
            viewVotes();
        }
    }
    story_1.setVotes = setVotes;
    function onSubmitStory() {
        var title = util.req("#story-title-input").value;
        var msg = {
            svc: services.estimate,
            cmd: command.client.addStory,
            param: { title: title }
        };
        socket.send(msg);
        return false;
    }
    story_1.onSubmitStory = onSubmitStory;
    function getActiveStory() {
        if (estimate.cache.activeStory === undefined) {
            console.warn("no active story");
            return undefined;
        }
        var curr = estimate.cache.stories.filter(function (x) { return x.id === estimate.cache.activeStory; });
        if (curr.length !== 1) {
            console.warn("cannot load active story [" + estimate.cache.activeStory + "]");
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
        util.req("#story-title").innerText = story.title;
        viewStoryStatus(story.status.key);
    }
    story_1.viewActiveStory = viewActiveStory;
    function viewStoryStatus(status) {
        switch (status) {
            case "pending":
                break;
            case "active":
                viewVotes();
                break;
            case "complete":
                viewVotes();
                break;
        }
        for (var _i = 0, _a = util.els(".story-status-section"); _i < _a.length; _i++) {
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
        var x = estimate.cache.votes;
        x = x.filter(function (v) { return v.userID != vote.userID || v.storyID != vote.storyID; });
        x.push(vote);
        estimate.cache.votes = x;
        if (vote.storyID === estimate.cache.activeStory) {
            viewVotes();
        }
    }
    story_1.onVoteUpdate = onVoteUpdate;
    function viewVotes() {
        var votes = estimate.cache.activeVotes();
        var activeVote = votes.filter(function (v) { return v.userID === system.cache.profile.userID; }).pop();
        viewActiveVotes(votes, activeVote);
        viewVoteResults(votes);
    }
    story_1.viewVotes = viewVotes;
    function viewActiveVotes(votes, activeVote) {
        var m = util.req("#story-vote-members");
        m.innerHTML = "";
        m.appendChild(story_1.renderVoteMembers(system.cache.members, votes));
        var c = util.req("#story-vote-choices");
        c.innerHTML = "";
        c.appendChild(story_1.renderVoteChoices(estimate.cache.detail.choices, activeVote === null || activeVote === void 0 ? void 0 : activeVote.choice));
    }
    function viewVoteResults(votes) {
        var c = util.req("#story-vote-results");
        c.innerHTML = "";
        c.appendChild(story_1.renderVoteResults(system.cache.members, votes));
    }
    function requestStoryStatus(s) {
        var story = getActiveStory();
        if (story === undefined) {
            console.log("no active story");
            return;
        }
        var msg = {
            svc: services.estimate,
            cmd: command.client.setStoryStatus,
            param: { storyID: story.id, status: s }
        };
        socket.send(msg);
    }
    story_1.requestStoryStatus = requestStoryStatus;
    function onStoryStatusChange(u) {
        util.req("#story-" + u.storyID + " .story-status").innerText = u.status.key;
        viewStoryStatus(u.status.key);
    }
    story_1.onStoryStatusChange = onStoryStatusChange;
    // noinspection JSUnusedGlobalSymbols
    function onSubmitVote(choice) {
        var msg = {
            svc: services.estimate,
            cmd: command.client.submitVote,
            param: { storyID: estimate.cache.activeStory, choice: choice }
        };
        socket.send(msg);
    }
    story_1.onSubmitVote = onSubmitVote;
})(story || (story = {}));
var util;
(function (util) {
    function els(selector, context) {
        return UIkit.util.$$(selector, context);
    }
    util.els = els;
    function req(selector, context) {
        var res = util.els(selector, context);
        if (res.length === 0) {
            console.error("no element found for selector [" + selector + "]");
        }
        return res[0];
    }
    util.req = req;
})(util || (util = {}));
var member;
(function (member_3) {
    function renderMember(member) {
        var profile = system.cache.profile;
        if (profile === undefined) {
            return JSX("div", { "class": "uk-margin-bottom" }, "error");
        }
        else {
            return JSX("div", null,
                JSX("div", { title: "user is offline", "class": "right uk-article-meta online-indicator" }, "offline"),
                JSX("a", { "class": profile.linkColor + "-fg", href: "", onclick: "return events.openModal('member', '" + member.userID + "');" }, member.name));
        }
    }
    function renderMembers(members) {
        if (members.length === 0) {
            return JSX("div", null,
                JSX("button", { "class": "uk-button uk-button-default", onclick: "events.openModal('invite');", type: "button" }, "Invite Members"));
        }
        else {
            return JSX("ul", { "class": "uk-list uk-list-divider" }, members.map(function (m) { return JSX("li", { id: "member-" + m.userID }, renderMember(m)); }));
        }
    }
    member_3.renderMembers = renderMembers;
})(member || (member = {}));
var story;
(function (story_2) {
    function renderStory(story) {
        var profile = system.cache.profile;
        if (profile === undefined) {
            return JSX("li", null, "profile error");
        }
        else {
            return JSX("li", { id: "story-" + story.id },
                JSX("div", { "class": "right uk-article-meta story-status" }, story.status.key),
                JSX("a", { "class": profile.linkColor + "-fg", href: "", onclick: "return events.openModal('story', '" + story.id + "');" }, story.title));
        }
    }
    function renderStories(stories) {
        if (stories.length === 0) {
            return JSX("div", null,
                JSX("button", { "class": "uk-button uk-button-default", onclick: "events.openModal('add-story');", type: "button" }, "Add Story"));
        }
        else {
            return JSX("ul", { "class": "uk-list uk-list-divider" }, stories.map(function (s) { return renderStory(s); }));
        }
    }
    story_2.renderStories = renderStories;
    function renderVoteMember(member, hasVote) {
        return JSX("div", { "class": "vote-member", title: member.name + " has " + (hasVote ? "voted" : "not voted") },
            JSX("div", null,
                JSX("span", { "data-uk-icon": "icon: " + (hasVote ? "check" : "minus") + "; ratio: 1.6" })),
            member.name);
    }
    function renderVoteMembers(members, votes) {
        return JSX("div", { "class": "uk-flex uk-flex-wrap uk-flex-around" }, members.map(function (m) { return renderVoteMember(m, votes.filter(function (v) { return v.userID == m.userID; }).length > 0); }));
    }
    story_2.renderVoteMembers = renderVoteMembers;
    function renderVoteChoices(choices, choice) {
        return JSX("div", { "class": "uk-flex uk-flex-wrap uk-flex-center" }, choices.map(function (c) { return JSX("div", { "class": "vote-choice uk-border-circle uk-box-shadow-hover-medium" + (c === choice ? " active " + system.cache.profile.linkColor + "-border" : ""), onclick: "story.onSubmitVote('" + c + "');" }, c); }));
    }
    story_2.renderVoteChoices = renderVoteChoices;
    function renderVoteResult(member, choice) {
        if (choice === undefined) {
            return JSX("div", { "class": "vote-member" },
                JSX("div", null,
                    JSX("span", { "data-uk-icon": "icon: minus; ratio: 1.6" })),
                " ",
                member.name);
        }
        return JSX("div", { "class": "vote-member" },
            JSX("div", null, choice),
            " ",
            member.name);
    }
    function renderVoteResults(members, votes) {
        return JSX("div", { "class": "uk-flex uk-flex-wrap uk-flex-around" }, members.map(function (m) {
            var vote = votes.filter(function (v) { return v.userID == m.userID; });
            return renderVoteResult(m, vote.length > 0 ? vote[0].choice : undefined);
        }));
    }
    story_2.renderVoteResults = renderVoteResults;
})(story || (story = {}));
