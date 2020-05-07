function delay(f: () => any) {
  setTimeout(f, 250);
}

function modalShow(key: string) {
  switch (key) {
    case "self":
      const selfInput = $req<HTMLInputElement>("#self-name-input");
      selfInput.value = $req("#member-self .member-name").innerText;
      delay(() => selfInput.focus());
      break;
    case "session":
      const sessionInput = $req<HTMLInputElement>("#model-title-input");
      sessionInput.value = $req("#model-title").innerText;
      delay(() => sessionInput.focus());
      break;
    case "invite":
      break;
    case "member":
      viewActiveMember();
      break;
    case "add-poll":
      const pollInput = $req<HTMLInputElement>("#poll-title-input");
      pollInput.value = "";
      delay(() => pollInput.focus());
      break;
    case "poll":
      viewActivePoll();
      break;
    default:
      console.debug("unhandled modal [" + key + "]");
  }
}

