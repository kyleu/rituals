declare const UIkit: any;

namespace dom {
  export function initDom(t: string, color: string) {
    try {
      style.themeLinks(color);
      style.setTheme(t);
    } catch (e) {
      console.warn("error setting style", e);
    }
    document.body.style.visibility = "visible";
    try {
      modal.wire();
    } catch (e) {
      console.warn("error wiring modals", e);
    }
  }

  export function els<T extends HTMLElement>(selector: string, context?: HTMLElement): T[] {
    return UIkit.util.$$(selector, context) as T[];
  }

  export function opt<T extends HTMLElement>(selector: string, context?: HTMLElement): T | undefined {
    return els<T>(selector, context).shift();
  }

  export function req<T extends HTMLElement>(selector: string, context?: HTMLElement): T {
    const res = opt<T>(selector, context);
    if (!res) {
      console.warn(`no element found for selector [${selector}]`);
    }
    return res!;
  }

  export function setHTML(el: string | HTMLElement, html: string) {
    if (typeof el === "string") {
      el = req(el);
    }
    el.innerHTML = html;
    return el;
  }

  export function setDisplay(el: string | HTMLElement, condition: boolean, v: string = "block") {
    if (typeof el === "string") {
      el = req(el);
    }

    el.style.display = condition ? v : "none";
    return el;
  }

  export function setContent(el: string | HTMLElement, e: JSX.Element) {
    if (typeof el === "string") {
      el = req(el);
    }
    el.innerHTML = "";
    el.appendChild(e);
    return el;
  }

  export function setText(el: string | HTMLElement, text: string): HTMLElement {
    if (typeof el === "string") {
      el = req(el);
    }
    el.innerText = text;
    return el;
  }
}
