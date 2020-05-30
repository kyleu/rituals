namespace comment {
  function renderComment(comment: Comment): JSX.Element {
    return <tr>
      <td>{system.getMemberName(comment.userID)}</td>
      <td dangerouslySetInnerHTML={{__html: comment.html}} />
      <td class="uk-table-shrink uk-text-nowrap">{date.toDateTimeString(new Date(comment.created))}</td>
    </tr>;
  }

  export function renderComments(comments: Comment[]): JSX.Element {
    if (comments.length === 0) {
      return <div>No comments available</div>;
    } else {
      return <table class="uk-table uk-table-divider uk-text-left">
        <thead>
          <tr>
            <th>User</th>
            <th>Content</th>
            <th>Created</th>
          </tr>
        </thead>
        <tbody>
          {comments.map(c => renderComment(c))}
        </tbody>
      </table>;
    }
  }

  export function renderCount(k: string, v: string) {
    const profile = system.cache.getProfile();
    return <div class="comment-count-container uk-margin-small-left left hidden" data-comment-type={k} data-comment-id={v}>
      <a class={`${profile.linkColor}-fg`} title="view comments">
        <div class="comment-count">
          <span class="uk-icon" data-uk-icon="comments" />
          <span class="text" />
        </div>
      </a>
    </div>;
  }
}
