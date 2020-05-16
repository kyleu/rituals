namespace profile {
  export function setNavColor(el: HTMLElement, c: string) {
    util.setValue("#navbar-color", c);
    let nb = util.req("#navbar");
    nb.className = (c + "-bg uk-navbar-container uk-navbar");
    let colors = document.querySelectorAll(".navbar_swatch");
    colors.forEach(function(i) {
      i.classList.remove("active");
    })
    el.classList.add("active");
  }

  export function setLinkColor(el: HTMLElement, c: string) {
    util.setValue("#link-color", c);
    let links = util.els(".profile-link");
    links.forEach(l => {
      l.classList.forEach(x => {
        if (x.indexOf("-fg") > -1) {
          l.classList.remove(x);
        }
        l.classList.add(c + "-fg");
      });
    });
    let colors = document.querySelectorAll(".link_swatch");
    colors.forEach(function(i) {
      i.classList.remove("active");
    })
    el.classList.add("active");
  }

  export function selectTheme(theme: string) {
    let card = util.els(".uk-card");
    switch (theme) {
      case "light":
        document.documentElement.classList.remove("uk-light");
        document.body.classList.remove("uk-light");
        document.documentElement.classList.add("uk-dark");
        document.body.classList.add("uk-dark");
        card.forEach(x => {
          x.classList.add("uk-card-default");
          x.classList.remove("uk-card-secondary");
        });
        break;
      case "dark":
        document.documentElement.classList.add("uk-light");
        document.body.classList.add("uk-light");
        document.documentElement.classList.remove("uk-dark");
        document.body.classList.remove("uk-dark");
        card.forEach(x => {
          x.classList.remove("uk-card-default");
          x.classList.add("uk-card-secondary");
        });
        break;
      default:
        console.warn("invalid theme");
        break;
    }
  }
}
