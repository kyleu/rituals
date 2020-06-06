namespace member {
  export function memberUpdateDom(nameChanged: boolean) {
    if (nameChanged) {
      switch (system.cache.currentService) {
        case services.team:
          break;
        case services.sprint:
          break;
        case services.estimate:
          if (estimate.cache.activeStory) {
            vote.viewVotes();
          }
          break;
        case services.standup:
          dom.setContent("#report-detail", report.renderReports(standup.cache.reports));
          if (standup.cache.activeReport) {
            report.viewActiveReport();
          }
          break;
        case services.retro:
          dom.setContent("#feedback-detail", feedback.renderFeedbackArray(retro.cache.feedback));
          if (retro.cache.activeFeedback) {
            feedback.viewActiveFeedback();
          }
          break;
      }
    }
  }

  export function activeMemberDom(member: member.Member) {
    const owner = selfCanEdit();
    dom.setDisplay("#modal-member .owner-form", owner);
    dom.setDisplay("#modal-member .member-form", !owner);
    dom.setDisplay("#modal-member .owner-actions", owner);
    dom.setDisplay("#modal-member .member-actions", !owner);
    dom.setSelectOption("#member-modal-role-select", member.role);
    dom.setText("#member-modal-name", member.name);
    dom.setText("#member-modal-role", member.role);
  }
}
