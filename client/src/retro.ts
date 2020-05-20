namespace retro {
  interface Detail extends rituals.Session {
    status: { key: string };
    categories: string[];
  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
    team?: team.Detail;
    sprint?: sprint.Detail;
    feedback: feedback.Feedback[];
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
        let sj = param as SessionJoined;
        rituals.onSessionJoin(sj);
        rituals.setTeam(sj.team);
        rituals.setSprint(sj.sprint);
        setRetroDetail(sj.session);
        feedback.setFeedback(sj.feedback);
        break;
      case command.server.sessionUpdate:
        setRetroDetail(param as Detail);
        break;
      case command.server.teamUpdate:
        const tm = param as team.Detail | undefined;
        if (retro.cache.detail) {
          retro.cache.detail.teamID = tm?.id;
        }
        rituals.setTeam(tm);
        break;
      case command.server.sprintUpdate:
        const spr = param as sprint.Detail | undefined;
        if (retro.cache.detail) {
          retro.cache.detail.sprintID = spr?.id;
        }
        rituals.setSprint(spr);
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
    dom.setValue("#model-categories-input", detail.categories.join(", "));
    dom.setOptions(dom.req("#retro-feedback-category"), detail.categories);
    dom.setOptions(dom.req("#retro-feedback-edit-category"), detail.categories);
    feedback.setFeedback(retro.cache.feedback);
    rituals.setDetail(detail);
  }

  export function onSubmitRetroSession() {
    const title = dom.req<HTMLInputElement>("#model-title-input").value;
    const categories = dom.req<HTMLInputElement>("#model-categories-input").value;
    const teamID = dom.req<HTMLSelectElement>("#model-team-select select").value;
    const sprintID = dom.req<HTMLSelectElement>("#model-sprint-select select").value;

    const msg = {svc: services.retro.key, cmd: command.client.updateSession, param: {title: title, categories: categories, teamID: teamID, sprintID: sprintID}};
    socket.send(msg);
  }
}
