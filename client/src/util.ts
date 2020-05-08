function $<T extends HTMLElement>(selector: string, context?: HTMLElement): T[] {
  return UIkit.util.$$(selector, context) as T[];
}

function $req<T extends HTMLElement>(selector: string): T {
  const res = $<T>(selector);
  if (res.length === 0) {
    console.error("no element found for selector [" + selector + "]");
  }
  return res[0];
}

function $id<T extends HTMLElement>(id: string): T {
  if (id.length > 0 && !(id[0] === "#")) {
    id = "#" + id;
  }
  return $req(id);
}

function init(svc: string, id: string) {
  window.onbeforeunload = function () {
    appUnloading = true;
  };

  socketConnect(svc, id);
}
