"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.themeInit = void 0;
// Content managed by Project Forge, see [projectforge.md] for details.
const keys = [];
function themeInit() {
    const x = document.querySelectorAll(".color-var");
    if (x.length > 0) {
        for (const el of Array.from(x)) {
            const i = el;
            const v = i.dataset["var"];
            const m = i.dataset["mode"];
            keys.push(v);
            if (!v || v.length === 0) {
                continue;
            }
            i.oninput = function () {
                set(m, v, i.value);
            };
        }
    }
}
exports.themeInit = themeInit;
function set(mode, key, v) {
    const mockup = document.querySelector("#mockup-" + mode);
    if (!mockup) {
        console.error("can't find mockup for mode [" + mode + "]");
        return;
    }
    switch (key) {
        case "color-foreground":
            setFG(mockup, ".mock-main", v);
            break;
        case "color-background":
            setBG(mockup, ".mock-main", v);
            break;
        case "color-foreground-muted":
            setFG(mockup, ".mock-main .mock-muted", v);
            break;
        case "color-background-muted":
            setBG(mockup, ".mock-main .mock-muted", v);
            break;
        case "color-link-foreground":
            setFG(mockup, ".mock-main .mock-link", v);
            break;
        case "color-link-visited-foreground":
            setFG(mockup, ".mock-main .mock-link-visited", v);
            break;
        case "color-nav-foreground":
            setFG(mockup, ".mock-nav", v);
            setFG(mockup, ".mock-nav .mock-link", v);
            break;
        case "color-nav-background":
            setBG(mockup, ".mock-nav", v);
            break;
        case "color-menu-foreground":
            setFG(mockup, ".mock-menu", v);
            setFG(mockup, ".mock-menu .mock-link", v);
            break;
        case "color-menu-background":
            setBG(mockup, ".mock-menu", v);
            break;
        case "color-menu-selected-foreground":
            setFG(mockup, ".mock-menu .mock-link-selected", v);
            break;
        case "color-menu-selected-background":
            setBG(mockup, ".mock-menu .mock-link-selected", v);
            break;
        default:
            console.error("invalid key [" + key + "]");
    }
}
function call(mockup, sel, f) {
    const q = mockup.querySelectorAll(sel);
    if (q.length == 0) {
        throw "empty query selector [" + sel + "]";
    }
    q.forEach(x => {
        f(x);
    });
}
function setBG(mockup, sel, v) {
    call(mockup, sel, el => el.style.backgroundColor = v);
}
function setFG(mockup, sel, v) {
    call(mockup, sel, el => el.style.color = v);
}
