(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-0e4c0bb5"],{1290:function(e,t,n){},"28a5":function(e,t,n){"use strict";var i=n("aae3"),a=n("cb7c"),r=n("ebd6"),c=n("0390"),s=n("9def"),o=n("5f1b"),l=n("520a"),u=n("79e5"),d=Math.min,h=[].push,f="split",p="length",v="lastIndex",g=4294967295,b=!u(function(){RegExp(g,"y")});n("214f")("split",2,function(e,t,n,u){var m;return m="c"=="abbc"[f](/(b)*/)[1]||4!="test"[f](/(?:)/,-1)[p]||2!="ab"[f](/(?:ab)*/)[p]||4!="."[f](/(.?)(.?)/)[p]||"."[f](/()()/)[p]>1||""[f](/.?/)[p]?function(e,t){var a=String(this);if(void 0===e&&0===t)return[];if(!i(e))return n.call(a,e,t);var r,c,s,o=[],u=(e.ignoreCase?"i":"")+(e.multiline?"m":"")+(e.unicode?"u":"")+(e.sticky?"y":""),d=0,f=void 0===t?g:t>>>0,b=new RegExp(e.source,u+"g");while(r=l.call(b,a)){if(c=b[v],c>d&&(o.push(a.slice(d,r.index)),r[p]>1&&r.index<a[p]&&h.apply(o,r.slice(1)),s=r[0][p],d=c,o[p]>=f))break;b[v]===r.index&&b[v]++}return d===a[p]?!s&&b.test("")||o.push(""):o.push(a.slice(d)),o[p]>f?o.slice(0,f):o}:"0"[f](void 0,0)[p]?function(e,t){return void 0===e&&0===t?[]:n.call(this,e,t)}:n,[function(n,i){var a=e(this),r=void 0==n?void 0:n[t];return void 0!==r?r.call(n,a,i):m.call(String(a),n,i)},function(e,t){var i=u(m,e,this,t,m!==n);if(i.done)return i.value;var l=a(e),h=String(this),f=r(l,RegExp),p=l.unicode,v=(l.ignoreCase?"i":"")+(l.multiline?"m":"")+(l.unicode?"u":"")+(b?"y":"g"),y=new f(b?l:"^(?:"+l.source+")",v),x=void 0===t?g:t>>>0;if(0===x)return[];if(0===h.length)return null===o(y,h)?[h]:[];var j=0,S=0,w=[];while(S<h.length){y.lastIndex=b?S:0;var k,T=o(y,b?h:h.slice(S));if(null===T||(k=d(s(y.lastIndex+(b?0:S)),h.length))===j)S=c(h,S,p);else{if(w.push(h.slice(j,S)),w.length===x)return w;for(var O=1;O<=T.length-1;O++)if(w.push(T[O]),w.length===x)return w;S=j=k}}return w.push(h.slice(j)),w}]})},"4f37":function(e,t,n){"use strict";n("aa77")("trim",function(e){return function(){return e(this,3)}})},"5b47":function(e,t,n){"use strict";var i=n("1290"),a=n.n(i);a.a},"84f8":function(e,t,n){"use strict";n.r(t);var i=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-container",[n("el-header",[n("el-card",{attrs:{height:60,"body-style":{padding:"10px 10px 10px 20px"}}},[n("el-row",[n("el-col",{staticStyle:{float:"right"},attrs:{span:3}},[n("el-button",{staticStyle:{float:"right"},on:{click:e.handleFullScreen}},[n("v-icon",{attrs:{small:""}},[e._v("fullscreen")])],1)],1)],1)],1)],1),n("el-container",[n("div",{staticStyle:{width:"100%"},attrs:{id:"shell"}})])],1)},a=[],r=(n("7f7f"),n("4f37"),n("28a5"),n("d225")),c=n("b0b4"),s=n("308d"),o=n("6bb5"),l=n("4e2b"),u=n("9ab4"),d=n("60a3"),h=n("d70b"),f=h["a"].url.basicUrl+"/api/v1/b",p=function(e){function t(){var e;return Object(r["a"])(this,t),e=Object(s["a"])(this,Object(o["a"])(t).apply(this,arguments)),e.height="500px",e}return Object(l["a"])(t,e),Object(c["a"])(t,[{key:"created",value:function(){}},{key:"mounted",value:function(){this.height=.8*window.innerHeight+"px",this.renderTerminal()}},{key:"renderTerminal",value:function(){var e=this;window.jQuery.terminal?jQuery(function(t,n){t("#shell").terminal(function(i,a){if(""!=i){var r="COMMANDS:\n    exit        exit fullscreen\n    call        Call a service endpoint using rpc\n    health      Query the health of a service\n    list        List items in registry\n    get         Get item from registry\n";try{var c=i.trim().split(/\s+/);switch(c[0]){case"exit":e.$xools.toggleFullScreen("shell");break;case"help":a.echo(r);break;case"list":if(1==c.length||"services"!=c[1])return void a.echo("COMMANDS:\n    services    List services in registry\n");t.ajax({dataType:"json",contentType:"application/json",url:f+"/services",data:{},success:function(e){for(var t=[],n=0;n<e.data.length;n++)t.push(e.data[n].name);a.echo(t.join("\n"))}});break;case"get":if(c.length<3||"service"!=c[1])return void a.echo("COMMANDS:\n    service    Get service from registry\n");t.ajax({dataType:"json",contentType:"application/json",url:f+"/service/"+c[2],data:{},success:function(e){if(0!=e.data.length){a.echo("service\t"+c[2]),a.echo(" ");for(var i={},r=0;r<e.data.length;r++){var s=e.data[r];a.echo("\nversion "+s.version),a.echo(" "),a.echo("Id\tAddress\tPort\tMetadata\n");for(var o=function(e){var n=s.nodes[e],i=[];t.each(n.metadata,function(e,t){i.push(e+"="+t)}),a.echo(n.id+"\t"+n.address+"\t"+n.port+"\t"+i.join(","))},l=0;l<s.nodes.length;l++)o(l);a.echo(" ");for(var u=0;s.endpoints&&u<s.endpoints.length;u++)i[s.endpoints[u].name]==n&&(i[s.endpoints[u].name]=s.endpoints[u])}t.each(i,function(e,n){a.echo("Endpoint: "+e);var i=[];t.each(n.metadata,function(e,t){i.push(e+"="+t)}),a.echo("Metadata: "+i.join(","))})}}});break;case"health":if(c.length<2)return void a.echo("USAGE:\n    health [service]");t.ajax({dataType:"json",contentType:"application/json",url:f+"/service/"+c[1],data:{},success:function(e){a.echo("service\t"+c[1]),a.echo(" ");for(var n=0;n<e.data.length;n++){var i=e.data[n];a.echo("\nversion "+i.version),a.echo(" "),a.echo("Id\tAddress:Port\tStatus\n");for(var r=0;r<i.nodes.length;r++){var s=i.nodes[r];t.ajax({dataType:"json",url:f+"/health",data:{service:i.name,address:s.address+":"+s.port},success:function(e){a.echo(s.id+"\t"+s.address+":"+s.port+"\t"+e.data.status)},error:function(e){a.echo(s.id+"\t"+s.address+":"+s.port+"\t"+e.data.status)}})}a.echo(" ")}}});break;case"call":if(c.length<3)return void a.echo("USAGE:\n    call [service] [endpoint] [request]");var s="{}";c.length>3&&(s=c.slice(3).join(" ")),t.ajax({method:"post",endpoint:"POST",dataType:"json",contentType:"application/json",url:f+"/rpc",data:JSON.stringify({service:c[1],endpoint:c[2],request:s}),success:function(e){a.echo(JSON.stringify(e,null,2))}});break;default:a.echo(i+": command not found"),a.echo(r)}}catch(o){a.error(new String(o))}}else a.echo("")},{greetings:"",name:"micro_cli",height:500,prompt:"micro:~$ "})}):setTimeout(this.renderTerminal,500)}},{key:"handleFullScreen",value:function(){this.$xools.toggleFullScreen("shell")}}]),t}(d["c"]);p=u["a"]([Object(d["a"])({components:{}})],p);var v=p,g=v,b=(n("5b47"),n("2877")),m=Object(b["a"])(g,i,a,!1,null,"34beddc2",null);t["default"]=m.exports},aa77:function(e,t,n){var i=n("5ca1"),a=n("be13"),r=n("79e5"),c=n("fdef"),s="["+c+"]",o="​",l=RegExp("^"+s+s+"*"),u=RegExp(s+s+"*$"),d=function(e,t,n){var a={},s=r(function(){return!!c[e]()||o[e]()!=o}),l=a[e]=s?t(h):c[e];n&&(a[n]=l),i(i.P+i.F*s,"String",a)},h=d.trim=function(e,t){return e=String(a(e)),1&t&&(e=e.replace(l,"")),2&t&&(e=e.replace(u,"")),e};e.exports=d},aae3:function(e,t,n){var i=n("d3f4"),a=n("2d95"),r=n("2b4c")("match");e.exports=function(e){var t;return i(e)&&(void 0!==(t=e[r])?!!t:"RegExp"==a(e))}},fdef:function(e,t){e.exports="\t\n\v\f\r   ᠎             　\u2028\u2029\ufeff"}}]);