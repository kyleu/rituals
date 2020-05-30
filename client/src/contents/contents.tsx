namespace contents {
  function renderSprintContent(svc: services.Service, session: session.Session): JSX.Element {
    const profile = system.cache.getProfile();
    return <tr>
      <td><a class={`${profile.linkColor}-fg`} href={`/${svc.key}/${session.slug}`}>{session.title}</a></td>
      <td class="uk-table-shrink uk-text-nowrap">{system.getMemberName(session.owner)}</td>
      <td class="uk-table-shrink uk-text-nowrap">{date.toDateTimeString(new Date(session.created))}</td>
    </tr>;
  }

  function toContent(svc: services.Service, sessions: session.Session[]) {
    return sessions.map(s => {
      return {svc: svc, session: s}
    })
  }

  export function renderContents(src: services.Service, tgt: services.Service, sessions: session.Session[]): JSX.Element {
    const contents = toContent(tgt, sessions);
    contents.sort((l, r) => (l.session.created > r.session.created ? -1 : 1));

    if (contents.length === 0) {
      return <div>{`No ${tgt.plural} in this ${src.key}`}</div>;
    } else {
      return <table class="uk-table uk-table-divider uk-text-left">
        <tbody>
        {contents.map(a => renderSprintContent(a.svc, a.session))}
        </tbody>
      </table>;
    }
  }

}
