package main

import (
	"encoding/json"
	"file-extractor/textExtract"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	e "file-extractor/errors"

	"code.sajari.com/docconv"
)

var resp = httptest.NewRecorder()

func TestScrapper(t *testing.T) {
	urls := []string{"https://www.google.com", "https://www.facebook.com/"}
	for _, url := range urls {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Errorf("Request error %v", err)
		}
		ScrappingProcess(resp, req, url)
	}
	if resp.Code != 200 {
		t.Errorf("we expected 200, but it returns %v", resp.Code)
	}
}

func TestPdf(t *testing.T) {
	pdfNames := []string{"./test_files/sample.pdf", "./test_files/test.pdf"}
	for _, pdf := range pdfNames {
		file, err := os.Open(pdf)
		if err != nil && ErrorCheck(resp, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue]) {
			return
		}

		_, _, err = docconv.ConvertPDF(file)
		if err != nil && ErrorCheck(resp, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue]) {
			return
		}
	}
	if resp.Code != 200 {
		t.Errorf("we expected 200, but it returns %v", resp.Code)
	}
}

func TestWord(t *testing.T) {
	WordNames := []string{"./test_files/word.docx"}
	for _, word := range WordNames {
		file, err := os.Open(word)
		if err != nil && ErrorCheck(resp, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue]) {
			return
		}
		switch filepath.Ext(word) {
		case ".doc":
			_, _, err = docconv.ConvertDoc(file)
			if err != nil && ErrorCheck(resp, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue]) {
				return
			}
		case ".docx":
			_, _, err = docconv.ConvertDocx(file)
			if err != nil && ErrorCheck(resp, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue]) {
				return
			}
		}
	}
	if resp.Code != 200 {
		t.Errorf("we expected 200, but it returns %v", resp.Code)
	}
}

func TestXml(t *testing.T) {
	XmlFiles := []string{"./test_files/test.xml", "./test_files/facebook.html"}

	for _, word := range XmlFiles {
		file, err := os.Open(word)
		if err != nil && ErrorCheck(resp, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue]) {
			return
		}

		_, err = docconv.XMLToText(file, []string{}, []string{}, false)
		if err != nil && ErrorCheck(resp, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue]) {
			return
		}
	}
	if resp.Code != 200 {
		t.Errorf("we expected 200, but it returns %v", resp.Code)
	}
}

func Test_HtmlExtractor(t *testing.T) {
	el := `<!doctype html>
<html lang="en-US" prefix="og: https://ogp.me/ns#" class="no-js" itemtype="https://schema.org/Blog" itemscope>
<head><meta charset="UTF-8"><script>if(navigator.userAgent.match(/MSIE|Internet Explorer/i)||navigator.userAgent.match(/Trident\/7\..*?rv:11/i)){var href=document.location.href;if(!href.match(/[?&]nowprocket/)){if(href.indexOf("?")==-1){if(href.indexOf("#")==-1){document.location.href=href+"?nowprocket=1"}else{document.location.href=href.replace("#","?nowprocket=1#")}}else{if(href.indexOf("#")==-1){document.location.href=href+"&nowprocket=1"}else{document.location.href=href.replace("#","&nowprocket=1#")}}}}</script><script>class RocketLazyLoadScripts{constructor(){this.v="1.2.3",this.triggerEvents=["keydown","mousedown","mousemove","touchmove","touchstart","touchend","wheel"],this.userEventHandler=this._triggerListener.bind(this),this.touchStartHandler=this._onTouchStart.bind(this),this.touchMoveHandler=this._onTouchMove.bind(this),this.touchEndHandler=this._onTouchEnd.bind(this),this.clickHandler=this._onClick.bind(this),this.interceptedClicks=[],window.addEventListener("pageshow",t=>{this.persisted=t.persisted}),window.addEventListener("DOMContentLoaded",()=>{this._preconnect3rdParties()}),this.delayedScripts={normal:[],async:[],defer:[]},this.trash=[],this.allJQueries=[]}_addUserInteractionListener(t){if(document.hidden){t._triggerListener();return}this.triggerEvents.forEach(e=>window.addEventListener(e,t.userEventHandler,{passive:!0})),window.addEventListener("touchstart",t.touchStartHandler,{passive:!0}),window.addEventListener("mousedown",t.touchStartHandler),document.addEventListener("visibilitychange",t.userEventHandler)}_removeUserInteractionListener(){this.triggerEvents.forEach(t=>window.removeEventListener(t,this.userEventHandler,{passive:!0})),document.removeEventListener("visibilitychange",this.userEventHandler)}_onTouchStart(t){"HTML"!==t.target.tagName&&(window.addEventListener("touchend",this.touchEndHandler),window.addEventListener("mouseup",this.touchEndHandler),window.addEventListener("touchmove",this.touchMoveHandler,{passive:!0}),window.addEventListener("mousemove",this.touchMoveHandler),t.target.addEventListener("click",this.clickHandler),this._renameDOMAttribute(t.target,"onclick","rocket-onclick"),this._pendingClickStarted())}_onTouchMove(t){window.removeEventListener("touchend",this.touchEndHandler),window.removeEventListener("mouseup",this.touchEndHandler),window.removeEventListener("touchmove",this.touchMoveHandler,{passive:!0}),window.removeEventListener("mousemove",this.touchMoveHandler),t.target.removeEventListener("click",this.clickHandler),this._renameDOMAttribute(t.target,"rocket-onclick","onclick"),this._pendingClickFinished()}_onTouchEnd(t){window.removeEventListener("touchend",this.touchEndHandler),window.removeEventListener("mouseup",this.touchEndHandler),window.removeEventListener("touchmove",this.touchMoveHandler,{passive:!0}),window.removeEventListener("mousemove",this.touchMoveHandler)}_onClick(t){t.target.removeEventListener("click",this.clickHandler),this._renameDOMAttribute(t.target,"rocket-onclick","onclick"),this.interceptedClicks.push(t),t.preventDefault(),t.stopPropagation(),t.stopImmediatePropagation(),this._pendingClickFinished()}_replayClicks(){window.removeEventListener("touchstart",this.touchStartHandler,{passive:!0}),window.removeEventListener("mousedown",this.touchStartHandler),this.interceptedClicks.forEach(t=>{t.target.dispatchEvent(new MouseEvent("click",{view:t.view,bubbles:!0,cancelable:!0}))})}_waitForPendingClicks(){return new Promise(t=>{this._isClickPending?this._pendingClickFinished=t:t()})}_pendingClickStarted(){this._isClickPending=!0}_pendingClickFinished(){this._isClickPending=!1}_renameDOMAttribute(t,e,r){t.hasAttribute&&t.hasAttribute(e)&&(event.target.setAttribute(r,event.target.getAttribute(e)),event.target.removeAttribute(e))}_triggerListener(){this._removeUserInteractionListener(this),"loading"===document.readyState?document.addEventListener("DOMContentLoaded",this._loadEverythingNow.bind(this)):this._loadEverythingNow()}_preconnect3rdParties(){let t=[];document.querySelectorAll("script[type=rocketlazyloadscript]").forEach(e=>{if(e.hasAttribute("src")){let r=new URL(e.src).origin;r!==location.origin&&t.push({src:r,crossOrigin:e.crossOrigin||"module"===e.getAttribute("data-rocket-type")})}}),t=[...new Map(t.map(t=>[JSON.stringify(t),t])).values()],this._batchInjectResourceHints(t,"preconnect")}async _loadEverythingNow(){this.lastBreath=Date.now(),this._delayEventListeners(this),this._delayJQueryReady(this),this._handleDocumentWrite(),this._registerAllDelayedScripts(),this._preloadAllScripts(),await this._loadScriptsFromList(this.delayedScripts.normal),await this._loadScriptsFromList(this.delayedScripts.defer),await this._loadScriptsFromList(this.delayedScripts.async);try{await this._triggerDOMContentLoaded(),await this._triggerWindowLoad()}catch(t){console.error(t)}window.dispatchEvent(new Event("rocket-allScriptsLoaded")),this._waitForPendingClicks().then(()=>{this._replayClicks()}),this._emptyTrash()}_registerAllDelayedScripts(){document.querySelectorAll("script[type=rocketlazyloadscript]").forEach(t=>{t.hasAttribute("data-rocket-src")?t.hasAttribute("async")&&!1!==t.async?this.delayedScripts.async.push(t):t.hasAttribute("defer")&&!1!==t.defer||"module"===t.getAttribute("data-rocket-type")?this.delayedScripts.defer.push(t):this.delayedScripts.normal.push(t):this.delayedScripts.normal.push(t)})}async _transformScript(t){return new Promise((await this._littleBreath(),navigator.userAgent.indexOf("Firefox/")>0||""===navigator.vendor)?e=>{let r=document.createElement("script");[...t.attributes].forEach(t=>{let e=t.nodeName;"type"!==e&&("data-rocket-type"===e&&(e="type"),"data-rocket-src"===e&&(e="src"),r.setAttribute(e,t.nodeValue))}),t.text&&(r.text=t.text),r.hasAttribute("src")?(r.addEventListener("load",e),r.addEventListener("error",e)):(r.text=t.text,e());try{t.parentNode.replaceChild(r,t)}catch(i){e()}}:async e=>{function r(){t.setAttribute("data-rocket-status","failed"),e()}try{let i=t.getAttribute("data-rocket-type"),n=t.getAttribute("data-rocket-src");t.text,i?(t.type=i,t.removeAttribute("data-rocket-type")):t.removeAttribute("type"),t.addEventListener("load",function r(){t.setAttribute("data-rocket-status","executed"),e()}),t.addEventListener("error",r),n?(t.removeAttribute("data-rocket-src"),t.src=n):t.src="data:text/javascript;base64,"+window.btoa(unescape(encodeURIComponent(t.text)))}catch(s){r()}})}async _loadScriptsFromList(t){let e=t.shift();return e&&e.isConnected?(await this._transformScript(e),this._loadScriptsFromList(t)):Promise.resolve()}_preloadAllScripts(){this._batchInjectResourceHints([...this.delayedScripts.normal,...this.delayedScripts.defer,...this.delayedScripts.async],"preload")}_batchInjectResourceHints(t,e){var r=document.createDocumentFragment();t.forEach(t=>{let i=t.getAttribute&&t.getAttribute("data-rocket-src")||t.src;if(i){let n=document.createElement("link");n.href=i,n.rel=e,"preconnect"!==e&&(n.as="script"),t.getAttribute&&"module"===t.getAttribute("data-rocket-type")&&(n.crossOrigin=!0),t.crossOrigin&&(n.crossOrigin=t.crossOrigin),t.integrity&&(n.integrity=t.integrity),r.appendChild(n),this.trash.push(n)}}),document.head.appendChild(r)}_delayEventListeners(t){let e={};function r(t,r){!function t(r){!e[r]&&(e[r]={originalFunctions:{add:r.addEventListener,remove:r.removeEventListener},eventsToRewrite:[]},r.addEventListener=function(){arguments[0]=i(arguments[0]),e[r].originalFunctions.add.apply(r,arguments)},r.removeEventListener=function(){arguments[0]=i(arguments[0]),e[r].originalFunctions.remove.apply(r,arguments)});function i(t){return e[r].eventsToRewrite.indexOf(t)>=0?"rocket-"+t:t}}(t),e[t].eventsToRewrite.push(r)}function i(t,e){let r=t[e];Object.defineProperty(t,e,{get:()=>r||function(){},set(i){t["rocket"+e]=r=i}})}r(document,"DOMContentLoaded"),r(window,"DOMContentLoaded"),r(window,"load"),r(window,"pageshow"),r(document,"readystatechange"),i(document,"onreadystatechange"),i(window,"onload"),i(window,"onpageshow")}_delayJQueryReady(t){let e;function r(r){if(r&&r.fn&&!t.allJQueries.includes(r)){r.fn.ready=r.fn.init.prototype.ready=function(e){return t.domReadyFired?e.bind(document)(r):document.addEventListener("rocket-DOMContentLoaded",()=>e.bind(document)(r)),r([])};let i=r.fn.on;r.fn.on=r.fn.init.prototype.on=function(){if(this[0]===window){function t(t){return t.split(" ").map(t=>"load"===t||0===t.indexOf("load.")?"rocket-jquery-load":t).join(" ")}"string"==typeof arguments[0]||arguments[0]instanceof String?arguments[0]=t(arguments[0]):"object"==typeof arguments[0]&&Object.keys(arguments[0]).forEach(e=>{let r=arguments[0][e];delete arguments[0][e],arguments[0][t(e)]=r})}return i.apply(this,arguments),this},t.allJQueries.push(r)}e=r}r(window.jQuery),Object.defineProperty(window,"jQuery",{get:()=>e,set(t){r(t)}})}async _triggerDOMContentLoaded(){this.domReadyFired=!0,await this._littleBreath(),document.dispatchEvent(new Event("rocket-DOMContentLoaded")),await this._littleBreath(),window.dispatchEvent(new Event("rocket-DOMContentLoaded")),await this._littleBreath(),document.dispatchEvent(new Event("rocket-readystatechange")),await this._littleBreath(),document.rocketonreadystatechange&&document.rocketonreadystatechange()}async _triggerWindowLoad(){await this._littleBreath(),window.dispatchEvent(new Event("rocket-load")),await this._littleBreath(),window.rocketonload&&window.rocketonload(),await this._littleBreath(),this.allJQueries.forEach(t=>t(window).trigger("rocket-jquery-load")),await this._littleBreath();let t=new Event("rocket-pageshow");t.persisted=this.persisted,window.dispatchEvent(t),await this._littleBreath(),window.rocketonpageshow&&window.rocketonpageshow({persisted:this.persisted})}_handleDocumentWrite(){let t=new Map;document.write=document.writeln=function(e){let r=document.currentScript;r||console.error("WPRocket unable to document.write this: "+e);let i=document.createRange(),n=r.parentElement,s=t.get(r);void 0===s&&(s=r.nextSibling,t.set(r,s));let a=document.createDocumentFragment();i.setStart(a,0),a.appendChild(i.createContextualFragment(e)),n.insertBefore(a,s)}}async _littleBreath(){Date.now()-this.lastBreath>45&&(await this._requestAnimFrame(),this.lastBreath=Date.now())}async _requestAnimFrame(){return document.hidden?new Promise(t=>setTimeout(t)):new Promise(t=>requestAnimationFrame(t))}_emptyTrash(){this.trash.forEach(t=>t.remove())}static run(){let t=new RocketLazyLoadScripts;t._addUserInteractionListener(t)}}RocketLazyLoadScripts.run();</script>
	
	<meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1">
	<link rel="alternate" hreflang="en" href="https://www.guru99.com/software-testing-introduction-importance.html" />
<link rel="alternate" hreflang="fr" href="https://www.guru99.com/fr/software-testing-introduction-importance.html" />
<title>What is Software Testing? Definition</title>
<link rel="preload" as="font" href="https://www.guru99.com/wp-content/fonts/source-sans-pro/6xK3dSBYKcSV-LCoeQqfX1RYOo3qOK7l.woff2" crossorigin><link rel="stylesheet" href="https://www.guru99.com/wp-content/cache/min/1/29af46a51ed0a43a870973597a1e1216.css" media="all" data-minify="1" />
<meta name="description" content="Testing in Software Engineering is defined as an activity to check whether the actual results match the expected results and to ensure that the software system is Defect free."/>
<meta name="robots" content="follow, index, max-image-preview:large"/>
<link rel="canonical" href="https://www.guru99.com/software-testing-introduction-importance.html" />
<meta name="keywords" content="testing"/>
<meta property="og:locale" content="en_US" />
<meta property="og:type" content="article" />
<meta property="og:title" content="What is Software Testing? Definition" />
<meta property="og:description" content="Testing in Software Engineering is defined as an activity to check whether the actual results match the expected results and to ensure that the software system is Defect free." />
<meta property="og:url" content="https://www.guru99.com/software-testing-introduction-importance.html" />
<meta property="og:site_name" content="Guru99" />
<meta property="article:publisher" content="https://www.facebook.com/Guru99Official" />
<meta property="og:updated_time" content="2023-09-19T17:07:31+05:30" />
<meta property="og:image" content="https://www.guru99.com/images/2/061920_1310_Whatissoftwaretesting1.png" />
<meta property="og:image:secure_url" content="https://www.guru99.com/images/2/061920_1310_Whatissoftwaretesting1.png" />
<meta property="og:image:alt" content="testing" />
<meta property="og:video" content="https://www.youtube.com/embed/TDynSmrzpXw?si=PpkvKDrY5bOKk3J6" />
<meta property="ya:ovs:adult" content="1" />
<meta property="ya:ovs:upload_date" content="2020-01-02GMT+053000:00:00+05:30" />
<meta property="ya:ovs:allow_embed" content="true" />
<meta name="twitter:card" content="summary_large_image" />
<meta name="twitter:title" content="What is Software Testing? Definition" />
<meta name="twitter:description" content="Testing in Software Engineering is defined as an activity to check whether the actual results match the expected results and to ensure that the software system is Defect free." />
<meta name="twitter:site" content="@guru99com" />
<meta name="twitter:creator" content="@guru99com" />
<meta name="twitter:image" content="https://www.guru99.com/images/2/061920_1310_Whatissoftwaretesting1.png" />
<meta name="twitter:label1" content="Written by" />
<meta name="twitter:data1" content="Thomas Hamilton" />
<meta name="twitter:label2" content="Time to read" />
<meta name="twitter:data2" content="5 minutes" />
<script type="application/ld+json" class="rank-math-schema-pro">{"@context":"https://schema.org","@graph":[{"@type":"Organization","@id":"https://www.guru99.com/#organization","name":"Guru99","sameAs":["https://www.facebook.com/Guru99Official","https://twitter.com/guru99com"],"logo":{"@type":"ImageObject","@id":"https://www.guru99.com/#logo","url":"https://www.guru99.com/images/guru99-logo-v1-150x59.png","contentUrl":"https://www.guru99.com/images/guru99-logo-v1-150x59.png","caption":"Guru99","inLanguage":"en-US"}},{"@type":"WebSite","@id":"https://www.guru99.com/#website","url":"https://www.guru99.com","name":"Guru99","publisher":{"@id":"https://www.guru99.com/#organization"},"inLanguage":"en-US"},{"@type":"ImageObject","@id":"https://www.guru99.com/images/software-testing-introduction.png","url":"https://www.guru99.com/images/software-testing-introduction.png","width":"368","height":"150","inLanguage":"en-US"},{"@type":"BreadcrumbList","@id":"https://www.guru99.com/software-testing-introduction-importance.html#breadcrumb","itemListElement":[{"@type":"ListItem","position":"1","item":{"@id":"https://www.guru99.com","name":"Home"}},{"@type":"ListItem","position":"2","item":{"@id":"https://www.guru99.com/softwaretesting","name":"Software Testing"}},{"@type":"ListItem","position":"3","item":{"@id":"https://www.guru99.com/software-testing-introduction-importance.html","name":"What is Software Testing? Definition"}}]},{"@type":"WebPage","@id":"https://www.guru99.com/software-testing-introduction-importance.html#webpage","url":"https://www.guru99.com/software-testing-introduction-importance.html","name":"What is Software Testing? Definition","dateModified":"2023-09-19T17:07:31+05:30","isPartOf":{"@id":"https://www.guru99.com/#website"},"primaryImageOfPage":{"@id":"https://www.guru99.com/images/software-testing-introduction.png"},"inLanguage":"en-US","breadcrumb":{"@id":"https://www.guru99.com/software-testing-introduction-importance.html#breadcrumb"}},{"@type":"Person","@id":"https://www.guru99.com/author/thomas","name":"Thomas Hamilton","url":"https://www.guru99.com/author/thomas","image":{"@type":"ImageObject","@id":"https://www.guru99.com/images/thomas-hamilton.jpg","url":"https://www.guru99.com/images/thomas-hamilton.jpg","caption":"Thomas Hamilton","inLanguage":"en-US"},"worksFor":{"@id":"https://www.guru99.com/#organization"}},{"@type":"Article","headline":"What is Software Testing? Definition","description":"Testing in Software Engineering is defined as an activity to check whether the actual results match the expected results and to ensure that the software system is Defect free.","author":{"@id":"https://www.guru99.com/author/thomas","name":"Thomas Hamilton"},"name":"What is Software Testing? Definition","articleSection":"Software Testing","@id":"https://www.guru99.com/software-testing-introduction-importance.html#schema-567486","isPartOf":{"@id":"https://www.guru99.com/software-testing-introduction-importance.html#webpage"},"publisher":{"@id":"https://www.guru99.com/#organization"},"image":{"@id":"https://www.guru99.com/images/software-testing-introduction.png"},"inLanguage":"en-US","mainEntityOfPage":{"@id":"https://www.guru99.com/software-testing-introduction-importance.html#webpage"}},{"@type":"VideoObject","embedUrl":"https://www.youtube.com/embed/TDynSmrzpXw?si=PpkvKDrY5bOKk3J6","name":"What is Software Testing? Definition","description":"Testing in Software Engineering is defined as an activity to check whether the actual results match the expected results and to ensure that the software system is Defect free.","uploadDate":"2020-01-02T00:00:00+05:30","thumbnailUrl":"https://www.guru99.com/images/2/061920_1310_Whatissoftwaretesting1.png","hasPart":[],"width":"640","height":"360","@id":"https://www.guru99.com/software-testing-introduction-importance.html#schema-567488","isPartOf":{"@id":"https://www.guru99.com/software-testing-introduction-importance.html#webpage"},"publisher":{"@id":"https://www.guru99.com/#organization"},"inLanguage":"en-US"}]}</script>
<link rel='dns-prefetch' href='//pagead2.googlesyndication.com' />
<link rel='dns-prefetch' href='//acdn.adnxs.com' />

			<script type="rocketlazyloadscript">document.documentElement.classList.remove( 'no-js' );</script>
			<style>
img.wp-smiley,
img.emoji {
	display: inline !important;
	border: none !important;
	box-shadow: none !important;
	height: 1em !important;
	width: 1em !important;
	margin: 0 0.07em !important;
	vertical-align: -0.1em !important;
	background: none !important;
	padding: 0 !important;
}
</style>
	


<style id='classic-theme-styles-inline-css'>
/*! This file is auto-generated */
.wp-block-button__link{color:#fff;background-color:#32373c;border-radius:9999px;box-shadow:none;text-decoration:none;padding:calc(.667em + 2px) calc(1.333em + 2px);font-size:1.125em}.wp-block-file__button{background:#32373c;color:#fff;text-decoration:none}
</style>
<style id='global-styles-inline-css'>
body{--wp--preset--color--black: #000000;--wp--preset--color--cyan-bluish-gray: #abb8c3;--wp--preset--color--white: #ffffff;--wp--preset--color--pale-pink: #f78da7;--wp--preset--color--vivid-red: #cf2e2e;--wp--preset--color--luminous-vivid-orange: #ff6900;--wp--preset--color--luminous-vivid-amber: #fcb900;--wp--preset--color--light-green-cyan: #7bdcb5;--wp--preset--color--vivid-green-cyan: #00d084;--wp--preset--color--pale-cyan-blue: #8ed1fc;--wp--preset--color--vivid-cyan-blue: #0693e3;--wp--preset--color--vivid-purple: #9b51e0;--wp--preset--color--theme-palette-1: #3182CE;--wp--preset--color--theme-palette-2: #2B6CB0;--wp--preset--color--theme-palette-3: #1A202C;--wp--preset--color--theme-palette-4: #2D3748;--wp--preset--color--theme-palette-5: #4A5568;--wp--preset--color--theme-palette-6: #718096;--wp--preset--color--theme-palette-7: #EDF2F7;--wp--preset--color--theme-palette-8: #F7FAFC;--wp--preset--color--theme-palette-9: #FFFFFF;--wp--preset--gradient--vivid-cyan-blue-to-vivid-purple: linear-gradient(135deg,rgba(6,147,227,1) 0%,rgb(155,81,224) 100%);--wp--preset--gradient--light-green-cyan-to-vivid-green-cyan: linear-gradient(135deg,rgb(122,220,180) 0%,rgb(0,208,130) 100%);--wp--preset--gradient--luminous-vivid-amber-to-luminous-vivid-orange: linear-gradient(135deg,rgba(252,185,0,1) 0%,rgba(255,105,0,1) 100%);--wp--preset--gradient--luminous-vivid-orange-to-vivid-red: linear-gradient(135deg,rgba(255,105,0,1) 0%,rgb(207,46,46) 100%);--wp--preset--gradient--very-light-gray-to-cyan-bluish-gray: linear-gradient(135deg,rgb(238,238,238) 0%,rgb(169,184,195) 100%);--wp--preset--gradient--cool-to-warm-spectrum: linear-gradient(135deg,rgb(74,234,220) 0%,rgb(151,120,209) 20%,rgb(207,42,186) 40%,rgb(238,44,130) 60%,rgb(251,105,98) 80%,rgb(254,248,76) 100%);--wp--preset--gradient--blush-light-purple: linear-gradient(135deg,rgb(255,206,236) 0%,rgb(152,150,240) 100%);--wp--preset--gradient--blush-bordeaux: linear-gradient(135deg,rgb(254,205,165) 0%,rgb(254,45,45) 50%,rgb(107,0,62) 100%);--wp--preset--gradient--luminous-dusk: linear-gradient(135deg,rgb(255,203,112) 0%,rgb(199,81,192) 50%,rgb(65,88,208) 100%);--wp--preset--gradient--pale-ocean: linear-gradient(135deg,rgb(255,245,203) 0%,rgb(182,227,212) 50%,rgb(51,167,181) 100%);--wp--preset--gradient--electric-grass: linear-gradient(135deg,rgb(202,248,128) 0%,rgb(113,206,126) 100%);--wp--preset--gradient--midnight: linear-gradient(135deg,rgb(2,3,129) 0%,rgb(40,116,252) 100%);--wp--preset--font-size--small: 14px;--wp--preset--font-size--medium: 24px;--wp--preset--font-size--large: 32px;--wp--preset--font-size--x-large: 42px;--wp--preset--font-size--larger: 40px;--wp--preset--spacing--20: 0.44rem;--wp--preset--spacing--30: 0.67rem;--wp--preset--spacing--40: 1rem;--wp--preset--spacing--50: 1.5rem;--wp--preset--spacing--60: 2.25rem;--wp--preset--spacing--70: 3.38rem;--wp--preset--spacing--80: 5.06rem;--wp--preset--shadow--natural: 6px 6px 9px rgba(0, 0, 0, 0.2);--wp--preset--shadow--deep: 12px 12px 50px rgba(0, 0, 0, 0.4);--wp--preset--shadow--sharp: 6px 6px 0px rgba(0, 0, 0, 0.2);--wp--preset--shadow--outlined: 6px 6px 0px -3px rgba(255, 255, 255, 1), 6px 6px rgba(0, 0, 0, 1);--wp--preset--shadow--crisp: 6px 6px 0px rgba(0, 0, 0, 1);}:where(.is-layout-flex){gap: 0.5em;}:where(.is-layout-grid){gap: 0.5em;}body .is-layout-flow > .alignleft{float: left;margin-inline-start: 0;margin-inline-end: 2em;}body .is-layout-flow > .alignright{float: right;margin-inline-start: 2em;margin-inline-end: 0;}body .is-layout-flow > .aligncenter{margin-left: auto !important;margin-right: auto !important;}body .is-layout-constrained > .alignleft{float: left;margin-inline-start: 0;margin-inline-end: 2em;}body .is-layout-constrained > .alignright{float: right;margin-inline-start: 2em;margin-inline-end: 0;}body .is-layout-constrained > .aligncenter{margin-left: auto !important;margin-right: auto !important;}body .is-layout-constrained > :where(:not(.alignleft):not(.alignright):not(.alignfull)){max-width: var(--wp--style--global--content-size);margin-left: auto !important;margin-right: auto !important;}body .is-layout-constrained > .alignwide{max-width: var(--wp--style--global--wide-size);}body .is-layout-flex{display: flex;}body .is-layout-flex{flex-wrap: wrap;align-items: center;}body .is-layout-flex > *{margin: 0;}body .is-layout-grid{display: grid;}body .is-layout-grid > *{margin: 0;}:where(.wp-block-columns.is-layout-flex){gap: 2em;}:where(.wp-block-columns.is-layout-grid){gap: 2em;}:where(.wp-block-post-template.is-layout-flex){gap: 1.25em;}:where(.wp-block-post-template.is-layout-grid){gap: 1.25em;}.has-black-color{color: var(--wp--preset--color--black) !important;}.has-cyan-bluish-gray-color{color: var(--wp--preset--color--cyan-bluish-gray) !important;}.has-white-color{color: var(--wp--preset--color--white) !important;}.has-pale-pink-color{color: var(--wp--preset--color--pale-pink) !important;}.has-vivid-red-color{color: var(--wp--preset--color--vivid-red) !important;}.has-luminous-vivid-orange-color{color: var(--wp--preset--color--luminous-vivid-orange) !important;}.has-luminous-vivid-amber-color{color: var(--wp--preset--color--luminous-vivid-amber) !important;}.has-light-green-cyan-color{color: var(--wp--preset--color--light-green-cyan) !important;}.has-vivid-green-cyan-color{color: var(--wp--preset--color--vivid-green-cyan) !important;}.has-pale-cyan-blue-color{color: var(--wp--preset--color--pale-cyan-blue) !important;}.has-vivid-cyan-blue-color{color: var(--wp--preset--color--vivid-cyan-blue) !important;}.has-vivid-purple-color{color: var(--wp--preset--color--vivid-purple) !important;}.has-black-background-color{background-color: var(--wp--preset--color--black) !important;}.has-cyan-bluish-gray-background-color{background-color: var(--wp--preset--color--cyan-bluish-gray) !important;}.has-white-background-color{background-color: var(--wp--preset--color--white) !important;}.has-pale-pink-background-color{background-color: var(--wp--preset--color--pale-pink) !important;}.has-vivid-red-background-color{background-color: var(--wp--preset--color--vivid-red) !important;}.has-luminous-vivid-orange-background-color{background-color: var(--wp--preset--color--luminous-vivid-orange) !important;}.has-luminous-vivid-amber-background-color{background-color: var(--wp--preset--color--luminous-vivid-amber) !important;}.has-light-green-cyan-background-color{background-color: var(--wp--preset--color--light-green-cyan) !important;}.has-vivid-green-cyan-background-color{background-color: var(--wp--preset--color--vivid-green-cyan) !important;}.has-pale-cyan-blue-background-color{background-color: var(--wp--preset--color--pale-cyan-blue) !important;}.has-vivid-cyan-blue-background-color{background-color: var(--wp--preset--color--vivid-cyan-blue) !important;}.has-vivid-purple-background-color{background-color: var(--wp--preset--color--vivid-purple) !important;}.has-black-border-color{border-color: var(--wp--preset--color--black) !important;}.has-cyan-bluish-gray-border-color{border-color: var(--wp--preset--color--cyan-bluish-gray) !important;}.has-white-border-color{border-color: var(--wp--preset--color--white) !important;}.has-pale-pink-border-color{border-color: var(--wp--preset--color--pale-pink) !important;}.has-vivid-red-border-color{border-color: var(--wp--preset--color--vivid-red) !important;}.has-luminous-vivid-orange-border-color{border-color: var(--wp--preset--color--luminous-vivid-orange) !important;}.has-luminous-vivid-amber-border-color{border-color: var(--wp--preset--color--luminous-vivid-amber) !important;}.has-light-green-cyan-border-color{border-color: var(--wp--preset--color--light-green-cyan) !important;}.has-vivid-green-cyan-border-color{border-color: var(--wp--preset--color--vivid-green-cyan) !important;}.has-pale-cyan-blue-border-color{border-color: var(--wp--preset--color--pale-cyan-blue) !important;}.has-vivid-cyan-blue-border-color{border-color: var(--wp--preset--color--vivid-cyan-blue) !important;}.has-vivid-purple-border-color{border-color: var(--wp--preset--color--vivid-purple) !important;}.has-vivid-cyan-blue-to-vivid-purple-gradient-background{background: var(--wp--preset--gradient--vivid-cyan-blue-to-vivid-purple) !important;}.has-light-green-cyan-to-vivid-green-cyan-gradient-background{background: var(--wp--preset--gradient--light-green-cyan-to-vivid-green-cyan) !important;}.has-luminous-vivid-amber-to-luminous-vivid-orange-gradient-background{background: var(--wp--preset--gradient--luminous-vivid-amber-to-luminous-vivid-orange) !important;}.has-luminous-vivid-orange-to-vivid-red-gradient-background{background: var(--wp--preset--gradient--luminous-vivid-orange-to-vivid-red) !important;}.has-very-light-gray-to-cyan-bluish-gray-gradient-background{background: var(--wp--preset--gradient--very-light-gray-to-cyan-bluish-gray) !important;}.has-cool-to-warm-spectrum-gradient-background{background: var(--wp--preset--gradient--cool-to-warm-spectrum) !important;}.has-blush-light-purple-gradient-background{background: var(--wp--preset--gradient--blush-light-purple) !important;}.has-blush-bordeaux-gradient-background{background: var(--wp--preset--gradient--blush-bordeaux) !important;}.has-luminous-dusk-gradient-background{background: var(--wp--preset--gradient--luminous-dusk) !important;}.has-pale-ocean-gradient-background{background: var(--wp--preset--gradient--pale-ocean) !important;}.has-electric-grass-gradient-background{background: var(--wp--preset--gradient--electric-grass) !important;}.has-midnight-gradient-background{background: var(--wp--preset--gradient--midnight) !important;}.has-small-font-size{font-size: var(--wp--preset--font-size--small) !important;}.has-medium-font-size{font-size: var(--wp--preset--font-size--medium) !important;}.has-large-font-size{font-size: var(--wp--preset--font-size--large) !important;}.has-x-large-font-size{font-size: var(--wp--preset--font-size--x-large) !important;}
.wp-block-navigation a:where(:not(.wp-element-button)){color: inherit;}
:where(.wp-block-post-template.is-layout-flex){gap: 1.25em;}:where(.wp-block-post-template.is-layout-grid){gap: 1.25em;}
:where(.wp-block-columns.is-layout-flex){gap: 2em;}:where(.wp-block-columns.is-layout-grid){gap: 2em;}
.wp-block-pullquote{font-size: 1.5em;line-height: 1.6;}
</style>



<style id='kadence-global-inline-css'>
/* Kadence Base CSS */
:root{--global-palette1:#3182CE;--global-palette2:#2B6CB0;--global-palette3:#1A202C;--global-palette4:#2D3748;--global-palette5:#4A5568;--global-palette6:#718096;--global-palette7:#EDF2F7;--global-palette8:#F7FAFC;--global-palette9:#FFFFFF;--global-palette9rgb:255, 255, 255;--global-palette-highlight:#0556f3;--global-palette-highlight-alt:#0556f3;--global-palette-highlight-alt2:var(--global-palette9);--global-palette-btn-bg:var(--global-palette1);--global-palette-btn-bg-hover:var(--global-palette1);--global-palette-btn:var(--global-palette9);--global-palette-btn-hover:var(--global-palette9);--global-body-font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen-Sans,Ubuntu,Cantarell,"Helvetica Neue",sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";--global-heading-font-family:'Source Sans Pro', sans-serif;--global-primary-nav-font-family:inherit;--global-fallback-font:sans-serif;--global-display-fallback-font:sans-serif;--global-content-width:1290px;--global-content-narrow-width:842px;--global-content-edge-padding:1.5rem;--global-content-boxed-padding:2rem;--global-calc-content-width:calc(1290px - var(--global-content-edge-padding) - var(--global-content-edge-padding) );--wp--style--global--content-size:var(--global-calc-content-width);}.wp-site-blocks{--global-vw:calc( 100vw - ( 0.5 * var(--scrollbar-offset)));}:root .has-theme-palette-1-background-color{background-color:var(--global-palette1);}:root .has-theme-palette-1-color{color:var(--global-palette1);}:root .has-theme-palette-2-background-color{background-color:var(--global-palette2);}:root .has-theme-palette-2-color{color:var(--global-palette2);}:root .has-theme-palette-3-background-color{background-color:var(--global-palette3);}:root .has-theme-palette-3-color{color:var(--global-palette3);}:root .has-theme-palette-4-background-color{background-color:var(--global-palette4);}:root .has-theme-palette-4-color{color:var(--global-palette4);}:root .has-theme-palette-5-background-color{background-color:var(--global-palette5);}:root .has-theme-palette-5-color{color:var(--global-palette5);}:root .has-theme-palette-6-background-color{background-color:var(--global-palette6);}:root .has-theme-palette-6-color{color:var(--global-palette6);}:root .has-theme-palette-7-background-color{background-color:var(--global-palette7);}:root .has-theme-palette-7-color{color:var(--global-palette7);}:root .has-theme-palette-8-background-color{background-color:var(--global-palette8);}:root .has-theme-palette-8-color{color:var(--global-palette8);}:root .has-theme-palette-9-background-color{background-color:var(--global-palette9);}:root .has-theme-palette-9-color{color:var(--global-palette9);}:root .has-theme-palette1-background-color{background-color:var(--global-palette1);}:root .has-theme-palette1-color{color:var(--global-palette1);}:root .has-theme-palette2-background-color{background-color:var(--global-palette2);}:root .has-theme-palette2-color{color:var(--global-palette2);}:root .has-theme-palette3-background-color{background-color:var(--global-palette3);}:root .has-theme-palette3-color{color:var(--global-palette3);}:root .has-theme-palette4-background-color{background-color:var(--global-palette4);}:root .has-theme-palette4-color{color:var(--global-palette4);}:root .has-theme-palette5-background-color{background-color:var(--global-palette5);}:root .has-theme-palette5-color{color:var(--global-palette5);}:root .has-theme-palette6-background-color{background-color:var(--global-palette6);}:root .has-theme-palette6-color{color:var(--global-palette6);}:root .has-theme-palette7-background-color{background-color:var(--global-palette7);}:root .has-theme-palette7-color{color:var(--global-palette7);}:root .has-theme-palette8-background-color{background-color:var(--global-palette8);}:root .has-theme-palette8-color{color:var(--global-palette8);}:root .has-theme-palette9-background-color{background-color:var(--global-palette9);}:root .has-theme-palette9-color{color:var(--global-palette9);}body{background:var(--global-palette9);}body, input, select, optgroup, textarea{font-style:normal;font-weight:400;font-size:20px;line-height:27px;font-family:var(--global-body-font-family);color:#222222;}.content-bg, body.content-style-unboxed .site{background:var(--global-palette9);}h1,h2,h3,h4,h5,h6{font-family:var(--global-heading-font-family);}h1{font-style:normal;font-weight:normal;font-size:33px;line-height:34px;font-family:'Source Sans Pro', sans-serif;color:#222222;}h2{font-style:normal;font-weight:normal;font-size:28px;line-height:40px;font-family:'Source Sans Pro', sans-serif;color:#222222;}h3{font-style:normal;font-weight:normal;font-size:24px;line-height:25px;font-family:'Source Sans Pro', sans-serif;color:#222222;}h4{font-style:normal;font-weight:normal;font-size:22px;line-height:21px;font-family:'Source Sans Pro', sans-serif;color:#222222;}h5{font-style:normal;font-weight:normal;font-size:19px;line-height:20px;font-family:'Source Sans Pro', sans-serif;color:#222222;}h6{font-style:normal;font-weight:normal;font-size:18px;line-height:1.5;font-family:'Source Sans Pro', sans-serif;color:#222222;}.entry-hero h1{font-style:normal;font-weight:normal;font-size:33px;line-height:34px;font-family:'Source Sans Pro', sans-serif;color:#222222;}.entry-hero .kadence-breadcrumbs, .entry-hero .search-form{font-style:normal;}.entry-hero .kadence-breadcrumbs{max-width:1290px;}.site-container, .site-header-row-layout-contained, .site-footer-row-layout-contained, .entry-hero-layout-contained, .comments-area, .alignfull > .wp-block-cover__inner-container, .alignwide > .wp-block-cover__inner-container{max-width:var(--global-content-width);}.content-width-narrow .content-container.site-container, .content-width-narrow .hero-container.site-container{max-width:var(--global-content-narrow-width);}@media all and (min-width: 1520px){.wp-site-blocks .content-container  .alignwide{margin-left:-115px;margin-right:-115px;width:unset;max-width:unset;}}@media all and (min-width: 1102px){.content-width-narrow .wp-site-blocks .content-container .alignwide{margin-left:-130px;margin-right:-130px;width:unset;max-width:unset;}}.content-style-boxed .wp-site-blocks .entry-content .alignwide{margin-left:calc( -1 * var( --global-content-boxed-padding ) );margin-right:calc( -1 * var( --global-content-boxed-padding ) );}.content-area{margin-top:5rem;margin-bottom:5rem;}@media all and (max-width: 1024px){.content-area{margin-top:3rem;margin-bottom:3rem;}}@media all and (max-width: 767px){.content-area{margin-top:2rem;margin-bottom:2rem;}}@media all and (max-width: 1024px){:root{--global-content-boxed-padding:2rem;}}@media all and (max-width: 767px){:root{--global-content-boxed-padding:1.5rem;}}.entry-content-wrap{padding:2rem;}@media all and (max-width: 1024px){.entry-content-wrap{padding:2rem;}}@media all and (max-width: 767px){.entry-content-wrap{padding:1.5rem;}}.entry.single-entry{box-shadow:0px 15px 15px -10px rgba(0,0,0,0.05);}.entry.loop-entry{box-shadow:0px 15px 15px -10px rgba(0,0,0,0.05);}.loop-entry .entry-content-wrap{padding:2rem;}@media all and (max-width: 1024px){.loop-entry .entry-content-wrap{padding:2rem;}}@media all and (max-width: 767px){.loop-entry .entry-content-wrap{padding:1.5rem;}}.primary-sidebar.widget-area .widget{margin-bottom:1.5em;color:var(--global-palette4);}.primary-sidebar.widget-area .widget-title{font-style:normal;font-weight:normal;font-size:20px;line-height:1.5;color:var(--global-palette3);}.primary-sidebar.widget-area .sidebar-inner-wrap a:where(:not(.button):not(.wp-block-button__link):not(.wp-element-button)):hover{color:#ec4747;}.primary-sidebar.widget-area{background:var(--global-palette9);}.has-sidebar.has-left-sidebar:not(.rtl) .primary-sidebar.widget-area, .rtl.has-sidebar:not(.has-left-sidebar) .primary-sidebar.widget-area{border-right:1px solid #e1e1e1;}.has-sidebar:not(.has-left-sidebar):not(.rtl) .primary-sidebar.widget-area, .rtl.has-sidebar.has-left-sidebar .primary-sidebar.widget-area{border-left:1px solid #e1e1e1;}button, .button, .wp-block-button__link, input[type="button"], input[type="reset"], input[type="submit"], .fl-button, .elementor-button-wrapper .elementor-button{box-shadow:0px 0px 0px -7px rgba(0,0,0,0);}button:hover, button:focus, button:active, .button:hover, .button:focus, .button:active, .wp-block-button__link:hover, .wp-block-button__link:focus, .wp-block-button__link:active, input[type="button"]:hover, input[type="button"]:focus, input[type="button"]:active, input[type="reset"]:hover, input[type="reset"]:focus, input[type="reset"]:active, input[type="submit"]:hover, input[type="submit"]:focus, input[type="submit"]:active, .elementor-button-wrapper .elementor-button:hover, .elementor-button-wrapper .elementor-button:focus, .elementor-button-wrapper .elementor-button:active{box-shadow:0px 15px 25px -7px rgba(0,0,0,0.1);}@media all and (min-width: 1025px){.transparent-header .entry-hero .entry-hero-container-inner{padding-top:49px;}}@media all and (max-width: 1024px){.mobile-transparent-header .entry-hero .entry-hero-container-inner{padding-top:49px;}}@media all and (max-width: 767px){.mobile-transparent-header .entry-hero .entry-hero-container-inner{padding-top:49px;}}.wp-site-blocks .entry-hero-container-inner{background:var(--global-palette9);}#colophon{background:#323a56;}.site-middle-footer-wrap .site-footer-row-container-inner{background:#323a56;font-style:normal;}.site-footer .site-middle-footer-wrap a:where(:not(.button):not(.wp-block-button__link):not(.wp-element-button)){color:var(--global-palette1);}.site-footer .site-middle-footer-wrap a:where(:not(.button):not(.wp-block-button__link):not(.wp-element-button)):hover{color:var(--global-palette1);}.site-middle-footer-inner-wrap{padding-top:0px;padding-bottom:30px;grid-column-gap:0px;grid-row-gap:0px;}.site-middle-footer-inner-wrap .widget{margin-bottom:30px;}.site-middle-footer-inner-wrap .widget-area .widget-title{font-style:normal;font-weight:400;}.site-middle-footer-inner-wrap .site-footer-section:not(:last-child):after{right:calc(-0px / 2);}.site-top-footer-wrap .site-footer-row-container-inner{background:#323a56;font-style:normal;color:var(--global-palette4);border-bottom:0px none transparent;}.site-footer .site-top-footer-wrap a:not(.button):not(.wp-block-button__link):not(.wp-element-button){color:var(--global-palette1);}.site-top-footer-inner-wrap{padding-top:0px;padding-bottom:0px;grid-column-gap:0px;grid-row-gap:0px;}.site-top-footer-inner-wrap .widget{margin-bottom:30px;}.site-top-footer-inner-wrap .site-footer-section:not(:last-child):after{border-right:0px none transparent;right:calc(-0px / 2);}@media all and (max-width: 767px){.site-top-footer-wrap .site-footer-row-container-inner{border-bottom:1px none #323a56;}.site-top-footer-inner-wrap .site-footer-section:not(:last-child):after{border-right:0px none transparent;}}.site-bottom-footer-wrap .site-footer-row-container-inner{background:var(--global-palette9);}.site-bottom-footer-inner-wrap{padding-top:30px;padding-bottom:30px;grid-column-gap:30px;}.site-bottom-footer-inner-wrap .widget{margin-bottom:30px;}.site-bottom-footer-inner-wrap .site-footer-section:not(:last-child):after{right:calc(-30px / 2);}.footer-social-wrap{margin:0px 0px 0px 0px;}.footer-social-wrap .footer-social-inner-wrap{font-size:1.28em;gap:0.3em;}.site-footer .site-footer-wrap .site-footer-section .footer-social-wrap .footer-social-inner-wrap .social-button{color:var(--global-palette9);border:2px none transparent;border-color:var(--global-palette9);border-radius:3px;}.site-footer .site-footer-wrap .site-footer-section .footer-social-wrap .footer-social-inner-wrap .social-button:hover{color:var(--global-palette9);border-color:var(--global-palette9);}#colophon .footer-html{font-style:normal;color:var(--global-palette9);}#colophon .site-footer-row-container .site-footer-row .footer-html a{color:var(--global-palette9);}#kt-scroll-up-reader, #kt-scroll-up{border-radius:0px 0px 0px 0px;color:var(--global-palette3);border-color:var(--global-palette4);bottom:30px;font-size:1.2em;padding:0.4em 0.4em 0.4em 0.4em;}#kt-scroll-up-reader.scroll-up-side-right, #kt-scroll-up.scroll-up-side-right{right:30px;}#kt-scroll-up-reader.scroll-up-side-left, #kt-scroll-up.scroll-up-side-left{left:30px;}#kt-scroll-up-reader:hover, #kt-scroll-up:hover{color:var(--global-palette2);border-color:var(--global-palette2);}#colophon .footer-navigation .footer-menu-container > ul > li > a{padding-left:calc(1.2em / 2);padding-right:calc(1.2em / 2);padding-top:calc(0.6em / 2);padding-bottom:calc(0.6em / 2);color:var(--global-palette5);}#colophon .footer-navigation .footer-menu-container > ul li a:hover{color:var(--global-palette-highlight);}#colophon .footer-navigation .footer-menu-container > ul li.current-menu-item > a{color:var(--global-palette3);}body.page{background:var(--global-palette9);}.entry-hero.page-hero-section .entry-header{min-height:200px;}.comment-metadata a:not(.comment-edit-link), .comment-body .edit-link:before{display:none;}.entry-hero.post-hero-section .entry-header{min-height:200px;}
/* Kadence Header CSS */
@media all and (max-width: 1024px){.mobile-transparent-header #masthead{position:absolute;left:0px;right:0px;z-index:100;}.kadence-scrollbar-fixer.mobile-transparent-header #masthead{right:var(--scrollbar-offset,0);}.mobile-transparent-header #masthead, .mobile-transparent-header .site-top-header-wrap .site-header-row-container-inner, .mobile-transparent-header .site-main-header-wrap .site-header-row-container-inner, .mobile-transparent-header .site-bottom-header-wrap .site-header-row-container-inner{background:transparent;}.site-header-row-tablet-layout-fullwidth, .site-header-row-tablet-layout-standard{padding:0px;}}@media all and (min-width: 1025px){.transparent-header #masthead{position:absolute;left:0px;right:0px;z-index:100;}.transparent-header.kadence-scrollbar-fixer #masthead{right:var(--scrollbar-offset,0);}.transparent-header #masthead, .transparent-header .site-top-header-wrap .site-header-row-container-inner, .transparent-header .site-main-header-wrap .site-header-row-container-inner, .transparent-header .site-bottom-header-wrap .site-header-row-container-inner{background:transparent;}}.site-branding a.brand img{max-width:135px;}.site-branding a.brand img.svg-logo-image{width:135px;}.site-branding{padding:0px 0px 0px 0px;}#masthead, #masthead .kadence-sticky-header.item-is-fixed:not(.item-at-start):not(.site-header-row-container):not(.site-main-header-wrap), #masthead .kadence-sticky-header.item-is-fixed:not(.item-at-start) > .site-header-row-container-inner{background:#ffffff;}.site-main-header-wrap .site-header-row-container-inner{border-bottom:1px solid #cccccc;}.site-main-header-inner-wrap{min-height:49px;}.site-top-header-wrap .site-header-row-container-inner{background:var(--global-palette1);}.site-top-header-inner-wrap{min-height:0px;}.site-bottom-header-inner-wrap{min-height:0px;}#masthead .kadence-sticky-header.item-is-fixed:not(.item-at-start):not(.site-header-row-container):not(.item-hidden-above):not(.site-main-header-wrap), #masthead .kadence-sticky-header.item-is-fixed:not(.item-at-start):not(.item-hidden-above) > .site-header-row-container-inner{background:var(--global-palette9);}#masthead .kadence-sticky-header.item-is-fixed:not(.item-at-start) .site-branding .site-title, #masthead .kadence-sticky-header.item-is-fixed:not(.item-at-start) .site-branding .site-description{color:var(--global-palette3);}.header-navigation[class*="header-navigation-style-underline"] .header-menu-container.primary-menu-container>ul>li>a:after{width:calc( 100% - 2em);}.main-navigation .primary-menu-container > ul > li.menu-item > a{padding-left:calc(2em / 2);padding-right:calc(2em / 2);padding-top:0em;padding-bottom:0em;color:#4a5568;}.main-navigation .primary-menu-container > ul > li.menu-item .dropdown-nav-special-toggle{right:calc(2em / 2);}.main-navigation .primary-menu-container > ul > li.menu-item > a:hover{color:#000000;}.main-navigation .primary-menu-container > ul > li.menu-item.current-menu-item > a{color:#1a202c;}.header-navigation[class*="header-navigation-style-underline"] .header-menu-container.secondary-menu-container>ul>li>a:after{width:calc( 100% - 1.2em);}.secondary-navigation .secondary-menu-container > ul > li.menu-item > a{padding-left:calc(1.2em / 2);padding-right:calc(1.2em / 2);padding-top:0.6em;padding-bottom:0.6em;color:var(--global-palette9);background:var(--global-palette9);}.secondary-navigation .primary-menu-container > ul > li.menu-item .dropdown-nav-special-toggle{right:calc(1.2em / 2);}.secondary-navigation .secondary-menu-container > ul > li.menu-item > a:hover{color:#323a56;background:#323a56;}.secondary-navigation .secondary-menu-container > ul > li.menu-item.current-menu-item > a{color:#323a56;background:#323a56;}.header-navigation .header-menu-container ul ul.sub-menu, .header-navigation .header-menu-container ul ul.submenu{background:#1a202c;box-shadow:0px 2px 13px 0px rgba(0,0,0,0.1);}.header-navigation .header-menu-container ul ul li.menu-item, .header-menu-container ul.menu > li.kadence-menu-mega-enabled > ul > li.menu-item > a{border-bottom:1px none rgba(255,255,255,0.1);}.header-navigation .header-menu-container ul ul li.menu-item > a{width:100px;padding-top:4px;padding-bottom:4px;color:var(--global-palette8);font-style:normal;font-size:15px;}.header-navigation .header-menu-container ul ul li.menu-item > a:hover{color:var(--global-palette9);background:#323a56;}.header-navigation .header-menu-container ul ul li.menu-item.current-menu-item > a{color:var(--global-palette9);background:#2d3748;}.mobile-toggle-open-container .menu-toggle-open{color:var(--global-palette3);padding:0.4em 0.6em 0.4em 0.6em;font-size:14px;}.mobile-toggle-open-container .menu-toggle-open.menu-toggle-style-bordered{border:1px solid currentColor;}.mobile-toggle-open-container .menu-toggle-open .menu-toggle-icon{font-size:29px;}.mobile-toggle-open-container .menu-toggle-open:hover, .mobile-toggle-open-container .menu-toggle-open:focus{color:#087deb;}.mobile-navigation ul li{font-size:14px;}.mobile-navigation ul li a{padding-top:1em;padding-bottom:1em;}.mobile-navigation ul li > a, .mobile-navigation ul li.menu-item-has-children > .drawer-nav-drop-wrap{color:#f7fafc;}.mobile-navigation ul li > a:hover, .mobile-navigation ul li.menu-item-has-children > .drawer-nav-drop-wrap:hover{color:var(--global-palette9);}.mobile-navigation ul li.current-menu-item > a, .mobile-navigation ul li.current-menu-item.menu-item-has-children > .drawer-nav-drop-wrap{color:var(--global-palette9);}.mobile-navigation ul li.menu-item-has-children .drawer-nav-drop-wrap, .mobile-navigation ul li:not(.menu-item-has-children) a{border-bottom:1px solid rgba(255,255,255,0.1);}.mobile-navigation:not(.drawer-navigation-parent-toggle-true) ul li.menu-item-has-children .drawer-nav-drop-wrap button{border-left:1px solid rgba(255,255,255,0.1);}#mobile-drawer .drawer-inner, #mobile-drawer.popup-drawer-layout-fullwidth.popup-drawer-animation-slice .pop-portion-bg, #mobile-drawer.popup-drawer-layout-fullwidth.popup-drawer-animation-slice.pop-animated.show-drawer .drawer-inner{background:#323a56;}#mobile-drawer .drawer-header .drawer-toggle{padding:0.6em 0.15em 0.6em 0.15em;font-size:24px;}#mobile-drawer .drawer-header .drawer-toggle, #mobile-drawer .drawer-header .drawer-toggle:focus{color:var(--global-palette9);}#mobile-drawer .drawer-header .drawer-toggle:hover, #mobile-drawer .drawer-header .drawer-toggle:focus:hover{color:#0887fc;}#main-header .header-button{color:var(--global-palette9);background:var(--global-palette9);border:2px none transparent;box-shadow:0px 0px 0px -7px rgba(0,0,0,0);}#main-header .header-button:hover{color:#323a56;background:#323a56;box-shadow:0px 15px 25px -7px rgba(0,0,0,0.1);}.header-social-wrap .header-social-inner-wrap{font-size:1em;gap:0.3em;}.header-social-wrap .header-social-inner-wrap .social-button{border:2px none transparent;border-radius:3px;}.header-mobile-social-wrap .header-mobile-social-inner-wrap{font-size:1em;gap:0.3em;}.header-mobile-social-wrap .header-mobile-social-inner-wrap .social-button{border:2px none transparent;border-radius:3px;}.search-toggle-open-container .search-toggle-open{color:var(--global-palette5);}.search-toggle-open-container .search-toggle-open.search-toggle-style-bordered{border:1px solid currentColor;}.search-toggle-open-container .search-toggle-open .search-toggle-icon{font-size:1em;}.search-toggle-open-container .search-toggle-open:hover, .search-toggle-open-container .search-toggle-open:focus{color:var(--global-palette-highlight);}#search-drawer .drawer-inner{background:rgba(9, 12, 16, 0.97);}.mobile-header-button-wrap .mobile-header-button-inner-wrap .mobile-header-button{border:2px none transparent;box-shadow:0px 0px 0px -7px rgba(0,0,0,0);}.mobile-header-button-wrap .mobile-header-button-inner-wrap .mobile-header-button:hover{box-shadow:0px 15px 25px -7px rgba(0,0,0,0.1);}
/* Kadence Pro Header CSS */
.header-navigation-dropdown-direction-left ul ul.submenu, .header-navigation-dropdown-direction-left ul ul.sub-menu{right:0px;left:auto;}.rtl .header-navigation-dropdown-direction-right ul ul.submenu, .rtl .header-navigation-dropdown-direction-right ul ul.sub-menu{left:0px;right:auto;}.header-account-button .nav-drop-title-wrap > .kadence-svg-iconset, .header-account-button > .kadence-svg-iconset{font-size:1.2em;}.site-header-item .header-account-button .nav-drop-title-wrap, .site-header-item .header-account-wrap > .header-account-button{display:flex;align-items:center;}.header-account-style-icon_label .header-account-label{padding-left:5px;}.header-account-style-label_icon .header-account-label{padding-right:5px;}.site-header-item .header-account-wrap .header-account-button{text-decoration:none;box-shadow:none;color:inherit;background:transparent;padding:0.6em 0em 0.6em 0em;}.header-mobile-account-wrap .header-account-button .nav-drop-title-wrap > .kadence-svg-iconset, .header-mobile-account-wrap .header-account-button > .kadence-svg-iconset{font-size:1.2em;}.header-mobile-account-wrap .header-account-button .nav-drop-title-wrap, .header-mobile-account-wrap > .header-account-button{display:flex;align-items:center;}.header-mobile-account-wrap.header-account-style-icon_label .header-account-label{padding-left:5px;}.header-mobile-account-wrap.header-account-style-label_icon .header-account-label{padding-right:5px;}.header-mobile-account-wrap .header-account-button{text-decoration:none;box-shadow:none;color:inherit;background:transparent;padding:0.6em 0em 0.6em 0em;}#login-drawer .drawer-inner .drawer-content{display:flex;justify-content:center;align-items:center;position:absolute;top:0px;bottom:0px;left:0px;right:0px;padding:0px;}#loginform p label{display:block;}#login-drawer #loginform{width:100%;}#login-drawer #loginform input{width:100%;}#login-drawer #loginform input[type="checkbox"]{width:auto;}#login-drawer .drawer-inner .drawer-header{position:relative;z-index:100;}#login-drawer .drawer-content_inner.widget_login_form_inner{padding:2em;width:100%;max-width:350px;border-radius:.25rem;background:var(--global-palette9);color:var(--global-palette4);}#login-drawer .lost_password a{color:var(--global-palette6);}#login-drawer .lost_password, #login-drawer .register-field{text-align:center;}#login-drawer .widget_login_form_inner p{margin-top:1.2em;margin-bottom:0em;}#login-drawer .widget_login_form_inner p:first-child{margin-top:0em;}#login-drawer .widget_login_form_inner label{margin-bottom:0.5em;}#login-drawer hr.register-divider{margin:1.2em 0;border-width:1px;}#login-drawer .register-field{font-size:90%;}@media all and (min-width: 1025px){#login-drawer hr.register-divider.hide-desktop{display:none;}#login-drawer p.register-field.hide-desktop{display:none;}}@media all and (max-width: 1024px){#login-drawer hr.register-divider.hide-mobile{display:none;}#login-drawer p.register-field.hide-mobile{display:none;}}@media all and (max-width: 767px){#login-drawer hr.register-divider.hide-mobile{display:none;}#login-drawer p.register-field.hide-mobile{display:none;}}.tertiary-navigation .tertiary-menu-container > ul > li.menu-item > a{padding-left:calc(1.2em / 2);padding-right:calc(1.2em / 2);padding-top:0.6em;padding-bottom:0.6em;color:var(--global-palette5);}.tertiary-navigation .tertiary-menu-container > ul > li.menu-item > a:hover{color:var(--global-palette-highlight);}.tertiary-navigation .tertiary-menu-container > ul > li.menu-item.current-menu-item > a{color:var(--global-palette3);}.quaternary-navigation .quaternary-menu-container > ul > li.menu-item > a{padding-left:calc(1.2em / 2);padding-right:calc(1.2em / 2);padding-top:0.6em;padding-bottom:0.6em;color:var(--global-palette5);}.quaternary-navigation .quaternary-menu-container > ul > li.menu-item > a:hover{color:var(--global-palette-highlight);}.quaternary-navigation .quaternary-menu-container > ul > li.menu-item.current-menu-item > a{color:var(--global-palette3);}#main-header .header-divider{border-right:1px solid var(--global-palette6);height:50%;}#main-header .header-divider2{border-right:1px solid var(--global-palette6);height:50%;}#main-header .header-divider3{border-right:1px solid var(--global-palette6);height:50%;}#mobile-header .header-mobile-divider, #mobile-drawer .header-mobile-divider{border-right:1px solid var(--global-palette6);height:50%;}#mobile-drawer .header-mobile-divider{border-top:1px solid var(--global-palette6);width:50%;}#mobile-header .header-mobile-divider2{border-right:1px solid var(--global-palette6);height:50%;}#mobile-drawer .header-mobile-divider2{border-top:1px solid var(--global-palette6);width:50%;}.header-item-search-bar form ::-webkit-input-placeholder{color:currentColor;opacity:0.5;}.header-item-search-bar form ::placeholder{color:currentColor;opacity:0.5;}.header-search-bar form{max-width:100%;width:240px;}.header-mobile-search-bar form{max-width:calc(100vw - var(--global-sm-spacing) - var(--global-sm-spacing));width:240px;}.header-widget-lstyle-normal .header-widget-area-inner a:not(.button){text-decoration:underline;}.element-contact-inner-wrap{display:flex;flex-wrap:wrap;align-items:center;margin-top:-0.6em;margin-left:calc(-0.6em / 2);margin-right:calc(-0.6em / 2);}.element-contact-inner-wrap .header-contact-item{display:inline-flex;flex-wrap:wrap;align-items:center;margin-top:0.6em;margin-left:calc(0.6em / 2);margin-right:calc(0.6em / 2);}.element-contact-inner-wrap .header-contact-item .kadence-svg-iconset{font-size:1em;}.header-contact-item img{display:inline-block;}.header-contact-item .contact-label{margin-left:0.3em;}.rtl .header-contact-item .contact-label{margin-right:0.3em;margin-left:0px;}.header-mobile-contact-wrap .element-contact-inner-wrap{display:flex;flex-wrap:wrap;align-items:center;margin-top:-0.6em;margin-left:calc(-0.6em / 2);margin-right:calc(-0.6em / 2);}.header-mobile-contact-wrap .element-contact-inner-wrap .header-contact-item{display:inline-flex;flex-wrap:wrap;align-items:center;margin-top:0.6em;margin-left:calc(0.6em / 2);margin-right:calc(0.6em / 2);}.header-mobile-contact-wrap .element-contact-inner-wrap .header-contact-item .kadence-svg-iconset{font-size:1em;}#main-header .header-button2{border:2px none transparent;box-shadow:0px 0px 0px -7px rgba(0,0,0,0);}#main-header .header-button2:hover{box-shadow:0px 15px 25px -7px rgba(0,0,0,0.1);}.mobile-header-button2-wrap .mobile-header-button-inner-wrap .mobile-header-button2{border:2px none transparent;box-shadow:0px 0px 0px -7px rgba(0,0,0,0);}.mobile-header-button2-wrap .mobile-header-button-inner-wrap .mobile-header-button2:hover{box-shadow:0px 15px 25px -7px rgba(0,0,0,0.1);}#widget-drawer.popup-drawer-layout-fullwidth .drawer-content .header-widget2, #widget-drawer.popup-drawer-layout-sidepanel .drawer-inner{max-width:400px;}#widget-drawer.popup-drawer-layout-fullwidth .drawer-content .header-widget2{margin:0 auto;}.widget-toggle-open{display:flex;align-items:center;background:transparent;box-shadow:none;}.widget-toggle-open:hover, .widget-toggle-open:focus{border-color:currentColor;background:transparent;box-shadow:none;}.widget-toggle-open .widget-toggle-icon{display:flex;}.widget-toggle-open .widget-toggle-label{padding-right:5px;}.rtl .widget-toggle-open .widget-toggle-label{padding-left:5px;padding-right:0px;}.widget-toggle-open .widget-toggle-label:empty, .rtl .widget-toggle-open .widget-toggle-label:empty{padding-right:0px;padding-left:0px;}.widget-toggle-open-container .widget-toggle-open{color:var(--global-palette5);padding:0.4em 0.6em 0.4em 0.6em;font-size:14px;}.widget-toggle-open-container .widget-toggle-open.widget-toggle-style-bordered{border:1px solid currentColor;}.widget-toggle-open-container .widget-toggle-open .widget-toggle-icon{font-size:20px;}.widget-toggle-open-container .widget-toggle-open:hover, .widget-toggle-open-container .widget-toggle-open:focus{color:var(--global-palette-highlight);}#widget-drawer .header-widget-2style-normal a:not(.button){text-decoration:underline;}#widget-drawer .header-widget-2style-plain a:not(.button){text-decoration:none;}#widget-drawer .header-widget2 .widget-title{color:var(--global-palette9);}#widget-drawer .header-widget2{color:var(--global-palette8);}#widget-drawer .header-widget2 a:not(.button), #widget-drawer .header-widget2 .drawer-sub-toggle{color:var(--global-palette8);}#widget-drawer .header-widget2 a:not(.button):hover, #widget-drawer .header-widget2 .drawer-sub-toggle:hover{color:var(--global-palette9);}#mobile-secondary-site-navigation ul li{font-size:14px;}#mobile-secondary-site-navigation ul li a{padding-top:1em;padding-bottom:1em;}#mobile-secondary-site-navigation ul li > a, #mobile-secondary-site-navigation ul li.menu-item-has-children > .drawer-nav-drop-wrap{color:var(--global-palette8);}#mobile-secondary-site-navigation ul li.current-menu-item > a, #mobile-secondary-site-navigation ul li.current-menu-item.menu-item-has-children > .drawer-nav-drop-wrap{color:var(--global-palette-highlight);}#mobile-secondary-site-navigation ul li.menu-item-has-children .drawer-nav-drop-wrap, #mobile-secondary-site-navigation ul li:not(.menu-item-has-children) a{border-bottom:1px solid rgba(255,255,255,0.1);}#mobile-secondary-site-navigation:not(.drawer-navigation-parent-toggle-true) ul li.menu-item-has-children .drawer-nav-drop-wrap button{border-left:1px solid rgba(255,255,255,0.1);}
</style>







<style id='kadence-blocks-global-variables-inline-css'>
:root {--global-kb-font-size-sm:clamp(0.8rem, 0.73rem + 0.217vw, 0.9rem);--global-kb-font-size-md:clamp(1.1rem, 0.995rem + 0.326vw, 1.25rem);--global-kb-font-size-lg:clamp(1.75rem, 1.576rem + 0.543vw, 2rem);--global-kb-font-size-xl:clamp(2.25rem, 1.728rem + 1.63vw, 3rem);--global-kb-font-size-xxl:clamp(2.5rem, 1.456rem + 3.26vw, 4rem);--global-kb-font-size-xxxl:clamp(2.75rem, 0.489rem + 7.065vw, 6rem);}
</style>
<style id='kadence_blocks_css-inline-css'>
.kb-table-of-content-nav.kb-table-of-content-id_1b49da-72 .kb-table-of-content-wrap{background-color:#edf2f7;border-top-width:1px;border-right-width:1px;border-bottom-width:1px;border-left-width:1px;box-shadow:rgba(0, 0, 0, 0.2) 0px 0px 14px 0px;max-width:450px;}.kb-table-of-content-nav.kb-table-of-content-id_1b49da-72 .kb-toggle-icon-style-basiccircle .kb-table-of-contents-icon-trigger:after, .kb-table-of-content-nav.kb-table-of-content-id_1b49da-72 .kb-toggle-icon-style-basiccircle .kb-table-of-contents-icon-trigger:before, .kb-table-of-content-nav.kb-table-of-content-id_1b49da-72 .kb-toggle-icon-style-arrowcircle .kb-table-of-contents-icon-trigger:after, .kb-table-of-content-nav.kb-table-of-content-id_1b49da-72 .kb-toggle-icon-style-arrowcircle .kb-table-of-contents-icon-trigger:before, .kb-table-of-content-nav.kb-table-of-content-id_1b49da-72 .kb-toggle-icon-style-xclosecircle .kb-table-of-contents-icon-trigger:after, .kb-table-of-content-nav.kb-table-of-content-id_1b49da-72 .kb-toggle-icon-style-xclosecircle .kb-table-of-contents-icon-trigger:before{background-color:#edf2f7;}
</style>
<link rel="https://api.w.org/" href="https://www.guru99.com/wp-json/" /><link rel="alternate" type="application/json" href="https://www.guru99.com/wp-json/wp/v2/posts/2" /><!-- / HREFLANG Tags by DCGWS Version 2.0.0 -->
<!-- / HREFLANG Tags by DCGWS -->
<style type="text/css">
			:root{
				--lasso-main: #5e36ca !important;
				--lasso-title: black !important;
				--lasso-button: #22baa0 !important;
				--lasso-secondary-button: #22baa0 !important;
				--lasso-button-text: white !important;
				--lasso-background: white !important;
				--lasso-pros: #22baa0 !important;
				--lasso-cons: #e06470 !important;
			}
			
			
		</style>
			<script type="rocketlazyloadscript" data-rocket-type="text/javascript">
				// Notice how this gets configured before we load Font Awesome
				window.FontAwesomeConfig = { autoReplaceSvg: false }
			</script>
		<link rel="pingback" href="https://www.guru99.com/xmlrpc.php">
<meta name="theme-color" content="#2cbaff">

<!-- PLACE THIS SECTION INSIDE OF YOUR HEAD TAGS -->
<!-- Below is a recommended list of pre-connections, which allow the network to establish each connection quicker, speeding up response times and improving ad performance. -->
<link rel="preconnect" href="https://a.pub.network/" crossorigin="">
<link rel="preconnect" href="https://b.pub.network/" crossorigin="">
<link rel="preconnect" href="https://c.pub.network/" crossorigin="">
<link rel="preconnect" href="https://d.pub.network/" crossorigin="">
<link rel="preconnect" href="https://c.amazon-adsystem.com" crossorigin="">
<link rel="preconnect" href="https://s.amazon-adsystem.com" crossorigin="">
<link rel="preconnect" href="https://secure.quantserve.com/" crossorigin="">
<link rel="preconnect" href="https://rules.quantcount.com/" crossorigin="">
<link rel="preconnect" href="https://pixel.quantserve.com/" crossorigin="">
<link rel="preconnect" href="https://cmp.quantcast.com/" crossorigin="">
<link rel="preconnect" href="https://btloader.com/" crossorigin="">
<link rel="preconnect" href="https://api.btloader.com/" crossorigin="">
<link rel="preconnect" href="https://confiant-integrations.global.ssl.fastly.net" crossorigin="">
<!-- Below is a link to a CSS file that accounts for Cumulative Layout Shift, a new Core Web Vitals subset that Google uses to help rank your site in search -->
<!-- The file is intended to eliminate the layout shifts that are seen when ads load into the page. If you don't want to use this, simply remove this file -->
<!-- To find out more about CLS, visit https://web.dev/vitals/ -->
<link rel="stylesheet" href="https://a.pub.network/guru99-com/cls.css">
<script data-cfasync="false" type="text/javascript">
  var freestar = freestar || {};
  freestar.queue = freestar.queue || [];
  freestar.config = freestar.config || {};
  freestar.config.disabledProducts = {
    googleInterstitial: true 
};
  freestar.config.enabled_slots = [];
  freestar.initCallback = function () { (freestar.config.enabled_slots.length === 0) ? freestar.initCallbackCalled = false : freestar.newAdSlots(freestar.config.enabled_slots) }
</script>
<script src="https://a.pub.network/guru99-com/pubfig.min.js" data-cfasync="false" async=""></script>

<link rel="preload" href="https://www.guru99.com/wp-content/fonts/source-sans-pro/6xK3dSBYKcSV-LCoeQqfX1RYOo3qOK7l.woff2" as="font" type="font/woff2" crossorigin>
<link rel="icon" href="https://www.guru99.com/images/favicon-new-logo.png" sizes="32x32" />
<link rel="icon" href="https://www.guru99.com/images/favicon-new-logo.png" sizes="192x192" />
<link rel="apple-touch-icon" href="https://www.guru99.com/images/favicon-new-logo.png" />
<meta name="msapplication-TileImage" content="https://www.guru99.com/images/favicon-new-logo.png" />
		<style id="wp-custom-css">
			body {
	--global-body-font-family: 'Source Sans Pro', sans-serif;
}

.content-wrap .entry img,
.content-wrap .entry p img {
    margin: 0 auto;
}
hr{
	border-bottom:none;
}
hr{
	border-top: 1px solid #eee;
	margin-top: 20px !important;
}
.entry-content a:hover {
background: #ffec54;
}
a{
	text-decoration:none;
}
table{
	border-spacing: 0 !important;
	border:0;
	border-collapse: collapse;
}td{
	padding: 0.5rem;
}
th{
	padding: 0.5rem;
	border:0;
	text-align: left !important;
}
.table td {
    border: 0px;
    border-top: 1px solid #eee;
}
tbody tr:nth-child(2n+1) td, tr:nth-child(2n+1) th {
    background: #f9f9f9;
}
.key-difference {
    border: 1px solid #d6d6d6;
    background-color: #e0f1f5;
    padding: 0.938rem;
	  margin-bottom: 20px;
}
.img_caption {
    text-align: center !important;
}
.alert.alert-error { 
	background-color: #f6e7e7;
	border: 1px solid #edd1d0;
	border-radius: 0.1875rem;
	box-sizing: inherit;
	color: #b94a48;
	margin: 1.5rem 0px;
	margin-bottom: 1.5rem;
	padding: 0.938rem;
	text-align: center;
	text-shadow: none;
}
.alert-error a {
    color: #000;
    font-weight: bold;
    text-decoration: none;
}
.alert.alert-success { 
	background-color: #dfeedf;
	border: 1px solid #c4e0c4;
	border-radius: 0.1875rem;
	box-sizing: inherit;
	color: #468847;
	list-style: outside none none;
	margin: 1.5rem 0px;
	margin-bottom: 1.5rem;
	padding: 0.938rem;
	text-align: center;
	text-shadow: none;
}
.alert-success a{
    color: #356635;
    font-weight: bold;
}
.alert.alert-info { 
	background-color: #e2eff5;
	border: 1px solid #c7e0ec;
	border-radius: 0.1875rem;
	border-top-left-radius: 3px;
	border-top-right-radius: 3px;
	box-sizing: inherit;
	color: #3a87ad;
	list-style: outside none none;
	margin: 1.5rem 0px;
	margin-bottom: 1.5rem;
	padding: 0.938rem;
	text-shadow: none;
}
.alert-info a{
color: #2d6987;
    font-weight: bold;
}
body p{
    margin: 0 0 1.3rem 0 !important;
}
.review-border{
border:1px solid #eee;
}
h1 a, h2 a, h3 a, h4 a, h5 a, h6 a{
color: #0556f3;
}
.alert.alert-warning { 
	background-color: #f8f4ec;
	border: 1px solid #eee4d2;
	border-radius: 0.1875rem;
	box-sizing: inherit;
	color: #c09853;
	list-style: outside none none;
	margin: 1.5rem 0px;
	margin-bottom: 1.5rem;
	padding: 0.938rem;
	text-shadow: none;
}
.alert-warning a{
    color: #6c5328;
    font-weight: bold;
}
code{
background-color: #f7f7f7;
color: #9c1d3d;
padding: 2px 4px;
border: 1px solid rgba(0,0,0,0.1);
font-size: 1rem;
border-radius: 0.1875rem;
}
.button1 {
    background: #2f81ff;
    color: #fff!important;
    font-size: 14px;
    padding: 8px 13px;
    text-align: center;
    text-transform: none;
    white-space: nowrap;
}
ul, ol, dl {
    margin-top: 1.5rem !important;
    margin-bottom: 1.5rem !important;
}
img{
display: inline-block;
}
h1{
	margin-top: 10px !important;
	}	
h2, h3, h4, h5{
	margin: 1.5rem 0 0.75rem 0 !important;
}
.with-ribbon {
	position: relative;
}

.with-ribbon figcaption {
	position: absolute;
	right: 0;
	top: 0;
	padding: 10px;
	display: inline-block;
	color: #fff;
	background: red;
}
.nav-link-center {
    order: 1;
}

.nav-previous {
    order: 0;
}

.nav-next {
    order: 2;
}
.single-content h2:first-child{
    margin-top: 0px !important;
}
.single-content h3{
margin-top: 0px;
}
.single-content h2{
margin-top: 0px !important;
}
.entry-content{
	margin-top: 0px !important;
}
.entry-meta{
	margin-bottom: 0px !important;
}
.entry-header{
	margin-bottom: 0px !important;
}
.tool-sticky th{
border:1px solid #eee !important;
background: #ffe !important;
}
.tool-sticky td{
border: 1px solid #eee !important;
}

.tool-sticky tbody tr:nth-child(2n+1) td{
background: #fff;
}
.button1 {
    background: #2f81ff;
    color: #fff!important;
    font-size: 14px;
    padding: 8px 13px;
    text-align: center;
    text-transform: none;
    white-space: nowrap;
}

th{
background: #f2f2f2;
}

@media only screen and (max-width: 1023px) {
table {
display: block;
overflow: scroll;
overflow-x: auto;
overflow-y: auto;
}
}
.pagenav{
    background: #df5035;
    font-size: 1rem;
    border-radius: 5px;
    border: 0px;
    padding: 0.8rem 1rem;
	color:#fff;
}
.comment-navigation .nav-previous:after, .post-navigation .nav-previous:after{
    position: inherit;
}
.header-menu-container ul.menu>li.kadence-menu-mega-columns-3>ul.sub-menu { 
grid-template-columns: 30% 30% 30%; 
}
.single-post .entry-header {
margin-bottom: 0px !important;
}
.comment-navigation .nav-links, .post-navigation .nav-links {
display: flex !important;
flex-flow: row !important;
justify-content: space-between !important;
}
.site-header-row {
display: flex !important;
justify-content: space-evenly;
}
.header-navigation ul {
margin: 0 !important;
}
.header-menu-container ul.menu>li.kadence-menu-mega-width-custom>ul.sub-menu {
transition-duration: .5s !important;
}
@media (max-width: 767px) {
  .hidden-phone { 
   display: none !important;
	}}
	
.vs-sticky{
  min-width: 100px;
  max-width: 300px;
  left: 0px;
  position: sticky;
  background-color: white !important;
}
@media (max-width: 767px){
.kt-row-column-wrap.kt-mobile-layout-row>.wp-block-kadence-column {
    margin-bottom: 0px !important;
}}

	.wp-has-aspect-ratio{	
--aspect-ratio:56.25% !important;
}
.wgs_wrapper td.gsib_a{
padding: 0px;
    background: none;
}
.wgs_wrapper .gsc-input-box{
border:1px solid black;
}

@media(max-width: 360px) { .responsivetable{ width: 45%; } }
@media screen and (max-width: 540px) and (min-width: 361px) { .responsivetable{ width: 44%; } }
@media screen and (max-width: 959px) and (min-width: 541px) { .responsivetable{ width: 30%; } }
@media screen and (max-width: 1599px) and (min-width: 960px) { .responsivetable{ width: 17%; } }
@media screen and (min-width: 1600px) { .responsivetable{ width: 17%; } }

h1, h2, h3, h4, h5, h6 {
	font-weight: 700 !important;
}
.wp-block-latest-posts.wp-block-latest-posts__list.is-grid li>a{
color:#0556f3;
}

div.w3-container.w3-half { 
	box-sizing: border-box;
	float: left;
	width: 100%;
}
div.w3-row.w3-border::after { 
	clear: both;
	content: "";
	display: table;
}
div.w3-row.w3-border::before { 
	clear: both;
	content: "";
	display: table;
}
@media (min-width: 601px) { 
	div.w3-container.w3-half { 
		width: 50%;
	}
}
.top-pros{
background:green;
color:#FFF;
margin-right: 10px !important;
padding:5px;
}
.top-cons{
background:darkred;
color:#FFF;
margin-left: 10px !important;
padding:5px;
}

.entry-content  a.nohover:hover {
background: transparent;
}
div.lasso-grid-row .lasso-description {
min-height: 10px;
}
div.lasso-grid-row .lasso-badge {
color: #fff;
background:#5e36ca !important;
}
div.lasso-grid-row .lasso-description {
font-size: 20px;
}
.lasso-grid-row .lasso-splash .lasso-title {
    min-height: 10px;
}
a.lasso-button-1{
	background: #2f81ff !important;
}
@media screen and (max-width: 1200px){
div.lasso-grid-row .lasso-description {
    min-height: 10px !important;
	}}

.hilr {
    background-color: #ffb1b5 !important;
}
.hilb {
    background-color: #c1f7ff !important;
}
.hilight {
    background-color: yellow !important;
}
a:hover.button1 {
background: #2f81ff !important;
}

.header-menu-container ul.menu>li.menu-item-has-children>ul.sub-menu {
	visibility: hidden !important;
}

.header-menu-container ul.menu>li.menu-item-has-children>ul.sub-menu.show {
    visibility: visible !important;
    opacity: 1;
    clip: auto;
    height: auto;
    overflow: visible;
}

.lasso-badge{
z-index: 10;
}
.header-menu-container ul.menu>li.kadence-menu-mega-enabled>ul a {
    width: 100% !important;
}

@media (max-width: 500px) {
.entry-meta-divider-customicon span.meta-label {
display: none;
}
}

@media (max-width: 1024px) {
.primary-sidebar.widget-area{
	display: none;
}}

.toolbutton {
background: #f68700 !important;
border-radius: 1000px;
padding: 10px 27px;
color: #ffffff !important;
display: inline-block;
font-weight: bold;
font-size: 27px;
letter-spacing: 0.8px;
}
a:hover.toolbutton {
background: #ff9f00 !important;
color: #ffffff !important;
}
.site-main-header-wrap .site-header-row-container-inner {
border-bottom: 1px solid #cccccc;
}
.box12{
border: 0.3px solid #eee;
    box-sizing: border-box;
    border-radius: 8px;
    padding-top: 10px;
    padding-left: 15px;
    line-height: 1.8em;
    background: #F6FCFF;
}

div.w3-topta-container1.w3-topta-half1 { 
	box-sizing: border-box;
	float: left;
	width: 100%;
	border: 1px solid #e0def5;
	margin: 5px;
	border-radius: 15px;
	padding: 10px;
	background-color: #f2f1fb;
}
div.w3-topta-row1.w3-topta-border1::after { 
	clear: both;
	content: "";
	display: table;
}
div.w3-topta-row1.w3-topta-border1::before { 
	clear: both;
	content: "";
	display: table;
}
@media (min-width: 766px) { 
	div.w3-topta-container1.w3-topta-half1 { 
		width: 18.5%;
	}
}
@media (min-width: 766px) { 
	div.topta-lastbox { 
		width: 19% !important;
	}
}

.topta-button2 {
    background: #2f81ff !important;
    color: #fff!important;
    font-size: 18px;
    padding: 10px 50px;
    text-align: center;
    text-transform: none;
    white-space: nowrap;
	border-radius: 1000px;
}
@media only screen and (min-width: 767px) and (max-width: 1023px)  {.topta-button2{ padding: 10px 15px !important; } }
@media only screen and (min-width: 1024px) and (max-width: 1149px)  {.topta-button2{ padding: 10px 30px !important; } }

div.elementor-widget-topta-container99 { 
	box-sizing: border-box;
	color: #111111;
	font-size: 15px;
	line-height: 25.5px;
	word-wrap: break-word;
margin-bottom:15px;
}

div.top-3__topta-best-choise99 { 
	align-items: center;
	background: #5e36ca;
	bottom: 0px;
	box-sizing: border-box;
	color: white;
	display: flex;
	font-size: 15px;
	font-weight: 600;
	height: 40px;
	justify-content: center;
	left: 0px;
	line-height: 25.5px;
	margin: -30px auto 0px;
	position: relative;
	right: 0px;
	text-align: center;
	text-transform: capitalize;
	top: 0px;
	width: 150px;
	word-wrap: break-word;
}

div.top-3__topta-best-choise99::before { 
	border-style: solid;
	border-width: 0px 0px 20px 20px;
	content: "";
	left: 0px;
	margin-left: -20px;
	position: absolute;
	top: 0px;
	border-color: transparent transparent #1e0b7c transparent;
}
@media (max-width: 766px) { 
div.top-3__topta-best-choise99{
	margin: -15px auto 0px !important;
}
}

#more1 {display: none;}

.kt-blocks-accordion-header {
    background: #f7f9fe !important;
}
.kt-blocks-accordion-header:hover{
    background: #ffffff !important;
}

.you-might-like{
font-size: 25px;
line-height: 25px;
font-weight: 700;
font-style: normal;
}
.circle1{
	width:22px;
	height:17px;
	display:inline-block;
    background-image: url("https://www.guru99.com/images/iconpros.png");
	background-repeat: no-repeat;
}
.circle2{
	display: inline-block;
	width: 22px;
    height: 17px;
    background-image: url("https://www.guru99.com/images/iconcons.png");
    background-repeat: no-repeat;
}
ul.pros1-icon{
	margin-top: 10px !important;
    padding-left: 25px !important;
}
ul.cons1-icon{
	margin-top: 10px !important;
    padding-left: 25px !important;
}
.prcons-add-bor{
	border-bottom: 1px solid #222222;
	padding-bottom: 8px;
	margin-top:0px !important;
}
div.w3-containerp99.w3-halfp99 { 
	box-sizing: border-box;
	float: left;
	width: 100%;
}
@media (min-width: 601px) {div.w3-containerp99.w3-halfp99 {width: 48%;}}
@media (min-width: 601px) {div.w3-pros-mar {margin-right:15px;}}
@media (min-width: 601px) {div.w3-cons-mal {margin-left:15px;}}

.coauthors {
    display: flex;
    align-items: center;
    font-weight: normal;
}
.coauthors img {
    display: block !important;
    border-radius: 20px;
    margin-left: 5px !important;
}
.coauthors a {
    display: flex;
    margin-right: 1em;
    font-weight: normal;
    grid-gap: 8px;
}
@media only screen and (max-width: 600px) {
.nav-previous>a:lang(fr) {
	width:155px!important;
}
}

@media only screen and (max-width: 600px) {
.nav-next>a:lang(fr) {
	width:130px!important;
}
}
.main-navigation .primary-menu-container > ul > li.menu-item > a:lang(fr){
padding-left:12px;
padding-right:12px;
}		</style>
		<noscript><style id="rocket-lazyload-nojs-css">.rll-youtube-player, [data-lazy-src]{display:none !important;}</style></noscript></head>

<body class="post-template-default single single-post postid-2 single-format-standard wp-custom-logo wp-embed-responsive lasso-v318 footer-on-bottom hide-focus-outline link-style-standard has-sidebar has-sticky-sidebar-widget content-title-style-normal content-width-normal content-style-unboxed content-vertical-padding-hide non-transparent-header mobile-non-transparent-header">
<div id="wrapper" class="site wp-site-blocks">
			<a class="skip-link screen-reader-text scroll-ignore" href="#main">Skip to content</a>
		<header id="masthead" class="site-header" role="banner" itemtype="https://schema.org/WPHeader" itemscope>
	<div id="main-header" class="site-header-wrap">
		<div class="site-header-inner-wrap">
			<div class="site-header-upper-wrap">
				<div class="site-header-upper-inner-wrap">
					<div class="site-main-header-wrap site-header-row-container site-header-focus-item site-header-row-layout-standard" data-section="kadence_customizer_header_main">
	<div class="site-header-row-container-inner">
				<div class="site-container">
			<div class="site-main-header-inner-wrap site-header-row site-header-row-has-sides site-header-row-no-center">
									<div class="site-header-main-section-left site-header-section site-header-section-left">
						<div class="site-header-item site-header-focus-item" data-section="title_tagline">
	<div class="site-branding branding-layout-standard site-brand-logo-only"><a class="brand has-logo-image" href="https://www.guru99.com/" rel="home" aria-label="Guru99"><img width="300" height="59" src="https://www.guru99.com/images/guru99-logo-v1.png" class="custom-logo" alt="Guru99" decoding="async" /></a></div></div><!-- data-section="title_tagline" -->
					</div>
																	<div class="site-header-main-section-right site-header-section site-header-section-right">
						<div class="site-header-item site-header-focus-item site-header-item-main-navigation header-navigation-layout-stretch-false header-navigation-layout-fill-stretch-false" data-section="kadence_customizer_primary_navigation">
		<nav id="site-navigation" class="main-navigation header-navigation nav--toggle-sub header-navigation-style-standard header-navigation-dropdown-animation-none" role="navigation" aria-label="Primary Navigation">
				<div class="primary-menu-container header-menu-container">
			<ul id="primary-menu" class="menu"><li id="menu-item-3172" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-3172"><a href="/">Home</a></li>
<li id="menu-item-3173" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3173 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-3 kadence-menu-mega-layout-equal"><a href="/software-testing.html"><span class="nav-drop-title-wrap">Testing<span class="dropdown-nav-toggle"><span class="kadence-svg-iconset svg-baseline"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></span></span></a>
<ul class="sub-menu">
	<li id="menu-item-4569" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4569"><a href="https://www.guru99.com/agile-testing-course.html">Agile Testing</a></li>
	<li id="menu-item-4572" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4572"><a href="/junit-tutorial.html">JUnit</a></li>
	<li id="menu-item-4579" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4579"><a href="/hp-alm-free-tutorial.html">Quality Center(ALM)</a></li>
	<li id="menu-item-4570" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4570"><a href="/bugzilla-tutorial-for-beginners.html">Bugzilla</a></li>
	<li id="menu-item-4584" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4584"><a href="/loadrunner-v12-tutorials.html">HP Loadrunner</a></li>
	<li id="menu-item-4593" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4593"><a href="/rpa-tutorial.html">RPA</a></li>
	<li id="menu-item-4571" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4571"><a href="https://www.guru99.com/cucumber-tutorials.html">Cucumber</a></li>
	<li id="menu-item-4600" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4600"><a href="/software-testing.html">Software Testing</a></li>
	<li id="menu-item-4606" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4606"><a href="/learn-sap-testing-create-your-first-sap-test-case.html">SAP Testing</a></li>
	<li id="menu-item-4608" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4608"><a href="https://www.guru99.com/data-testing.html">Database Testing</a></li>
	<li id="menu-item-4616" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4616"><a href="/mobile-testing.html">Mobile Testing</a></li>
	<li id="menu-item-4622" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4622"><a href="/selenium-tutorial.html">Selenium</a></li>
	<li id="menu-item-4626" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4626"><a href="/utlimate-guide-etl-datawarehouse-testing.html">ETL Testing</a></li>
	<li id="menu-item-4628" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4628"><a href="https://www.guru99.com/mantis-bug-tracker-tutorial.html">Mantis</a></li>
	<li id="menu-item-4635" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4635"><a href="https://www.guru99.com/soapui-tutorial.html">SoapUI</a></li>
	<li id="menu-item-4640" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4640"><a href="https://www.guru99.com/jmeter-tutorials.html">JMeter</a></li>
	<li id="menu-item-4646" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4646"><a href="https://www.guru99.com/postman-tutorial.html">Postman</a></li>
	<li id="menu-item-4653" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4653"><a href="https://www.guru99.com/test-management.html">TEST Management</a></li>
	<li id="menu-item-4658" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4658"><a href="https://www.guru99.com/jira-tutorial-a-complete-guide-for-beginners.html">JIRA</a></li>
	<li id="menu-item-4663" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4663"><a href="https://www.guru99.com/quick-test-professional-qtp-tutorial.html">QTP</a></li>
	<li id="menu-item-4665" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4665"><a href="https://www.guru99.com/testlink-tutorial-complete-guide.html">TestLink</a></li>
</ul>
</li>
<li id="menu-item-3174" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3174 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-3 kadence-menu-mega-layout-equal"><a href="/what-is-sap.html"><span class="nav-drop-title-wrap">SAP<span class="dropdown-nav-toggle"><span class="kadence-svg-iconset svg-baseline"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></span></span></a>
<ul class="sub-menu">
	<li id="menu-item-4678" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4678"><a href="https://www.guru99.com/abap-tutorial.html">ABAP</a></li>
	<li id="menu-item-4681" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4681"><a href="https://www.guru99.com/sap-crm-training.html">CRM</a></li>
	<li id="menu-item-4683" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4683"><a href="https://www.guru99.com/sap-pi-process-integration-tutorial.html">PI/PO</a></li>
	<li id="menu-item-4685" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4685"><a href="https://www.guru99.com/overview-of-sap-apo.html">APO</a></li>
	<li id="menu-item-4689" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4689"><a href="https://www.guru99.com/crystal-reports-tutorial.html">Crystal Reports</a></li>
	<li id="menu-item-4695" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4695"><a href="https://www.guru99.com/sap-pp-tutorials.html">PP</a></li>
	<li id="menu-item-4699" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4699"><a href="https://www.guru99.com/what-is-sap.html">Beginners</a></li>
	<li id="menu-item-4707" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4707"><a href="https://www.guru99.com/sap-fico-training-tutorials.html">FICO</a></li>
	<li id="menu-item-4709" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4709"><a href="https://www.guru99.com/free-sap-sd-training-course.html">SD</a></li>
	<li id="menu-item-4713" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4713"><a href="https://www.guru99.com/sap-basis-training-tutorials.html">Basis</a></li>
	<li id="menu-item-4717" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4717"><a href="https://www.guru99.com/sap-hana-tutorial.html">HANA</a></li>
	<li id="menu-item-4721" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4721"><a href="https://www.guru99.com/sapui5-tutorial.html">SAPUI5</a></li>
	<li id="menu-item-4732" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4732"><a href="https://www.guru99.com/sap-bods-tutorial.html">BODS</a></li>
	<li id="menu-item-4739" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4739"><a href="https://www.guru99.com/sap-hcm.html">HR</a></li>
	<li id="menu-item-4744" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4744"><a href="https://www.guru99.com/overview-of-sap-security.html">Security Tutorial</a></li>
	<li id="menu-item-4748" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4748"><a href="https://www.guru99.com/sap-bi.html">BI/BW</a></li>
	<li id="menu-item-4759" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4759"><a href="https://www.guru99.com/sap-mm-training-tutorials.html">MM</a></li>
	<li id="menu-item-4762" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4762"><a href="https://www.guru99.com/overview-of-sap-solution-manager.html">Solution Manager</a></li>
	<li id="menu-item-4764" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4764"><a href="https://www.guru99.com/sap-bpc.html">BPC</a></li>
	<li id="menu-item-4768" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4768"><a href="https://www.guru99.com/sap-quality-management-qm-tutorial.html">QM</a></li>
	<li id="menu-item-4771" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4771"><a href="https://www.guru99.com/sap-successfactor.html">Successfactors</a></li>
	<li id="menu-item-4773" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4773"><a href="https://www.guru99.com/co-tutorials.html">CO</a></li>
	<li id="menu-item-4775" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4775"><a href="https://www.guru99.com/sap-payroll.html">Payroll</a></li>
	<li id="menu-item-4778" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4778"><a href="https://www.guru99.com/sap-training-hub.html">SAP Courses</a></li>
</ul>
</li>
<li id="menu-item-3175" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3175 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-4 kadence-menu-mega-layout-equal"><a href="/java-tutorial.html"><span class="nav-drop-title-wrap">Web<span class="dropdown-nav-toggle"><span class="kadence-svg-iconset svg-baseline"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></span></span></a>
<ul class="sub-menu">
	<li id="menu-item-4793" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4793"><a href="https://www.guru99.com/apache.html">Apache</a></li>
	<li id="menu-item-4796" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4796"><a href="https://www.guru99.com/java-tutorial.html">Java</a></li>
	<li id="menu-item-4799" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4799"><a href="https://www.guru99.com/php-tutorials.html">PHP</a></li>
	<li id="menu-item-4800" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4800"><a href="https://www.guru99.com/ms-sql-server-tutorial.html">SQL Server</a></li>
	<li id="menu-item-4802" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4802"><a href="https://www.guru99.com/angularjs-tutorial.html">AngularJS</a></li>
	<li id="menu-item-4805" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4805"><a href="https://www.guru99.com/jsp-tutorial.html">JSP</a></li>
	<li id="menu-item-4806" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4806"><a href="https://www.guru99.com/pl-sql-tutorials.html">PL/SQL</a></li>
	<li id="menu-item-4809" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4809"><a href="https://www.guru99.com/uml-tutorial.html">UML</a></li>
	<li id="menu-item-4811" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4811"><a href="https://www.guru99.com/asp-net-tutorial.html">ASP.NET</a></li>
	<li id="menu-item-4817" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4817"><a href="https://www.guru99.com/kotlin-tutorial.html">Kotlin</a></li>
	<li id="menu-item-4819" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4819"><a href="https://www.guru99.com/postgresql-tutorial.html">PostgreSQL</a></li>
	<li id="menu-item-4824" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4824"><a href="https://www.guru99.com/vb-net-tutorial.html">VB.NET</a></li>
	<li id="menu-item-4827" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4827"><a href="https://www.guru99.com/c-programming-tutorial.html">C</a></li>
	<li id="menu-item-4830" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4830"><a href="https://www.guru99.com/unix-linux-tutorial.html">Linux</a></li>
	<li id="menu-item-4833" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4833"><a href="https://www.guru99.com/python-tutorials.html">Python</a></li>
	<li id="menu-item-4835" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4835"><a href="https://www.guru99.com/vbscript-tutorials-for-beginners.html">VBScript</a></li>
	<li id="menu-item-4838" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4838"><a href="https://www.guru99.com/c-sharp-tutorial.html">C#</a></li>
	<li id="menu-item-4845" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4845"><a href="https://www.guru99.com/mariadb-tutorial-install.html">MariaDB</a></li>
	<li id="menu-item-4846" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4846"><a href="https://www.guru99.com/reactjs-tutorial.html">ReactJS</a></li>
	<li id="menu-item-4847" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4847"><a href="https://www.guru99.com/web-services-tutorial.html">Web Services</a></li>
	<li id="menu-item-4850" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4850"><a href="https://www.guru99.com/cpp-programming-tutorial.html">C++</a></li>
	<li id="menu-item-4852" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4852"><a href="https://www.guru99.com/ms-access-tutorial.html">MS Access</a></li>
	<li id="menu-item-4854" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4854"><a href="https://www.guru99.com/ruby-on-rails-tutorial.html">Ruby &#038; Rails</a></li>
	<li id="menu-item-4857" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4857"><a href="https://www.guru99.com/wpf-tutorial.html">WPF</a></li>
	<li id="menu-item-4863" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4863"><a href="https://www.guru99.com/codeigniter-tutorial.html">CodeIgniter</a></li>
	<li id="menu-item-4864" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4864"><a href="https://www.guru99.com/mysql-tutorial.html">MySQL</a></li>
	<li id="menu-item-4868" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4868"><a href="https://www.guru99.com/scala-tutorial.html">Scala</a></li>
	<li id="menu-item-4887" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4887"><a href="https://www.guru99.com/sqlite-tutorial.html">SQLite</a></li>
	<li id="menu-item-4871" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4871"><a href="https://www.guru99.com/dbms-tutorial.html">DBMS</a></li>
	<li id="menu-item-4875" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4875"><a href="https://www.guru99.com/node-js-tutorial.html">Node.js</a></li>
	<li id="menu-item-4877" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4877"><a href="https://www.guru99.com/sql.html">SQL</a></li>
	<li id="menu-item-4884" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4884"><a href="https://www.guru99.com/perl-tutorials.html">Perl</a></li>
	<li id="menu-item-4880" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4880"><a href="https://www.guru99.com/interactive-javascript-tutorials.html">JavaScript</a></li>
</ul>
</li>
<li id="menu-item-3176" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3176 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-3 kadence-menu-mega-layout-equal"><a href="/design-analysis-algorithms-tutorial.html"><span class="nav-drop-title-wrap">Must Learn<span class="dropdown-nav-toggle"><span class="kadence-svg-iconset svg-baseline"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></span></span></a>
<ul class="sub-menu">
	<li id="menu-item-4895" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4895"><a href="https://www.guru99.com/accounting.html">Accounting</a></li>
	<li id="menu-item-4897" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4897"><a href="https://www.guru99.com/embedded-systems-tutorial.html">Embedded Systems</a></li>
	<li id="menu-item-4980" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4980"><a href="https://www.guru99.com/os-tutorial.html">Operating System</a></li>
	<li id="menu-item-4906" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4906"><a href="https://www.guru99.com/design-analysis-algorithms-tutorial.html">Algorithms</a></li>
	<li id="menu-item-4909" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4909"><a href="https://www.guru99.com/ethical-hacking-tutorials.html">Ethical Hacking</a></li>
	<li id="menu-item-4911" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4911"><a href="https://www.guru99.com/pmp-tutorial.html">PMP</a></li>
	<li id="menu-item-4914" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4914"><a href="https://www.guru99.com/android-tutorial.html">Android</a></li>
	<li id="menu-item-4918" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4918"><a href="https://www.guru99.com/excel-tutorials.html">Excel Tutorial</a></li>
	<li id="menu-item-4919" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4919"><a href="https://www.guru99.com/photoshop-tutorials.html">Photoshop</a></li>
	<li id="menu-item-15774" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-15774"><a href="https://www.guru99.com/cryptocurrency-tutorial.html">Blockchain</a></li>
	<li id="menu-item-4923" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4923"><a href="https://www.guru99.com/google-go-tutorial.html">Go Programming</a></li>
	<li id="menu-item-4927" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4927"><a href="https://www.guru99.com/project-management-tutorial.html">Project Management</a></li>
	<li id="menu-item-4930" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4930"><a href="https://www.guru99.com/business-analyst-tutorial-course.html">Business Analyst</a></li>
	<li id="menu-item-4934" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4934"><a href="https://www.guru99.com/iot-tutorial.html">IoT</a></li>
	<li id="menu-item-4937" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4937"><a href="https://www.guru99.com/best-hard-disks.html">Reviews</a></li>
	<li id="menu-item-17034" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-17034"><a href="https://www.guru99.com/web-design-and-development-tutorial.html">Build Website</a></li>
	<li id="menu-item-4948" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4948"><a href="https://www.guru99.com/itil-framework-process.html">ITIL</a></li>
	<li id="menu-item-4951" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4951"><a href="https://www.guru99.com/salesforce-tutorial.html">Salesforce</a></li>
	<li id="menu-item-4953" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4953"><a href="https://www.guru99.com/cloud-computing-for-beginners.html">Cloud Computing</a></li>
	<li id="menu-item-4958" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4958"><a href="https://www.guru99.com/jenkins-tutorial.html">Jenkins</a></li>
	<li id="menu-item-4960" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4960"><a href="https://www.guru99.com/seo-tutorial.html">SEO</a></li>
	<li id="menu-item-4964" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4964"><a href="https://www.guru99.com/learn-cobol-programming-tutorial.html">COBOL</a></li>
	<li id="menu-item-4965" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4965"><a href="https://www.guru99.com/mis-tutorial.html">MIS</a></li>
	<li id="menu-item-4968" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4968"><a href="https://www.guru99.com/software-engineering-tutorial.html">Software Engineering</a></li>
	<li id="menu-item-10495" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-10495"><a href="https://www.guru99.com/compiler-tutorial.html">Compiler Design</a></li>
	<li id="menu-item-4976" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4976"><a href="https://www.guru99.com/anime-websites-watch-online-free.html">Movie</a></li>
	<li id="menu-item-4975" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4975"><a href="https://www.guru99.com/vba-tutorial.html">VBA</a></li>
	<li id="menu-item-16691" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-16691"><a href="https://www.guru99.com/online-courses.html">Courses</a></li>
	<li id="menu-item-4971" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4971"><a href="https://www.guru99.com/data-communication-computer-network-tutorial.html">Networking</a></li>
	<li id="menu-item-16961" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-16961"><a href="https://www.guru99.com/best-vpn.html">VPN</a></li>
</ul>
</li>
<li id="menu-item-3177" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3177 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-3 kadence-menu-mega-layout-equal"><a href="/bigdata-tutorials.html"><span class="nav-drop-title-wrap">Big Data<span class="dropdown-nav-toggle"><span class="kadence-svg-iconset svg-baseline"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></span></span></a>
<ul class="sub-menu">
	<li id="menu-item-4990" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4990"><a href="https://www.guru99.com/aws-tutorial.html">AWS</a></li>
	<li id="menu-item-4993" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4993"><a href="https://www.guru99.com/hive-tutorials.html">Hive</a></li>
	<li id="menu-item-4996" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4996"><a href="https://www.guru99.com/power-bi-tutorial.html">Power BI</a></li>
	<li id="menu-item-4997" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4997"><a href="https://www.guru99.com/bigdata-tutorials.html">Big Data</a></li>
	<li id="menu-item-5000" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5000"><a href="https://www.guru99.com/informatica-tutorials.html">Informatica</a></li>
	<li id="menu-item-5001" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5001"><a href="https://www.guru99.com/qlikview-tutorial.html">Qlikview</a></li>
	<li id="menu-item-5005" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5005"><a href="https://www.guru99.com/cassandra-tutorial.html">Cassandra</a></li>
	<li id="menu-item-5008" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5008"><a href="https://www.guru99.com/microstrategy-tutorial.html">MicroStrategy</a></li>
	<li id="menu-item-5011" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5011"><a href="https://www.guru99.com/tableau-tutorial.html">Tableau</a></li>
	<li id="menu-item-5016" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5016"><a href="https://www.guru99.com/cognos-tutorial.html">Cognos</a></li>
	<li id="menu-item-5018" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5018"><a href="https://www.guru99.com/mongodb-tutorials.html">MongoDB</a></li>
	<li id="menu-item-5020" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5020"><a href="https://www.guru99.com/talend-tutorial.html">Talend</a></li>
	<li id="menu-item-5024" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5024"><a href="https://www.guru99.com/data-warehousing-tutorial.html">Data Warehousing</a></li>
	<li id="menu-item-5030" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5030"><a href="https://www.guru99.com/apache-nifi-tutorial.html">NiFi</a></li>
	<li id="menu-item-5036" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5036"><a href="https://www.guru99.com/zookeeper-tutorial.html">ZooKeeper</a></li>
	<li id="menu-item-5039" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5039"><a href="https://www.guru99.com/devops-tutorial.html">DevOps</a></li>
	<li id="menu-item-5041" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5041"><a href="https://www.guru99.com/obiee-tutorial.html">OBIEE</a></li>
	<li id="menu-item-5049" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5049"><a href="https://www.guru99.com/pentaho-tutorial.html">Pentaho</a></li>
	<li id="menu-item-5046" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5046"><a href="https://www.guru99.com/hbase-tutorials.html">HBase</a></li>
</ul>
</li>
<li id="menu-item-3178" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3178 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-2 kadence-menu-mega-layout-equal"><a href="/live-testing-project.html"><span class="nav-drop-title-wrap">Live Project<span class="dropdown-nav-toggle"><span class="kadence-svg-iconset svg-baseline"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></span></span></a>
<ul class="sub-menu">
	<li id="menu-item-5064" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5064"><a href="https://www.guru99.com/live-agile-testing-project.html">Live Agile Testing</a></li>
	<li id="menu-item-5067" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5067"><a href="https://www.guru99.com/live-selenium-project.html">Live Selenium Project</a></li>
	<li id="menu-item-5070" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5070"><a href="https://www.guru99.com/live-interactive-exercise-hp-alm.html">Live HP ALM</a></li>
	<li id="menu-item-5076" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5076"><a href="https://www.guru99.com/live-ecommerce-project.html">Live Selenium 2</a></li>
	<li id="menu-item-5080" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5080"><a href="https://www.guru99.com/live-java-project.html">Live Java Project</a></li>
	<li id="menu-item-5082" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5082"><a href="https://www.guru99.com/live-penetration-testing-project.html">Live Security Testing</a></li>
	<li id="menu-item-5086" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5086"><a href="https://www.guru99.com/live-mobile-testing-project.html">Live Mobile Testing</a></li>
	<li id="menu-item-5089" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5089"><a href="https://www.guru99.com/live-testing-project.html">Live Testing Project</a></li>
	<li id="menu-item-5092" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5092"><a href="https://www.guru99.com/live-payment-gateway-project.html">Live Payment Gateway</a></li>
	<li id="menu-item-5093" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5093"><a href="https://www.guru99.com/live-insurance-testing-project.html">Live Testing 2</a></li>
	<li id="menu-item-5096" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5096"><a href="https://www.guru99.com/live-php-project-learn-complete-web-development-cycle.html">Live PHP  Project</a></li>
	<li id="menu-item-5100" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5100"><a href="https://www.guru99.com/live-telecom-project.html">Live Telecom</a></li>
	<li id="menu-item-5101" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5101"><a href="https://www.guru99.com/live-projects.html">Live Projects Hub</a></li>
	<li id="menu-item-5103" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5103"><a href="https://www.guru99.com/live-uft-testing.html">Live UFT/QTP Testing</a></li>
	<li id="menu-item-5108" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5108"><a href="https://www.guru99.com/live-python-project.html">Live Python Project</a></li>
	<li id="menu-item-5111" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5111"><a href="https://www.guru99.com/live-seo-project.html">Live SEO Project</a></li>
</ul>
</li>
<li id="menu-item-3179" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3179 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-2 kadence-menu-mega-layout-equal"><a href="/artificial-intelligence-tutorial.html"><span class="nav-drop-title-wrap">AI<span class="dropdown-nav-toggle"><span class="kadence-svg-iconset svg-baseline"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></span></span></a>
<ul class="sub-menu">
	<li id="menu-item-16679" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-16679"><a href="https://www.guru99.com/ai-tutorial.html">Artificial Intelligence</a></li>
	<li id="menu-item-5120" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5120"><a href="https://www.guru99.com/pytorch-tutorial.html">PyTorch</a></li>
	<li id="menu-item-16520" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-16520"><a href="https://www.guru99.com/data-science-tutorial-for-beginners.html">Data Science</a></li>
	<li id="menu-item-5124" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5124"><a href="https://www.guru99.com/r-tutorial.html">R Programming</a></li>
	<li id="menu-item-5127" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5127"><a href="https://www.guru99.com/keras-tutorial.html">Keras</a></li>
	<li id="menu-item-5128" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5128"><a href="https://www.guru99.com/tensorflow-tutorial.html">TensorFlow</a></li>
	<li id="menu-item-5129" class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5129"><a href="https://www.guru99.com/nltk-tutorial.html">NLTK</a></li>
</ul>
</li>
</ul>		</div>
	</nav><!-- #site-navigation -->
	</div><!-- data-section="primary_navigation" -->
<div class="site-header-item site-header-focus-item" data-section="kadence_customizer_header_search">
		<div class="search-toggle-open-container">
						<button class="search-toggle-open drawer-toggle search-toggle-style-default" aria-label="View Search Form" data-toggle-target="#search-drawer" data-toggle-body-class="showing-popup-drawer-from-full" aria-expanded="false" data-set-focus="#search-drawer .search-field"
					>
						<span class="search-toggle-icon"><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-search-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="26" height="28" viewBox="0 0 26 28"><title>Search</title><path d="M18 13c0-3.859-3.141-7-7-7s-7 3.141-7 7 3.141 7 7 7 7-3.141 7-7zM26 26c0 1.094-0.906 2-2 2-0.531 0-1.047-0.219-1.406-0.594l-5.359-5.344c-1.828 1.266-4.016 1.937-6.234 1.937-6.078 0-11-4.922-11-11s4.922-11 11-11 11 4.922 11 11c0 2.219-0.672 4.406-1.937 6.234l5.359 5.359c0.359 0.359 0.578 0.875 0.578 1.406z"></path>
				</svg></span></span>
		</button>
	</div>
	</div><!-- data-section="header_search" -->
					</div>
							</div>
		</div>
	</div>
</div>
				</div>
			</div>
					</div>
	</div>
	
<div id="mobile-header" class="site-mobile-header-wrap">
	<div class="site-header-inner-wrap">
		<div class="site-header-upper-wrap">
			<div class="site-header-upper-inner-wrap">
			<div class="site-main-header-wrap site-header-focus-item site-header-row-layout-standard site-header-row-tablet-layout-default site-header-row-mobile-layout-default ">
	<div class="site-header-row-container-inner">
		<div class="site-container">
			<div class="site-main-header-inner-wrap site-header-row site-header-row-has-sides site-header-row-no-center">
									<div class="site-header-main-section-left site-header-section site-header-section-left">
						<div class="site-header-item site-header-focus-item" data-section="title_tagline">
	<div class="site-branding mobile-site-branding branding-layout-standard branding-tablet-layout-inherit site-brand-logo-only branding-mobile-layout-inherit"><a class="brand has-logo-image" href="https://www.guru99.com/" rel="home" aria-label="Guru99"><img width="300" height="59" src="https://www.guru99.com/images/guru99-logo-v1.png" class="custom-logo" alt="Guru99" decoding="async" /></a></div></div><!-- data-section="title_tagline" -->
					</div>
																	<div class="site-header-main-section-right site-header-section site-header-section-right">
						<div class="site-header-item site-header-focus-item site-header-item-navgation-popup-toggle" data-section="kadence_customizer_mobile_trigger">
		<div class="mobile-toggle-open-container">
						<button id="mobile-toggle" class="menu-toggle-open drawer-toggle menu-toggle-style-default" aria-label="Open menu" data-toggle-target="#mobile-drawer" data-toggle-body-class="showing-popup-drawer-from-right" aria-expanded="false" data-set-focus=".menu-toggle-close"
					>
						<span class="menu-toggle-icon"><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-menu-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Toggle Menu</title><path d="M3 13h18c0.552 0 1-0.448 1-1s-0.448-1-1-1h-18c-0.552 0-1 0.448-1 1s0.448 1 1 1zM3 7h18c0.552 0 1-0.448 1-1s-0.448-1-1-1h-18c-0.552 0-1 0.448-1 1s0.448 1 1 1zM3 19h18c0.552 0 1-0.448 1-1s-0.448-1-1-1h-18c-0.552 0-1 0.448-1 1s0.448 1 1 1z"></path>
				</svg></span></span>
		</button>
	</div>
	</div><!-- data-section="mobile_trigger" -->
					</div>
							</div>
		</div>
	</div>
</div>
			</div>
		</div>
			</div>
</div>
</header><!-- #masthead -->

	<div id="inner-wrap" class="wrap hfeed kt-clear">
		<div id="primary" class="content-area">
	<div class="content-container site-container">
		<main id="main" class="site-main" role="main">
						<div class="content-wrap">
				<article id="post-2" class="entry content-bg single-entry post-2 post type-post status-publish format-standard has-post-thumbnail hentry category-softwaretesting tag-non-amp">
	<div class="entry-content-wrap">
		<header class="entry-header post-title title-align-left title-tablet-align-inherit title-mobile-align-inherit">
	<h1 class="entry-title">What is Software Testing? Definition</h1><div class="entry-meta entry-meta-divider-customicon">
	<div class="coauthors"> 
            By :  
            <a href="https://www.guru99.com/author/thomas" class="author-link" title="Posts by Thomas Hamilton"> 
                <img src="data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%2025%2025'%3E%3C/svg%3E" width="25" height="25" alt="Thomas Hamilton" class="avatar avatar-25 wp-user-avatar wp-user-avatar-25 alignnone photo" data-lazy-src="https://www.guru99.com/images/thomas-hamilton-120x120.jpg" /><noscript><img src="https://www.guru99.com/images/thomas-hamilton-120x120.jpg" width="25" height="25" alt="Thomas Hamilton" class="avatar avatar-25 wp-user-avatar wp-user-avatar-25 alignnone photo" /></noscript> 
                Thomas Hamilton 
            </a> 
            </div>					<span class="updated-on">
						<span class="kadence-svg-iconset"><svg class="kadence-svg-icon kadence-hours-alt-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Hours</title><path d="M23 12c0-3.037-1.232-5.789-3.222-7.778s-4.741-3.222-7.778-3.222-5.789 1.232-7.778 3.222-3.222 4.741-3.222 7.778 1.232 5.789 3.222 7.778 4.741 3.222 7.778 3.222 5.789-1.232 7.778-3.222 3.222-4.741 3.222-7.778zM21 12c0 2.486-1.006 4.734-2.636 6.364s-3.878 2.636-6.364 2.636-4.734-1.006-6.364-2.636-2.636-3.878-2.636-6.364 1.006-4.734 2.636-6.364 3.878-2.636 6.364-2.636 4.734 1.006 6.364 2.636 2.636 3.878 2.636 6.364zM11 6v6c0 0.389 0.222 0.727 0.553 0.894l4 2c0.494 0.247 1.095 0.047 1.342-0.447s0.047-1.095-0.447-1.342l-3.448-1.723v-5.382c0-0.552-0.448-1-1-1s-1 0.448-1 1z"></path>
				</svg></span><span class="meta-label">Updated</span><time class="entry-date published updated" datetime="2023-09-19T17:07:31+05:30" itemprop="dateModified">September 19, 2023</time>					</span>
					</div><!-- .entry-meta -->
</header><!-- .entry-header -->

<div class="entry-content single-content">
	<div class='code-block code-block-1' style='margin: 8px 0; clear: both;'>
<style>
 .float-ad-left {
  margin-right: 6px;
  width: 345px;
  min-height: 100px;
 }
 @media(min-width: 768px) {
  .float-ad-left {
   float: left;
   width: 345px;
   min-height: 280px;
  }
 }
</style>

<div data-freestar-ad="__336x280 __336x280" id="guru99_top_banner" class="float-ad-left">
 <script data-cfasync="false" type="text/javascript">
  freestar.config.enabled_slots.push({
   placementName: "guru99_top_banner",
   slotId: "guru99_top_banner"
  });
 </script>
</div></div>

<h2>Software Testing</h2>
<p><strong>Software Testing</strong> is a method to check whether the actual software product matches expected requirements and to ensure that software product is<a href="/defect-management-process.html" data-lasso-id="186522"> Defect </a>free. It involves execution of software/system components using manual or automated tools to evaluate one or more properties of interest. The purpose of software testing is to identify errors, gaps or missing requirements in contrast to actual requirements.</p>
<p>Some prefer saying Software testing definition as a <a href="/white-box-testing.html" data-lasso-id="186523">White Box</a> and <a href="/black-box-testing.html" data-lasso-id="186524">Black Box Testing</a>. In simple terms, Software Testing means the Verification of Application Under Test (AUT). This Software Testing course introduces testing software to the audience and justifies the importance of software testing.</p>

<nav class="wp-block-kadence-tableofcontents kb-table-of-content-nav kb-table-of-content-id_1b49da-72 kb-collapsible-toc kb-toc-toggle-hidden" role="navigation" aria-label="Table of Contents"><div class="kb-table-of-content-wrap"><div class="kb-table-of-contents-title-wrap kb-toggle-icon-style-arrowcircle"><span class="kb-table-of-contents-title"><strong>Table of Content:</strong></span><button class="kb-table-of-contents-icon-trigger kb-table-of-contents-toggle" aria-expanded="false" aria-label="Expand Table of Contents"></button></div><ul class="kb-table-of-content-list kb-table-of-content-list-columns-1 kb-table-of-content-list-style-none kb-table-of-content-link-style-plain"><li><a class="kb-table-of-contents__entry" href="#software-testing">Software Testing</a></li><li><a class="kb-table-of-contents__entry" href="#why-software-testing-is-important">Why Software Testing is Important?</a></li><li><a class="kb-table-of-contents__entry" href="#what-are-the-benefits-of-software-testing">What are the benefits of Software Testing?</a></li><li><a class="kb-table-of-contents__entry" href="#testing-in-software-engineering">Testing in Software Engineering</a></li><li><a class="kb-table-of-contents__entry" href="#types-of-software-testing">Types of Software Testing</a></li><li><a class="kb-table-of-contents__entry" href="#testing-strategies-in-software-engineering">Testing Strategies in Software Engineering</a></li><li><a class="kb-table-of-contents__entry" href="#program-testing">Program Testing</a></li><li><a class="kb-table-of-contents__entry" href="#summary-of-software-testing-basics">Summary of Software Testing Basics</a></li></ul></div></nav>

<div class='code-block code-block-2' style='margin: 8px 0; clear: both;'>
<!-- Tag ID: guru99_incontent_1 -->
<div align="center" data-freestar-ad="__728x250" id="guru99_incontent_1">
  <script data-cfasync="false" type="text/javascript">
    freestar.config.enabled_slots.push({ placementName: "guru99_incontent_1", slotId: "guru99_incontent_1" });
  </script>
</div></div>

<h2>Why Software Testing is Important?</h2>
<p><strong>Software Testing is Important</strong> because if there are any bugs or errors in the software, it can be identified early and can be solved before delivery of the software product. Properly tested software product ensures reliability, security and high performance which further results in time saving, cost effectiveness and customer satisfaction.</p>
<!-- [element-39837] -->
<style>
.featured-partners-container
{
background-color: #fcf9f1;
    padding: 8px 14px 0;
}

div.w3-container21.w3-half21 { 
	box-sizing: border-box;
	float: left;
	width: 100%;
	border: 1px solid #c3cce5;
	margin: 5px;
	border-radius: 5px;
	padding: 15px 10px;
	background: #fff;

}
div.w3-row21.w3-border21::after { 
	clear: both;
	content: "";
	display: table;
}
div.w3-row21.w3-border21::before { 
	clear: both;
	content: "";
	display: table;
}
@media (min-width: 766px) { 
	div.w3-container21.w3-half21 { 
		width: 32%;
	}
}
@media only screen and (min-width: 767px) and (max-width: 1023px){ 
	div.w3-container21.w3-half21 { 
		width: 31.5% !important;
	}
}

.button22 {
    background: #2f81ff !important;
    color: #fff!important;
    font-size: 18px;
    padding: 10px 50px;
    text-align: center;
    text-transform: none;
    white-space: nowrap;
	border-radius: 1000px;
}
@media only screen and (min-width: 767px) and (max-width: 1023px)  {.button22{ padding: 10px 15px !important; } }
@media only screen and (min-width: 1024px) and (max-width: 1149px)  {.button22{ padding: 10px 30px !important; } }

.card-number{

    color: #fff;
    letter-spacing: -.5px;
    min-width: 24px;
    height: 24px;
    border-radius: 50%;
    margin: 2px 0 0;
    background: #3d4a85;
    display: flex;
    align-items: center;
    justify-content: center;
    flex: 0 0 auto;
}
.btn-wrapper{
margin-top: 12px;
}
.get-estimate-cta{
min-height: 50px;
    padding: 10px;
    width: 100%;
}
.btn-wrapper .get-estimate-cta{
box-shadow: 0 4px 12px rgba(0,0,0,.2);
    border-radius: 10px;
    background: #35b782 !important;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all .2s ease;
    padding: 10px 8px;
    text-decoration: none;
	color: #fff;
}
.partners-stats .stats-wrapper{
	display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #c3cce6;
    padding: 10px 0;
}
.partners-stats .stats-wrapper .label {
    font-size: 13px;
    color: #333;
    padding-right: 10px;
    padding-bottom: 0;
    opacity: .6;
    margin: 0;
}
.partners-stats .stats-wrapper .value {
    width: 50%;
    font-size: 16px;
    color: #333;
    margin: 0;
    text-align: right;
}
@media(max-width: 766px){.last-widget-hide { display:none !important;}}
</style>

<div class="featured-partners-container">
<div class="w3-row21 w3-border21">                      

        <div class="w3-container21 w3-half21">  
		
		<span class="card-number" style="width: 11%;float:left;">1</span>
		<span class="card-title" style="height: 26px;margin-left: 10px;font-size: 22px;"><strong>Monday</strong></span>
		
        <p style="text-align:center;margin-top: 30px !important;margin-bottom: 0px !important;"><a href="https://guru99.live/FglzuJ" target="_blank" rel="sponsored noopener"><img decoding="async" src="data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%20160%20100'%3E%3C/svg%3E" width="160" height="100" data-lazy-src="https://www.guru99.com/images/monday-logo-v1.png"><noscript><img decoding="async" src="https://www.guru99.com/images/monday-logo-v1.png" width="160" height="100"></noscript></a></p>   
		
		
		<div class="btn-wrapper" style="height: 75px;"><a href="https://guru99.live/FglzuJ" class="get-estimate-cta fp-widget-cta-event" target="_blank" rel="sponsored noopener"><span>Learn More </span></a><p class="cta-subtext" style="font-size: 13px;text-align: center;">On Monday&#8217;s Website</p></div>
		
		<div class="partners-stats">
		<div class="stats-wrapper"><div class="tooltip-wrapper label">Time Tracking </div><span class="value">Yes</span></div>
		<div class="stats-wrapper"><div class="tooltip-wrapper label">Drag &amp; Drop </div><span class="value">Yes</span></div>
		<div class="stats-wrapper"><div class="tooltip-wrapper label">Free Trial </div><span class="value">Forever Free Plan</span></div>
		</div>
		
        </div>
                    

        <div class="w3-container21 w3-half21">  
		
		<span class="card-number" style="width: 11%;float:left;">2</span>
		<span class="card-title" style="height: 26px;margin-left: 10px;font-size: 22px;"><strong>JIRA Software</strong></span>
		
        <p style="text-align:center;margin-top: 21px !important;margin-bottom: 0px !important;"><a href="https://guru99.live/13Uel6" target="_blank" rel="sponsored noopener"><img decoding="async" src="data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%20225%20100'%3E%3C/svg%3E" width="225" height="100" data-lazy-src="https://www.guru99.com/images/jira-software-logo-v1.png"><noscript><img decoding="async" src="https://www.guru99.com/images/jira-software-logo-v1.png" width="225" height="100"></noscript></a></p>   
		
		
		<div class="btn-wrapper" style="height: 75px;"><a href="https://guru99.live/13Uel6" class="get-estimate-cta fp-widget-cta-event" target="_blank" rel="sponsored noopener"><span>Learn More </span></a><p class="cta-subtext" style="font-size: 13px;text-align: center;">On Jira Software Website</p></div>
		
		<div class="partners-stats">
		<div class="stats-wrapper"><div class="tooltip-wrapper label">Time Tracking </div><span class="value">Yes</span></div>
		<div class="stats-wrapper"><div class="tooltip-wrapper label">Drag &amp; Drop </div><span class="value">Yes</span></div>
		<div class="stats-wrapper"><div class="tooltip-wrapper label">Free Trial </div><span class="value">Forever Free Plan</span></div>
		</div>
		
        </div>
		
		<div class="w3-container21 w3-half21 last-widget-hide">  
		
		<span class="card-number" style="width: 11%;float:left;">3</span>
		<span class="card-title" style="height: 26px;margin-left: 10px;font-size: 22px;"><strong>Smartsheet</strong></span>
		
        <p style="text-align:center;margin-top: 25px !important;margin-bottom: 0px !important;"><a href="https://guru99.live/kWYQAI" target="_blank" rel="sponsored noopener"><img decoding="async" src="data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%20250%2045'%3E%3C/svg%3E" style="width: 90%;" width="250" height="45" data-lazy-src="https://www.guru99.com/images/smartsheet-logo-v3.png"><noscript><img decoding="async" src="https://www.guru99.com/images/smartsheet-logo-v3.png" style="width: 90%;" width="250" height="45"></noscript></a></p>   
		
		
		<div class="btn-wrapper" style="height: 75px;"><a href="https://guru99.live/kWYQAI" class="get-estimate-cta fp-widget-cta-event" target="_blank" rel="sponsored noopener"><span>Learn More </span></a><p class="cta-subtext" style="font-size: 13px;text-align: center;">On Smartsheet&#8217;s Website</p></div>
		
		<div class="partners-stats">
		<div class="stats-wrapper"><div class="tooltip-wrapper label">Time Tracking </div><span class="value">Yes</span></div>
		<div class="stats-wrapper"><div class="tooltip-wrapper label">Drag &amp; Drop </div><span class="value">Yes</span></div>
		<div class="stats-wrapper"><div class="tooltip-wrapper label">Free Trial </div><span class="value">Forever Free Plan</span></div>
		</div>
		
        </div>
   
 </div>
 </div>
<!-- [/element-39837] -->
<p><strong>What is the need of Testing?</strong></p>
<p>Testing is important because software bugs could be expensive or even dangerous. Software bugs can potentially cause monetary and human loss, and history is full of such examples.</p>
<ul>
<li>In April 2015, Bloomberg terminal in London crashed due to software glitch affected more than 300,000 traders on financial markets. It forced the government to postpone a 3bn pound debt sale.</li>
<li>Nissan cars recalled over 1 million cars from the market due to software failure in the airbag sensory detectors. There has been reported two accident due to this software failure.</li>
<li>Starbucks was forced to close about 60 percent of stores in the U.S and Canada due to software failure in its POS system. At one point, the store served coffee for free as they were unable to process the transaction.</li>
<li>Some of Amazon&#8217;s third-party retailers saw their product price is reduced to 1p due to a software glitch. They were left with heavy losses.</li>
<li>Vulnerability in Windows 10. This bug enables users to escape from security sandboxes through a flaw in the win32k system.</li>
<li>In 2015 fighter plane F-35 fell victim to a software bug, making it unable to detect targets correctly.</li>
<li>China Airlines Airbus A300 crashed due to a software bug on April 26, 1994, killing 264 innocents live</li>
<li>In 1985, Canada&#8217;s Therac-25 radiation therapy machine malfunctioned due to software bug and delivered lethal radiation doses to patients, leaving 3 people dead and critically injuring 3 others.</li>
<li>In April of 1999, a software bug caused the failure of a $1.2 billion military satellite launch, the costliest accident in history</li>
<li>In May of 1996, a software bug caused the bank accounts of 823 customers of a major U.S. bank to be credited with 920 million US dollars.</li>
</ul>
<p align="center"><iframe loading="lazy" src="about:blank" width="640" height="360" frameborder="0" allowfullscreen="allowfullscreen" data-rocket-lazyload="fitvidscompatible" data-lazy-src="//www.youtube.com/embed/TDynSmrzpXw"><p>Your browser does not support iframes. Please <a href="https://www.example.com/">click here</a> to view the content.</p>
    <div>Another block-level fallback message.</div>
</iframe><noscript><iframe src="//www.youtube.com/embed/TDynSmrzpXw" width="640" height="360" frameborder="0" allowfullscreen="allowfullscreen"></iframe></noscript></p>

<p align="center">Click <a href="/faq#faq1">here</a> if the video is not accessible <br> </p>

<h2>What are the benefits of Software Testing?</h2>
<p>Here are the benefits of using software testing:</p>
<ul>
<li><strong>Cost-Effective: </strong>It is one of the important advantages of software testing. Testing any IT project on time helps you to save your money for the long term. In case if the bugs caught in the earlier stage of software testing, it costs less to fix.</li>
<li><strong>Security: </strong>It is the most vulnerable and sensitive benefit of software testing. People are looking for trusted products. It helps in removing risks and problems earlier.</li>
<li><strong>Product quality: </strong>It is an essential requirement of any software product. Testing ensures a quality product is delivered to customers.</li>
<li><strong>Customer Satisfaction: </strong>The main aim of any product is to give satisfaction to their customers. UI/UX Testing ensures the best user experience.</li>
</ul>
<p><strong>» Also check:</strong> <a href="/software-testing-service-providers.html" data-lasso-id="300785">Best Software Testing Services Companies</a></p>
<div class='code-block code-block-3' style='margin: 8px 0; clear: both;'>
<!-- Tag ID: guru99_static_3 -->
<div data-freestar-ad="__300x250 __336x280" id="guru99_incontent_2">
 <script data-cfasync="false" type="text/javascript">
  freestar.config.enabled_slots.push({ placementName: "guru99_incontent_2", slotId: "guru99_incontent_2" });
 </script>
</div></div>

<h2>Testing in Software Engineering</h2>
<p>As per ANSI/IEEE 1059, <strong>Testing in Software Engineering</strong> is a process of evaluating a software product to find whether the current software product meets the required conditions or not. The testing process involves evaluating the features of the software product for requirements in terms of any missing requirements, bugs or errors, security, reliability and performance.</p>
<h2>Types of Software Testing</h2>
<p>Here are the software testing types:</p>
<p>Typically Testing is classified into three categories.</p>
<ul>
<li>Functional Testing</li>
<li>Non-Functional Testing or <a href="/performance-testing.html" data-lasso-id="186529">Performance Testing</a></li>
<li>Maintenance (Regression and Maintenance)</li>
</ul>
<figure style="text-align:center;"><a href="https://www.guru99.com/images/2/061920_1310_Whatissoftwaretesting1.png" data-lasso-id="512249"><img decoding="async" fetchpriority="high" width="628" height="275" src="data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%20628%20275'%3E%3C/svg%3E" alt="Types of Software Testing in Software Engineering" data-lazy-src="https://www.guru99.com/images/2/061920_1310_Whatissoftwaretesting1.png"><noscript><img decoding="async" fetchpriority="high" width="628" height="275" src="https://www.guru99.com/images/2/061920_1310_Whatissoftwaretesting1.png" alt="Types of Software Testing in Software Engineering"></noscript></a><figcaption style="text-align:center;">Types of Software Testing in Software Engineering</figcaption></figure>
<table class="table table-striped">
<tbody>
<tr>
<th>Testing Category</th>
<th>Types of Testing</th>
</tr>
<tr>
<td>Functional Testing</td>
<td>
<ul>
<li><a href="/unit-testing-guide.html" data-lasso-id="186530">Unit Testing</a></li>
<li><a href="/integration-testing.html" data-lasso-id="186531">Integration Testing</a></li>
<li>Smoke</li>
<li>UAT ( User Acceptance Testing)</li>
<li>Localization</li>
<li>Globalization</li>
<li>Interoperability</li>
<li>So on</li>
</ul>
</td>
</tr>
<tr>
<td>Non-Functional Testing</td>
<td>
<ul>
<li>Performance</li>
<li>Endurance</li>
<li>Load</li>
<li>Volume</li>
<li>Scalability</li>
<li>Usability</li>
<li>So on</li>
</ul>
</td>
</tr>
<tr>
<td>Maintenance</td>
<td>
<ul>
<li>Regression</li>
<li>Maintenance</li>
</ul>
</td>
</tr>
</tbody>
</table>
<p>This is not the complete list as there are more than <a href="/types-of-software-testing.html" data-lasso-id="186532">150 types of testing</a> types and still adding. Also, note that not all testing types are applicable to all projects but depend on the nature &amp; scope of the project. To explore a variety of testing tools and find the ones that suit your project requirements, visit this <a href="/testing-tools.html" data-lasso-id="454382">list of testing tools</a>.</p>
<h2>Testing Strategies in Software Engineering</h2>
<p>Here are important strategies in software engineering:</p>
<p><strong>Unit Testing: </strong> This software testing basic approach is followed by the programmer to test the unit of the program. It helps developers to know whether the individual unit of the code is working properly or not.</p>
<p><strong>Integration testing: </strong>It focuses on the construction and design of the software. You need to see that the integrated units are working without errors or not.</p>
<p><strong>System testing: </strong>In this method, your software is compiled as a whole and then tested as a whole. This testing strategy checks the functionality, security, portability, amongst others.</p>
<h2>Program Testing</h2>
<p><strong>Program Testing</strong> in software testing is a method of executing an actual software program with the aim of testing program behavior and finding errors. The software program is executed with test case data to analyse the program behavior or response to the test data. A good program testing is one which has high chances of finding bugs.</p>
<div class='code-block code-block-4' style='margin: 8px 0; clear: both;'>
<!-- Tag ID: guru99_static_4 -->
<div data-freestar-ad="__300x250 __336x280" id="guru99_incontent_3">
 <script data-cfasync="false" type="text/javascript">
  freestar.config.enabled_slots.push({ placementName: "guru99_incontent_3", slotId: "guru99_incontent_3" });
 </script>
</div></div>

<h2>Summary of Software Testing Basics</h2>
<ul>
<li>Define Software Testing: Software testing is defined as an activity to check whether the actual results match the expected results and to ensure that the software system is Defect free.</li>
<li>Testing is important because software bugs could be expensive or even dangerous.</li>
<li>The important reasons for using software testing are: cost-effective, security, product quality, and customer satisfaction.</li>
<li>Typically Testing is classified into three categories <a href="/functional-testing.html" data-lasso-id="186535">functional testing</a>, non-functional testing or performance testing, and maintenance.</li>
<li>The important strategies in software engineering are: unit testing, integration testing, validation testing, and system testing.</li>
</ul><div class='yarpp yarpp-related yarpp-related-website yarpp-template-list'>
<!-- YARPP List -->
<strong class="you-might-like">You Might Like:</strong><ul>
<li><a href="https://www.guru99.com/software-testing-seven-principles.html" rel="bookmark" title="7 Principles of Software Testing with Examples">7 Principles of Software Testing with Examples</a></li>
<li><a href="https://www.guru99.com/v-model-software-testing.html" rel="bookmark" title="V-Model in Software Testing">V-Model in Software Testing</a></li>
<li><a href="https://www.guru99.com/software-testing-life-cycle.html" rel="bookmark" title="STLC (Software Testing Life Cycle) Phases, Entry, Exit Criteria">STLC (Software Testing Life Cycle) Phases, Entry, Exit Criteria</a></li>
<li><a href="https://www.guru99.com/manual-testing.html" rel="bookmark" title="Manual Testing Tutorial: What is, Types, Concepts">Manual Testing Tutorial: What is, Types, Concepts</a></li>
<li><a href="https://www.guru99.com/automation-testing.html" rel="bookmark" title="What is Automation Testing? Test Tutorial">What is Automation Testing? Test Tutorial</a></li>
</ul>
</div>
<!-- AI CONTENT END 1 -->
</div><!-- .entry-content -->
	</div>
</article><!-- #post-2 -->


	<nav class="navigation post-navigation" role="navigation" aria-label="Posts">
    <span class="screen-reader-text">Post navigation</span>
    <div class="nav-links">
	<div class="nav-link-center"> <a href="https://form.jotform.me/72391811797466" target="_blank" style="color:#0556f3;font-size:15px;" rel="nofollow"> Report a Bug </a> </div>
	<div class="nav-previous"><a href="https://www.guru99.com/live-insurance-testing-project.html" rel="prev"><div><big class="pagenav"><span class="kadence-svg-iconset svg-baseline"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-left-alt-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="29" height="28" viewBox="0 0 29 28"><title>Previous</title><path d="M28 12.5v3c0 0.281-0.219 0.5-0.5 0.5h-19.5v3.5c0 0.203-0.109 0.375-0.297 0.453s-0.391 0.047-0.547-0.078l-6-5.469c-0.094-0.094-0.156-0.219-0.156-0.359v0c0-0.141 0.063-0.281 0.156-0.375l6-5.531c0.156-0.141 0.359-0.172 0.547-0.094 0.172 0.078 0.297 0.25 0.297 0.453v3.5h19.5c0.281 0 0.5 0.219 0.5 0.5z"></path>
				</svg></span>Prev</big></div></a></div><div class="nav-next"><a href="https://www.guru99.com/software-testing-career-complete-guide.html" rel="next"><div><big class="pagenav">Next<span class="kadence-svg-iconset svg-baseline"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-right-alt-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="27" height="28" viewBox="0 0 27 28"><title>Continue</title><path d="M27 13.953c0 0.141-0.063 0.281-0.156 0.375l-6 5.531c-0.156 0.141-0.359 0.172-0.547 0.094-0.172-0.078-0.297-0.25-0.297-0.453v-3.5h-19.5c-0.281 0-0.5-0.219-0.5-0.5v-3c0-0.281 0.219-0.5 0.5-0.5h19.5v-3.5c0-0.203 0.109-0.375 0.297-0.453s0.391-0.047 0.547 0.078l6 5.469c0.094 0.094 0.156 0.219 0.156 0.359v0z"></path>
				</svg></span></big></div></a></div>
	</div>
</nav>
	<div class='code-block code-block-5' style='margin: 8px 0; clear: both;'>
<!-- Tag ID: guru99_static_5 -->
<div data-freestar-ad="__300x250 __336x280" id="guru99_banner_near_footer">
 <script data-cfasync="false" type="text/javascript">
  freestar.config.enabled_slots.push({ placementName: "guru99_banner_near_footer", slotId: "guru99_banner_near_footer" });
 </script>
</div></div>
			</div>
					</main><!-- #main -->
		<aside id="secondary" role="complementary" class="primary-sidebar widget-area sidebar-slug-sidebar-primary sidebar-link-style-normal">
	<div class="sidebar-inner-wrap">
		<section id="block-4" class="widget widget_block">
<div class="wp-block-group is-layout-flow wp-block-group-is-layout-flow"><div class="wp-block-group__inner-container">
<!-- Tag ID: guru99_right_rail_1 -->
<div data-freestar-ad="__300x250 __336x280" style="margin-left: 10px;margin-top: 10px;margin-bottom: 10px;" id="guru99_right_rail_1">
 <script data-cfasync="false" type="text/javascript">
  freestar.config.enabled_slots.push({ placementName: "guru99_right_rail_1", slotId: "guru99_right_rail_1" });
 </script>
</div>

<!-- Tag ID: guru99_right_rail_2 -->
<div data-freestar-ad="__300x250 __336x280" style="margin-left: 10px;" id="guru99_right_rail_2">
 <script data-cfasync="false" type="text/javascript">
  freestar.config.enabled_slots.push({ placementName: "guru99_right_rail_2", slotId: "guru99_right_rail_2" });
 </script>
</div>
</div></div>
</section><section id="block-7" class="widget widget_block"><!-- Tag ID: guru99_right_rail_3 -->
<div data-freestar-ad="__300x250 __336x280" style="margin-left: 10px;" id="guru99_right_rail_3">
 <script data-cfasync="false" type="text/javascript">
  freestar.config.enabled_slots.push({ placementName: "guru99_right_rail_3", slotId: "guru99_right_rail_3" });
 </script>
</div></section>	</div>
</aside><!-- #secondary -->
	</div>
</div><!-- #primary -->
	</div><!-- #inner-wrap -->
	<footer id="colophon" class="site-footer" role="contentinfo">
	<div class="site-footer-wrap">
		<div class="site-top-footer-wrap site-footer-row-container site-footer-focus-item site-footer-row-layout-standard site-footer-row-tablet-layout-default site-footer-row-mobile-layout-default" data-section="kadence_customizer_footer_top">
	<div class="site-footer-row-container-inner">
				<div class="site-container">
			<div class="site-top-footer-inner-wrap site-footer-row site-footer-row-columns-1 site-footer-row-column-layout-row site-footer-row-tablet-column-layout-default site-footer-row-mobile-column-layout-row ft-ro-dir-column ft-ro-collapse-normal ft-ro-t-dir-default ft-ro-m-dir-default ft-ro-lstyle-plain">
									<div class="site-footer-top-section-1 site-footer-section footer-section-inner-items-1">
						<div class="footer-widget-area widget-area site-footer-focus-item footer-widget1 content-align-left content-tablet-align-default content-mobile-align-default content-valign-top content-tablet-valign-default content-mobile-valign-default" data-section="sidebar-widgets-footer1">
	<div class="footer-widget-area-inner site-info-inner">
		<section id="block-9" class="widget widget_block">

<style id='kadence-blocks-advancedheading-inline-css'>
.wp-block-kadence-advancedheading mark{background:transparent;border-style:solid;border-width:0}.wp-block-kadence-advancedheading mark.kt-highlight{color:#f76a0c;}.kb-adv-heading-icon{display: inline-flex;justify-content: center;align-items: center;}.single-content .kadence-advanced-heading-wrapper h1, .single-content .kadence-advanced-heading-wrapper h2, .single-content .kadence-advanced-heading-wrapper h3, .single-content .kadence-advanced-heading-wrapper h4, .single-content .kadence-advanced-heading-wrapper h5, .single-content .kadence-advanced-heading-wrapper h6 {margin: 1.5em 0 .5em;}.single-content .kadence-advanced-heading-wrapper+* { margin-top:0;}
</style>


<style>.kb-row-layout-idblock-9_950f4c-bc > .kt-row-column-wrap{align-content:start;}:where(.kb-row-layout-idblock-9_950f4c-bc > .kt-row-column-wrap) > .wp-block-kadence-column{justify-content:start;}.kb-row-layout-idblock-9_950f4c-bc > .kt-row-column-wrap{column-gap:var(--global-kb-gap-md, 2rem);row-gap:var(--global-kb-gap-md, 2rem);padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;grid-template-columns:minmax(0, 1fr);}.kb-row-layout-idblock-9_950f4c-bc > .kt-row-layout-overlay{opacity:1;background:linear-gradient(90deg, var(--global-palette1) 12%, var(--global-palette2) 100%);}.kb-row-layout-idblock-9_950f4c-bc .kt-row-layout-bottom-sep{height:120px;}.kb-row-layout-idblock-9_950f4c-bc .kt-row-layout-bottom-sep svg{width:100%;}.kb-row-layout-idblock-9_950f4c-bc .kt-row-layout-bottom-sep svg{fill:var(--global-palette9, #ffffff)!important;}@media all and (max-width: 1024px){.kb-row-layout-idblock-9_950f4c-bc > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}}@media all and (max-width: 1024px){.kb-row-layout-idblock-9_950f4c-bc .kt-row-layout-bottom-sep{height:px;}}@media all and (max-width: 1024px){.kb-row-layout-idblock-9_950f4c-bc .kt-row-layout-bottom-sep svg{width:%;}}@media all and (max-width: 767px){.kb-row-layout-idblock-9_950f4c-bc > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}.kb-row-layout-idblock-9_950f4c-bc .kt-row-layout-bottom-sep{height:px;}.kb-row-layout-idblock-9_950f4c-bc .kt-row-layout-bottom-sep svg{width:%;}}</style><div class="kb-row-layout-wrap kb-row-layout-idblock-9_950f4c-bc alignnone kt-row-has-bg wp-block-kadence-rowlayout"><div class="kt-row-layout-overlay kt-row-overlay-gradient"></div><div class="kt-row-column-wrap kt-has-1-columns kt-row-layout-equal kt-tab-layout-inherit kt-mobile-layout-row kt-row-valign-top">
<style>.kadence-column713b1a-e5 > .kt-inside-inner-col,.kadence-column713b1a-e5 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column713b1a-e5 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column713b1a-e5{position:relative;}</style>
<div class="wp-block-kadence-column kadence-column713b1a-e5 inner-column-1"><div class="kt-inside-inner-col">
<style>#kt-layout-id_f155c7-50 > .kt-row-column-wrap{align-content:start;}:where(#kt-layout-id_f155c7-50 > .kt-row-column-wrap) > .wp-block-kadence-column{justify-content:start;}#kt-layout-id_f155c7-50 > .kt-row-column-wrap{column-gap:var(--global-kb-gap-md, 2rem);row-gap:var(--global-kb-gap-md, 2rem);padding-top:var( --global-kb-row-default-top, 25px );padding-bottom:var( --global-kb-row-default-bottom, 25px );padding-top:0px;padding-bottom:0px;padding-left:0px;padding-right:0px;grid-template-columns:minmax(0, 1fr);}#kt-layout-id_f155c7-50 > .kt-row-layout-overlay{opacity:1;background-image:linear-gradient(90deg, var(--global-palette1, #3182CE) 12%, var(--global-palette2, #2B6CB0) 100%);}#kt-layout-id_f155c7-50 .kt-row-layout-bottom-sep{height:120px;}#kt-layout-id_f155c7-50 .kt-row-layout-bottom-sep svg{width:100%;}#kt-layout-id_f155c7-50 .kt-row-layout-bottom-sep svg{fill:var(--global-palette9, #ffffff)!important;}@media all and (max-width: 1024px){#kt-layout-id_f155c7-50 > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}}@media all and (max-width: 1024px){#kt-layout-id_f155c7-50 .kt-row-layout-bottom-sep{height:px;}}@media all and (max-width: 1024px){#kt-layout-id_f155c7-50 .kt-row-layout-bottom-sep svg{width:%;}}@media all and (max-width: 767px){#kt-layout-id_f155c7-50 > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}#kt-layout-id_f155c7-50 .kt-row-layout-bottom-sep{height:px;}#kt-layout-id_f155c7-50 .kt-row-layout-bottom-sep svg{width:%;}}</style>
<div class="wp-block-kadence-rowlayout alignnone"><div id="kt-layout-id_f155c7-50" class="kt-row-layout-inner kt-row-has-bg kt-layout-id_f155c7-50"><div class="kt-row-layout-overlay kt-row-overlay-gradient"></div><div class="kt-row-column-wrap kt-has-1-columns kt-gutter-default kt-v-gutter-default kt-row-valign-top kt-row-layout-equal kt-tab-layout-inherit kt-m-colapse-left-to-right kt-mobile-layout-row"><style>.kadence-column_4c580c-55 > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_4c580c-55 > .kt-inside-inner-col,.kadence-column_4c580c-55 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_4c580c-55 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_4c580c-55{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-1 kadence-column_4c580c-55"><div class="kt-inside-inner-col">
<style>#kt-layout-id_20363b-bd > .kt-row-column-wrap{align-content:start;}:where(#kt-layout-id_20363b-bd > .kt-row-column-wrap) > .wp-block-kadence-column{justify-content:start;}#kt-layout-id_20363b-bd > .kt-row-column-wrap{column-gap:var(--global-kb-gap-md, 2rem);row-gap:var(--global-kb-gap-md, 2rem);padding-top:var( --global-kb-row-default-top, 25px );padding-bottom:var( --global-kb-row-default-bottom, 25px );padding-top:0px;padding-bottom:0px;padding-left:0px;padding-right:0px;grid-template-columns:minmax(0, 1fr);}#kt-layout-id_20363b-bd > .kt-row-layout-overlay{opacity:1;background-image:linear-gradient(90deg, var(--global-palette1, #3182CE) 12%, var(--global-palette2, #2B6CB0) 100%);}#kt-layout-id_20363b-bd .kt-row-layout-bottom-sep{height:120px;}#kt-layout-id_20363b-bd .kt-row-layout-bottom-sep svg{width:100%;}#kt-layout-id_20363b-bd .kt-row-layout-bottom-sep svg{fill:var(--global-palette9, #ffffff)!important;}@media all and (max-width: 1024px){#kt-layout-id_20363b-bd > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}}@media all and (max-width: 1024px){#kt-layout-id_20363b-bd .kt-row-layout-bottom-sep{height:px;}}@media all and (max-width: 1024px){#kt-layout-id_20363b-bd .kt-row-layout-bottom-sep svg{width:%;}}@media all and (max-width: 767px){#kt-layout-id_20363b-bd > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}#kt-layout-id_20363b-bd .kt-row-layout-bottom-sep{height:px;}#kt-layout-id_20363b-bd .kt-row-layout-bottom-sep svg{width:%;}}</style>
<div class="wp-block-kadence-rowlayout alignnone"><div id="kt-layout-id_20363b-bd" class="kt-row-layout-inner kt-row-has-bg kt-layout-id_20363b-bd"><div class="kt-row-layout-overlay kt-row-overlay-gradient"></div><div class="kt-row-column-wrap kt-has-1-columns kt-gutter-default kt-v-gutter-default kt-row-valign-top kt-row-layout-equal kt-tab-layout-inherit kt-m-colapse-left-to-right kt-mobile-layout-row"><style>.kadence-column_1bd349-e8 > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_1bd349-e8 > .kt-inside-inner-col,.kadence-column_1bd349-e8 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_1bd349-e8 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_1bd349-e8{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-1 kadence-column_1bd349-e8"><div class="kt-inside-inner-col"></div></div>
</div></div></div>


<style>#kt-layout-id_d38890-a0 > .kt-row-column-wrap{z-index:10;position:relative;align-content:start;}:where(#kt-layout-id_d38890-a0 > .kt-row-column-wrap) > .wp-block-kadence-column{justify-content:start;}#kt-layout-id_d38890-a0 > .kt-row-column-wrap{column-gap:var(--global-kb-gap-md, 2rem);row-gap:var(--global-kb-gap-md, 2rem);max-width:var( --global-content-width, 1290px );padding-left:var(--global-content-edge-padding);padding-right:var(--global-content-edge-padding);padding-top:var( --global-kb-row-default-top, 25px );padding-bottom:var( --global-kb-row-default-bottom, 25px );padding-top:0px;padding-bottom:0px;padding-left:0px;padding-right:0px;grid-template-columns:minmax(0, 1fr);}#kt-layout-id_d38890-a0 > .kt-row-layout-overlay{opacity:0.30;}#kt-layout-id_d38890-a0 .kt-row-layout-bottom-sep{height:120px;}#kt-layout-id_d38890-a0 .kt-row-layout-bottom-sep svg{width:100%;}#kt-layout-id_d38890-a0 .kt-row-layout-bottom-sep svg{fill:#ffffff!important;}#kt-layout-id_d38890-a0 .kt-row-layout-top-sep{height:100px;}#kt-layout-id_d38890-a0 .kt-row-layout-top-sep svg{fill:var(--global-palette9, #ffffff)!important;}@media all and (max-width: 1024px){#kt-layout-id_d38890-a0 > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}}@media all and (max-width: 1024px){#kt-layout-id_d38890-a0 .kt-row-layout-bottom-sep{height:px;}}@media all and (max-width: 1024px){#kt-layout-id_d38890-a0 .kt-row-layout-bottom-sep svg{width:%;}}@media all and (max-width: 767px){#kt-layout-id_d38890-a0 > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}#kt-layout-id_d38890-a0 .kt-row-layout-bottom-sep{height:px;}#kt-layout-id_d38890-a0 .kt-row-layout-bottom-sep svg{width:%;}}</style>
<div class="wp-block-kadence-rowlayout alignnone"><div id="kt-layout-id_d38890-a0" class="kt-row-layout-inner kt-row-has-bg kt-layout-id_d38890-a0 has-theme-palette-8-background-color"><div class="kt-row-layout-top-sep kt-row-sep-type-mtns"><svg style="fill:var(--global-palette9)" viewbox="0 0 1000 100" preserveaspectratio="none"><path d="M1000,50l-182.69,-45.286l-292.031,61.197l-190.875,-41.075l-143.748,28.794l-190.656,-23.63l0,70l1000,0l0,-50Z" style="opacity:0.4"></path><path d="M1000,57l-152.781,-22.589l-214.383,19.81l-159.318,-21.471l-177.44,25.875l-192.722,5.627l-103.356,-27.275l0,63.023l1000,0l0,-43Z"></path></svg></div><div class="kt-row-column-wrap kt-has-1-columns kt-gutter-default kt-v-gutter-default kt-row-valign-top kt-row-layout-equal kt-tab-layout-inherit kt-m-colapse-left-to-right kt-mobile-layout-row kb-theme-content-width"><style>.kadence-column_d1271a-a3 > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_d1271a-a3 > .kt-inside-inner-col,.kadence-column_d1271a-a3 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_d1271a-a3 > .kt-inside-inner-col{background-color:#323a56;}.kadence-column_d1271a-a3 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_d1271a-a3{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-1 kadence-column_d1271a-a3"><div class="kt-inside-inner-col" style="background:#323a56"><style>.wp-block-kadence-advancedheading.kt-adv-heading_766798-d5, .wp-block-kadence-advancedheading.kt-adv-heading_766798-d5[data-kb-block="kb-adv-heading_766798-d5"]{text-align:center;font-size:26px;font-weight:normal;font-style:normal;}.wp-block-kadence-advancedheading.kt-adv-heading_766798-d5 mark, .wp-block-kadence-advancedheading.kt-adv-heading_766798-d5[data-kb-block="kb-adv-heading_766798-d5"] mark{font-weight:normal;font-style:normal;color:#f76a0c;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;}</style>
<p class="kt-adv-heading_766798-d5 hidden-phone wp-block-kadence-advancedheading has-white-color has-text-color" data-kb-block="kb-adv-heading_766798-d5" style="font-size: 26px;line-height: 40px;font-weight: 700;font-style: normal;margin: 1.5rem 0 0.75rem 0 !important;text-align:center;">Top Tutorials</p>


<style>#kt-layout-id_cd17bb-57{margin-top:40px;}#kt-layout-id_cd17bb-57 > .kt-row-column-wrap{align-content:start;}:where(#kt-layout-id_cd17bb-57 > .kt-row-column-wrap) > .wp-block-kadence-column{justify-content:start;}#kt-layout-id_cd17bb-57 > .kt-row-column-wrap{column-gap:var(--global-kb-gap-none, 0 );row-gap:var(--global-kb-gap-md, 2rem);padding-top:var( --global-kb-row-default-top, 25px );padding-bottom:var( --global-kb-row-default-bottom, 25px );padding-top:0px;padding-bottom:0px;padding-left:0px;padding-right:0px;grid-template-columns:repeat(4, minmax(0, 1fr));}#kt-layout-id_cd17bb-57 > .kt-row-layout-overlay{opacity:0.30;background-image:linear-gradient(180deg,  0%, #00B5E2 100%);}#kt-layout-id_cd17bb-57 .kt-row-layout-bottom-sep{height:100px;}#kt-layout-id_cd17bb-57 .kt-row-layout-bottom-sep svg{width:100%;}#kt-layout-id_cd17bb-57 .kt-row-layout-bottom-sep svg{fill:#ffffff!important;}@media all and (max-width: 1024px){#kt-layout-id_cd17bb-57 > .kt-row-column-wrap{grid-template-columns:repeat(4, minmax(0, 1fr));}}@media all and (max-width: 1024px){#kt-layout-id_cd17bb-57 .kt-row-layout-bottom-sep{height:px;}}@media all and (max-width: 1024px){#kt-layout-id_cd17bb-57 .kt-row-layout-bottom-sep svg{width:%;}}@media all and (max-width: 767px){#kt-layout-id_cd17bb-57 > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}#kt-layout-id_cd17bb-57 .kt-row-layout-bottom-sep{height:px;}#kt-layout-id_cd17bb-57 .kt-row-layout-bottom-sep svg{width:%;}}</style>
<div class="wp-block-kadence-rowlayout alignnone"><div id="kt-layout-id_cd17bb-57" class="kt-row-layout-inner kt-layout-id_cd17bb-57"><div class="kt-row-column-wrap kt-has-4-columns kt-gutter-none kt-v-gutter-default kt-row-valign-top kt-row-layout-equal kt-tab-layout-inherit kt-m-colapse-left-to-right kt-mobile-layout-row"><style>.wp-block-kadence-column.kadence-column_fa94df-79 > .kt-inside-inner-col{margin-left:25px;}.kadence-column_fa94df-79 > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_fa94df-79 > .kt-inside-inner-col,.kadence-column_fa94df-79 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_fa94df-79 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_fa94df-79{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-1 kadence-column_fa94df-79"><div class="kt-inside-inner-col">


<style>.kt-svg-icons_640e01-75 .kt-svg-item-0 .kb-svg-icon-wrap{color:var(--global-palette9, #ffffff);font-size:30px;}.kt-svg-icons_640e01-75 .kt-svg-item-0:hover .kb-svg-icon-wrap{color:var(--global-palette9, #ffffff);}.kt-svg-icons_640e01-75 .kt-svg-item-1 .kb-svg-icon-wrap{color:var(--global-palette9, #ffffff);font-size:30px;}.kt-svg-icons_640e01-75 .kt-svg-item-1:hover .kb-svg-icon-wrap{color:var(--global-palette9, #ffffff);}.kt-svg-icons_640e01-75 .kt-svg-item-2 .kb-svg-icon-wrap{color:var(--global-palette9, #ffffff);font-size:30px;}.kt-svg-icons_640e01-75 .kt-svg-item-2:hover .kb-svg-icon-wrap{color:var(--global-palette9, #ffffff);}.kt-svg-icons_640e01-75 .kt-svg-item-3 .kb-svg-icon-wrap{color:var(--global-palette9, #ffffff);font-size:30px;}.kt-svg-icons_640e01-75 .kt-svg-item-3:hover .kb-svg-icon-wrap{color:var(--global-palette9, #ffffff);}.kt-svg-icons_640e01-75 .kt-svg-item-4 .kb-svg-icon-wrap{color:var(--global-palette9, #ffffff);font-size:30px;}.kt-svg-icons_640e01-75 .kt-svg-item-4:hover .kb-svg-icon-wrap{color:var(--global-palette9, #ffffff);}.wp-block-kadence-icon.kt-svg-icons_640e01-75{justify-content:flex-start;}</style>
<div class="wp-block-kadence-icon kt-svg-icons kt-svg-icons_640e01-75 alignnone" style="text-align:left">

<div class="kt-svg-style-default kt-svg-icon-wrap kt-svg-item-0"><a href="https://www.facebook.com/Guru99Official" target="_blank" class="kt-svg-icon-link" style="margin-right:4px"><div style="display:inline-flex;justify-content:center;align-items:center;color:var(--global-palette9)" class="kt-svg-icon kt-svg-icon-fa_facebook-square"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 448 512" height="30" width="30" fill="currentColor" xmlns="http://www.w3.org/2000/svg" preserveaspectratio="xMinYMin meet" aria-hidden="true"><path d="M448 80v352c0 26.5-21.5 48-48 48h-85.3V302.8h60.6l8.7-67.6h-69.3V192c0-19.6 5.4-32.9 33.5-32.9H384V98.7c-6.2-.8-27.4-2.7-52.2-2.7-51.6 0-87 31.5-87 89.4v49.9H184v67.6h60.9V480H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48z"></path></svg></div></a></div>


<div class="kt-svg-style-default kt-svg-icon-wrap kt-svg-item-1"><a href="https://twitter.com/guru99com" target="_blank" class="kt-svg-icon-link" style="margin-right:4px"><div style="display:inline-flex;justify-content:center;align-items:center;color:var(--global-palette9)" class="kt-svg-icon kt-svg-icon-fa_twitter-square"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 448 512" height="30" width="30" fill="currentColor" xmlns="http://www.w3.org/2000/svg" preserveaspectratio="xMinYMin meet" aria-hidden="true"><path d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48zm-48.9 158.8c.2 2.8.2 5.7.2 8.5 0 86.7-66 186.6-186.6 186.6-37.2 0-71.7-10.8-100.7-29.4 5.3.6 10.4.8 15.8.8 30.7 0 58.9-10.4 81.4-28-28.8-.6-53-19.5-61.3-45.5 10.1 1.5 19.2 1.5 29.6-1.2-30-6.1-52.5-32.5-52.5-64.4v-.8c8.7 4.9 18.9 7.9 29.6 8.3a65.447 65.447 0 0 1-29.2-54.6c0-12.2 3.2-23.4 8.9-33.1 32.3 39.8 80.8 65.8 135.2 68.6-9.3-44.5 24-80.6 64-80.6 18.9 0 35.9 7.9 47.9 20.7 14.8-2.8 29-8.3 41.6-15.8-4.9 15.2-15.2 28-28.8 36.1 13.2-1.4 26-5.1 37.8-10.2-8.9 13.1-20.1 24.7-32.9 34z"></path></svg></div></a></div><div class="kt-svg-style-default kt-svg-icon-wrap kt-svg-item-2"><a href="https://www.linkedin.com/company/guru99/" target="_blank" class="kt-svg-icon-link" style="margin-right:4px"><div style="display:inline-flex;justify-content:center;align-items:center;color:var(--global-palette9)" class="kt-svg-icon kt-svg-icon-fa_linkedin"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 448 512" height="30" width="30" fill="currentColor" xmlns="http://www.w3.org/2000/svg" preserveaspectratio="xMinYMin meet" aria-hidden="true"><path d="M416 32H31.9C14.3 32 0 46.5 0 64.3v383.4C0 465.5 14.3 480 31.9 480H416c17.6 0 32-14.5 32-32.3V64.3c0-17.8-14.4-32.3-32-32.3zM135.4 416H69V202.2h66.5V416zm-33.2-243c-21.3 0-38.5-17.3-38.5-38.5S80.9 96 102.2 96c21.2 0 38.5 17.3 38.5 38.5 0 21.3-17.2 38.5-38.5 38.5zm282.1 243h-66.4V312c0-24.8-.5-56.7-34.5-56.7-34.6 0-39.9 27-39.9 54.9V416h-66.4V202.2h63.7v29.2h.9c8.9-16.8 30.6-34.5 62.9-34.5 67.2 0 79.7 44.3 79.7 101.9V416z"></path></svg></div></a></div><div class="kt-svg-style-default kt-svg-icon-wrap kt-svg-item-3"><a href="https://www.youtube.com/channel/UC19i1XD6k88KqHlET8atqFQ" target="_blank" class="kt-svg-icon-link" style="margin-right:4px"><div style="display:inline-flex;justify-content:center;align-items:center;color:var(--global-palette9)" class="kt-svg-icon kt-svg-icon-fa_youtube"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 576 512" height="30" width="30" fill="currentColor" xmlns="http://www.w3.org/2000/svg" preserveaspectratio="xMinYMin meet" aria-hidden="true"><path d="M549.655 124.083c-6.281-23.65-24.787-42.276-48.284-48.597C458.781 64 288 64 288 64S117.22 64 74.629 75.486c-23.497 6.322-42.003 24.947-48.284 48.597-11.412 42.867-11.412 132.305-11.412 132.305s0 89.438 11.412 132.305c6.281 23.65 24.787 41.5 48.284 47.821C117.22 448 288 448 288 448s170.78 0 213.371-11.486c23.497-6.321 42.003-24.171 48.284-47.821 11.412-42.867 11.412-132.305 11.412-132.305s0-89.438-11.412-132.305zm-317.51 213.508V175.185l142.739 81.205-142.739 81.201z"></path></svg></div></a></div><div class="kt-svg-style-default kt-svg-icon-wrap kt-svg-item-4"><a href="https://forms.aweber.com/form/46/724807646.htm" target="_blank" class="kt-svg-icon-link"><div style="display:inline-flex;justify-content:center;align-items:center;color:var(--global-palette9)" class="kt-svg-icon kt-svg-icon-fe_mail"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 24 24" height="30" width="30" fill="none" stroke="currentColor" xmlns="http://www.w3.org/2000/svg" stroke-width="1" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path><polyline points="22,6 12,13 2,6"></polyline></svg></div></a></div></div>

<span style="margin: 0px;padding-bottom:9px;padding-top:18px;color:#fff;font-size:22px;"><strong>About</strong></span><br>
<a href="/about-us" data-lasso-id="36133" style="color:#fff;">About Us</a><br>
<a href="/advertise-us" data-lasso-id="36134" style="color:#fff;">Advertise with Us</a><br>
<a href="/become-an-instructor" data-lasso-id="36135" style="color:#fff;">Write For Us</a><br>
<a href="/contact-us" data-lasso-id="36136" style="color:#fff;">Contact Us</a>
</div></div>


<style>.kadence-column_923a74-19 > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_923a74-19 > .kt-inside-inner-col,.kadence-column_923a74-19 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_923a74-19 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_923a74-19{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-2 kadence-column_923a74-19 hidden-phone"><div class="kt-inside-inner-col"><style>#kt-info-box_c022d7-f5 .kt-blocks-info-box-link-wrap{border-color:#323a56;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;background:#323a56;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;}#kt-info-box_c022d7-f5 .kt-blocks-info-box-link-wrap:hover{border-color:#323a56;background:#323a56;}#kt-info-box_c022d7-f5 .kt-info-svg-icon, #kt-info-box_c022d7-f5 .kt-info-svg-icon-flip, #kt-info-box_c022d7-f5 .kt-blocks-info-box-number{font-size:100px;}#kt-info-box_c022d7-f5 .kt-blocks-info-box-media{color:#000000;background:#ffe34c;border-color:#444444;border-radius:100px;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;padding-top:23px;padding-right:23px;padding-bottom:23px;padding-left:23px;}#kt-info-box_c022d7-f5 .kt-blocks-info-box-media-container{margin-top:0px;margin-right:0px;margin-bottom:0px;margin-left:0px;}#kt-info-box_c022d7-f5 .kt-blocks-info-box-media .kadence-info-box-image-intrisic img{border-radius:100px;}#kt-info-box_c022d7-f5 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-media{color:#000000;background:#ffe34c;border-color:#444444;}#kt-info-box_c022d7-f5 h5.kt-blocks-info-box-title{color:#ffffff;font-size:20px;font-family:'Source Sans Pro';font-style:normal;font-weight:700;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;margin-top:0px;margin-right:0px;margin-bottom:10px;margin-left:0px;}#kt-info-box_c022d7-f5 .kt-blocks-info-box-link-wrap:hover h5.kt-blocks-info-box-title{color:#ffffff;}#kt-info-box_c022d7-f5 .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);font-size:16px;font-style:normal;font-weight:normal;}#kt-info-box_c022d7-f5 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);}#kt-info-box_c022d7-f5 .kt-blocks-info-box-learnmore{background:transparent;border-color:#555555;border-width:0px 0px 0px 0px;padding-top:4px;padding-right:8px;padding-bottom:4px;padding-left:8px;margin:10px 0px 10px 0px;}#kt-info-box_c022d7-f5 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-learnmore{color:#ffffff;background:#444444;border-color:#444444;}</style>
<div id="kt-info-box_c022d7-f5" class="wp-block-kadence-infobox"><a class="kt-blocks-info-box-link-wrap info-box-link kt-blocks-info-box-media-align-top kt-info-halign-center" href="https://www.guru99.com/python-tutorials.html"><div class="kt-blocks-info-box-media-container"><div class="kt-blocks-info-box-media kt-info-media-animate-none"><div class="kadence-info-box-icon-container kt-info-icon-animate-none"><div class="kadence-info-box-icon-inner-container"><span style="display:block;justify-content:center;align-items:center" class="kt-info-svg-icon kt-info-svg-icon-icon-star-empty"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 1024 1024" height="100" width="100" fill="currentColor" xmlns="http://www.w3.org/2000/svg"><path d="M1024 397.050l-353.78-51.408-158.22-320.582-158.216 320.582-353.784 51.408 256 249.538-60.432 352.352 316.432-166.358 316.432 166.358-60.434-352.352 256.002-249.538zM512 753.498l-223.462 117.48 42.676-248.83-180.786-176.222 249.84-36.304 111.732-226.396 111.736 226.396 249.836 36.304-180.788 176.222 42.678 248.83-223.462-117.48z"></path></svg></span></div></div></div></div><div class="kt-infobox-textcontent" style="margin: 18px 0 8px 0 !important;"><span class="kt-blocks-info-box-title" style="font-size: 20px;line-height: 20px;font-weight: 700;font-style: normal;color: #fff;">Python</span></div></a></div>
</div></div>


<style>.kadence-column_06e0ae-f5 > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_06e0ae-f5 > .kt-inside-inner-col,.kadence-column_06e0ae-f5 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_06e0ae-f5 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_06e0ae-f5{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-3 kadence-column_06e0ae-f5 hidden-phone"><div class="kt-inside-inner-col"><style>#kt-info-box_b7d75a-43 .kt-blocks-info-box-link-wrap{border-color:#323a56;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;background:#323a56;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;}#kt-info-box_b7d75a-43 .kt-blocks-info-box-link-wrap:hover{border-color:#323a56;background:#323a56;}#kt-info-box_b7d75a-43 .kt-info-svg-icon, #kt-info-box_b7d75a-43 .kt-info-svg-icon-flip, #kt-info-box_b7d75a-43 .kt-blocks-info-box-number{font-size:100px;}#kt-info-box_b7d75a-43 .kt-blocks-info-box-media{color:#000000;background:#ff7f50;border-color:#444444;border-radius:100px;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;padding-top:23px;padding-right:23px;padding-bottom:23px;padding-left:23px;}#kt-info-box_b7d75a-43 .kt-blocks-info-box-media-container{margin-top:0px;margin-right:0px;margin-bottom:0px;margin-left:0px;}#kt-info-box_b7d75a-43 .kt-blocks-info-box-media .kadence-info-box-image-intrisic img{border-radius:100px;}#kt-info-box_b7d75a-43 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-media{color:#000000;background:#ff7f50;border-color:#444444;}#kt-info-box_b7d75a-43 h5.kt-blocks-info-box-title{color:#ffffff;font-size:20px;font-family:'Source Sans Pro';font-style:normal;font-weight:700;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;margin-top:0px;margin-right:0px;margin-bottom:10px;margin-left:0px;}#kt-info-box_b7d75a-43 .kt-blocks-info-box-link-wrap:hover h5.kt-blocks-info-box-title{color:#ffffff;}#kt-info-box_b7d75a-43 .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);font-size:16px;font-style:normal;font-weight:normal;}#kt-info-box_b7d75a-43 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);}#kt-info-box_b7d75a-43 .kt-blocks-info-box-learnmore{background:transparent;border-color:#555555;border-width:0px 0px 0px 0px;padding-top:4px;padding-right:8px;padding-bottom:4px;padding-left:8px;margin:10px 0px 10px 0px;}#kt-info-box_b7d75a-43 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-learnmore{color:#ffffff;background:#444444;border-color:#444444;}#kt-info-box_b7d75a-43 .kt-blocks-info-box-link-wrap{box-shadow:0px 0px 30px 0px rgba(0, 0, 0, 0.03);}#kt-info-box_b7d75a-43 .kt-blocks-info-box-link-wrap:hover{box-shadow:0px 0px 40px 0px rgba(0, 0, 0, 0.07);}</style>
<div id="kt-info-box_b7d75a-43" class="wp-block-kadence-infobox"><a class="kt-blocks-info-box-link-wrap info-box-link kt-blocks-info-box-media-align-top kt-info-halign-center" href="https://www.guru99.com/software-testing.html"><div class="kt-blocks-info-box-media-container"><div class="kt-blocks-info-box-media kt-info-media-animate-none"><div class="kadence-info-box-icon-container kt-info-icon-animate-none"><div class="kadence-info-box-icon-inner-container"><span style="display:block;justify-content:center;align-items:center" class="kt-info-svg-icon kt-info-svg-icon-icon-bug"><svg style="display:inline-block;vertical-align:middle" viewbox="-36 0 1024 951" height="100" width="100" fill="currentColor" xmlns="http://www.w3.org/2000/svg" preserveaspectratio="xMinYMin meet"><path d="M932.571 548.571c0 20-16.571 36.571-36.571 36.571h-128c0 71.429-15.429 125.143-38.286 165.714l118.857 119.429c14.286 14.286 14.286 37.143 0 51.429-6.857 7.429-16.571 10.857-25.714 10.857s-18.857-3.429-25.714-10.857l-113.143-112.571s-74.857 68.571-172 68.571v-512h-73.143v512c-103.429 0-178.857-75.429-178.857-75.429l-104.571 118.286c-7.429 8-17.143 12-27.429 12-8.571 0-17.143-2.857-24.571-9.143-14.857-13.714-16-36.571-2.857-52l115.429-129.714c-20-39.429-33.143-90.286-33.143-156.571h-128c-20 0-36.571-16.571-36.571-36.571s16.571-36.571 36.571-36.571h128v-168l-98.857-98.857c-14.286-14.286-14.286-37.143 0-51.429s37.143-14.286 51.429 0l98.857 98.857h482.286l98.857-98.857c14.286-14.286 37.143-14.286 51.429 0s14.286 37.143 0 51.429l-98.857 98.857v168h128c20 0 36.571 16.571 36.571 36.571zM658.286 219.429h-365.714c0-101.143 81.714-182.857 182.857-182.857s182.857 81.714 182.857 182.857z"></path></svg></span></div></div></div></div><div class="kt-infobox-textcontent" style="margin: 18px 0 8px 0 !important;"><span class="kt-blocks-info-box-title" style="font-size: 20px;line-height: 20px;font-weight: 700;font-style: normal;color: #fff;">Testing</span></div></a></div>
</div></div>


<style>.kadence-column_6232ca-f9 > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_6232ca-f9 > .kt-inside-inner-col,.kadence-column_6232ca-f9 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_6232ca-f9 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_6232ca-f9{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-4 kadence-column_6232ca-f9 hidden-phone"><div class="kt-inside-inner-col"><style>#kt-info-box_0d77ac-96 .kt-blocks-info-box-link-wrap{border-color:#323a56;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;background:#323a56;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;}#kt-info-box_0d77ac-96 .kt-blocks-info-box-link-wrap:hover{border-color:#323a56;background:#323a56;}#kt-info-box_0d77ac-96 .kt-info-svg-icon, #kt-info-box_0d77ac-96 .kt-info-svg-icon-flip, #kt-info-box_0d77ac-96 .kt-blocks-info-box-number{font-size:100px;}#kt-info-box_0d77ac-96 .kt-blocks-info-box-media{color:#000000;background:#b5eaaa;border-color:#444444;border-radius:100px;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;padding-top:23px;padding-right:23px;padding-bottom:23px;padding-left:23px;}#kt-info-box_0d77ac-96 .kt-blocks-info-box-media-container{margin-top:0px;margin-right:0px;margin-bottom:0px;margin-left:0px;}#kt-info-box_0d77ac-96 .kt-blocks-info-box-media .kadence-info-box-image-intrisic img{border-radius:100px;}#kt-info-box_0d77ac-96 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-media{color:#000000;background:#b5eaaa;border-color:#444444;}#kt-info-box_0d77ac-96 h5.kt-blocks-info-box-title{color:#ffffff;font-size:20px;font-family:'Source Sans Pro';font-style:normal;font-weight:700;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;margin-top:0px;margin-right:0px;margin-bottom:10px;margin-left:0px;}#kt-info-box_0d77ac-96 .kt-blocks-info-box-link-wrap:hover h5.kt-blocks-info-box-title{color:#ffffff;}#kt-info-box_0d77ac-96 .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);font-size:16px;font-style:normal;font-weight:normal;}#kt-info-box_0d77ac-96 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);}#kt-info-box_0d77ac-96 .kt-blocks-info-box-learnmore{background:transparent;border-color:#555555;border-width:0px 0px 0px 0px;padding-top:4px;padding-right:8px;padding-bottom:4px;padding-left:8px;margin:10px 0px 10px 0px;}#kt-info-box_0d77ac-96 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-learnmore{color:#ffffff;background:#444444;border-color:#444444;}#kt-info-box_0d77ac-96 .kt-blocks-info-box-link-wrap{box-shadow:0px 0px 30px 0px rgba(0, 0, 0, 0.03);}#kt-info-box_0d77ac-96 .kt-blocks-info-box-link-wrap:hover{box-shadow:0px 0px 40px 0px rgba(0, 0, 0, 0.07);}</style>
<div id="kt-info-box_0d77ac-96" class="wp-block-kadence-infobox"><a class="kt-blocks-info-box-link-wrap info-box-link kt-blocks-info-box-media-align-top kt-info-halign-center" href="https://www.guru99.com/ethical-hacking-tutorials.html"><div class="kt-blocks-info-box-media-container"><div class="kt-blocks-info-box-media kt-info-media-animate-none"><div class="kadence-info-box-icon-container kt-info-icon-animate-none"><div class="kadence-info-box-icon-inner-container"><span style="display:block;justify-content:center;align-items:center" class="kt-info-svg-icon kt-info-svg-icon-icon-laptop"><svg style="display:inline-block;vertical-align:middle" viewbox="0 -36 1024 1097" height="100" width="100" fill="currentColor" xmlns="http://www.w3.org/2000/svg" preserveaspectratio="xMinYMin meet"><path d="M237.714 731.429c-50.286 0-91.429-41.143-91.429-91.429v-402.286c0-50.286 41.143-91.429 91.429-91.429h621.714c50.286 0 91.429 41.143 91.429 91.429v402.286c0 50.286-41.143 91.429-91.429 91.429h-621.714zM219.429 237.714v402.286c0 9.714 8.571 18.286 18.286 18.286h621.714c9.714 0 18.286-8.571 18.286-18.286v-402.286c0-9.714-8.571-18.286-18.286-18.286h-621.714c-9.714 0-18.286 8.571-18.286 18.286zM1005.714 768h91.429v54.857c0 30.286-41.143 54.857-91.429 54.857h-914.286c-50.286 0-91.429-24.571-91.429-54.857v-54.857h1005.714zM594.286 822.857c5.143 0 9.143-4 9.143-9.143s-4-9.143-9.143-9.143h-91.429c-5.143 0-9.143 4-9.143 9.143s4 9.143 9.143 9.143h91.429z"></path></svg></span></div></div></div></div><div class="kt-infobox-textcontent" style="margin: 18px 0 8px 0 !important;"><span class="kt-blocks-info-box-title" style="font-size: 20px;line-height: 20px;font-weight: 700;font-style: normal;color: #fff;">Hacking</span></div></a></div>
</div></div>
</div></div></div>


<style>#kt-layout-id_5c1b91-c0{margin-top:40px;}#kt-layout-id_5c1b91-c0 > .kt-row-column-wrap{align-content:start;}:where(#kt-layout-id_5c1b91-c0 > .kt-row-column-wrap) > .wp-block-kadence-column{justify-content:start;}#kt-layout-id_5c1b91-c0 > .kt-row-column-wrap{column-gap:var(--global-kb-gap-none, 0 );row-gap:var(--global-kb-gap-md, 2rem);padding-top:var( --global-kb-row-default-top, 25px );padding-bottom:var( --global-kb-row-default-bottom, 25px );padding-top:0px;padding-bottom:0px;padding-left:0px;padding-right:0px;grid-template-columns:repeat(4, minmax(0, 1fr));}#kt-layout-id_5c1b91-c0 > .kt-row-layout-overlay{opacity:0.30;}@media all and (max-width: 1024px){#kt-layout-id_5c1b91-c0 > .kt-row-column-wrap{grid-template-columns:repeat(4, minmax(0, 1fr));}}@media all and (max-width: 767px){#kt-layout-id_5c1b91-c0 > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}}</style>
<div class="wp-block-kadence-rowlayout alignnone"><div id="kt-layout-id_5c1b91-c0" class="kt-row-layout-inner kt-layout-id_5c1b91-c0"><div class="kt-row-column-wrap kt-has-4-columns kt-gutter-none kt-v-gutter-default kt-row-valign-top kt-row-layout-equal kt-tab-layout-inherit kt-m-colapse-left-to-right kt-mobile-layout-row"><style>.wp-block-kadence-column.kadence-column_3ee1cc-0b > .kt-inside-inner-col{margin-left:25px;}.kadence-column_3ee1cc-0b > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_3ee1cc-0b > .kt-inside-inner-col,.kadence-column_3ee1cc-0b > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_3ee1cc-0b > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_3ee1cc-0b{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-1 kadence-column_3ee1cc-0b"><div class="kt-inside-inner-col">
<span style="margin: 0px;padding-bottom:9px;padding-top:18px;color:#fff;font-size:22px;"><strong>Career Suggestion</strong></span><br>
<a href="/best-sap-module.html" data-lasso-id="36137" style="color:#fff;">SAP Career Suggestion Tool</a><br>
<a href="/software-testing-career-complete-guide.html" data-lasso-id="36138" style="color:#fff;">Software Testing as a Career</a>
<br><br>
<span style="margin: 0px;padding-bottom:9px;padding-top:18px;color:#fff;font-size:22px;"><strong>Interesting</strong></span><br>
<a href="/ebook-pdf.html" data-lasso-id="36140" style="color:#fff;">eBook</a><br>
<a href="/blog" data-lasso-id="36141" style="color:#fff;">Blog</a><br>
<a href="/tests.html" data-lasso-id="36142" style="color:#fff;">Quiz</a><br>
<a href="/sap-ebook-pdf.html" data-lasso-id="36143" style="color:#fff;">SAP eBook</a><br>
</div></div>


<style>.kadence-column_7285ee-c5 > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_7285ee-c5 > .kt-inside-inner-col,.kadence-column_7285ee-c5 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_7285ee-c5 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_7285ee-c5{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-2 kadence-column_7285ee-c5 hidden-phone"><div class="kt-inside-inner-col"><style>#kt-info-box_7010b5-54 .kt-blocks-info-box-link-wrap{border-color:#323a56;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;background:#323a56;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;}#kt-info-box_7010b5-54 .kt-blocks-info-box-link-wrap:hover{border-color:#323a56;background:#323a56;}#kt-info-box_7010b5-54 .kt-info-svg-icon, #kt-info-box_7010b5-54 .kt-info-svg-icon-flip, #kt-info-box_7010b5-54 .kt-blocks-info-box-number{font-size:100px;}#kt-info-box_7010b5-54 .kt-blocks-info-box-media{color:#000000;background:#72a81a;border-color:#444444;border-radius:100px;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;padding-top:23px;padding-right:23px;padding-bottom:23px;padding-left:23px;}#kt-info-box_7010b5-54 .kt-blocks-info-box-media-container{margin-top:0px;margin-right:0px;margin-bottom:0px;margin-left:0px;}#kt-info-box_7010b5-54 .kt-blocks-info-box-media .kadence-info-box-image-intrisic img{border-radius:100px;}#kt-info-box_7010b5-54 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-media{color:#000000;background:#72a81a;border-color:#444444;}#kt-info-box_7010b5-54 h5.kt-blocks-info-box-title{color:#ffffff;font-size:20px;font-family:'Source Sans Pro';font-style:normal;font-weight:700;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;margin-top:0px;margin-right:0px;margin-bottom:10px;margin-left:0px;}#kt-info-box_7010b5-54 .kt-blocks-info-box-link-wrap:hover h5.kt-blocks-info-box-title{color:#ffffff;}#kt-info-box_7010b5-54 .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);font-size:16px;font-style:normal;font-weight:normal;}#kt-info-box_7010b5-54 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);}#kt-info-box_7010b5-54 .kt-blocks-info-box-learnmore{background:transparent;border-color:#555555;border-width:0px 0px 0px 0px;padding-top:4px;padding-right:8px;padding-bottom:4px;padding-left:8px;margin:10px 0px 10px 0px;}#kt-info-box_7010b5-54 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-learnmore{color:#ffffff;background:#444444;border-color:#444444;}</style>
<div id="kt-info-box_7010b5-54" class="wp-block-kadence-infobox"><a class="kt-blocks-info-box-link-wrap info-box-link kt-blocks-info-box-media-align-top kt-info-halign-center" href="https://www.guru99.com/sap-training-hub.html"><div class="kt-blocks-info-box-media-container"><div class="kt-blocks-info-box-media kt-info-media-animate-none"><div class="kadence-info-box-icon-container kt-info-icon-animate-none"><div class="kadence-info-box-icon-inner-container"><span style="display:block;justify-content:center;align-items:center" class="kt-info-svg-icon kt-info-svg-icon-fe_aperture"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 24 24" height="100" width="100" fill="none" stroke="currentColor" xmlns="http://www.w3.org/2000/svg" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="14.31" y1="8" x2="20.05" y2="17.94"></line><line x1="9.69" y1="8" x2="21.17" y2="8"></line><line x1="7.38" y1="12" x2="13.12" y2="2.06"></line><line x1="9.69" y1="16" x2="3.95" y2="6.06"></line><line x1="14.31" y1="16" x2="2.83" y2="16"></line><line x1="16.62" y1="12" x2="10.88" y2="21.94"></line></svg></span></div></div></div></div><div class="kt-infobox-textcontent" style="margin: 18px 0 8px 0 !important;"><span class="kt-blocks-info-box-title" style="font-size: 20px;line-height: 20px;font-weight: 700;font-style: normal;color: #fff;">SAP</span></div></a></div>
</div></div>


<style>.kadence-column_08932f-1f > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_08932f-1f > .kt-inside-inner-col,.kadence-column_08932f-1f > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_08932f-1f > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_08932f-1f{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-3 kadence-column_08932f-1f hidden-phone"><div class="kt-inside-inner-col"><style>#kt-info-box_28eb14-10 .kt-blocks-info-box-link-wrap{border-color:#323a56;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;background:#323a56;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;}#kt-info-box_28eb14-10 .kt-blocks-info-box-link-wrap:hover{border-color:#323a56;background:#323a56;}#kt-info-box_28eb14-10 .kt-info-svg-icon, #kt-info-box_28eb14-10 .kt-info-svg-icon-flip, #kt-info-box_28eb14-10 .kt-blocks-info-box-number{font-size:100px;}#kt-info-box_28eb14-10 .kt-blocks-info-box-media{color:#000000;background:#ed9600;border-color:#444444;border-radius:100px;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;padding-top:23px;padding-right:23px;padding-bottom:23px;padding-left:23px;}#kt-info-box_28eb14-10 .kt-blocks-info-box-media-container{margin-top:0px;margin-right:0px;margin-bottom:0px;margin-left:0px;}#kt-info-box_28eb14-10 .kt-blocks-info-box-media .kadence-info-box-image-intrisic img{border-radius:100px;}#kt-info-box_28eb14-10 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-media{color:#000000;background:#ed9600;border-color:#444444;}#kt-info-box_28eb14-10 h5.kt-blocks-info-box-title{color:#ffffff;font-size:20px;font-family:'Source Sans Pro';font-style:normal;font-weight:700;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;margin-top:0px;margin-right:0px;margin-bottom:10px;margin-left:0px;}#kt-info-box_28eb14-10 .kt-blocks-info-box-link-wrap:hover h5.kt-blocks-info-box-title{color:#ffffff;}#kt-info-box_28eb14-10 .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);font-size:16px;font-style:normal;font-weight:normal;}#kt-info-box_28eb14-10 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);}#kt-info-box_28eb14-10 .kt-blocks-info-box-learnmore{background:transparent;border-color:#555555;border-width:0px 0px 0px 0px;padding-top:4px;padding-right:8px;padding-bottom:4px;padding-left:8px;margin:10px 0px 10px 0px;}#kt-info-box_28eb14-10 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-learnmore{color:#ffffff;background:#444444;border-color:#444444;}#kt-info-box_28eb14-10 .kt-blocks-info-box-link-wrap{box-shadow:0px 0px 30px 0px rgba(0, 0, 0, 0.03);}#kt-info-box_28eb14-10 .kt-blocks-info-box-link-wrap:hover{box-shadow:0px 0px 40px 0px rgba(0, 0, 0, 0.07);}</style>
<div id="kt-info-box_28eb14-10" class="wp-block-kadence-infobox"><a class="kt-blocks-info-box-link-wrap info-box-link kt-blocks-info-box-media-align-top kt-info-halign-center" href="https://www.guru99.com/java-tutorial.html"><div class="kt-blocks-info-box-media-container"><div class="kt-blocks-info-box-media kt-info-media-animate-none"><div class="kadence-info-box-icon-container kt-info-icon-animate-none"><div class="kadence-info-box-icon-inner-container"><span style="display:block;justify-content:center;align-items:center" class="kt-info-svg-icon kt-info-svg-icon-fas_coffee"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 640 512" height="100" width="100" fill="currentColor" xmlns="http://www.w3.org/2000/svg"><path d="M192 384h192c53 0 96-43 96-96h32c70.6 0 128-57.4 128-128S582.6 32 512 32H120c-13.3 0-24 10.7-24 24v232c0 53 43 96 96 96zM512 96c35.3 0 64 28.7 64 64s-28.7 64-64 64h-32V96h32zm47.7 384H48.3c-47.6 0-61-64-36-64h583.3c25 0 11.8 64-35.9 64z"></path></svg></span></div></div></div></div><div class="kt-infobox-textcontent" style="margin: 18px 0 8px 0 !important;"><span class="kt-blocks-info-box-title" style="font-size: 20px;line-height: 20px;font-weight: 700;font-style: normal;color: #fff;">Java</span></div></a></div>
</div></div>


<style>.kadence-column_cf0881-7c > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_cf0881-7c > .kt-inside-inner-col,.kadence-column_cf0881-7c > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_cf0881-7c > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_cf0881-7c{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-4 kadence-column_cf0881-7c hidden-phone"><div class="kt-inside-inner-col"><style>#kt-info-box_fa331f-84 .kt-blocks-info-box-link-wrap{border-color:#323a56;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;background:#323a56;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;}#kt-info-box_fa331f-84 .kt-blocks-info-box-link-wrap:hover{border-color:#323a56;background:#323a56;}#kt-info-box_fa331f-84 .kt-info-svg-icon, #kt-info-box_fa331f-84 .kt-info-svg-icon-flip, #kt-info-box_fa331f-84 .kt-blocks-info-box-number{font-size:100px;}#kt-info-box_fa331f-84 .kt-blocks-info-box-media{color:#000000;background:#ffe87c;border-color:#444444;border-radius:100px;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;padding-top:23px;padding-right:23px;padding-bottom:23px;padding-left:23px;}#kt-info-box_fa331f-84 .kt-blocks-info-box-media-container{margin-top:0px;margin-right:0px;margin-bottom:0px;margin-left:0px;}#kt-info-box_fa331f-84 .kt-blocks-info-box-media .kadence-info-box-image-intrisic img{border-radius:100px;}#kt-info-box_fa331f-84 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-media{color:#000000;background:#ffe87c;border-color:#444444;}#kt-info-box_fa331f-84 h5.kt-blocks-info-box-title{color:#ffffff;font-size:20px;font-family:'Source Sans Pro';font-style:normal;font-weight:700;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;margin-top:0px;margin-right:0px;margin-bottom:10px;margin-left:0px;}#kt-info-box_fa331f-84 .kt-blocks-info-box-link-wrap:hover h5.kt-blocks-info-box-title{color:#ffffff;}#kt-info-box_fa331f-84 .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);font-size:16px;font-style:normal;font-weight:normal;}#kt-info-box_fa331f-84 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);}#kt-info-box_fa331f-84 .kt-blocks-info-box-learnmore{background:transparent;border-color:#555555;border-width:0px 0px 0px 0px;padding-top:4px;padding-right:8px;padding-bottom:4px;padding-left:8px;margin:10px 0px 10px 0px;}#kt-info-box_fa331f-84 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-learnmore{color:#ffffff;background:#444444;border-color:#444444;}#kt-info-box_fa331f-84 .kt-blocks-info-box-link-wrap{box-shadow:0px 0px 30px 0px rgba(0, 0, 0, 0.03);}#kt-info-box_fa331f-84 .kt-blocks-info-box-link-wrap:hover{box-shadow:0px 0px 40px 0px rgba(0, 0, 0, 0.07);}</style>
<div id="kt-info-box_fa331f-84" class="wp-block-kadence-infobox"><a class="kt-blocks-info-box-link-wrap info-box-link kt-blocks-info-box-media-align-top kt-info-halign-center" href="https://www.guru99.com/sql.html"><div class="kt-blocks-info-box-media-container"><div class="kt-blocks-info-box-media kt-info-media-animate-none"><div class="kadence-info-box-icon-container kt-info-icon-animate-none"><div class="kadence-info-box-icon-inner-container"><span style="display:block;justify-content:center;align-items:center" class="kt-info-svg-icon kt-info-svg-icon-fe_code"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 24 24" height="100" width="100" fill="none" stroke="currentColor" xmlns="http://www.w3.org/2000/svg" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="16 18 22 12 16 6"></polyline><polyline points="8 6 2 12 8 18"></polyline></svg></span></div></div></div></div><div class="kt-infobox-textcontent" style="margin: 18px 0 8px 0 !important;"><span class="kt-blocks-info-box-title" style="font-size: 20px;line-height: 20px;font-weight: 700;font-style: normal;color: #fff;">SQL</span></div></a></div>
</div></div>
</div></div></div>


<style>#kt-layout-id_a47002-e9{margin-top:40px;}#kt-layout-id_a47002-e9 > .kt-row-column-wrap{align-content:start;}:where(#kt-layout-id_a47002-e9 > .kt-row-column-wrap) > .wp-block-kadence-column{justify-content:start;}#kt-layout-id_a47002-e9 > .kt-row-column-wrap{column-gap:var(--global-kb-gap-none, 0 );row-gap:var(--global-kb-gap-md, 2rem);padding-top:var( --global-kb-row-default-top, 25px );padding-bottom:var( --global-kb-row-default-bottom, 25px );padding-top:0px;padding-bottom:0px;padding-left:0px;padding-right:0px;grid-template-columns:repeat(4, minmax(0, 1fr));}#kt-layout-id_a47002-e9 > .kt-row-layout-overlay{opacity:0.30;}@media all and (max-width: 1024px){#kt-layout-id_a47002-e9 > .kt-row-column-wrap{grid-template-columns:repeat(4, minmax(0, 1fr));}}@media all and (max-width: 767px){#kt-layout-id_a47002-e9 > .kt-row-column-wrap{grid-template-columns:minmax(0, 1fr);}}</style>
<div class="wp-block-kadence-rowlayout alignnone"><div id="kt-layout-id_a47002-e9" class="kt-row-layout-inner kt-layout-id_a47002-e9"><div class="kt-row-column-wrap kt-has-4-columns kt-gutter-none kt-v-gutter-default kt-row-valign-top kt-row-layout-equal kt-tab-layout-inherit kt-m-colapse-left-to-right kt-mobile-layout-row"><style>.wp-block-kadence-column.kadence-column_7d5a66-f9 > .kt-inside-inner-col{margin-left:25px;}.kadence-column_7d5a66-f9 > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_7d5a66-f9 > .kt-inside-inner-col,.kadence-column_7d5a66-f9 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_7d5a66-f9 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_7d5a66-f9{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-1 kadence-column_7d5a66-f9"><div class="kt-inside-inner-col">
<span style="margin: 0px;padding-bottom:9px;padding-top:18px;color:#fff;font-size:22px;"><strong>Execute online</strong></span><br>
<a href="/try-java-editor.html" data-lasso-id="36146" style="color:#fff;">Execute Java Online</a><br> 
<a href="/execute-javascript-online.html" data-lasso-id="36147" style="color:#fff;">Execute Javascript</a><br> 
<a href="/execute-html-online.html" data-lasso-id="36148" style="color:#fff;">Execute HTML</a><br> 
<a href="/execute-python-online.html" data-lasso-id="36149" style="color:#fff;">Execute Python</a><br>
</div></div>


<style>.kadence-column_f38d98-3a > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_f38d98-3a > .kt-inside-inner-col,.kadence-column_f38d98-3a > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_f38d98-3a > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_f38d98-3a{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-2 kadence-column_f38d98-3a hidden-phone"><div class="kt-inside-inner-col"><style>#kt-info-box_eaeb48-11 .kt-blocks-info-box-link-wrap{border-color:#323a56;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;background:#323a56;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;}#kt-info-box_eaeb48-11 .kt-blocks-info-box-link-wrap:hover{border-color:#323a56;background:#323a56;}#kt-info-box_eaeb48-11 .kt-info-svg-icon, #kt-info-box_eaeb48-11 .kt-info-svg-icon-flip, #kt-info-box_eaeb48-11 .kt-blocks-info-box-number{font-size:100px;}#kt-info-box_eaeb48-11 .kt-blocks-info-box-media{color:#000000;background:#2282e8;border-color:#444444;border-radius:100px;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;padding-top:23px;padding-right:23px;padding-bottom:23px;padding-left:23px;}#kt-info-box_eaeb48-11 .kt-blocks-info-box-media-container{margin-top:0px;margin-right:0px;margin-bottom:0px;margin-left:0px;}#kt-info-box_eaeb48-11 .kt-blocks-info-box-media .kadence-info-box-image-intrisic img{border-radius:100px;}#kt-info-box_eaeb48-11 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-media{color:#000000;background:#2282e8;border-color:#444444;}#kt-info-box_eaeb48-11 h5.kt-blocks-info-box-title{color:#ffffff;font-size:20px;font-family:'Source Sans Pro';font-style:normal;font-weight:700;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;margin-top:0px;margin-right:0px;margin-bottom:10px;margin-left:0px;}#kt-info-box_eaeb48-11 .kt-blocks-info-box-link-wrap:hover h5.kt-blocks-info-box-title{color:#ffffff;}#kt-info-box_eaeb48-11 .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);font-size:16px;font-style:normal;font-weight:normal;}#kt-info-box_eaeb48-11 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);}#kt-info-box_eaeb48-11 .kt-blocks-info-box-learnmore{background:transparent;border-color:#555555;border-width:0px 0px 0px 0px;padding-top:4px;padding-right:8px;padding-bottom:4px;padding-left:8px;margin:10px 0px 10px 0px;}#kt-info-box_eaeb48-11 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-learnmore{color:#ffffff;background:#444444;border-color:#444444;}</style>
<div id="kt-info-box_eaeb48-11" class="wp-block-kadence-infobox"><a class="kt-blocks-info-box-link-wrap info-box-link kt-blocks-info-box-media-align-top kt-info-halign-center" href="https://www.guru99.com/selenium-tutorial.html"><div class="kt-blocks-info-box-media-container"><div class="kt-blocks-info-box-media kt-info-media-animate-none"><div class="kadence-info-box-icon-container kt-info-icon-animate-none"><div class="kadence-info-box-icon-inner-container"><span style="display:block;justify-content:center;align-items:center" class="kt-info-svg-icon kt-info-svg-icon-fas_book-open"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 576 512" height="100" width="100" fill="currentColor" xmlns="http://www.w3.org/2000/svg"><path d="M542.22 32.05c-54.8 3.11-163.72 14.43-230.96 55.59-4.64 2.84-7.27 7.89-7.27 13.17v363.87c0 11.55 12.63 18.85 23.28 13.49 69.18-34.82 169.23-44.32 218.7-46.92 16.89-.89 30.02-14.43 30.02-30.66V62.75c.01-17.71-15.35-31.74-33.77-30.7zM264.73 87.64C197.5 46.48 88.58 35.17 33.78 32.05 15.36 31.01 0 45.04 0 62.75V400.6c0 16.24 13.13 29.78 30.02 30.66 49.49 2.6 149.59 12.11 218.77 46.95 10.62 5.35 23.21-1.94 23.21-13.46V100.63c0-5.29-2.62-10.14-7.27-12.99z"></path></svg></span></div></div></div></div><div class="kt-infobox-textcontent" style="margin: 18px 0 8px 0 !important;"><span class="kt-blocks-info-box-title" style="font-size: 20px;line-height: 20px;font-weight: 700;font-style: normal;color: #fff;">Selenium</span></div></a></div>
</div></div>


<style>.kadence-column_98252f-54 > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_98252f-54 > .kt-inside-inner-col,.kadence-column_98252f-54 > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_98252f-54 > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_98252f-54{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-3 kadence-column_98252f-54 hidden-phone"><div class="kt-inside-inner-col"><style>#kt-info-box_8d6b07-7e .kt-blocks-info-box-link-wrap{border-color:#323a56;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;background:#323a56;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;}#kt-info-box_8d6b07-7e .kt-blocks-info-box-link-wrap:hover{border-color:#323a56;background:#323a56;}#kt-info-box_8d6b07-7e .kt-info-svg-icon, #kt-info-box_8d6b07-7e .kt-info-svg-icon-flip, #kt-info-box_8d6b07-7e .kt-blocks-info-box-number{font-size:100px;}#kt-info-box_8d6b07-7e .kt-blocks-info-box-media{color:#000000;background:#bba7db;border-color:#444444;border-radius:100px;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;padding-top:23px;padding-right:23px;padding-bottom:23px;padding-left:23px;}#kt-info-box_8d6b07-7e .kt-blocks-info-box-media-container{margin-top:0px;margin-right:0px;margin-bottom:0px;margin-left:0px;}#kt-info-box_8d6b07-7e .kt-blocks-info-box-media .kadence-info-box-image-intrisic img{border-radius:100px;}#kt-info-box_8d6b07-7e .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-media{color:#000000;background:#bba7db;border-color:#444444;}#kt-info-box_8d6b07-7e h5.kt-blocks-info-box-title{color:#ffffff;font-size:20px;font-family:'Source Sans Pro';font-style:normal;font-weight:700;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;margin-top:0px;margin-right:0px;margin-bottom:10px;margin-left:0px;}#kt-info-box_8d6b07-7e .kt-blocks-info-box-link-wrap:hover h5.kt-blocks-info-box-title{color:#ffffff;}#kt-info-box_8d6b07-7e .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);font-size:16px;font-style:normal;font-weight:normal;}#kt-info-box_8d6b07-7e .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);}#kt-info-box_8d6b07-7e .kt-blocks-info-box-learnmore{background:transparent;border-color:#555555;border-width:0px 0px 0px 0px;padding-top:4px;padding-right:8px;padding-bottom:4px;padding-left:8px;margin:10px 0px 10px 0px;}#kt-info-box_8d6b07-7e .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-learnmore{color:#ffffff;background:#444444;border-color:#444444;}#kt-info-box_8d6b07-7e .kt-blocks-info-box-link-wrap{box-shadow:0px 0px 30px 0px rgba(0, 0, 0, 0.03);}#kt-info-box_8d6b07-7e .kt-blocks-info-box-link-wrap:hover{box-shadow:0px 0px 40px 0px rgba(0, 0, 0, 0.07);}</style>
<div id="kt-info-box_8d6b07-7e" class="wp-block-kadence-infobox"><a class="kt-blocks-info-box-link-wrap info-box-link kt-blocks-info-box-media-align-top kt-info-halign-center" href="https://www.guru99.com/web-design-and-development-tutorial.html"><div class="kt-blocks-info-box-media-container"><div class="kt-blocks-info-box-media kt-info-media-animate-none"><div class="kadence-info-box-icon-container kt-info-icon-animate-none"><div class="kadence-info-box-icon-inner-container"><span style="display:block;justify-content:center;align-items:center" class="kt-info-svg-icon kt-info-svg-icon-fe_briefcase"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 24 24" height="100" width="100" fill="none" stroke="currentColor" xmlns="http://www.w3.org/2000/svg" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="7" width="20" height="14" rx="2" ry="2"></rect><path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"></path></svg></span></div></div></div></div><div class="kt-infobox-textcontent" style="margin: 18px 0 8px 0 !important;"><span class="kt-blocks-info-box-title" style="font-size: 20px;line-height: 20px;font-weight: 700;font-style: normal;color: #fff;">Build Website</span></div></a></div>
</div></div>


<style>.kadence-column_b8bdb5-8e > .kt-inside-inner-col{border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;}.kadence-column_b8bdb5-8e > .kt-inside-inner-col,.kadence-column_b8bdb5-8e > .kt-inside-inner-col:before{border-top-left-radius:0px;border-top-right-radius:0px;border-bottom-right-radius:0px;border-bottom-left-radius:0px;}.kadence-column_b8bdb5-8e > .kt-inside-inner-col:before{opacity:0.3;}.kadence-column_b8bdb5-8e{position:relative;}</style>
<div class="wp-block-kadence-column inner-column-4 kadence-column_b8bdb5-8e hidden-phone"><div class="kt-inside-inner-col"><style>#kt-info-box_98787e-43 .kt-blocks-info-box-link-wrap{border-color:#323a56;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;background:#323a56;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;}#kt-info-box_98787e-43 .kt-blocks-info-box-link-wrap:hover{border-color:#323a56;background:#323a56;}#kt-info-box_98787e-43 .kt-info-svg-icon, #kt-info-box_98787e-43 .kt-info-svg-icon-flip, #kt-info-box_98787e-43 .kt-blocks-info-box-number{font-size:100px;}#kt-info-box_98787e-43 .kt-blocks-info-box-media{color:#000000;background:#c36241;border-color:#444444;border-radius:100px;border-top-width:0px;border-right-width:0px;border-bottom-width:0px;border-left-width:0px;padding-top:23px;padding-right:23px;padding-bottom:23px;padding-left:23px;}#kt-info-box_98787e-43 .kt-blocks-info-box-media-container{margin-top:0px;margin-right:0px;margin-bottom:0px;margin-left:0px;}#kt-info-box_98787e-43 .kt-blocks-info-box-media .kadence-info-box-image-intrisic img{border-radius:100px;}#kt-info-box_98787e-43 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-media{color:#000000;background:#c36241;border-color:#444444;}#kt-info-box_98787e-43 h5.kt-blocks-info-box-title{color:#ffffff;font-size:20px;font-family:'Source Sans Pro';font-style:normal;font-weight:700;padding-top:0px;padding-right:0px;padding-bottom:0px;padding-left:0px;margin-top:0px;margin-right:0px;margin-bottom:10px;margin-left:0px;}#kt-info-box_98787e-43 .kt-blocks-info-box-link-wrap:hover h5.kt-blocks-info-box-title{color:#ffffff;}#kt-info-box_98787e-43 .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);font-size:16px;font-style:normal;font-weight:normal;}#kt-info-box_98787e-43 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-text{color:var(--global-palette5, #4A5568);}#kt-info-box_98787e-43 .kt-blocks-info-box-learnmore{background:transparent;border-color:#555555;border-width:0px 0px 0px 0px;padding-top:4px;padding-right:8px;padding-bottom:4px;padding-left:8px;margin:10px 0px 10px 0px;}#kt-info-box_98787e-43 .kt-blocks-info-box-link-wrap:hover .kt-blocks-info-box-learnmore{color:#ffffff;background:#444444;border-color:#444444;}#kt-info-box_98787e-43 .kt-blocks-info-box-link-wrap{box-shadow:0px 0px 30px 0px rgba(0, 0, 0, 0.03);}#kt-info-box_98787e-43 .kt-blocks-info-box-link-wrap:hover{box-shadow:0px 0px 40px 0px rgba(0, 0, 0, 0.07);}</style>
<div id="kt-info-box_98787e-43" class="wp-block-kadence-infobox"><a class="kt-blocks-info-box-link-wrap info-box-link kt-blocks-info-box-media-align-top kt-info-halign-center" href="https://www.guru99.com/best-vpn.html"><div class="kt-blocks-info-box-media-container"><div class="kt-blocks-info-box-media kt-info-media-animate-none"><div class="kadence-info-box-icon-container kt-info-icon-animate-none"><div class="kadence-info-box-icon-inner-container"><span style="display:block;justify-content:center;align-items:center" class="kt-info-svg-icon kt-info-svg-icon-fas_cloud"><svg style="display:inline-block;vertical-align:middle" viewbox="0 0 640 512" height="100" width="100" fill="currentColor" xmlns="http://www.w3.org/2000/svg"><path d="M537.6 226.6c4.1-10.7 6.4-22.4 6.4-34.6 0-53-43-96-96-96-19.7 0-38.1 6-53.3 16.2C367 64.2 315.3 32 256 32c-88.4 0-160 71.6-160 160 0 2.7.1 5.4.2 8.1C40.2 219.8 0 273.2 0 336c0 79.5 64.5 144 144 144h368c70.7 0 128-57.3 128-128 0-61.9-44-113.6-102.4-125.4z"></path></svg></span></div></div></div></div><div class="kt-infobox-textcontent" style="margin: 18px 0 8px 0 !important;"><span class="kt-blocks-info-box-title" style="font-size: 20px;line-height: 20px;font-weight: 700;font-style: normal;color: #fff;">VPNs</span></div></a></div>
</div></div>
</div></div></div>



<p></p>
</div></div>
</div></div></div>

</div></div>
</div></div></div>

</div></div>

</div></div></section>	</div>
</div><!-- .footer-widget1 -->
					</div>
								</div>
		</div>
	</div>
</div>
<div class="site-middle-footer-wrap site-footer-row-container site-footer-focus-item site-footer-row-layout-standard site-footer-row-tablet-layout-default site-footer-row-mobile-layout-default" data-section="kadence_customizer_footer_middle">
	<div class="site-footer-row-container-inner">
				<div class="site-container">
			<div class="site-middle-footer-inner-wrap site-footer-row site-footer-row-columns-3 site-footer-row-column-layout-center-wide site-footer-row-tablet-column-layout-default site-footer-row-mobile-column-layout-row ft-ro-dir-column ft-ro-collapse-normal ft-ro-t-dir-default ft-ro-m-dir-default ft-ro-lstyle-normal">
									<div class="site-footer-middle-section-1 site-footer-section footer-section-inner-items-1">
						<div class="footer-widget-area widget-area site-footer-focus-item footer-widget6 content-align-default content-tablet-align-default content-mobile-align-default content-valign-default content-tablet-valign-default content-mobile-valign-default" data-section="sidebar-widgets-footer6">
	<div class="footer-widget-area-inner site-info-inner">
		<section id="block-11" class="widget widget_block"><div class="gtranslate_wrapper" id="gt-wrapper-16748402"></div></section>	</div>
</div><!-- .footer-widget6 -->
					</div>
										<div class="site-footer-middle-section-2 site-footer-section footer-section-inner-items-1">
						
<div class="footer-widget-area site-info site-footer-focus-item content-align-default content-tablet-align-default content-mobile-align-default content-valign-default content-tablet-valign-default content-mobile-valign-default" data-section="kadence_customizer_footer_html">
	<div class="footer-widget-area-inner site-info-inner">
		<div class="footer-html inner-link-style-plain"><div class="footer-html-inner"><p>&copy; Copyright - Guru99 2023         <a href="/privacy-policy">Privacy Policy</a>  |  <a href="/affiliate-earning-disclaimer">Affiliate Disclaimer</a>  |  <a href="/terms-of-service">ToS</a></p>
</div></div>	</div>
</div><!-- .site-info -->
					</div>
										<div class="site-footer-middle-section-3 site-footer-section footer-section-inner-items-0">
											</div>
								</div>
		</div>
	</div>
</div>
	</div>
</footer><!-- #colophon -->


<script type="rocketlazyloadscript">
document.querySelectorAll('a img').forEach( function(item, index){
	item.closest('a').classList.add('nohover');
});
</script>
</div><!-- #wrapper -->

<style id='kadence_mega_menu_inline-inline-css'>
#menu-item-3173.kadence-menu-mega-enabled > .sub-menu{width:600px;}.header-navigation[class*="header-navigation-dropdown-animation-fade"] #menu-item-3173.kadence-menu-mega-enabled > .sub-menu{margin-left:-300px;}#menu-item-3173.kadence-menu-mega-enabled > .sub-menu{background-color:#323a56;}.header-navigation .header-menu-container #menu-item-3173.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{color:var(--global-palette3);}.header-navigation .header-menu-container #menu-item-3173.kadence-menu-mega-enabled > .sub-menu li.menu-item > a{background:#323a56;}.header-navigation .header-menu-container #menu-item-3173.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{background:var(--global-palette9);}#menu-item-3174.kadence-menu-mega-enabled > .sub-menu{width:600px;}.header-navigation[class*="header-navigation-dropdown-animation-fade"] #menu-item-3174.kadence-menu-mega-enabled > .sub-menu{margin-left:-300px;}#menu-item-3174.kadence-menu-mega-enabled > .sub-menu{background-color:#323a56;}.header-navigation .header-menu-container #menu-item-3174.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{color:var(--global-palette3);}.header-navigation .header-menu-container #menu-item-3174.kadence-menu-mega-enabled > .sub-menu li.menu-item > a{background:#323a56;}.header-navigation .header-menu-container #menu-item-3174.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{background:var(--global-palette9);}#menu-item-3175.kadence-menu-mega-enabled > .sub-menu{width:600px;}.header-navigation[class*="header-navigation-dropdown-animation-fade"] #menu-item-3175.kadence-menu-mega-enabled > .sub-menu{margin-left:-300px;}#menu-item-3175.kadence-menu-mega-enabled > .sub-menu{background-color:#323a56;}.header-navigation .header-menu-container #menu-item-3175.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{color:var(--global-palette3);}.header-navigation .header-menu-container #menu-item-3175.kadence-menu-mega-enabled > .sub-menu li.menu-item > a{background:#323a56;}.header-navigation .header-menu-container #menu-item-3175.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{background:var(--global-palette9);}#menu-item-3176.kadence-menu-mega-enabled > .sub-menu{width:600px;}.header-navigation[class*="header-navigation-dropdown-animation-fade"] #menu-item-3176.kadence-menu-mega-enabled > .sub-menu{margin-left:-300px;}#menu-item-3176.kadence-menu-mega-enabled > .sub-menu{background-color:#323a56;}.header-navigation .header-menu-container #menu-item-3176.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{color:var(--global-palette3);}.header-navigation .header-menu-container #menu-item-3176.kadence-menu-mega-enabled > .sub-menu li.menu-item > a{background:#323a56;}.header-navigation .header-menu-container #menu-item-3176.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{background:var(--global-palette9);}#menu-item-3177.kadence-menu-mega-enabled > .sub-menu{width:600px;}.header-navigation[class*="header-navigation-dropdown-animation-fade"] #menu-item-3177.kadence-menu-mega-enabled > .sub-menu{margin-left:-300px;}#menu-item-3177.kadence-menu-mega-enabled > .sub-menu{background-color:#323a56;}.header-navigation .header-menu-container #menu-item-3177.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{color:var(--global-palette3);}.header-navigation .header-menu-container #menu-item-3177.kadence-menu-mega-enabled > .sub-menu li.menu-item > a{background:#323a56;}.header-navigation .header-menu-container #menu-item-3177.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{background:var(--global-palette9);}#menu-item-3178.kadence-menu-mega-enabled > .sub-menu{width:400px;}.header-navigation[class*="header-navigation-dropdown-animation-fade"] #menu-item-3178.kadence-menu-mega-enabled > .sub-menu{margin-left:-400px;}#menu-item-3178.kadence-menu-mega-enabled > .sub-menu{background-color:#323a56;}.header-navigation .header-menu-container #menu-item-3178.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{color:var(--global-palette3);}.header-navigation .header-menu-container #menu-item-3178.kadence-menu-mega-enabled > .sub-menu li.menu-item > a{background:#323a56;}.header-navigation .header-menu-container #menu-item-3178.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{background:var(--global-palette9);}#menu-item-3179.kadence-menu-mega-enabled > .sub-menu{width:400px;}.header-navigation[class*="header-navigation-dropdown-animation-fade"] #menu-item-3179.kadence-menu-mega-enabled > .sub-menu{margin-left:-400px;}#menu-item-3179.kadence-menu-mega-enabled > .sub-menu{background-color:#323a56;}.header-navigation .header-menu-container #menu-item-3179.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{color:var(--global-palette3);}.header-navigation .header-menu-container #menu-item-3179.kadence-menu-mega-enabled > .sub-menu li.menu-item > a{background:#323a56;}.header-navigation .header-menu-container #menu-item-3179.kadence-menu-mega-enabled > .sub-menu li.menu-item > a:hover{background:var(--global-palette9);}
</style>
			<script type="rocketlazyloadscript">document.documentElement.style.setProperty('--scrollbar-offset', window.innerWidth - document.documentElement.clientWidth + 'px' );</script>
			<a id="kt-scroll-up" tabindex="-1" aria-hidden="true" aria-label="Scroll to top" href="#wrapper" class="kadence-scroll-to-top scroll-up-wrap scroll-ignore scroll-up-side-right scroll-up-style-outline vs-lg-true vs-md-true vs-sm-false"><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-up2-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="26" height="28" viewBox="0 0 26 28"><title>Scroll to top</title><path d="M25.172 15.172c0 0.531-0.219 1.031-0.578 1.406l-1.172 1.172c-0.375 0.375-0.891 0.594-1.422 0.594s-1.047-0.219-1.406-0.594l-4.594-4.578v11c0 1.125-0.938 1.828-2 1.828h-2c-1.062 0-2-0.703-2-1.828v-11l-4.594 4.578c-0.359 0.375-0.875 0.594-1.406 0.594s-1.047-0.219-1.406-0.594l-1.172-1.172c-0.375-0.375-0.594-0.875-0.594-1.406s0.219-1.047 0.594-1.422l10.172-10.172c0.359-0.375 0.875-0.578 1.406-0.578s1.047 0.203 1.422 0.578l10.172 10.172c0.359 0.375 0.578 0.891 0.578 1.422z"></path>
				</svg></span></a><button id="kt-scroll-up-reader" href="#wrapper" aria-label="Scroll to top" class="kadence-scroll-to-top scroll-up-wrap scroll-ignore scroll-up-side-right scroll-up-style-outline vs-lg-true vs-md-true vs-sm-false"><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-up2-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="26" height="28" viewBox="0 0 26 28"><title>Scroll to top</title><path d="M25.172 15.172c0 0.531-0.219 1.031-0.578 1.406l-1.172 1.172c-0.375 0.375-0.891 0.594-1.422 0.594s-1.047-0.219-1.406-0.594l-4.594-4.578v11c0 1.125-0.938 1.828-2 1.828h-2c-1.062 0-2-0.703-2-1.828v-11l-4.594 4.578c-0.359 0.375-0.875 0.594-1.406 0.594s-1.047-0.219-1.406-0.594l-1.172-1.172c-0.375-0.375-0.594-0.875-0.594-1.406s0.219-1.047 0.594-1.422l10.172-10.172c0.359-0.375 0.875-0.578 1.406-0.578s1.047 0.203 1.422 0.578l10.172 10.172c0.359 0.375 0.578 0.891 0.578 1.422z"></path>
				</svg></span></button>
<script type="rocketlazyloadscript">
var is_mobile = /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent); 
if(is_mobile == false){
!function(e,t){(e=t.createElement("script")).src="https://cdn.convertbox.com/convertbox/js/embed.js",e.id="app-convertbox-script",e.async=true,e.dataset.uuid="37ed7c7c-8ffa-41a9-9264-26aac657f888",document.getElementsByTagName("head")[0].appendChild(e)}(window,document);
}
</script>

<!-- Google tag (gtag.js) -->
<script async src="https://www.googletagmanager.com/gtag/js?id=G-30SEBTJ46Q"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'G-30SEBTJ46Q');
</script>

<!-- Below is the code required to invoke the on call sticky footer. -->
<script>
 window.freestar.queue.push(function(){
   window.freestar.newStickyFooter("guru99_sticky_footer");
 });
</script>
	<div id="mobile-drawer" class="popup-drawer popup-drawer-layout-sidepanel popup-drawer-animation-fade popup-drawer-side-right" data-drawer-target-string="#mobile-drawer"
			>
		<div class="drawer-overlay" data-drawer-target-string="#mobile-drawer"></div>
		<div class="drawer-inner">
						<div class="drawer-header">
				<button class="menu-toggle-close drawer-toggle" aria-label="Close menu"  data-toggle-target="#mobile-drawer" data-toggle-body-class="showing-popup-drawer-from-right" aria-expanded="false" data-set-focus=".menu-toggle-open"
							>
					<span class="toggle-close-bar"></span>
					<span class="toggle-close-bar"></span>
				</button>
			</div>
			<div class="drawer-content mobile-drawer-content content-align-left content-valign-top">
								<div class="site-header-item site-header-focus-item site-header-item-mobile-navigation mobile-navigation-layout-stretch-false" data-section="kadence_customizer_mobile_navigation">
		<nav id="mobile-site-navigation" class="mobile-navigation drawer-navigation drawer-navigation-parent-toggle-false" role="navigation" aria-label="Primary Mobile Navigation">
				<div class="mobile-menu-container drawer-menu-container">
			<ul id="mobile-menu" class="menu has-collapse-sub-nav"><li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-3172"><a href="/">Home</a></li>
<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3173 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-3 kadence-menu-mega-layout-equal"><div class="drawer-nav-drop-wrap"><a href="/software-testing.html">Testing</a><button class="drawer-sub-toggle" data-toggle-duration="10" data-toggle-target="#mobile-menu .menu-item-3173 &gt; .sub-menu" aria-expanded="false"><span class="screen-reader-text">Expand child menu</span><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></button></div>
<ul class="sub-menu">
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4569"><a href="https://www.guru99.com/agile-testing-course.html">Agile Testing</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4572"><a href="/junit-tutorial.html">JUnit</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4579"><a href="/hp-alm-free-tutorial.html">Quality Center(ALM)</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4570"><a href="/bugzilla-tutorial-for-beginners.html">Bugzilla</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4584"><a href="/loadrunner-v12-tutorials.html">HP Loadrunner</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4593"><a href="/rpa-tutorial.html">RPA</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4571"><a href="https://www.guru99.com/cucumber-tutorials.html">Cucumber</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4600"><a href="/software-testing.html">Software Testing</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4606"><a href="/learn-sap-testing-create-your-first-sap-test-case.html">SAP Testing</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4608"><a href="https://www.guru99.com/data-testing.html">Database Testing</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4616"><a href="/mobile-testing.html">Mobile Testing</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4622"><a href="/selenium-tutorial.html">Selenium</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-4626"><a href="/utlimate-guide-etl-datawarehouse-testing.html">ETL Testing</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4628"><a href="https://www.guru99.com/mantis-bug-tracker-tutorial.html">Mantis</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4635"><a href="https://www.guru99.com/soapui-tutorial.html">SoapUI</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4640"><a href="https://www.guru99.com/jmeter-tutorials.html">JMeter</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4646"><a href="https://www.guru99.com/postman-tutorial.html">Postman</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4653"><a href="https://www.guru99.com/test-management.html">TEST Management</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4658"><a href="https://www.guru99.com/jira-tutorial-a-complete-guide-for-beginners.html">JIRA</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4663"><a href="https://www.guru99.com/quick-test-professional-qtp-tutorial.html">QTP</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4665"><a href="https://www.guru99.com/testlink-tutorial-complete-guide.html">TestLink</a></li>
</ul>
</li>
<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3174 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-3 kadence-menu-mega-layout-equal"><div class="drawer-nav-drop-wrap"><a href="/what-is-sap.html">SAP</a><button class="drawer-sub-toggle" data-toggle-duration="10" data-toggle-target="#mobile-menu .menu-item-3174 &gt; .sub-menu" aria-expanded="false"><span class="screen-reader-text">Expand child menu</span><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></button></div>
<ul class="sub-menu">
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4678"><a href="https://www.guru99.com/abap-tutorial.html">ABAP</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4681"><a href="https://www.guru99.com/sap-crm-training.html">CRM</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4683"><a href="https://www.guru99.com/sap-pi-process-integration-tutorial.html">PI/PO</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4685"><a href="https://www.guru99.com/overview-of-sap-apo.html">APO</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4689"><a href="https://www.guru99.com/crystal-reports-tutorial.html">Crystal Reports</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4695"><a href="https://www.guru99.com/sap-pp-tutorials.html">PP</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4699"><a href="https://www.guru99.com/what-is-sap.html">Beginners</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4707"><a href="https://www.guru99.com/sap-fico-training-tutorials.html">FICO</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4709"><a href="https://www.guru99.com/free-sap-sd-training-course.html">SD</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4713"><a href="https://www.guru99.com/sap-basis-training-tutorials.html">Basis</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4717"><a href="https://www.guru99.com/sap-hana-tutorial.html">HANA</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4721"><a href="https://www.guru99.com/sapui5-tutorial.html">SAPUI5</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4732"><a href="https://www.guru99.com/sap-bods-tutorial.html">BODS</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4739"><a href="https://www.guru99.com/sap-hcm.html">HR</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4744"><a href="https://www.guru99.com/overview-of-sap-security.html">Security Tutorial</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4748"><a href="https://www.guru99.com/sap-bi.html">BI/BW</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4759"><a href="https://www.guru99.com/sap-mm-training-tutorials.html">MM</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4762"><a href="https://www.guru99.com/overview-of-sap-solution-manager.html">Solution Manager</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4764"><a href="https://www.guru99.com/sap-bpc.html">BPC</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4768"><a href="https://www.guru99.com/sap-quality-management-qm-tutorial.html">QM</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4771"><a href="https://www.guru99.com/sap-successfactor.html">Successfactors</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4773"><a href="https://www.guru99.com/co-tutorials.html">CO</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4775"><a href="https://www.guru99.com/sap-payroll.html">Payroll</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4778"><a href="https://www.guru99.com/sap-training-hub.html">SAP Courses</a></li>
</ul>
</li>
<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3175 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-4 kadence-menu-mega-layout-equal"><div class="drawer-nav-drop-wrap"><a href="/java-tutorial.html">Web</a><button class="drawer-sub-toggle" data-toggle-duration="10" data-toggle-target="#mobile-menu .menu-item-3175 &gt; .sub-menu" aria-expanded="false"><span class="screen-reader-text">Expand child menu</span><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></button></div>
<ul class="sub-menu">
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4793"><a href="https://www.guru99.com/apache.html">Apache</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4796"><a href="https://www.guru99.com/java-tutorial.html">Java</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4799"><a href="https://www.guru99.com/php-tutorials.html">PHP</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4800"><a href="https://www.guru99.com/ms-sql-server-tutorial.html">SQL Server</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4802"><a href="https://www.guru99.com/angularjs-tutorial.html">AngularJS</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4805"><a href="https://www.guru99.com/jsp-tutorial.html">JSP</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4806"><a href="https://www.guru99.com/pl-sql-tutorials.html">PL/SQL</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4809"><a href="https://www.guru99.com/uml-tutorial.html">UML</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4811"><a href="https://www.guru99.com/asp-net-tutorial.html">ASP.NET</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4817"><a href="https://www.guru99.com/kotlin-tutorial.html">Kotlin</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4819"><a href="https://www.guru99.com/postgresql-tutorial.html">PostgreSQL</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4824"><a href="https://www.guru99.com/vb-net-tutorial.html">VB.NET</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4827"><a href="https://www.guru99.com/c-programming-tutorial.html">C</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4830"><a href="https://www.guru99.com/unix-linux-tutorial.html">Linux</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4833"><a href="https://www.guru99.com/python-tutorials.html">Python</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4835"><a href="https://www.guru99.com/vbscript-tutorials-for-beginners.html">VBScript</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4838"><a href="https://www.guru99.com/c-sharp-tutorial.html">C#</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4845"><a href="https://www.guru99.com/mariadb-tutorial-install.html">MariaDB</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4846"><a href="https://www.guru99.com/reactjs-tutorial.html">ReactJS</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4847"><a href="https://www.guru99.com/web-services-tutorial.html">Web Services</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4850"><a href="https://www.guru99.com/cpp-programming-tutorial.html">C++</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4852"><a href="https://www.guru99.com/ms-access-tutorial.html">MS Access</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4854"><a href="https://www.guru99.com/ruby-on-rails-tutorial.html">Ruby &#038; Rails</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4857"><a href="https://www.guru99.com/wpf-tutorial.html">WPF</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4863"><a href="https://www.guru99.com/codeigniter-tutorial.html">CodeIgniter</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4864"><a href="https://www.guru99.com/mysql-tutorial.html">MySQL</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4868"><a href="https://www.guru99.com/scala-tutorial.html">Scala</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4887"><a href="https://www.guru99.com/sqlite-tutorial.html">SQLite</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4871"><a href="https://www.guru99.com/dbms-tutorial.html">DBMS</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4875"><a href="https://www.guru99.com/node-js-tutorial.html">Node.js</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4877"><a href="https://www.guru99.com/sql.html">SQL</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4884"><a href="https://www.guru99.com/perl-tutorials.html">Perl</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4880"><a href="https://www.guru99.com/interactive-javascript-tutorials.html">JavaScript</a></li>
</ul>
</li>
<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3176 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-3 kadence-menu-mega-layout-equal"><div class="drawer-nav-drop-wrap"><a href="/design-analysis-algorithms-tutorial.html">Must Learn</a><button class="drawer-sub-toggle" data-toggle-duration="10" data-toggle-target="#mobile-menu .menu-item-3176 &gt; .sub-menu" aria-expanded="false"><span class="screen-reader-text">Expand child menu</span><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></button></div>
<ul class="sub-menu">
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4895"><a href="https://www.guru99.com/accounting.html">Accounting</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4897"><a href="https://www.guru99.com/embedded-systems-tutorial.html">Embedded Systems</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4980"><a href="https://www.guru99.com/os-tutorial.html">Operating System</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4906"><a href="https://www.guru99.com/design-analysis-algorithms-tutorial.html">Algorithms</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4909"><a href="https://www.guru99.com/ethical-hacking-tutorials.html">Ethical Hacking</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4911"><a href="https://www.guru99.com/pmp-tutorial.html">PMP</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4914"><a href="https://www.guru99.com/android-tutorial.html">Android</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4918"><a href="https://www.guru99.com/excel-tutorials.html">Excel Tutorial</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4919"><a href="https://www.guru99.com/photoshop-tutorials.html">Photoshop</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-15774"><a href="https://www.guru99.com/cryptocurrency-tutorial.html">Blockchain</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4923"><a href="https://www.guru99.com/google-go-tutorial.html">Go Programming</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4927"><a href="https://www.guru99.com/project-management-tutorial.html">Project Management</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4930"><a href="https://www.guru99.com/business-analyst-tutorial-course.html">Business Analyst</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4934"><a href="https://www.guru99.com/iot-tutorial.html">IoT</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4937"><a href="https://www.guru99.com/best-hard-disks.html">Reviews</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-17034"><a href="https://www.guru99.com/web-design-and-development-tutorial.html">Build Website</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4948"><a href="https://www.guru99.com/itil-framework-process.html">ITIL</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4951"><a href="https://www.guru99.com/salesforce-tutorial.html">Salesforce</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4953"><a href="https://www.guru99.com/cloud-computing-for-beginners.html">Cloud Computing</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4958"><a href="https://www.guru99.com/jenkins-tutorial.html">Jenkins</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4960"><a href="https://www.guru99.com/seo-tutorial.html">SEO</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4964"><a href="https://www.guru99.com/learn-cobol-programming-tutorial.html">COBOL</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4965"><a href="https://www.guru99.com/mis-tutorial.html">MIS</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4968"><a href="https://www.guru99.com/software-engineering-tutorial.html">Software Engineering</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-10495"><a href="https://www.guru99.com/compiler-tutorial.html">Compiler Design</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4976"><a href="https://www.guru99.com/anime-websites-watch-online-free.html">Movie</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4975"><a href="https://www.guru99.com/vba-tutorial.html">VBA</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-16691"><a href="https://www.guru99.com/online-courses.html">Courses</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4971"><a href="https://www.guru99.com/data-communication-computer-network-tutorial.html">Networking</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-16961"><a href="https://www.guru99.com/best-vpn.html">VPN</a></li>
</ul>
</li>
<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3177 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-3 kadence-menu-mega-layout-equal"><div class="drawer-nav-drop-wrap"><a href="/bigdata-tutorials.html">Big Data</a><button class="drawer-sub-toggle" data-toggle-duration="10" data-toggle-target="#mobile-menu .menu-item-3177 &gt; .sub-menu" aria-expanded="false"><span class="screen-reader-text">Expand child menu</span><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></button></div>
<ul class="sub-menu">
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4990"><a href="https://www.guru99.com/aws-tutorial.html">AWS</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4993"><a href="https://www.guru99.com/hive-tutorials.html">Hive</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4996"><a href="https://www.guru99.com/power-bi-tutorial.html">Power BI</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-4997"><a href="https://www.guru99.com/bigdata-tutorials.html">Big Data</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5000"><a href="https://www.guru99.com/informatica-tutorials.html">Informatica</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5001"><a href="https://www.guru99.com/qlikview-tutorial.html">Qlikview</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5005"><a href="https://www.guru99.com/cassandra-tutorial.html">Cassandra</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5008"><a href="https://www.guru99.com/microstrategy-tutorial.html">MicroStrategy</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5011"><a href="https://www.guru99.com/tableau-tutorial.html">Tableau</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5016"><a href="https://www.guru99.com/cognos-tutorial.html">Cognos</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5018"><a href="https://www.guru99.com/mongodb-tutorials.html">MongoDB</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5020"><a href="https://www.guru99.com/talend-tutorial.html">Talend</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5024"><a href="https://www.guru99.com/data-warehousing-tutorial.html">Data Warehousing</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5030"><a href="https://www.guru99.com/apache-nifi-tutorial.html">NiFi</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5036"><a href="https://www.guru99.com/zookeeper-tutorial.html">ZooKeeper</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5039"><a href="https://www.guru99.com/devops-tutorial.html">DevOps</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5041"><a href="https://www.guru99.com/obiee-tutorial.html">OBIEE</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5049"><a href="https://www.guru99.com/pentaho-tutorial.html">Pentaho</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5046"><a href="https://www.guru99.com/hbase-tutorials.html">HBase</a></li>
</ul>
</li>
<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3178 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-2 kadence-menu-mega-layout-equal"><div class="drawer-nav-drop-wrap"><a href="/live-testing-project.html">Live Project</a><button class="drawer-sub-toggle" data-toggle-duration="10" data-toggle-target="#mobile-menu .menu-item-3178 &gt; .sub-menu" aria-expanded="false"><span class="screen-reader-text">Expand child menu</span><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></button></div>
<ul class="sub-menu">
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5064"><a href="https://www.guru99.com/live-agile-testing-project.html">Live Agile Testing</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5067"><a href="https://www.guru99.com/live-selenium-project.html">Live Selenium Project</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5070"><a href="https://www.guru99.com/live-interactive-exercise-hp-alm.html">Live HP ALM</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5076"><a href="https://www.guru99.com/live-ecommerce-project.html">Live Selenium 2</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5080"><a href="https://www.guru99.com/live-java-project.html">Live Java Project</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5082"><a href="https://www.guru99.com/live-penetration-testing-project.html">Live Security Testing</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5086"><a href="https://www.guru99.com/live-mobile-testing-project.html">Live Mobile Testing</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5089"><a href="https://www.guru99.com/live-testing-project.html">Live Testing Project</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5092"><a href="https://www.guru99.com/live-payment-gateway-project.html">Live Payment Gateway</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5093"><a href="https://www.guru99.com/live-insurance-testing-project.html">Live Testing 2</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5096"><a href="https://www.guru99.com/live-php-project-learn-complete-web-development-cycle.html">Live PHP  Project</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5100"><a href="https://www.guru99.com/live-telecom-project.html">Live Telecom</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5101"><a href="https://www.guru99.com/live-projects.html">Live Projects Hub</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5103"><a href="https://www.guru99.com/live-uft-testing.html">Live UFT/QTP Testing</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5108"><a href="https://www.guru99.com/live-python-project.html">Live Python Project</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5111"><a href="https://www.guru99.com/live-seo-project.html">Live SEO Project</a></li>
</ul>
</li>
<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-has-children menu-item-3179 kadence-menu-mega-enabled kadence-menu-mega-width-custom kadence-menu-mega-columns-2 kadence-menu-mega-layout-equal"><div class="drawer-nav-drop-wrap"><a href="/artificial-intelligence-tutorial.html">AI</a><button class="drawer-sub-toggle" data-toggle-duration="10" data-toggle-target="#mobile-menu .menu-item-3179 &gt; .sub-menu" aria-expanded="false"><span class="screen-reader-text">Expand child menu</span><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-arrow-down-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Expand</title><path d="M5.293 9.707l6 6c0.391 0.391 1.024 0.391 1.414 0l6-6c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span></button></div>
<ul class="sub-menu">
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-16679"><a href="https://www.guru99.com/ai-tutorial.html">Artificial Intelligence</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5120"><a href="https://www.guru99.com/pytorch-tutorial.html">PyTorch</a></li>
	<li class="menu-item menu-item-type-custom menu-item-object-custom menu-item-16520"><a href="https://www.guru99.com/data-science-tutorial-for-beginners.html">Data Science</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5124"><a href="https://www.guru99.com/r-tutorial.html">R Programming</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5127"><a href="https://www.guru99.com/keras-tutorial.html">Keras</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5128"><a href="https://www.guru99.com/tensorflow-tutorial.html">TensorFlow</a></li>
	<li class="menu-item menu-item-type-post_type menu-item-object-post menu-item-5129"><a href="https://www.guru99.com/nltk-tutorial.html">NLTK</a></li>
</ul>
</li>
</ul>		</div>
	</nav><!-- #site-navigation -->
	</div><!-- data-section="mobile_navigation" -->
							</div>
		</div>
	</div>
	
<script type="rocketlazyloadscript" id="rocket-browser-checker-js-after">
"use strict";var _createClass=function(){function defineProperties(target,props){for(var i=0;i<props.length;i++){var descriptor=props[i];descriptor.enumerable=descriptor.enumerable||!1,descriptor.configurable=!0,"value"in descriptor&&(descriptor.writable=!0),Object.defineProperty(target,descriptor.key,descriptor)}}return function(Constructor,protoProps,staticProps){return protoProps&&defineProperties(Constructor.prototype,protoProps),staticProps&&defineProperties(Constructor,staticProps),Constructor}}();function _classCallCheck(instance,Constructor){if(!(instance instanceof Constructor))throw new TypeError("Cannot call a class as a function")}var RocketBrowserCompatibilityChecker=function(){function RocketBrowserCompatibilityChecker(options){_classCallCheck(this,RocketBrowserCompatibilityChecker),this.passiveSupported=!1,this._checkPassiveOption(this),this.options=!!this.passiveSupported&&options}return _createClass(RocketBrowserCompatibilityChecker,[{key:"_checkPassiveOption",value:function(self){try{var options={get passive(){return!(self.passiveSupported=!0)}};window.addEventListener("test",null,options),window.removeEventListener("test",null,options)}catch(err){self.passiveSupported=!1}}},{key:"initRequestIdleCallback",value:function(){!1 in window&&(window.requestIdleCallback=function(cb){var start=Date.now();return setTimeout(function(){cb({didTimeout:!1,timeRemaining:function(){return Math.max(0,50-(Date.now()-start))}})},1)}),!1 in window&&(window.cancelIdleCallback=function(id){return clearTimeout(id)})}},{key:"isDataSaverModeOn",value:function(){return"connection"in navigator&&!0===navigator.connection.saveData}},{key:"supportsLinkPrefetch",value:function(){var elem=document.createElement("link");return elem.relList&&elem.relList.supports&&elem.relList.supports("prefetch")&&window.IntersectionObserver&&"isIntersecting"in IntersectionObserverEntry.prototype}},{key:"isSlowConnection",value:function(){return"connection"in navigator&&"effectiveType"in navigator.connection&&("2g"===navigator.connection.effectiveType||"slow-2g"===navigator.connection.effectiveType)}}]),RocketBrowserCompatibilityChecker}();
</script>
<script id='rocket-preload-links-js-extra'>
var RocketPreloadLinksConfig = {"excludeUris":"\/(?:.+\/)?feed(?:\/(?:.+\/?)?)?$|\/(?:.+\/)?embed\/|\/(index.php\/)?(.*)wp-json(\/.*|$)|\/refer\/|\/go\/|\/recommend\/|\/recommends\/","usesTrailingSlash":"","imageExt":"jpg|jpeg|gif|png|tiff|bmp|webp|avif|pdf|doc|docx|xls|xlsx|php","fileExt":"jpg|jpeg|gif|png|tiff|bmp|webp|avif|pdf|doc|docx|xls|xlsx|php|html|htm","siteUrl":"https:\/\/www.guru99.com","onHoverDelay":"100","rateThrottle":"3"};
</script>
<script type="rocketlazyloadscript" id="rocket-preload-links-js-after">
(function() {
"use strict";var r="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},e=function(){function i(e,t){for(var n=0;n<t.length;n++){var i=t[n];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(e,i.key,i)}}return function(e,t,n){return t&&i(e.prototype,t),n&&i(e,n),e}}();function i(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}var t=function(){function n(e,t){i(this,n),this.browser=e,this.config=t,this.options=this.browser.options,this.prefetched=new Set,this.eventTime=null,this.threshold=1111,this.numOnHover=0}return e(n,[{key:"init",value:function(){!this.browser.supportsLinkPrefetch()||this.browser.isDataSaverModeOn()||this.browser.isSlowConnection()||(this.regex={excludeUris:RegExp(this.config.excludeUris,"i"),images:RegExp(".("+this.config.imageExt+")$","i"),fileExt:RegExp(".("+this.config.fileExt+")$","i")},this._initListeners(this))}},{key:"_initListeners",value:function(e){-1<this.config.onHoverDelay&&document.addEventListener("mouseover",e.listener.bind(e),e.listenerOptions),document.addEventListener("mousedown",e.listener.bind(e),e.listenerOptions),document.addEventListener("touchstart",e.listener.bind(e),e.listenerOptions)}},{key:"listener",value:function(e){var t=e.target.closest("a"),n=this._prepareUrl(t);if(null!==n)switch(e.type){case"mousedown":case"touchstart":this._addPrefetchLink(n);break;case"mouseover":this._earlyPrefetch(t,n,"mouseout")}}},{key:"_earlyPrefetch",value:function(t,e,n){var i=this,r=setTimeout(function(){if(r=null,0===i.numOnHover)setTimeout(function(){return i.numOnHover=0},1e3);else if(i.numOnHover>i.config.rateThrottle)return;i.numOnHover++,i._addPrefetchLink(e)},this.config.onHoverDelay);t.addEventListener(n,function e(){t.removeEventListener(n,e,{passive:!0}),null!==r&&(clearTimeout(r),r=null)},{passive:!0})}},{key:"_addPrefetchLink",value:function(i){return this.prefetched.add(i.href),new Promise(function(e,t){var n=document.createElement("link");n.rel="prefetch",n.href=i.href,n.onload=e,n.onerror=t,document.head.appendChild(n)}).catch(function(){})}},{key:"_prepareUrl",value:function(e){if(null===e||"object"!==(void 0===e?"undefined":r(e))||!1 in e||-1===["http:","https:"].indexOf(e.protocol))return null;var t=e.href.substring(0,this.config.siteUrl.length),n=this._getPathname(e.href,t),i={original:e.href,protocol:e.protocol,origin:t,pathname:n,href:t+n};return this._isLinkOk(i)?i:null}},{key:"_getPathname",value:function(e,t){var n=t?e.substring(this.config.siteUrl.length):e;return n.startsWith("/")||(n="/"+n),this._shouldAddTrailingSlash(n)?n+"/":n}},{key:"_shouldAddTrailingSlash",value:function(e){return this.config.usesTrailingSlash&&!e.endsWith("/")&&!this.regex.fileExt.test(e)}},{key:"_isLinkOk",value:function(e){return null!==e&&"object"===(void 0===e?"undefined":r(e))&&(!this.prefetched.has(e.href)&&e.origin===this.config.siteUrl&&-1===e.href.indexOf("?")&&-1===e.href.indexOf("#")&&!this.regex.excludeUris.test(e.href)&&!this.regex.images.test(e.href))}}],[{key:"run",value:function(){"undefined"!=typeof RocketPreloadLinksConfig&&new n(new RocketBrowserCompatibilityChecker({capture:!0,passive:!0}),RocketPreloadLinksConfig).init()}}]),n}();t.run();
}());
</script>
<script type="rocketlazyloadscript" data-rocket-src='https://www.guru99.com/wp-content/themes/kadence/assets/js/simplelightbox.min.js' id='kadence-simplelightbox-js' async></script>
<script type="rocketlazyloadscript" data-rocket-src='https://www.guru99.com/wp-content/themes/kadence/assets/js/lightbox-init.min.js' id='kadence-lightbox-init-js' async></script>
<script id='kadence-navigation-js-extra'>
var kadenceConfig = {"screenReader":{"expand":"Expand child menu","expandOf":"Expand child menu of","collapse":"Collapse child menu","collapseOf":"Collapse child menu of"},"breakPoints":{"desktop":"1024","tablet":768},"scrollOffset":"0"};
</script>
<script type="rocketlazyloadscript" data-rocket-src='https://www.guru99.com/wp-content/themes/kadence/assets/js/navigation.min.js' id='kadence-navigation-js' async></script>
<script id='kadence-blocks-tableofcontents-js-extra'>
var kadence_blocks_toc = {"headings":"[{\"anchor\":\"software-testing\",\"content\":\"Software Testing\",\"level\":2,\"page\":1},{\"anchor\":\"why-software-testing-is-important\",\"content\":\"Why Software Testing is Important?\",\"level\":2,\"page\":1},{\"anchor\":\"what-are-the-benefits-of-software-testing\",\"content\":\"What are the benefits of Software Testing?\",\"level\":2,\"page\":1},{\"anchor\":\"testing-in-software-engineering\",\"content\":\"Testing in Software Engineering\",\"level\":2,\"page\":1},{\"anchor\":\"types-of-software-testing\",\"content\":\"Types of Software Testing\",\"level\":2,\"page\":1},{\"anchor\":\"testing-strategies-in-software-engineering\",\"content\":\"Testing Strategies in Software Engineering\",\"level\":2,\"page\":1},{\"anchor\":\"program-testing\",\"content\":\"Program Testing\",\"level\":2,\"page\":1},{\"anchor\":\"summary-of-software-testing-basics\",\"content\":\"Summary of Software Testing Basics\",\"level\":2,\"page\":1}]","expandText":"Expand Table of Contents","collapseText":"Collapse Table of Contents"};
var kadence_blocks_toc = {"headings":"[{\"anchor\":\"software-testing\",\"content\":\"Software Testing\",\"level\":2,\"page\":1},{\"anchor\":\"why-software-testing-is-important\",\"content\":\"Why Software Testing is Important?\",\"level\":2,\"page\":1},{\"anchor\":\"what-are-the-benefits-of-software-testing\",\"content\":\"What are the benefits of Software Testing?\",\"level\":2,\"page\":1},{\"anchor\":\"testing-in-software-engineering\",\"content\":\"Testing in Software Engineering\",\"level\":2,\"page\":1},{\"anchor\":\"types-of-software-testing\",\"content\":\"Types of Software Testing\",\"level\":2,\"page\":1},{\"anchor\":\"testing-strategies-in-software-engineering\",\"content\":\"Testing Strategies in Software Engineering\",\"level\":2,\"page\":1},{\"anchor\":\"program-testing\",\"content\":\"Program Testing\",\"level\":2,\"page\":1},{\"anchor\":\"summary-of-software-testing-basics\",\"content\":\"Summary of Software Testing Basics\",\"level\":2,\"page\":1}]","expandText":"Expand Table of Contents","collapseText":"Collapse Table of Contents"};
</script>
<script type="rocketlazyloadscript" data-rocket-src='https://www.guru99.com/wp-content/plugins/kadence-blocks/includes/assets/js/kb-table-of-contents.min.js' id='kadence-blocks-tableofcontents-js'></script>
<script type="rocketlazyloadscript" data-rocket-src='https://www.guru99.com/wp-content/plugins/kadence-pro/dist/mega-menu/kadence-mega-menu.min.js' id='kadence-mega-menu-js'></script>
<script type="rocketlazyloadscript" id="gt_widget_script_16748402-js-before">
window.gtranslateSettings = /* document.write */ window.gtranslateSettings || {};window.gtranslateSettings['16748402'] = {"default_language":"en","languages":["en","fr"],"url_structure":"sub_directory","flag_style":"3d","flag_size":24,"wrapper_selector":"#gt-wrapper-16748402","alt_flags":[],"switcher_open_direction":"top","switcher_horizontal_position":"inline","switcher_text_color":"#ffffff","switcher_arrow_color":"#ffffff","switcher_border_color":"#ccc","switcher_background_color":"#0556f3","switcher_background_shadow_color":"#0556f3","switcher_background_hover_color":"#000000","dropdown_text_color":"#ffffff","dropdown_hover_color":"#000000","dropdown_background_color":"#0556f3","custom_css":".gt_option {\r\nposition:absolute!important;\r\nbottom:41px!important;\r\n}\r\n.gt_selected{\r\nmargin-top:28px;\r\n}","flags_location":"\/wp-content\/plugins\/gtranslate\/flags\/"};
</script><script type="rocketlazyloadscript" data-rocket-src="https://www.guru99.com/wp-content/plugins/gtranslate/js/dwf.js" data-no-optimize="1" data-no-minify="1" data-gt-orig-url="/software-testing-introduction-importance.html" data-gt-orig-domain="www.guru99.com" data-gt-widget-id="16748402" defer></script>	<div id="search-drawer" class="popup-drawer popup-drawer-layout-fullwidth" data-drawer-target-string="#search-drawer"
			>
		<div class="drawer-overlay" data-drawer-target-string="#search-drawer"></div>
		<div class="drawer-inner">
			<div class="drawer-header">
				<button class="search-toggle-close drawer-toggle" aria-label="Close search"  data-toggle-target="#search-drawer" data-toggle-body-class="showing-popup-drawer-from-full" aria-expanded="false" data-set-focus=".search-toggle-open"
							>
					<span class="kadence-svg-iconset"><svg class="kadence-svg-icon kadence-close-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><title>Toggle Menu Close</title><path d="M5.293 6.707l5.293 5.293-5.293 5.293c-0.391 0.391-0.391 1.024 0 1.414s1.024 0.391 1.414 0l5.293-5.293 5.293 5.293c0.391 0.391 1.024 0.391 1.414 0s0.391-1.024 0-1.414l-5.293-5.293 5.293-5.293c0.391-0.391 0.391-1.024 0-1.414s-1.024-0.391-1.414 0l-5.293 5.293-5.293-5.293c-0.391-0.391-1.024-0.391-1.414 0s-0.391 1.024 0 1.414z"></path>
				</svg></span>				</button>
			</div>
			<div class="drawer-content">
				<form role="search" method="get" class="search-form" action="https://www.guru99.com/">
				<label>
					<span class="screen-reader-text">Search for:</span>
					<input type="search" class="search-field" placeholder="Search &hellip;" value="" name="s" />
				</label>
				<input type="submit" class="search-submit" value="Search" />
			<div class="kadence-search-icon-wrap"><span class="kadence-svg-iconset"><svg aria-hidden="true" class="kadence-svg-icon kadence-search-svg" fill="currentColor" version="1.1" xmlns="http://www.w3.org/2000/svg" width="26" height="28" viewBox="0 0 26 28"><title>Search</title><path d="M18 13c0-3.859-3.141-7-7-7s-7 3.141-7 7 3.141 7 7 7 7-3.141 7-7zM26 26c0 1.094-0.906 2-2 2-0.531 0-1.047-0.219-1.406-0.594l-5.359-5.344c-1.828 1.266-4.016 1.937-6.234 1.937-6.078 0-11-4.922-11-11s4.922-11 11-11 11 4.922 11 11c0 2.219-0.672 4.406-1.937 6.234l5.359 5.359c0.359 0.359 0.578 0.875 0.578 1.406z"></path>
				</svg></span></div></form>			</div>
		</div>
	</div>
	<style>
	.header-menu-container ul.menu>li.menu-item-has-children>ul.sub-menu,
.header-menu-container ul.menu>li.menu-item-has-children>ul.sub-menu * {
	visibility: hidden !important;
	transition: none;
}

.header-menu-container ul.menu>li.menu-item-has-children>ul.sub-menu.clicked {
    visibility: visible !important;
    opacity: 1 !important;
    clip: auto !important;
    height: auto !important;
    overflow: visible !important;
	margin-top: 10px !important;
}

.header-menu-container ul.menu>li.menu-item-has-children>ul.sub-menu.clicked * {
	visibility: visible !important;
}


.header-menu-container ul.menu>li.menu-item-has-children.menu-item-object-custom>ul.sub-menu.clicked {
	display: grid;
}
</style>
<script type="rocketlazyloadscript">
const megaMenu = document.querySelectorAll('.kadence-menu-mega-enabled > a');
const subMenu = document.querySelectorAll('.kadence-menu-mega-enabled > .sub-menu');
const siteNavigation = document.getElementById('site-navigation');


megaMenu.forEach( function(menu){
	menu.addEventListener('click', function(e){
		e.preventDefault();
		if ( menu.parentElement.querySelector('.sub-menu').classList.contains('clicked') ) {
			subMenu.forEach( function(submenu){
				submenu.classList.remove('clicked');
			});
		} else {
			subMenu.forEach( function(submenu){
				submenu.classList.remove('clicked');
			});
			menu.parentElement.querySelector('.sub-menu').classList.add('clicked');
		}
	});
});

document.addEventListener('click', function(event) {
    let isClickInsideElement = siteNavigation.contains(event.target);
    if (!isClickInsideElement) {
        subMenu.forEach( function(submenu){
			submenu.classList.remove('clicked');
		});
    }
});
</script>
<script>window.lazyLoadOptions=[{elements_selector:"img[data-lazy-src],.rocket-lazyload,iframe[data-lazy-src]",data_src:"lazy-src",data_srcset:"lazy-srcset",data_sizes:"lazy-sizes",class_loading:"lazyloading",class_loaded:"lazyloaded",threshold:300,callback_loaded:function(element){if(element.tagName==="IFRAME"&&element.dataset.rocketLazyload=="fitvidscompatible"){if(element.classList.contains("lazyloaded")){if(typeof window.jQuery!="undefined"){if(jQuery.fn.fitVids){jQuery(element).parent().fitVids()}}}}}},{elements_selector:".rocket-lazyload",data_src:"lazy-src",data_srcset:"lazy-srcset",data_sizes:"lazy-sizes",class_loading:"lazyloading",class_loaded:"lazyloaded",threshold:300,}];window.addEventListener('LazyLoad::Initialized',function(e){var lazyLoadInstance=e.detail.instance;if(window.MutationObserver){var observer=new MutationObserver(function(mutations){var image_count=0;var iframe_count=0;var rocketlazy_count=0;mutations.forEach(function(mutation){for(var i=0;i<mutation.addedNodes.length;i++){if(typeof mutation.addedNodes[i].getElementsByTagName!=='function'){continue}
if(typeof mutation.addedNodes[i].getElementsByClassName!=='function'){continue}
images=mutation.addedNodes[i].getElementsByTagName('img');is_image=mutation.addedNodes[i].tagName=="IMG";iframes=mutation.addedNodes[i].getElementsByTagName('iframe');is_iframe=mutation.addedNodes[i].tagName=="IFRAME";rocket_lazy=mutation.addedNodes[i].getElementsByClassName('rocket-lazyload');image_count+=images.length;iframe_count+=iframes.length;rocketlazy_count+=rocket_lazy.length;if(is_image){image_count+=1}
if(is_iframe){iframe_count+=1}}});if(image_count>0||iframe_count>0||rocketlazy_count>0){lazyLoadInstance.update()}});var b=document.getElementsByTagName("body")[0];var config={childList:!0,subtree:!0};observer.observe(b,config)}},!1)</script><script data-no-minify="1" async src="https://www.guru99.com/wp-content/plugins/wp-rocket/assets/js/lazyload/17.8.3/lazyload.min.js"></script></body>
</html>

<!-- Cached for great performance - Debug: cached@1696245505 -->`
	htmlText, err := textExtract.ExtractHtml(el)
	if err != nil {
		fmt.Printf("error is %s", err.Error())
		return
	}
	text, err := json.Marshal(htmlText)
	if err != nil {
		fmt.Printf("err is %s", err.Error())
	}
	fmt.Println(string(text))
	a := string(text)
	if a != `"Skip to content\nHome\nTestingExpand\nAgile Testing\nJUnit\nQuality Center(ALM)\nBugzilla\nHP Loadrunner\nRPA\nCucumber\nSoftware Testing\nSAP Testing\nDatabase Testing\nMobile Testing\nSelenium\nETL Testing\nMantis\nSoapUI\nJMeter\nPostman\nTEST Management\nJIRA\nQTP\nTestLink\nSAPExpand\nABAP\nCRM\nPI/PO\nAPO\nCrystal Reports\nPP\nBeginners\nFICO\nSD\nBasis\nHANA\nSAPUI5\nBODS\nHR\nSecurity Tutorial\nBI/BW\nMM\nSolution Manager\nBPC\nQM\nSuccessfactors\nCO\nPayroll\nSAP Courses\nWebExpand\nApache\nJava\nPHP\nSQL Server\nAngularJS\nJSP\nPL/SQL\nUML\nASP.NET\nKotlin\nPostgreSQL\nVB.NET\nC\nLinux\nPython\nVBScript\nC#\nMariaDB\nReactJS\nWeb Services\nC++\nMS Access\nRuby \u0026 Rails\nWPF\nCodeIgniter\nMySQL\nScala\nSQLite\nDBMS\nNode.js\nSQL\nPerl\nJavaScript\nMust LearnExpand\nAccounting\nEmbedded Systems\nOperating System\nAlgorithms\nEthical Hacking\nPMP\nAndroid\nExcel Tutorial\nPhotoshop\nBlockchain\nGo Programming\nProject Management\nBusiness Analyst\nIoT\nReviews\nBuild Website\nITIL\nSalesforce\nCloud Computing\nJenkins\nSEO\nCOBOL\nMIS\nSoftware Engineering\nCompiler Design\nMovie\nVBA\nCourses\nNetworking\nVPN\nBig DataExpand\nAWS\nHive\nPower BI\nBig Data\nInformatica\nQlikview\nCassandra\nMicroStrategy\nTableau\nCognos\nMongoDB\nTalend\nData Warehousing\nNiFi\nZooKeeper\nDevOps\nOBIEE\nPentaho\nHBase\nLive ProjectExpand\nLive Agile Testing\nLive Selenium Project\nLive HP ALM\nLive Selenium 2\nLive Java Project\nLive Security Testing\nLive Mobile Testing\nLive Testing Project\nLive Payment Gateway\nLive Testing 2\nLive PHP Project\nLive Telecom\nLive Projects Hub\nLive UFT/QTP Testing\nLive Python Project\nLive SEO Project\nAIExpand\nArtificial Intelligence\nPyTorch\nData Science\nR Programming\nKeras\nTensorFlow\nNLTK\nSearch\nToggle Menu\nWhat is Software Testing? Definition\nBy : \u003cimg src=\"https://www.guru99.com/images/thomas-hamilton-120x120.jpg\" width=\"25\" height=\"25\" alt=\"Thomas Hamilton\" class=\"avatar avatar-25 wp-user-avatar wp-user-avatar-25 alignnone photo\" /\u003e Thomas Hamilton\nHours UpdatedSeptember 19, 2023\nSoftware Testing\nSoftware Testing is a method to check whether the actual software product matches expected requirements and to ensure that software product is Defect free. It involves execution of software/system components using manual or automated tools to evaluate one or more properties of interest. The purpose of software testing is to identify errors, gaps or missing requirements in contrast to actual requirements.\nSome prefer saying Software testing definition as a White Box and Black Box Testing . In simple terms, Software Testing means the Verification of Application Under Test (AUT). This Software Testing course introduces testing software to the audience and justifies the importance of software testing.\nTable of Content:\nSoftware Testing\nWhy Software Testing is Important?\nWhat are the benefits of Software Testing?\nTesting in Software Engineering\nTypes of Software Testing\nTesting Strategies in Software Engineering\nProgram Testing\nSummary of Software Testing Basics\nWhy Software Testing is Important?\nSoftware Testing is Important because if there are any bugs or errors in the software, it can be identified early and can be solved before delivery of the software product. Properly tested software product ensures reliability, security and high performance which further results in time saving, cost effectiveness and customer satisfaction.\n1 Monday\n\u003cimg decoding=\"async\" src=\"https://www.guru99.com/images/monday-logo-v1.png\" width=\"160\" height=\"100\"\u003e\nLearn More\nOn Monday’s Website\nTime Tracking\nYes\nDrag \u0026 Drop\nYes\nFree Trial\nForever Free Plan\n2 JIRA Software\n\u003cimg decoding=\"async\" src=\"https://www.guru99.com/images/jira-software-logo-v1.png\" width=\"225\" height=\"100\"\u003e\nLearn More\nOn Jira Software Website\nTime Tracking\nYes\nDrag \u0026 Drop\nYes\nFree Trial\nForever Free Plan\n3 Smartsheet\n\u003cimg decoding=\"async\" src=\"https://www.guru99.com/images/smartsheet-logo-v3.png\" style=\"width: 90%;\" width=\"250\" height=\"45\"\u003e\nLearn More\nOn Smartsheet’s Website\nTime Tracking\nYes\nDrag \u0026 Drop\nYes\nFree Trial\nForever Free Plan\nWhat is the need of Testing?\nTesting is important because software bugs could be expensive or even dangerous. Software bugs can potentially cause monetary and human loss, and history is full of such examples.\nIn April 2015, Bloomberg terminal in London crashed due to software glitch affected more than 300,000 traders on financial markets. It forced the government to postpone a 3bn pound debt sale.\nNissan cars recalled over 1 million cars from the market due to software failure in the airbag sensory detectors. There has been reported two accident due to this software failure.\nStarbucks was forced to close about 60 percent of stores in the U.S and Canada due to software failure in its POS system. At one point, the store served coffee for free as they were unable to process the transaction.\nSome of Amazon’s third-party retailers saw their product price is reduced to 1p due to a software glitch. They were left with heavy losses.\nVulnerability in Windows 10. This bug enables users to escape from security sandboxes through a flaw in the win32k system.\nIn 2015 fighter plane F-35 fell victim to a software bug, making it unable to detect targets correctly.\nChina Airlines Airbus A300 crashed due to a software bug on April 26, 1994, killing 264 innocents live\nIn 1985, Canada’s Therac-25 radiation therapy machine malfunctioned due to software bug and delivered lethal radiation doses to patients, leaving 3 people dead and critically injuring 3 others.\nIn April of 1999, a software bug caused the failure of a $1.2 billion military satellite launch, the costliest accident in history\nIn May of 1996, a software bug caused the bank accounts of 823 customers of a major U.S. bank to be credited with 920 million US dollars.\nClick here if the video is not accessible\nWhat are the benefits of Software Testing?\nHere are the benefits of using software testing:\nCost-Effective: It is one of the important advantages of software testing. Testing any IT project on time helps you to save your money for the long term. In case if the bugs caught in the earlier stage of software testing, it costs less to fix.\nSecurity: It is the most vulnerable and sensitive benefit of software testing. People are looking for trusted products. It helps in removing risks and problems earlier.\nProduct quality: It is an essential requirement of any software product. Testing ensures a quality product is delivered to customers.\nCustomer Satisfaction: The main aim of any product is to give satisfaction to their customers. UI/UX Testing ensures the best user experience.\n» Also check: Best Software Testing Services Companies\nTesting in Software Engineering\nAs per ANSI/IEEE 1059, Testing in Software Engineering is a process of evaluating a software product to find whether the current software product meets the required conditions or not. The testing process involves evaluating the features of the software product for requirements in terms of any missing requirements, bugs or errors, security, reliability and performance.\nTypes of Software Testing\nHere are the software testing types:\nTypically Testing is classified into three categories.\nFunctional Testing\nNon-Functional Testing or Performance Testing\nMaintenance (Regression and Maintenance)\n\u003cimg decoding=\"async\" fetchpriority=\"high\" width=\"628\" height=\"275\" src=\"https://www.guru99.com/images/2/061920_1310_Whatissoftwaretesting1.png\" alt=\"Types of Software Testing in Software Engineering\"\u003e\nTypes of Software Testing in Software Engineering\nTesting Category\nTypes of Testing\nFunctional Testing\nUnit Testing\nIntegration Testing\nSmoke\nUAT ( User Acceptance Testing)\nLocalization\nGlobalization\nInteroperability\nSo on\nNon-Functional Testing\nPerformance\nEndurance\nLoad\nVolume\nScalability\nUsability\nSo on\nMaintenance\nRegression\nMaintenance\nThis is not the complete list as there are more than 150 types of testing types and still adding. Also, note that not all testing types are applicable to all projects but depend on the nature \u0026 scope of the project. To explore a variety of testing tools and find the ones that suit your project requirements, visit this list of testing tools .\nTesting Strategies in Software Engineering\nHere are important strategies in software engineering:\nUnit Testing: This software testing basic approach is followed by the programmer to test the unit of the program. It helps developers to know whether the individual unit of the code is working properly or not.\nIntegration testing: It focuses on the construction and design of the software. You need to see that the integrated units are working without errors or not.\nSystem testing: In this method, your software is compiled as a whole and then tested as a whole. This testing strategy checks the functionality, security, portability, amongst others.\nProgram Testing\nProgram Testing in software testing is a method of executing an actual software program with the aim of testing program behavior and finding errors. The software program is executed with test case data to analyse the program behavior or response to the test data. A good program testing is one which has high chances of finding bugs.\nSummary of Software Testing Basics\nDefine Software Testing: Software testing is defined as an activity to check whether the actual results match the expected results and to ensure that the software system is Defect free.\nTesting is important because software bugs could be expensive or even dangerous.\nThe important reasons for using software testing are: cost-effective, security, product quality, and customer satisfaction.\nTypically Testing is classified into three categories functional testing , non-functional testing or performance testing, and maintenance.\nThe important strategies in software engineering are: unit testing, integration testing, validation testing, and system testing.\nYou Might Like:\n7 Principles of Software Testing with Examples\nV-Model in Software Testing\nSTLC (Software Testing Life Cycle) Phases, Entry, Exit Criteria\nManual Testing Tutorial: What is, Types, Concepts\nWhat is Automation Testing? Test Tutorial\nPost navigation\nReport a Bug\nPrevious Prev\nNextContinue\nTop Tutorials\nAbout About Us Advertise with Us Write For Us Contact Us\nPython\nTesting\nHacking\nCareer Suggestion SAP Career Suggestion Tool Software Testing as a Career Interesting eBook Blog Quiz SAP eBook\nSAP\nJava\nSQL\nExecute online Execute Java Online Execute Javascript Execute HTML Execute Python\nSelenium\nBuild Website\nVPNs\n© Copyright - Guru99 2023 Privacy Policy | Affiliate Disclaimer | ToS\nScroll to top\nScroll to top\nHome\nTesting\nExpand child menu Expand\nAgile Testing\nJUnit\nQuality Center(ALM)\nBugzilla\nHP Loadrunner\nRPA\nCucumber\nSoftware Testing\nSAP Testing\nDatabase Testing\nMobile Testing\nSelenium\nETL Testing\nMantis\nSoapUI\nJMeter\nPostman\nTEST Management\nJIRA\nQTP\nTestLink\nSAP\nExpand child menu Expand\nABAP\nCRM\nPI/PO\nAPO\nCrystal Reports\nPP\nBeginners\nFICO\nSD\nBasis\nHANA\nSAPUI5\nBODS\nHR\nSecurity Tutorial\nBI/BW\nMM\nSolution Manager\nBPC\nQM\nSuccessfactors\nCO\nPayroll\nSAP Courses\nWeb\nExpand child menu Expand\nApache\nJava\nPHP\nSQL Server\nAngularJS\nJSP\nPL/SQL\nUML\nASP.NET\nKotlin\nPostgreSQL\nVB.NET\nC\nLinux\nPython\nVBScript\nC#\nMariaDB\nReactJS\nWeb Services\nC++\nMS Access\nRuby \u0026 Rails\nWPF\nCodeIgniter\nMySQL\nScala\nSQLite\nDBMS\nNode.js\nSQL\nPerl\nJavaScript\nMust Learn\nExpand child menu Expand\nAccounting\nEmbedded Systems\nOperating System\nAlgorithms\nEthical Hacking\nPMP\nAndroid\nExcel Tutorial\nPhotoshop\nBlockchain\nGo Programming\nProject Management\nBusiness Analyst\nIoT\nReviews\nBuild Website\nITIL\nSalesforce\nCloud Computing\nJenkins\nSEO\nCOBOL\nMIS\nSoftware Engineering\nCompiler Design\nMovie\nVBA\nCourses\nNetworking\nVPN\nBig Data\nExpand child menu Expand\nAWS\nHive\nPower BI\nBig Data\nInformatica\nQlikview\nCassandra\nMicroStrategy\nTableau\nCognos\nMongoDB\nTalend\nData Warehousing\nNiFi\nZooKeeper\nDevOps\nOBIEE\nPentaho\nHBase\nLive Project\nExpand child menu Expand\nLive Agile Testing\nLive Selenium Project\nLive HP ALM\nLive Selenium 2\nLive Java Project\nLive Security Testing\nLive Mobile Testing\nLive Testing Project\nLive Payment Gateway\nLive Testing 2\nLive PHP Project\nLive Telecom\nLive Projects Hub\nLive UFT/QTP Testing\nLive Python Project\nLive SEO Project\nAI\nExpand child menu Expand\nArtificial Intelligence\nPyTorch\nData Science\nR Programming\nKeras\nTensorFlow\nNLTK\nToggle Menu Close\nSearch for:\nSearch"` {
		t.Errorf("incorrect html text extracted")
	}
}
