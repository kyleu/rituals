namespace retro {
  interface Detail extends rituals.Session {
    categories: string[];
  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
    feedback: feedback.Feedback[];
  }

  class Cache {
    detail?: Detail;
    feedback: feedback.Feedback[] = [];

    activeFeedback?: string;
  }

  export const cache = new Cache();

  export function onRetroMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.retro, param as string);
        break;
      case command.server.sessionJoined:
        let sj = param as SessionJoined
        rituals.onSessionJoin(sj);
        setRetroDetail(sj.session);
        feedback.setFeedback(sj.feedback);
        break;
      case command.server.sessionUpdate:
        setRetroDetail(param as Detail);
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
    const msg = {
      svc: services.retro,
      cmd: command.client.updateSession,
      param: {
        title: title,
        categories: categories,
      },
    };
    socket.send(msg);
  }
}
