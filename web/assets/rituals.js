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
                vote.setVotes(sj.votes);
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
                vote.onVoteUpdate(param);
                break;
            default:
                console.warn("unhandled command [" + cmd + "] for estimate");
        }
    }
    estimate.onEstimateMessage = onEstimateMessage;
    function setEstimateDetail(detail) {
        estimate.cache.detail = detail;
        util.setValue("#model-choices-input", detail.choices.join(", "));
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
            case "session":
                var sessionInput_1 = util.setValue("#model-title-input", util.req("#model-title").innerText);
                delay(function () { return sessionInput_1.focus(); });
                break;
            // member
            case "self":
                var selfInput_1 = util.setValue("#self-name-input", util.req("#member-self .member-name").innerText);
                delay(function () { return selfInput_1.focus(); });
                break;
            case "invite":
                break;
            case "member":
                system.cache.activeMember = id;
                member.viewActiveMember();
                break;
            // estimate
            case "add-story":
                var storyInput_1 = util.setValue("#story-title-input", "");
                delay(function () { return storyInput_1.focus(); });
                break;
            case "story":
                estimate.cache.activeStory = id;
                story.viewActiveStory();
                break;
            // standup
            case "add-report":
                util.setValue("#standup-report-date", dateToYMD(new Date()));
                var reportContent_1 = util.setValue("#standup-report-content", "");
                util.wireTextarea(reportContent_1);
                delay(function () { return reportContent_1.focus(); });
                break;
            case "report":
                standup.cache.activeReport = id;
                report.viewActiveReport();
                var reportEditContent_1 = util.req("#standup-report-edit-content");
                delay(function () {
                    util.wireTextarea(reportEditContent_1);
                    reportEditContent_1.focus();
                });
                break;
            // default
            default:
                console.debug("unhandled modal [" + key + "]");
        }
        UIkit.modal("#modal-" + key).show();
        return false;
    }
    events.openModal = openModal;
    function dateToYMD(date) {
        var d = date.getDate();
        var m = date.getMonth() + 1;
        var y = date.getFullYear();
        return '' + y + '-' + (m <= 9 ? '0' + m : m) + '-' + (d <= 9 ? '0' + d : d);
    }
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
        Cache.prototype.getProfile = function () {
            if (this.profile === undefined) {
                throw "no active profile";
            }
            return this.profile;
        };
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
        submitVote: "submit-vote",
        addReport: "add-report",
        editReport: "edit-report"
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
        voteUpdate: "vote-update",
        reportUpdate: "report-update"
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
            util.setText("#member-self .member-name", self[0].name);
            util.setValue("#self-name-input", self[0].name);
            util.setText("#member-self .member-role", self[0].role.key);
        }
        else if (self.length === 0) {
            console.warn("self not found among members");
        }
        else {
            console.warn("multiple self entries found among members");
        }
        var others = system.cache.members.filter(function (x) { return !isSelf(x); });
        util.setContent("#member-detail", member_1.renderMembers(others));
        renderOnline();
    }
    member_1.setMembers = setMembers;
    function onMemberUpdate(member) {
        if (isSelf(member)) {
            UIkit.modal("#modal-self").hide();
        }
        var x = system.cache.members;
        var curr = x.filter(function (m) { return m.userID === member.userID; });
        var nameChanged = curr.length == 1 && curr[0].name != member.name;
        x = x.filter(function (m) { return m.userID !== member.userID; });
        if (x.length === system.cache.members.length) {
            UIkit.notification(member.name + " has joined", { status: "success", pos: "top-right" });
        }
        x.push(member);
        x = x.sort(function (l, r) { return (l.name > r.name) ? 1 : -1; });
        system.cache.members = x;
        setMembers();
        if (nameChanged) {
            if (system.cache.currentService == services.estimate) {
                if (estimate.cache.activeStory) {
                    vote.viewVotes();
                }
            }
            if (system.cache.currentService == services.standup) {
                util.setContent("#report-detail", report.renderReports(standup.cache.reports));
                if (standup.cache.activeReport) {
                    report.viewActiveReport();
                }
            }
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
        util.setText("#member-modal-name", member.name);
        util.setText("#member-modal-role", member.role.key);
    }
    member_1.viewActiveMember = viewActiveMember;
})(member || (member = {}));
var report;
(function (report_1) {
    function onSubmitReport() {
        var d = util.req("#standup-report-date").value;
        var content = util.req("#standup-report-content").value;
        var msg = {
            svc: services.standup,
            cmd: command.client.addReport,
            param: { d: d, content: content }
        };
        socket.send(msg);
        return false;
    }
    report_1.onSubmitReport = onSubmitReport;
    function onEditReport() {
        var d = util.req("#standup-report-edit-date").value;
        var content = util.req("#standup-report-edit-content").value;
        var msg = {
            svc: services.standup,
            cmd: command.client.editReport,
            param: { id: standup.cache.activeReport, d: d, content: content }
        };
        socket.send(msg);
        return false;
    }
    report_1.onEditReport = onEditReport;
    function getActiveReport() {
        if (standup.cache.activeReport === undefined) {
            console.warn("no active report");
            return undefined;
        }
        var curr = standup.cache.reports.filter(function (x) { return x.id === standup.cache.activeReport; });
        if (curr.length !== 1) {
            console.log("cannot load active report [" + standup.cache.activeReport + "]");
            return undefined;
        }
        return curr[0];
    }
    function viewActiveReport() {
        var profile = system.cache.getProfile();
        var report = getActiveReport();
        if (report === undefined) {
            return;
        }
        util.setText("#report-title", report.d + " / " + member.getMemberName(report.author));
        var contentEdit = util.req("#modal-report .content-edit");
        var contentEditDate = util.req("#standup-report-edit-date", contentEdit);
        var contentEditTextarea = util.req("#standup-report-edit-content", contentEdit);
        var contentView = util.req("#modal-report .content-view");
        var buttonsEdit = util.req("#modal-report .buttons-edit");
        var buttonsView = util.req("#modal-report .buttons-view");
        if (report.author === profile.userID) {
            contentEdit.style.display = "block";
            util.setValue(contentEditDate, report.d);
            util.setValue(contentEditTextarea, report.content);
            util.wireTextarea(contentEditTextarea);
            contentView.style.display = "none";
            util.setHTML(contentView, "");
            buttonsEdit.style.display = "block";
            buttonsView.style.display = "none";
        }
        else {
            contentEdit.style.display = "none";
            util.setValue(contentEditDate, "");
            util.setValue(contentEditTextarea, "");
            contentView.style.display = "block";
            util.setHTML(contentView, report.html);
            buttonsEdit.style.display = "none";
            buttonsView.style.display = "block";
        }
    }
    report_1.viewActiveReport = viewActiveReport;
    function setReports(reports) {
        standup.cache.reports = reports;
        util.setContent("#report-detail", report_1.renderReports(reports));
        UIkit.modal("#modal-add-report").hide();
    }
    report_1.setReports = setReports;
    function getReportDates(reports) {
        function distinct(v, i, s) {
            return s.indexOf(v) === i;
        }
        function toCollection(d) {
            return {
                "d": d,
                "reports": reports.filter(function (r) { return r.d === d; }).sort(function (l, r) { return (l.created > r.created ? -1 : 1); })
            };
        }
        return reports.map(function (r) { return r.d; }).filter(distinct).sort().reverse().map(toCollection);
    }
    report_1.getReportDates = getReportDates;
})(report || (report = {}));
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
        util.setText("#model-title", session.title);
        util.setValue("#model-title-input", session.title);
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
            this.reports = [];
        }
        return Cache;
    }());
    standup.cache = new Cache();
    function onStandupMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.standup, param);
                break;
            case command.server.sessionJoined:
                var sj = param;
                rituals.onSessionJoin(sj);
                setStandupDetail(sj.session);
                report.setReports(sj.reports);
                break;
            case command.server.sessionUpdate:
                setStandupDetail(param);
                break;
            case command.server.reportUpdate:
                onReportUpdate(param);
                break;
            default:
                console.warn("unhandled command [" + cmd + "] for standup");
        }
    }
    standup.onStandupMessage = onStandupMessage;
    function setStandupDetail(detail) {
        standup.cache.detail = detail;
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
    function onReportUpdate(r) {
        var x = standup.cache.reports;
        x = x.filter(function (p) { return p.id !== r.id; });
        x.push(r);
        report.setReports(x);
        if (r.id === standup.cache.activeReport) {
            UIkit.modal("#modal-report").hide();
        }
    }
})(standup || (standup = {}));
var story;
(function (story_1) {
    function setStories(stories) {
        estimate.cache.stories = stories;
        util.setContent("#story-detail", story_1.renderStories(stories));
        stories.forEach(function (s) { return setStoryStatus(s.id, s.status.key, s, false); });
        showTotalIfNeeded();
        UIkit.modal("#modal-add-story").hide();
    }
    story_1.setStories = setStories;
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
            return undefined;
        }
        var curr = estimate.cache.stories.filter(function (x) { return x.id === estimate.cache.activeStory; });
        if (curr.length !== 1) {
            console.warn("cannot load active story [" + estimate.cache.activeStory + "]");
            return undefined;
        }
        return curr[0];
    }
    story_1.getActiveStory = getActiveStory;
    function viewActiveStory() {
        var s = getActiveStory();
        if (s === undefined) {
            console.log("no active story");
            return;
        }
        util.setText("#story-title", s.title);
        viewStoryStatus(s.status.key);
    }
    story_1.viewActiveStory = viewActiveStory;
    function viewStoryStatus(status) {
        function setActive(el, status) {
            var s = el.id.substr(el.id.lastIndexOf("-") + 1);
            if (s === status) {
                el.classList.add("active");
            }
            else {
                el.classList.remove("active");
            }
        }
        for (var _i = 0, _a = util.els(".story-status-body"); _i < _a.length; _i++) {
            var el = _a[_i];
            setActive(el, status);
        }
        for (var _b = 0, _c = util.els(".story-status-actions"); _b < _c.length; _b++) {
            var el = _c[_b];
            setActive(el, status);
        }
        var txt = "";
        switch (status) {
            case "pending":
                txt = "Story";
                break;
            case "active":
                txt = "Voting";
                break;
            case "complete":
                txt = "Results";
                break;
        }
        util.setText("#story-status", txt);
        vote.viewVotes();
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
    function setStoryStatus(storyID, status, currStory, calcTotal) {
        if (currStory !== null && currStory.status.key == "complete") {
            if (currStory.finalVote.length > 0) {
                status = currStory.finalVote;
            }
        }
        util.setContent("#story-" + storyID + " .story-status", story_1.renderStatus(status));
        if (calcTotal) {
            showTotalIfNeeded();
        }
    }
    function onStoryStatusChange(u) {
        var currStory = null;
        estimate.cache.stories.forEach(function (s) {
            if (s.id == u.storyID) {
                currStory = s;
                s.finalVote = u.finalVote;
                s.status = u.status;
            }
        });
        setStoryStatus(u.storyID, u.status.key, currStory, true);
        if (u.storyID === estimate.cache.activeStory) {
            viewStoryStatus(u.status.key);
        }
    }
    story_1.onStoryStatusChange = onStoryStatusChange;
    function showTotalIfNeeded() {
        var stories = estimate.cache.stories;
        var strings = stories.filter(function (s) { return s.status.key === "complete"; }).map(function (s) { return s.finalVote; }).filter(function (c) { return c.length > 0; });
        var floats = strings.map(function (c) { return parseFloat(c); }).filter(function (f) { return !isNaN(f); });
        var sum = 0;
        floats.forEach(function (f) { return sum += f; });
        var curr = util.opt("#story-total");
        var panel = util.req("#story-list");
        if (curr !== null) {
            panel.removeChild(curr);
        }
        if (sum > 0) {
            panel.appendChild(story_1.renderTotal(sum));
        }
    }
})(story || (story = {}));
var util;
(function (util) {
    function els(selector, context) {
        return UIkit.util.$$(selector, context);
    }
    util.els = els;
    function opt(selector, context) {
        var res = util.els(selector, context);
        if (res.length === 0) {
            return null;
        }
        return res[0];
    }
    util.opt = opt;
    function req(selector, context) {
        var res = util.opt(selector, context);
        if (res === null) {
            console.error("no element found for selector [" + selector + "]");
        }
        return res;
    }
    util.req = req;
    function setHTML(el, html) {
        if (typeof el === "string") {
            el = util.req(el);
        }
        el.innerHTML = html;
        return el;
    }
    util.setHTML = setHTML;
    function setContent(el, e) {
        if (typeof el === "string") {
            el = util.req(el);
        }
        el.innerHTML = "";
        el.appendChild(e);
        return el;
    }
    util.setContent = setContent;
    function setText(el, text) {
        if (typeof el === "string") {
            el = util.req(el);
        }
        el.innerText = text;
        return el;
    }
    util.setText = setText;
    function setValue(el, text) {
        if (typeof el === "string") {
            el = util.req(el);
        }
        el.value = text;
        return el;
    }
    util.setValue = setValue;
    function wireTextarea(text) {
        function resize() {
            text.style.height = 'auto';
            text.style.height = (text.scrollHeight < 64 ? 64 : (text.scrollHeight + 6)) + 'px';
        }
        function delayedResize() {
            window.setTimeout(resize, 0);
        }
        var x = text.dataset["autoresize"];
        if (x === undefined) {
            text.dataset["autoresize"] = "true";
            text.addEventListener('change', resize, false);
            text.addEventListener('cut', delayedResize, false);
            text.addEventListener('paste', delayedResize, false);
            text.addEventListener('drop', delayedResize, false);
            text.addEventListener('keydown', delayedResize, false);
            text.focus();
            text.select();
        }
        resize();
    }
    util.wireTextarea = wireTextarea;
})(util || (util = {}));
var vote;
(function (vote) {
    function setVotes(votes) {
        estimate.cache.votes = votes;
        viewVotes();
    }
    vote.setVotes = setVotes;
    function onVoteUpdate(v) {
        var x = estimate.cache.votes;
        x = x.filter(function (v) { return v.userID != v.userID || v.storyID != v.storyID; });
        x.push(v);
        estimate.cache.votes = x;
        if (v.storyID === estimate.cache.activeStory) {
            viewVotes();
        }
    }
    vote.onVoteUpdate = onVoteUpdate;
    function viewVotes() {
        var s = story.getActiveStory();
        if (s === undefined) {
            return;
        }
        var votes = estimate.cache.activeVotes();
        var activeVote = votes.filter(function (v) { return v.userID === system.cache.getProfile().userID; }).pop();
        if (s.status.key == "active") {
            viewActiveVotes(votes, activeVote);
        }
        if (s.status.key == "complete") {
            viewVoteResults(votes);
        }
    }
    vote.viewVotes = viewVotes;
    function viewActiveVotes(votes, activeVote) {
        util.setContent("#story-vote-members", vote.renderVoteMembers(system.cache.members, votes));
        util.setContent("#story-vote-choices", vote.renderVoteChoices(estimate.cache.detail.choices, activeVote === null || activeVote === void 0 ? void 0 : activeVote.choice));
    }
    function viewVoteResults(votes) {
        util.setContent("#story-vote-results", vote.renderVoteResults(system.cache.members, votes));
        util.setContent("#story-vote-summary", vote.renderVoteSummary(votes));
    }
    // noinspection JSUnusedGlobalSymbols
    function onSubmitVote(choice) {
        var msg = {
            svc: services.estimate,
            cmd: command.client.submitVote,
            param: { storyID: estimate.cache.activeStory, choice: choice }
        };
        socket.send(msg);
    }
    vote.onSubmitVote = onSubmitVote;
    function getVoteResults(votes) {
        var floats = votes.map(function (v) {
            var n = parseFloat(v.choice);
            if (isNaN(n)) {
                return -1;
            }
            return n;
        }).filter(function (x) { return x !== -1; }).sort();
        var count = floats.length;
        var min = Math.min.apply(Math, floats);
        var max = Math.max.apply(Math, floats);
        var sum = floats.reduce(function (x, y) { return x + y; }, 0);
        var mode = floats.reduce(function (current, item) {
            var val = current.numMapping[item] = (current.numMapping[item] || 0) + 1;
            if (val > current.greatestFreq) {
                current.greatestFreq = val;
                current.mode = item;
            }
            return current;
        }, { mode: null, greatestFreq: -Infinity, numMapping: {} }).mode;
        return {
            count: count,
            min: min,
            max: max,
            sum: sum,
            mean: count == 0 ? 0 : sum / count,
            median: count == 0 ? 0 : floats[Math.floor(floats.length / 2)],
            mode: count == 0 ? 0 : mode
        };
    }
    vote.getVoteResults = getVoteResults;
})(vote || (vote = {}));
var member;
(function (member_3) {
    function renderMember(member) {
        var profile = system.cache.getProfile();
        return JSX("div", { "class": "section", onclick: "events.openModal('member', '" + member.userID + "');" },
            JSX("div", { title: "user is offline", "class": "right uk-article-meta online-indicator" }, "offline"),
            JSX("div", { "class": profile.linkColor + "-fg section-link" }, member.name));
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
    function getMemberName(id) {
        var ret = system.cache.members.filter(function (m) { return m.userID === id; });
        if (ret.length === 0) {
            return id;
        }
        return ret[0].name;
    }
    member_3.getMemberName = getMemberName;
})(member || (member = {}));
var report;
(function (report) {
    function renderReport(model) {
        var profile = system.cache.getProfile();
        var ret = JSX("div", { id: "report-" + model.id, "class": "report-detail uk-border-rounded section", onclick: "events.openModal('report', '" + model.id + "');" },
            JSX("a", { "class": profile.linkColor + "-fg section-link" }, member.getMemberName(model.author)),
            JSX("div", { "class": "report-content" }, "loading..."));
        if (model.html.length > 0) {
            util.setHTML(util.req(".report-content", ret), model.html).style.display = "block";
        }
        return ret;
    }
    function renderReports(reports) {
        if (reports.length === 0) {
            return JSX("div", null,
                JSX("button", { "class": "uk-button uk-button-default", onclick: "events.openModal('add-report');", type: "button" }, "Add Report"));
        }
        else {
            var dates = report.getReportDates(reports);
            return JSX("ul", { "class": "uk-list" }, dates.map(function (day) { return JSX("li", { id: "report-date-" + day.d },
                JSX("div", null, day.d),
                day.reports.map(function (r) { return JSX("li", null, renderReport(r)); })); }));
        }
    }
    report.renderReports = renderReports;
})(report || (report = {}));
var story;
(function (story_2) {
    function renderStory(story) {
        var profile = system.cache.getProfile();
        return JSX("li", { id: "story-" + story.id, "class": "section", onclick: "events.openModal('story', '" + story.id + "');" },
            JSX("div", { "class": "right uk-article-meta story-status" }, story.status.key),
            JSX("div", { "class": profile.linkColor + "-fg section-link" }, story.title));
    }
    function renderStories(stories) {
        if (stories.length === 0) {
            return JSX("div", null,
                JSX("button", { "class": "uk-button uk-button-default", onclick: "events.openModal('add-story');", type: "button" }, "Add Story"));
        }
        else {
            return JSX("ul", { id: "story-list", "class": "uk-list uk-list-divider" }, stories.map(function (s) { return renderStory(s); }));
        }
    }
    story_2.renderStories = renderStories;
    function renderStatus(status) {
        switch (status) {
            case "pending":
                return JSX("span", null, status);
            case "active":
                return JSX("span", null, status);
            default:
                return JSX("span", { "class": "vote-badge uk-border-rounded" }, status);
        }
    }
    story_2.renderStatus = renderStatus;
    function renderTotal(sum) {
        return JSX("li", { id: "story-total" },
            JSX("div", { "class": "right uk-article-meta" },
                JSX("span", { "class": "vote-badge uk-border-rounded" }, sum)),
            " Total");
    }
    story_2.renderTotal = renderTotal;
})(story || (story = {}));
var vote;
(function (vote_1) {
    function renderVoteMember(member, hasVote) {
        return JSX("div", { "class": "vote-member", title: member.name + " has " + (hasVote ? "voted" : "not voted") },
            JSX("div", null,
                JSX("span", { "data-uk-icon": "icon: " + (hasVote ? "check" : "minus") + "; ratio: 1.6" })),
            member.name);
    }
    function renderVoteMembers(members, votes) {
        return JSX("div", { "class": "uk-flex uk-flex-wrap uk-flex-around" }, members.map(function (m) { return renderVoteMember(m, votes.filter(function (v) { return v.userID == m.userID; }).length > 0); }));
    }
    vote_1.renderVoteMembers = renderVoteMembers;
    function renderVoteChoices(choices, choice) {
        return JSX("div", { "class": "uk-flex uk-flex-wrap uk-flex-center" }, choices.map(function (c) { return JSX("div", { "class": "vote-choice uk-border-circle uk-box-shadow-hover-medium" + (c === choice ? " active " + system.cache.getProfile().linkColor + "-border" : ""), onclick: "vote.onSubmitVote('" + c + "');" }, c); }));
    }
    vote_1.renderVoteChoices = renderVoteChoices;
    function renderVoteResult(member, choice) {
        if (choice === undefined) {
            return JSX("div", { "class": "vote-result" },
                JSX("div", null,
                    JSX("span", { "class": "uk-border-circle" },
                        JSX("span", { "data-uk-icon": "icon: minus; ratio: 1.6" }))),
                " ",
                member.name);
        }
        return JSX("div", { "class": "vote-result" },
            JSX("div", null,
                JSX("span", { "class": "uk-border-circle" }, choice)),
            " ",
            member.name);
    }
    function renderVoteResults(members, votes) {
        return JSX("div", { "class": "uk-flex uk-flex-wrap uk-flex-around" }, members.map(function (m) {
            var vote = votes.filter(function (v) { return v.userID == m.userID; });
            return renderVoteResult(m, length > 0 ? vote[0].choice : undefined);
        }));
    }
    vote_1.renderVoteResults = renderVoteResults;
    function renderVoteSummary(votes) {
        var results = vote_1.getVoteResults(votes);
        function trim(n) { return n.toString().substr(0, 4); }
        return JSX("div", { "class": "uk-flex uk-flex-wrap uk-flex-center result-container" },
            JSX("div", { "class": "result" },
                JSX("div", { "class": "secondary uk-border-circle" },
                    trim(results.count),
                    " / ",
                    trim(votes.length)),
                " ",
                JSX("div", null, "votes counted")),
            JSX("div", { "class": "result" },
                JSX("div", { "class": "secondary uk-border-circle" },
                    trim(results.min),
                    "-",
                    trim(results.max)),
                " ",
                JSX("div", null, "vote range")),
            JSX("div", { "class": "result mean-result" },
                JSX("div", { "class": "mean uk-border-circle " + system.cache.getProfile().linkColor + "-border" }, trim(results.mean)),
                " ",
                JSX("div", null, "average")),
            JSX("div", { "class": "result" },
                JSX("div", { "class": "secondary uk-border-circle" }, trim(results.median)),
                " ",
                JSX("div", null, "median")),
            JSX("div", { "class": "result" },
                JSX("div", { "class": "secondary uk-border-circle" }, trim(results.mode)),
                " ",
                JSX("div", null, "mode")));
    }
    vote_1.renderVoteSummary = renderVoteSummary;
})(vote || (vote = {}));
