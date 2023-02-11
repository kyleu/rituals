import {JSX} from "./jsx"; // eslint-disable-line @typescript-eslint/no-unused-vars
import {Feedback} from "./feedback";
import {username} from "./member";
import {snippetCommentsModal, snippetCommentsModalLink} from "./comments";

export function snippetFeedbackContainer(cat: string) {
  return <div id={"category-" + cat} data-category={cat} class="category">
    <div class="right"><a class="add-feedback-link" data-category={cat} href={"#modal-feedback--add-" + cat}>
      <button>New</button>
    </a></div>
    <h4><a href={"#modal-feedback--add-" + cat}>{cat}</a></h4>
    <div class="clear"></div>
    <div class="feedback-list"></div>
    <div id={"modal-feedback--add-" + cat} class="modal modal-feedback-add" data-category={cat} style="display: none;">
      <a class="backdrop" href="#"></a>
      <div class="modal-content">
        <div class="modal-header">
          <a href="#" class="modal-close">×</a>
          <h2>New Feedback</h2>
        </div>
        <div class="modal-body">
          <form action="#" method="post">
            <input type="hidden" name="action" value="child-add"/>
            <input type="hidden" name="action" value="child-add"/>
            <div class="mb expanded">
              <label for={"category-" + cat + "-select"}><em class="title">Category</em></label>
              <div>
                <select name="category" id={"category-" + cat + "-select"}>
                  <option selected="selected" value="TODO">TODO</option>
                </select>
              </div>
            </div>
            <div class="mb expanded">
              <label for={"feedback-add-content-" + cat}><em class="title">Content</em></label>
              <div><textarea rows="8" id={"feedback-add-content-" + cat} name="content" placeholder="HTML and Markdown supported"></textarea></div>
            </div>
            <button type="submit">Add Feedback</button>
          </form>
        </div>
      </div>
    </div>
  </div>;
}

export function snippetFeedback(f: Feedback) {
  return <div id={"feedback-" + f.id} class="feedback mt clear">
    <div class="right">{snippetCommentsModalLink("feedback", f.id)}</div>
    {snippetCommentsModal("feedback", f.id, f.id)}
    <a href={"#modal-feedback-" + f.id} class="clean modal-feedback-edit-link" data-id={f.id}>
      <div>{username(f.userID)}</div>
      <div class="pt feedback-content">{f.html}</div>
    </a>
  </div>;
}
