import {JSX} from "./jsx"; // eslint-disable-line @typescript-eslint/no-unused-vars
import {Comment} from "./comment";
import {relativeTime, utc} from "./time";

export function snippetComment(c: Comment, un: string) {
  const li = <li></li>;

  const d = utc(new Date());
  const timeSpan = <span class="nowrap reltime" data-time={d}>just now</span>;
  relativeTime(d, timeSpan);

  const time = <div class="right"></div>;
  time.appendChild(timeSpan);
  li.appendChild(time);

  li.appendChild(<div>{ c.content }</div>);
  li.appendChild(<div><em>{ un }</em></div>);
  return li;
}
