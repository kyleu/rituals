function onEstimateMessage(cmd: String, param: any) {
  switch(cmd) {
    case "detail":
      break;
    case "members":
      break;
    case "polls":
      break;
    default:
      console.warn("Unhandled command [" + cmd + "] for estimate")
  }
}
