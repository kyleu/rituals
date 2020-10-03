namespace retro {
  interface Detail extends session.Session {
    readonly categories: string[];
  }

  interface SessionJoined extends session.SessionJoined {
    readonly session: Detail;
    readonly team?: team.Detail;
    readonly sprint?: sprint.Detail;
    readonly feedback: feedback.Feedback[];
  }

  class Cache {
    activeFeedback?: string;

    detail?: Detail;
    sprint?: sprint.Detail;

    feedback: feedback.Feedback[] = [];
  }

  export const cache = new Cache();

  export function onRetroMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.retro.key, param as string);
        break;
      case command.server.sessionJoined:
        const sj = param as SessionJoined;
        session.onSessionJoin(sj);
        setRetroDetail(sj.session);
        feedback.setFeedback(sj.feedback);
        session.showWelcomeMessage(sj.session.status, sj.members.length);
        break;
      case command.server.sessionUpdate:
        setRetroDetail(param as Detail);
        break;
      case command.server.sessionRemove:
        system.onSessionRemove(services.retro);
        break;
      case command.server.permissionsUpdate:
        system.setPermissions(param as permission.Permission[]);
        break;
      case command.server.teamUpdate:
        const tm = param as team.Detail | undefined;
        if (retro.cache.detail) {
          retro.cache.detail.teamID = tm?.id;
        }
        session.setTeam(tm);
        break;
      case command.server.sprintUpdate:
        const spr = param as sprint.Detail | undefined;
        if (retro.cache.detail) {
          retro.cache.detail.sprintID = spr?.id;
        }
        session.setSprint(spr);
        break;
      case command.server.feedbackUpdate:
        feedback.onFeedbackUpdate(param as feedback.Feedback);
        break;
      case command.server.feedbackRemove:
        feedback.onFeedbackRemoved(param as string);
        break;
      default:
        console.warn(`unhandled command [${cmd}] for retro`);
    }
  }

  function setRetroDetail(detail: Detail) {
    cache.detail = detail;
    const cs = detail.categories.join(", ");
    dom.setValue("#model-categories-input", cs);
    dom.setContent("#session-view-section .categories", tags.renderTagsView(detail.categories));
    dom.setOptions("#feedback-category", detail.categories);
    dom.setOptions("#feedback-edit-category", detail.categories);
    feedback.setFeedback(retro.cache.feedback);
    session.setDetail(detail);
  }

  export function onSubmitRetroSession() {
    const title = dom.req<HTMLInputElement>("#model-title-input").value;
    const categories = dom.req<HTMLInputElement>("#model-categories-input").value;
    const teamID = dom.req<HTMLSelectElement>("#model-team-select select").value;
    const sprintID = dom.req<HTMLSelectElement>("#model-sprint-select select").value;
    const permissions = permission.readPermissions();

    const msg = { svc: services.retro.key, cmd: command.client.updateSession, param: { title, categories, teamID, sprintID, permissions } };
    socket.send(msg);
  }
}
