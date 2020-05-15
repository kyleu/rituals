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

      // actions
      case "actions":
        action.loadActions();
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
        util.setValue("#standup-report-date", util.dateToYMD(new Date()));
        const reportContent = util.setValue("#standup-report-content", "");
        util.wireTextarea(reportContent as HTMLTextAreaElement);
        delay(() => reportContent.focus());
        break;
      case "report":
        standup.cache.activeReport = id;
        report.viewActiveReport();
        const reportEditContent = util.req("#standup-report-edit-content");
        delay(() => {
          util.wireTextarea(reportEditContent as HTMLTextAreaElement);
          reportEditContent.focus();
        });
        break;

      // retro
      case "add-feedback":
        util.setSelectOption(util.req("#retro-feedback-category"), id);
        const feedbackContent = util.setValue("#retro-feedback-content", "");
        util.wireTextarea(feedbackContent as HTMLTextAreaElement);
        delay(() => feedbackContent.focus());
        break;
      case "feedback":
        retro.cache.activeFeedback = id;
        feedback.viewActiveFeedback();
        const feedbackEditContent = util.req("#retro-feedback-edit-content");
        delay(() => {
          util.wireTextarea(feedbackEditContent as HTMLTextAreaElement);
          feedbackEditContent.focus();
        });
        break;

      // default
      default:
        console.debug("unhandled modal [" + key + "]");
    }
    UIkit.modal("#modal-" + key).show();
    return false;
  }
}
