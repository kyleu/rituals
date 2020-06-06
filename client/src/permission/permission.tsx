namespace permission {
  export function renderEmails(key: string, emails: ReadonlyArray<Email>) {
    const cls = `uk-checkbox uk-margin-small-right perm-${key}-email`;
    const oc = `permission.onChanged('email', '${key}', this.checked)`;
    return <ul>
      {emails.map(e => {
        return <li>
          <label>
            {e.matched ? <input class={cls} type="checkbox" value={e.domain} checked="checked" onchange={oc} /> : <input class={cls} type="checkbox" value={e.domain} onchange={oc} />}
            Using email address {e.domain}
          </label>
        </li>;
      })}
    </ul>;
  }

  function renderMessage(p: collection.Group<string, permission.Permission>) {
    switch(p.key) {
      case services.team.key:
        return <li>Must be a member of this session's team</li>
      case services.sprint.key:
        return <li>Must be a member of this session's team</li>
      default:
        let x = collection.flatten(p.members.map(x => x.k.split(",").map(x => x.trim()).filter(x => x.length > 0)));
        if (x.length === 0) {
          return <li>Must sign in with {p.key}</li>
        }
        return <li>Must sign in with {p.key} using an email address from {x.join(" or ")}</li>
    }
  }

  export function renderView(perms: collection.Group<string, Permission>[]) {
    if (perms.length === 0) {
      return <div>public</div>;
    }
    return <ul>{perms.map(p => renderMessage(p))}</ul>;
  }
}
