"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.init = void 0;
// Content managed by Project Forge, see [projectforge.md] for details.
require("./client.css");
const jsx_1 = require("./jsx");
const menu_1 = require("./menu");
const mode_1 = require("./mode");
const flash_1 = require("./flash");
const link_1 = require("./link");
const modal_1 = require("./modal");
const editor_1 = require("./editor");
const theme_1 = require("./theme");
const socket_1 = require("./socket");
const app_1 = require("./app");
function init() {
    window.rituals = {};
    window.JSX = jsx_1.JSX;
    (0, menu_1.menuInit)();
    (0, mode_1.modeInit)();
    (0, flash_1.flashInit)();
    (0, link_1.linkInit)();
    (0, modal_1.modalInit)();
    (0, editor_1.editorInit)();
    (0, theme_1.themeInit)();
    (0, socket_1.socketInit)();
    (0, app_1.appInit)();
}
exports.init = init;
document.addEventListener("DOMContentLoaded", init);
