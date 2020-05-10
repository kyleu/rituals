namespace events {
  function delay(f: () => any) {
    setTimeout(f, 250);
  }

  export function openModal(key: string, id?: string) {
    switch (key) {
      case "session":
        const sessionInput = util.req<HTMLInputElement>("#model-title-input");
        sessionInput.value = util.req("#model-title").innerText;
        delay(() => sessionInput.focus());
        break;

      // member
      case "self":
        const selfInput = util.req<HTMLInputElement>("#self-name-input");
        selfInput.value = util.req("#member-self .member-name").innerText;
        delay(() => selfInput.focus());
        break;
      case "invite":
        break;
      case "member":
        system.cache.activeMember = id;
        member.viewActiveMember();
        break;

      // estimate
      case "add-story":
        const storyInput = util.req<HTMLInputElement>("#story-title-input");
        storyInput.value = "";
        delay(() => storyInput.focus());
        break;
      case "story":
        estimate.cache.activeStory = id;
        story.viewActiveStory();
        break;

      // standup
      case "add-update":
        const updateDate = util.req<HTMLInputElement>("#standup-update-date");
        updateDate.value = dateToYMD(new Date());
        const updateInput = util.req<HTMLInputElement>("#standup-update-input");
        updateInput.value = "";
        delay(() => updateInput.focus());
        break;
      case "update":
        standup.cache.activeUpdate = id;
        update.viewActiveUpdate();
        break;

      // default
      default:
        console.debug("unhandled modal [" + key + "]");
    }
    UIkit.modal("#modal-" + key).show();
    return false;
  }

  function dateToYMD(date: Date) {
    var d = date.getDate();
    var m = date.getMonth() + 1;
    var y = date.getFullYear();
    return '' + y + '-' + (m <= 9 ? '0' + m : m) + '-' + (d <= 9 ? '0' + d : d);
  }
}
