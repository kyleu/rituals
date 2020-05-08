namespace events {
  function delay(f: () => any) {
    setTimeout(f, 250);
  }

  export function openModal(key: string, id?: string) {
    switch (key) {
      case "self":
        const selfInput = util.req<HTMLInputElement>("#self-name-input");
        selfInput.value = util.req("#member-self .member-name").innerText;
        delay(() => selfInput.focus());
        break;
      case "session":
        const sessionInput = util.req<HTMLInputElement>("#model-title-input");
        sessionInput.value = util.req("#model-title").innerText;
        delay(() => sessionInput.focus());
        break;
      case "invite":
        break;
      case "member":
        system.cache.activeMember = id;
        member.viewActiveMember();
        break;
      case "add-story":
        const storyInput = util.req<HTMLInputElement>("#story-title-input");
        storyInput.value = "";
        delay(() => storyInput.focus());
        break;
      case "story":
        estimate.cache.activeStory = id;
        story.viewActiveStory();
        break;
      default:
        console.debug("unhandled modal [" + key + "]");
    }
    UIkit.modal("#modal-" + key).show();
    return false;
  }
}
