"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.appInit = void 0;
// $PF_IGNORE$
const x_1 = require("./x");
const dom_1 = require("./dom");
function appInit() {
    (0, dom_1.req)("table").appendChild((0, x_1.test)());
}
exports.appInit = appInit;
