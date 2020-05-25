namespace estimate {
  interface Detail extends rituals.Session {
    readonly status: { key: string };
    readonly choices: string[];
  }

  export interface StoryStatusChange {
    readonly storyID: string;
    readonly status: string;
    readonly finalVote: string;
  }

  interface SessionJoined extends rituals.SessionJoined {
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
        rituals.onError(services.estimate.key, param as string);
        break;
      case command.server.sessionJoined:
        const sj = param as SessionJoined;
        rituals.onSessionJoin(sj);
        rituals.setTeam(sj.team);
        rituals.setSprint(sj.sprint);
        setEstimateDetail(sj.session);
        story.setStories(sj.stories);
        vote.setVotes(sj.votes);
        rituals.showWelcomeMessage(sj.members.length);
        break;
      case command.server.sessionUpdate:
        setEstimateDetail(param as Detail);
        break;
      case command.server.permissionsUpdate:
        system.setPermissions(param as permission.Permission[]);
        break;
      case command.server.teamUpdate:
        const tm = param as team.Detail | undefined;
        if (estimate.cache.detail) {
          estimate.cache.detail.teamID = tm?.id;
        }
        rituals.setTeam(tm);
        break;
      case command.server.sprintUpdate:
        const spr = param as sprint.Detail | undefined;
        if (estimate.cache.detail) {
          estimate.cache.detail.sprintID = spr?.id;
        }
        rituals.setSprint(spr)
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
    dom.setValue("#model-choices-input", detail.choices.join(", "));
    story.viewActiveStory();
    rituals.setDetail(detail);
  }

  export function onSubmitEstimateSession() {
    const title = dom.req<HTMLInputElement>("#model-title-input").value;
    const choices = dom.req<HTMLInputElement>("#model-choices-input").value;
    const teamID = dom.req<HTMLSelectElement>("#model-team-select select").value;
    const sprintID = dom.req<HTMLSelectElement>("#model-sprint-select select").value;
    const permissions = permission.readPermissions();

    const msg = {svc: services.estimate.key, cmd: command.client.updateSession, param: {title, choices, teamID, sprintID, permissions}};
    socket.send(msg);
  }

  export function onStoryUpdate(s: story.Story) {
    const x = preUpdate(s.id);
    x.push(s);
    if (s.id === estimate.cache.activeStory) {
      dom.setText("#story-title", s.title);
    }
    story.setStories(x);
  }

  export function onStoryRemove(id: string) {
    const x = preUpdate(id);
    story.setStories(x);
    if (id === estimate.cache.activeStory) {
      UIkit.modal("#modal-story").hide();
    }
    UIkit.notification("story has been deleted", {status: "success", pos: "top-right"});
  }

  function preUpdate(id: string) {
    return estimate.cache.stories.filter((p) => p.id !== id);
  }
}
