"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[21],{68456:function(e,n,a){a.d(n,{QT:function(){return o},YH:function(){return c},mo:function(){return s}});var t=a(61889),i=a(75952),r=a(80184),s=function(){return(0,r.jsxs)(t.ZP,{container:!0,columnGap:1,children:[(0,r.jsx)(t.ZP,{children:(0,r.jsx)(i.gyG,{width:"16px",height:"16px"})}),(0,r.jsx)(t.ZP,{item:!0,children:"Open ID"})]})},o=function(){return(0,r.jsxs)(t.ZP,{container:!0,columnGap:1,children:[(0,r.jsx)(t.ZP,{children:(0,r.jsx)(i.vcZ,{width:"16px",height:"16px"})}),(0,r.jsx)(t.ZP,{item:!0,children:"LDAP / Active Directory"})]})},c=function(){return(0,r.jsxs)(t.ZP,{container:!0,columnGap:1,children:[(0,r.jsx)(t.ZP,{children:(0,r.jsx)(i.oyc,{width:"16px",height:"16px"})}),(0,r.jsx)(t.ZP,{item:!0,children:"Built-in"})]})}},37021:function(e,n,a){a.r(n);var t=a(93433),i=a(29439),r=a(1413),s=a(72791),o=a(78687),c=a(51691),l=a(20890),d=a(96040),u=a(13400),m=a(75952),h=a(61889),p=a(11135),x=a(25787),v=a(20165),f=a(3579),_=a(23814),Z=a(68456),j=a(84741),g=a(40968),D=a(87995),b=a(41320),N=a(83679),S=a(21435),I=a(37516),k=a(40306),y=a(81207),A=a(42419),R=a(27247),C=a(50896),w=a(80184);function P(){return null}var F=(0,o.$j)((function(e){return{loadingTenant:e.tenants.loadingTenant,selectedTenant:e.tenants.currentTenant,tenant:e.tenants.tenantInfo}}),null);n.default=(0,x.Z)((function(e){return(0,p.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)({adUserDnRows:{display:"flex",marginBottom:10},buttonTray:{marginLeft:10,display:"flex",height:38,"& button":{background:"#EAEAEA"}}},_.oZ),_.bK),{},{loaderAlign:{textAlign:"center"}},_.Bz),_.QV),_.DF),_.oO),_.AK))}))(F((function(e){var n=e.classes,a=(0,b.TL)(),r=(0,o.v9)((function(e){return e.tenants.tenantInfo})),p=(0,o.v9)((function(e){return e.tenants.loadingTenant})),x=(0,s.useState)(!1),_=(0,i.Z)(x,2),F=_[0],L=_[1],T=(0,s.useState)(!1),U=(0,i.Z)(T,2),O=U[0],B=U[1],G=(0,s.useState)("Built-in"),q=(0,i.Z)(G,2),z=q[0],K=q[1],E=(0,s.useState)(""),H=(0,i.Z)(E,2),V=H[0],M=H[1],Q=(0,s.useState)(""),Y=(0,i.Z)(Q,2),W=Y[0],$=Y[1],J=(0,s.useState)(""),X=(0,i.Z)(J,2),ee=X[0],ne=X[1],ae=(0,s.useState)(!1),te=(0,i.Z)(ae,2),ie=te[0],re=te[1],se=(0,s.useState)(""),oe=(0,i.Z)(se,2),ce=oe[0],le=oe[1],de=(0,s.useState)(""),ue=(0,i.Z)(de,2),me=ue[0],he=ue[1],pe=(0,s.useState)(""),xe=(0,i.Z)(pe,2),ve=xe[0],fe=xe[1],_e=(0,s.useState)(""),Ze=(0,i.Z)(_e,2),je=Ze[0],ge=Ze[1],De=(0,s.useState)(""),be=(0,i.Z)(De,2),Ne=be[0],Se=be[1],Ie=(0,s.useState)(""),ke=(0,i.Z)(Ie,2),ye=ke[0],Ae=ke[1],Re=(0,s.useState)(!1),Ce=(0,i.Z)(Re,2),we=Ce[0],Pe=Ce[1],Fe=(0,s.useState)(""),Le=(0,i.Z)(Fe,2),Te=Le[0],Ue=Le[1],Oe=(0,s.useState)(""),Be=(0,i.Z)(Oe,2),Ge=Be[0],qe=Be[1],ze=(0,s.useState)(""),Ke=(0,i.Z)(ze,2),Ee=Ke[0],He=Ke[1],Ve=(0,s.useState)(""),Me=(0,i.Z)(Ve,2),Qe=Me[0],Ye=Me[1],We=(0,s.useState)(!1),$e=(0,i.Z)(We,2),Je=$e[0],Xe=$e[1],en=(0,s.useState)(!1),nn=(0,i.Z)(en,2),an=nn[0],tn=nn[1],rn=(0,s.useState)(!1),sn=(0,i.Z)(rn,2),on=sn[0],cn=sn[1],ln=(0,s.useState)([""]),dn=(0,i.Z)(ln,2),un=dn[0],mn=dn[1],hn=(0,s.useState)([""]),pn=(0,i.Z)(hn,2),xn=pn[0],vn=pn[1],fn=(0,s.useState)({}),_n=(0,i.Z)(fn,2),Zn=_n[0],jn=_n[1],gn=function(e){jn((0,j.he)(Zn,e))},Dn=(0,s.useState)(!1),bn=(0,i.Z)(Dn,2),Nn=bn[0],Sn=bn[1];(0,s.useEffect)((function(){var e=[];"OpenID"===z&&(e=[].concat((0,t.Z)(e),[{fieldKey:"openID_CONFIGURATION_URL",required:!0,value:V},{fieldKey:"openID_clientID",required:!0,value:W},{fieldKey:"openID_secretID",required:!0,value:ee},{fieldKey:"openID_claimName",required:!0,value:me}])),"AD"===z&&(e=[].concat((0,t.Z)(e),[{fieldKey:"AD_URL",required:!0,value:je},{fieldKey:"ad_lookupBindDN",required:!0,value:Ne}]));var n=(0,g.R)(e);Sn(0===Object.keys(n).length),jn(n)}),[z,V,W,ee,me,je,Ne]);var In=(0,s.useCallback)((function(){y.Z.invoke("GET","/api/v1/namespaces/".concat(null===r||void 0===r?void 0:r.namespace,"/tenants/").concat(null===r||void 0===r?void 0:r.name,"/identity-provider")).then((function(e){e&&(e.oidc?(K("OpenID"),M(e.oidc.configuration_url),$(e.oidc.client_id),ne(e.oidc.secret_id),le(e.oidc.callback_url),he(e.oidc.claim_name),fe(e.oidc.scopes)):e.active_directory&&(K("AD"),ge(e.active_directory.url),Se(e.active_directory.lookup_bind_dn),Ae(e.active_directory.lookup_bind_password),Ue(e.active_directory.user_dn_search_base_dn),qe(e.active_directory.user_dn_search_filter),He(e.active_directory.group_search_base_dn),Ye(e.active_directory.group_search_filter),Xe(e.active_directory.skip_tls_verification),tn(e.active_directory.server_insecure),cn(e.active_directory.server_start_tls)))})).catch((function(e){a((0,D.Ih)(e))}))}),[r,a]);(0,s.useEffect)((function(){r&&In()}),[r,In]);return(0,w.jsxs)(s.Fragment,{children:[(0,w.jsx)(k.Z,{title:"Save and Restart",confirmText:"Restart",cancelText:"Cancel",titleIcon:(0,w.jsx)(m.EjK,{}),isLoading:F,onClose:function(){return B(!1)},isOpen:O,onConfirm:function(){L(!0);var e={};switch(z){case"AD":e.active_directory={url:je,lookup_bind_dn:Ne,lookup_bind_password:ye,user_dn_search_base_dn:Te,user_dn_search_filter:Ge,group_search_base_dn:Ee,group_search_filter:Qe,skip_tls_verification:Je,server_insecure:an,server_start_tls:on};break;case"OpenID":e.oidc={configuration_url:V,client_id:W,secret_id:ee,callback_url:ce,claim_name:me,scopes:ve}}y.Z.invoke("POST","/api/v1/namespaces/".concat(null===r||void 0===r?void 0:r.namespace,"/tenants/").concat(null===r||void 0===r?void 0:r.name,"/identity-provider"),e).then((function(){L(!1),B(!1),In()})).catch((function(e){a((0,D.Ih)(e)),L(!1)}))},confirmationContent:(0,w.jsx)(c.Z,{children:"Are you sure you want to save the changes and restart the service?"})}),p?(0,w.jsx)("div",{className:n.loaderAlign,children:(0,w.jsx)(m.aNw,{})}):(0,w.jsxs)(s.Fragment,{children:[(0,w.jsxs)(h.ZP,{item:!0,xs:12,children:[(0,w.jsx)("h1",{className:n.sectionTitle,children:"Identity Provider"}),(0,w.jsx)(P,{})]}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.protocolRadioOptions,paddingBottom:1,children:(0,w.jsx)(N.Z,{currentSelection:z,id:"idp-options",name:"idp-options",label:"Protocol",onChange:function(e){K(e.target.value)},selectorOptions:[{label:(0,w.jsx)(Z.YH,{}),value:"Built-in"},{label:(0,w.jsx)(Z.mo,{}),value:"OpenID"},{label:(0,w.jsx)(Z.QT,{}),value:"AD"}]})}),"OpenID"===z&&(0,w.jsxs)(s.Fragment,{children:[(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"openID_CONFIGURATION_URL",name:"openID_CONFIGURATION_URL",onChange:function(e){M(e.target.value),gn("openID_CONFIGURATION_URL")},label:"Configuration URL",value:V,placeholder:"https://your-identity-provider.com/.well-known/openid-configuration",error:Zn.openID_CONFIGURATION_URL||"",required:!0})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"openID_clientID",name:"openID_clientID",onChange:function(e){$(e.target.value),gn("openID_clientID")},label:"Client ID",value:W,error:Zn.openID_clientID||"",required:!0})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{type:ie?"text":"password",id:"openID_secretID",name:"openID_secretID",onChange:function(e){ne(e.target.value),gn("openID_secretID")},label:"Secret ID",value:ee,error:Zn.openID_secretID||"",required:!0,overlayIcon:ie?(0,w.jsx)(v.Z,{}):(0,w.jsx)(f.Z,{}),overlayAction:function(){return re(!ie)}})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"openID_callbackURL",name:"openID_callbackURL",onChange:function(e){le(e.target.value),gn("openID_callbackURL")},label:"Callback URL",value:ce,placeholder:"https://your-console-endpoint:9443/oauth_callback",error:Zn.openID_callbackURL||""})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"openID_claimName",name:"openID_claimName",onChange:function(e){he(e.target.value),gn("openID_claimName")},label:"Claim Name",value:me,error:Zn.openID_claimName||"",required:!0})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"openID_scopes",name:"openID_scopes",onChange:function(e){fe(e.target.value),gn("openID_scopes")},label:"Scopes",value:ve})})]}),"AD"===z&&(0,w.jsxs)(s.Fragment,{children:[(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"AD_URL",name:"AD_URL",onChange:function(e){ge(e.target.value),gn("AD_URL")},label:"LDAP Server Address",value:je,placeholder:"ldap-server:636",error:Zn.AD_URL||"",required:!0})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(I.Z,{value:"ad_skipTLS",id:"ad_skipTLS",name:"ad_skipTLS",checked:Je,onChange:function(e){var n=e.target.checked;Xe(n)},label:"Skip TLS Verification"})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(I.Z,{value:"ad_serverInsecure",id:"ad_serverInsecure",name:"ad_serverInsecure",checked:an,onChange:function(e){var n=e.target.checked;tn(n)},label:"Server Insecure"})}),an?(0,w.jsxs)(h.ZP,{item:!0,xs:12,children:[(0,w.jsx)(l.Z,{className:n.error,variant:"caption",display:"block",gutterBottom:!0,children:"Warning: All traffic with Active Directory will be unencrypted"}),(0,w.jsx)("br",{})]}):null,(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(I.Z,{value:"ad_serverStartTLS",id:"ad_serverStartTLS",name:"ad_serverStartTLS",checked:on,onChange:function(e){var n=e.target.checked;cn(n)},label:"Start TLS connection to AD/LDAP server"})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"ad_lookupBindDN",name:"ad_lookupBindDN",onChange:function(e){Se(e.target.value),gn("ad_lookupBindDN")},label:"Lookup Bind DN",value:Ne,placeholder:"cn=admin,dc=min,dc=io",error:Zn.ad_lookupBindDN||"",required:!0})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{type:we?"text":"password",id:"ad_lookupBindPassword",name:"ad_lookupBindPassword",onChange:function(e){Ae(e.target.value)},label:"Lookup Bind Password",value:ye,placeholder:"admin",overlayIcon:we?(0,w.jsx)(v.Z,{}):(0,w.jsx)(f.Z,{}),overlayAction:function(){return Pe(!we)}})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"ad_userDNSearchBaseDN",name:"ad_userDNSearchBaseDN",onChange:function(e){Ue(e.target.value)},label:"User DN Search Base DN",value:Te,placeholder:"dc=min,dc=io"})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"ad_userDNSearchFilter",name:"ad_userDNSearchFilter",onChange:function(e){qe(e.target.value)},label:"User DN Search Filter",value:Ge,placeholder:"(sAMAcountName=%s)"})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"ad_groupSearchBaseDN",name:"ad_groupSearchBaseDN",onChange:function(e){He(e.target.value)},label:"Group Search Base DN",value:Ee,placeholder:"ou=hwengg,dc=min,dc=io;ou=swengg,dc=min,dc=io"})}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.formFieldRow,children:(0,w.jsx)(S.Z,{id:"ad_groupSearchFilter",name:"ad_groupSearchFilter",onChange:function(e){Ye(e.target.value)},label:"Group Search Filter",value:Qe,placeholder:"(&(objectclass=groupOfNames)(member=%s))"})})]}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.buttonContainer,children:(0,w.jsx)(m.zxk,{id:"save-idp",type:"submit",variant:"callAction",color:"primary",disabled:!Nn||F,onClick:function(){return B(!0)},label:"Save"})}),"AD"===z&&(0,w.jsxs)(s.Fragment,{children:[(0,w.jsx)(C.Z,{children:"User & Group management"}),(0,w.jsx)("br",{}),(0,w.jsxs)("fieldset",{className:n.fieldGroup,children:[(0,w.jsx)("legend",{className:n.descriptionText,children:"List of user DNs (Distinguished Names) to be added as Tenant Administrators"}),(0,w.jsx)(h.ZP,{item:!0,xs:12,children:un.map((function(e,a){return(0,w.jsx)(s.Fragment,{children:(0,w.jsxs)("div",{className:n.adUserDnRows,children:[(0,w.jsx)(S.Z,{id:"ad-userdn-".concat(a.toString()),label:"",placeholder:"",name:"ad-userdn-".concat(a.toString()),value:un[a],onChange:function(e){mn(un.map((function(n,t){return t===a?e.target.value:n})))},index:a,error:Zn["ad-userdn-".concat(a.toString())]||""},"csv-ad-userdn-".concat(a.toString())),(0,w.jsxs)("div",{className:n.buttonTray,children:[(0,w.jsx)(d.Z,{title:"Add User","aria-label":"add",children:(0,w.jsx)(u.Z,{size:"small",onClick:function(){mn([].concat((0,t.Z)(un),[""]))},children:(0,w.jsx)(A.Z,{})})}),(0,w.jsx)(d.Z,{title:"Remove","aria-label":"add",children:(0,w.jsx)(u.Z,{size:"small",style:{marginLeft:16},onClick:function(){un.length>1&&mn(un.filter((function(e,n){return n!==a})))},children:(0,w.jsx)(R.Z,{})})})]})]})},"identityField-".concat(a.toString()))}))})]}),(0,w.jsxs)("fieldset",{className:n.fieldGroup,children:[(0,w.jsx)("legend",{className:n.descriptionText,children:"List of group DNs (Distinguished Names) to be added as Tenant Administrators"}),(0,w.jsx)(h.ZP,{item:!0,xs:12,children:xn.map((function(e,a){return(0,w.jsx)(s.Fragment,{children:(0,w.jsxs)("div",{className:n.adUserDnRows,children:[(0,w.jsx)(S.Z,{id:"ad-groupdn-".concat(a.toString()),label:"",placeholder:"",name:"ad-groupdn-".concat(a.toString()),value:xn[a],onChange:function(e){vn(xn.map((function(n,t){return t===a?e.target.value:n})))},index:a,error:Zn["ad-groupdn-".concat(a.toString())]||""},"csv-ad-groupdn-".concat(a.toString())),(0,w.jsxs)("div",{className:n.buttonTray,children:[(0,w.jsx)(d.Z,{title:"Add Group","aria-label":"add",children:(0,w.jsx)(u.Z,{size:"small",onClick:function(){vn([].concat((0,t.Z)(xn),[""]))},children:(0,w.jsx)(A.Z,{})})}),(0,w.jsx)(d.Z,{title:"Remove","aria-label":"add",children:(0,w.jsx)(u.Z,{size:"small",style:{marginLeft:16},onClick:function(){xn.length>1&&vn(xn.filter((function(e,n){return n!==a})))},children:(0,w.jsx)(R.Z,{})})})]})]})},"identityField-".concat(a.toString()))}))})]}),(0,w.jsx)("br",{}),(0,w.jsx)(h.ZP,{item:!0,xs:12,className:n.buttonContainer,children:(0,w.jsx)(m.zxk,{id:"add-additional-dns",type:"submit",variant:"callAction",disabled:!Nn||F,onClick:function(){return function(){L(!0);var e={};"AD"===z&&(e={user_dns:un.filter((function(e){return""!==e.trim()})),group_dns:xn.filter((function(e){return""!==e.trim()}))});y.Z.invoke("POST","/api/v1/namespaces/".concat(null===r||void 0===r?void 0:r.namespace,"/tenants/").concat(null===r||void 0===r?void 0:r.name,"/set-administrators"),e).then((function(){L(!1),vn([""]),mn([""]),In(),a((0,D.y1)("Administrators added successfully"))})).catch((function(e){a((0,D.Ih)(e)),L(!1)}))}()},label:"Add additional DNs"})})]})]})]})})))},42419:function(e,n,a){var t=a(64836);n.Z=void 0;var i=t(a(45649)),r=a(80184),s=(0,i.default)((0,r.jsx)("path",{d:"M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z"}),"Add");n.Z=s},27247:function(e,n,a){var t=a(64836);n.Z=void 0;var i=t(a(45649)),r=a(80184),s=(0,i.default)((0,r.jsx)("path",{d:"M6 19c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V7H6v12zM19 4h-3.5l-1-1h-5l-1 1H5v2h14V4z"}),"Delete");n.Z=s}}]);
//# sourceMappingURL=21.f346171d.chunk.js.map