import {JSX} from "./jsx"; // eslint-disable-line @typescript-eslint/no-unused-vars
import {Comment, initCommentsLink} from "./comment";
import {relativeTime, utc} from "./time";
import {svg} from "./util";

export function snippetComment(c: Comment, un: string) {
  const li = <li></li>;

  const d = utc(new Date());
  const timeSpan = <span class="nowrap reltime" data-time={d}>just now</span>;
  relativeTime(d, timeSpan);

  const time = <div class="right"></div>;
  time.appendChild(timeSpan);
  li.appendChild(time);

  li.appendChild(<div>{c.content}</div>);
  li.appendChild(<div><em>{un}</em></div>);
  return li;
}

export function snippetCommentsModalLink(svc: string, id: string) {
  const x = "comment-link-" + svc + "-" + id;
  const href = "#modal-" + svc + "-" + id + "-comments";
  const a = <a id={x} class="comment-link" data-key={svc + "-" + id} href={href} title="0 comment" dangerouslySetInnerHTML={svg("comment-alt")}></a>;
  initCommentsLink(a);
  return a;
}

export function snippetCommentsModal(svc: string, id: string, title: string) {
  return <div id={"modal-" + svc + "-" + id + "-comments"} class="modal comments" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>{title} Comments</h2>
      </div>
      <div class="modal-body">
        <ul id={"comment-list-" + svc + "-" + id} class="comment-list">
        </ul>
        <form action="#" method="post" class="expanded">
          <input type="hidden" name="action" value="comment"/>
          <input type="hidden" name="svc" value={svc}/>
          <input type="hidden" name="modelID" value={id}/>
          <div><textarea name="content" placeholder="Add a comment, Markdown and HTML supported"></textarea></div>
          <div class="mt right">
            <button type="submit">Add Comment</button>
          </div>
        </form>
      </div>
    </div>
  </div>
}
