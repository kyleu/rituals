"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.clear = exports.setText = exports.setContent = exports.setDisplay = exports.setHTML = exports.req = exports.opt = exports.els = void 0;
function els(selector, context) {
    let result;
    if (context) {
        result = context.querySelectorAll(selector);
    }
    else {
        result = document.querySelectorAll(selector);
    }
    const ret = [];
    result.forEach(v => {
        ret.push(v);
    });
    return ret;
}
exports.els = els;
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
exports.opt = opt;
function req(selector, context) {
    const res = opt(selector, context);
    if (!res) {
        console.warn(`no element found for selector [${selector}]`);
    }
    return res;
}
exports.req = req;
function setHTML(el, html) {
    if (typeof el === "string") {
        el = req(el);
    }
    el.innerHTML = html;
    return el;
}
exports.setHTML = setHTML;
function setDisplay(el, condition, v = "block") {
    if (typeof el === "string") {
        el = req(el);
    }
    el.style.display = condition ? v : "none";
    return el;
}
exports.setDisplay = setDisplay;
function setContent(el, e) {
    if (typeof el === "string") {
        el = req(el);
    }
    clear(el);
    if (Array.isArray(e)) {
        e.forEach(x => el.appendChild(x));
    }
    else {
        el.appendChild(e);
    }
    return el;
}
exports.setContent = setContent;
function setText(el, text) {
    if (typeof el === "string") {
        el = req(el);
    }
    el.innerText = text;
    return el;
}
exports.setText = setText;
function clear(el) {
    return setHTML(el, "");
}
exports.clear = clear;
