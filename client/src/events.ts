function delay(f: () => any) {
  setTimeout(f, 250);
}

function modalShow(key: string) {
  switch (key) {
    case "self":
      let input = $req<HTMLInputElement>("#self-name-input");
      input.value = $req("#member-self .member-name").innerText;
      delay(() => input.focus());
      break;
    case "session":
      break;
    case "invite":
      break;
    case "poll":
      break;
    default:
      console.debug("unhandled modal [" + key + "]");
  }
}

