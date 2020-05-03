let $: (selector: string, context?: any) => [HTMLElement] = UIkit.util.$$;

function $id(id: string): HTMLElement {
  if(id.length > 0 && !(id[0] === '#')) {
    id = "#" + id;
  }
  return $(id)[0]
}

let appInitialized = false;
let appUnloading = false;

function init(svc: string, id: any) {
  appInitialized = true;

  window.onbeforeunload = function() {
    appUnloading = true;
  };

  socketConnect(svc, id);
}
