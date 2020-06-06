namespace team {
  export function renderTeamLink(tm: team.Detail, bare?: boolean) {
    const profile = system.cache.getProfile();
    const a = <a class={`${profile.linkColor}-fg`} href={`/team/${tm.slug}`}>{tm.title}</a>
    if (bare) {
      return a;
    }
    return <span>in {a}</span>;
  }

  export function renderTeamSelect(teams: ReadonlyArray<team.Detail>, activeID: string | undefined) {
    return <select class="uk-select" onchange="permission.setModelPerms('team')">
      <option value="">- no team -</option>
      { teams.map(t => {
        return t.id === activeID ? <option selected="selected" value={t.id}>{t.title}</option> : <option value={t.id}>{t.title}</option>;
      }) }
    </select>
  }
}
