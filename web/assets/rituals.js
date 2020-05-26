"use strict";
var action;
(function (action) {
    function loadActions() {
        const msg = { svc: services.system.key, cmd: command.client.getActions, param: null };
        socket.send(msg);
    }
    action.loadActions = loadActions;
    function viewActions(actions) {
        dom.setContent("#action-list", action.renderActions(actions));
    }
    action.viewActions = viewActions;
})(action || (action = {}));
var auth;
(function (auth) {
    const github = { key: "github", title: "GitHub" };
    const google = { key: "google", title: "Google" };
    const slack = { key: "slack", title: "Slack" };
    const amazon = { key: "amazon", title: "Amazon" };
    const microsoft = { key: "microsoft", title: "Microsoft" };
    auth.allProviders = [github, google, slack, amazon, microsoft];
})(auth || (auth = {}));
var collection;
(function (collection) {
    class Group {
        constructor(key) {
            this.members = [];
            this.key = key;
        }
    }
    collection.Group = Group;
    function groupBy(list, func) {
        const res = [];
        let group;
        if (list) {
            list.forEach((o) => {
                const groupName = func(o);
                if (!group) {
                    group = new Group(groupName);
                }
                if (groupName != group.key) {
                    res.push(group);
                    group = new Group(groupName);
                }
                group.members.push(o);
            });
        }
        if (group) {
            res.push(group);
        }
        return res;
    }
    collection.groupBy = groupBy;
    function findGroup(groups, key) {
        for (const g of groups) {
            if (g.key === key) {
                return g.members;
            }
        }
        return [];
    }
    collection.findGroup = findGroup;
    function flatten(a) {
        const ret = [];
        a.forEach(v => ret.push(...v));
        return ret;
    }
    collection.flatten = flatten;
})(collection || (collection = {}));
var command;
(function (command) {
    command.client = {
        error: "error",
        ping: "ping",
        connect: "connect",
        updateSession: "update-session",
        getActions: "get-actions",
        getTeams: "get-teams",
        getSprints: "get-sprints",
        updateProfile: "update-profile",
        removeMember: "remove-member",
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
        removeFeedback: "remove-feedback",
    };
    command.server = {
        error: "error",
        pong: "pong",
        sessionJoined: "session-joined",
        sessionUpdate: "session-update",
        permissionsUpdate: "permissions-update",
        teamUpdate: "team-update",
        sprintUpdate: "sprint-update",
        contentUpdate: "content-update",
        actions: "actions",
        teams: "teams",
        sprints: "sprints",
        memberUpdate: "member-update",
        onlineUpdate: "online-update",
        storyUpdate: "story-update",
        storyRemove: "story-remove",
        storyStatusChange: "story-status-change",
        voteUpdate: "vote-update",
        reportUpdate: "report-update",
        reportRemove: "report-remove",
        feedbackUpdate: "feedback-update",
        feedbackRemove: "feedback-remove",
    };
})(command || (command = {}));
var contents;
(function (contents) {
    function onContentDisplay(key, same, content, html) {
        dom.setDisplay(`#modal-${key} .content-edit`, same);
        dom.setDisplay(`#modal-${key} .buttons-edit`, same);
        dom.setHTML(dom.setDisplay(`#modal-${key} .content-view`, !same), !same ? "" : html);
        dom.setDisplay(`#modal-${key} .buttons-view`, !same);
        const contentEditTextarea = dom.req(`#${key}-edit-content`);
        dom.setValue(contentEditTextarea, same ? content : "");
        if (same) {
            dom.wireTextarea(contentEditTextarea);
        }
    }
    contents.onContentDisplay = onContentDisplay;
})(contents || (contents = {}));
var date;
(function (date_1) {
    function dateToYMD(date) {
        const d = date.getDate();
        const m = date.getMonth() + 1;
        const y = date.getFullYear();
        return `${y}-${m <= 9 ? `0${m}` : m}-${d <= 9 ? `0${d}` : d}`;
    }
    date_1.dateToYMD = dateToYMD;
    function dateFromYMD(s) {
        const d = new Date(s);
        return new Date(d.getTime() + (d.getTimezoneOffset() * 60000));
    }
    date_1.dateFromYMD = dateFromYMD;
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
    date_1.dow = dow;
    function toDateString(d) {
        return d.toLocaleDateString();
    }
    date_1.toDateString = toDateString;
    function toTimeString(d) {
        return d.toLocaleTimeString().slice(0, 8);
    }
    date_1.toTimeString = toTimeString;
    function toDateTimeString(d) {
        return `${toDateString(d)} ${toTimeString(d)}`;
    }
    date_1.toDateTimeString = toDateTimeString;
})(date || (date = {}));
var dom;
(function (dom) {
    function els(selector, context) {
        return UIkit.util.$$(selector, context);
    }
    dom.els = els;
    function opt(selector, context) {
        return els(selector, context).shift();
    }
    dom.opt = opt;
    function req(selector, context) {
        const res = opt(selector, context);
        if (!res) {
            console.warn(`no element found for selector [${selector}]`);
        }
        return res;
    }
    dom.req = req;
    function setHTML(el, html) {
        if (typeof el === "string") {
            el = req(el);
        }
        el.innerHTML = html;
        return el;
    }
    dom.setHTML = setHTML;
    function setDisplay(el, condition, v = "block") {
        if (typeof el === "string") {
            el = req(el);
        }
        el.style.display = condition ? v : "none";
        return el;
    }
    dom.setDisplay = setDisplay;
    function setContent(el, e) {
        if (typeof el === "string") {
            el = req(el);
        }
        el.innerHTML = "";
        el.appendChild(e);
        return el;
    }
    dom.setContent = setContent;
    function setText(el, text) {
        if (typeof el === "string") {
            el = req(el);
        }
        el.innerText = text;
        return el;
    }
    dom.setText = setText;
    function setValue(el, text) {
        if (typeof el === "string") {
            el = req(el);
        }
        el.value = text;
        return el;
    }
    dom.setValue = setValue;
    function wireTextarea(text) {
        function resize() {
            text.style.height = "auto";
            text.style.height = `${text.scrollHeight < 64 ? 64 : (text.scrollHeight + 6)}px`;
        }
        function delayedResize() {
            window.setTimeout(resize, 0);
        }
        const x = text.dataset["autoresize"];
        if (!x) {
            text.dataset["autoresize"] = "true";
            text.addEventListener("change", resize, false);
            text.addEventListener("cut", delayedResize, false);
            text.addEventListener("paste", delayedResize, false);
            text.addEventListener("drop", delayedResize, false);
            text.addEventListener("keydown", delayedResize, false);
            text.focus();
            text.select();
        }
        resize();
    }
    dom.wireTextarea = wireTextarea;
    function setOptions(el, categories) {
        if (typeof el === "string") {
            el = req(el);
        }
        el.innerHTML = "";
        for (const c of categories) {
            const opt = document.createElement("option");
            opt.value = c;
            opt.innerText = c;
            el.appendChild(opt);
        }
    }
    dom.setOptions = setOptions;
    function setSelectOption(el, o) {
        if (typeof el === "string") {
            el = req(el);
        }
        for (let i = 0; i < el.children.length; i++) {
            const e = el.children.item(i);
            e.selected = e.value === o;
        }
    }
    dom.setSelectOption = setSelectOption;
})(dom || (dom = {}));
var estimate;
(function (estimate) {
    class Cache {
        constructor() {
            this.stories = [];
            this.votes = [];
        }
        activeVotes() {
            if (!this.activeStory) {
                return [];
            }
            return this.votes.filter(x => x.storyID === this.activeStory);
        }
    }
    estimate.cache = new Cache();
    function onEstimateMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.estimate.key, param);
                break;
            case command.server.sessionJoined:
                const sj = param;
                rituals.onSessionJoin(sj);
                rituals.setTeam(sj.team);
                rituals.setSprint(sj.sprint);
                setEstimateDetail(sj.session);
                story.setStories(sj.stories);
                vote.setVotes(sj.votes);
                rituals.showWelcomeMessage(sj.members.length);
                break;
            case command.server.sessionUpdate:
                setEstimateDetail(param);
                break;
            case command.server.permissionsUpdate:
                system.setPermissions(param);
                break;
            case command.server.teamUpdate:
                const tm = param;
                if (estimate.cache.detail) {
                    estimate.cache.detail.teamID = tm === null || tm === void 0 ? void 0 : tm.id;
                }
                rituals.setTeam(tm);
                break;
            case command.server.sprintUpdate:
                const spr = param;
                if (estimate.cache.detail) {
                    estimate.cache.detail.sprintID = spr === null || spr === void 0 ? void 0 : spr.id;
                }
                rituals.setSprint(spr);
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
                console.warn(`unhandled command [${cmd}] for estimate`);
        }
    }
    estimate.onEstimateMessage = onEstimateMessage;
    function setEstimateDetail(detail) {
        estimate.cache.detail = detail;
        dom.setValue("#model-choices-input", detail.choices.join(", "));
        story.viewActiveStory();
        rituals.setDetail(detail);
    }
    function onSubmitEstimateSession() {
        const title = dom.req("#model-title-input").value;
        const choices = dom.req("#model-choices-input").value;
        const teamID = dom.req("#model-team-select select").value;
        const sprintID = dom.req("#model-sprint-select select").value;
        const permissions = permission.readPermissions();
        const msg = { svc: services.estimate.key, cmd: command.client.updateSession, param: { title, choices, teamID, sprintID, permissions } };
        socket.send(msg);
    }
    estimate.onSubmitEstimateSession = onSubmitEstimateSession;
    function onStoryUpdate(s) {
        const x = preUpdate(s.id);
        x.push(s);
        if (s.id === estimate.cache.activeStory) {
            dom.setText("#story-title", s.title);
        }
        story.setStories(x);
    }
    estimate.onStoryUpdate = onStoryUpdate;
    function onStoryRemove(id) {
        const x = preUpdate(id);
        story.setStories(x);
        if (id === estimate.cache.activeStory) {
            UIkit.modal("#modal-story").hide();
        }
        UIkit.notification("story has been deleted", { status: "success", pos: "top-right" });
    }
    estimate.onStoryRemove = onStoryRemove;
    function preUpdate(id) {
        return estimate.cache.stories.filter((p) => p.id !== id);
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
                const sessionInput = dom.setValue("#model-title-input", dom.req("#model-title").innerText);
                delay(() => sessionInput.focus());
                team.refreshTeams();
                sprint.refreshSprints();
                break;
            // member
            case "self":
                const selfInput = dom.setValue("#self-name-input", dom.req("#member-self .member-name").innerText);
                delay(() => selfInput.focus());
                break;
            case "invitation":
                break;
            case "member":
                system.cache.activeMember = id;
                member.viewActiveMember();
                break;
            case "welcome":
                break;
            // actions
            case "actions":
                action.loadActions();
                break;
            // estimate
            case "add-story":
                const storyInput = dom.setValue("#story-title-input", "");
                delay(() => storyInput.focus());
                break;
            case "story":
                estimate.cache.activeStory = id;
                story.viewActiveStory();
                break;
            // standup
            case "add-report":
                dom.setValue("#report-date", date.dateToYMD(new Date()));
                const reportContent = dom.setValue("#report-content", "");
                dom.wireTextarea(reportContent);
                delay(() => reportContent.focus());
                break;
            case "report":
                standup.cache.activeReport = id;
                report.viewActiveReport();
                const reportEditContent = dom.req("#report-edit-content");
                delay(() => {
                    dom.wireTextarea(reportEditContent);
                    reportEditContent.focus();
                });
                break;
            // retro
            case "add-feedback":
                dom.setSelectOption("#feedback-category", id);
                const feedbackContent = dom.setValue("#feedback-content", "");
                dom.wireTextarea(feedbackContent);
                delay(() => feedbackContent.focus());
                break;
            case "feedback":
                retro.cache.activeFeedback = id;
                feedback.viewActiveFeedback();
                const feedbackEditContent = dom.req("#feedback-edit-content");
                delay(() => {
                    dom.wireTextarea(feedbackEditContent);
                    feedbackEditContent.focus();
                });
                break;
            // default
            default:
                console.warn(`unhandled modal [${key}]`);
        }
        UIkit.modal(`#modal-${key}`).show();
        return false;
    }
    events.openModal = openModal;
})(events || (events = {}));
var feedback;
(function (feedback_1) {
    function setFeedback(feedback) {
        retro.cache.feedback = feedback;
        dom.setContent("#feedback-detail", feedback_1.renderFeedbackArray(feedback));
        UIkit.modal("#modal-add-feedback").hide();
    }
    feedback_1.setFeedback = setFeedback;
    function onSubmitFeedback() {
        const category = dom.req("#feedback-category").value;
        const content = dom.req("#feedback-content").value;
        const msg = { svc: services.retro.key, cmd: command.client.addFeedback, param: { category, content } };
        socket.send(msg);
        return false;
    }
    feedback_1.onSubmitFeedback = onSubmitFeedback;
    function onEditFeedback() {
        const id = retro.cache.activeFeedback;
        const category = dom.req("#feedback-edit-category").value;
        const content = dom.req("#feedback-edit-content").value;
        const msg = { svc: services.retro.key, cmd: command.client.updateFeedback, param: { id, category, content } };
        socket.send(msg);
        return false;
    }
    feedback_1.onEditFeedback = onEditFeedback;
    function onRemoveFeedback() {
        const id = retro.cache.activeFeedback;
        if (id) {
            UIkit.modal.confirm("Delete this feedback?").then(function () {
                const msg = { svc: services.retro.key, cmd: command.client.removeFeedback, param: id };
                socket.send(msg);
                UIkit.modal("#modal-feedback").hide();
            });
        }
        return false;
    }
    feedback_1.onRemoveFeedback = onRemoveFeedback;
    function getActiveFeedback() {
        if (!retro.cache.activeFeedback) {
            return undefined;
        }
        const curr = retro.cache.feedback.filter(x => x.id === retro.cache.activeFeedback).shift();
        if (!curr) {
            console.warn(`cannot load active Feedback [${retro.cache.activeFeedback}]`);
        }
        return curr;
    }
    feedback_1.getActiveFeedback = getActiveFeedback;
    function viewActiveFeedback() {
        const profile = system.cache.getProfile();
        const fb = getActiveFeedback();
        if (!fb) {
            console.warn("no active feedback");
            return;
        }
        const same = fb.authorID === profile.userID;
        dom.setText("#feedback-title", `${fb.category} / ${system.getMemberName(fb.authorID)}`);
        dom.setSelectOption("#feedback-edit-category", same ? fb.category : undefined);
        contents.onContentDisplay("report", same, fb.content, fb.html);
    }
    feedback_1.viewActiveFeedback = viewActiveFeedback;
    function onFeedbackUpdate(r) {
        const x = preUpdate(r.id);
        x.push(r);
        postUpdate(x, r.id);
    }
    feedback_1.onFeedbackUpdate = onFeedbackUpdate;
    function onFeedbackRemoved(id) {
        const x = preUpdate(id);
        postUpdate(x, id);
        UIkit.notification("feedback has been deleted", { status: "success", pos: "top-right" });
    }
    feedback_1.onFeedbackRemoved = onFeedbackRemoved;
    function preUpdate(id) {
        return retro.cache.feedback.filter((p) => p.id !== id);
    }
    function postUpdate(x, id) {
        feedback.setFeedback(x);
        if (id === retro.cache.activeFeedback) {
            UIkit.modal("#modal-feedback").hide();
        }
    }
    function getFeedbackCategories(feedback, categories) {
        function toCollection(c) {
            const reports = feedback.filter(r => r.category === c).sort((l, r) => (l.created > r.created ? -1 : 1));
            return { category: c, feedback: reports };
        }
        const ret = categories.map(toCollection);
        const extras = feedback.filter(r => !categories.find(x => x === r.category));
        if (extras.length > 0) {
            ret.push({ category: "unknown", feedback: extras });
        }
        return ret;
    }
    feedback_1.getFeedbackCategories = getFeedbackCategories;
})(feedback || (feedback = {}));
// noinspection JSUnusedGlobalSymbols
function JSX(tag, attrs) {
    const e = document.createElement(tag);
    for (const name in attrs) {
        if (name && attrs.hasOwnProperty(name)) {
            const v = attrs[name];
            if (v === true) {
                e.setAttribute(name, name);
            }
            else if (v !== false && v !== null && v !== undefined) {
                e.setAttribute(name, v.toString());
            }
        }
    }
    for (let i = 2; i < arguments.length; i++) {
        let child = arguments[i];
        if (Array.isArray(child)) {
            child.forEach(c => {
                e.appendChild(c);
            });
        }
        else {
            if (!child.nodeType) {
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
        var _a;
        return x.userID === ((_a = system.cache.profile) === null || _a === void 0 ? void 0 : _a.userID);
    }
    function setMembers() {
        const self = system.cache.members.filter(isSelf).shift();
        if (self) {
            dom.setText("#member-self .member-name", self.name);
            dom.setValue("#self-name-input", self.name);
            dom.setText("#member-self .member-role", self.role);
        }
        const others = system.cache.members.filter(x => !isSelf(x));
        dom.setContent("#member-detail", member_1.renderMembers(others));
        renderOnline();
    }
    member_1.setMembers = setMembers;
    function onMemberUpdate(member) {
        var _a;
        if (isSelf(member)) {
            UIkit.modal("#modal-self").hide();
        }
        const unfiltered = system.cache.members;
        const curr = unfiltered.filter(m => m.userID === member.userID).shift();
        const nameChanged = (curr === null || curr === void 0 ? void 0 : curr.name) !== member.name;
        const ms = unfiltered.filter(m => m.userID !== member.userID);
        if (ms.length === system.cache.members.length) {
            UIkit.notification(`${member.name} has joined`, { status: "success", pos: "top-right" });
        }
        if (member.name === "::delete") {
            if (member.userID === ((_a = system.cache.profile) === null || _a === void 0 ? void 0 : _a.userID)) {
                UIkit.modal("#modal-self").hide();
                UIkit.notification(`you have left this ${system.cache.currentService}`, { status: "success", pos: "top-right" });
                document.location.href = "/";
            }
            else {
                UIkit.modal("#modal-member").hide();
            }
        }
        else {
            ms.push(member);
        }
        ms.sort((l, r) => (l.name > r.name) ? 1 : -1);
        system.cache.members = ms;
        setMembers();
        if (nameChanged) {
            switch (system.cache.currentService) {
                case services.team.key:
                    break;
                case services.sprint.key:
                    break;
                case services.estimate.key:
                    if (estimate.cache.activeStory) {
                        vote.viewVotes();
                    }
                    break;
                case services.standup.key:
                    dom.setContent("#report-detail", report.renderReports(standup.cache.reports));
                    if (standup.cache.activeReport) {
                        report.viewActiveReport();
                    }
                    break;
                case services.retro.key:
                    dom.setContent("#feedback-detail", feedback.renderFeedbackArray(retro.cache.feedback));
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
            if (!system.cache.online.find(x => x === update.userID)) {
                system.cache.online.push(update.userID);
            }
        }
        else {
            system.cache.online = system.cache.online.filter(x => x !== update.userID);
        }
        renderOnline();
    }
    member_1.onOnlineUpdate = onOnlineUpdate;
    function renderOnline() {
        for (const member of system.cache.members) {
            const el = dom.opt(`#member-${member.userID} .online-indicator`);
            if (el) {
                if (!system.cache.online.find(x => x === member.userID)) {
                    el.classList.add("offline");
                }
                else {
                    el.classList.remove("offline");
                }
            }
        }
    }
    function onSubmitSelf() {
        const name = dom.req("#self-name-input").value;
        const choice = dom.req("#self-name-choice-global").checked ? "global" : "local";
        const msg = { svc: services.system.key, cmd: command.client.updateProfile, param: { name, choice } };
        socket.send(msg);
    }
    member_1.onSubmitSelf = onSubmitSelf;
    function getActiveMember() {
        if (!system.cache.activeMember) {
            console.warn("no active member");
            return undefined;
        }
        const curr = system.cache.members.filter(x => x.userID === system.cache.activeMember).shift();
        if (curr) {
            console.warn(`cannot load active member [${system.cache.activeMember}]`);
        }
        return curr;
    }
    function viewActiveMember() {
        const member = getActiveMember();
        if (!member) {
            return;
        }
        dom.setText("#member-modal-name", member.name);
        dom.setText("#member-modal-role", member.role);
    }
    member_1.viewActiveMember = viewActiveMember;
    function removeMember(id = system.cache.activeMember) {
        var _a;
        if (!id) {
            console.warn(`cannot load active member [${system.cache.activeMember}]`);
        }
        if (id == "self") {
            id = (_a = system.cache.profile) === null || _a === void 0 ? void 0 : _a.userID;
        }
        if (confirm(`Are you sure you wish to leave this ${system.cache.currentService}?`)) {
            const msg = { svc: system.cache.currentService, cmd: command.client.removeMember, param: id };
            socket.send(msg);
        }
    }
    member_1.removeMember = removeMember;
})(member || (member = {}));
var permission;
(function (permission) {
    function setPerms() {
        ["team", "sprint"].forEach(setModelPerms);
        if (system.cache.auths != null) {
            auth.allProviders.forEach(setProviderPerms);
        }
    }
    permission.setPerms = setPerms;
    function setModelPerms(key) {
        const el = dom.opt(`#model-${key}-select select`);
        if (el) {
            const perms = collection.findGroup(system.cache.permissions, key);
            const section = dom.opt(`#perm-${key}-section`);
            if (section) {
                const checkbox = dom.req(`#perm-${key}-checkbox`);
                checkbox.checked = perms.length > 0;
                dom.setDisplay(section, el.value != "");
            }
            collection.findGroup(system.cache.permissions, key);
        }
    }
    permission.setModelPerms = setModelPerms;
    function onChanged(k, v, checked) {
        switch (k) {
            case "email":
                return onEmailChanged(v, checked);
            case "provider":
                return onProviderChanged(v, checked);
        }
    }
    permission.onChanged = onChanged;
    function onEmailChanged(key, checked) {
        const checkbox = dom.opt(`#perm-${key}-checkbox`);
        if (checkbox && checked && !checkbox.checked) {
            checkbox.checked = true;
        }
    }
    function onProviderChanged(key, checked) {
        dom.els(`.perm-${key}-email`).forEach(el => {
            el.disabled = !checked;
            if (!checked) {
                el.checked = false;
            }
        });
    }
    function setProviderPerms(p) {
        const perms = collection.findGroup(system.cache.permissions, p.key);
        const auths = system.cache.auths.filter(a => a.provider == p.key);
        const section = dom.opt(`#perm-${p.key}-section`);
        if (section) {
            const checkbox = dom.req(`#perm-${p.key}-checkbox`);
            checkbox.checked = perms.length > 0;
            const emailContainer = dom.req(`#perm-${p.key}-email-container`);
            const emails = collection.flatten(perms.map(x => x.v.split(",").filter(x => x.length > 0))).map(x => ({ matched: true, domain: x }));
            const additional = auths.filter(a => emails.filter(e => a.email.endsWith(e.domain)).length == 0).map(m => {
                return { matched: false, domain: getDomain(m.email) };
            });
            emails.push(...additional);
            emails.sort();
            dom.setDisplay(emailContainer, emails.length > 0);
            dom.setContent(emailContainer, emails.length == 0 ? document.createElement("span") : permission.renderEmails(p.key, emails));
        }
    }
    function getDomain(email) {
        const idx = email.lastIndexOf("@");
        if (idx == -1) {
            return email;
        }
        return email.substr(idx);
    }
})(permission || (permission = {}));
var profile;
(function (profile) {
    // noinspection JSUnusedGlobalSymbols
    function setNavColor(el, c) {
        dom.setValue("#navbar-color", c);
        const nb = dom.req("#navbar");
        nb.className = `${c}-bg uk-navbar-container uk-navbar`;
        const colors = document.querySelectorAll(".navbar_swatch");
        colors.forEach(function (i) {
            i.classList.remove("active");
        });
        el.classList.add("active");
    }
    profile.setNavColor = setNavColor;
    // noinspection JSUnusedGlobalSymbols
    function setLinkColor(el, c) {
        dom.setValue("#link-color", c);
        const links = dom.els(".profile-link");
        links.forEach(l => {
            l.classList.forEach(x => {
                if (x.indexOf("-fg") > -1) {
                    l.classList.remove(x);
                }
                l.classList.add(`${c}-fg`);
            });
        });
        const colors = document.querySelectorAll(".link_swatch");
        colors.forEach(function (i) {
            i.classList.remove("active");
        });
        el.classList.add("active");
    }
    profile.setLinkColor = setLinkColor;
    function selectTheme(theme) {
        const card = dom.els(".uk-card");
        switch (theme) {
            case "light":
                document.documentElement.classList.remove("uk-light");
                document.body.classList.remove("uk-light");
                document.documentElement.classList.add("uk-dark");
                document.body.classList.add("uk-dark");
                card.forEach(x => {
                    x.classList.add("uk-card-default");
                    x.classList.remove("uk-card-secondary");
                });
                break;
            case "dark":
                document.documentElement.classList.add("uk-light");
                document.body.classList.add("uk-light");
                document.documentElement.classList.remove("uk-dark");
                document.body.classList.remove("uk-dark");
                card.forEach(x => {
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
        const d = dom.req("#report-date").value;
        const content = dom.req("#report-content").value;
        const msg = { svc: services.standup.key, cmd: command.client.addReport, param: { d, content } };
        socket.send(msg);
        return false;
    }
    report_1.onSubmitReport = onSubmitReport;
    function onEditReport() {
        const d = dom.req("#report-edit-date").value;
        const content = dom.req("#report-edit-content").value;
        const msg = { svc: services.standup.key, cmd: command.client.updateReport, param: { id: standup.cache.activeReport, d, content } };
        socket.send(msg);
        return false;
    }
    report_1.onEditReport = onEditReport;
    function onRemoveReport() {
        const id = standup.cache.activeReport;
        if (id) {
            UIkit.modal.confirm("Delete this report?").then(function () {
                const msg = { svc: services.standup.key, cmd: command.client.removeReport, param: id };
                socket.send(msg);
                UIkit.modal("#modal-report").hide();
            });
        }
        return false;
    }
    report_1.onRemoveReport = onRemoveReport;
    function getActiveReport() {
        if (!standup.cache.activeReport) {
            console.warn("no active report");
            return undefined;
        }
        const curr = standup.cache.reports.filter(x => x.id === standup.cache.activeReport).shift();
        if (!curr) {
            console.warn(`cannot load active report [${standup.cache.activeReport}]`);
        }
        return curr;
    }
    function viewActiveReport() {
        const profile = system.cache.getProfile();
        const report = getActiveReport();
        if (!report) {
            console.warn("no active report");
            return;
        }
        dom.setText("#report-title", `${report.d} / ${system.getMemberName(report.authorID)}`);
        setFor(report, profile.userID);
    }
    report_1.viewActiveReport = viewActiveReport;
    function setReports(reports) {
        standup.cache.reports = reports;
        dom.setContent("#report-detail", report_1.renderReports(reports));
        UIkit.modal("#modal-add-report").hide();
    }
    report_1.setReports = setReports;
    function getReportDates(reports) {
        function distinct(v, i, s) {
            return s.indexOf(v) === i;
        }
        function toCollection(d) {
            const sorted = reports.filter(r => r.d === d).sort((l, r) => (l.created > r.created ? -1 : 1));
            return { "d": d, "reports": sorted };
        }
        return reports.map(r => r.d).filter(distinct).sort().reverse().map(toCollection);
    }
    report_1.getReportDates = getReportDates;
    function setFor(report, userID) {
        const same = report.authorID === userID;
        contents.onContentDisplay("report", same, report.content, report.html);
        dom.setValue(dom.req("#report-edit-date"), same ? report.d : "");
    }
})(report || (report = {}));
var retro;
(function (retro) {
    class Cache {
        constructor() {
            this.feedback = [];
        }
    }
    retro.cache = new Cache();
    function onRetroMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.retro.key, param);
                break;
            case command.server.sessionJoined:
                const sj = param;
                rituals.onSessionJoin(sj);
                rituals.setTeam(sj.team);
                rituals.setSprint(sj.sprint);
                setRetroDetail(sj.session);
                feedback.setFeedback(sj.feedback);
                rituals.showWelcomeMessage(sj.members.length);
                break;
            case command.server.sessionUpdate:
                setRetroDetail(param);
                break;
            case command.server.permissionsUpdate:
                system.setPermissions(param);
                break;
            case command.server.teamUpdate:
                const tm = param;
                if (retro.cache.detail) {
                    retro.cache.detail.teamID = tm === null || tm === void 0 ? void 0 : tm.id;
                }
                rituals.setTeam(tm);
                break;
            case command.server.sprintUpdate:
                const spr = param;
                if (retro.cache.detail) {
                    retro.cache.detail.sprintID = spr === null || spr === void 0 ? void 0 : spr.id;
                }
                rituals.setSprint(spr);
                break;
            case command.server.feedbackUpdate:
                feedback.onFeedbackUpdate(param);
                break;
            case command.server.feedbackRemove:
                feedback.onFeedbackRemoved(param);
                break;
            default:
                console.warn(`unhandled command [${cmd}] for retro`);
        }
    }
    retro.onRetroMessage = onRetroMessage;
    function setRetroDetail(detail) {
        retro.cache.detail = detail;
        dom.setValue("#model-categories-input", detail.categories.join(", "));
        dom.setOptions("#feedback-category", detail.categories);
        dom.setOptions("#feedback-edit-category", detail.categories);
        feedback.setFeedback(retro.cache.feedback);
        rituals.setDetail(detail);
    }
    function onSubmitRetroSession() {
        const title = dom.req("#model-title-input").value;
        const categories = dom.req("#model-categories-input").value;
        const teamID = dom.req("#model-team-select select").value;
        const sprintID = dom.req("#model-sprint-select select").value;
        const permissions = permission.readPermissions();
        const msg = { svc: services.retro.key, cmd: command.client.updateSession, param: { title, categories, teamID, sprintID, permissions } };
        socket.send(msg);
    }
    retro.onSubmitRetroSession = onSubmitRetroSession;
})(retro || (retro = {}));
var rituals;
(function (rituals) {
    function onSocketMessage(msg) {
        if (debug) {
            console.debug("message received");
            console.debug(msg);
        }
        switch (msg.svc) {
            case services.system.key:
                onSystemMessage(msg.cmd, msg.param);
                break;
            case services.team.key:
                team.onTeamMessage(msg.cmd, msg.param);
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
                console.warn(`unhandled message for service [${msg.svc}]`);
        }
    }
    rituals.onSocketMessage = onSocketMessage;
    function setDetail(session) {
        system.cache.session = session;
        dom.setText("#model-title", session.title);
        dom.setValue("#model-title-input", session.title);
        const items = dom.els("#navbar .uk-navbar-item");
        if (items.length > 0) {
            items[items.length - 1].innerText = session.title;
        }
        UIkit.modal("#modal-session").hide();
    }
    rituals.setDetail = setDetail;
    function onError(svc, err) {
        console.warn(`${svc}: ${err}`);
        const idx = err.lastIndexOf(":");
        if (idx > -1) {
            err = err.substr(idx + 1);
        }
        UIkit.notification(`${svc} error: ${err}`, { status: "danger", pos: "top-right" });
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
            case command.server.teams:
                team.viewTeams(param);
                break;
            case command.server.sprints:
                sprint.viewSprints(param);
                break;
            case command.server.memberUpdate:
                member.onMemberUpdate(param);
                break;
            case command.server.onlineUpdate:
                member.onOnlineUpdate(param);
                break;
            default:
                console.warn(`unhandled system message for command [${cmd}]`);
        }
    }
    function onSessionJoin(param) {
        system.cache.session = param.session;
        system.cache.profile = param.profile;
        system.cache.auths = param.auths;
        permission.applyPermissions(param.permissions);
        permission.setPerms();
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
    function setSprint(spr) {
        UIkit.modal("#modal-session").hide();
        const lc = dom.req("#sprint-link-container");
        lc.innerHTML = "";
        if (spr) {
            lc.appendChild(sprint.renderSprintLink(spr));
            dom.req("#sprint-warning-name").innerText = spr.title;
        }
    }
    rituals.setSprint = setSprint;
    function setTeam(tm) {
        UIkit.modal("#modal-session").hide();
        const container = dom.req("#team-link-container");
        container.innerHTML = "";
        if (tm) {
            container.appendChild(team.renderTeamLink(tm));
        }
    }
    rituals.setTeam = setTeam;
    function showWelcomeMessage(count) {
        if (count === 1) {
            setTimeout(() => events.openModal("welcome"), 300);
        }
    }
    rituals.showWelcomeMessage = showWelcomeMessage;
})(rituals || (rituals = {}));
var services;
(function (services) {
    services.system = { key: "system", title: "System", plural: "systems", icon: "close" };
    services.team = { key: "team", title: "Team", plural: "teams", icon: "users" };
    services.sprint = { key: "sprint", title: "Sprint", plural: "sprints", icon: "git-fork" };
    services.estimate = { key: "estimate", title: "Estimate Session", plural: "estimates", icon: "settings" };
    services.standup = { key: "standup", title: "Daily Standup", plural: "standups", icon: "future" };
    services.retro = { key: "retro", title: "Retrospective", plural: "retros", icon: "history" };
})(services || (services = {}));
var socket;
(function (socket_1) {
    let socket;
    let appUnloading = false;
    function socketUrl() {
        const l = document.location;
        let protocol = "ws";
        if (l.protocol === "https:") {
            protocol = "wss";
        }
        return protocol + `://${l.host}/s`;
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
            if (debug) {
                console.debug("socket connected");
            }
            const msg = { svc: svc, cmd: command.client.connect, param: id };
            send(msg);
        };
        socket.onmessage = function (event) {
            const msg = JSON.parse(event.data);
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
        if (debug) {
            console.debug("sending message");
            console.debug(msg);
        }
        socket.send(JSON.stringify(msg));
    }
    socket_1.send = send;
    function onSocketClose() {
        function disconnect(seconds) {
            if (debug) {
                console.info(`socket closed, reconnecting in ${seconds} seconds`);
            }
            setTimeout(() => {
                socketConnect(system.cache.currentService, system.cache.currentID);
            }, seconds * 1000);
        }
        if (!appUnloading) {
            disconnect(10);
        }
    }
})(socket || (socket = {}));
var sprint;
(function (sprint) {
    class Cache {
    }
    sprint.cache = new Cache();
    function onSprintMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.sprint.key, param);
                break;
            case command.server.sessionJoined:
                const sj = param;
                rituals.onSessionJoin(sj);
                setSprintDetail(sj.session);
                rituals.setTeam(sj.team);
                setSprintContents(sj);
                rituals.showWelcomeMessage(sj.members.length);
                break;
            case command.server.teamUpdate:
                const tm = param;
                if (sprint.cache.detail) {
                    sprint.cache.detail.teamID = tm === null || tm === void 0 ? void 0 : tm.id;
                }
                rituals.setTeam(tm);
                break;
            case command.server.sessionUpdate:
                setSprintDetail(param);
                break;
            case command.server.permissionsUpdate:
                system.setPermissions(param);
                break;
            case command.server.contentUpdate:
                socket.socketConnect(system.cache.currentService, system.cache.currentID);
                break;
            default:
                console.warn(`unhandled command [${cmd}] for sprint`);
        }
    }
    sprint.onSprintMessage = onSprintMessage;
    function setSprintDetail(detail) {
        var _a, _b;
        sprint.cache.detail = detail;
        const s = ((_a = detail.startDate) === null || _a === void 0 ? void 0 : _a.length) === 0 ? undefined : new Date(detail.startDate);
        const e = ((_b = detail.endDate) === null || _b === void 0 ? void 0 : _b.length) === 0 ? undefined : new Date(detail.endDate);
        dom.setContent("#sprint-date-display", sprint.renderSprintDates(s, e));
        dom.setValue("#sprint-start-date-input", s ? date.dateToYMD(s) : "");
        dom.setValue("#sprint-end-date-input", e ? date.dateToYMD(e) : "");
        rituals.setDetail(detail);
    }
    function setSprintContents(sj) {
        dom.setContent("#sprint-estimate-list", contents.renderContents(services.sprint, services.estimate, sj.estimates));
        dom.setContent("#sprint-standup-list", contents.renderContents(services.sprint, services.standup, sj.standups));
        dom.setContent("#sprint-retro-list", contents.renderContents(services.sprint, services.retro, sj.retros));
    }
    function onSubmitSprintSession() {
        var _a, _b;
        const title = dom.req("#model-title-input").value;
        const teamID = dom.req("#model-team-select select").value;
        const startDate = (_a = dom.opt("#model-start-date-input")) === null || _a === void 0 ? void 0 : _a.value;
        const endDate = (_b = dom.opt("#model-end-date-input")) === null || _b === void 0 ? void 0 : _b.value;
        const permissions = permission.readPermissions();
        const msg = { svc: services.sprint.key, cmd: command.client.updateSession, param: { title, startDate, endDate, teamID, permissions } };
        socket.send(msg);
    }
    sprint.onSubmitSprintSession = onSubmitSprintSession;
    function refreshSprints() {
        const sprintSelect = dom.opt("#model-sprint-select");
        if (sprintSelect) {
            socket.send({ svc: services.system.key, cmd: command.client.getSprints, param: null });
        }
    }
    sprint.refreshSprints = refreshSprints;
    function viewSprints(sprints) {
        var _a;
        const c = dom.opt("#model-sprint-container");
        if (c) {
            dom.setDisplay(c, sprints.length > 0);
            dom.setContent("#model-sprint-select", sprint.renderSprintSelect(sprints, (_a = system.cache.session) === null || _a === void 0 ? void 0 : _a.sprintID));
            permission.setModelPerms("sprint");
        }
    }
    sprint.viewSprints = viewSprints;
})(sprint || (sprint = {}));
var standup;
(function (standup) {
    class Cache {
        constructor() {
            this.reports = [];
        }
    }
    standup.cache = new Cache();
    function onStandupMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.standup.key, param);
                break;
            case command.server.sessionJoined:
                const sj = param;
                rituals.onSessionJoin(sj);
                rituals.setTeam(sj.team);
                rituals.setSprint(sj.sprint);
                setStandupDetail(sj.session);
                report.setReports(sj.reports);
                rituals.showWelcomeMessage(sj.members.length);
                break;
            case command.server.sessionUpdate:
                setStandupDetail(param);
                break;
            case command.server.permissionsUpdate:
                system.setPermissions(param);
                break;
            case command.server.teamUpdate:
                const tm = param;
                if (standup.cache.detail) {
                    standup.cache.detail.teamID = tm === null || tm === void 0 ? void 0 : tm.id;
                }
                rituals.setTeam(tm);
                break;
            case command.server.sprintUpdate:
                const x = param;
                if (standup.cache.detail) {
                    standup.cache.detail.sprintID = x === null || x === void 0 ? void 0 : x.id;
                }
                rituals.setSprint(x);
                break;
            case command.server.reportUpdate:
                onReportUpdate(param);
                break;
            case command.server.reportRemove:
                onReportRemoved(param);
                break;
            default:
                console.warn(`unhandled command [${cmd}] for standup`);
        }
    }
    standup.onStandupMessage = onStandupMessage;
    function setStandupDetail(detail) {
        standup.cache.detail = detail;
        rituals.setDetail(detail);
    }
    function onSubmitStandupSession() {
        const title = dom.req("#model-title-input").value;
        const teamID = dom.req("#model-team-select select").value;
        const sprintID = dom.req("#model-sprint-select select").value;
        const permissions = permission.readPermissions();
        const msg = { svc: services.standup.key, cmd: command.client.updateSession, param: { title, teamID, sprintID, permissions } };
        socket.send(msg);
    }
    standup.onSubmitStandupSession = onSubmitStandupSession;
    function onReportUpdate(r) {
        const x = preUpdate(r.id);
        x.push(r);
        postUpdate(x, r.id);
    }
    function onReportRemoved(id) {
        const x = preUpdate(id);
        postUpdate(x, id);
        UIkit.notification("report has been deleted", { status: "success", pos: "top-right" });
    }
    function preUpdate(id) {
        return standup.cache.reports.filter((p) => p.id !== id);
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
            const s = el.id.substr(el.id.lastIndexOf("-") + 1);
            if (s === status) {
                el.classList.add("active");
            }
            else {
                el.classList.remove("active");
            }
        }
        for (const el of dom.els(".story-status-body")) {
            setActive(el, status);
        }
        for (const el of dom.els(".story-status-actions")) {
            setActive(el, status);
        }
        let txt = "";
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
        dom.setText("#story-status", txt);
        vote.viewVotes();
    }
    story_1.viewStoryStatus = viewStoryStatus;
    function requestStoryStatus(status) {
        const story = story_1.getActiveStory();
        if (!story) {
            return;
        }
        const msg = { svc: services.estimate.key, cmd: command.client.setStoryStatus, param: { storyID: story.id, status } };
        socket.send(msg);
    }
    story_1.requestStoryStatus = requestStoryStatus;
    function setStoryStatus(storyID, status, currStory, calcTotal) {
        if (currStory && currStory.status === "complete") {
            if (currStory.finalVote.length > 0) {
                status = currStory.finalVote;
            }
        }
        dom.setContent(`#story-${storyID} .story-status`, story_1.renderStatus(status));
        if (calcTotal) {
            story_1.showTotalIfNeeded();
        }
    }
    story_1.setStoryStatus = setStoryStatus;
    function onStoryStatusChange(u) {
        let currStory;
        estimate.cache.stories.forEach(s => {
            if (s.id === u.storyID) {
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
        dom.setContent("#story-detail", story.renderStories(stories));
        stories.forEach(s => story.setStoryStatus(s.id, s.status, s, false));
        showTotalIfNeeded();
        UIkit.modal("#modal-add-story").hide();
    }
    story.setStories = setStories;
    function onSubmitStory() {
        const title = dom.req("#story-title-input").value;
        const msg = { svc: services.estimate.key, cmd: command.client.addStory, param: { title } };
        socket.send(msg);
        return false;
    }
    story.onSubmitStory = onSubmitStory;
    function beginEditStory() {
        const s = getActiveStory();
        const title = prompt("Edit your story", s.title);
        if (title && title !== s.title) {
            const msg = { svc: services.estimate.key, cmd: command.client.updateStory, param: { id: s.id, title } };
            socket.send(msg);
        }
        return false;
    }
    story.beginEditStory = beginEditStory;
    function onRemoveStory() {
        const id = estimate.cache.activeStory;
        if (id) {
            UIkit.modal.confirm("Delete this story?").then(function () {
                const msg = { svc: services.estimate.key, cmd: command.client.removeStory, param: id };
                socket.send(msg);
                UIkit.modal("#modal-story").hide();
            });
        }
        return false;
    }
    story.onRemoveStory = onRemoveStory;
    function getActiveStory() {
        if (!estimate.cache.activeStory) {
            return undefined;
        }
        const curr = estimate.cache.stories.filter(x => x.id === estimate.cache.activeStory).shift();
        if (curr) {
            console.warn(`cannot load active story [${estimate.cache.activeStory}]`);
        }
        return curr;
    }
    story.getActiveStory = getActiveStory;
    function viewActiveStory() {
        const s = getActiveStory();
        if (!s) {
            return;
        }
        dom.setText("#story-title", s.title);
        story.viewStoryStatus(s.status);
    }
    story.viewActiveStory = viewActiveStory;
    function showTotalIfNeeded() {
        const stories = estimate.cache.stories;
        const strings = stories.filter(s => s.status === "complete").map(s => s.finalVote).filter(c => c.length > 0);
        const floats = strings.map(c => parseFloat(c)).filter(f => !isNaN(f));
        let sum = 0;
        floats.forEach(f => sum += f);
        const curr = dom.opt("#story-total");
        const panel = dom.req("#story-list");
        if (curr !== undefined) {
            panel.removeChild(curr);
        }
        if (sum > 0) {
            panel.appendChild(story.renderTotal(sum));
        }
    }
    story.showTotalIfNeeded = showTotalIfNeeded;
})(story || (story = {}));
const debug = true;
var system;
(function (system) {
    class Cache {
        constructor() {
            this.currentService = "";
            this.currentID = "";
            this.connectTime = 0;
            this.permissions = [];
            this.auths = [];
            this.members = [];
            this.online = [];
        }
        getProfile() {
            if (!this.profile) {
                throw "no active profile";
            }
            return this.profile;
        }
    }
    function getMemberName(id) {
        const ret = system.cache.members.filter(m => m.userID === id).shift();
        if (ret) {
            return ret.name;
        }
        return "{former member}";
    }
    system.getMemberName = getMemberName;
    system.cache = new Cache();
    function setPermissions(perms) {
        permission.applyPermissions(perms);
        permission.setPerms();
    }
    system.setPermissions = setPermissions;
    function setAuths(auths) {
        system.cache.auths = auths;
        permission.setPerms();
    }
    system.setAuths = setAuths;
})(system || (system = {}));
var team;
(function (team) {
    class Cache {
    }
    team.cache = new Cache();
    function onTeamMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.team.key, param);
                break;
            case command.server.sessionJoined:
                const sj = param;
                rituals.onSessionJoin(sj);
                setTeamDetail(sj.session);
                setTeamHistory(sj);
                rituals.showWelcomeMessage(sj.members.length);
                break;
            case command.server.sessionUpdate:
                setTeamDetail(param);
                break;
            case command.server.permissionsUpdate:
                system.setPermissions(param);
                break;
            case command.server.contentUpdate:
                socket.socketConnect(system.cache.currentService, system.cache.currentID);
                break;
            default:
                console.warn(`unhandled command [${cmd}] for team`);
        }
    }
    team.onTeamMessage = onTeamMessage;
    function setTeamDetail(detail) {
        team.cache.detail = detail;
        rituals.setDetail(detail);
    }
    function setTeamHistory(sj) {
        dom.setContent("#team-sprint-list", contents.renderContents(services.team, services.sprint, sj.sprints));
        dom.setContent("#team-estimate-list", contents.renderContents(services.team, services.estimate, sj.estimates));
        dom.setContent("#team-standup-list", contents.renderContents(services.team, services.standup, sj.standups));
        dom.setContent("#team-retro-list", contents.renderContents(services.team, services.retro, sj.retros));
    }
    function onSubmitTeamSession() {
        const title = dom.req("#model-title-input").value;
        const permissions = permission.readPermissions();
        const msg = { svc: services.team.key, cmd: command.client.updateSession, param: { title, permissions } };
        socket.send(msg);
    }
    team.onSubmitTeamSession = onSubmitTeamSession;
    function refreshTeams() {
        const teamSelect = dom.opt("#model-team-select");
        if (teamSelect) {
            socket.send({ svc: services.system.key, cmd: command.client.getTeams, param: null });
        }
    }
    team.refreshTeams = refreshTeams;
    function viewTeams(teams) {
        var _a;
        const c = dom.opt("#model-team-container");
        if (c) {
            dom.setDisplay(c, teams.length > 0);
            dom.setContent("#model-team-select", team.renderTeamSelect(teams, (_a = system.cache.session) === null || _a === void 0 ? void 0 : _a.teamID));
            permission.setModelPerms("team");
        }
    }
    team.viewTeams = viewTeams;
})(team || (team = {}));
var vote;
(function (vote) {
    function setVotes(votes) {
        estimate.cache.votes = votes;
        viewVotes();
    }
    vote.setVotes = setVotes;
    function onVoteUpdate(v) {
        let x = estimate.cache.votes;
        x = x.filter(v => v.userID !== v.userID || v.storyID !== v.storyID);
        x.push(v);
        estimate.cache.votes = x;
        if (v.storyID === estimate.cache.activeStory) {
            viewVotes();
        }
    }
    vote.onVoteUpdate = onVoteUpdate;
    function viewVotes() {
        var _a;
        const s = story.getActiveStory();
        if (!s) {
            return;
        }
        const votes = estimate.cache.activeVotes();
        const activeVote = votes.filter(v => v.userID === system.cache.getProfile().userID).pop();
        switch (s.status) {
            case "pending":
                const same = ((_a = system.cache.profile) === null || _a === void 0 ? void 0 : _a.userID) === s.authorID;
                dom.setDisplay("#story-edit-section", same);
                dom.setDisplay("#story-view-section", !same);
                break;
            case "active":
                viewActiveVotes(votes, activeVote);
                break;
            case "complete":
                viewVoteResults(votes);
                break;
            default:
                console.warn(`invalid story status [${s.status}]`);
        }
    }
    vote.viewVotes = viewVotes;
    function viewActiveVotes(votes, activeVote) {
        dom.setContent("#story-vote-members", vote.renderVoteMembers(system.cache.members, votes));
        dom.setContent("#story-vote-choices", vote.renderVoteChoices(estimate.cache.detail.choices, activeVote === null || activeVote === void 0 ? void 0 : activeVote.choice));
    }
    function viewVoteResults(votes) {
        dom.setContent("#story-vote-results", vote.renderVoteResults(system.cache.members, votes));
        dom.setContent("#story-vote-summary", vote.renderVoteSummary(votes));
    }
    // noinspection JSUnusedGlobalSymbols
    function onSubmitVote(choice) {
        const msg = { svc: services.estimate.key, cmd: command.client.submitVote, param: { storyID: estimate.cache.activeStory, choice } };
        socket.send(msg);
    }
    vote.onSubmitVote = onSubmitVote;
    function getVoteResults(votes) {
        const floats = votes.map(v => {
            const n = parseFloat(v.choice);
            if (isNaN(n)) {
                return -1;
            }
            return n;
        }).filter(x => x !== -1).sort();
        const count = floats.length;
        const min = Math.min(...floats);
        const max = Math.max(...floats);
        const sum = floats.reduce((x, y) => x + y, 0);
        const mode = floats.reduce(function (current, item) {
            const val = current.numMapping[item] = (current.numMapping[item] || 0) + 1;
            if (val > current.greatestFreq) {
                current.greatestFreq = val;
                current.mode = item;
            }
            return current;
        }, { mode: null, greatestFreq: -Infinity, numMapping: {} }).mode;
        return {
            count, min, max, sum,
            mean: count === 0 ? 0 : sum / count,
            median: count === 0 ? 0 : floats[Math.floor(floats.length / 2)],
            mode: count === 0 ? 0 : mode
        };
    }
    vote.getVoteResults = getVoteResults;
})(vote || (vote = {}));
var action;
(function (action_1) {
    function renderAction(action) {
        const c = JSON.stringify(action.content, null, 2);
        return JSX("tr", null,
            JSX("td", null, system.getMemberName(action.authorID)),
            JSX("td", null, action.act),
            JSX("td", null, c === "null" ? "" : JSX("pre", null, c)),
            JSX("td", null, action.note),
            JSX("td", { class: "uk-table-shrink uk-text-nowrap" }, date.toDateTimeString(new Date(action.occurred))));
    }
    function renderActions(actions) {
        if (actions.length === 0) {
            return JSX("div", null, "No actions available");
        }
        else {
            return JSX("table", { class: "uk-table uk-table-divider uk-text-left" },
                JSX("thead", null,
                    JSX("tr", null,
                        JSX("th", null, "Author"),
                        JSX("th", null, "Act"),
                        JSX("th", null, "Content"),
                        JSX("th", null, "Note"),
                        JSX("th", null, "Occurred"))),
                JSX("tbody", null, actions.map(a => renderAction(a))));
        }
    }
    action_1.renderActions = renderActions;
})(action || (action = {}));
var contents;
(function (contents_1) {
    function renderSprintContent(svc, session) {
        const profile = system.cache.getProfile();
        return JSX("tr", null,
            JSX("td", null,
                JSX("a", { class: `${profile.linkColor}-fg`, href: `/${svc.key}/${session.slug}` }, session.title)),
            JSX("td", { class: "uk-table-shrink uk-text-nowrap" }, system.getMemberName(session.owner)),
            JSX("td", { class: "uk-table-shrink uk-text-nowrap" }, date.toDateTimeString(new Date(session.created))));
    }
    function toContent(svc, sessions) {
        return sessions.map(s => {
            return { svc: svc, session: s };
        });
    }
    function renderContents(src, tgt, sessions) {
        const contents = toContent(tgt, sessions);
        contents.sort((l, r) => (l.session.created > r.session.created ? -1 : 1));
        if (contents.length === 0) {
            return JSX("div", null, `No ${tgt.plural} in this ${src.key}`);
        }
        else {
            return JSX("table", { class: "uk-table uk-table-divider uk-text-left" },
                JSX("tbody", null, contents.map(a => renderSprintContent(a.svc, a.session))));
        }
    }
    contents_1.renderContents = renderContents;
})(contents || (contents = {}));
var feedback;
(function (feedback) {
    function renderFeedback(model) {
        const profile = system.cache.getProfile();
        const ret = JSX("div", { id: `feedback-${model.id}`, class: "feedback-detail uk-border-rounded section", onclick: `events.openModal('feedback', '${model.id}');` },
            JSX("a", { class: `${profile.linkColor}-fg section-link` }, system.getMemberName(model.authorID)),
            JSX("div", { class: "feedback-content" }, "loading..."));
        if (model.html.length > 0) {
            dom.setHTML(dom.req(".feedback-content", ret), model.html).style.display = "block";
        }
        return ret;
    }
    function renderFeedbackArray(f) {
        var _a;
        if (f.length === 0) {
            return JSX("div", null,
                JSX("button", { class: "uk-button uk-button-default", onclick: "events.openModal('add-feedback');", type: "button" }, "Add Feedback"));
        }
        else {
            const cats = feedback.getFeedbackCategories(f, ((_a = retro.cache.detail) === null || _a === void 0 ? void 0 : _a.categories) || []);
            const profile = system.cache.getProfile();
            return JSX("div", { class: "uk-grid-small uk-grid-match uk-child-width-expand@m uk-grid-divider", "data-uk-grid": true }, cats.map(cat => JSX("div", { class: "feedback-list uk-transition-toggle" },
                JSX("div", { class: "feedback-category-header" },
                    JSX("span", { class: "right" },
                        JSX("a", { class: `${profile.linkColor}-fg uk-icon-button uk-transition-fade`, "data-uk-icon": "plus", onclick: `events.openModal('add-feedback', '${cat.category}');`, title: "edit feedback" })),
                    JSX("span", { class: "feedback-category-title", onclick: `events.openModal('add-feedback', '${cat.category}');` }, cat.category)),
                JSX("div", null, cat.feedback.map(fb => JSX("div", null, renderFeedback(fb)))))));
        }
    }
    feedback.renderFeedbackArray = renderFeedbackArray;
})(feedback || (feedback = {}));
var member;
(function (member_2) {
    function renderMember(member) {
        const profile = system.cache.getProfile();
        return JSX("div", { class: "section", onclick: `events.openModal('member', '${member.userID}');` },
            JSX("div", { title: "user is offline", class: "right uk-article-meta online-indicator" }, "offline"),
            JSX("div", { class: `${profile.linkColor}-fg section-link` }, member.name));
    }
    function renderMembers(members) {
        if (members.length === 0) {
            return JSX("div", null,
                JSX("button", { class: "uk-button uk-button-default", onclick: "events.openModal('invitation');", type: "button" }, "Invite Members"));
        }
        else {
            return JSX("ul", { class: "uk-list uk-list-divider" }, members.map(m => JSX("li", { id: `member-${m.userID}` }, renderMember(m))));
        }
    }
    member_2.renderMembers = renderMembers;
})(member || (member = {}));
var permission;
(function (permission) {
    function renderEmails(key, emails) {
        const cls = `uk-checkbox uk-margin-small-right perm-${key}-email`;
        const oc = `permission.onChanged('email', '${key}', this.checked)`;
        return JSX("ul", null, emails.map(e => {
            return JSX("li", null,
                JSX("label", null,
                    e.matched ? JSX("input", { class: cls, type: "checkbox", value: e.domain, checked: "checked", onchange: oc }) : JSX("input", { class: cls, type: "checkbox", value: e.domain, onchange: oc }),
                    "Using email address ",
                    e.domain));
        }));
    }
    permission.renderEmails = renderEmails;
    function readPermission(k) {
        const checkbox = dom.opt(`#perm-${k}-checkbox`);
        if (!checkbox || !checkbox.checked) {
            return [];
        }
        const emails = dom.els(`.perm-${k}-email`);
        const v = emails.filter(e => e.checked).map(e => e.value).join(",");
        const access = "member";
        return [{ k, v, access }];
    }
    function readPermissions() {
        const ret = [];
        ret.push(...readPermission("team"));
        ret.push(...readPermission("sprint"));
        ret.push(...readPermission("github"));
        ret.push(...readPermission("google"));
        ret.push(...readPermission("slack"));
        ret.push(...readPermission("amazon"));
        ret.push(...readPermission("microsoft"));
        return ret;
    }
    permission.readPermissions = readPermissions;
    function applyPermissions(perms) {
        system.cache.permissions = collection.groupBy(perms, x => x.k);
        dom.setDisplay("#public-link-container", perms === null || perms.length === 0);
        dom.setDisplay("#private-link-container", perms !== null && perms.length > 0);
    }
    permission.applyPermissions = applyPermissions;
})(permission || (permission = {}));
var report;
(function (report) {
    function renderReport(model) {
        const profile = system.cache.getProfile();
        const ret = JSX("div", { id: `report-${model.id}`, class: "report-detail uk-border-rounded section", onclick: `events.openModal('report', '${model.id}');` },
            JSX("a", { class: `${profile.linkColor}-fg section-link` }, system.getMemberName(model.authorID)),
            JSX("div", { class: "report-content" }, "loading..."));
        if (model.html.length > 0) {
            dom.setHTML(dom.req(".report-content", ret), model.html).style.display = "block";
        }
        return ret;
    }
    function renderReports(reports) {
        if (reports.length === 0) {
            return JSX("div", null,
                JSX("button", { class: "uk-button uk-button-default", onclick: "events.openModal('add-report');", type: "button" }, "Add Report"));
        }
        else {
            const dates = report.getReportDates(reports);
            return JSX("ul", { class: "uk-list" }, dates.map(day => JSX("li", { id: `report-date-${day.d}` },
                JSX("h5", null,
                    JSX("div", { class: "right uk-article-meta" }, date.dow(date.dateFromYMD(day.d).getDay())),
                    date.toDateString(date.dateFromYMD(day.d))),
                JSX("ul", null, day.reports.map(r => JSX("li", null, renderReport(r)))))));
        }
    }
    report.renderReports = renderReports;
})(report || (report = {}));
var sprint;
(function (sprint) {
    function renderSprintDates(startDate, endDate) {
        function f(p, d) {
            return JSX("span", null,
                p,
                " ",
                JSX("span", { class: "sprint-date", onclick: "events.openModal('session');" }, d ? date.toDateString(d) : ""));
        }
        const s = f("starts", startDate);
        const e = f("ends", endDate);
        if (startDate) {
            if (endDate) {
                return JSX("span", null,
                    s,
                    ", ",
                    e);
            }
            else {
                return s;
            }
        }
        else {
            if (endDate) {
                return e;
            }
            else {
                return JSX("span", null, "Sprint");
            }
        }
    }
    sprint.renderSprintDates = renderSprintDates;
    function renderSprintLink(spr) {
        const profile = system.cache.getProfile();
        return JSX("span", null,
            JSX("a", { class: `${profile.linkColor}-fg`, href: `/sprint/${spr.slug}` }, spr.title),
            "\u00A0");
    }
    sprint.renderSprintLink = renderSprintLink;
    function renderSprintSelect(sprints, activeID) {
        return JSX("select", { class: "uk-select", onchange: "permission.setModelPerms('sprint')" },
            JSX("option", { value: "" }, "- no sprint -"),
            sprints.map(s => {
                return s.id === activeID ? JSX("option", { selected: "selected", value: s.id }, s.title) : JSX("option", { value: s.id }, s.title);
            }));
    }
    sprint.renderSprintSelect = renderSprintSelect;
})(sprint || (sprint = {}));
var story;
(function (story_2) {
    function renderStory(story) {
        const profile = system.cache.getProfile();
        return JSX("li", { id: `story-${story.id}`, class: "section", onclick: `events.openModal('story', '${story.id}');` },
            JSX("div", { class: "right uk-article-meta story-status" }, story.status),
            JSX("div", { class: `${profile.linkColor}-fg section-link` }, story.title));
    }
    function renderStories(stories) {
        if (stories.length === 0) {
            return JSX("div", { id: "story-list" },
                JSX("button", { class: "uk-button uk-button-default", onclick: "events.openModal('add-story');", type: "button" }, "Add Story"));
        }
        else {
            return JSX("ul", { id: "story-list", class: "uk-list uk-list-divider" }, stories.map(s => renderStory(s)));
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
                return JSX("span", { class: "vote-badge uk-border-rounded" }, status);
        }
    }
    story_2.renderStatus = renderStatus;
    function renderTotal(sum) {
        return JSX("li", { id: "story-total" },
            JSX("div", { class: "right uk-article-meta" },
                JSX("span", { class: "vote-badge uk-border-rounded" }, sum)),
            " Total");
    }
    story_2.renderTotal = renderTotal;
})(story || (story = {}));
var team;
(function (team) {
    function renderTeamLink(tm) {
        const profile = system.cache.getProfile();
        return JSX("span", null,
            " in ",
            JSX("a", { class: `${profile.linkColor}-fg`, href: `/team/${tm.slug}` }, tm.title));
    }
    team.renderTeamLink = renderTeamLink;
    function renderTeamSelect(teams, activeID) {
        return JSX("select", { class: "uk-select", onchange: "permission.setModelPerms('team')" },
            JSX("option", { value: "" }, "- no team -"),
            teams.map(t => {
                return t.id === activeID ? JSX("option", { selected: "selected", value: t.id }, t.title) : JSX("option", { value: t.id }, t.title);
            }));
    }
    team.renderTeamSelect = renderTeamSelect;
})(team || (team = {}));
var vote;
(function (vote_1) {
    function renderVoteMember(member, hasVote) {
        return JSX("div", { class: "vote-member", title: `${member.name} has ${hasVote ? "voted" : "not voted"}` },
            JSX("div", null,
                JSX("span", { "data-uk-icon": `icon: ${hasVote ? "check" : "minus"}; ratio: 1.6` })),
            member.name);
    }
    function renderVoteMembers(members, votes) {
        return JSX("div", { class: "uk-flex uk-flex-wrap uk-flex-around" }, members.map(m => renderVoteMember(m, votes.filter(v => v.userID === m.userID).length > 0)));
    }
    vote_1.renderVoteMembers = renderVoteMembers;
    function renderVoteChoices(choices, choice) {
        return JSX("div", { class: "uk-flex uk-flex-wrap uk-flex-center" }, choices.map(c => JSX("div", { class: `vote-choice uk-border-circle uk-box-shadow-hover-medium${(c === choice ? ` active ${system.cache.getProfile().linkColor}-border` : "")}`, onclick: `vote.onSubmitVote('${c}');` }, c)));
    }
    vote_1.renderVoteChoices = renderVoteChoices;
    function renderVoteResult(member, choice) {
        if (!choice) {
            return JSX("div", { class: "vote-result" },
                JSX("div", null,
                    JSX("span", { class: "uk-border-circle" },
                        JSX("span", { "data-uk-icon": "icon: minus; ratio: 1.6" }))),
                " ",
                member.name);
        }
        return JSX("div", { class: "vote-result" },
            JSX("div", null,
                JSX("span", { class: "uk-border-circle" }, choice)),
            " ",
            member.name);
    }
    function renderVoteResults(members, votes) {
        return JSX("div", { class: "uk-flex uk-flex-wrap uk-flex-around" }, members.map(m => {
            const vote = votes.filter(v => {
                return v.userID === m.userID;
            });
            return renderVoteResult(m, vote.length > 0 ? vote[0].choice : undefined);
        }));
    }
    vote_1.renderVoteResults = renderVoteResults;
    function renderVoteSummary(votes) {
        const results = vote_1.getVoteResults(votes);
        function trim(n) { return n.toString().substr(0, 4); }
        return JSX("div", { class: "uk-flex uk-flex-wrap uk-flex-center result-container" },
            JSX("div", { class: "result" },
                JSX("div", { class: "secondary uk-border-circle" },
                    trim(results.count),
                    " / ",
                    trim(votes.length)),
                " ",
                JSX("div", null, "votes counted")),
            JSX("div", { class: "result" },
                JSX("div", { class: "secondary uk-border-circle" },
                    trim(results.min),
                    "-",
                    trim(results.max)),
                " ",
                JSX("div", null, "vote range")),
            JSX("div", { class: "result mean-result" },
                JSX("div", { class: `mean uk-border-circle ${system.cache.getProfile().linkColor}-border` }, trim(results.mean)),
                " ",
                JSX("div", null, "average")),
            JSX("div", { class: "result" },
                JSX("div", { class: "secondary uk-border-circle" }, trim(results.median)),
                " ",
                JSX("div", null, "median")),
            JSX("div", { class: "result" },
                JSX("div", { class: "secondary uk-border-circle" }, trim(results.mode)),
                " ",
                JSX("div", null, "mode")));
    }
    vote_1.renderVoteSummary = renderVoteSummary;
})(vote || (vote = {}));
