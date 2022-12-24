(function(){var e,t,n,r,i="undefined",o="none",a="header",l=typeof window.onpagehide!==i?"pagehide":"beforeunload",c="Enter",u=function(){};e=typeof console!==i?function(e){console.error(e)}:u,document.body.classList?(t=function(e,t){return e&&e.classList.contains(t)},n=function(e,t){e&&e.classList.add(t)},r=function(e,t){e&&e.classList.remove(t)}):(t=function(e,t){if(e)return new RegExp("\\b"+t+"\\b").test(e.className)},n=function(e,t){if(e){var n=e.className;new RegExp("\\b"+t+"\\b").test(n)||(e.className=n+" "+t)}},r=function(e,t){if(e){var n=e.className,r=new RegExp("^\\s*"+t+"\\s+|\\s+"+t+"\\b","g"),i=n.replace(r,"");n!==i&&(e.className=i)}});var s=!1;try{typeof sessionStorage!==i&&(s=!0)}catch(f){}(function(){var e;if(document.querySelector){if(e=document.body.querySelector(".filter"))if(e.addEventListener){var t=e.querySelector("input");if(t){var i,u,f=String.prototype.trim?function(e){return e.trim()}:(i=/^\s+|\s+$/g,function(e){return e.replace(i,"")}),d=e.querySelector("button"),v="."+o,y=".item-list > li:not(."+a+"):not(.parent)",p=y+v,m=y+":not(.none)",g="",h=function(){var e=f(t.value).toLowerCase();if(e!==g){var i,a,l;if(e)for(d&&(d.style.display="block"),i=e.indexOf(g)>=0?m:g.indexOf(e)>=0?p:y,l=(a=document.body.querySelectorAll(i)).length-1;l>=0;l--){var c=a[l],u=c.querySelector(".name");u&&u.textContent.toLowerCase().indexOf(e)<0?n(c,o):r(c,o)}else for(d&&(d.style.display=""),i=p,l=(a=document.body.querySelectorAll(i)).length-1;l>=0;l--)r(a[l],o);g=e}},E=function(){clearTimeout(u),u=setTimeout(h,350)};t.addEventListener("input",E,!1),t.addEventListener("change",E,!1);var b=function(){clearTimeout(u),t.blur(),h()},S=function(){clearTimeout(u),t.value="",h()};if(t.addEventListener("keydown",(function(e){if(e.key)switch(e.key){case c:b(),e.preventDefault();break;case"Escape":case"Esc":S(),e.preventDefault()}else if(e.keyCode)switch(e.keyCode){case 13:b(),e.preventDefault();break;case 27:S(),e.preventDefault()}}),!1),d&&d.addEventListener("click",(function(){clearTimeout(u),t.value="",t.focus(),h()})),s){var L=sessionStorage.getItem(location.pathname);sessionStorage.removeItem(location.pathname),window.addEventListener(l,(function(){t.value&&sessionStorage.setItem(location.pathname,t.value)}),!1),L&&(t.value=L)}t.value&&h()}}else e.className+=" none"}else(e=document.getElementById&&document.getElementById("panel-filter"))&&(e.className+=" none")})(),function(){if(document.querySelector&&document.addEventListener&&document.body.parentElement){var e=document.body.querySelector(".path-list"),n=document.body.querySelector(".item-list");if(e||n){var r,i,l,c,u,s,f="li:not(.none):not(."+a+")",d=["INPUT","BUTTON","TEXTAREA"],v=navigator.platform,y=v.indexOf("Mac")>=0||v.indexOf("iPhone")>=0||v.indexOf("iPad")>=0||v.indexOf("iPod")>=0;h(),y?(u=function(e){return!(e.ctrlKey||e.shiftKey||e.metaKey)},s=function(e){return e.altKey}):(u=function(e){return!(e.altKey||e.shiftKey||e.metaKey)},s=function(e){return e.ctrlKey}),document.addEventListener("keydown",(function(t){var r=function(t){if(!(d.indexOf(t.target.tagName)>=0))if(t.key){if(u(t))switch(t.key){case"Left":case"ArrowLeft":return s(t)?m(e):p(e,!0);case"Right":case"ArrowRight":return s(t)?g(e):p(e,!1);case"Up":case"ArrowUp":return s(t)?m(n):p(n,!0);case"Down":case"ArrowDown":return s(t)?g(n):p(n,!1)}if(!t.ctrlKey&&(!t.altKey||y)&&!t.metaKey&&1===t.key.length)return E(t.key)}else if(t.keyCode){if(u(t))switch(t.keyCode){case 37:return s(t)?m(e):p(e,!0);case 39:return s(t)?g(e):p(e,!1);case 38:return s(t)?m(n):p(n,!0);case 40:return s(t)?g(n):p(n,!1)}if(!t.ctrlKey&&(!t.altKey||y)&&!t.metaKey&&t.keyCode>=32&&t.keyCode<=126)return E(String.fromCharCode(t.keyCode))}}(t);r&&(t.preventDefault(),r.focus())}))}}function p(e,n,r){if(e){r||(r=e.querySelector(":focus"));for(var i=r;i&&"LI"!==i.tagName;)i=i.parentElement;if(i||(i=n?e.firstElementChild:e.lastElementChild),i){var l=i;do{n?(l=l.previousElementSibling)||(l=e.lastElementChild):(l=l.nextElementSibling)||(l=e.firstElementChild)}while(l!==i&&(t(l,o)||t(l,a)));if(l)return l.querySelector("a")}}}function m(e){var t=e.querySelector(f);return t&&t.querySelector("a")}function g(e){var t=e.querySelector("li a");return t=p(e,!0,t)}function h(){r=undefined,i="",l=null}function E(e){var t;return(e=e.toLowerCase())===r?t=n.querySelector(":focus"):(l||(l=n.querySelector(":focus")),t=l,r=r===undefined?e:""),i+=e,clearTimeout(c),c=setTimeout(h,850),function(e,t,n,r){var i,o,a=1===r.length,l=n;do{if(a)a=!1;else if(l){if(i){if(i===l)return;if(o){if(o===l)return}else o=l}else i=l;var c=(l.querySelector(".name")||l).textContent.toLowerCase();if(r.length<=c.length&&c.substring(0,r.length)===r)return l}}while(l=p(e,t,l))}(n,!1,t,r||i)}}(),function(){if(document.querySelector&&document.addEventListener){var t=document.body.querySelector(".upload");if(t){var a=t.querySelector("form");if(a){var f=a.querySelector(".file");if(f){var d=document.body.querySelector(".upload-type");if(d){var v,y="file",p="dirfile",m="innerdirfile",g=d.querySelector("."+y),h=d.querySelector("."+p),E=d.querySelector("."+m),b=g,S=Boolean(h),L=String.prototype.padStart?function(e,t,n){return e.padStart(t,n)}:function(e,t,n){var r=e.length;if(r>=t)return e;var i,o=t-r,a=Math.ceil(o/n.length);if(String.prototype.repeat)i=n.repeat(a);else{i="";for(var l=0;l<a;l++)i+=n}return i.length>o&&(i=i.substring(0,o)),i+e};if("https:"!==location.protocol||typeof FileSystemHandle===i||DataTransferItem.prototype.webkitGetAsEntry)v=function(t,n,r,i){var o=[],a=!1;if(!t||!t.length||!t[0].webkitGetAsEntry)return r(o,a);for(var l=[],c=0,u=t.length;c<u;c++){var s=t[c],f=s.webkitGetAsEntry();if(f)if(f.isFile){var d=s.getAsFile();o.push({file:d,relativePath:d.name})}else if(f.isDirectory){if(a=!0,!n)return i();l.push(f)}}(function v(t,n,r){var i=t.length,o=0;if(!i)return r();function a(){++o===i&&r()}function l(e,t,n){e.readEntries((function(r){if(!r.length)return n();v(r,t,(function(){l(e,t,n)}))}),n)}t.forEach((function(t){if(t.isFile){var r=t.fullPath;"/"===r[0]&&(r=r.substring(1)),t.file((function(e){n.push({file:e,relativePath:r}),a()}),(function(t){a(),e(t)}))}else if(t.isDirectory){l(t.createReader(),n,a)}}))})(l,o,(function(){r(o,a)}))};else{var w="directory",k={mode:"read"};v=function(t,n,r,i){function o(t,n,r){return Promise.all(t.map((function(t){var i=t.value;return"file"===i.kind?i.queryPermission(k).then((function(e){if("prompt"===e)return i.requestPermission(k)})).then((function(){return i.getFile()})).then((function(e){var t=r+e.name;n.push({file:e,relativePath:t})}))["catch"]((function(t){e(t)})):i.kind===w?new Promise((function(e){var t=[],a=i.values();function l(){t=null,a=null,e()}(function c(){a.next().then((function(e){e.done?t.length?o(t,n,r+i.name+"/").then(l):l():(t.push(e),c())}))})()})):void 0})))}var a=[],l=!1;if(!t||!t.length)return r(a,l);var c=Array.prototype.slice.call(t);Promise.all(c.map((function(e){return e.getAsFileSystemHandle()}))).then((function(e){if(e=e.filter(Boolean),(l=e.some((function(e){return e.kind===w})))&&!n)return i();o(e.map((function(e){return{value:e,done:!1}})),a,"").then((function(){r(a,l)}))}))}}var q=u,A=u;(function(){var e="hidden",t="active";function a(e,i){if(e!==b)return r(b,t),n(b=e,t),i&&(f.value=""),!0}function u(e){a(g,Boolean(e))&&(f.name=y,f.webkitdirectory=!1)}function v(){a(h,b===g)&&(f.name=p,f.webkitdirectory=!0)}function S(){a(E,b===g)&&(f.name=m,f.webkitdirectory=!0)}function L(e){switch(e.key){case c:case" ":if(e.ctrlKey||e.altKey||e.metaKey||e.shiftKey)break;if(e.preventDefault(),e.stopPropagation(),e.target===b)break;e.target.click()}}if(typeof f.webkitdirectory!==i){if(h&&r(h,e),E&&r(E,e),g&&(g.addEventListener("click",u),g.addEventListener("keydown",L)),h&&(h.addEventListener("click",v),h.addEventListener("keydown",L)),E&&(E.addEventListener("click",S),E.addEventListener("keydown",L)),s){var w="upload-type",k=sessionStorage.getItem(w);sessionStorage.removeItem(w),window.addEventListener(l,(function(){var e=f.name;e!==y&&sessionStorage.setItem(w,e)}),!1),k===p?h&&h.click():k===m&&E&&E.click()}g&&f.addEventListener("change",(function(e){if(b!==g){var t=e.target.files;if(t&&t.length){var n=Array.prototype.slice.call(t).every((function(e){return!e.webkitRelativePath||e.webkitRelativePath.indexOf("/")<0}));n&&u()}}})),q=function(){g&&b!==g&&(g.focus(),u(!0))},A=function(){h?b!==h&&(h.focus(),v()):E&&b!==E&&(E.focus(),S())}}else n(d,o)})();var D=function(){if(typeof FormData!==i){var e=!1,t=[],o="uploading",l="failed",c=document.body.querySelector(".upload-status"),u=c&&c.querySelector(".progress"),s=c&&c.querySelector(".warn .message");return function(i){i&&i.length&&(e?t.push(i):(e=!0,r(c,l),n(c,o),g(i)))}}function d(){u&&(u.style.width="")}function v(){if(t.length)return g(t.shift());e=!1,r(c,o)}function y(e){r(c,o),n(c,l),s&&(s.textContent=" - "+e.type),t.length=0}function p(){var t=this.status;t>=200&&t<=299?!e&&location.reload():y({type:this.statusText||t})}function m(e){if(e.lengthComputable){var t=100*e.loaded/e.total;u.style.width=t+"%"}}function g(e){var t=f.name,n=new FormData;e.forEach((function(e){var r;e.file?(r=e.relativePath,e=e.file):e.webkitRelativePath&&(r=e.webkitRelativePath),r||(r=e.name),n.append(t,e,r)}));var r=new XMLHttpRequest;r.addEventListener("error",d),r.addEventListener("error",y),r.addEventListener("abort",d),r.addEventListener("abort",y),r.addEventListener("load",d),r.addEventListener("load",v),r.addEventListener("load",p),u&&r.upload.addEventListener("progress",m),r.open(a.method,a.action),r.send(n)}}();D?(function(e){a.addEventListener("submit",(function(t){t.stopPropagation(),t.preventDefault();var n=Array.prototype.slice.call(f.files);e(n)})),f.addEventListener("change",(function(){var t=Array.prototype.slice.call(f.files);e(t)}))}(D),function(e){var t,n="text/plain",r="text.txt";Blob&&Blob.prototype.msClose?t=function(e){var t=new Blob([e],{type:n});return t.name=r,t}:File&&(t=function(e){return new File([e],r,{type:n})});var o=["hidden","radio","checkbox","button","reset","submit","image"];function a(t){q();var n,r,i,o,a=(n=new Date,r=String(1e4*n.getFullYear()+100*(n.getMonth()+1)+n.getDate()),i=String(1e4*n.getHours()+100*n.getMinutes()+n.getSeconds()),o=String(n.getMilliseconds()),"-"+(r=L(r,8,"0"))+"-"+(i=L(i,6,"0"))+"-"+L(o,3,"0"));t=t.map((function(e,t){var n=e.name,r=n.lastIndexOf(".");return r<0&&(r=n.length),{file:e,relativePath:n=n.substring(0,r)+a+"-"+t+n.substring(r)}})),e(t)}function l(e){var r,i;if(e.files&&e.files.length?r=Array.prototype.slice.call(e.files):e.items&&e.items.length?(i=Array.prototype.slice.call(e.items),r=i.map((function(e){return e.getAsFile()})).filter(Boolean)):r=[],r.length)a(r);else if(t&&i)for(var o=0,l=0,c=i.length;l<c;l++)e.types[l]===n&&(o++,i[l].getAsString((function(e){var n=t(e);r.push(n),r.length===o&&a(r)})))}document.documentElement.addEventListener("paste",(function(t){var n=t.target.tagName;if(!("INPUT"===n&&o.indexOf(t.target.type)<0)&&"TEXTAREA"!==n){var r=t.clipboardData;if(r){var a=r.items;a&&a.length?v(a,S,(function(t,n){t.length?1!==t.length||"image/png"!==t[0].file.type?(n?A():q(),e(t)):l({files:t=t.map((function(e){return e&&e.file}))}):l(r)}),(function(){typeof showUploadDirFailMessage!==i&&showUploadDirFailMessage()})):l(r)}}}))}(D)):document.documentElement.addEventListener("paste",(function(e){var t=e.clipboardData;t&&t.files&&t.files.length&&(q(),f.files=t.files,a.submit())})),function(e){var t=!1,o="dragging";function l(e){t||(e.stopPropagation(),e.preventDefault(),n(e.currentTarget,o))}document.body.addEventListener("dragstart",(function(){t=!0})),document.body.addEventListener("dragend",(function(){t=!1}));var c=document.documentElement;c.addEventListener("dragenter",l),c.addEventListener("dragover",l),c.addEventListener("dragleave",(function(e){e.target===e.currentTarget&&r(e.currentTarget,o)})),c.addEventListener("drop",(function(t){t.stopPropagation(),t.preventDefault(),r(t.currentTarget,o),f.value="",t.dataTransfer&&t.dataTransfer.files&&t.dataTransfer.files.length&&function(e,t,n,r){v(e.items,t,(function(t,r){0===t.length&&e.files&&e.files.length&&(t=Array.prototype.slice.call(e.files)),n(t,r)}),r)}(t.dataTransfer,S&&Boolean(e),(function(t,n){n?(A(),e(t)):(q(),e?e(t):(f.files=t,a.submit()))}),(function(){typeof showUploadDirFailMessage!==i&&showUploadDirFailMessage()}))}))}(D)}}}}}}()})(),function(){if(document.querySelector&&document.addEventListener&&document.body.classList){var e=document.body.querySelector("form.item-list-action");if(e){var t=e.querySelector(".actions");if(t){var n,r,i="selecting",o=document.body.closest?function(e,t){return e.closest(t)}:function(e,t){for(var n=t.toUpperCase(),r=e;r&&r.tagName!==n;)r=r.parentNode;return r};(function(){var n=t.querySelector(".start-select");n&&n.addEventListener("click",(function(t){t.preventDefault(),e.reset(),e.classList.add(i)}));var r=t.querySelector(".cancel-select");r&&r.addEventListener("click",(function(){e.classList.remove(i)}))})(),(n=e.querySelector("input.select-all"))&&n.addEventListener("change",(function(t){var n=t.target.checked,r=e.querySelectorAll(".item-list li:not(.none) label.select input");Array.isArray(r)||(r=Array.prototype.slice.call(r)),r.forEach((function(e){e.checked=n}))})),function(){if("function"==typeof confirmDelete){var e=t.querySelector(".delete");e&&e.addEventListener("click",(function(e){confirmDelete(this.form)||(e.preventDefault(),e.stopImmediatePropagation&&e.stopImmediatePropagation())}))}}(),(r=e.querySelector("button.delete"))&&r.addEventListener("click",(function(t){if(!t.defaultPrevented){t.preventDefault(),t.stopPropagation();var n=t.target.formAction,r=Array.prototype.slice.call(e.querySelectorAll("input:checked[type=checkbox][name=name]")),i=r.map((function(e){return"name="+encodeURIComponent(e.value)})).join("&"),a=new XMLHttpRequest;a.open("POST",n),a.setRequestHeader("Content-Type","application/x-www-form-urlencoded"),a.addEventListener("load",(function(){r.forEach((function(e){var t=o(e,"li"),n=t.parentNode;n&&n.removeChild(t)}))})),a.send(i)}})),function(){function t(){e.classList.remove(i)}var n=e.querySelectorAll("button[formaction]");Array.isArray(n)||(n=Array.prototype.slice.call(n)),n.forEach((function(e){e.addEventListener("click",t)}))}()}}}}();