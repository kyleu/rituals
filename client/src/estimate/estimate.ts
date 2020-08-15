namespace estimate {
  interface Detail extends session.Session {
    readonly choices: readonly string[];
  }

  export interface StoryStatusChange {
    readonly storyID: string;
    readonly status: string;
    readonly finalVote: string;
  }

  interface SessionJoined extends session.SessionJoined {
    readonly session: Detail;
    readonly team?: team.Detail;
    readonly sprint?: sprint.Detail;
    readonly stories: story.Story[];
    readonly votes: vote.Vote[];
  }

  class Cache {
    activeStory?: string;

    detail?: Detail;
    sprint?: sprint.Detail;

    stories: story.Story[] = [];
    votes: vote.Vote[] = [];

    public activeVotes(): vote.Vote[] {
      if (!this.activeStory) {
        return [];
      }
      return this.votes.filter(x => x.storyID === this.activeStory);
    }
  }

  export const cache = new Cache();

  export function onEstimateMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.estimate, param as string);
        break;
      case command.server.sessionJoined:
        const sj = param as SessionJoined;
        session.onSessionJoin(sj);
        setEstimateDetail(sj.session);
        story.setStories(sj.stories);
        vote.setVotes(sj.votes);
        session.showWelcomeMessage(sj.session.status, sj.members.length);
        break;
      case command.server.sessionUpdate:
        setEstimateDetail(param as Detail);
        break;
      case command.server.sessionRemove:
        system.onSessionRemove(services.estimate);
        break;
      case command.server.permissionsUpdate:
        system.setPermissions(param as permission.Permission[]);
        break;
      case command.server.teamUpdate:
        const tm = param as team.Detail | undefined;
        if (estimate.cache.detail) {
          estimate.cache.detail.teamID = tm?.id;
        }
        session.setTeam(tm);
        break;
      case command.server.sprintUpdate:
        const spr = param as sprint.Detail | undefined;
        if (estimate.cache.detail) {
          estimate.cache.detail.sprintID = spr?.id;
        }
        session.setSprint(spr);
        break;
      case command.server.storyUpdate:
        onStoryUpdate(param as story.Story);
        break;
      case command.server.storyRemove:
        onStoryRemove(param as string);
        break;
      case command.server.storyStatusChange:
        story.onStoryStatusChange(param as StoryStatusChange);
        break;
      case command.server.voteUpdate:
        vote.onVoteUpdate(param as vote.Vote);
        break;
      default:
        console.warn(`unhandled command [${cmd}] for estimate`);
    }
  }

  function setEstimateDetail(detail: Detail) {
    cache.detail = detail;
    const cs = detail.choices.join(", ");
    dom.setValue("#model-choices-input", cs);
    dom.setContent("#session-view-section .choices", tags.renderTagsView(detail.choices));
    story.viewActiveStory();
    session.setDetail(detail);
  }

  export function onSubmitEstimateSession() {
    const title = dom.req<HTMLInputElement>("#model-title-input").value;
    const choices = dom.req<HTMLInputElement>("#model-choices-input").value;
    const teamID = dom.req<HTMLSelectElement>("#model-team-select select").value;
    const sprintID = dom.req<HTMLSelectElement>("#model-sprint-select select").value;
    const permissions = permission.readPermissions();

    const msg = { svc: services.estimate.key, cmd: command.client.updateSession, param: { title, choices, teamID, sprintID, permissions } };
    socket.send(msg);
  }

  export function onStoryUpdate(s: story.Story) {
    const x = preUpdate(s.id);
    x.push(s);
    x.sort(s => s.idx);
    if (s.id === estimate.cache.activeStory) {
      dom.setText("#story-title", s.title);
    }
    story.setStories(x);
  }

  export function onStoryRemove(id: string) {
    const x = preUpdate(id);
    story.setStories(x);
    if (id === estimate.cache.activeStory) {
      modal.hide("story");
    }
    notify.notify("story has been deleted", true);
  }

  function preUpdate(id: string) {
    return estimate.cache.stories.filter(p => p.id !== id);
  }
}
