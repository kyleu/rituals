declare var UIkit: any;

namespace util {
  export function els<T extends HTMLElement>(selector: string, context?: HTMLElement): T[] {
    return UIkit.util.$$(selector, context) as T[];
  }

  export function opt<T extends HTMLElement>(selector: string, context?: HTMLElement): T | null {
    const res = util.els<T>(selector, context);
    if (res.length === 0) {
      return null;
    }
    return res[0];
  }

  export function req<T extends HTMLElement>(selector: string, context?: HTMLElement): T {
    const res = util.opt<T>(selector, context);
    if (res === null) {
      console.warn("no element found for selector [" + selector + "]");
    }
    return res!;
  }

  export function setHTML(el: string | HTMLElement, html: string) {
    if (typeof el === "string") {
      el = util.req(el);
    }
    el.innerHTML = html;
    return el;
  }

  export function setContent(el: string | HTMLElement, e: JSX.Element) {
    if (typeof el === "string") {
      el = util.req(el);
    }
    el.innerHTML = "";
    el.appendChild(e);
    return el;
  }

  export function setText(el: string | HTMLElement, text: string): HTMLElement {
    if (typeof el === "string") {
      el = util.req(el);
    }
    el.innerText = text;
    return el;
  }

  export function setValue(el: string | HTMLInputElement | HTMLTextAreaElement, text: string): HTMLInputElement | HTMLTextAreaElement {
    if (typeof el === "string") {
      el = util.req<HTMLInputElement>(el);
    }
    el.value = text;
    return el;
  }

  export function wireTextarea(text: HTMLTextAreaElement) {
    function resize() {
      text.style.height = 'auto';
      text.style.height = (text.scrollHeight < 64 ? 64 : (text.scrollHeight + 6)) + 'px';
    }

    function delayedResize() {
      window.setTimeout(resize, 0);
    }

    const x = text.dataset["autoresize"];
    if(x === undefined) {
      text.dataset["autoresize"] = "true";

      text.addEventListener('change', resize, false);
      text.addEventListener('cut', delayedResize, false);
      text.addEventListener('paste', delayedResize, false);
      text.addEventListener('drop', delayedResize, false);
      text.addEventListener('keydown', delayedResize, false);

      text.focus();
      text.select();
    }

    resize();
  }

  export function dateToYMD(date: Date) {
    var d = date.getDate();
    var m = date.getMonth() + 1;
    var y = date.getFullYear();
    return '' + y + '-' + (m <= 9 ? '0' + m : m) + '-' + (d <= 9 ? '0' + d : d);
  }

  export function dateFromYMD(s: string) {
    let d = new Date(s);
    d = new Date(d.getTime() + (d.getTimezoneOffset() * 60000));
    return d;
  }

  export function dow(i: number) {
    switch (i) {
      case 0:
        return "Sun";
      case 1:
        return "Mon";
      case 2:
        return "Tue";
      case 3:
        return "Wed";
      case 4:
        return "Thu";
      case 5:
        return "Fri";
      case 6:
        return "Sat";
      default:
        return "???";
    }
  }

  export function setOptions(el: HTMLSelectElement, categories: string[]) {
    el.innerHTML = ""
    for(const c of categories) {
      const opt = document.createElement("option");
      opt.value = c;
      opt.innerText = c
      el.appendChild(opt)
    }
  }

  export function setSelectOption(el: HTMLSelectElement, o: string | undefined) {
    for(let i = 0; i < el.children.length; i ++) {
      const e = el.children.item(i) as HTMLOptionElement;
      e.selected = e.value === o;
    }
  }
}
