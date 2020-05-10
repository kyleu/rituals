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
      console.error("no element found for selector [" + selector + "]");
    }
    return res!;
  }

  export function setContent(path: string, el: JSX.Element) {
    const detail = util.req(path);
    detail.innerHTML = "";
    detail.appendChild(el);
  }
}
