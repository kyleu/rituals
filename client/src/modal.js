"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.modalInit = void 0;
// Content managed by Project Forge, see [projectforge.md] for details.
function modalInit() {
    document.addEventListener('keydown', (event) => {
        if (event.key === 'Escape') {
            if (document.location.hash.startsWith("#modal-")) {
                document.location.hash = "";
            }
        }
    });
}
exports.modalInit = modalInit;
