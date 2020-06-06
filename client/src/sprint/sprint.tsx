namespace sprint {
  export function renderSprintDates(startDate: Date | undefined, endDate: Date | undefined) {
    function f(p: string, d: Date | undefined) {
      return <span>{p} <span class="sprint-date" onclick="modal.open('session');">{d ? date.toDateString(d) : ""}</span></span>;
    }
    const s = f("starts", startDate)
    const e = f("ends", endDate)
    if (startDate) {
      if (endDate) {
        return <span>{s}, {e}</span>;
      } else {
        return s
      }
    } else {
      if (endDate) {
        return e
      } else {
        return <span>Sprint</span>
      }
    }
  }

  export function renderSprintLink(spr: sprint.Detail, bare?: boolean) {
    const profile = system.cache.getProfile();
    const a = <a class={`${profile.linkColor}-fg`} href={`/sprint/${spr.slug}`}>{spr.title}</a>
    if(bare) {
      return a;
    }
    return <span>{a} </span>;
  }

  export function renderSprintSelect(sprints: ReadonlyArray<sprint.Detail>, activeID: string | undefined) {
    return <select class="uk-select" onchange="permission.setModelPerms('sprint')">
      <option value="">- no sprint -</option>
      {sprints.map(s => {
        return s.id === activeID ? <option selected="selected" value={s.id}>{s.title}</option> : <option value={s.id}>{s.title}</option>;
      })}
    </select>
  }
}
