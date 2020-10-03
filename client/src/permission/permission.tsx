namespace permission {
  export function renderEmails(key: string, emails: readonly Email[]) {
    const cls = `uk-checkbox uk-margin-small-right perm-${key}-email`;
    const oc = `permission.onChanged('email', '${key}', this.checked)`;
    return (
      <ul>
        {emails.map(e => {
          return (
            <li>
              <label>
                {e.matched ? (
                  <input class={cls} type="checkbox" value={e.domain} checked="checked" onchange={oc} />
                ) : (
                  <input class={cls} type="checkbox" value={e.domain} onchange={oc} />
                )}
                Using email address {e.domain}
              </label>
            </li>
          );
        })}
      </ul>
    );
  }

  function renderMessage(p: group.Group<string, permission.Permission>) {
    switch (p.key) {
      case services.team.key:
        return <li>Must be a member of this session's team</li>;
      case services.sprint.key:
        return <li>Must be a member of this session's sprint</li>;
      default:
        let col = group.flatten(p.members.map(x => x.k.split(",").map(y => y.trim()).filter(z => z.length > 0)));
        if (col.length === 0) {
          return <li>Must sign in with {p.key}</li>;
        }
        return <li>Must sign in with {p.key} using an email address from {col.join(" or ")}</li>;
    }
  }

  export function renderView(perms: group.Group<string, Permission>[]) {
    if (perms.length === 0) {
      return <div>public</div>;
    }
    return <ul>{perms.map(p => renderMessage(p))}</ul>;
  }
}
