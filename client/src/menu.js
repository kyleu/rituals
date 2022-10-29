"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.menuInit = void 0;
// Content managed by Project Forge, see [projectforge.md] for details.
function menuInit() {
    for (const n of Array.from(document.querySelectorAll(".menu-container .final"))) {
        n.scrollIntoView({ block: "nearest" });
    }
}
exports.menuInit = menuInit;
