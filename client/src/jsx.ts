declare namespace JSX {
  interface Element extends Node {}
  interface IntrinsicElements {
    [elemName: string]: any;
  }
}

function JSX(tag: string, attrs: any) {
  const e = document.createElement(tag);
  for (const name in attrs) {
    if (name && attrs.hasOwnProperty(name)) {
      const v = attrs[name];
      if (v === true) {
        e.setAttribute(name, name);
      } else if (v !== false && v !== null && v !== undefined) {
        e.setAttribute(name, v.toString());
      }
    }
  }
  for (let i = 2; i < arguments.length; i++) {
    let child = arguments[i];
    if (Array.isArray(child)) {
      child.forEach(c => {
        e.appendChild(c);
      });
    } else {
      if (child.nodeType === null || child.nodeType === undefined) {
        child = document.createTextNode(child.toString());
      }
      e.appendChild(child);
    }
  }
  return e;
}
