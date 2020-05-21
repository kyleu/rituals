namespace permission {
  function renderPerm(perm: Permission): JSX.Element {
    return <div>{perm.k}:{perm.v}</div>;
  }

  function basicPerms(title: string, perms: Permission[]) {
    return <div>
      <div>{title}</div>
      {perms.map(renderPerm)}
    </div>
  }

  function teamPerms(perms: Permission[]) {
    return basicPerms("Team", perms)
  }

  function sprintPerms(perms: Permission[]) {
    return basicPerms("Sprint", perms)
  }

  function invitationPerms(perms: Permission[]) {
    return basicPerms("Invitation", perms)
  }

  function googlePerms(perms: Permission[]) {
    return basicPerms("Google", perms)
  }

  function githubPerms(perms: Permission[]) {
    return basicPerms("GitHub", perms)
  }

  function slackPerms(perms: Permission[]) {
    return basicPerms("Slack", perms)
  }

  export function renderPermissions(perms: Permission[]): JSX.Element {
    const g = collection.groupBy(perms, x => x.k);
    return <ul class="uk-list uk-list-divider">
      {teamPerms(findGroup("team", g))}
      {sprintPerms(findGroup("sprint", g))}
      {invitationPerms(findGroup("invitation", g))}
      {googlePerms(findGroup("google", g))}
      {githubPerms(findGroup("github", g))}
      {slackPerms(findGroup("slack", g))}
    </ul>;
  }

  function findGroup(key: string, groups: collection.Group<Permission>[]): Permission[] {
    let ret: Permission[] = [];
    for (const g of groups) {
      if (g.key == key) {
        ret = g.members;
        break;
      }
    }
    return ret
  }
}
