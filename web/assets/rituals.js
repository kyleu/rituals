"use strict";
var rituals;
(function (rituals) {
    function onError(svc, err) {
        console.error(`${svc.key}: ${err}`);
        const idx = err.lastIndexOf(":");
        if (idx > -1) {
            err = err.substr(idx + 1);
        }
        notify.notify(`${svc.key} error: ${err}`, false);
    }
    rituals.onError = onError;
    function init(svc, id) {
        window.onbeforeunload = function () {
            socket.setAppUnloading();
        };
        socket.socketConnect(services.fromKey(svc), id);
    }
    rituals.init = init;
})(rituals || (rituals = {}));
var action;
(function (action) {
    function loadActions() {
        socket.send({ svc: services.system.key, cmd: command.client.getActions, param: null });
    }
    action.loadActions = loadActions;
    function viewActions(actions) {
        dom.setContent("#action-list", action.renderActions(actions));
    }
    action.viewActions = viewActions;
})(action || (action = {}));
var action;
(function (action_1) {
    function renderAction(action) {
        const c = JSON.stringify(action.content, null, 2);
        return JSX("tr", null,
            JSX("td", null, member.renderTitle(member.getMember(action.userID))),
            JSX("td", null, action.act),
            JSX("td", null, c === "null" ? "" : JSX("pre", null, c)),
            JSX("td", null, action.note),
            JSX("td", { class: "uk-table-shrink uk-text-nowrap" }, date.toDateTimeString(new Date(action.created))));
    }
    function renderActions(actions) {
        if (actions.length === 0) {
            return JSX("div", null, "No actions available");
        }
        else {
            return JSX("table", { class: "uk-table uk-table-divider uk-text-left" },
                JSX("thead", null,
                    JSX("tr", null,
                        JSX("th", null, "User"),
                        JSX("th", null, "Act"),
                        JSX("th", null, "Content"),
                        JSX("th", null, "Note"),
                        JSX("th", null, "Created"))),
                JSX("tbody", null, actions.map(a => renderAction(a))));
        }
    }
    action_1.renderActions = renderActions;
})(action || (action = {}));
var auth;
(function (auth) {
    const github = { key: "github", title: "GitHub" };
    const google = { key: "google", title: "Google" };
    const slack = { key: "slack", title: "Slack" };
    const facebook = { key: "facebook", title: "Facebook" };
    const amazon = { key: "amazon", title: "Amazon" };
    const microsoft = { key: "microsoft", title: "Microsoft" };
    auth.allProviders = [github, google, slack, facebook, amazon, microsoft];
    let auths = [];
    function applyAuths(as) {
        auths = as;
    }
    auth.applyAuths = applyAuths;
    function active() {
        if (!auths) {
            return [];
        }
        return auths;
    }
    auth.active = active;
})(auth || (auth = {}));
var comment;
(function (comment) {
    let activeComments = [];
    let activeType;
    let activeID;
    function applyComments(comments) {
        activeComments = comments;
    }
    comment.applyComments = applyComments;
    function setActive(t, id) {
        activeType = t;
        activeID = id;
    }
    comment.setActive = setActive;
    function show(t) {
        modal.open("comment", t);
    }
    comment.show = show;
    function add(t) {
        const textarea = dom.req(`#comment-add-content-${t}`);
        const v = textarea.value;
        textarea.value = "";
        const param = { targetType: activeType, targetID: activeID, content: v };
        socket.send({ svc: services.system.key, cmd: command.client.addComment, param: param });
    }
    comment.add = add;
    function onCommentUpdate(u) {
        activeComments.push(u);
        setCounts();
        load();
    }
    comment.onCommentUpdate = onCommentUpdate;
    function find(t, id) {
        if ((!t) || t === "modal") {
            t = activeType;
            if (!id) {
                id = activeID;
            }
        }
        if (t === "root") {
            t = "";
        }
        if (id) {
            return activeComments.filter(x => x.targetType === t && x.targetID == id);
        }
        return activeComments.filter(x => x.targetType === t);
    }
    function load(t, id) {
        if ((!t) || t === "modal") {
            t = activeType;
            if (!id) {
                id = activeID;
            }
        }
        if (!t) {
            console.warn(`invalid comment type [${t}]`);
            return;
        }
        activeType = t;
        activeID = id;
        const comments = find(t, id);
        if (t !== "root") {
            t = "modal";
        }
        const el = dom.req(`#drop-comment-${t} .uk-comment-list`);
        dom.setContent(el, comment.renderComments(comments, system.cache.getProfile()));
        el.scrollTop = el.scrollHeight;
    }
    comment.load = load;
    function setCounts() {
        const containers = dom.els(`.comment-count-container`);
        let matchedModal = false;
        const modalCount = dom.opt(`#comment-count-modal`);
        containers.forEach(cc => {
            const t = cc.dataset["commentType"];
            const id = cc.dataset["commentId"];
            if (!t) {
                throw `invalid comment type [${t}] with id [${id}]`;
            }
            let comments = find(t, id);
            setCount(t, comments, cc);
            if (activeType === t) {
                if (modalCount) {
                    setCount(t, comments, modalCount);
                    matchedModal = true;
                }
            }
        });
        if (!matchedModal && modalCount) {
            setCount("modal", [], modalCount, true);
        }
    }
    comment.setCounts = setCounts;
    function closeModal() {
        if (activeType === "story") {
            modal.openSoon("story");
        }
        else if (activeType === "report") {
            modal.openSoon("report");
        }
        else if (activeType === "feedback") {
            modal.openSoon("feedback");
        }
        activeType = undefined;
    }
    comment.closeModal = closeModal;
    function remove(id) {
        if (confirm("remove this comment?")) {
            socket.send({ svc: services.system.key, cmd: command.client.removeComment, param: id });
        }
        return false;
    }
    comment.remove = remove;
    function setCount(t, comments, cc, force) {
        dom.req(".text", cc).innerText = comments.length.toString();
        if (t !== "root" && t !== "modal" && t.length !== 0) {
            dom.setDisplay(cc, (comments.length !== 0) || force === true);
        }
    }
    function onCommentRemoved(id) {
        activeComments = activeComments.filter((c) => c.id !== id);
        setCounts();
        load();
    }
    comment.onCommentRemoved = onCommentRemoved;
})(comment || (comment = {}));
var comment;
(function (comment_1) {
    function renderComment(comment, profile) {
        let close = comment.userID === profile.userID ? JSX("div", { class: "right" },
            JSX("a", { class: `${profile.linkColor}-fg`, "data-uk-icon": "close", href: "", onclick: `return comment.remove('${comment.id}');`, title: "remove your comment" })) : JSX("span", null);
        return JSX("li", null,
            JSX("article", { class: "uk-comment uk-visible-toggle uk-transition-toggle", tabindex: "-1" },
                member.renderHeader(member.getMember(comment.userID), comment.created, close),
                JSX("div", { class: "uk-comment-body" },
                    JSX("div", { dangerouslySetInnerHTML: { __html: comment.html } })),
                JSX("hr", null)));
    }
    function renderComments(comments, profile) {
        if (comments.length === 0) {
            return JSX("div", null, "No comments available");
        }
        else {
            return JSX("div", null, comments.map(c => renderComment(c, profile)));
        }
    }
    comment_1.renderComments = renderComments;
    function renderCount(k, v) {
        const profile = system.cache.getProfile();
        return JSX("div", { class: "comment-count-container uk-margin-small-left left hidden", "data-comment-type": k, "data-comment-id": v },
            JSX("a", { class: `${profile.linkColor}-fg`, title: "view comments" },
                JSX("div", { class: "comment-count" },
                    JSX("span", { class: "uk-icon", "data-uk-icon": "comment" }),
                    JSX("span", { class: "text" }))));
    }
    comment_1.renderCount = renderCount;
})(comment || (comment = {}));
var contents;
(function (contents) {
    function onContentDisplay(key, same, content, html) {
        dom.setDisplay(`#modal-${key} .content-edit`, same);
        dom.setDisplay(`#modal-${key} .buttons-edit`, same);
        const v = dom.setDisplay(`#modal-${key} .content-view`, !same);
        dom.setHTML(v, same ? "" : html);
        dom.setDisplay(`#modal-${key} .buttons-view`, !same);
        const contentEditTextarea = dom.req(`#${key}-edit-content`);
        dom.setValue(contentEditTextarea, same ? content : "");
        if (same) {
            dom.wireTextarea(contentEditTextarea);
        }
    }
    contents.onContentDisplay = onContentDisplay;
})(contents || (contents = {}));
var contents;
(function (contents_1) {
    function renderSprintContent(svc, session) {
        const profile = system.cache.getProfile();
        return JSX("tr", null,
            JSX("td", null,
                JSX("a", { class: `${profile.linkColor}-fg`, href: `/${svc.key}/${session.slug}` }, session.title)),
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
                rituals.onError(services.estimate, param);
                break;
            case command.server.sessionJoined:
                const sj = param;
                session.onSessionJoin(sj);
                setEstimateDetail(sj.session);
                story.setStories(sj.stories);
                vote.setVotes(sj.votes);
                session.showWelcomeMessage(sj.members.length);
                break;
            case command.server.sessionUpdate:
                setEstimateDetail(param);
                break;
            case command.server.sessionRemove:
                system.onSessionRemove(services.estimate);
                break;
            case command.server.permissionsUpdate:
                system.setPermissions(param);
                break;
            case command.server.teamUpdate:
                const tm = param;
                if (estimate.cache.detail) {
                    estimate.cache.detail.teamID = tm === null || tm === void 0 ? void 0 : tm.id;
                }
                session.setTeam(tm);
                break;
            case command.server.sprintUpdate:
                const spr = param;
                if (estimate.cache.detail) {
                    estimate.cache.detail.sprintID = spr === null || spr === void 0 ? void 0 : spr.id;
                }
                session.setSprint(spr);
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
        session.setDetail(detail);
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
        x.sort(s => s.idx);
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
            modal.hide("story");
        }
        notify.notify("story has been deleted", true);
    }
    estimate.onStoryRemove = onStoryRemove;
    function preUpdate(id) {
        return estimate.cache.stories.filter((p) => p.id !== id);
    }
})(estimate || (estimate = {}));
var feedback;
(function (feedback_1) {
    function setFeedback(feedback) {
        retro.cache.feedback = feedback;
        dom.setContent("#feedback-detail", feedback_1.renderFeedbackArray(feedback));
        comment.setCounts();
        modal.hide("add-feedback");
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
                socket.send({ svc: services.retro.key, cmd: command.client.removeFeedback, param: id });
                modal.hide("feedback");
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
    function viewActiveFeedback(id) {
        var _a;
        const profile = system.cache.getProfile();
        if (id) {
            retro.cache.activeFeedback = id;
        }
        const fb = getActiveFeedback();
        if (!fb) {
            console.warn("no active feedback");
            return;
        }
        const same = fb.userID === profile.userID;
        dom.setText("#feedback-title", `${fb.category} / ${(_a = member.getMember(fb.userID)) === null || _a === void 0 ? void 0 : _a.name}`);
        dom.setSelectOption("#feedback-edit-category", same ? fb.category : undefined);
        contents.onContentDisplay("feedback", same, fb.content, fb.html);
        comment.setActive("feedback", fb.id);
        comment.setCounts();
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
        notify.notify("feedback has been deleted", true);
    }
    feedback_1.onFeedbackRemoved = onFeedbackRemoved;
    function preUpdate(id) {
        return retro.cache.feedback.filter((p) => p.id !== id);
    }
    function postUpdate(x, id) {
        feedback.setFeedback(x);
        if (id === retro.cache.activeFeedback) {
            modal.hide("feedback");
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
    function viewAddFeedback(p) {
        dom.setSelectOption("#feedback-category", p);
        const feedbackContent = dom.setValue("#feedback-content", "");
        dom.wireTextarea(feedbackContent);
        setTimeout((() => feedbackContent.focus()), 250);
    }
    feedback_1.viewAddFeedback = viewAddFeedback;
    function viewFeedback(p) {
        feedback.viewActiveFeedback(p);
        const feedbackEditContent = dom.req("#feedback-edit-content");
        setTimeout(() => {
            dom.wireTextarea(feedbackEditContent);
            feedbackEditContent.focus();
        }, 250);
    }
    feedback_1.viewFeedback = viewFeedback;
})(feedback || (feedback = {}));
var feedback;
(function (feedback) {
    function renderFeedback(model) {
        const profile = system.cache.getProfile();
        const ret = JSX("div", { id: `feedback-${model.id}`, class: "feedback-detail section", onclick: `modal.open('feedback', '${model.id}');` },
            JSX("div", { class: "feedback-comments right" }, comment.renderCount("feedback", model.id)),
            JSX("div", { class: "left" },
                JSX("a", { class: `${profile.linkColor}-fg section-link` }, member.renderTitle(member.getMember(model.userID)))),
            JSX("div", { class: "clear" }),
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
                JSX("button", { class: "uk-button uk-button-default", onclick: "modal.open('add-feedback');", type: "button" }, "Add Feedback"));
        }
        else {
            const cats = feedback.getFeedbackCategories(f, ((_a = retro.cache.detail) === null || _a === void 0 ? void 0 : _a.categories) || []);
            const profile = system.cache.getProfile();
            return JSX("div", { class: "uk-grid-small uk-grid-match uk-child-width-expand@m uk-grid-divider", "data-uk-grid": true }, cats.map(cat => JSX("div", { class: "feedback-list uk-transition-toggle" },
                JSX("div", { class: "feedback-category-header" },
                    JSX("span", { class: "right" },
                        JSX("a", { class: `${profile.linkColor}-fg uk-icon-button uk-transition-fade`, "data-uk-icon": "plus", onclick: `modal.open('add-feedback', '${cat.category}');`, title: "add feedback" })),
                    JSX("span", { class: "feedback-category-title", onclick: `modal.open('add-feedback', '${cat.category}');` }, cat.category)),
                JSX("div", null, cat.feedback.map(fb => JSX("div", null, renderFeedback(fb)))))));
        }
    }
    feedback.renderFeedbackArray = renderFeedbackArray;
})(feedback || (feedback = {}));
var member;
(function (member_1) {
    function memberUpdateDom(nameChanged) {
        if (nameChanged) {
            switch (system.cache.currentService) {
                case services.team:
                    break;
                case services.sprint:
                    break;
                case services.estimate:
                    if (estimate.cache.activeStory) {
                        vote.viewVotes();
                    }
                    break;
                case services.standup:
                    dom.setContent("#report-detail", report.renderReports(standup.cache.reports));
                    if (standup.cache.activeReport) {
                        report.viewActiveReport();
                    }
                    break;
                case services.retro:
                    dom.setContent("#feedback-detail", feedback.renderFeedbackArray(retro.cache.feedback));
                    if (retro.cache.activeFeedback) {
                        feedback.viewActiveFeedback();
                    }
                    break;
            }
        }
    }
    member_1.memberUpdateDom = memberUpdateDom;
    function activeMemberDom(member) {
        const owner = member_1.selfCanEdit();
        dom.setDisplay("#modal-member .owner-form", owner);
        dom.setDisplay("#modal-member .member-form", !owner);
        dom.setDisplay("#modal-member .owner-actions", owner);
        dom.setDisplay("#modal-member .member-actions", !owner);
        dom.setSelectOption("#member-modal-role-select", member.role);
        dom.setText("#member-modal-name", member.name);
        dom.setText("#member-modal-role", member.role);
    }
    member_1.activeMemberDom = activeMemberDom;
})(member || (member = {}));
var member;
(function (member_2) {
    let members = [];
    let activeMember;
    function getMember(id) {
        return members.filter(m => m.userID === id).shift();
    }
    member_2.getMember = getMember;
    function getMembers() {
        return members;
    }
    member_2.getMembers = getMembers;
    function setMembers() {
        member_2.updateSelf(members.filter(member_2.isSelf).shift());
        const others = members.filter(x => !member_2.isSelf(x));
        dom.setContent("#member-detail", member_2.renderMembers(others));
        if (others.length > 0) {
            modal.hide('welcome');
        }
        member_2.renderOnline();
    }
    member_2.setMembers = setMembers;
    function onMemberUpdate(member) {
        if (member_2.isSelf(member)) {
            modal.hide("self");
        }
        else {
            modal.hide("member");
        }
        const unfiltered = members;
        const curr = unfiltered.filter(m => m.userID === member.userID).shift();
        const nameChanged = (curr === null || curr === void 0 ? void 0 : curr.name) !== member.name;
        const ms = unfiltered.filter(m => m.userID !== member.userID);
        if (ms.length === members.length) {
            notify.notify(`${member.name} has joined`, true);
        }
        ms.push(member);
        ms.sort((l, r) => (l.name > r.name) ? 1 : -1);
        members = ms;
        setMembers();
        member_2.memberUpdateDom(nameChanged);
    }
    member_2.onMemberUpdate = onMemberUpdate;
    function onMemberRemove(member) {
        var _a;
        if (member === system.cache.getProfile().userID) {
            notify.notify(`you have left this ${(_a = system.cache.currentService) === null || _a === void 0 ? void 0 : _a.key}`, true);
            document.location.href = "/";
        }
        else {
            modal.hide("member");
            const unfiltered = members;
            const ms = unfiltered.filter(m => m.userID !== member);
            ms.sort((l, r) => (l.name > r.name) ? 1 : -1);
            members = ms;
            setMembers();
        }
    }
    member_2.onMemberRemove = onMemberRemove;
    function viewActiveMember(p) {
        if (p) {
            activeMember = p;
        }
        const member = getActiveMember();
        if (!member) {
            return;
        }
        member_2.activeMemberDom(member);
    }
    member_2.viewActiveMember = viewActiveMember;
    function removeMember(id = activeMember) {
        if (!id) {
            console.warn(`cannot load active member [${activeMember}]`);
        }
        if (id === "self") {
            id = system.cache.getProfile().userID;
        }
        const svc = system.cache.currentService;
        if (confirm(`Are you sure you wish to leave this ${svc.key}?`)) {
            const msg = { svc: svc.key, cmd: command.client.removeMember, param: id };
            socket.send(msg);
        }
    }
    member_2.removeMember = removeMember;
    function saveRole() {
        const curr = getActiveMember();
        if (!curr) {
            console.warn("no active member");
            return;
        }
        const src = curr.role;
        const tgt = dom.req("#member-modal-role-select").value;
        if (src === tgt) {
            modal.hide("member");
        }
        else {
            const svc = system.cache.currentService;
            const msg = { svc: svc.key, cmd: command.client.updateMember, param: { id: curr.userID, role: tgt } };
            socket.send(msg);
        }
    }
    member_2.saveRole = saveRole;
    function applyMembers(m) {
        members = m;
    }
    member_2.applyMembers = applyMembers;
    function getActiveMember() {
        if (!activeMember) {
            console.warn("no active member");
            return undefined;
        }
        const curr = members.filter(x => x.userID === activeMember).shift();
        if (!curr) {
            console.warn(`cannot load active member [${activeMember}]`);
        }
        return curr;
    }
})(member || (member = {}));
var member;
(function (member_3) {
    function renderMember(member) {
        const profile = system.cache.getProfile();
        return JSX("div", { class: "section", onclick: `modal.open('member', '${member.userID}');` },
            JSX("div", { title: "user is offline", class: "right uk-article-meta online-indicator" }, "offline"),
            JSX("div", { class: `${profile.linkColor}-fg section-link` }, renderTitle(member)));
    }
    function renderMembers(members) {
        if (members.length === 0) {
            return JSX("div", null,
                JSX("button", { class: "uk-button uk-button-default", onclick: "modal.open('invitation');", type: "button" }, "Invite Members"));
        }
        else {
            return JSX("ul", { class: "uk-list uk-list-divider" }, members.map(m => JSX("li", { id: `member-${m.userID}` }, renderMember(m))));
        }
    }
    member_3.renderMembers = renderMembers;
    function renderTitle(member) {
        if (!member) {
            return JSX("span", null, "{former member}");
        }
        if (member.picture && member.picture.length > 0 && member.picture != "none") {
            return JSX("div", null,
                JSX("div", { class: "profile-image" },
                    JSX("img", { class: "uk-border-circle", src: member.picture, alt: member.name })),
                JSX("div", { class: "left" }, member.name),
                JSX("div", { class: "clear" }));
        }
        return JSX("div", null,
            JSX("div", { class: "profile-image" },
                JSX("span", { class: "profile-icon uk-icon", "data-uk-icon": "user" })),
            JSX("div", { class: "left" }, member.name));
    }
    member_3.renderTitle = renderTitle;
    function renderHeader(m, t, close) {
        return JSX("header", { class: "uk-comment-header uk-position-relative" },
            close ? close : JSX("span", null),
            JSX("div", { class: "uk-grid-collapse uk-flex-middle", "uk-grid": true },
                JSX("div", { class: "uk-width-auto" }, (m && m.picture && m.picture.length > 0) ? JSX("div", null,
                    JSX("div", { class: "profile-image" },
                        JSX("img", { class: "uk-border-circle", src: m.picture, alt: m.name }))) : JSX("div", null,
                    JSX("div", { class: "profile-image" },
                        JSX("span", { class: "profile-icon uk-icon", "data-uk-icon": "user" })))),
                JSX("div", { class: "uk-width-expand" },
                    JSX("h4", { class: "uk-comment-title uk-margin-remove" }, m === null || m === void 0 ? void 0 : m.name),
                    JSX("p", { class: "uk-comment-meta uk-margin-remove-top" }, date.toDateTimeString(new Date(t))))));
    }
    member_3.renderHeader = renderHeader;
    function viewSelf() {
        const selfInput = dom.setValue("#self-name-input", dom.req("#member-self .member-name").innerText);
        setTimeout(() => selfInput.focus(), 250);
    }
    member_3.viewSelf = viewSelf;
    function setPicture(url) {
        if (url && url.length > 0 && url != "none") {
            return JSX("div", { class: "model-icon profile-image" },
                JSX("img", { class: "uk-border-circle", src: url, alt: "your picture" }));
        }
        return JSX("span", { class: "model-icon h3-icon", onclick: "modal.open('self');", "data-uk-icon": "icon: user;" });
    }
    member_3.setPicture = setPicture;
})(member || (member = {}));
var member;
(function (member_4) {
    let online = [];
    function onOnlineUpdate(update) {
        if (update.connected) {
            if (!online.find(x => x === update.userID)) {
                online.push(update.userID);
            }
        }
        else {
            online = online.filter(x => x !== update.userID);
        }
        renderOnline();
    }
    member_4.onOnlineUpdate = onOnlineUpdate;
    function applyOnline(o) {
        online = o;
    }
    member_4.applyOnline = applyOnline;
    function renderOnline() {
        for (const member of member_4.getMembers()) {
            const el = dom.opt(`#member-${member.userID} .online-indicator`);
            if (el) {
                if (!online.find(x => x === member.userID)) {
                    el.classList.add("offline");
                }
                else {
                    el.classList.remove("offline");
                }
            }
        }
    }
    member_4.renderOnline = renderOnline;
    function canEdit(m) {
        return m.role == "owner";
    }
})(member || (member = {}));
var member;
(function (member) {
    let me;
    function isSelf(x) {
        return x.userID === system.cache.getProfile().userID;
    }
    member.isSelf = isSelf;
    function selfCanEdit() {
        return me !== undefined && canEdit(me);
    }
    member.selfCanEdit = selfCanEdit;
    function updateSelf(self) {
        if (self) {
            me = self;
            dom.setContent("#self-picture", member.setPicture(self.picture));
            dom.setText("#member-self .member-name", self.name);
            dom.setValue("#self-name-input", self.name);
            dom.setText("#member-self .member-role", self.role);
            const e = canEdit(self);
            dom.setDisplay("#history-container", e);
            dom.setDisplay("#session-edit-section", e);
            dom.setDisplay("#session-view-section", !e);
        }
    }
    member.updateSelf = updateSelf;
    function onSubmitSelf() {
        const name = dom.req("#self-name-input").value;
        const choice = dom.req("#self-name-choice-global").checked ? "global" : "local";
        const picture = dom.req("#self-picture-input").value;
        const msg = { svc: services.system.key, cmd: command.client.updateProfile, param: { name, choice, picture } };
        socket.send(msg);
    }
    member.onSubmitSelf = onSubmitSelf;
    function canEdit(m) {
        return m.role == "owner";
    }
})(member || (member = {}));
var command;
(function (command) {
    command.client = {
        ping: "ping",
        connect: "connect",
        getActions: "get-actions",
        getTeams: "get-teams",
        getSprints: "get-sprints",
        updateSession: "update-session",
        addComment: "add-comment",
        updateComment: "update-comment",
        removeComment: "remove-comment",
        updateProfile: "update-profile",
        updateMember: "update-member",
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
        sessionRemove: "session-remove",
        commentUpdate: "comment-update",
        commentRemove: "comment-remove",
        permissionsUpdate: "permissions-update",
        teamUpdate: "team-update",
        sprintUpdate: "sprint-update",
        contentUpdate: "content-update",
        actions: "actions",
        teams: "teams",
        sprints: "sprints",
        memberUpdate: "member-update",
        memberRemove: "member-remove",
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
var services;
(function (services) {
    services.system = { key: "system", title: "System", plural: "systems", icon: "close" };
    services.team = { key: "team", title: "Team", plural: "teams", icon: "users" };
    services.sprint = { key: "sprint", title: "Sprint", plural: "sprints", icon: "git-fork" };
    services.estimate = { key: "estimate", title: "Estimate Session", plural: "estimates", icon: "settings" };
    services.standup = { key: "standup", title: "Daily Standup", plural: "standups", icon: "future" };
    services.retro = { key: "retro", title: "Retrospective", plural: "retros", icon: "history" };
    const allServices = [services.system, services.team, services.sprint, services.estimate, services.standup, services.retro];
    function fromKey(key) {
        const ret = allServices.find(s => s.key === key);
        if (!ret) {
            throw `invalid service [${key}]`;
        }
        return ret;
    }
    services.fromKey = fromKey;
})(services || (services = {}));
var permission;
(function (permission) {
    let permissions = [];
    function setPerms() {
        dom.setDisplay("#public-link-container", permissions === null || permissions.length === 0);
        dom.setDisplay("#private-link-container", permissions !== null && permissions.length > 0);
        ["team", "sprint"].forEach(setModelPerms);
        auth.allProviders.forEach(setProviderPerms);
    }
    permission.setPerms = setPerms;
    function setModelPerms(key) {
        const el = dom.opt(`#model-${key}-select select`);
        if (el) {
            const perms = collection.findGroup(permissions, key);
            const section = dom.opt(`#perm-${key}-section`);
            if (section) {
                const checkbox = dom.req(`#perm-${key}-checkbox`);
                checkbox.checked = perms.length > 0;
                dom.setDisplay(section, el.value.length !== 0);
            }
            collection.findGroup(permissions, key);
        }
    }
    permission.setModelPerms = setModelPerms;
    // noinspection JSUnusedGlobalSymbols
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
        const perms = collection.findGroup(permissions, p.key);
        const auths = auth.active().filter(a => a.provider === p.key);
        const section = dom.opt(`#perm-${p.key}-section`);
        if (section) {
            const checkbox = dom.req(`#perm-${p.key}-checkbox`);
            checkbox.checked = perms.length > 0;
            const emailContainer = dom.req(`#perm-${p.key}-email-container`);
            const emails = collection.flatten(perms.map(x => x.v.split(",").filter(x => x.length > 0))).map(x => ({ matched: true, domain: x }));
            const additional = auths.filter(a => emails.filter(e => a.email.endsWith(e.domain)).length === 0).map(m => {
                return { matched: false, domain: getDomain(m.email) };
            });
            emails.push(...additional);
            emails.sort();
            dom.setDisplay(emailContainer, emails.length > 0);
            dom.setContent(emailContainer, emails.length === 0 ? document.createElement("span") : permission.renderEmails(p.key, emails));
        }
    }
    function getDomain(email) {
        const idx = email.lastIndexOf("@");
        if (idx === -1) {
            return email;
        }
        return email.substr(idx);
    }
    function applyPermissions(perms) {
        permissions = collection.groupBy(perms, x => x.k).groups;
    }
    permission.applyPermissions = applyPermissions;
})(permission || (permission = {}));
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
        ret.push(...readPermission("facebook"));
        ret.push(...readPermission("amazon"));
        ret.push(...readPermission("microsoft"));
        return ret;
    }
    permission.readPermissions = readPermissions;
})(permission || (permission = {}));
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
                modal.hide("report");
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
    function viewActiveReport(id) {
        var _a;
        const profile = system.cache.getProfile();
        if (id) {
            standup.cache.activeReport = id;
        }
        const report = getActiveReport();
        if (!report) {
            console.warn("no active report");
            return;
        }
        dom.setText("#report-title", `${date.toDateString(date.dateFromYMD(report.d))} - ${(_a = member.getMember(report.userID)) === null || _a === void 0 ? void 0 : _a.name}`);
        setFor(report, profile.userID);
    }
    report_1.viewActiveReport = viewActiveReport;
    function setReports(reports) {
        standup.cache.reports = reports;
        dom.setContent("#report-detail", report_1.renderReports(reports));
        comment.setCounts();
        modal.hide("add-report");
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
        const same = report.userID === userID;
        dom.setValue(dom.req("#report-edit-date"), same ? report.d : "");
        contents.onContentDisplay("report", same, report.content, report.html);
        comment.setActive("report", report.id);
        comment.setCounts();
    }
    function viewAddReport() {
        dom.setValue("#report-date", date.dateToYMD(new Date()));
        const reportContent = dom.setValue("#report-content", "");
        dom.wireTextarea(reportContent);
        setTimeout(() => reportContent.focus(), 250);
    }
    report_1.viewAddReport = viewAddReport;
    function viewReport(p) {
        report.viewActiveReport(p);
        const reportEditContent = dom.req("#report-edit-content");
        setTimeout(() => {
            dom.wireTextarea(reportEditContent);
            reportEditContent.focus();
        }, 250);
    }
    report_1.viewReport = viewReport;
})(report || (report = {}));
var report;
(function (report) {
    function renderReport(model) {
        const profile = system.cache.getProfile();
        const ret = JSX("div", { id: `report-${model.id}`, class: "report-detail section", onclick: `modal.open('report', '${model.id}');` },
            JSX("div", { class: "report-comments right" }, comment.renderCount("report", model.id)),
            JSX("div", { class: "left" },
                JSX("a", { class: `${profile.linkColor}-fg section-link` }, member.renderTitle(member.getMember(model.userID)))),
            JSX("div", { class: "clear" }),
            JSX("div", { class: "report-content" }, "loading..."));
        if (model.html.length > 0) {
            dom.setHTML(dom.req(".report-content", ret), model.html).style.display = "block";
        }
        return ret;
    }
    function renderReports(reports) {
        if (reports.length === 0) {
            return JSX("div", null,
                JSX("button", { class: "uk-button uk-button-default", onclick: "modal.open('add-report');", type: "button" }, "Add Report"));
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
                rituals.onError(services.retro, param);
                break;
            case command.server.sessionJoined:
                const sj = param;
                session.onSessionJoin(sj);
                setRetroDetail(sj.session);
                feedback.setFeedback(sj.feedback);
                session.showWelcomeMessage(sj.members.length);
                break;
            case command.server.sessionUpdate:
                setRetroDetail(param);
                break;
            case command.server.sessionRemove:
                system.onSessionRemove(services.retro);
                break;
            case command.server.permissionsUpdate:
                system.setPermissions(param);
                break;
            case command.server.teamUpdate:
                const tm = param;
                if (retro.cache.detail) {
                    retro.cache.detail.teamID = tm === null || tm === void 0 ? void 0 : tm.id;
                }
                session.setTeam(tm);
                break;
            case command.server.sprintUpdate:
                const spr = param;
                if (retro.cache.detail) {
                    retro.cache.detail.sprintID = spr === null || spr === void 0 ? void 0 : spr.id;
                }
                session.setSprint(spr);
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
        session.setDetail(detail);
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
var session;
(function (session_1) {
    function setDetail(session) {
        var _a;
        const oldSlug = ((_a = system.cache.session) === null || _a === void 0 ? void 0 : _a.slug) || "invalid";
        system.cache.session = session;
        document.title = session.title;
        dom.setText("#model-title", session.title);
        dom.setValue("#model-title-input", session.title);
        const items = dom.els("#navbar .uk-navbar-item");
        if (items.length > 0) {
            items[items.length - 1].innerText = session.title;
        }
        if (oldSlug !== session.slug) {
            window.history.replaceState(null, "", document.location.href.replace(oldSlug, session.slug));
            console.log("slugChanged!!!!!");
        }
        modal.hide("session");
    }
    session_1.setDetail = setDetail;
    function onSessionJoin(param) {
        system.cache.apply(param);
        permission.setPerms();
        member.setMembers();
        comment.setCounts();
    }
    session_1.onSessionJoin = onSessionJoin;
    function showWelcomeMessage(memberCount) {
        if (memberCount === 1) {
            setTimeout(() => modal.open("welcome"), 300);
        }
    }
    session_1.showWelcomeMessage = showWelcomeMessage;
    function setSprint(spr) {
        modal.hide("session");
        const lc = dom.req("#sprint-link-container");
        lc.innerHTML = "";
        if (spr) {
            lc.appendChild(sprint.renderSprintLink(spr));
            dom.req("#sprint-warning-name").innerText = spr.title;
        }
    }
    session_1.setSprint = setSprint;
    function setTeam(tm) {
        modal.hide("session");
        const container = dom.req("#team-link-container");
        container.innerHTML = "";
        if (tm) {
            container.appendChild(team.renderTeamLink(tm));
        }
    }
    session_1.setTeam = setTeam;
    function onModalOpen(param) {
        const sessionInput = dom.setValue("#model-title-input", dom.req("#model-title").innerText);
        setTimeout(() => sessionInput.focus(), 250);
        team.refreshTeams();
        sprint.refreshSprints();
    }
    session_1.onModalOpen = onModalOpen;
})(session || (session = {}));
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
            send({ svc: svc.key, cmd: command.client.connect, param: id });
        };
        socket.onmessage = function (event) {
            const msg = JSON.parse(event.data);
            onSocketMessage(msg);
        };
        socket.onerror = function (event) {
            rituals.onError(services.system, event.type);
        };
        socket.onclose = function () {
            onSocketClose();
        };
    }
    socket_1.socketConnect = socketConnect;
    function send(msg) {
        if (debug) {
            console.debug("out", msg);
        }
        socket.send(JSON.stringify(msg));
    }
    socket_1.send = send;
    function onSocketMessage(msg) {
        if (debug) {
            console.debug("in", msg);
        }
        switch (msg.svc) {
            case services.system.key:
                system.onSystemMessage(msg.cmd, msg.param);
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
    socket_1.onSocketMessage = onSocketMessage;
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
                rituals.onError(services.sprint, param);
                break;
            case command.server.sessionJoined:
                const sj = param;
                session.onSessionJoin(sj);
                setSprintDetail(sj.session);
                setSprintContents(sj);
                session.showWelcomeMessage(sj.members.length);
                break;
            case command.server.teamUpdate:
                const tm = param;
                if (sprint.cache.detail) {
                    sprint.cache.detail.teamID = tm === null || tm === void 0 ? void 0 : tm.id;
                }
                session.setTeam(tm);
                break;
            case command.server.sessionUpdate:
                setSprintDetail(param);
                break;
            case command.server.sessionRemove:
                system.onSessionRemove(services.sprint);
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
        sprint.cache.detail = detail;
        const s = detail.startDate ? date.utcDate(detail.startDate) : undefined;
        const e = detail.endDate ? date.utcDate(detail.endDate) : undefined;
        dom.setContent("#sprint-date-display", sprint.renderSprintDates(s, e));
        dom.setValue("#sprint-start-date-input", s ? date.dateToYMD(s) : "");
        dom.setValue("#sprint-end-date-input", e ? date.dateToYMD(e) : "");
        session.setDetail(detail);
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
        const startDate = (_a = dom.opt("#sprint-start-date-input")) === null || _a === void 0 ? void 0 : _a.value;
        const endDate = (_b = dom.opt("#sprint-end-date-input")) === null || _b === void 0 ? void 0 : _b.value;
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
            // dom.setDisplay(c, sprints.length > 0)
            dom.setContent("#model-sprint-select", sprint.renderSprintSelect(sprints, (_a = system.cache.session) === null || _a === void 0 ? void 0 : _a.sprintID));
            permission.setModelPerms("sprint");
        }
    }
    sprint.viewSprints = viewSprints;
})(sprint || (sprint = {}));
var sprint;
(function (sprint) {
    function renderSprintDates(startDate, endDate) {
        function f(p, d) {
            return JSX("span", null,
                p,
                " ",
                JSX("span", { class: "sprint-date", onclick: "modal.open('session');" }, d ? date.toDateString(d) : ""));
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
                rituals.onError(services.standup, param);
                break;
            case command.server.sessionJoined:
                const sj = param;
                session.onSessionJoin(sj);
                setStandupDetail(sj.session);
                report.setReports(sj.reports);
                session.showWelcomeMessage(sj.members.length);
                break;
            case command.server.sessionUpdate:
                setStandupDetail(param);
                break;
            case command.server.sessionRemove:
                system.onSessionRemove(services.standup);
                break;
            case command.server.permissionsUpdate:
                system.setPermissions(param);
                break;
            case command.server.teamUpdate:
                const tm = param;
                if (standup.cache.detail) {
                    standup.cache.detail.teamID = tm === null || tm === void 0 ? void 0 : tm.id;
                }
                session.setTeam(tm);
                break;
            case command.server.sprintUpdate:
                const x = param;
                if (standup.cache.detail) {
                    standup.cache.detail.sprintID = x === null || x === void 0 ? void 0 : x.id;
                }
                session.setSprint(x);
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
        session.setDetail(detail);
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
        notify.notify("report has been deleted", true);
    }
    function preUpdate(id) {
        return standup.cache.reports.filter((p) => p.id !== id);
    }
    function postUpdate(x, id) {
        report.setReports(x);
        if (id === standup.cache.activeReport) {
            modal.hide("report");
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
        const param = { storyID: story.id, status };
        socket.send({ svc: services.estimate.key, cmd: command.client.setStoryStatus, param: param });
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
        let currStory = undefined;
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
        comment.setCounts();
        modal.hide("add-story");
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
        if (title === null) {
            return false;
        }
        if (title && title !== s.title) {
            const msg = { svc: services.estimate.key, cmd: command.client.updateStory, param: { storyID: s.id, title } };
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
                modal.hide("story");
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
        if (!curr) {
            console.warn(`cannot load active story [${estimate.cache.activeStory}]`);
        }
        return curr;
    }
    story.getActiveStory = getActiveStory;
    function viewActiveStory(id) {
        if (id) {
            estimate.cache.activeStory = id;
        }
        const s = getActiveStory();
        if (!s) {
            return;
        }
        dom.setText("#story-title", s.title);
        story.viewStoryStatus(s.status);
        comment.setActive("story", s.id);
        comment.setCounts();
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
var story;
(function (story_2) {
    function renderStory(story) {
        const profile = system.cache.getProfile();
        return JSX("li", { id: `story-${story.id}`, class: "section", onclick: `modal.open('story', '${story.id}');` },
            JSX("div", { class: "right uk-article-meta story-status" }, story.status),
            JSX("div", { class: `${profile.linkColor}-fg section-link left` }, story.title),
            JSX("div", { class: "left", style: "margin-top: 4px;" }, comment.renderCount("story", story.id)));
    }
    function renderStories(stories) {
        if (stories.length === 0) {
            return JSX("div", { id: "story-list" },
                JSX("button", { class: "uk-button uk-button-default", onclick: "modal.open('add-story');", type: "button" }, "Add Story"));
        }
        else {
            return JSX("table", { class: "uk-table uk-table-divider uk-table-small" },
                JSX("ul", { id: "story-list", class: "uk-list uk-list-divider" }, stories.map(s => renderStory(s))));
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
                return JSX("span", { class: "vote-badge" }, status);
        }
    }
    story_2.renderStatus = renderStatus;
    function renderTotal(sum) {
        return JSX("li", { id: "story-total" },
            JSX("div", { class: "right uk-article-meta" },
                JSX("span", { class: "vote-badge" }, sum)),
            " Total");
    }
    story_2.renderTotal = renderTotal;
    function viewAddStory() {
        const storyInput = dom.setValue("#story-title-input", "");
        setTimeout(() => storyInput.focus(), 250);
    }
    story_2.viewAddStory = viewAddStory;
})(story || (story = {}));
const debug = true;
var system;
(function (system) {
    class Cache {
        constructor() {
            this.currentID = "";
            this.connectTime = 0;
        }
        getProfile() {
            if (!this.profile) {
                throw "no active profile";
            }
            return this.profile;
        }
        apply(sj) {
            system.cache.session = sj.session;
            system.cache.profile = sj.profile;
            auth.applyAuths(sj.auths);
            permission.applyPermissions(sj.permissions);
            member.applyMembers(sj.members);
            member.applyOnline(sj.online);
            comment.applyComments(sj.comments);
            if (sj.team) {
                session.setTeam(sj.team);
            }
            if (sj.sprint) {
                session.setSprint(sj.sprint);
            }
        }
    }
    system.cache = new Cache();
    function setPermissions(perms) {
        permission.applyPermissions(perms);
        permission.setPerms();
    }
    system.setPermissions = setPermissions;
    // noinspection JSUnusedGlobalSymbols
    function setAuths(auths) {
        auth.applyAuths(auths);
        permission.setPerms();
    }
    system.setAuths = setAuths;
    function onSystemMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.system, param);
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
            case command.server.memberRemove:
                member.onMemberRemove(param);
                break;
            case command.server.onlineUpdate:
                member.onOnlineUpdate(param);
                break;
            case command.server.commentUpdate:
                comment.onCommentUpdate(param);
                break;
            case command.server.commentRemove:
                comment.onCommentRemoved(param);
                break;
            default:
                console.warn(`unhandled system message for command [${cmd}]`);
        }
    }
    system.onSystemMessage = onSystemMessage;
    function onSessionRemove(svc) {
        document.location.reload();
    }
    system.onSessionRemove = onSessionRemove;
})(system || (system = {}));
var team;
(function (team) {
    class Cache {
    }
    team.cache = new Cache();
    function onTeamMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                rituals.onError(services.team, param);
                break;
            case command.server.sessionJoined:
                const sj = param;
                session.onSessionJoin(sj);
                setTeamDetail(sj.session);
                setTeamHistory(sj);
                session.showWelcomeMessage(sj.members.length);
                break;
            case command.server.sessionUpdate:
                setTeamDetail(param);
                break;
            case command.server.sessionRemove:
                system.onSessionRemove(services.team);
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
        session.setDetail(detail);
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
            // dom.setDisplay(c, teams.length > 0)
            dom.setContent("#model-team-select", team.renderTeamSelect(teams, (_a = system.cache.session) === null || _a === void 0 ? void 0 : _a.teamID));
            permission.setModelPerms("team");
        }
    }
    team.viewTeams = viewTeams;
})(team || (team = {}));
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
var profile;
(function (profile) {
    // noinspection JSUnusedGlobalSymbols
    function setNavColor(el, c) {
        dom.setValue("#nav-color", c);
        const nb = dom.req("#navbar");
        nb.className = `${c}-bg uk-navbar-container uk-navbar`;
        const colors = document.querySelectorAll(".nav_swatch");
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
    function setPicture(p) {
        dom.setValue('#self-picture-input', p);
        return false;
    }
    profile.setPicture = setPicture;
})(profile || (profile = {}));
var collection;
(function (collection) {
    class Group {
        constructor(key) {
            this.members = [];
            this.key = key;
        }
    }
    collection.Group = Group;
    class GroupSet {
        constructor() {
            this.groups = [];
        }
        findOrInsert(key) {
            const ret = this.groups.find(x => x.key === key);
            if (ret) {
                return ret;
            }
            const n = new Group(key);
            this.groups.push(n);
            return n;
        }
    }
    collection.GroupSet = GroupSet;
    function groupBy(list, func) {
        const res = new GroupSet();
        if (list) {
            list.forEach((o) => {
                const group = res.findOrInsert(func(o));
                group.members.push(o);
            });
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
    const tzOffset = new Date().getTimezoneOffset() * 60000;
    function utcDate(s) {
        return new Date(Date.parse(s) + tzOffset);
    }
    date_1.utcDate = utcDate;
})(date || (date = {}));
var dom;
(function (dom) {
    function initDom(t, color) {
        try {
            style.themeLinks(color);
            style.setTheme(t);
        }
        catch (e) {
            console.warn("error setting style", e);
        }
        try {
            modal.wire();
        }
        catch (e) {
            console.warn("error wiring modals", e);
        }
    }
    dom.initDom = initDom;
    function els(selector, context) {
        return UIkit.util.$$(selector, context);
    }
    dom.els = els;
    function opt(selector, context) {
        const e = els(selector, context);
        switch (e.length) {
            case 0:
                return undefined;
            case 1:
                return e[0];
            default:
                console.warn(`found [${e.length}] elements with selector [${selector}], wanted zero or one`);
        }
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
})(dom || (dom = {}));
var dom;
(function (dom) {
    function setValue(el, text) {
        if (typeof el === "string") {
            el = dom.req(el);
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
            el = dom.req(el);
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
            el = dom.req(el);
        }
        for (let i = 0; i < el.children.length; i++) {
            const e = el.children.item(i);
            e.selected = e.value === o;
        }
    }
    dom.setSelectOption = setSelectOption;
    function insertAtCaret(e, text) {
        if (e.selectionStart || e.selectionStart === 0) {
            var startPos = e.selectionStart;
            var endPos = e.selectionEnd;
            e.value = e.value.substring(0, startPos) + text + e.value.substring(endPos, e.value.length);
            e.selectionStart = startPos + text.length;
            e.selectionEnd = startPos + text.length;
        }
        else {
            e.value += text;
        }
    }
    dom.insertAtCaret = insertAtCaret;
})(dom || (dom = {}));
// noinspection JSUnusedGlobalSymbols
function JSX(tag, attrs) {
    const e = document.createElement(tag);
    for (const name in attrs) {
        if (name && attrs.hasOwnProperty(name)) {
            const v = attrs[name];
            if (name === "dangerouslySetInnerHTML") {
                e.innerHTML = v["__html"];
            }
            else if (v === true) {
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
        else if (child === undefined || child === null) {
            throw `child for tag [${tag}] is ${child}`;
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
var modal;
(function (modal) {
    let openEvents = new Map();
    let closeEvents = new Map();
    let activeParam;
    function register(key, o, c) {
        if (!o) {
            o = () => { };
        }
        openEvents.set(key, o);
        if (c) {
            closeEvents.set(key, c);
        }
    }
    modal.register = register;
    function wire() {
        UIkit.util.on(".drop", "show", onDropOpen);
        UIkit.util.on(".drop", "beforehide", onDropBeforeHide);
        UIkit.util.on(".drop", "hide", onDropHide);
        UIkit.util.on(".modal", "show", onModalOpen);
        UIkit.util.on(".modal", "hide", onModalHide);
        register("welcome");
        // session
        register("session", session.onModalOpen);
        register("action", action.loadActions);
        register("comment", comment.load, comment.closeModal);
        // member
        register("self", member.viewSelf);
        register("invitation");
        register("member", member.viewActiveMember);
        // estimate
        register("add-story", story.viewAddStory);
        register("story", story.viewActiveStory);
        // standup
        register("add-report", report.viewAddReport);
        register("report", report.viewReport);
        // retro
        register("add-feedback", feedback.viewAddFeedback);
        register("feedback", feedback.viewFeedback);
    }
    modal.wire = wire;
    function open(key, param) {
        activeParam = param;
        const m = UIkit.modal(`#modal-${key}`);
        if (!m) {
            console.warn(`no modal available with key [${key}]`);
        }
        m.show();
        return false;
    }
    modal.open = open;
    function openSoon(key) {
        setTimeout(() => open(key), 0);
    }
    modal.openSoon = openSoon;
    function hide(key) {
        const m = UIkit.modal(`#modal-${key}`);
        const el = m.$el;
        if (el.classList.contains("uk-open")) {
            m.hide();
        }
    }
    modal.hide = hide;
    function onModalOpen(e) {
        if (!e.target) {
            return;
        }
        const el = e.target;
        if (el.id.indexOf("modal") !== 0) {
            return;
        }
        const key = el.id.substr("modal-".length);
        const f = openEvents.get(key);
        if (f) {
            f(activeParam);
        }
        else {
            console.warn(`no modal open handler registered for [${key}]`);
        }
        activeParam = undefined;
    }
    function onModalHide(e) {
        if (!e.target) {
            return;
        }
        const el = e.target;
        if (el.classList.contains("uk-open")) {
            const key = el.id.substr("modal-".length);
            const f = closeEvents.get(key);
            if (f) {
                f(activeParam);
            }
            activeParam = undefined;
        }
    }
    function onDropOpen(e) {
        if (!e.target) {
            return;
        }
        const el = e.target;
        const key = el.dataset["key"] || "";
        let t = el.dataset["t"] || "";
        const f = openEvents.get(key);
        if (f) {
            f(t);
        }
        else {
            console.warn(`no modal open handler registered for [${key}]`);
        }
    }
    function onDropHide(e) {
        if (!e.target) {
            return;
        }
        const el = e.target;
        if (el.classList.contains("uk-open")) {
            const key = el.dataset["key"] || "";
            const t = el.dataset["t"] || "";
            const f = closeEvents.get(key);
            if (f) {
                f(t);
            }
        }
    }
    let emojiPicked = false;
    function onEmojiPicked() {
        emojiPicked = true;
        setTimeout(() => emojiPicked = false, 200);
    }
    modal.onEmojiPicked = onEmojiPicked;
    function onDropBeforeHide(e) {
        if (emojiPicked) {
            e.preventDefault();
        }
    }
})(modal || (modal = {}));
var notify;
(function (notify_1) {
    function notify(msg, status) {
        UIkit.notification(msg, { status: status ? "success" : "danger", pos: "top-right" });
    }
    notify_1.notify = notify;
})(notify || (notify = {}));
var style;
(function (style) {
    function setTheme(theme) {
        wireEmoji(theme);
        const card = dom.els(".uk-card");
        switch (theme) {
            case "auto":
                let t = "light";
                if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
                    t = "dark";
                }
                setTheme(t);
                fetch("/profile/theme/" + t).then(r => r.text()).then(body => {
                    console.log(`Set theme to [${t}]`);
                });
                break;
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
    style.setTheme = setTheme;
    function themeLinks(color) {
        dom.els(".theme").forEach(el => {
            el.classList.add(`${color}-fg`);
        });
    }
    style.themeLinks = themeLinks;
    function wireEmoji(t) {
        if (typeof EmojiButton === 'undefined') {
            dom.els(".picker-toggle").forEach(el => dom.setDisplay(el, false));
            return;
        }
        const opts = { position: "bottom-end", theme: t, zIndex: 1021 };
        dom.els(".textarea-emoji").forEach(el => {
            const toggle = dom.req(".picker-toggle", el);
            toggle.addEventListener("click", () => {
                const textarea = dom.req(".uk-textarea", el);
                const picker = new EmojiButton(opts);
                picker.on('emoji', (emoji) => {
                    modal.onEmojiPicked();
                    dom.insertAtCaret(textarea, emoji);
                });
                picker.togglePicker(toggle);
            }, false);
        });
    }
})(style || (style = {}));
var vote;
(function (vote) {
    function setVotes(votes) {
        estimate.cache.votes = votes;
        viewVotes();
    }
    vote.setVotes = setVotes;
    function onVoteUpdate(v) {
        let x = estimate.cache.votes;
        x = x.filter(vt => vt.userID !== v.userID || vt.storyID !== v.storyID);
        x.push(v);
        estimate.cache.votes = x;
        if (v.storyID === estimate.cache.activeStory) {
            viewVotes();
        }
    }
    vote.onVoteUpdate = onVoteUpdate;
    function viewVotes() {
        const s = story.getActiveStory();
        if (!s) {
            return;
        }
        const votes = estimate.cache.activeVotes();
        const activeVote = votes.filter(v => v.userID === system.cache.getProfile().userID).pop();
        switch (s.status) {
            case "pending":
                const same = system.cache.getProfile().userID === s.userID;
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
        console.log("!!!!!!!!");
        console.log(votes);
        dom.setContent("#story-vote-members", vote.renderVoteMembers(member.getMembers(), votes));
        dom.setContent("#story-vote-choices", vote.renderVoteChoices(estimate.cache.detail.choices, activeVote === null || activeVote === void 0 ? void 0 : activeVote.choice));
    }
    function viewVoteResults(votes) {
        dom.setContent("#story-vote-results", vote.renderVoteResults(member.getMembers(), votes));
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
//# sourceMappingURL=rituals.js.map