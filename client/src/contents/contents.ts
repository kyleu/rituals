namespace contents {
  export function onContentDisplay(key: string, same: boolean, content: string, html: string) {
    dom.setDisplay(`#modal-${key} .content-edit`, same);
    dom.setDisplay(`#modal-${key} .buttons-edit`, same);

    const v = dom.setDisplay(`#modal-${key} .content-view`, !same);
    dom.setHTML(v, same ? "" : html);
    dom.setDisplay(`#modal-${key} .buttons-view`, !same);

    const contentEditTextarea = dom.req<HTMLTextAreaElement>(`#${key}-edit-content`);
    dom.setValue(contentEditTextarea, same ? content : "");
    if (same) {
      dom.wireTextarea(contentEditTextarea);
    }
  }
}