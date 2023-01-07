(()=>{function g(t,n){let e;n?e=n.querySelectorAll(t):e=document.querySelectorAll(t);let o=[];return e.forEach(r=>{o.push(r)}),o}function M(t,n){let e=g(t,n);switch(e.length){case 0:return;case 1:return e[0];default:console.warn(`found [${e.length}] elements with selector [${t}], wanted zero or one`)}}function d(t,n){let e=M(t,n);return e||console.warn(`no element found for selector [${t}]`),e}function N(t,n){return typeof t=="string"&&(t=d(t)),t.innerHTML=n,t}function h(t,n,e="block"){return typeof t=="string"&&(t=d(t)),t.style.display=n?e:"none",t}function D(t,n){let e=document.createElement(t);for(let o in n)if(o&&n.hasOwnProperty(o)){let r=n[o];o==="dangerouslySetInnerHTML"?N(e,r.__html):r===!0?e.setAttribute(o,o):r!==!1&&r!==null&&r!==void 0&&e.setAttribute(o,r.toString())}for(let o=2;o<arguments.length;o++){let r=arguments[o];if(Array.isArray(r))r.forEach(i=>{if(r==null)throw`child array for tag [${t}] is ${r}
${e.outerHTML}`;if(i==null)throw`child for tag [${t}] is ${i}
${e.outerHTML}`;typeof i=="string"&&(i=document.createTextNode(i)),e.appendChild(i)});else{if(r==null)throw`child for tag [${t}] is ${r}
${e.outerHTML}`;r.nodeType||(r=document.createTextNode(r.toString())),e.appendChild(r)}}return e}function $(){for(let t of Array.from(document.querySelectorAll(".menu-container .final")))t.scrollIntoView({block:"nearest"})}var L="mode-light",x="mode-dark";function O(){for(let t of Array.from(document.getElementsByClassName("mode-input"))){let n=t;n.onclick=function(){switch(n.value){case"":document.body.classList.remove(L),document.body.classList.remove(x);break;case"light":document.body.classList.add(L),document.body.classList.remove(x);break;case"dark":document.body.classList.remove(L),document.body.classList.add(x);break}}}}function q(){let t=document.getElementById("flash-container");if(t===null)return;let n=t.querySelectorAll(".flash");n.length>0&&setTimeout(()=>{for(let e of n){let o=e;o.style.opacity="0",setTimeout(()=>o.remove(),500)}},3e3)}function J(){for(let t of Array.from(document.getElementsByClassName("link-confirm"))){let n=t;n.onclick=function(){let e=n.dataset.message;return e&&e.length===0&&(e="Are you sure?"),confirm(e)}}}function U(){g(".reltime").forEach(t=>{let n=t.dataset.time||"";t.innerText=ne(n)})}function ne(t){let n=new Date((t||"").replace(/-/g,"/").replace(/[TZ]/g," ")+" UTC"),e=(new Date().getTime()-n.getTime())/1e3,o=Math.floor(e/86400),r=n.getFullYear(),i=n.getMonth()+1,s=n.getDate();return isNaN(o)||o<0||o>=31?(console.log("### big",o),r.toString()+"-"+(i<10?"0"+i.toString():i.toString())+"-"+(s<10?"0"+s.toString():s.toString())):o==0&&(e<60&&"just now"||e<120&&"1 minute ago"||e<3600&&Math.floor(e/60)+" minutes ago"||e<7200&&"1 hour ago"||e<86400&&Math.floor(e/3600)+" hours ago")||o==1&&"Yesterday"||o<7&&o+" days ago"||o<31&&Math.ceil(o/7)+" weeks ago"||""}function W(){window.rituals.autocomplete=oe}function oe(t,n,e,o,r){if(!t)return;let i=t.id+"-list",s=document.createElement("datalist"),u=document.createElement("option");u.value="",u.innerText="Loading...",s.appendChild(u),s.id=i,t.parentElement?.prepend(s),t.setAttribute("autocomplete","off"),t.setAttribute("list",i);let l={},E="";function Z(c){let a=n;return a.includes("?")?a+"&t=json&"+e+"="+encodeURIComponent(c):a+"?t=json&"+e+"="+encodeURIComponent(c)}function I(c){let a=l[c];!a||!a.frag||(E=c,s.replaceChildren(a.frag.cloneNode(!0)))}function ee(){let c=t.value;if(c.length===0)return;let a=Z(c),v=!c||!E;if(!v){let m=l[E];m&&(v=!m.data.find(p=>c===r(p)))}if(!!v){if(l[c]&&l[c].url===a){I(c);return}fetch(a).then(m=>m.json()).then(m=>{if(!m)return;let p=Array.isArray(m)?m:[m],S=document.createDocumentFragment(),A=10;for(let y=0;y<p.length&&A>0;y++){let C=r(p[y]),te=o(p[y]);if(C){let b=document.createElement("option");b.value=C,b.innerText=te,S.appendChild(b),A--}}l[c]={url:a,data:p,frag:S,complete:!1},I(c)})}}t.oninput=re(ee,250),console.log("managing ["+t.id+"] autocomplete")}function re(t,n){let e=0;return function(...o){e!==0&&window.clearTimeout(e),e=window.setTimeout(function(){t.apply(null,o)},n)}}function j(){document.addEventListener("keydown",t=>{t.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}function w(t,n){return`<svg class="icon" style="width: ${n}px; height: ${n}px;"><use xlink:href="#svg-${t}"></use></svg>`}function B(){for(let t of g(".tag-editor")){let n=d("input.result",t),e=d(".tags",t),o=n.value.split(",").map(i=>i.trim()).filter(i=>i!=="");h(n,!1),e.innerHTML="";for(let i of o)e.appendChild(R(i,t));M(".add-item",t)?.remove();let r=document.createElement("div");r.className="add-item",r.innerHTML=w("plus",22),r.onclick=function(){ce(e,t)},t.insertBefore(r,d(".clear",t))}}function ie(t,n){return t.parentElement!==n.parentElement?null:t===n?0:t.compareDocumentPosition(n)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var T;function R(t,n){let e=document.createElement("div");e.className="item",e.draggable=!0,e.ondragstart=function(s){s.dataTransfer.setDragImage(document.createElement("div"),0,0),e.classList.add("dragging"),T=e},e.ondragover=function(s){let u=ie(e,T);if(!u)return;let l=u===-1?e:e.nextSibling;T.parentElement.insertBefore(T,l),H(n)},e.ondrop=function(s){s.preventDefault()},e.ondragend=function(s){e.classList.remove("dragging"),s.preventDefault()};let o=document.createElement("div");o.innerText=t,o.className="value",o.onclick=function(){X(e)},e.appendChild(o);let r=document.createElement("input");r.className="editor",e.appendChild(r);let i=document.createElement("div");return i.innerHTML=w("times",13),i.className="close",i.onclick=function(){se(e)},e.appendChild(i),e}function se(t){let n=t.parentElement.parentElement;t.remove(),H(n)}function ce(t,n){let e=R("",n);t.appendChild(e),X(e)}function X(t){let n=d(".value",t),e=d(".editor",t);e.value=n.innerText;let o=function(){if(e.value===""){t.remove();return}n.innerText=e.value,h(n,!0),h(e,!1);let r=t.parentElement.parentElement;H(r)};e.onblur=o,e.onkeydown=function(r){if(r.code==="Enter")return r.preventDefault(),o(),!1},h(n,!1),h(e,!0),e.focus()}function H(t){let n=[],e=g(".item .value",t);for(let r of e)n.push(r.innerText);let o=d("input.result",t);o.value=n.join(", ")}var F="--selected";function ae(t){let n=t.parentElement.parentElement.querySelector("input");if(!n)throw"no associated input found";n.value="\u2205"}function _(){window.rituals.setSiblingToNull=ae;let t={},n={};for(let e of Array.from(document.querySelectorAll("form.editor"))){let o=e,r=()=>{t={},n={};for(let i of o.elements){let s=i;if(s.name.length>0)if(s.name.endsWith(F))n[s.name]=s;else{(s.type!=="radio"||s.checked)&&(t[s.name]=s.value);let u=()=>{let l=n[s.name+F];l&&(l.checked=t[s.name]!==s.value)};s.onchange=u,s.onkeyup=u}}};o.onreset=r,r()}}var le=[];function G(){let t=document.querySelectorAll(".color-var");if(t.length>0)for(let n of Array.from(t)){let e=n,o=e.dataset.var,r=e.dataset.mode;le.push(o),!(!o||o.length===0)&&(e.oninput=function(){ue(r,o,e.value)})}}function ue(t,n,e){let o=document.querySelector("#mockup-"+t);if(!o){console.error("can't find mockup for mode ["+t+"]");return}switch(n){case"color-foreground":f(o,".mock-main",e);break;case"color-background":k(o,".mock-main",e);break;case"color-foreground-muted":f(o,".mock-main .mock-muted",e);break;case"color-background-muted":k(o,".mock-main .mock-muted",e);break;case"color-link-foreground":f(o,".mock-main .mock-link",e);break;case"color-link-visited-foreground":f(o,".mock-main .mock-link-visited",e);break;case"color-nav-foreground":f(o,".mock-nav",e),f(o,".mock-nav .mock-link",e);break;case"color-nav-background":k(o,".mock-nav",e);break;case"color-menu-foreground":f(o,".mock-menu",e),f(o,".mock-menu .mock-link",e);break;case"color-menu-background":k(o,".mock-menu",e);break;case"color-menu-selected-foreground":f(o,".mock-menu .mock-link-selected",e);break;case"color-menu-selected-background":k(o,".mock-menu .mock-link-selected",e);break;default:console.error("invalid key ["+n+"]")}}function P(t,n,e){let o=t.querySelectorAll(n);if(o.length==0)throw"empty query selector ["+n+"]";o.forEach(r=>{e(r)})}function k(t,n,e){P(t,n,o=>o.style.backgroundColor=e)}function f(t,n,e){P(t,n,o=>o.style.color=e)}function Y(){window.rituals.Socket=K}var z=!1,K=class{constructor(n,e,o,r,i){this.debug=n,this.open=e,this.recv=o,this.err=r,this.url=Q(i),this.connected=!1,this.pauseSeconds=1,this.pendingMessages=[],this.connect()}connect(){window.onbeforeunload=function(){z=!0},this.connectTime=Date.now(),this.sock=new WebSocket(Q(this.url));let n=this;this.sock.onopen=()=>{n.connected=!0,n.pendingMessages.forEach(n.send),n.pendingMessages=[],n.debug&&console.log("WebSocket connected"),n.open()},this.sock.onmessage=e=>{let o=JSON.parse(e.data);n.debug&&console.debug("in",o),n.recv(o)},this.sock.onerror=e=>()=>{n.err("socket",e.type)},this.sock.onclose=()=>{if(z)return;n.connected=!1;let e=n.connectTime?Date.now()-n.connectTime:0;0<e&&e<2e3?(n.pauseSeconds=n.pauseSeconds*2,n.debug&&console.debug(`socket closed immediately, reconnecting in ${n.pauseSeconds} seconds`),setTimeout(()=>{n.connect()},n.pauseSeconds*1e3)):(console.debug("socket closed after ["+e+"ms]"),setTimeout(()=>{n.connect()},n.pauseSeconds*500))}}disconnect(){}send(n){if(this.debug&&console.debug("out",n),!this.sock)throw"not initialized";if(this.connected){let e=JSON.stringify(n,null,2);this.sock.send(e)}else this.pendingMessages.push(n)}};function Q(t){if(t||(t="/connect"),t.indexOf("ws")==0)return t;let n=document.location,e="ws";return n.protocol==="https:"&&(e="wss"),t.indexOf("/")!=0&&(t="/"+t),e+`://${n.host}${t}`}function me(){console.log("[socket]: open")}function de(t){console.log("[socket]: receive "+JSON.stringify(t,null,2));let n=document.getElementById("socket-list");if(n){let e=document.createElement("pre");e.innerText=JSON.stringify(t,null,2),n.append(e)}}function fe(t){console.log("[socket error]: "+t)}function pe(t,n,e,o){new window.rituals.Socket(!0,me,de,fe,"/"+t+"/"+n.id+"/connect"),console.log("loaded ["+t+"] workspace with ["+e?.length+"] members and ["+o?.length+"] permissions")}function V(){window.rituals.initWorkspace=pe}function ge(){window.rituals={},window.JSX=D,$(),O(),q(),J(),U(),W(),j(),B(),_(),G(),Y(),V()}document.addEventListener("DOMContentLoaded",ge);})();
//# sourceMappingURL=client.js.map
