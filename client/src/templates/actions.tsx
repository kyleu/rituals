namespace action {
  function renderAction(action: Action): JSX.Element {
    const c = JSON.stringify(action.content, null, 2);
    return <tr>
      <td>{system.getMemberName(action.authorID)}</td>
      <td>{action.act}</td>
      <td>{c === "null" ? "" : <pre>{c}</pre>}</td>
      <td>{action.note}</td>
      <td class="uk-table-shrink uk-text-nowrap">{date.toDateTimeString(new Date(action.created))}</td>
    </tr>;
  }

  export function renderActions(actions: Action[]): JSX.Element {
    if (actions.length === 0) {
      return <div>No actions available</div>;
    } else {
      return <table class="uk-table uk-table-divider uk-text-left">
        <thead>
          <tr>
            <th>Author</th>
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
