"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.flashInit = void 0;
// Content managed by Project Forge, see [projectforge.md] for details.
function flashInit() {
    const container = document.getElementById("flash-container");
    if (container === null) {
        return;
    }
    const x = container.querySelectorAll(".flash");
    if (x.length > 0) {
        setTimeout(() => {
            for (const f of x) {
                const el = f;
                el.style.opacity = "0";
                setTimeout(() => el.remove(), 500);
            }
        }, 3000);
    }
}
exports.flashInit = flashInit;
