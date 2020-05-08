declare var UIkit: any;

namespace util {
  export function els<T extends HTMLElement>(selector: string, context?: HTMLElement): T[] {
    return UIkit.util.$$(selector, context) as T[];
  }

  export function req<T extends HTMLElement>(selector: string, context?: HTMLElement): T {
    const res = util.els<T>(selector, context);
    if (res.length === 0) {
      console.error("no element found for selector [" + selector + "]");
    }
    return res[0];
  }
}
