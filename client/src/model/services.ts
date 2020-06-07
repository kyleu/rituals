namespace services {
  export interface Service {
    readonly key: string;
    readonly title: string;
    readonly plural: string;
    readonly icon: string;
  }

  export const system: Service = { key: "system", title: "System", plural: "systems", icon: "close" };
  export const team: Service = { key: "team", title: "Team", plural: "teams", icon: "users" };
  export const sprint: Service = { key: "sprint", title: "Sprint", plural: "sprints", icon: "git-fork" };
  export const estimate: Service = { key: "estimate", title: "Estimate Session", plural: "estimates", icon: "settings" };
  export const standup: Service = { key: "standup", title: "Daily Standup", plural: "standups", icon: "future" };
  export const retro: Service = { key: "retro", title: "Retrospective", plural: "retros", icon: "history" };

  const allServices = [system, team, sprint, estimate, standup, retro];

  export function fromKey(key: string) {
    const ret = allServices.find(s => s.key === key);
    if (!ret) {
      throw `invalid service [${key}]`;
    }
    return ret;
  }
}
