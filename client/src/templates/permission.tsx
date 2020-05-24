namespace permission {
  export function renderEmails(key: string, emails: Email[]): JSX.Element {
    const cls = `uk-checkbox uk-margin-small-right perm-${key}-email`;
    const oc = `permission.onChanged('email', '${key}', this.checked)`;
    return <ul>
      {emails.map(e => {
        return <li>
          <label>
            {e.matched ? <input class={cls} type="checkbox" value={e.domain} checked="checked" onchange={oc} /> : <input class={cls} type="checkbox" value={e.domain} onclick={oc} />}
            Using email address {e.domain}
          </label>
        </li>;
      })}
    </ul>;
  }

  function readPermission(k: string): Permission[] {
    if(!dom.req<HTMLInputElement>(`#perm-${k}-checkbox`).checked) {
      return [];
    }

    const emails = dom.els<HTMLInputElement>(`.perm-${k}-email`);
    const v = emails.filter(e => e.checked).map(e => e.value).join(",");

    const access = "member";

    return [{k, v, access}];
  }

  export function readPermissions(): Permission[] {
    const ret = [];

    ret.push(...readPermission("team"));
    ret.push(...readPermission("sprint"));
    ret.push(...readPermission("github"));
    ret.push(...readPermission("google"));
    ret.push(...readPermission("slack"));

    return ret;
  }
}
