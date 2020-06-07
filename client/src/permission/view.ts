namespace permission {
  export function updateView() {
    const ci = dom.req("#session-view-section .perms");
    dom.clear(ci);
    ci.appendChild(permission.renderView(permissions));
  }
}
