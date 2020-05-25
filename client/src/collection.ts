namespace collection {
  export class Group<T> {
    readonly key: string;
    readonly members: T[] = [];

    constructor(key: string) {
      this.key = key;
    }
  }

  export function groupBy<T>(list: T[] | null, func: (x: T) => string): Group<T>[] {
    const res: Group<T>[] = [];
    let group: Group<T> | undefined;

    if (list) {
      list.forEach((o) => {
        const groupName = func(o);
        if (!group) {
          group = new Group<T>(groupName);
        }
        if (groupName != group.key) {
          res.push(group);
          group = new Group<T>(groupName);
        }
        group.members.push(o);
      });
    }
    if (group) {
      res.push(group);
    }
    return res;
  }

  export function findGroup<T>(groups: collection.Group<T>[], key: string): T[] {
    for (const g of groups) {
      if (g.key === key) {
        return g.members;
      }
    }
    return []
  }

  export function flatten<T>(a: T[][]): T[] {
    const ret: T[] = [];
    a.forEach(v => ret.push(...v));
    return ret;
  }
}
