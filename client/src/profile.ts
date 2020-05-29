namespace profile {
  // noinspection JSUnusedGlobalSymbols
  export function setNavColor(el: HTMLElement, c: string) {
    dom.setValue("#nav-color", c);
    const nb = dom.req("#navbar");
    nb.className = `${c}-bg uk-navbar-container uk-navbar`;
    const colors = document.querySelectorAll(".nav_swatch");
    colors.forEach(function(i) {
      i.classList.remove("active");
    });
    el.classList.add("active");
  }

  // noinspection JSUnusedGlobalSymbols
  export function setLinkColor(el: HTMLElement, c: string) {
    dom.setValue("#link-color", c);
    const links = dom.els(".profile-link");
    links.forEach(l => {
      l.classList.forEach(x => {
        if (x.indexOf("-fg") > -1) {
          l.classList.remove(x);
        }
        l.classList.add(`${c}-fg`);
      });
    });
    const colors = document.querySelectorAll(".link_swatch");
    colors.forEach(function(i) {
      i.classList.remove("active");
    });
    el.classList.add("active");
  }
}
