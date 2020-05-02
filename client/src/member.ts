interface Member {
  userID: string;
  name: string;
  role: { key: string; };
  created: string;
}

function setMembers(members: [Member]) {
  const detail = $id("member-detail");
  detail.innerHTML = "";
  detail.appendChild(renderMembers(members));
}
