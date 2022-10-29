"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.modeInit = void 0;
// Content managed by Project Forge, see [projectforge.md] for details.
const l = "mode-light";
const d = "mode-dark";
function modeInit() {
    for (const el of Array.from(document.getElementsByClassName("mode-input"))) {
        const i = el;
        i.onclick = function () {
            switch (i.value) {
                case "":
                    document.body.classList.remove(l);
                    document.body.classList.remove(d);
                    break;
                case "light":
                    document.body.classList.add(l);
                    document.body.classList.remove(d);
                    break;
                case "dark":
                    document.body.classList.remove(l);
                    document.body.classList.add(d);
                    break;
            }
        };
    }
}
exports.modeInit = modeInit;
