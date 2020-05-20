namespace contents {
  function renderSprintContent(svc: services.Service, session: rituals.Session): JSX.Element {
    const profile = system.cache.getProfile();
    return <tr>
      <td><a class={`${profile.linkColor}-fg`} href={`/${svc.key}/${session.slug}`}>{session.title}</a></td>
      <td class="uk-table-shrink uk-text-nowrap">{system.getMemberName(session.owner)}</td>
      <td class="uk-table-shrink uk-text-nowrap">{date.toDateTimeString(new Date(session.created))}</td>
    </tr>;
  }

  function toContent(svc: services.Service, sessions: rituals.Session[]) {
    return sessions.map(s => {
      return {svc: svc, session: s}
    })
  }

  export function renderContents(svc: services.Service, sessions: rituals.Session[]): JSX.Element {
    const contents = toContent(svc, sessions);
    contents.sort((l, r) => (l.session.created > r.session.created ? -1 : 1));

    if (contents.length === 0) {
      return <div>{`No ${svc.plural} in this sprint`}</div>;
    } else {
      return <table class="uk-table uk-table-divider uk-text-left">
        <tbody>
        {contents.map(a => renderSprintContent(a.svc, a.session))}
        </tbody>
      </table>;
    }
  }

}
