namespace contents {
  export function onContentDisplay(key: string, same: boolean, content: string, html: string) {
    dom.setDisplay(`#modal-${key} .content-edit`, same);
    dom.setDisplay(`#modal-${key} .buttons-edit`, same);

    dom.setHTML(dom.setDisplay(`#modal-${key} .content-view`, !same), !same ? "" : html);
    dom.setDisplay(`#modal-${key} .buttons-view`, !same);

    const contentEditTextarea = dom.req<HTMLTextAreaElement>(`#${key}-edit-content`);
    dom.setValue(contentEditTextarea, same ? content : "");
    if (same) {
      dom.wireTextarea(contentEditTextarea);
    }
  }
}
