namespace events {
  function delay(f: () => any) {
    setTimeout(f, 250);
  }

  export function openModal(key: string, id?: string) {
    switch (key) {
      case "session":
        const sessionInput = util.setValue("#model-title-input", util.req("#model-title").innerText);
        delay(() => sessionInput.focus());
        break;

      // member
      case "self":
        const selfInput = util.setValue("#self-name-input", util.req("#member-self .member-name").innerText);
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
        const storyInput = util.setValue("#story-title-input", "");
        delay(() => storyInput.focus());
        break;
      case "story":
        estimate.cache.activeStory = id;
        story.viewActiveStory();
        break;

      // standup
      case "add-report":
        util.setValue("#standup-report-date", dateToYMD(new Date()));
        const reportContent = util.setValue("#standup-report-content", "");
        util.wireTextarea(reportContent as HTMLTextAreaElement);
        delay(() => reportContent.focus());
        break;
      case "report":
        standup.cache.activeReport = id;
        report.viewActiveReport();
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
