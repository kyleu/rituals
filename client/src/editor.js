"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.editorInit = exports.setSiblingToNull = void 0;
// Content managed by Project Forge, see [projectforge.md] for details.
const selected = "--selected";
function setSiblingToNull(el) {
    const i = el.parentElement.parentElement.querySelector("input");
    if (!i) {
        throw "no associated input found";
    }
    i.value = "∅";
}
exports.setSiblingToNull = setSiblingToNull;
function editorInit() {
    window.rituals.setSiblingToNull = setSiblingToNull;
    let editorCache = {};
    let selectedCache = {};
    for (const n of Array.from(document.querySelectorAll(".editor"))) {
        const frm = n;
        const buildCache = () => {
            editorCache = {};
            selectedCache = {};
            for (const el of frm.elements) {
                const input = el;
                if (input.name.length > 0) {
                    if (input.name.endsWith(selected)) {
                        selectedCache[input.name] = input;
                    }
                    else {
                        if ((input.type !== "radio") || input.checked) {
                            editorCache[input.name] = input.value;
                        }
                        const evt = () => {
                            const cv = selectedCache[input.name + selected];
                            if (cv) {
                                cv.checked = editorCache[input.name] !== input.value;
                            }
                        };
                        input.onchange = evt;
                        input.onkeyup = evt;
                    }
                }
            }
        };
        frm.onreset = buildCache;
        buildCache();
    }
}
exports.editorInit = editorInit;
