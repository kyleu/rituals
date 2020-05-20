namespace team {
  export function renderTeamLink(tm: team.Detail) {
    const profile = system.cache.getProfile();
    return <span> in <a class={`${profile.linkColor}-fg`} href={`/team/${tm.slug}`}>{tm.title}</a></span>
  }

  export function renderTeamSelect(teams: team.Detail[], activeID: string | undefined) {
    return <select class="uk-select">
      <option value="">- no team -</option>
      { teams.map(t => {
        return t.id === activeID ? <option selected="selected" value={t.id}>{t.title}</option> : <option value={t.id}>{t.title}</option>;
      }) }
    </select>
  }
}
