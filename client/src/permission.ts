import {els, req} from "./dom";

export type Permission = {
  key: string;
  value: string;
}

export function permissionsUpdate(perms: Permission[]) {
  console.log("TODO: permissionsUpdate", perms);
}

export function permissionsTeamToggle(set: boolean) {
  req(".permission-config-team").style.display = set ? "block" : "none";
}

export function permissionsSprintToggle(set: boolean) {
  req(".permission-config-sprint").style.display = set ? "block" : "none";
}

export function loadPermsForm(frm: HTMLFormElement) {
  type StrMap = { [key: string]: string };
  const ret: StrMap = {};
  for (const el of els<HTMLInputElement>(".perm-option", frm)) {
    if (el.checked) {
      ret[el.name] = el.value;
    }
  }
  return ret;
}

export function initPermissions(t: HTMLSelectElement, s?: HTMLSelectElement) {
  permissionsTeamToggle(t.value !== "");
  if (s) {
    permissionsSprintToggle(s.value !== "");
  }
}
