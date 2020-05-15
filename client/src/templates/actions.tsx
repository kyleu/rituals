namespace action {
  function renderAction(action: Action): JSX.Element {
    return <tr>
      <td>{getMemberName(action.authorID)}</td>
      <td>{action.act}</td>
      <td><pre>{JSON.stringify(action.content, null, 2)}</pre></td>
      <td>{action.note}</td>
      <td>{new Date(action.occurred).toLocaleDateString()} {new Date(action.occurred).toLocaleTimeString().slice(0, 8)}</td>
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
            <th>Occurred</th>
          </tr>
        </thead>
        <tbody>
          {actions.map(a => renderAction(a))}
        </tbody>
      </table>;
    }
  }

  export function getMemberName(id: string) {
    const ret = system.cache.members.filter(m => m.userID === id);
    if(ret.length === 0) {
      return id;
    }
    return ret[0].name;
  }
}
