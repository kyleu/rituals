(()=>{function c(e,t){let n;t?n=t.querySelectorAll(e):n=document.querySelectorAll(e);let o=[];return n.forEach(i=>{o.push(i)}),o}function M(e,t){let n=c(e,t);switch(n.length){case 0:return;case 1:return n[0];default:console.warn(`found [${n.length}] elements with selector [${e}], wanted zero or one`)}}function r(e,t){let n=M(e,t);if(!n)throw`no element found for selector [${e}]`;return n}function J(e,t){return typeof e=="string"&&(e=r(e)),e.innerHTML=t,e}function H(e,t,n="block"){return typeof e=="string"&&(e=r(e)),e.style.display=t?n:"none",e}function a(e,t,...n){let o=document.createElement(e);for(let i in t)if(i&&t.hasOwnProperty(i)){let s=t[i];i==="dangerouslySetInnerHTML"?J(o,s.__html):s===!0?o.setAttribute(i,i):s!==!1&&s!==null&&s!==void 0&&o.setAttribute(i,s.toString())}for(let i of n)if(Array.isArray(i))i.forEach(s=>{if(i==null)throw`child array for tag [${e}] is ${i}
${o.outerHTML}`;if(s==null)throw`child for tag [${e}] is ${s}
${o.outerHTML}`;typeof s=="string"&&(s=document.createTextNode(s)),o.appendChild(s)});else{if(i==null)throw`child for tag [${e}] is ${i}
${o.outerHTML}`;i.nodeType||(i=document.createTextNode(i.toString())),o.appendChild(i)}return o}function X(){for(let e of Array.from(document.querySelectorAll(".menu-container .final")))e.scrollIntoView({block:"nearest"})}var W="mode-light",_="mode-dark";function ee(){for(let e of Array.from(document.getElementsByClassName("mode-input"))){let t=e;t.onclick=function(){switch(t.value){case"":document.body.classList.remove(W),document.body.classList.remove(_);break;case"light":document.body.classList.add(W),document.body.classList.remove(_);break;case"dark":document.body.classList.remove(W),document.body.classList.add(_);break}}}}function h(e,t,n){let o=document.getElementById("flash-container");o===null&&(o=document.createElement("div"),o.id="flash-container",document.body.insertAdjacentElement("afterbegin",o));let i=document.createElement("div");i.className="flash";let s=document.createElement("input");s.type="radio",s.style.display="none",s.id="hide-flash-"+e,i.appendChild(s);let m=document.createElement("label");m.htmlFor="hide-flash-"+e;let d=document.createElement("span");d.innerHTML="\xD7",m.appendChild(d),i.appendChild(m);let u=document.createElement("div");u.className="content flash-"+t,u.innerText=n,i.appendChild(u),o.appendChild(i),ne(i)}function te(){let e=document.getElementById("flash-container");if(e===null)return h;let t=e.querySelectorAll(".flash");if(t.length>0)for(let n of t)ne(n);return h}function ne(e){setTimeout(()=>{e.style.opacity="0",setTimeout(()=>e.remove(),500)},5e3)}function oe(){for(let e of Array.from(document.getElementsByClassName("link-confirm"))){let t=e;t.onclick=function(){let n=t.dataset.message;return n&&n.length===0&&(n="Are you sure?"),confirm(n)}}}function re(){return c(".reltime").forEach(e=>{w(e.dataset.time||"",e)}),w}function ie(e){let t=Date.UTC(e.getUTCFullYear(),e.getUTCMonth(),e.getUTCDate(),e.getUTCHours(),e.getUTCMinutes(),e.getUTCSeconds());return new Date(t).toISOString().substring(0,19).replace("T"," ")}function w(e,t){let n=(e||"").replace(/-/g,"/").replace(/[TZ]/g," ")+" UTC",o=new Date(n),i=(new Date().getTime()-o.getTime())/1e3,s=Math.floor(i/86400),m=o.getFullYear(),d=o.getMonth()+1,u=o.getDate();if(isNaN(s)||s<0||s>=31)return m.toString()+"-"+(d<10?"0"+d.toString():d.toString())+"-"+(u<10?"0"+u.toString():u.toString());let l="",f=0;return s==0?i<5?(f=1,l="just now"):i<60?(f=1,l=Math.floor(i)+" seconds ago"):i<120?(f=10,l="1 minute ago"):i<3600?(f=30,l=Math.floor(i/60)+" minutes ago"):i<7200?(f=60,l="1 hour ago"):(f=60,l=Math.floor(i/3600)+" hours ago"):s==1?(f=600,l="yesterday"):s<7?(f=600,l=s+" days ago"):(f=6e3,l=Math.ceil(s/7)+" weeks ago"),t&&(t.innerText=l,setTimeout(()=>w(e,t),f*1e3)),l}function se(){return et}function et(e,t,n,o,i){if(!e)return;let s=e.id+"-list",m=document.createElement("datalist"),d=document.createElement("option");d.value="",d.innerText="Loading...",m.appendChild(d),m.id=s,e.parentElement?.prepend(m),e.setAttribute("autocomplete","off"),e.setAttribute("list",s);let u={},l="";function f(g){let v=t;return v.includes("?")?v+"&t=json&"+n+"="+encodeURIComponent(g):v+"?t=json&"+n+"="+encodeURIComponent(g)}function O(g){let v=u[g];!v||!v.frag||(l=g,m.replaceChildren(v.frag.cloneNode(!0)))}function Je(){let g=e.value;if(g.length===0)return;let v=f(g),B=!g||!l;if(!B){let b=u[l];b&&(B=!b.data.find(L=>g===i(L)))}if(!!B){if(u[g]&&u[g].url===v){O(g);return}fetch(v).then(b=>b.json()).then(b=>{if(!b)return;let L=Array.isArray(b)?b:[b],K=document.createDocumentFragment(),Q=10;for(let A=0;A<L.length&&Q>0;A++){let Z=i(L[A]),Xe=o(L[A]);if(Z){let j=document.createElement("option");j.value=Z,j.innerText=Xe,K.appendChild(j),Q--}}u[g]={url:v,data:L,frag:K,complete:!1},O(g)})}}e.oninput=tt(Je,250),console.log("managing ["+e.id+"] autocomplete")}function tt(e,t){let n=0;return function(...o){n!==0&&window.clearTimeout(n),n=window.setTimeout(function(){e(null,...o)},t)}}function ae(){document.addEventListener("keydown",e=>{e.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}function T(e,t){return`<svg class="icon" style="width: ${t}px; height: ${t}px;"><use xlink:href="#svg-${e}"></use></svg>`}function C(e,t){return{__html:`<svg class="${t||""}" style="width: 18px; height: 18px;"><use href="#svg-${e}"></use></svg>`}}function x(e){let t=r("input.result",e),n=r(".tags",e),o=t.value.split(",").map(s=>s.trim()).filter(s=>s!=="");H(t,!1),n.innerHTML="";for(let s of o)n.appendChild(ce(s,e));M(".add-item",e)?.remove();let i=document.createElement("div");i.className="add-item",i.innerHTML=T("plus",22),i.onclick=function(){rt(n,e)},e.insertBefore(i,r(".clear",e))}function me(){for(let e of c(".tag-editor"))x(e);return x}function nt(e,t){return e.parentElement!==t.parentElement?null:e===t?0:e.compareDocumentPosition(t)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var S;function ce(e,t){let n=document.createElement("div");n.className="item",n.draggable=!0,n.ondragstart=function(m){m.dataTransfer?.setDragImage(document.createElement("div"),0,0),n.classList.add("dragging"),S=n},n.ondragover=function(){let m=nt(n,S);if(!m)return;let d=m===-1?n:n.nextSibling;S.parentElement?.insertBefore(S,d),V(t)},n.ondrop=function(m){m.preventDefault()},n.ondragend=function(m){n.classList.remove("dragging"),m.preventDefault()};let o=document.createElement("div");o.innerText=e,o.className="value",o.onclick=function(){le(n)},n.appendChild(o);let i=document.createElement("input");i.className="editor",n.appendChild(i);let s=document.createElement("div");return s.innerHTML=T("times",13),s.className="close",s.onclick=function(){ot(n)},n.appendChild(s),n}function ot(e){let t=e.parentElement?.parentElement;e.remove(),t&&V(t)}function rt(e,t){let n=ce("",t);e.appendChild(n),le(n)}function le(e){let t=r(".value",e),n=r(".editor",e);n.value=t.innerText;let o=function(){if(n.value===""){e.remove();return}t.innerText=n.value,H(t,!0),H(n,!1);let i=e.parentElement?.parentElement;i&&V(i)};n.onblur=o,n.onkeydown=function(i){if(i.code==="Enter")return i.preventDefault(),o(),!1},H(t,!1),H(n,!0),n.focus()}function V(e){let t=[],n=c(".item .value",e);for(let i of n)t.push(i.innerText);let o=r("input.result",e);o.value=t.join(", ")}var de="--selected";function it(e){let t=e.parentElement?.parentElement?.querySelector("input");if(!t)throw"no associated input found";t.value="\u2205"}function G(e){e.onreset=()=>G(e);let t={},n={};for(let o of e.elements){let i=o;if(i.name.length>0)if(i.name.endsWith(de))n[i.name]=i;else{(i.type!=="radio"||i.checked)&&(t[i.name]=i.value);let s=()=>{let m=n[i.name+de];m&&(m.checked=t[i.name]!==i.value)};i.onchange=s,i.onkeyup=s}}}function ue(){for(let e of Array.from(document.querySelectorAll("form.editor")))G(e);return[it,G]}var st=[];function pe(){let e=document.querySelectorAll(".color-var");if(e.length>0)for(let t of Array.from(e)){let n=t,o=n.dataset.var,i=n.dataset.mode;st.push(o),!(!o||o.length===0)&&(n.oninput=function(){at(i,o,n.value)})}}function at(e,t,n){let o=document.querySelector("#mockup-"+e);if(!o){console.error("can't find mockup for mode ["+e+"]");return}switch(t){case"color-foreground":y(o,".mock-main",n);break;case"color-background":D(o,".mock-main",n);break;case"color-foreground-muted":y(o,".mock-main .mock-muted",n);break;case"color-background-muted":D(o,".mock-main .mock-muted",n);break;case"color-link-foreground":y(o,".mock-main .mock-link",n);break;case"color-link-visited-foreground":y(o,".mock-main .mock-link-visited",n);break;case"color-nav-foreground":y(o,".mock-nav",n),y(o,".mock-nav .mock-link",n);break;case"color-nav-background":D(o,".mock-nav",n);break;case"color-menu-foreground":y(o,".mock-menu",n),y(o,".mock-menu .mock-link",n);break;case"color-menu-background":D(o,".mock-menu",n);break;case"color-menu-selected-foreground":y(o,".mock-menu .mock-link-selected",n);break;case"color-menu-selected-background":D(o,".mock-menu .mock-link-selected",n);break;default:console.error("invalid key ["+t+"]")}}function fe(e,t,n){let o=e.querySelectorAll(t);if(o.length==0)throw"empty query selector ["+t+"]";o.forEach(i=>n(i))}function D(e,t,n){fe(e,t,o=>o.style.backgroundColor=n)}function y(e,t,n){fe(e,t,o=>o.style.color=n)}function ge(){return R}var he=!1,R=class{constructor(t,n,o,i,s){this.debug=t,this.open=n,this.recv=o,this.err=i,this.url=Te(s),this.connected=!1,this.pauseSeconds=1,this.pendingMessages=[],this.connect()}connect(){window.onbeforeunload=function(){he=!0},this.connectTime=Date.now(),this.sock=new WebSocket(Te(this.url));let t=this;this.sock.onopen=()=>{t.connected=!0,t.pendingMessages.forEach(t.send),t.pendingMessages=[],t.debug&&console.log("WebSocket connected"),t.open()},this.sock.onmessage=n=>{let o=JSON.parse(n.data);t.debug&&console.debug("[socket]: receive",o),t.recv(o)},this.sock.onerror=n=>()=>{t.err("socket",n.type)},this.sock.onclose=()=>{if(he)return;t.connected=!1;let n=t.connectTime?Date.now()-t.connectTime:0;0<n&&n<2e3?(t.pauseSeconds=t.pauseSeconds*2,t.debug&&console.debug(`socket closed immediately, reconnecting in ${t.pauseSeconds} seconds`),setTimeout(()=>{t.connect()},t.pauseSeconds*1e3)):(console.debug("socket closed after ["+n+"ms]"),setTimeout(()=>{t.connect()},t.pauseSeconds*500))}}disconnect(){}send(t){if(this.debug&&console.debug("out",t),!this.sock)throw"not initialized";if(this.connected){let n=JSON.stringify(t,null,2);this.sock.send(n)}else this.pendingMessages.push(t)}};function Te(e){if(e||(e="/connect"),e.indexOf("ws")==0)return e;let t=document.location,n="ws";return t.protocol==="https:"&&(n="wss"),e.indexOf("/")!=0&&(e="/"+e),n+`://${t.host}${e}`}function ve(e,t,n){return a("tr",{id:"member-"+e,class:"member","data-id":e},a("td",null,a("a",{href:"#modal-member-"+e},a("div",{class:"left",style:"padding-right: var(--padding-small);",dangerouslySetInnerHTML:C("profile")}),a("span",{class:"member-name"},t))),a("td",{class:"shrink",style:"text-align: right"},a("em",{class:"member-status"},n)),a("td",{class:"shrink online-status",title:"offline",dangerouslySetInnerHTML:C("circle","right")}))}function Me(e,t,n){let o=[["owner","Owner"],["member","Member"],["observer","Observer"]];return a("div",{id:"modal-member-"+e,"data-id":e,class:"modal modal-member",style:"display: none;"},a("a",{class:"backdrop",href:"#"}),a("div",{class:"modal-content"},a("div",{class:"modal-header"},a("a",{href:"#",class:"modal-close"},"\xD7"),a("h2",null,t)),a("div",{class:"modal-body"},a("form",{action:document.location.pathname,method:"post",class:"expanded"},a("input",{type:"hidden",name:"userID",value:e}),a("em",null,"Role"),a("br",null),a("select",{class:"input-role",name:"role"},o.map(i=>i[0]==n?a("option",{selected:"selected",value:i[0]},i[1]):a("option",{value:i[0]},i[1]))),a("hr",null),a("div",{class:"right"},a("button",{class:"member-update",type:"submit",name:"action",value:"member-update"},"Save")),a("button",{type:"submit",class:"member-remove",name:"action",value:"member-remove",onClick:"return confirm('Are you sure you wish to remove this user?');"},"Remove")))))}var F,k;function Ee(e){if(!e)return"System";let t=k[e];return t||"Unknown User"}function be(){return F}function ye(){mt(),ct(),P()}function mt(){let e=r("#modal-self"),t=r("form",e);t.onsubmit=function(){let n=r('input[name="name"]',t),o=r('input[name="choice"]:checked',t);return p("self",{name:n.value,choice:o.value}),r("#self-name").innerText=n.value,document.location.hash="",!1}}function ct(){let e=c(".modal-member");for(let t of e)ke(t)}function ke(e){let t=r("form",e),n=function(o){let i=r('input[name="userID"]',t).value,s=r('select[name="role"]',t).value;p(o,{userID:i,role:s});let m=r("#member-"+i);return o==="member-update"?r(".member-role",m).innerText=s:o==="member-remove"&&(m.remove(),P()),document.location.hash="",!1};r(".member-update",t).onclick=()=>n("member-update"),r(".member-remove",t).onclick=()=>confirm("Are you sure you wish to remove this user?")?n("member-remove"):!1}function P(){k={},F=r("#self-id").innerText,k[F]=r("#self-name").innerText;let e=r("#panel-members"),t=c(".member",e);for(let n of t){let o=n.dataset.id;o&&(k[o]=r(".member-name",n).innerText)}}function Le(e,t,n){if(M("#member-"+e)||e===F)return Y(e,t,n);let i=r("#panel-members table tbody"),s=-1;for(let l=0;l<i.children.length;l++){let f=i.children.item(l);if(r(".member-name",f).innerText.localeCompare(t,void 0,{sensitivity:"accent"})>0){s=l;break}}let m=ve(e,t,n);s==-1?i.appendChild(m):i.insertBefore(m,i.children[s]);let d=r("#member-modals"),u=Me(e,t,n);d.appendChild(u),ke(u),k[e]=t}function Y(e,t,n){if(e===F)r("#self-name").innerText=t,r("#self-role").innerText=n;else{let o=r("#member-"+e);r(".member-name",o).innerText=t,r(".member-role",o).innerText=n;let i=r("#modal-member-"+e);r('select[name="role"]',i).value=n}if(k[e]!==t){k[e]=t;let o=r("#panel-members table tbody"),i=o.children,s=[];for(let m in i)i[m].nodeType==1&&s.push(i[m]);s.sort((m,d)=>{let u=r(".member-name",m).innerText,l=r(".member-name",d).innerText;return u.localeCompare(l,void 0,{sensitivity:"accent"})}),o.replaceChildren(...s)}}function He(e){r("#member-"+e).remove(),P()}function xe(e,t){let n=M("#member-"+e+" .online-status");if(!n)throw"missing panel #member-"+e;n.title=t?"online":"offline";let o=t?"check-circle":"circle";n.innerHTML='<svg style="width: 18px; height: 18px;" class="right"><use xlink:href="#svg-'+o+'"></use></svg>'}function Ie(e,t){let n=a("li",null),o=ie(new Date),i=a("span",{class:"nowrap reltime","data-time":o},"just now");w(o,i);let s=a("div",{class:"right"});return s.appendChild(i),n.appendChild(s),n.appendChild(a("div",null,e.content)),n.appendChild(a("div",null,a("em",null,t))),n}function we(){let e=c(".modal.comments");for(let t of e){let n=r("form",t);n.onsubmit=function(){let o={svc:r('input[name="svc"]',n).value,modelID:r('input[name="modelID"]',n).value},i=r("textarea",n);return o.content=i.value,o.userID=be(),p("comment",o),z(o),!1}}}function z(e){let t=r("#comment-list-"+e.svc+"-"+e.modelID),n=Ee(e.userID),o=Ie(e,n);t.appendChild(o);let i=t.childNodes.length-1,s=r("#comment-link-"+e.svc+"-"+e.modelID);s.title=i+(i==1?" comment":" comments"),s.innerHTML.indexOf("comment-dots")==-1&&(s.innerHTML='<svg style="width: 18px; height: 18px;" class="right"><use xlink:href="#svg-comment-dots"></use></svg>')}function E(e,t,n,o){let i="";if(o){let s=M(`select[name="sprint"] option[value="${o}"]`,t);if(s){let m=s.innerText;i+=`<a href="/sprint/${o}">${m}</a> `}}if(i+=e,n){let s=M(`select[name="team"] option[value="${n}"]`,t);if(s){let m=s.innerText;i+=` in <a href="/team/${n}">${m}</a>`}}return i}function $(e){let t=r(`#${e.type}-list tbody`),n=M(".empty",t);n&&n.remove();let o=document.createElement("tr");o.id=`${e.type}-list-${e.id}`;let i=document.createElement("td"),s=document.createElement("a");s.href=e.path,s.innerText=e.title,i.appendChild(s),o.appendChild(i),t.appendChild(o)}function N(e){r(`#${e.type}-list-${e.id}`).remove();let t=r(`#${e.type}-list tbody`);if(t.children.length===0){let n=document.createElement("tr");n.classList.add("empty");let o=document.createElement("em");o.innerText="no "+e.type+"s",n.appendChild(o),t.appendChild(n)}}function Ce(){let e=r("#modal-team-config form");e.onsubmit=function(){let t=r('input[name="title"]',e).value,n=r('input[name="icon"]:checked',e).value;return p("update",{title:t,icon:n}),document.location.hash="",!1}}function De(e){switch(e.cmd){case"update":return lt(e.param);case"child-add":return $(e.param);case"child-remove":return N(e.param);default:throw"invalid team command ["+e.cmd+"]"}}function lt(e){let t=r("#modal-team-config form");r('input[name="title"]',t).value=e.title;for(let n of c('input[name="icon"]',t))n.checked=e.icon===n.value;r("#model-title").innerText=e.title,r("#model-icon").innerHTML=T(e.icon,20),r("#model-banner").innerHTML=E("team",t,"",""),h("team","success","team updated")}function Fe(){let e=r("#modal-sprint-config form");e.onsubmit=function(){let t=r('input[name="title"]',e).value,n=r('input[name="icon"]:checked',e).value,o=r('input[name="startDate"]',e).value,i=r('input[name="endDate"]',e).value,s=r('select[name="team"]',e).value;return p("update",{title:t,icon:n,startDate:o,endDate:i,team:s}),document.location.hash="",!1}}function Ae(e){switch(e.cmd){case"update":return dt(e.param);case"child-add":return $(e.param);case"child-remove":return N(e.param);default:throw"invalid sprint command ["+e.cmd+"]"}}function dt(e){let t=r("#modal-sprint-config form");r('input[name="title"]',t).value=e.title;for(let n of c('input[name="icon"]',t))n.checked=e.icon===n.value;r('input[name="startDate"]',t).value=U(e.startDate),r('input[name="endDate"]',t).value=U(e.endDate),r('select[name="team"]',t).value=e.teamID?e.teamID:"",r("#model-title").innerText=e.title,r("#model-icon").innerHTML=T(e.icon,20),r("#model-summary").innerText=ut(e),r("#model-banner").innerHTML=E("sprint",t,e.teamID,""),h("sprint","success","sprint updated")}function ut(e){let t="";return e.startDate&&(t+="starts ",t+=U(e.startDate),e.endDate&&(t+=", ")),e.endDate&&(t+="ends ",t+=U(e.endDate)),t}function U(e){return`${e}`.split("T")[0]}function Se(e){let t="comment-link-"+e.estimateID,n="#modal-"+e.estimateID+"-comments";return a("tr",{class:"story-row",id:"story-row-{%s s.ID.String() %}","data-idx":"s.Idx"},a("td",null,a("a",{href:"#modal-story-"+e.id,class:"story-title"},e.title)),a("td",null,e.status),a("td",null,e.finalVote?e.finalVote:"-"),a("td",null,a("a",{id:t,href:n,title:"0 comments",dangerouslySetInnerHTML:C("comment-alt")})))}function Re(e){let t=e.dataset.id;if(!t){console.warn("no id in dataset",e);return}c(".status-new-form-delete",e).forEach(n=>ht(t,n)),c(".status-new-form-next",e).forEach(n=>$e(t,n)),c(".status-active-form-prev",e).forEach(n=>pt(t,n)),c(".status-active-form-next",e).forEach(n=>ft(t,n)),c(".status-active-form-vote",e).forEach(n=>gt(t,n)),c(".status-complete-form-prev",e).forEach(n=>$e(t,n))}function pt(e,t){t.onsubmit=function(){return console.log("new"),!1}}function $e(e,t){t.onsubmit=function(){return console.log("active"),!1}}function ft(e,t){t.onsubmit=function(){return console.log("complete"),!1}}function gt(e,t){t.onsubmit=function(){return console.log("vote"),!1}}function ht(e,t){t.onsubmit=function(){return console.log("delete"),!1}}function Ne(){c(".add-story-link").forEach(n=>n.onclick=function(){return setTimeout(()=>r("#story-add-title").focus(),100),!0});let e=r("#modal-story--add"),t=r("form",e);t.onsubmit=function(){let n=r('input[name="title"]',t),o=n.value;return n.value="",p("child-add",{title:o}),!1},c("#story-modals .modal-story").forEach(Re)}function Ue(e){let t=r("#panel-detail table tbody"),n=-1;for(let m=0;m<t.children.length;m++){let d=t.children.item(m),u=r(".story-title",d).innerText,l=d.dataset.index;if(l){if(parseInt(l,10)>=e.idx){n=m;break}else if(u.localeCompare(e.title,void 0,{sensitivity:"accent"})>=0){n=m;break}}}let o=Se(e);n==-1?t.appendChild(o):t.insertBefore(o,t.children[n]);let s=r("#modal-story-new").cloneNode(!0);s.id="modal-story-"+e.id,s.dataset.id=e.id,s.dataset.status=e.status,s.classList.add("modal-story"),r("#story-modals").appendChild(s),(document.location.hash==="modal-story--add"||document.location.hash==="")&&(document.location.hash="modal-story-"+e.id)}function qe(){let e=r("#modal-estimate-config form");e.onsubmit=function(){let t=r('input[name="title"]',e).value,n=r('input[name="icon"]:checked',e).value,o=r('input[name="choices"]',e).value,i=r('select[name="team"]',e).value,s=r('select[name="sprint"]',e).value;return p("update",{title:t,icon:n,choices:o,team:i,sprint:s}),document.location.hash="",!1},Ne()}function Oe(e){switch(e.cmd){case"update":return Tt(e.param);case"child-add":return Ue(e.param);default:throw"invalid estimate command ["+e.cmd+"]"}}function Tt(e){let t=r("#modal-estimate-config form");r('input[name="title"]',t).value=e.title;for(let o of c('input[name="icon"]',t))o.checked=e.icon===o.value;let n=r('input[name="choices"]',t);n.value=e.choices,n.parentElement&&x(n.parentElement),r('select[name="team"]',t).value=e.teamID?e.teamID:"",r('select[name="sprint"]',t).value=e.sprintID?e.sprintID:"",r("#model-title").innerText=e.title,r("#model-icon").innerHTML=T(e.icon,20),r("#model-banner").innerHTML=E("estimate",t,e.teamID,e.sprintID),h("estimate","success","estimate updated")}function Be(){c(".add-report-link").forEach(n=>n.onclick=function(){return setTimeout(()=>r("#report-add-content").focus(),100),!0});let e=r("#modal-report--add"),t=r("form",e);t.onsubmit=function(){let n=r('input[name="day"]',t).value,o=r('textarea[name="content"]',t),i=o.value;return o.value="",p("child-add",{day:n,content:i}),document.location.hash="",!1}}function je(e){console.log("TODO: reportAdd")}function We(){let e=r("#modal-standup-config form");e.onsubmit=function(){let t=r('input[name="title"]',e).value,n=r('input[name="icon"]:checked',e).value,o=r('select[name="team"]',e).value,i=r('select[name="sprint"]',e).value;return p("update",{title:t,icon:n,team:o,sprint:i}),document.location.hash="",!1},Be()}function _e(e){switch(e.cmd){case"update":return vt(e.param);case"child-add":return je(e.param);default:throw"invalid standup command ["+e.cmd+"]"}}function vt(e){let t=r("#modal-standup-config form");r('input[name="title"]',t).value=e.title;for(let n of c('input[name="icon"]',t))n.checked=e.icon===n.value;r('select[name="team"]',t).value=e.teamID?e.teamID:"",r('select[name="sprint"]',t).value=e.sprintID?e.sprintID:"",r("#model-title").innerText=e.title,r("#model-icon").innerHTML=T(e.icon,20),r("#model-banner").innerHTML=E("standup",t,e.teamID,e.sprintID),h("standup","success","standup updated")}function Ve(e){return a("div",null,"TODO")}function Ge(){c(".add-feedback-link").forEach(e=>e.onclick=function(){let t=e.dataset.category;return setTimeout(()=>r("#feedback-add-content-"+t).focus(),100),!0});for(let e of c(".modal-feedback")){let t=r("form",e);t.onsubmit=function(){let n=r('select[name="category"]',t).value,o=r('textarea[name="content"]',t),i=o.value;return o.value="",p("child-add",{category:n,content:i}),document.location.hash="",!1}}}function Pe(e){let t=r("#category-"+e.category+" .feedback-list"),n=-1;for(let i=0;i<t.children.length;i++){let s=t.children.item(i),m=r(".feedback-content",s).innerText,d=s.dataset.index;if(d){if(parseInt(d,10)>=e.idx){n=i;break}else if(m.localeCompare(e.content,void 0,{sensitivity:"accent"})>=0){n=i;break}}}let o=Ve(e);n==-1?t.appendChild(o):t.insertBefore(o,t.children[n])}function Ye(){let e=r("#modal-retro-config form");e.onsubmit=function(){let t=r('input[name="title"]',e).value,n=r('input[name="icon"]:checked',e).value,o=r('input[name="categories"]',e).value,i=r('select[name="team"]',e).value,s=r('select[name="sprint"]',e).value;return p("update",{title:t,icon:n,categories:o,team:i,sprint:s}),document.location.hash="",!1},Ge()}function ze(e){switch(e.cmd){case"update":return Mt(e.param);case"child-add":return Pe(e.param);default:throw"invalid retro command ["+e.cmd+"]"}}function Mt(e){let t=r("#modal-retro-config form");r('input[name="title"]',t).value=e.title;for(let o of c('input[name="icon"]',t))o.checked=e.icon===o.value;let n=r('input[name="categories"]',t);n.value=e.categories,n.parentElement&&x(n.parentElement),r('select[name="team"]',t).value=e.teamID?e.teamID:"",r('select[name="sprint"]',t).value=e.sprintID?e.sprintID:"",r("#model-title").innerText=e.title,r("#model-icon").innerHTML=T(e.icon,20),r("#model-banner").innerHTML=E("retro",t,e.teamID,e.sprintID),h("retro","success","retro updated")}function Ke(e,t){switch(t.cmd){case"error":return Et(t.param.message);case"comment":return z(t.param);case"online-update":let n=t.param;return xe(n.userID,n.connected);case"member-add":let o=t.param;return Le(o.userID,o.name,o.role);case"member-update":let i=t.param;return Y(i.userID,i.name,i.role);case"member-remove":return He(t.param)}switch(e){case"team":return De(t);case"sprint":return Ae(t);case"estimate":return Oe(t);case"standup":return _e(t);case"retro":return ze(t);default:throw"invalid service ["+e+"]"}}function Et(e){h("error","error",e)}var Qe,I,q;function bt(){console.log("[socket]: open")}function yt(e){let t=document.getElementById("socket-list");if(t){let n=document.createElement("pre");n.innerText=JSON.stringify(e,null,2),t.append(n)}Ke(I,e)}function kt(e){console.log("[socket error]: "+e)}function Lt(e,t){switch(I=e,q=t,we(),ye(),I){case"team":Ce();break;case"sprint":Fe();break;case"estimate":qe();break;case"standup":We();break;case"retro":Ye();break}Qe=new R(!0,bt,yt,kt,"/"+I+"/"+q+"/connect"),console.log("loaded ["+I+"] workspace ["+q+"]")}function p(e,t){Qe.send({channel:I+":"+q,cmd:e,param:t})}function Ze(){window.initWorkspace=Lt}function Ht(){let[e,t]=ue();window.rituals={relativeTime:re(),autocomplete:se(),setSiblingToNull:e,initForm:t,flash:te(),tags:me(),Socket:ge()},X(),ee(),oe(),ae(),pe(),window.JSX=a,Ze()}document.addEventListener("DOMContentLoaded",Ht);})();
//# sourceMappingURL=client.js.map
