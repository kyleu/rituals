namespace style {
  export function setTheme(theme: string) {
    const card = dom.els(".uk-card");
    switch (theme) {
      case "default":
        if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
          setTheme("dark");
        } else {
          setTheme("light");
        }
        break;
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
