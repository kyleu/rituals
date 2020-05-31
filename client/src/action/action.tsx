namespace action {
  function renderAction(action: Action) {
    const c = JSON.stringify(action.content, null, 2);
    return <tr>
      <td>{member.renderTitle(member.getMember(action.userID))}</td>
      <td>{action.act}</td>
      <td>{c === "null" ? "" : <pre>{c}</pre>}</td>
      <td>{action.note}</td>
      <td class="uk-table-shrink uk-text-nowrap">{date.toDateTimeString(new Date(action.created))}</td>
    </tr>;
  }

  export function renderActions(actions: ReadonlyArray<Action>) {
    if (actions.length === 0) {
      return <div>No actions available</div>;
    } else {
      return <table class="uk-table uk-table-divider uk-text-left">
        <thead>
          <tr>
            <th>User</th>
            <th>Act</th>
            <th>Content</th>
            <th>Note</th>
            <th>Created</th>
          </tr>
        </thead>
        <tbody>
          {actions.map(a => renderAction(a))}
        </tbody>
      </table>;
    }
  }
}
