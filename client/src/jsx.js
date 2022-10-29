"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.JSX = void 0;
// Content managed by Project Forge, see [projectforge.md] for details.
const dom_1 = require("./dom");
// noinspection JSUnusedGlobalSymbols
function JSX(tag, attrs) {
    const e = document.createElement(tag);
    for (const name in attrs) {
        if (name && attrs.hasOwnProperty(name)) {
            const v = attrs[name];
            if (name === "dangerouslySetInnerHTML") {
                (0, dom_1.setHTML)(e, v["__html"]);
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
                if (child === undefined || child === null) {
                    throw `child array for tag [${tag}] is ${child}\n${e.outerHTML}`;
                }
                if (c === undefined || c === null) {
                    throw `child for tag [${tag}] is ${c}\n${e.outerHTML}`;
                }
                if (typeof c === "string") {
                    c = document.createTextNode(c);
                }
                e.appendChild(c);
            });
        }
        else if (child === undefined || child === null) {
            throw `child for tag [${tag}] is ${child}\n${e.outerHTML}`;
            // debugger;
            // child = document.createTextNode("NULL!");
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
exports.JSX = JSX;
