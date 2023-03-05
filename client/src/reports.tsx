import {JSX} from "./jsx"; // eslint-disable-line @typescript-eslint/no-unused-vars
import type {Report} from "./report";
import {username} from "./member";
import {snippetCommentsModal, snippetCommentsModalLink} from "./comments";
import {expandCollapse} from "./util";

export function snippetReportContainer(day: string) {
  return <li id={"report-group-" + day}>
    <input id={"accordion-" + day} type="checkbox" hidden checked="checked"/>
    <label for={"accordion-" + day} dangerouslySetInnerHTML={expandCollapse(day)}></label>
    <div class="bd"></div>
  </li>;
}

export function snippetReport(r: Report): HTMLElement {
  return <div class="report" id={"report-" + r.id}>
    <div>
      <div class="right">{snippetCommentsModalLink("report", r.id)}</div>
      {snippetCommentsModal("report", r.id, r.id)}
      <a href={"#modal-report-" + r.id} data-id={r.id} class="clean">
        <h4 class={"username member-" + r.userID + "-name"}>{username(r.userID)}</h4>
        <div class="pt">{r.html}</div>
      </a>
    </div>
  </div>;
}

export function snippetReportModalView(r: Report): HTMLElement {
  return <div id={"modal-report-" + r.id} class="modal modal-report-view" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>{r.day}: <span class={"member-" + r.userID + "-name"}>{username(r.userID)}</span></h2>
      </div>
      <div class="modal-body" dangerouslySetInnerHTML={r.html}></div>
    </div>
  </div>;
}

export function snippetReportModalEdit(r: Report): HTMLElement {
  return <div id={"modal-report-" + r.id} class="modal modal-report-edit" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>{r.day}: <span className={"member-" + r.userID + "-name"}>{username(r.userID)}</span></h2>
      </div>
      <div class="modal-body">
        <form action="#" method="post">
          <input type="hidden" name="reportID" value={ r.id } />
          <div class="mb expanded">
            <label for={"input-day-" + r.id}><em class="title">Day</em></label>
            <div><input type="date" id={"input-day-" + r.id} name="day" value={r.day} /></div>
          </div>
          <div class="mb expanded">
            <label for={"input-content-" + r.id}><em class="title">Content</em></label>
            <div><textarea id={"input-content-" + r.id} name="content">{ r.content }</textarea></div>
          </div>
          <div class="right"><button class="report-edit-save" type="submit" name="action" value="child-update">Save Changes</button></div>
          <button class="report-edit-delete" type="submit" name="action" value="child-remove" onclick="return confirm('Are you sure you want to delete this report?');">Delete</button>
        </form>
      </div>
    </div>
  </div>;
}
