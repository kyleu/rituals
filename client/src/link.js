"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.linkInit = void 0;
// Content managed by Project Forge, see [projectforge.md] for details.
function linkInit() {
    for (const l of Array.from(document.getElementsByClassName("link-confirm"))) {
        const el = l;
        el.onclick = function () {
            let msg = el.dataset.message;
            if (msg && msg.length === 0) {
                msg = "Are you sure?";
            }
            return confirm(msg);
        };
    }
}
exports.linkInit = linkInit;
