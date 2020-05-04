function $<T extends HTMLElement>(selector: string, context?: any): T[] {
  return UIkit.util.$$(selector, context) as T[];
};

function $req<T extends HTMLElement>(selector: string): T {
  let res = $<T>(selector)
  if (res.length == 0) {
    console.error("no element found for selector [" + selector + "]")
  }
  return res[0];
}

function $id(id: string): HTMLElement {
  if (id.length > 0 && !(id[0] === '#')) {
    id = "#" + id;
  }
  return $req(id);
}

let appInitialized = false;
let appUnloading = false;

function init(svc: string, id: any) {
  appInitialized = true;

  window.onbeforeunload = function () {
    appUnloading = true;
  };

  socketConnect(svc, id);
}
