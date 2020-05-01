function nonNull(val: any, fallback: any) {
  return Boolean(val) ? val : fallback;
}

declare namespace JSX {
  interface Element { }
  interface IntrinsicElements {
    [elemName: string]: any;
  }
}

function JSXparseChildren(children: any[]) {
  return children.map(child => {
    if(typeof child === 'string') {
      return document.createTextNode(child);
    }
    return child;
  })
}

function JSXparseNode(element: any, properties: any, children: any[]) {
  const el = document.createElement(element);
  Object.keys(nonNull(properties, {})).forEach(key => {
    el[key] = properties[key];
  })
  JSXparseChildren(children).forEach(child => {
    el.appendChild(child);
  });
  return el;
}

function JSX(element: any, properties: any, ...children: any[][]) {
  if(typeof element === 'function') {
    return element({
      ...nonNull(properties, {}),
      children
    });
  }
  return JSXparseNode(element, properties, children);
}
