"use strict";
var action;
(function (action) {
    function loadActions() {
        var msg = {
            svc: services.system.key,
            cmd: command.client.getActions,
            param: null
        };
        socket.send(msg);
    }
    action.loadActions = loadActions;
    function viewActions(actions) {
        util.setContent("#action-list", action.renderActions(actions));
    }
    action.viewActions = viewActions;
})(action || (action = {}));
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
                rituals.onError(services.estimate.key, param);
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
            case command.server.storyRemove:
                onStoryRemove(param);
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
        story.viewActiveStory();
        rituals.setDetail(detail);
    }
    function onSubmitEstimateSession() {
        var title = util.req("#model-title-input").value;
        var choices = util.req("#model-choices-input").value;
        var msg = {
            svc: services.estimate.key,
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
        var x = preUpdate(s.id);
        x.push(s);
        if (s.id === estimate.cache.activeStory) {
            util.setText("#story-title", s.title);
        }
        story.setStories(x);
    }
    estimate.onStoryUpdate = onStoryUpdate;
    function onStoryRemove(id) {
        var x = preUpdate(id);
        story.setStories(x);
        if (id === estimate.cache.activeStory) {
            UIkit.modal("#modal-story").hide();
        }
        UIkit.notification("story has been deleted", { status: "success", pos: "top-right" });
    }
    estimate.onStoryRemove = onStoryRemove;
    function preUpdate(id) {
        return estimate.cache.stories.filter(function (p) { return p.id !== id; });
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
            // actions
            case "actions":
                action.loadActions();
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
                util.setValue("#standup-report-date", util.dateToYMD(new Date()));
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
            // retro
            case "add-feedback":
                util.setSelectOption(util.req("#retro-feedback-category"), id);
                var feedbackContent_1 = util.setValue("#retro-feedback-content", "");
                util.wireTextarea(feedbackContent_1);
                delay(function () { return feedbackContent_1.focus(); });
                break;
            case "feedback":
                retro.cache.activeFeedback = id;
                feedback.viewActiveFeedback();
                var feedbackEditContent_1 = util.req("#retro-feedback-edit-content");
                delay(function () {
                    util.wireTextarea(feedbackEditContent_1);
                    feedbackEditContent_1.focus();
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
})(events || (events = {}));
var feedback;
(function (feedback_1) {
    function setFeedback(feedback) {
        retro.cache.feedback = feedback;
        util.setContent("#feedback-detail", feedback_1.renderFeedbackArray(feedback));
        UIkit.modal("#modal-add-feedback").hide();
    }
    feedback_1.setFeedback = setFeedback;
    function onSubmitFeedback() {
        var category = util.req("#retro-feedback-category").value;
        var content = util.req("#retro-feedback-content").value;
        var msg = {
            svc: services.retro.key,
            cmd: command.client.addFeedback,
            param: { category: category, content: content }
        };
        socket.send(msg);
        return false;
    }
    feedback_1.onSubmitFeedback = onSubmitFeedback;
    function onEditFeedback() {
        var id = retro.cache.activeFeedback;
        var category = util.req("#retro-feedback-edit-category").value;
        var content = util.req("#retro-feedback-edit-content").value;
        var msg = {
            svc: services.retro.key,
            cmd: command.client.updateFeedback,
            param: { id: id, category: category, content: content }
        };
        socket.send(msg);
        return false;
    }
    feedback_1.onEditFeedback = onEditFeedback;
    function onRemoveFeedback() {
        var id = retro.cache.activeFeedback;
        if (id && confirm("Delete this feedback?")) {
            var msg = {
                svc: services.retro.key,
                cmd: command.client.removeFeedback,
                param: id
            };
            socket.send(msg);
            UIkit.modal("#modal-feedback").hide();
        }
        return false;
    }
    feedback_1.onRemoveFeedback = onRemoveFeedback;
    function getActiveFeedback() {
        if (retro.cache.activeFeedback === undefined) {
            return undefined;
        }
        var curr = retro.cache.feedback.filter(function (x) { return x.id === retro.cache.activeFeedback; });
        if (curr.length !== 1) {
            console.warn("cannot load active Feedback [" + retro.cache.activeFeedback + "]");
            return undefined;
        }
        return curr[0];
    }
    feedback_1.getActiveFeedback = getActiveFeedback;
    function viewActiveFeedback() {
        var profile = system.cache.getProfile();
        var fb = getActiveFeedback();
        if (fb === undefined) {
            console.warn("no active feedback");
            return;
        }
        util.setText("#feedback-title", fb.category + " / " + system.getMemberName(fb.authorID));
        var contentEdit = util.req("#modal-feedback .content-edit");
        var contentEditCategory = util.req("#retro-feedback-edit-category", contentEdit);
        var contentEditTextarea = util.req("#retro-feedback-edit-content", contentEdit);
        var contentView = util.req("#modal-feedback .content-view");
        var buttonsEdit = util.req("#modal-feedback .buttons-edit");
        var buttonsView = util.req("#modal-feedback .buttons-view");
        if (fb.authorID === profile.userID) {
            contentEdit.style.display = "block";
            util.setSelectOption(contentEditCategory, fb.category);
            util.setValue(contentEditTextarea, fb.content);
            util.wireTextarea(contentEditTextarea);
            contentView.style.display = "none";
            util.setHTML(contentView, "");
            buttonsEdit.style.display = "block";
            buttonsView.style.display = "none";
        }
        else {
            contentEdit.style.display = "none";
            util.setSelectOption(contentEditCategory, undefined);
            util.setValue(contentEditTextarea, "");
            contentView.style.display = "block";
            util.setHTML(contentView, fb.html);
            buttonsEdit.style.display = "none";
            buttonsView.style.display = "block";
        }
    }
    feedback_1.viewActiveFeedback = viewActiveFeedback;
    function onFeedbackUpdate(r) {
        var x = preUpdate(r.id);
        x.push(r);
        postUpdate(x, r.id);
    }
    feedback_1.onFeedbackUpdate = onFeedbackUpdate;
    function onFeedbackRemoved(id) {
        var x = preUpdate(id);
        postUpdate(x, id);
        UIkit.notification("feedback has been deleted", { status: "success", pos: "top-right" });
    }
    feedback_1.onFeedbackRemoved = onFeedbackRemoved;
    function preUpdate(id) {
        return retro.cache.feedback.filter(function (p) { return p.id !== id; });
    }
    function postUpdate(x, id) {
        feedback.setFeedback(x);
        if (id === retro.cache.activeFeedback) {
            UIkit.modal("#modal-feedback").hide();
        }
    }
    function getFeedbackCategories(feedback, categories) {
        function toCollection(c) {
            var reports = feedback.filter(function (r) { return r.category === c; }).sort(function (l, r) { return (l.created > r.created ? -1 : 1); });
            return { category: c, feedback: reports };
        }
        var ret = categories.map(toCollection);
        var extras = feedback.filter(function (r) { return categories.indexOf(r.category) == -1; });
        if (extras.length > 0) {
            ret.push({ category: "unknown", feedback: extras });
        }
        return ret;
    }
    feedback_1.getFeedbackCategories = getFeedbackCategories;
})(feedback || (feedback = {}));
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
    function getMemberName(id) {
        var ret = system.cache.members.filter(function (m) { return m.userID === id; });
        if (ret.length === 0) {
            return id;
        }
        return ret[0].name;
    }
    system.getMemberName = getMemberName;
    system.cache = new Cache();
})(system || (system = {}));
var services;
(function (services) {
    services.system = {
        key: "system",
        title: "System",
        plural: "systems",
        icon: "close"
    };
    services.sprint = {
        key: "sprint",
        title: "Sprint",
        plural: "sprints",
        icon: "git-fork"
    };
    services.estimate = {
        key: "estimate",
        title: "Estimate",
        plural: "estimates",
        icon: "settings"
    };
    services.standup = {
        key: "standup",
        title: "Standup",
        plural: "standups",
        icon: "future"
    };
    services.retro = {
        key: "retro",
        title: "Retrospective",
        plural: "retros",
        icon: "history"
    };
})(services || (services = {}));
var command;
(function (command) {
    command.client = {
        error: "error",
        ping: "ping",
        connect: "connect",
        getActions: "get-actions",
        updateProfile: "update-profile",
        updateSession: "update-session",
        addStory: "add-story",
        updateStory: "update-story",
        removeStory: "remove-story",
        setStoryStatus: "set-story-status",
        submitVote: "submit-vote",
        addReport: "add-report",
        updateReport: "update-report",
        removeReport: "remove-report",
        addFeedback: "add-feedback",
        updateFeedback: "update-feedback",
        removeFeedback: "remove-feedback"
    };
    command.server = {
        error: "error",
        pong: "pong",
        actions: "actions",
        sessionJoined: "session-joined",
        sessionUpdate: "session-update",
        memberUpdate: "member-update",
        onlineUpdate: "online-update",
        storyUpdate: "story-update",
        storyRemove: "story-remove",
        storyStatusChange: "story-status-change",
        voteUpdate: "vote-update",
        reportUpdate: "report-update",
        reportRemove: "report-remove",
        feedbackUpdate: "feedback-update",
        feedbackRemove: "feedback-remove"
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
            util.setText("#member-self .member-role", self[0].role);
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
            switch (system.cache.currentService) {
                case services.sprint.key:
                    break;
                case services.estimate.key:
                    if (estimate.cache.activeStory) {
                        vote.viewVotes();
                    }
                    break;
                case services.standup.key:
                    util.setContent("#report-detail", report.renderReports(standup.cache.reports));
                    if (standup.cache.activeReport) {
                        report.viewActiveReport();
                    }
                    break;
                case services.retro.key:
                    util.setContent("#feedback-detail", feedback.renderFeedbackArray(retro.cache.feedback));
                    if (retro.cache.activeFeedback) {
                        feedback.viewActiveFeedback();
                    }
                    break;
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
            svc: services.system.key,
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
            console.warn("cannot load active member [" + system.cache.activeMember + "]");
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
        util.setText("#member-modal-role", member.role);
    }
    member_1.viewActiveMember = viewActiveMember;
})(member || (member = {}));
var profile;
(function (profile) {
    function setNavColor(el, c) {
        util.setValue("#navbar-color", c);
        var nb = util.req("#navbar");
        nb.className = (c + "-bg uk-navbar-container uk-navbar");
        var colors = document.querySelectorAll(".navbar_swatch");
        colors.forEach(function (i) {
            i.classList.remove("active");
        });
        el.classList.add("active");
    }
    profile.setNavColor = setNavColor;
    function setLinkColor(el, c) {
        util.setValue("#link-color", c);
        var links = util.els(".profile-link");
        links.forEach(function (l) {
            l.classList.forEach(function (x) {
                if (x.indexOf("-fg") > -1) {
                    l.classList.remove(x);
                }
                l.classList.add(c + "-fg");
            });
        });
        var colors = document.querySelectorAll(".link_swatch");
        colors.forEach(function (i) {
            i.classList.remove("active");
        });
        el.classList.add("active");
    }
    profile.setLinkColor = setLinkColor;
    function selectTheme(theme) {
        var card = util.els(".uk-card");
        switch (theme) {
            case "light":
                document.documentElement.classList.remove("uk-light");
                document.body.classList.remove("uk-light");
                document.documentElement.classList.add("uk-dark");
                document.body.classList.add("uk-dark");
                card.forEach(function (x) {
                    x.classList.add("uk-card-default");
                    x.classList.remove("uk-card-secondary");
                });
                break;
            case "dark":
                document.documentElement.classList.add("uk-light");
                document.body.classList.add("uk-light");
                document.documentElement.classList.remove("uk-dark");
                document.body.classList.remove("uk-dark");
                card.forEach(function (x) {
                    x.classList.remove("uk-card-default");
                    x.classList.add("uk-card-secondary");
                });
                break;
            default:
                console.warn("invalid theme");
                break;
        }
    }
    profile.selectTheme = selectTheme;
})(profile || (profile = {}));
var report;
(function (report_1) {
    function onSubmitReport() {
        var d = util.req("#standup-report-date").value;
        var content = util.req("#standup-report-content").value;
        var msg = {
            svc: services.standup.key,
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
            svc: services.standup.key,
            cmd: command.client.updateReport,
            param: { id: standup.cache.activeReport, d: d, content: content }
        };
        socket.send(msg);
        return false;
    }
    report_1.onEditReport = onEditReport;
    function onRemoveReport() {
        var id = standup.cache.activeReport;
        if (id && confirm("Delete this report?")) {
            var msg = {
                svc: services.standup.key,
                cmd: command.client.removeReport,
                param: id
            };
            socket.send(msg);
            UIkit.modal("#modal-report").hide();
        }
        return false;
    }
    report_1.onRemoveReport = onRemoveReport;
    function getActiveReport() {
        if (standup.cache.activeReport === undefined) {
            console.warn("no active report");
            return undefined;
        }
        var curr = standup.cache.reports.filter(function (x) { return x.id === standup.cache.activeReport; });
        if (curr.length !== 1) {
            console.warn("cannot load active report [" + standup.cache.activeReport + "]");
            return undefined;
        }
        return curr[0];
    }
    function viewActiveReport() {
        var profile = system.cache.getProfile();
        var report = getActiveReport();
        if (report === undefined) {
            console.warn("no active report");
            return;
        }
        util.setText("#report-title", report.d + " / " + system.getMemberName(report.authorID));
        var contentEdit = util.req("#modal-report .content-edit");
        var contentEditDate = util.req("#standup-report-edit-date", contentEdit);
        var contentEditTextarea = util.req("#standup-report-edit-content", contentEdit);
        var contentView = util.req("#modal-report .content-view");
        var buttonsEdit = util.req("#modal-report .buttons-edit");
        var buttonsView = util.req("#modal-report .buttons-view");
        if (report.authorID === profile.userID) {
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
            this.feedback = [];
        }
        return Cache;
    }());
    retro.cache = new Cache();
    function onRetroMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.retro.key, param);
                break;
            case command.server.sessionJoined:
                var sj = param;
                rituals.onSessionJoin(sj);
                setRetroDetail(sj.session);
                feedback.setFeedback(sj.feedback);
                break;
            case command.server.sessionUpdate:
                setRetroDetail(param);
                break;
            case command.server.feedbackUpdate:
                feedback.onFeedbackUpdate(param);
                break;
            case command.server.feedbackRemove:
                feedback.onFeedbackRemoved(param);
                break;
            default:
                console.warn("unhandled command [" + cmd + "] for retro");
        }
    }
    retro.onRetroMessage = onRetroMessage;
    function setRetroDetail(detail) {
        retro.cache.detail = detail;
        util.setValue("#model-categories-input", detail.categories.join(", "));
        util.setOptions(util.req("#retro-feedback-category"), detail.categories);
        util.setOptions(util.req("#retro-feedback-edit-category"), detail.categories);
        feedback.setFeedback(retro.cache.feedback);
        rituals.setDetail(detail);
    }
    function onSubmitRetroSession() {
        var title = util.req("#model-title-input").value;
        var categories = util.req("#model-categories-input").value;
        var msg = {
            svc: services.retro.key,
            cmd: command.client.updateSession,
            param: {
                title: title,
                categories: categories
            }
        };
        socket.send(msg);
    }
    retro.onSubmitRetroSession = onSubmitRetroSession;
})(retro || (retro = {}));
var rituals;
(function (rituals) {
    function onSocketMessage(msg) {
        console.debug("message received");
        console.debug(msg);
        switch (msg.svc) {
            case services.system.key:
                onSystemMessage(msg.cmd, msg.param);
                break;
            case services.sprint.key:
                sprint.onSprintMessage(msg.cmd, msg.param);
                break;
            case services.estimate.key:
                estimate.onEstimateMessage(msg.cmd, msg.param);
                break;
            case services.standup.key:
                standup.onStandupMessage(msg.cmd, msg.param);
                break;
            case services.retro.key:
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
            case command.server.actions:
                action.viewActions(param);
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
        console.debug("sending message");
        console.debug(msg);
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
var sprint;
(function (sprint) {
    var Cache = /** @class */ (function () {
        function Cache() {
            this.estimates = [];
            this.standups = [];
            this.retros = [];
        }
        return Cache;
    }());
    sprint.cache = new Cache();
    function onSprintMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.sprint.key, param);
                break;
            case command.server.sessionJoined:
                var sj = param;
                rituals.onSessionJoin(sj);
                setSprintDetail(sj.session);
                setSprintContents(sj);
                break;
            case command.server.sessionUpdate:
                setSprintDetail(param);
                break;
            default:
                console.warn("unhandled command [" + cmd + "] for sprint");
        }
    }
    sprint.onSprintMessage = onSprintMessage;
    function setSprintDetail(detail) {
        sprint.cache.detail = detail;
        rituals.setDetail(detail);
    }
    function setSprintContents(sj) {
        viewEstimates(sj.estimates);
        viewStandups(sj.standups);
        viewRetros(sj.retros);
    }
    function viewEstimates(estimates) {
        sprint.cache.estimates = estimates;
        util.setContent("#sprint-estimate-list", sprint.renderContents(services.estimate, sprint.cache.estimates));
    }
    function viewStandups(standups) {
        sprint.cache.standups = standups;
        util.setContent("#sprint-standup-list", sprint.renderContents(services.standup, sprint.cache.standups));
    }
    function viewRetros(retros) {
        sprint.cache.retros = retros;
        util.setContent("#sprint-retro-list", sprint.renderContents(services.retro, sprint.cache.retros));
    }
    function onSubmitSprintSession() {
        var title = util.req("#model-title-input").value;
        var msg = {
            svc: services.sprint.key,
            cmd: command.client.updateSession,
            param: {
                title: title
            }
        };
        socket.send(msg);
    }
    sprint.onSubmitSprintSession = onSubmitSprintSession;
})(sprint || (sprint = {}));
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
                rituals.onError(services.standup.key, param);
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
            case command.server.reportRemove:
                onReportRemoved(param);
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
            svc: services.standup.key,
            cmd: command.client.updateSession,
            param: {
                title: title
            }
        };
        socket.send(msg);
    }
    standup.onSubmitStandupSession = onSubmitStandupSession;
    function onReportUpdate(r) {
        var x = preUpdate(r.id);
        x.push(r);
        postUpdate(x, r.id);
    }
    function onReportRemoved(id) {
        var x = preUpdate(id);
        postUpdate(x, id);
        UIkit.notification("report has been deleted", { status: "success", pos: "top-right" });
    }
    function preUpdate(id) {
        return standup.cache.reports.filter(function (p) { return p.id !== id; });
    }
    function postUpdate(x, id) {
        report.setReports(x);
        if (id === standup.cache.activeReport) {
            UIkit.modal("#modal-report").hide();
        }
    }
})(standup || (standup = {}));
var story;
(function (story_1) {
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
    story_1.viewStoryStatus = viewStoryStatus;
    function requestStoryStatus(s) {
        var story = story_1.getActiveStory();
        if (story === undefined) {
            return;
        }
        var msg = {
            svc: services.estimate.key,
            cmd: command.client.setStoryStatus,
            param: { storyID: story.id, status: s }
        };
        socket.send(msg);
    }
    story_1.requestStoryStatus = requestStoryStatus;
    function setStoryStatus(storyID, status, currStory, calcTotal) {
        if (currStory !== null && currStory.status == "complete") {
            if (currStory.finalVote.length > 0) {
                status = currStory.finalVote;
            }
        }
        util.setContent("#story-" + storyID + " .story-status", story_1.renderStatus(status));
        if (calcTotal) {
            story_1.showTotalIfNeeded();
        }
    }
    story_1.setStoryStatus = setStoryStatus;
    function onStoryStatusChange(u) {
        var currStory = null;
        estimate.cache.stories.forEach(function (s) {
            if (s.id == u.storyID) {
                currStory = s;
                s.finalVote = u.finalVote;
                s.status = u.status;
            }
        });
        setStoryStatus(u.storyID, u.status, currStory, true);
        if (u.storyID === estimate.cache.activeStory) {
            viewStoryStatus(u.status);
        }
    }
    story_1.onStoryStatusChange = onStoryStatusChange;
})(story || (story = {}));
var story;
(function (story) {
    function setStories(stories) {
        estimate.cache.stories = stories;
        util.setContent("#story-detail", story.renderStories(stories));
        stories.forEach(function (s) { return story.setStoryStatus(s.id, s.status, s, false); });
        showTotalIfNeeded();
        UIkit.modal("#modal-add-story").hide();
    }
    story.setStories = setStories;
    function onSubmitStory() {
        var title = util.req("#story-title-input").value;
        var msg = {
            svc: services.estimate.key,
            cmd: command.client.addStory,
            param: { title: title }
        };
        socket.send(msg);
        return false;
    }
    story.onSubmitStory = onSubmitStory;
    function beginEditStory() {
        var s = getActiveStory();
        var x = prompt("Edit your story", s.title);
        if (x !== null && x !== s.title) {
            var msg = {
                svc: services.estimate.key,
                cmd: command.client.updateStory,
                param: { id: s.id, title: x }
            };
            socket.send(msg);
        }
        return false;
    }
    story.beginEditStory = beginEditStory;
    function onRemoveStory() {
        var id = estimate.cache.activeStory;
        if (id && confirm("Delete this story?")) {
            var msg = {
                svc: services.estimate.key,
                cmd: command.client.removeStory,
                param: id
            };
            socket.send(msg);
            UIkit.modal("#modal-story").hide();
        }
        return false;
    }
    story.onRemoveStory = onRemoveStory;
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
    story.getActiveStory = getActiveStory;
    function viewActiveStory() {
        var s = getActiveStory();
        if (s === undefined) {
            return;
        }
        util.setText("#story-title", s.title);
        story.viewStoryStatus(s.status);
    }
    story.viewActiveStory = viewActiveStory;
    function showTotalIfNeeded() {
        var stories = estimate.cache.stories;
        var strings = stories.filter(function (s) { return s.status === "complete"; }).map(function (s) { return s.finalVote; }).filter(function (c) { return c.length > 0; });
        var floats = strings.map(function (c) { return parseFloat(c); }).filter(function (f) { return !isNaN(f); });
        var sum = 0;
        floats.forEach(function (f) { return sum += f; });
        var curr = util.opt("#story-total");
        var panel = util.req("#story-list");
        if (curr !== null) {
            panel.removeChild(curr);
        }
        if (sum > 0) {
            panel.appendChild(story.renderTotal(sum));
        }
    }
    story.showTotalIfNeeded = showTotalIfNeeded;
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
            console.warn("no element found for selector [" + selector + "]");
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
    function dateToYMD(date) {
        var d = date.getDate();
        var m = date.getMonth() + 1;
        var y = date.getFullYear();
        return '' + y + '-' + (m <= 9 ? '0' + m : m) + '-' + (d <= 9 ? '0' + d : d);
    }
    util.dateToYMD = dateToYMD;
    function dateFromYMD(s) {
        var d = new Date(s);
        d = new Date(d.getTime() + (d.getTimezoneOffset() * 60000));
        return d;
    }
    util.dateFromYMD = dateFromYMD;
    function dow(i) {
        switch (i) {
            case 0:
                return "Sun";
            case 1:
                return "Mon";
            case 2:
                return "Tue";
            case 3:
                return "Wed";
            case 4:
                return "Thu";
            case 5:
                return "Fri";
            case 6:
                return "Sat";
            default:
                return "???";
        }
    }
    util.dow = dow;
    function setOptions(el, categories) {
        el.innerHTML = "";
        for (var _i = 0, categories_1 = categories; _i < categories_1.length; _i++) {
            var c = categories_1[_i];
            var opt_1 = document.createElement("option");
            opt_1.value = c;
            opt_1.innerText = c;
            el.appendChild(opt_1);
        }
    }
    util.setOptions = setOptions;
    function setSelectOption(el, o) {
        for (var i = 0; i < el.children.length; i++) {
            var e = el.children.item(i);
            e.selected = e.value === o;
        }
    }
    util.setSelectOption = setSelectOption;
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
        switch (s.status) {
            case "pending":
                var uID = system.cache.getProfile().userID;
                var e = util.req("#story-edit-section");
                var v = util.req("#story-view-section");
                if (uID === s.authorID) {
                    e.style.display = "block";
                    v.style.display = "none";
                }
                else {
                    e.style.display = "none";
                    v.style.display = "block";
                }
                break;
            case "active":
                viewActiveVotes(votes, activeVote);
                break;
            case "complete":
                viewVoteResults(votes);
                break;
            default:
                console.warn("invalid story status [" + s.status + "]");
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
            svc: services.estimate.key,
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
var action;
(function (action_1) {
    function renderAction(action) {
        var c = JSON.stringify(action.content, null, 2);
        return JSX("tr", null,
            JSX("td", null, system.getMemberName(action.authorID)),
            JSX("td", null, action.act),
            JSX("td", null, c === "null" ? "" : JSX("pre", null, c)),
            JSX("td", null, action.note),
            JSX("td", null,
                new Date(action.occurred).toLocaleDateString(),
                " ",
                new Date(action.occurred).toLocaleTimeString().slice(0, 8)));
    }
    function renderActions(actions) {
        if (actions.length === 0) {
            return JSX("div", null, "No actions available");
        }
        else {
            return JSX("table", { "class": "uk-table uk-table-divider uk-text-left" },
                JSX("thead", null,
                    JSX("tr", null,
                        JSX("th", null, "Author"),
                        JSX("th", null, "Act"),
                        JSX("th", null, "Content"),
                        JSX("th", null, "Note"),
                        JSX("th", null, "Occurred"))),
                JSX("tbody", null, actions.map(function (a) { return renderAction(a); })));
        }
    }
    action_1.renderActions = renderActions;
})(action || (action = {}));
var feedback;
(function (feedback) {
    function renderFeedback(model) {
        var profile = system.cache.getProfile();
        var ret = JSX("div", { id: "feedback-" + model.id, "class": "feedback-detail uk-border-rounded section", onclick: "events.openModal('feedback', '" + model.id + "');" },
            JSX("a", { "class": profile.linkColor + "-fg section-link" }, system.getMemberName(model.authorID)),
            JSX("div", { "class": "feedback-content" }, "loading..."));
        if (model.html.length > 0) {
            util.setHTML(util.req(".feedback-content", ret), model.html).style.display = "block";
        }
        return ret;
    }
    function renderFeedbackArray(f) {
        var _a;
        if (f.length === 0) {
            return JSX("div", null,
                JSX("button", { "class": "uk-button uk-button-default", onclick: "events.openModal('add-feedback');", type: "button" }, "Add Feedback"));
        }
        else {
            var cats = feedback.getFeedbackCategories(f, ((_a = retro.cache.detail) === null || _a === void 0 ? void 0 : _a.categories) || []);
            var profile_1 = system.cache.getProfile();
            return JSX("div", { "class": "uk-grid-small uk-grid-match uk-child-width-expand@m uk-grid-divider", "uk-grid": "" }, cats.map(function (cat) { return JSX("div", { "class": "feedback-list uk-transition-toggle" },
                JSX("div", { "class": "feedback-category-header" },
                    JSX("span", { "class": "right" },
                        JSX("a", { "class": profile_1.linkColor + "-fg uk-icon-button uk-transition-fade", "data-uk-icon": "plus", onclick: "events.openModal('add-feedback', '" + cat.category + "');", title: "edit session" })),
                    JSX("span", { "class": "feedback-category-title", onclick: "events.openModal('add-feedback', '" + cat.category + "');" }, cat.category)),
                JSX("div", null, cat.feedback.map(function (fb) { return JSX("div", null, renderFeedback(fb)); }))); }));
        }
    }
    feedback.renderFeedbackArray = renderFeedbackArray;
})(feedback || (feedback = {}));
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
})(member || (member = {}));
var report;
(function (report) {
    function renderReport(model) {
        var profile = system.cache.getProfile();
        var ret = JSX("div", { id: "report-" + model.id, "class": "report-detail uk-border-rounded section", onclick: "events.openModal('report', '" + model.id + "');" },
            JSX("a", { "class": profile.linkColor + "-fg section-link" }, system.getMemberName(model.authorID)),
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
                JSX("h5", null,
                    JSX("div", { "class": "right uk-article-meta" }, util.dow(util.dateFromYMD(day.d).getDay())),
                    util.dateFromYMD(day.d).toLocaleDateString()),
                JSX("ul", null, day.reports.map(function (r) { return JSX("li", null, renderReport(r)); }))); }));
        }
    }
    report.renderReports = renderReports;
})(report || (report = {}));
var sprint;
(function (sprint) {
    function renderSprintContent(svc, session) {
        var profile = system.cache.getProfile();
        return JSX("tr", null,
            JSX("td", null,
                JSX("a", { "class": profile.linkColor + "-fg", href: "/" + svc.key + "/" + session.slug }, session.slug)),
            JSX("td", { "class": "uk-table-shrink uk-text-nowrap" }, system.getMemberName(session.owner)),
            JSX("td", { "class": "uk-table-shrink uk-text-nowrap" },
                new Date(session.created).toLocaleDateString(),
                " ",
                new Date(session.created).toLocaleTimeString().slice(0, 8)));
    }
    function toContent(svc, sessions) {
        return sessions.map(function (s) { return { svc: svc, session: s }; });
    }
    function renderContents(svc, sessions) {
        var contents = toContent(svc, sessions);
        contents.sort(function (l, r) { return (l.session.created > r.session.created ? -1 : 1); });
        if (contents.length === 0) {
            return JSX("div", null, "No " + svc.plural + " in this sprint");
        }
        else {
            return JSX("table", { "class": "uk-table uk-table-divider uk-text-left" },
                JSX("tbody", null, contents.map(function (a) { return renderSprintContent(a.svc, a.session); })));
        }
    }
    sprint.renderContents = renderContents;
})(sprint || (sprint = {}));
var story;
(function (story_2) {
    function renderStory(story) {
        var profile = system.cache.getProfile();
        return JSX("li", { id: "story-" + story.id, "class": "section", onclick: "events.openModal('story', '" + story.id + "');" },
            JSX("div", { "class": "right uk-article-meta story-status" }, story.status),
            JSX("div", { "class": profile.linkColor + "-fg section-link" }, story.title));
    }
    function renderStories(stories) {
        if (stories.length === 0) {
            return JSX("div", { id: "story-list" },
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
