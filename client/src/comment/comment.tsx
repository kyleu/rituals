namespace comment {
  function renderComment(comment: comment.Comment, profile: profile.Profile) {
    let close = comment.userID === profile.userID ? <div class="right">
      <a class={`${profile.linkColor}-fg`} data-uk-icon="close" href="" onclick={`return comment.remove('${comment.id}');`} title="remove your comment"/>
    </div> : <span />
    return <li>
      <article class="uk-comment uk-visible-toggle uk-transition-toggle" tabindex="-1">
        {member.renderHeader(member.getMember(comment.userID), comment.created, close)}
        <div class="uk-comment-body">
          <div dangerouslySetInnerHTML={{__html: comment.html}}/>
        </div>
        <hr/>
      </article>
    </li>;
  }

  export function renderComments(comments: ReadonlyArray<comment.Comment>, profile: profile.Profile) {
    if (comments.length === 0) {
      return <div>No comments available</div>;
    } else {
      return <div>{comments.map(c => renderComment(c, profile))}</div>;
    }
  }

  export function renderCount(k: string, v: string) {
    const profile = system.cache.getProfile();
    return <div class="comment-count-container uk-margin-small-left left hidden" data-comment-type={k} data-comment-id={v}>
      <a class={`${profile.linkColor}-fg`} title="view comments">
        <div class="comment-count">
          <span class="uk-icon" data-uk-icon="comment"/>
          <span class="text"/>
        </div>
      </a>
    </div>;
  }
}
