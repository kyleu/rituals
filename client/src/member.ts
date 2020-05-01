interface Member {
  userID: string;
  name: string;
  role: { key: string; };
  created: string;
}

function setMembers(members: [Member]) {
  const detail = $id("member-detail");
  detail.innerHTML = "";
  for(const member of members) {
    console.log(member);
    detail.appendChild(renderMember(member));
  }
}
