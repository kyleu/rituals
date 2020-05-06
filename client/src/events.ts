function delay(f: () => any) {
  setTimeout(f, 250);
}

function modalShow(key: string) {
  switch (key) {
    case "self":
      let selfInput = $req<HTMLInputElement>("#self-name-input");
      selfInput.value = $req("#member-self .member-name").innerText;
      delay(() => selfInput.focus());
      break;
    case "session":
      let sessionInput = $req<HTMLInputElement>("#model-title-input");
      sessionInput.value = $req("#model-title").innerText;
      delay(() => sessionInput.focus());
      break;
    case "invite":
      break;
    case "member":
      viewActiveMember();
      break;
    case "poll":
      let pollInput = $req<HTMLInputElement>("#poll-title-input");
      pollInput.value = "";
      delay(() => pollInput.focus());
      break;
    default:
      console.debug("unhandled modal [" + key + "]");
  }
}

