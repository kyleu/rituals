function delay(f: () => any) {
  setTimeout(f, 250);
}

function openModal(key: string) {
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
  case "add-story":
    const storyInput = $req<HTMLInputElement>("#story-title-input");
    storyInput.value = "";
    delay(() => storyInput.focus());
    break;
  case "story":
    viewActiveStory();
    break;
  default:
    console.debug("unhandled modal [" + key + "]");
  }
  UIkit.modal("#modal-" + key).show();
}

