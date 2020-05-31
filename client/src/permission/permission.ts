namespace permission {
  export interface Permission {
    readonly k: string;
    readonly v: string;
    readonly access: string;
  }

  let permissions: collection.Group<string, Permission>[] = [];

  export interface Email {
    readonly matched: boolean,
    readonly domain: string
  }

  export function setPerms() {
    dom.setDisplay("#public-link-container", permissions === null || permissions.length === 0);
    dom.setDisplay("#private-link-container", permissions !== null && permissions.length > 0);
    ["team", "sprint"].forEach(setModelPerms);
    auth.allProviders.forEach(setProviderPerms);
  }

  export function setModelPerms(key: string) {
    const el = dom.opt<HTMLSelectElement>(`#model-${key}-select select`);
    if (el) {
      const perms = collection.findGroup(permissions, key);
      const section = dom.opt(`#perm-${key}-section`);
      if (section) {
        const checkbox = dom.req<HTMLInputElement>(`#perm-${key}-checkbox`);
        checkbox.checked = perms.length > 0;
        dom.setDisplay(section, el.value !== "");
      }
      collection.findGroup(permissions, key);
    }
  }

  // noinspection JSUnusedGlobalSymbols
  export function onChanged(k: string, v: string, checked: boolean) {
    switch (k) {
      case "email":
        return onEmailChanged(v, checked);
      case "provider":
        return onProviderChanged(v, checked);
    }
  }

  function onEmailChanged(key: string, checked: boolean) {
    const checkbox = dom.opt<HTMLInputElement>(`#perm-${key}-checkbox`);
    if(checkbox && checked && !checkbox.checked) {
      checkbox.checked = true;
    }
  }

  function onProviderChanged(key: string, checked: boolean) {
    dom.els<HTMLInputElement>(`.perm-${key}-email`).forEach(el => {
      el.disabled = !checked;
      if (!checked) {
        el.checked = false;
      }
    })
  }

  function setProviderPerms(p: auth.Provider) {
    const perms = collection.findGroup(permissions, p.key);
    const auths = auth.active().filter(a => a.provider === p.key);

    const section = dom.opt(`#perm-${p.key}-section`);
    if (section) {
      const checkbox = dom.req<HTMLInputElement>(`#perm-${p.key}-checkbox`);
      checkbox.checked = perms.length > 0;

      const emailContainer = dom.req(`#perm-${p.key}-email-container`);
      const emails = collection.flatten(perms.map(x => x.v.split(",").filter(x => x.length > 0))).map(x => ({matched: true, domain: x}));

      const additional = auths.filter(a => emails.filter(e => a.email.endsWith(e.domain)).length === 0).map(m => {
        return {matched: false, domain: getDomain(m.email)};
      });
      emails.push(...additional);
      emails.sort();

      dom.setDisplay(emailContainer, emails.length > 0);
      dom.setContent(emailContainer, emails.length === 0 ? document.createElement("span") : permission.renderEmails(p.key, emails))
    }
  }

  function getDomain(email: string) {
    const idx = email.lastIndexOf("@");
    if (idx === -1) {
      return email;
    }
    return email.substr(idx);
  }

  export function applyPermissions(perms: permission.Permission[] | null) {
    permissions = collection.groupBy(perms, x => x.k).groups;
  }
}
