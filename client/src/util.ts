let $: (selector: string, context?: any) => [HTMLElement] = UIkit.util.$$;

function $id(id: string): HTMLElement {
  if(id.length > 0 && !(id[0] === '#')) {
    id = "#" + id;
  }
  return $(id)[0]
}
