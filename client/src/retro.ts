namespace retro {
  interface Detail extends rituals.Session {
    status: { key: string };
    categories: string[];
  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
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
        let sj = param as SessionJoined
        rituals.onSessionJoin(sj);
        rituals.setSprint(sj.sprint)
        setRetroDetail(sj.session);
        feedback.setFeedback(sj.feedback);
        break;
      case command.server.sessionUpdate:
        setRetroDetail(param as Detail);
        break;
      case command.server.sprintUpdate:
        const x = param as sprint.Detail | undefined
        if (retro.cache.detail) {
          retro.cache.detail.sprintID = x?.id;
        }
        rituals.setSprint(x)
        break;
      case command.server.feedbackUpdate:
        feedback.onFeedbackUpdate(param as feedback.Feedback);
        break;
      case command.server.feedbackRemove:
        feedback.onFeedbackRemoved(param as string);
        break;
      default:
        console.warn("unhandled command [" + cmd + "] for retro");
    }
  }

  function setRetroDetail(detail: Detail) {
    cache.detail = detail;
    util.setValue("#model-categories-input", detail.categories.join(", "));
    util.setOptions(util.req("#retro-feedback-category"), detail.categories)
    util.setOptions(util.req("#retro-feedback-edit-category"), detail.categories)
    feedback.setFeedback(retro.cache.feedback);
    rituals.setDetail(detail);
  }

  export function onSubmitRetroSession() {
    const title = util.req<HTMLInputElement>("#model-title-input").value;
    const categories = util.req<HTMLInputElement>("#model-categories-input").value;
    const sprintID = util.req<HTMLSelectElement>("#model-sprint-select select").value;

    const msg = {svc: services.retro.key, cmd: command.client.updateSession, param: {title: title, categories: categories, sprintID: sprintID}};
    socket.send(msg);
  }
}
