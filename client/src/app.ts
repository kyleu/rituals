// $PF_IGNORE$
import {test} from "./x";
import {req} from "./dom";

export function appInit(): void {
  req("table").appendChild(test());
}
