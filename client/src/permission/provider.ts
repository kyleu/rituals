namespace permission {
  export function setProviderPerms(p: auth.Provider) {
    const perms = group.findGroup(permissions, p.key);
    const auths = auth.active().filter(a => a.provider === p.key);

    const section = dom.opt(`#perm-${p.key}-section`);
    if (section) {
      const checkbox = dom.req<HTMLInputElement>(`#perm-${p.key}-checkbox`);
      checkbox.checked = perms.length > 0;

      const emailContainer = dom.req(`#perm-${p.key}-email-container`);
      const emails = group.flatten(perms.map(x => x.v.split(",").filter(y => y.length > 0))).map(z => ({ matched: true, domain: z }));

      const additional = auths.filter(a => emails.filter(e => a.email.endsWith(e.domain)).length === 0).map(m => (
        { matched: false, domain: getDomain(m.email) }
      ));
      emails.push(...additional);
      emails.sort();

      dom.setDisplay(emailContainer, emails.length > 0);
      dom.setContent(emailContainer, emails.length === 0 ? document.createElement("span") : permission.renderEmails(p.key, emails));
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
    if (checkbox && checked && !checkbox.checked) {
      checkbox.checked = true;
    }
  }

  function onProviderChanged(key: string, checked: boolean) {
    dom.els<HTMLInputElement>(`.perm-${key}-email`).forEach(el => {
      el.disabled = !checked;
      if (!checked) {
        el.checked = false;
      }
    });
  }

  function getDomain(email: string) {
    const idx = email.lastIndexOf("@");
    if (idx === -1) {
      return email;
    }
    return email.substr(idx);
  }
}
