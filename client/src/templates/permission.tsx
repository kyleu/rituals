namespace permission {
  function renderPerm(perm: Permission): JSX.Element {
    return <div>{perm.k}:{perm.v}</div>;
  }

  function basicPerms(title: string, perms: permission.Permission[], auths: auth.Auth[]) {
    return <li>
      <div>{title}</div>
      {perms.map(renderPerm)}
    </li>
  }

  function teamPerms(teamID: string | undefined, perms: Permission[]) {
    if(teamID) {
      return basicPerms("Team", perms, [])
    }
    return <span />
  }

  function sprintPerms(sprintID: string | undefined, perms: Permission[]) {
    if(sprintID) {
      return basicPerms("Sprint", perms, [])
    }
    return <span />
  }

  function invitationPerms(perms: Permission[]) {
    return basicPerms("Invitation", perms, [])
  }

  function googlePerms(perms: permission.Permission[], auths: auth.Auth[]) {
    return basicPerms("Google", perms, auths)
  }

  function githubPerms(perms: permission.Permission[], auths: auth.Auth[]) {
    return basicPerms("GitHub", perms, auths)
  }

  function slackPerms(perms: permission.Permission[], auths: auth.Auth[]) {
    return basicPerms("Slack", perms, auths)
  }

  function dumpAuth(auths: auth.Auth[]) {
    if (auths.length === 0) {
      return <li>Not signed in</li>
    }
    return <li>Signed in on {auths.map(x => x.provider).join(", ")}</li>
  }

  export function renderPermissions(teamID: string | undefined, sprintID: string | undefined, perms: permission.Permission[], auths: auth.Auth[]): JSX.Element {
    const p = auths.map(x => x.provider)
    const g = collection.groupBy(perms, x => x.k);
    return <ul class="uk-list">
      {teamPerms(teamID, findGroup("team", g))}
      {sprintPerms(sprintID, findGroup("sprint", g))}
      {invitationPerms(findGroup("invitation", g))}
      {googlePerms(findGroup("google", g), auths.filter(a => a.provider == "google"))}
      {githubPerms(findGroup("github", g), auths.filter(a => a.provider == "google"))}
      {slackPerms(findGroup("slack", g), auths.filter(a => a.provider == "google"))}
      {dumpAuth(auths)}
    </ul>;
  }

  function findGroup(key: string, groups: collection.Group<Permission>[]): Permission[] {
    let ret: Permission[] = [];
    for (const g of groups) {
      if (g.key === key) {
        ret = g.members;
        break;
      }
    }
    return ret
  }
}
