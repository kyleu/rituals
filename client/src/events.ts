namespace events {
  function delay(f: () => any) {
    setTimeout(f, 250);
  }

  export function openModal(key: string, id?: string) {
    switch (key) {
      case "session":
        const sessionInput = dom.setValue("#model-title-input", dom.req("#model-title").innerText);
        delay(() => sessionInput.focus());
        team.refreshTeams();
        sprint.refreshSprints();
        break;

      // member
      case "self":
        const selfInput = dom.setValue("#self-name-input", dom.req("#member-self .member-name").innerText);
        delay(() => selfInput.focus());
        break;
      case "invitation":
        break;
      case "member":
        system.cache.activeMember = id;
        member.viewActiveMember();
        break;
      case "welcome":
        break;

      // actions
      case "actions":
        action.loadActions();
        break;

      // estimate
      case "add-story":
        const storyInput = dom.setValue("#story-title-input", "");
        delay(() => storyInput.focus());
        break;
      case "story":
        estimate.cache.activeStory = id;
        story.viewActiveStory();
        break;

      // standup
      case "add-report":
        dom.setValue("#report-date", date.dateToYMD(new Date()));
        const reportContent = dom.setValue("#report-content", "");
        dom.wireTextarea(reportContent as HTMLTextAreaElement);
        delay(() => reportContent.focus());
        break;
      case "report":
        standup.cache.activeReport = id;
        report.viewActiveReport();
        const reportEditContent = dom.req("#report-edit-content");
        delay(() => {
          dom.wireTextarea(reportEditContent as HTMLTextAreaElement);
          reportEditContent.focus();
        });
        break;

      // retro
      case "add-feedback":
        dom.setSelectOption("#feedback-category", id);
        const feedbackContent = dom.setValue("#feedback-content", "");
        dom.wireTextarea(feedbackContent as HTMLTextAreaElement);
        delay(() => feedbackContent.focus());
        break;
      case "feedback":
        retro.cache.activeFeedback = id;
        feedback.viewActiveFeedback();
        const feedbackEditContent = dom.req("#feedback-edit-content");
        delay(() => {
          dom.wireTextarea(feedbackEditContent as HTMLTextAreaElement);
          feedbackEditContent.focus();
        });
        break;

      // default
      default:
        console.warn(`unhandled modal [${key}]`);
    }
    UIkit.modal(`#modal-${key}`).show();
    return false;
  }

  export function hideModal(key: string) {
    UIkit.modal(`#modal-${key}`).hide();
  }
}
