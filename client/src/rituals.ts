namespace rituals {
  export function onError(svc: services.Service, err: string) {
    console.error(`${svc.key}: ${err}`);
    const idx = err.lastIndexOf(":");
    if (idx > -1) {
      err = err.substr(idx + 1);
    }
    notify.notify(`${svc.key} error: ${err}`, false);
  }

  export function init(svc: string, id: string) {
    window.onbeforeunload = function () {
      socket.setAppUnloading();
    };

    socket.socketConnect(services.fromKey(svc), id);
  }
}
