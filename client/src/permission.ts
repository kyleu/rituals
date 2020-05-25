namespace permission {
  export interface Permission {
    readonly k: string;
    readonly v: string;
    readonly access: string;
  }

  export interface Auth {
    readonly provider: string;
    readonly email: string;
  }

  export interface Provider {
    readonly key: string,
    readonly title: string
  }

  const github: Provider = {key: "github", title: "GitHub"};
  const google: Provider = {key: "google", title: "Google"};
  const slack: Provider = {key: "slack", title: "Slack"};
  const amazon: Provider = {key: "amazon", title: "Amazon"};
  const microsoft: Provider = {key: "microsoft", title: "Microsoft"};

  export const allProviders = [github, google, slack, amazon, microsoft];

  export interface Email {
    readonly matched: boolean,
    readonly domain: string
  }

  export function setPerms() {
    ["team", "sprint"].forEach(setModelPerms);
    if (system.cache.auths != null) {
      allProviders.forEach(setProviderPerms);
    }
  }

  export function setModelPerms(key: string) {
    const el = dom.opt<HTMLSelectElement>(`#model-${key}-select select`);
    if (el) {
      const perms = collection.findGroup(system.cache.permissions, key);
      const section = dom.opt(`#perm-${key}-section`);
      if (section) {
        const checkbox = dom.req<HTMLInputElement>(`#perm-${key}-checkbox`);
        checkbox.checked = perms.length > 0;
        dom.setDisplay(section, el.value != "");
      }
      collection.findGroup(system.cache.permissions, key);
    }
  }

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

  function setProviderPerms(p: Provider) {
    const perms = collection.findGroup(system.cache.permissions, p.key);
    const auths = system.cache.auths.filter(a => a.provider == p.key);

    const section = dom.opt(`#perm-${p.key}-section`);
    if (section) {
      const checkbox = dom.req<HTMLInputElement>(`#perm-${p.key}-checkbox`);
      checkbox.checked = perms.length > 0;

      const emailContainer = dom.req(`#perm-${p.key}-email-container`);
      const emails = collection.flatten(perms.map(x => x.v.split(",").filter(x => x.length > 0))).map(x => ({matched: true, domain: x}));

      const additional = auths.filter(a => emails.filter(e => a.email.endsWith(e.domain)).length == 0).map(m => {
        return {matched: false, domain: getDomain(m.email)};
      });
      emails.push(...additional);
      emails.sort();

      dom.setDisplay(emailContainer, emails.length > 0);
      dom.setContent(emailContainer, emails.length == 0 ? document.createElement("span") : permission.renderEmails(p.key, emails))
    }
  }

  function getDomain(email: string) {
    const idx = email.lastIndexOf("@");
    if (idx == -1) {
      return email;
    }
    return email.substr(idx);
  }
}

