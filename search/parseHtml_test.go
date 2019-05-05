package search

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

const searchResult = `    <!DOCTYPE html>
    <html data-pagecluster="Suchergebnisseite" class="">
    <head>
        <title>Suchergebnis für waschmaschiene | OTTO</title>
        <meta charset="utf-8">
        <meta http-equiv="content-language" content="de">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <meta name="robots" content="noindex,nofollow"/>

                <meta name="head_static_files">
<script>
  "use strict";

  var o_global = o_global || {};

  o_global.toggles = Object.freeze({
      "EMR_PXC": true,
      "RUM_PageCluster": true,
      "RUM_Viewport": true,
      "EMR_MPATHY": true,
      "FT6_CALL_CORE_WITH_DRESON_RULE": true,
      "RUM_CustomEvents": true,
      "SCALE_SendUserTimings": true,
      "RUM_Visibility": false,
      "LHAS_2048_TRACK_COOKIES": false,
      "RUM_RoundTripTime": true,
      "FT6_CAMPAIGN_ENABLE_ASSETS_DELIVERY": true,
      "RUM_DetailedPerformance": true,
      "RUM_BasicInformation": true,
      "RUM_AssetX": true,
      "EMR_WEBTRENDS": true,
      "RUM_DeviceInformation": true,
      "RUM_ASSETS_45_SEND_TO_SCALE": true,
      "RUM_JavascriptErrors": true,
      "RUM_ScreenSize": false,
      "EMR_RUM": true,
      "RUM_Connection": false,
      "RUM_ImportantPerformance": true,
      "SCALE_SendRumData": true
  });
</script>

<!--[if IE 9]><script src="/assets-static/global-resources/assets.global-resources.js-cc-ie9.5328efcc.js" integrity="sha256-X+X3a9e6SoxI2GjOBbgaDVXq+dIiF6xqyfWod+nZk+U=" crossorigin="anonymous"></script><![endif]-->
<script>!function(){"use strict";var e=function(e){e.setAttribute("rel","stylesheet"),e.setAttribute("type","text/css"),e.setAttribute("media","all"),e.removeAttribute("as")},t=function(e){var t=arguments.length>1&&void 0!==arguments[1]&&arguments[1];e.setAttribute("preloaded",t?"error":"true"),e.removeEventListener("load",window.invokePreload.onLoad),e.removeAttribute("onload"),e.removeAttribute("onerror"),e.onload=null};!function(){try{new Function("(a = 0) => a")}catch(e){return!1}}();window.invokePreload=window.invokePreload||{},invokePreload.onLoad=t,invokePreload.onScriptLoad=t,invokePreload.onScriptError=function(e){return t(e,!0)},invokePreload.onStyleLoad=function(t){return-1===[].map.call(document.styleSheets,function(e){try{return"all"===e.media.mediaText?e.href:null}catch(e){return null}}).indexOf(t.href)&&(window.requestAnimationFrame?window.requestAnimationFrame(function(){return e(t)}):e(t)),t.removeAttribute("onload"),t}}();

</script><link rel="stylesheet" href="/assets-static/global-resources/assets.global-resources.head.84c9fe0d.css" integrity="sha256-94EiJS3Bwu5zRnxvxPVMZ3H0h2Go38vw1YeI/uyX25g=" crossorigin="anonymous"/>
<script src="/assets-static/global-resources/assets.global-resources.head.6de51c72.js" integrity="sha256-nXgoE2VtVbCZvQtREg3dfURE9XyDoq6AbcaaFUGOmjc=" crossorigin="anonymous"></script><link href="/assets-static/global-resources/assets.global-resources.fonts.d3d1351b.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)"/><noscript><link rel="stylesheet" href="/assets-static/global-resources/assets.global-resources.fonts.d3d1351b.css" crossorigin="anonymous"/></noscript>
<link rel="preload" name="OttoSans" as="font" weight="400" crossorigin="anonymous" type="font/woff2" href="/assets-static/global-resources/assets/otto-sans_88e5e760.woff2"/>
<link rel="preload" name="OttoSans" as="font" weight="700" crossorigin="anonymous" type="font/woff2" href="/assets-static/global-resources/assets/otto-sans-bold_89d38c21.woff2"/>
<link rel="preload" name="OttoSansThin" as="font" weight="400" crossorigin="anonymous" type="font/woff2" href="/assets-static/global-resources/assets/otto-sans-thin_be4e47f1.woff2"/>
<link rel="preload" name="OttoIconFonts" as="font" weight="400" crossorigin="anonymous" type="font/woff2" href="/assets-static/global-resources/assets/OTTO-Icons_120f8d51.woff2"/>
<link rel="preload" name="OttoIcons" as="font" weight="400" crossorigin="anonymous" type="font/woff2" href="/assets-static/global-resources/assets/OTTO-Icons-v2_4953e1b.woff2"/>
<link href="/assets-static/global-pattern/assets.global-pattern.main.64bff9b5.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)"/><noscript><link rel="stylesheet" href="/assets-static/global-pattern/assets.global-pattern.main.64bff9b5.css" crossorigin="anonymous"/></noscript>
<script src="/assets-static/global-pattern/assets.global-pattern.main.f52deb74.js" integrity="sha256-O5tJ1RvGMHLPQOXAy1MYlj+ympHN3FY+5gf7nWERe7Q=" crossorigin="anonymous"></script>
<link rel="shortcut icon" type="image/x-icon" href="/assets-static/global-favicons/favicon.ico">
<link rel="icon" type="image/x-icon" href="/assets-static/global-favicons/favicon.ico">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon.png">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-57x57.png" sizes="57x57">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-60x60.png" sizes="60x60">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-72x72.png" sizes="72x72">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-76x76.png" sizes="76x76">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-114x114.png" sizes="114x114">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-120x120.png" sizes="120x120">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-128x128.png" sizes="128x128">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-144x144.png" sizes="144x144">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-152x152.png" sizes="152x152">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-180x180.png" sizes="180x180">
<link rel="apple-touch-icon" href="/assets-static/global-favicons/apple-touch-icon-precomposed.png">
<link rel="icon" type="image/png" href="/assets-static/global-favicons/favicon-16x16.png" sizes="16x16">
<link rel="icon" type="image/png" href="/assets-static/global-favicons/favicon-32x32.png" sizes="32x32">
<link rel="icon" type="image/png" href="/assets-static/global-favicons/favicon-96x96.png" sizes="96x96">
<link rel="icon" type="image/png" href="/assets-static/global-favicons/favicon-160x160.png" sizes="160x160">
<link rel="icon" type="image/png" href="/assets-static/global-favicons/favicon-192x192.png" sizes="192x192">
<link rel="icon" type="image/png" href="/assets-static/global-favicons/favicon-196x196.png" sizes="196x196">
<link rel="manifest" href="/assets-static/global-favicons/manifest.json">
<meta name="msapplication-TileImage" content="/assets-static/global-favicons/ms-icon-144x144.png">
<meta name="msapplication-TileColor" content="#ffffff">
<meta name="msapplication-navbutton-color" content="#d52b1e">
<meta name="msapplication-square70x70logo" content="/assets-static/global-favicons/ms-icon-70x70.png">
<meta name="msapplication-square144x144logo" content="/assets-static/global-favicons/ms-icon-144x144.png">
<meta name="msapplication-square150x150logo" content="/assets-static/global-favicons/ms-icon-150x150.png">
<meta name="msapplication-wide310x150logo" content="/assets-static/global-favicons/ms-icon-310x150.png">
<meta name="msapplication-square310x310logo" content="/assets-static/global-favicons/ms-icon-310x310.png">

  <script type="text/javascript">var o_global=o_global||{};o_global.fonts=o_global.fonts||{};o_global.dirhashes={"js":{"all":"1573a9aac30427f5","head":"4e904ef14c3d1483","thirdparty":"d41d8cd98f00b204","aftersales-order-tracking":"7ea4a16bfd9b90ea","api-authorization":"256609b564b90541","api-awesome-oauth2-app":"23f4ceca286b8310","order":"f046e64c00c9cb4f","san":"ddd8f69445adbfa5","user":"0aef3c01eaa7a22a","user-benefit-offers":"5fc9e421725ae149","user-campaignmanagement":"9cebf12434ba15d7"},"css":{"all":"62b37af2ce11c3ae","aftersales-order-tracking":"16a26c864b88d2b3","aftersales":"fc9df76b4426509c","api-authorization":"a366b6a6500ca271","order":"58afb47fc0cf7329","san":"bf57282c937da1a4","user-benefit-offers":"6d2369610cac42ca","user-campaignmanagement":"112c77a322abc354","user":"55f5eb34b0a279ca"},"img":{"aftersales":"914086475dc9b41e","aftersales-order-tracking":"da39a3ee5e6b4b0d","api-authorization":"da39a3ee5e6b4b0d","api-awesome-oauth2-app":"da39a3ee5e6b4b0d","apps-backend":"da39a3ee5e6b4b0d","global-resources":"9cc2d231044d4951","order":"d4179016602616be","san":"c71f419a039a76b8","san-nav-view":"da39a3ee5e6b4b0d","san-productlister":"da39a3ee5e6b4b0d","san-srch-productfacets":"da39a3ee5e6b4b0d","user":"507a6e392b4ccfc6","user-benefit-offers":"da39a3ee5e6b4b0d","user-campaignmanagement":"da39a3ee5e6b4b0d"}};o_global.fonts.conf={}</script>
  <script type="text/javascript" src="https://www.otto.de/static/all/js/4e904ef14c3d1483/public_head_min.js" crossorigin></script>
<meta name="theme-color" content="#D4021D">
  <link href="https://www.otto.de/static/all/css/62b37af2ce11c3ae/files/public_critical_min.css" rel="stylesheet">
  <link href="https://www.otto.de/static/san/css/bf57282c937da1a4/files/private_critical_min.css" rel="stylesheet">
    </head>
    <body class="san-system" itemscope itemtype="http://schema.org/WebPage">
    <div class="ts-bct" data-ts_sfid="dcbc0c3beacf9cb88bd4175ed422c5d085067498"></div>
    <div class="gridAndInfoContainer">
    <div class="gridContainer reducedOuterPadding wrapper">

    <script type="text/javascript">
        if (!!o_san.searchHandler) {
            o_san.searchHandler.init();
    }
    </script>

    <header class="withSubMenu">
    <div id="sanHeadWrp" >
    <div id="mobilMenuWrp">
        <label id="burgerMenu">
            <p class="p_iconFont">=</p>

            <p>Sortiment</p>
        </label>
    </div>
    <div id="logoWrp">
        <a class="san_logo" href="/" data-tracking="{&quot;san_Header&quot;:&quot;logo&quot;}" title="zur Homepage">
            <svg class="san-svg" width="100%" height="100%" viewBox="0 0 140 52">
                <path d="M121.136,2.75c-6.876,0-12.469,2.12-16.42,6.508c0.562-1.523,0.858-3.381,0.897-5.569H81.517 c-5.258,0-8.263,2.315-10.329,8.262l1.065-8.262H47.905c-4.458,0-7.518,2.131-8.624,6.305C36.264,5.159,30.902,2.75,24.12,2.75 c-12.893,0-21.281,7.449-22.783,23.346L1.15,28.1c-1.314,14.209,6.134,21.219,17.712,21.219c12.895,0,21.282-7.513,22.784-23.41 l0.188-2.002c0.339-3.646,0.091-6.814-0.655-9.515h10.919l-3.129,26.79c-0.626,5.32,2.253,7.761,7.072,7.761 c2.942,0,3.943-0.188,5.07-0.375l3.881-34.176h12.707l-3.131,26.79c-0.625,5.32,2.254,7.761,7.073,7.761 c2.941,0,3.943-0.188,5.069-0.375l3.881-34.176h6.26c2.165,0,3.913-0.503,5.264-1.498c-1.97,3.473-3.257,7.858-3.763,13.203 L98.166,28.1c-1.314,14.209,6.133,21.219,17.712,21.219c12.896,0,21.281-7.513,22.783-23.41l0.188-2.002 C140.164,9.76,132.716,2.75,121.136,2.75z M28.94,23.03l-0.126,1.502c-0.875,10.765-4.381,14.083-8.699,14.083 c-3.881,0-6.634-2.628-6.071-9.638l0.124-1.565c0.877-10.704,4.382-14.083,8.701-14.083C26.687,13.328,29.503,16.02,28.94,23.03z M125.957,23.03l-0.126,1.502c-0.877,10.765-4.381,14.083-8.699,14.083c-3.882,0-6.637-2.628-6.072-9.638l0.125-1.565 c0.876-10.704,4.381-14.083,8.7-14.083C123.702,13.328,126.519,16.02,125.957,23.03z"/>
                <image class="san-svg" src="/san/resources/san/img/header/otto_logo_2015.png"/>
            </svg>
        </a>
    </div>
    <div id="searchAndIconWrp">
            <span id="shoppingLink">&nbsp;</span>

        <div id="searchAndIconBg">
<div class="san_searchField__Wrapper">
    <form class="p_form js_searchForm focus" action="/suche" data-article-number-search="/p/search/" autocomplete="off" autocorrect="off" spellcheck="false" role="search">
        <input placeholder="Suchbegriff / Artikelnr. eingeben" data-error="Bitte mind. ein Zeichen eingeben" class="p_form__input js_searchField sanSearchInput" type="text" autocomplete="off" autocorrect="off" maxlength="50" disabled>
        <button class="sanSearchDelBtn p_symbolBtn100--4th" type="reset"><i>X</i></button>
        <button class="js_submitButton sanSearchButton" type="submit" title="Suche" disabled ><span>&raquo;</span></button>
    </form>

    <div class="san_suggestLayer san_newSuggest"
         data-suggestserveruri="/san-srch-suggest-api/completion"
         data-suggestscope="//catalog01/de_DE"></div>

</div>
            <script type="text/javascript">
                if (!!o_san.searchHandler) {
                    o_san.searchHandler.init();
                }
            </script>
            <ul id="sanIconWrp">
                <li id="searchGlassWrp"><span id="searchGlass" class="p_iconFont">»</span></li>
                <li id="serviceLinkWrp"><a id="serviceLink"
                                           href="/shoppages/service/"
                                           data-tracking="{&quot;san_Header&quot;:&quot;service&quot;}"><span
                        class="p_iconFont">s</span> Service</a></li>
                <li id="loginWrp">
                    
<div class="us_loginAreaContainerWrapper" style="visibility: hidden">
    <div id="us_js_id_loginAreaContainerToReplace" class="us_loginAreaContainerBackground">
        <a class="us_loginAreaFallbackLink" href="/user/accountOverview">
            <span class="p_icons us_loginAreaContainerIcon">Θ</span>
            <span class="us_iconSubtitle">Mein Konto</span>
        </a>
    </div>
<link href="/user/assets/ft4.user.login-area.139d419c.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)"/><noscript><link rel="stylesheet" href="/user/assets/ft4.user.login-area.139d419c.css" crossorigin="anonymous"/></noscript>
<link rel="preload" crossorigin="anonymous" href="/user/assets/ft4.user.login-area.6bb7750d.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/></div>

                </li>
                <li id="favouritesWrp">
                    
<link href="/order-wishlist/statics/ft1.wishlist-core.common.scala.dd7daae8.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)"/><noscript><link rel="stylesheet" href="/order-wishlist/statics/ft1.wishlist-core.common.scala.dd7daae8.css" crossorigin="anonymous"/></noscript>
<link rel="preload" critical="critical" crossorigin="anonymous" href="/order-wishlist/statics/ft1.wishlist-core.common.scala.3932c7a9.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/>
<link rel="stylesheet" href="/order-wishlist/statics/ft1.wishlist-core.miniWishlist.scala.074d6195.css" integrity="sha256-J5HxY2LDZcu7KvgTZJ0gyZDm18y51sW8dcyhTVGdLPA=" crossorigin="anonymous"/>
<link rel="preload" crossorigin="anonymous" href="/order-wishlist/statics/ft1.wishlist-core.miniWishlist.scala.9fc03450.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/>

    <div class="wl_mini wl_js_mini_link ts-link ub64e"
    data-ub64e="L29yZGVyLXdpc2hsaXN0L3dpc2hsaXN0L3N5c3RlbQ=="
    data-ts-link='{"san_Header": "wishlist", "order_MerkzettelAufruf": "header"}'
    title="Mein Merkzettel">
        <div class="wl_toggles" data-toggles="MIGRATE_VARIATION_IDS_FROM_CSV_ON_THE_FLY,INPUT_VALIDATION,HMAC_ACTIVE,TEILSTRECKEN_PROFILING"></div>
        <span class="p_icons wl_mini__icon">&hearts;</span>
        <span id="wl_mini_amount" class="wl_mini__badge p_badge100 wl_mini__badge--empty wl_js_mini_amount"
        data-qa="miniWishlistAmount"
        data-loadurl="/order-wishlist/wishlist/amount.json"></span>
        <span class="wl_mini__text">Merkzettel</span>
    </div>

                </li>
                <li id="basketWrp">
                    
<link rel="preload" critical="critical" crossorigin="anonymous" href="/order/statics/ft1.order-core.common-public.68d1bc02.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/><link rel="stylesheet" href="/order/statics/ft1.order-core.minibasket.21c208fd.css" integrity="sha256-U15sc6o2bTB05zp9TWnAYzoEYJK1fNzlvLQBf12biN0=" crossorigin="anonymous"/>
<link rel="preload" crossorigin="anonymous" href="/order/statics/ft1.order-core.minibasket.5eddf6d1.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/><div class="or_minis or_minibasket order_js_minibasket_link ub64e ts-link"
     data-ub64e="L29yZGVyL2Jhc2tldA=="
     data-ts-link='{"san_Header":"basket"}'
     title="Zum Warenkorb">

    <span class="p_icons or_minis__icon">+</span>
    <span class="or_minis__badge or_minis__badge--empty p_badge100 order_js_minibasket_amount"
          data-loadurl="/order/basket/amount.json"
          data-qa="miniBasketAmount"></span>
    <span class="or_minis__text">Warenkorb</span>
</div>


<div id="order_js_settings" data-settings="A_PLUS_DISPOSAL" ></div>
                </li>
            </ul>
        </div>
    </div>
</div>

<div id="navigationWrp" class="innaActivated">
        <nav>
            <ul class="firstRow">
                    <li class="san_navItem inspiration" data-flyout="0">
                        <a href="/inspiration/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;inspiration&quot;}" class="navItem">
                            <span class="js_shortname">Inspiration</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem damen" data-flyout="1">
                        <a href="/damen/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;damen&quot;}" class="navItem">
                            <span class="js_shortname">Damen</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem herren" data-flyout="2">
                        <a href="/herren/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;herren&quot;}" class="navItem">
                            <span class="js_shortname">Herren</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem kinder" data-flyout="3">
                        <a href="/kinder/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;kinder&quot;}" class="navItem">
                            <span class="js_shortname">Kinder</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem waesche-bademode" data-flyout="4">
                        <a href="/mode/waesche-bademode/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;waesche-bademode&quot;}" class="navItem">
                            <span class="js_shortname">Wäsche/Bademode</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem sport" data-flyout="5">
                        <a href="/sport/?ansicht=einstieg" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;sport&quot;}" class="navItem">
                            <span class="js_shortname">Sport</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem schuhe" data-flyout="6">
                        <a href="/mode/schuhe/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;schuhe&quot;}" class="navItem">
                            <span class="js_shortname">Schuhe</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem grosse-groessen" data-flyout="7">
                        <a href="/mode/?grossegroesse" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;grosse-groessen&quot;}" class="navItem">
                            <span class="js_shortname">Große Größen</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    <li class="split-line"></li>
                    
                    <li class="san_navItem multimedia" data-flyout="8">
                        <a href="/technik/multimedia/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;multimedia&quot;}" class="navItem">
                            <span class="js_shortname">Multimedia</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem haushalt" data-flyout="9">
                        <a href="/haushalt/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;haushalt&quot;}" class="navItem">
                            <span class="js_shortname">Haushalt</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem kueche" data-flyout="10">
                        <a href="/moebel/?ansicht=einstieg&amp;thema=kueche" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;kueche&quot;}" class="navItem">
                            <span class="js_shortname">Küche</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem moebel" data-flyout="11">
                        <a href="/moebel/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;moebel&quot;}" class="navItem">
                            <span class="js_shortname">Möbel</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem heimtextilien" data-flyout="12">
                        <a href="/heimtextilien/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;heimtextilien&quot;}" class="navItem">
                            <span class="js_shortname">Heimtextilien</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem baumarkt" data-flyout="13">
                        <a href="/baumarkt/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;baumarkt&quot;}" class="navItem">
                            <span class="js_shortname">Baumarkt</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem spielzeug" data-flyout="14">
                        <a href="/spielzeug/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;spielzeug&quot;}" class="navItem">
                            <span class="js_shortname">Spielzeug</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem marken" data-flyout="15">
                        <a href="/marken/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;marken&quot;}" class="navItem">
                            <span class="js_shortname">Marken</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    <li class="san_square" data-dot="."></li>
                    <li class="san_navItem sale" data-flyout="16">
                        <a href="/sale/" data-tracking="{&quot;san_Navigation&quot;:&quot;global&quot;, &quot;wk.nav_GlobalNavigation&quot;:&quot;sale&quot;}" class="navItem">
                            <span class="js_shortname">%Sale%</span>
                            <div class="overState"></div>
                        </a>
                        <div class="flyOut shadow">
                        </div>
                    </li>
                    
                    
            </ul>
        </nav>
        <div id="flyout-data" data-url="/san/menuitems/views/flyout2"
              data-trackingkey="{&quot;san_Navigation&quot;:&quot;global&quot;}"></div>
</div>

    </header>
    <div class="content contentWithSidebar">
    




    <div class="subnavWrapper">
        <table class="subnav">
            <tr>
                <td class="breadcrumb " itemprop="breadcrumb">
    <span class="js_san_dynamicBreadcrumb dynamicBreadcrumb"
          data-dynamic-breadcrumb-url="/san/breadcrumbs/external/?referrer=">
                <span class="item" itemscope itemtype="http://data-vocabulary.org/Breadcrumb">
                    <a href="/" itemprop="url"
                       data-tracking='{"san_Navigation":"Breadcrumb"}'>
                        <span class="itemTitle" itemprop="title">Startseite</span>
                    </a>
                </span>
        </span>
    <span class="item">
            <span class="divide">|</span>
            <h1>Suchergebnis für waschmaschiene</h1>
        <strong>(432)</strong>
        </span>
                </td>
                    <td class="paging"><div class="san_paging">
<ul class="san_paging_block">
<li id="san_pageInfo" class="san_paging__item"><span class="san_paging__btn">Seite 1 von 15</span></li>
<li id="san_pagingTopNext" class="san_paging__item"><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wPTImcHM9MzA=" data-tracking='{"san_NaviPaging":"2"}'> <p class="p_btn50--4th san_paging__btn" title="weiter"><i>&gt;</i></p> </span></li> </ul>
</div></td>
            </tr>
        </table>
    </div>

        <form class="san_filterToggle">
            <label for="san-filterToggleButton" class="san_filterToggle__label">Auswahl verfeinern:</label>
            <button id="san-filterToggleButton" type="button"
                    class="san_filterToggle__button san-sidebar san-filter-container"><span
                    class="san_filterToggle__caption">Filtern</span></button>
        </form>


    <div class="sanSidebarWrapper sanNavAccordion san_filterToggle__content">
        <dl class="san-sidebar p_accordion100 p_accordion--1st">
            <dd class="san-sidebar p_accordion__content">
                    <div class="filterWrapper">
                        <dl class="p_accordion100 p_accordion--3rd">









        <dt class="p_accordion__header p_accordion__header--open"
            id="facet-assortmentsearch"
            data-title="assortmentsearch"
            data-shortname="assortmentsearch">
                Sortiment
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;assortmentsearch&quot;}">
        <div class="san-scroll-wrapper">
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPXRlY2huaWs=" data-merge-tracking='{"san_FilterValue":"technik"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="374">Technik</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPWluc3RhbGxhdGlvbg==" data-merge-tracking='{"san_FilterValue":"installation"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="27">Installation</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPWt1ZWNoZW5hdXNzdGF0dHVuZw==" data-merge-tracking='{"san_FilterValue":"kuechenausstattung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="20">Küchenausstattung</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPWJla2xlaWR1bmc=" data-merge-tracking='{"san_FilterValue":"bekleidung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="10">Bekleidung</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPWF1ZmJld2FocnVuZw==" data-merge-tracking='{"san_FilterValue":"aufbewahrung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="8">Aufbewahrung</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPXNjaHV0eg==" data-merge-tracking='{"san_FilterValue":"schutz"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="6">Schutz</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPWF1c3J1ZXN0dW5n" data-merge-tracking='{"san_FilterValue":"ausruestung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="5">Ausrüstung</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPW1vZWJlbA==" data-merge-tracking='{"san_FilterValue":"moebel"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="5">Möbel</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPWhlaW10ZXh0aWxpZW4=" data-merge-tracking='{"san_FilterValue":"heimtextilien"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="4">Heimtextilien</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPXNwaWVsemV1Zw==" data-merge-tracking='{"san_FilterValue":"spielzeug"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Spielzeug</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPXBmbGVnZW1pdHRlbA==" data-merge-tracking='{"san_FilterValue":"pflegemittel"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Pflegemittel</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPWFybWF0dXJlbg==" data-merge-tracking='{"san_FilterValue":"armaturen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Armaturen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hc3NvcnRtZW50c2VhcmNoPXNjaHJlaWJ3YXJlbg==" data-merge-tracking='{"san_FilterValue":"schreibwaren"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Schreibwaren</span> </span> </span>    </li>
            </ul>
        </div>
        </dd>
        <dt class="p_accordion__header p_accordion__header--open"
            id="facet-kategorie"
            data-title="kategorie"
            data-shortname="kategorie">
                Kategorie
        </dt>
        <dd class="p_accordion__content facet categoryAll"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;kategorie&quot;}">
        <div class="san-scroll-wrapper">
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9d2FzY2htYXNjaGluZW4=" data-merge-tracking='{"san_FilterValue":"waschmaschinen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="371">Waschmaschinen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9dHVlY2hlcg==" data-merge-tracking='{"san_FilterValue":"tuecher"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="20">Tücher</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9d2FzY2htYXNjaGluZW5zY2hsYWV1Y2hl" data-merge-tracking='{"san_FilterValue":"waschmaschinenschlaeuche"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="20">Waschmaschinenschläuche</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9Z2VzY2hpcnJzcHVlbGVy" data-merge-tracking='{"san_FilterValue":"geschirrspueler"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="17">Geschirrspüler</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9enViZWhvZXItZnVlci13YXNjaG1hc2NoaW5lbg==" data-merge-tracking='{"san_FilterValue":"zubehoer-fuer-waschmaschinen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="12">Zubehör für Waschmaschinen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9dHJvY2tuZXI=" data-merge-tracking='{"san_FilterValue":"trockner"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="11">Trockner</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9cHV0em1pdHRlbA==" data-merge-tracking='{"san_FilterValue":"putzmittel"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="6">Putzmittel</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9d2FzY2htYXNjaGluZW4tdW50ZXJnZXN0ZWxsZQ==" data-merge-tracking='{"san_FilterValue":"waschmaschinen-untergestelle"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="6">Waschmaschinen-Untergestelle</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9bmV0emU=" data-merge-tracking='{"san_FilterValue":"netze"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="4">Netze</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9YmFkbW9lYmVs" data-merge-tracking='{"san_FilterValue":"badmoebel"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Badmöbel</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9a2lzc2Vu" data-merge-tracking='{"san_FilterValue":"kissen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Kissen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9a3VlaGxzY2hyYWVua2U=" data-merge-tracking='{"san_FilterValue":"kuehlschraenke"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Kühlschränke</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9cHVsbG92ZXI=" data-merge-tracking='{"san_FilterValue":"pullover"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Pullover</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9c3RhZW5kZXI=" data-merge-tracking='{"san_FilterValue":"staender"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Ständer</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9dHJvY2tuZXItdW50ZXJnZXN0ZWxsZQ==" data-merge-tracking='{"san_FilterValue":"trockner-untergestelle"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Trockner-Untergestelle</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9d2FzY2h0cm9ja25lcg==" data-merge-tracking='{"san_FilterValue":"waschtrockner"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Waschtrockner</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9d2Flc2NoZXNhbW1sZXI=" data-merge-tracking='{"san_FilterValue":"waeschesammler"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Wäschesammler</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9d2Flc2NoZXNvcnRpZXJlcg==" data-merge-tracking='{"san_FilterValue":"waeschesortierer"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Wäschesortierer</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9d2Flc2NoZXN0YWVuZGVy" data-merge-tracking='{"san_FilterValue":"waeschestaender"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Wäscheständer</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9enViZWhvZXItZnVlci1rdWVobHNjaHJhZW5rZQ==" data-merge-tracking='{"san_FilterValue":"zubehoer-fuer-kuehlschraenke"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Zubehör für Kühlschränke</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9Ymhz" data-merge-tracking='{"san_FilterValue":"bhs"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">BHs</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9Z2VmcmllcnNjaHJhZW5rZQ==" data-merge-tracking='{"san_FilterValue":"gefrierschraenke"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Gefrierschränke</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9Z2VzY2hlbmthcnRpa2Vs" data-merge-tracking='{"san_FilterValue":"geschenkartikel"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Geschenkartikel</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9amFja2Vu" data-merge-tracking='{"san_FilterValue":"jacken"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Jacken</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9a2luZGVyLWhhdXNoYWx0c2dlcmFldGU=" data-merge-tracking='{"san_FilterValue":"kinder-haushaltsgeraete"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Kinder-Haushaltsgeräte</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9cmVnYWxl" data-merge-tracking='{"san_FilterValue":"regale"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Regale</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9c2NocmFlbmtl" data-merge-tracking='{"san_FilterValue":"schraenke"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Schränke</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9c2hpcnRz" data-merge-tracking='{"san_FilterValue":"shirts"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Shirts</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9d2FzY2htYXNjaGluZQ==" data-merge-tracking='{"san_FilterValue":"waschmaschine"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Waschmaschine</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9enViZWhvZXItZnVlci1nZWZyaWVyc2NocmFlbmtl" data-merge-tracking='{"san_FilterValue":"zubehoer-fuer-gefrierschraenke"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Zubehör für Gefrierschränke</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9YmFieWJlZGFyZg==" data-merge-tracking='{"san_FilterValue":"babybedarf"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Babybedarf</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9YmFkYXJtYXR1cmVu" data-merge-tracking='{"san_FilterValue":"badarmaturen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Badarmaturen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9YmV0dHdhZXNjaGU=" data-merge-tracking='{"san_FilterValue":"bettwaesche"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Bettwäsche</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9ZGVja2Vu" data-merge-tracking='{"san_FilterValue":"decken"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Decken</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9ZXRpa2V0dGVu" data-merge-tracking='{"san_FilterValue":"etiketten"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Etiketten</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9Zml0bmVzc2dlcmFldGU=" data-merge-tracking='{"san_FilterValue":"fitnessgeraete"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Fitnessgeräte</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9aGFuZHR1ZWNoZXI=" data-merge-tracking='{"san_FilterValue":"handtuecher"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Handtücher</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9aG9zZW4=" data-merge-tracking='{"san_FilterValue":"hosen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Hosen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9a2F0emVuYmVkYXJm" data-merge-tracking='{"san_FilterValue":"katzenbedarf"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Katzenbedarf</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9a2luZGVyc2l0emU=" data-merge-tracking='{"san_FilterValue":"kindersitze"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Kindersitze</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9a2xlaWRlcmF1ZmJld2FocnVuZw==" data-merge-tracking='{"san_FilterValue":"kleideraufbewahrung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Kleideraufbewahrung</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9a2xlaWRlcnNhZWNrZQ==" data-merge-tracking='{"san_FilterValue":"kleidersaecke"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Kleidersäcke</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9c2FuaXRhZXJ0ZWNobmlr" data-merge-tracking='{"san_FilterValue":"sanitaertechnik"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Sanitärtechnik</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9c3BpZWx6ZXVnZmlndXJlbg==" data-merge-tracking='{"san_FilterValue":"spielzeugfiguren"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Spielzeugfiguren</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9dHJvY2tuZXJmaXhpZXJwbGF0dGVu" data-merge-tracking='{"san_FilterValue":"trocknerfixierplatten"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Trocknerfixierplatten</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9rYXRlZ29yaWU9enViZWhvZXItZnVlci1jYW1waW5n" data-merge-tracking='{"san_FilterValue":"zubehoer-fuer-camping"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Zubehör für Camping</span> </span> </span>    </li>
            </ul>
        </div>
        </dd>
        <dt class="p_accordion__header"
            id="facet-produkttyp"
            data-title="produkttyp"
            data-shortname="produkttyp">
                Produkttyp
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;produkttyp&quot;}">
        <div class="san-scroll-wrapper">
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWZyb250bGFkZXI=" data-merge-tracking='{"san_FilterValue":"frontlader"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="265">Frontlader</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXRvcGxhZGVy" data-merge-tracking='{"san_FilterValue":"toplader"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="34">Toplader</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPW11bHRpZnVua3Rpb25zdHVlY2hlcg==" data-merge-tracking='{"san_FilterValue":"multifunktionstuecher"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="20">Multifunktionstücher</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXp3aXNjaGVuYmF1c2FldHpl" data-merge-tracking='{"san_FilterValue":"zwischenbausaetze"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="17">Zwischenbausätze</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXVudGVyYmF1YmxlY2hl" data-merge-tracking='{"san_FilterValue":"unterbaubleche"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="11">Unterbaubleche</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWdlc2NoaXJyc3B1ZWxlcnNjaGxhZXVjaGU=" data-merge-tracking='{"san_FilterValue":"geschirrspuelerschlaeuche"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="10">Geschirrspülerschläuche</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXdhZXNjaGVuZXR6ZQ==" data-merge-tracking='{"san_FilterValue":"waeschenetze"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="6">Wäschenetze</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWVpbmJhdXdhc2NobWFzY2hpbmVu" data-merge-tracking='{"san_FilterValue":"einbauwaschmaschinen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="5">Einbauwaschmaschinen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXZpYnJhdGlvbnNkYWVtcGZlcg==" data-merge-tracking='{"san_FilterValue":"vibrationsdaempfer"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Vibrationsdämpfer</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXdhbmR0cm9ja25lcg==" data-merge-tracking='{"san_FilterValue":"wandtrockner"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Wandtrockner</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWJoLWFjY2Vzc29pcmVz" data-merge-tracking='{"san_FilterValue":"bh-accessoires"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">BH-Accessoires</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWJhZC1ob2Noc2NocmFlbmtl" data-merge-tracking='{"san_FilterValue":"bad-hochschraenke"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Bad-Hochschränke</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWVudGthbGtlcg==" data-merge-tracking='{"san_FilterValue":"entkalker"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Entkalker</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWtpbmRlci13YXNjaG1hc2NoaW5l" data-merge-tracking='{"san_FilterValue":"kinder-waschmaschine"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Kinder-Waschmaschine</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXJlaXNla2lzc2Vu" data-merge-tracking='{"san_FilterValue":"reisekissen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Reisekissen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXN0cmlja2phY2tlbg==" data-merge-tracking='{"san_FilterValue":"strickjacken"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Strickjacken</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXYtcHVsbG92ZXI=" data-merge-tracking='{"san_FilterValue":"v-pullover"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">V-Pullover</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXdhc3Nlcm1lbGRlcg==" data-merge-tracking='{"san_FilterValue":"wassermelder"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Wassermelder</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWJhZHJlZ2FsZQ==" data-merge-tracking='{"san_FilterValue":"badregale"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Badregale</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWJhdW13b2xsYmV0dHdhZXNjaGU=" data-merge-tracking='{"san_FilterValue":"baumwollbettwaesche"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Baumwollbettwäsche</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWJhdW13b2xsZGVja2Vu" data-merge-tracking='{"san_FilterValue":"baumwolldecken"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Baumwolldecken</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWNhcmRpZ2Fucw==" data-merge-tracking='{"san_FilterValue":"cardigans"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Cardigans</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWNhcmdvaG9zZW4=" data-merge-tracking='{"san_FilterValue":"cargohosen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Cargohosen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWRla29raXNzZW4=" data-merge-tracking='{"san_FilterValue":"dekokissen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Dekokissen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWRla29raXNzZW4tdW5p" data-merge-tracking='{"san_FilterValue":"dekokissen-uni"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Dekokissen uni</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWhhbmR0dWVjaGVyLXBhY2t1bmc=" data-merge-tracking='{"san_FilterValue":"handtuecher-packung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Handtücher (Packung)</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWhhdXNoYWx0c3JlZ2FsZQ==" data-merge-tracking='{"san_FilterValue":"haushaltsregale"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Haushaltsregale</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWthdHplbmhvZWhsZW4=" data-merge-tracking='{"san_FilterValue":"katzenhoehlen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Katzenhöhlen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWtlbm56ZWljaG51bmdzLWV0aWtldHRlbg==" data-merge-tracking='{"san_FilterValue":"kennzeichnungs-etiketten"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Kennzeichnungs-Etiketten</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWt1ZWNoZW5hdWZiZXdhaHJ1bmc=" data-merge-tracking='{"san_FilterValue":"kuechenaufbewahrung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Küchenaufbewahrung</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWxhbmdhcm1zaGlydA==" data-merge-tracking='{"san_FilterValue":"langarmshirt"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Langarmshirt</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPWxhbmdhcm1zaGlydHM=" data-merge-tracking='{"san_FilterValue":"langarmshirts"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Langarmshirts</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXBlcmthbC1iZXR0d2Flc2NoZQ==" data-merge-tracking='{"san_FilterValue":"perkal-bettwaesche"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Perkal-Bettwäsche</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXJvbGxlbnRyYWluZXI=" data-merge-tracking='{"san_FilterValue":"rollentrainer"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Rollentrainer</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXJ1bmRoYWxzcHVsbG92ZXI=" data-merge-tracking='{"san_FilterValue":"rundhalspullover"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Rundhalspullover</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXJ1bmRoYWxzc2hpcnRz" data-merge-tracking='{"san_FilterValue":"rundhalsshirts"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Rundhalsshirts</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXJ1ZWNrZW5sZWhuZW50YXNjaGVu" data-merge-tracking='{"san_FilterValue":"rueckenlehnentaschen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Rückenlehnentaschen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXN0ZWNrcmVnYWxl" data-merge-tracking='{"san_FilterValue":"steckregale"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Steckregale</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXN0cmFuZHNoaXJ0cw==" data-merge-tracking='{"san_FilterValue":"strandshirts"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Strandshirts</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXQtc2hpcnRz" data-merge-tracking='{"san_FilterValue":"t-shirts"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">T-Shirts</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXRlbGVza29wcmVnYWxl" data-merge-tracking='{"san_FilterValue":"teleskopregale"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Teleskopregale</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXRyYW5zcG9ydHJvbGxlcg==" data-merge-tracking='{"san_FilterValue":"transportroller"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Transportroller</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXdpY2tlbGF1ZnNhZXR6ZQ==" data-merge-tracking='{"san_FilterValue":"wickelaufsaetze"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Wickelaufsätze</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXdpY2tlbGJyZXR0ZXI=" data-merge-tracking='{"san_FilterValue":"wickelbretter"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Wickelbretter</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wcm9kdWt0dHlwPXdpY2tlbHRpc2NoZQ==" data-merge-tracking='{"san_FilterValue":"wickeltische"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Wickeltische</span> </span> </span>    </li>
            </ul>
        </div>
        </dd>
        <dt class="p_accordion__header"
            id="facet-marke"
            data-title="marke"
            data-shortname="marke">
                Marke
        </dt>
        <dd class="p_accordion__content facet brand js_filteredFacet"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;marke&quot;}">
                <div class="san-filter-search">
                    <form class="p_form" name="Marke_form">
                        <input class="p_form__input" type="text" maxlength="50" data-defaultValue="Marke suchen"/>
                        <span class="p_iconFont">»</span>
                    </form>
                </div>
                <div class="san-scroll-wrapper clickable">
                <ul class="san_facetValues san_facetValues--checkbox js_filteredFacetValues">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT14YXZheA==" data-merge-tracking='{"san_FilterValue":"xavax"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="42">Xavax</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1hZWc=" data-merge-tracking='{"san_FilterValue":"aeg"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="40">AEG</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1zaWVtZW5z" data-merge-tracking='{"san_FilterValue":"siemens"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="39">SIEMENS</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1taWVsZQ==" data-merge-tracking='{"san_FilterValue":"miele"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="36">Miele</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1ib3NjaA==" data-merge-tracking='{"san_FilterValue":"bosch"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="33">BOSCH</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1iYXVrbmVjaHQ=" data-merge-tracking='{"san_FilterValue":"bauknecht"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="31">BAUKNECHT</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1zYW1zdW5n" data-merge-tracking='{"san_FilterValue":"samsung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="24">Samsung</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1yb2xlZmY=" data-merge-tracking='{"san_FilterValue":"roleff"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="20">roleff</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1sZw==" data-merge-tracking='{"san_FilterValue":"lg"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="18">LG</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1iZWtv" data-merge-tracking='{"san_FilterValue":"beko"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="16">BEKO</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1oYW5zZWF0aWM=" data-merge-tracking='{"san_FilterValue":"hanseatic"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="16">Hanseatic</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1wcml2aWxlZw==" data-merge-tracking='{"san_FilterValue":"privileg"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="12">Privileg</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1nb3Jlbmpl" data-merge-tracking='{"san_FilterValue":"gorenje"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="10">GORENJE</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1hbWljYQ==" data-merge-tracking='{"san_FilterValue":"amica"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="9">Amica</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1leWNraGF1cy1iYXRoLXJlbGF4aW5n" data-merge-tracking='{"san_FilterValue":"eyckhaus-bath-relaxing"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="6">Eyckhaus bath \& relaxing</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1ncnVuZGln" data-merge-tracking='{"san_FilterValue":"grundig"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="6">Grundig</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1oYWllcg==" data-merge-tracking='{"san_FilterValue":"haier"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="6">Haier</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT13cHJv" data-merge-tracking='{"san_FilterValue":"wpro"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="5">Wpro</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1hZ29u" data-merge-tracking='{"san_FilterValue":"agon"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="5">agon®</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1oaXNlbnNl" data-merge-tracking='{"san_FilterValue":"hisense"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="4">Hisense</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1ob292ZXI=" data-merge-tracking='{"san_FilterValue":"hoover"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="4">Hoover</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1pbmRlc2l0" data-merge-tracking='{"san_FilterValue":"indesit"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="4">Indesit</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1sYXNjYW5h" data-merge-tracking='{"san_FilterValue":"lascana"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Lascana</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1sZWlmaGVpdA==" data-merge-tracking='{"san_FilterValue":"leifheit"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Leifheit</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1zei1tZXRhbGw=" data-merge-tracking='{"san_FilterValue":"sz-metall"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">SZ Metall</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT13ZW5rbw==" data-merge-tracking='{"san_FilterValue":"wenko"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">WENKO</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1jYW5keQ==" data-merge-tracking='{"san_FilterValue":"candy"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Candy</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1lYWdsZS1jcmVlaw==" data-merge-tracking='{"san_FilterValue":"eagle-creek"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Eagle Creek</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1mYW1pbHktZWRpdGlvbg==" data-merge-tracking='{"san_FilterValue":"family-edition"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Family Edition</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1kb2Jhcg==" data-merge-tracking='{"san_FilterValue":"dobar"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">dobar</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1hcml6b25h" data-merge-tracking='{"san_FilterValue":"arizona"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Arizona</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1iYWJ5Z28=" data-merge-tracking='{"san_FilterValue":"babygo"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">BabyGo</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1idWZmYWxv" data-merge-tracking='{"san_FilterValue":"buffalo"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Buffalo</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1jbGVhbm1heHg=" data-merge-tracking='{"san_FilterValue":"cleanmaxx"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">CLEANmaxx</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1kZXNpZ3VhbA==" data-merge-tracking='{"san_FilterValue":"desigual"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Desigual</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1lcG9jaC10cmF1bXdpZXNlbg==" data-merge-tracking='{"san_FilterValue":"epoch-traumwiesen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">EPOCH Traumwiesen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1lbGl0ZQ==" data-merge-tracking='{"san_FilterValue":"elite"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Elite</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1naWdhc2V0" data-merge-tracking='{"san_FilterValue":"gigaset"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Gigaset</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1nb29kLWxpZmU=" data-merge-tracking='{"san_FilterValue":"good-life"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Good Life</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1ob3Rwb2ludA==" data-merge-tracking='{"san_FilterValue":"hotpoint"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Hotpoint</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1rZXNwZXI=" data-merge-tracking='{"san_FilterValue":"kesper"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Kesper</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1rbGVpbg==" data-merge-tracking='{"san_FilterValue":"klein"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Klein</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1ubWM=" data-merge-tracking='{"san_FilterValue":"nmc"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">NMC</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1wYWQ=" data-merge-tracking='{"san_FilterValue":"pad"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">PAD</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1wbGF5Z28=" data-merge-tracking='{"san_FilterValue":"playgo"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Playgo</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1wcm9jb250b3Vy" data-merge-tracking='{"san_FilterValue":"procontour"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Procontour</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT12Y20=" data-merge-tracking='{"san_FilterValue":"vcm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">VCM</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT16d2Vja2Zvcm1hdmVyeQ==" data-merge-tracking='{"san_FilterValue":"zweckformavery"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">ZWECKFORMAVERY</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1oZWluZS1ob21l" data-merge-tracking='{"san_FilterValue":"heine-home"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">heine home</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9tYXJrZT1teS13YWxs" data-merge-tracking='{"san_FilterValue":"my-wall"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">my wall</span> </span> </span>    </li>
                    <li class="notFound hidden">keine Treffer</li>
                </ul>
                </div>
        </dd>
        <dt class="p_accordion__header"
            id="facet-farbe"
            data-title="farbe"
            data-shortname="farbe">
                Farbe
        </dt>
        <dd class="p_accordion__content facet color farbe"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;farbe&quot;}">
            <ul class="p_selTiles100 p_selTiles100--checkmark san_filter_farbe">
                    <li class="colorThumb weiß js_hasPaliTooltip "
                        data-tooltip="weiß <span>(615)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT13ZWlzcw==" data-merge-tracking='{"san_FilterValue":"weiss"}'></span>                    </li>
                    <li class="colorThumb blau js_hasPaliTooltip "
                        data-tooltip="blau <span>(56)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1ibGF1" data-merge-tracking='{"san_FilterValue":"blau"}'></span>                    </li>
                    <li class="colorThumb grau js_hasPaliTooltip "
                        data-tooltip="grau <span>(41)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1ncmF1" data-merge-tracking='{"san_FilterValue":"grau"}'></span>                    </li>
                    <li class="colorThumb rot js_hasPaliTooltip "
                        data-tooltip="rot <span>(35)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1yb3Q=" data-merge-tracking='{"san_FilterValue":"rot"}'></span>                    </li>
                    <li class="colorThumb schwarz js_hasPaliTooltip "
                        data-tooltip="schwarz <span>(35)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1zY2h3YXJ6" data-merge-tracking='{"san_FilterValue":"schwarz"}'></span>                    </li>
                    <li class="colorThumb grün js_hasPaliTooltip "
                        data-tooltip="grün <span>(29)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1ncnVlbg==" data-merge-tracking='{"san_FilterValue":"gruen"}'></span>                    </li>
                    <li class="colorThumb natur js_hasPaliTooltip "
                        data-tooltip="natur <span>(19)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1uYXR1cg==" data-merge-tracking='{"san_FilterValue":"natur"}'></span>                    </li>
                    <li class="colorThumb silberfarben js_hasPaliTooltip "
                        data-tooltip="silberfarben <span>(18)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1zaWxiZXJmYXJiZW4=" data-merge-tracking='{"san_FilterValue":"silberfarben"}'></span>                    </li>
                    <li class="colorThumb orange js_hasPaliTooltip "
                        data-tooltip="orange <span>(17)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1vcmFuZ2U=" data-merge-tracking='{"san_FilterValue":"orange"}'></span>                    </li>
                    <li class="colorThumb bunt js_hasPaliTooltip "
                        data-tooltip="bunt <span>(11)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1idW50" data-merge-tracking='{"san_FilterValue":"bunt"}'></span>                    </li>
                    <li class="colorThumb braun js_hasPaliTooltip "
                        data-tooltip="braun <span>(4)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1icmF1bg==" data-merge-tracking='{"san_FilterValue":"braun"}'></span>                    </li>
                    <li class="colorThumb rosa js_hasPaliTooltip "
                        data-tooltip="rosa <span>(4)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1yb3Nh" data-merge-tracking='{"san_FilterValue":"rosa"}'></span>                    </li>
                    <li class="colorThumb gelb js_hasPaliTooltip "
                        data-tooltip="gelb <span>(1)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1nZWxi" data-merge-tracking='{"san_FilterValue":"gelb"}'></span>                    </li>
                    <li class="colorThumb lila js_hasPaliTooltip "
                        data-tooltip="lila <span>(1)</span>"
                        data-tooltip-touch="false" data-tooltip-pos="mouse">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9mYXJiZT1saWxh" data-merge-tracking='{"san_FilterValue":"lila"}'></span>                    </li>
            </ul>
        </dd>
<dt class="p_accordion__header"
id="facet-groesse"
data-title="groesse"
data-shortname="groesse">
Größe
</dt>
<dd class="p_accordion__content facet san_sizeFilter"
>
<select name="sizeGroups" class="san_sizeGroups sizeGroups">
<option value="kindernormalgroessen"
>Kindernormalgrößen</option>
<option value="normalgroessen"
>Normalgrößen</option>
<option value="us-groessen"
selected="selected">US-Größen</option>
</select>
<div class="san_sizeSystem sizeSystems">
<div class="kindernormalgroessen hidden">
<p
class="san_sizeSystem--headline">Größe / Kindernormalgrößen</p>
<ul class="san_sizeFacet p_selTiles100 p_selTiles100--checkmark p_selTiles100--wide"
data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;kindernormalgroessen&quot;}">
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="152 <span>(1)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9raW5kZXJub3JtYWxncm9lc3Nlbj0xNTI=" data-merge-tracking='{"san_FilterValue":"152"}'> <span class="san_sizeFacet--value">152</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="158 <span>(1)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9raW5kZXJub3JtYWxncm9lc3Nlbj0xNTg=" data-merge-tracking='{"san_FilterValue":"158"}'> <span class="san_sizeFacet--value">158</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="164 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9raW5kZXJub3JtYWxncm9lc3Nlbj0xNjQ=" data-merge-tracking='{"san_FilterValue":"164"}'> <span class="san_sizeFacet--value">164</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="170 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9raW5kZXJub3JtYWxncm9lc3Nlbj0xNzA=" data-merge-tracking='{"san_FilterValue":"170"}'> <span class="san_sizeFacet--value">170</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="176 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9raW5kZXJub3JtYWxncm9lc3Nlbj0xNzY=" data-merge-tracking='{"san_FilterValue":"176"}'> <span class="san_sizeFacet--value">176</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="182 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9raW5kZXJub3JtYWxncm9lc3Nlbj0xODI=" data-merge-tracking='{"san_FilterValue":"182"}'> <span class="san_sizeFacet--value">182</span> </span> </li>
</ul>
</div>
<div class="normalgroessen hidden">
<p
class="san_sizeSystem--headline">Größe / Normalgrößen</p>
<ul class="san_sizeFacet p_selTiles100 p_selTiles100--checkmark p_selTiles100--wide"
data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;normalgroessen&quot;}">
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="32 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ub3JtYWxncm9lc3Nlbj0zMg==" data-merge-tracking='{"san_FilterValue":"32"}'> <span class="san_sizeFacet--value">32</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="34 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ub3JtYWxncm9lc3Nlbj0zNA==" data-merge-tracking='{"san_FilterValue":"34"}'> <span class="san_sizeFacet--value">34</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="36 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ub3JtYWxncm9lc3Nlbj0zNg==" data-merge-tracking='{"san_FilterValue":"36"}'> <span class="san_sizeFacet--value">36</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="38 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ub3JtYWxncm9lc3Nlbj0zOA==" data-merge-tracking='{"san_FilterValue":"38"}'> <span class="san_sizeFacet--value">38</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="40 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ub3JtYWxncm9lc3Nlbj00MA==" data-merge-tracking='{"san_FilterValue":"40"}'> <span class="san_sizeFacet--value">40</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="42 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ub3JtYWxncm9lc3Nlbj00Mg==" data-merge-tracking='{"san_FilterValue":"42"}'> <span class="san_sizeFacet--value">42</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="44 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ub3JtYWxncm9lc3Nlbj00NA==" data-merge-tracking='{"san_FilterValue":"44"}'> <span class="san_sizeFacet--value">44</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="46 <span>(2)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ub3JtYWxncm9lc3Nlbj00Ng==" data-merge-tracking='{"san_FilterValue":"46"}'> <span class="san_sizeFacet--value">46</span> </span> </li>
</ul>
</div>
<div class="us-groessen">
<p
class="san_sizeSystem--headline">Größe / US-Größen</p>
<ul class="san_sizeFacet p_selTiles100 p_selTiles100--checkmark p_selTiles100--wide"
data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;us-groessen&quot;}">
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="XS <span>(26)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz91cy1ncm9lc3Nlbj14cw==" data-merge-tracking='{"san_FilterValue":"xs"}'> <span class="san_sizeFacet--value">XS</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="S <span>(25)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz91cy1ncm9lc3Nlbj1z" data-merge-tracking='{"san_FilterValue":"s"}'> <span class="san_sizeFacet--value">S</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="M <span>(46)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz91cy1ncm9lc3Nlbj1t" data-merge-tracking='{"san_FilterValue":"m"}'> <span class="san_sizeFacet--value">M</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="L <span>(25)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz91cy1ncm9lc3Nlbj1s" data-merge-tracking='{"san_FilterValue":"l"}'> <span class="san_sizeFacet--value">L</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="XL <span>(26)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz91cy1ncm9lc3Nlbj14bA==" data-merge-tracking='{"san_FilterValue":"xl"}'> <span class="san_sizeFacet--value">XL</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="XXL <span>(25)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz91cy1ncm9lc3Nlbj14eGw=" data-merge-tracking='{"san_FilterValue":"xxl"}'> <span class="san_sizeFacet--value">XXL</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="XXXL <span>(9)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz91cy1ncm9lc3Nlbj14eHhs" data-merge-tracking='{"san_FilterValue":"xxxl"}'> <span class="san_sizeFacet--value">XXXL</span> </span> </li>
<li class="san_sizeFacet--box js_hasPaliTooltip"
data-tooltip="4XL <span>(7)</span>"
data-tooltip-pos="mouse" data-tooltip-touch="false">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz91cy1ncm9lc3Nlbj00eGw=" data-merge-tracking='{"san_FilterValue":"4xl"}'> <span class="san_sizeFacet--value">4XL</span> </span> </li>
</ul>
</div>
</div>
</dd>

        <dt class="p_accordion__header"
            id="facet-preis-in-eur"
            data-title="preis-in-eur"
            data-shortname="preis-in-eur">
                Preis
        </dt>
        <dd class="p_accordion__content facet retailprice"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;radio&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;preis-in-eur&quot;}">
    <ul class="san_facetValues san_facetValues--radio">
        <li class="san_facetValues__item customRange">
            <form class="san_facetValues__link san_customRange js_san_customRange"
                  action="/suche/waschmaschiene/?luf=preis-in-eur&amp;preis-in-eur=_from_-bis-_till_"
                  data-action-lowhigh="/suche/waschmaschiene/?luf=preis-in-eur&amp;preis-in-eur=_from_-bis-_till_"
                  data-action-low="/suche/waschmaschiene/?luf=preis-in-eur&amp;preis-in-eur=ab-_from_"
                  data-action-high="/suche/waschmaschiene/?luf=preis-in-eur&amp;preis-in-eur=bis-_till_"
                data-merge-tracking-lowhigh='{"san_FilterMethod":"free_text","san_FilterValue":"_from_-_till_"}'
                data-merge-tracking-low='{"san_FilterMethod":"free_text","san_FilterValue":"ab-_from_"}'
                data-merge-tracking-high='{"san_FilterMethod":"free_text","san_FilterValue":"bis-_till_"}'
                  novalidate>
                <input maxlength="5"
                       class="p_form__input san_customRange__input"
                       type="tel"
                       placeholder="3"
                       value=""/>
                <span class="san_customRange__dash">-</span>
                <input maxlength="5"
                       class="p_form__input san_customRange__input"
                       type="tel"
                       placeholder="1999"
                       value=""/>
                <button disabled class="p_btn150--1st san_customRange__button" type="submit"><i>&gt;</i></button>
            </form>
            <div class="p_message p_message--error san_customRange__error hidden"></div>
        </li>
    </ul>
                    <ul class="san_facetValues san_facetValues--checkbox" data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;one&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;reduziert&quot;}"}>
                            <li class="san_facetValues__item sale">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9yZWR1emllcnQ=" data-merge-tracking='{"san_FilterValue":"reduzierte-artikel"}'> <span class="san_facetValues__link">reduzierte Artikel</span> </span>                            </li>
                    </ul>
        </dd>
        <dt class="p_accordion__header"
            id="facet-verkaeufer"
            data-title="verkaeufer"
            data-shortname="verkaeufer">
                Verkäufer
        </dt>
        <dd class="p_accordion__content facet retailerFacet"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;radio&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;verkaeufer&quot;}">
                <ul class="san_facetValues san_facetValues--radio">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz92ZXJrYWV1ZmVyPW90dG8=" data-merge-tracking='{"san_FilterValue":"otto"}'> <span class="san_facetValues__link"> <span class="san-count" data-count="897">OTTO</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz92ZXJrYWV1ZmVyPWdlc2NoZW5raWRlZS1kZQ==" data-merge-tracking='{"san_FilterValue":"geschenkidee-de"}'> <span class="san_facetValues__link"> <span class="san-count" data-count="4">Geschenkidee.de</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz92ZXJrYWV1ZmVyPWludGVybmV0c3RvcmVz" data-merge-tracking='{"san_FilterValue":"internetstores"}'> <span class="san_facetValues__link"> <span class="san-count" data-count="3">Internetstores</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz92ZXJrYWV1ZmVyPXZjbQ==" data-merge-tracking='{"san_FilterValue":"vcm"}'> <span class="san_facetValues__link"> <span class="san-count" data-count="2">VCM</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz92ZXJrYWV1ZmVyPW15dG95cw==" data-merge-tracking='{"san_FilterValue":"mytoys"}'> <span class="san_facetValues__link"> <span class="san-count" data-count="2">myToys</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz92ZXJrYWV1ZmVyPW90dG8tb2ZmaWNl" data-merge-tracking='{"san_FilterValue":"otto-office"}'> <span class="san_facetValues__link"> <span class="san-count" data-count="1">OTTO Office</span> </span> </span>    </li>
                </ul>
        </dd>
        <dt class="p_accordion__header"
            id="facet-ladevolumen-waschen"
            data-title="ladevolumen-waschen"
            data-shortname="ladevolumen-waschen">
                Ladevolumen Waschen
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;ladevolumen-waschen&quot;}">
        
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9sYWRldm9sdW1lbi13YXNjaGVuPTgta2c=" data-merge-tracking='{"san_FilterValue":"8-kg"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="64">8 kg</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9sYWRldm9sdW1lbi13YXNjaGVuPTcta2c=" data-merge-tracking='{"san_FilterValue":"7-kg"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="48">7 kg</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9sYWRldm9sdW1lbi13YXNjaGVuPTYta2c=" data-merge-tracking='{"san_FilterValue":"6-kg"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="37">6 kg</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9sYWRldm9sdW1lbi13YXNjaGVuPTkta2c=" data-merge-tracking='{"san_FilterValue":"9-kg"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="27">9 kg</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9sYWRldm9sdW1lbi13YXNjaGVuPWFiLTEwLWtn" data-merge-tracking='{"san_FilterValue":"ab-10-kg"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="7">ab 10 kg</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9sYWRldm9sdW1lbi13YXNjaGVuPTUta2c=" data-merge-tracking='{"san_FilterValue":"5-kg"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="6">5 kg</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9sYWRldm9sdW1lbi13YXNjaGVuPXVudGVyLTUta2c=" data-merge-tracking='{"san_FilterValue":"unter-5-kg"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">unter 5 kg</span> </span> </span>    </li>
            </ul>
        
        </dd>
        <dt class="p_accordion__header"
            id="facet-breite"
            data-title="breite"
            data-shortname="breite">
                Breite
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;breite&quot;}">
        
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9icmVpdGU9NTYtNjAtY20=" data-merge-tracking='{"san_FilterValue":"56-60-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="184">56-60 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9icmVpdGU9YmlzLTU1LWNt" data-merge-tracking='{"san_FilterValue":"bis-55-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="26">Bis 55 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9icmVpdGU9NjUtNjktY20=" data-merge-tracking='{"san_FilterValue":"65-69-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">65 - 69 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9icmVpdGU9NjAtNjQtY20=" data-merge-tracking='{"san_FilterValue":"60-64-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">60 - 64 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9icmVpdGU9YWItNjEtY20=" data-merge-tracking='{"san_FilterValue":"ab-61-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Ab 61 cm</span> </span> </span>    </li>
            </ul>
        
        </dd>
        <dt class="p_accordion__header"
            id="facet-ausstattung"
            data-title="ausstattung"
            data-shortname="ausstattung">
                Ausstattung
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;ausstattung&quot;}">
        <div class="san-scroll-wrapper">
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz13YXNzZXJzY2h1dHo=" data-merge-tracking='{"san_FilterValue":"wasserschutz"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="184">Wasserschutz</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz1zY2huZWxscHJvZ3JhbW0=" data-merge-tracking='{"san_FilterValue":"schnellprogramm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="139">Schnellprogramm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz1pbnZlcnRlcm1vdG9y" data-merge-tracking='{"san_FilterValue":"invertermotor"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="131">Invertermotor</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz1nZW5hdWUtemVpdHBsYW51bmc=" data-merge-tracking='{"san_FilterValue":"genaue-zeitplanung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="106">genaue Zeitplanung</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz1uYWNobGVnZWZ1bmt0aW9u" data-merge-tracking='{"san_FilterValue":"nachlegefunktion"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="41">Nachlegefunktion</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz1hdXRvZG9zaWVydW5n" data-merge-tracking='{"san_FilterValue":"autodosierung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="29">Autodosierung</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz1zdGFydHplaXR2b3J3YWhs" data-merge-tracking='{"san_FilterValue":"startzeitvorwahl"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="28">Startzeitvorwahl</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz1hcHBmYWVoaWc=" data-merge-tracking='{"san_FilterValue":"appfaehig"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="27">appfähig</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz1kYW1wZmZ1bmt0aW9u" data-merge-tracking='{"san_FilterValue":"dampffunktion"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="23">Dampffunktion</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz11bnRlcmJhdWZhZWhpZw==" data-merge-tracking='{"san_FilterValue":"unterbaufaehig"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="8">Unterbaufähig</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz1kaXJla3RhbnRyaWVi" data-merge-tracking='{"san_FilterValue":"direktantrieb"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Direktantrieb</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9hdXNzdGF0dHVuZz1taXQtZGlzcGxheQ==" data-merge-tracking='{"san_FilterValue":"mit-display"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Mit Display</span> </span> </span>    </li>
            </ul>
        </div>
        </dd>
        <dt class="p_accordion__header"
            id="facet-touren"
            data-title="touren"
            data-shortname="touren">
                Touren
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;touren&quot;}">
        
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90b3VyZW49YmlzLTE0MDAtdS1taW4=" data-merge-tracking='{"san_FilterValue":"bis-1400-u-min"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="141">Bis 1400 U/Min.</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90b3VyZW49YmlzLTE2MDAtdS1taW4=" data-merge-tracking='{"san_FilterValue":"bis-1600-u-min"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="33">Bis 1600 U/Min.</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90b3VyZW49YmlzLTEyMDAtdS1taW4=" data-merge-tracking='{"san_FilterValue":"bis-1200-u-min"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="25">Bis 1200 U/Min.</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90b3VyZW49YmlzLTEwMDAtdS1taW4=" data-merge-tracking='{"san_FilterValue":"bis-1000-u-min"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="8">Bis 1000 U/Min.</span> </span> </span>    </li>
            </ul>
        
        </dd>
        <dt class="p_accordion__header"
            id="facet-tiefe"
            data-title="tiefe"
            data-shortname="tiefe">
                Tiefe
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;tiefe&quot;}">
        
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90aWVmZT01MS02MC1jbQ==" data-merge-tracking='{"san_FilterValue":"51-60-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="138">51 - 60 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90aWVmZT1hYi02MS1jbQ==" data-merge-tracking='{"san_FilterValue":"ab-61-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="55">ab 61 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90aWVmZT1iaXMtNTAtY20=" data-merge-tracking='{"san_FilterValue":"bis-50-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="20">bis 50 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90aWVmZT0yNS0yOS1jbQ==" data-merge-tracking='{"san_FilterValue":"25-29-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">25 - 29 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90aWVmZT0yMC0yNC1jbQ==" data-merge-tracking='{"san_FilterValue":"20-24-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">20 - 24 cm</span> </span> </span>    </li>
            </ul>
        
        </dd>
        <dt class="p_accordion__header"
            id="facet-hoehe"
            data-title="hoehe"
            data-shortname="hoehe">
                Höhe
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;hoehe&quot;}">
        
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ob2VoZT03MS04NS1jbQ==" data-merge-tracking='{"san_FilterValue":"71-85-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="183">71 - 85 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ob2VoZT1hYi04Ni1jbQ==" data-merge-tracking='{"san_FilterValue":"ab-86-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="27">ab 86 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ob2VoZT1iaXMtNzAtY20=" data-merge-tracking='{"san_FilterValue":"bis-70-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">bis 70 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ob2VoZT1hYi0xODAtY20=" data-merge-tracking='{"san_FilterValue":"ab-180-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">ab 180 cm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ob2VoZT0xNjAtMTc5LWNt" data-merge-tracking='{"san_FilterValue":"160-179-cm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">160 - 179 cm</span> </span> </span>    </li>
            </ul>
        
        </dd>
        <dt class="p_accordion__header"
            id="facet-energieeffizienz"
            data-title="energieeffizienz"
            data-shortname="energieeffizienz">
                Energieeffizienz
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;energieeffizienz&quot;}">
        
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9lbmVyZ2llZWZmaXppZW56PWFwbHVzcGx1c3BsdXM=" data-merge-tracking='{"san_FilterValue":"aplusplusplus"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="546">A+++</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9lbmVyZ2llZWZmaXppZW56PWFwbHVzcGx1cw==" data-merge-tracking='{"san_FilterValue":"aplusplus"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="28">A++</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9lbmVyZ2llZWZmaXppZW56PWFwbHVz" data-merge-tracking='{"san_FilterValue":"aplus"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="16">A+</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9lbmVyZ2llZWZmaXppZW56PWE=" data-merge-tracking='{"san_FilterValue":"a"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="6">A</span> </span> </span>    </li>
            </ul>
        
        </dd>
        <dt class="p_accordion__header"
            id="facet-bewertung"
            data-title="bewertung"
            data-shortname="bewertung">
                Bewertung
        </dt>
        <dd class="p_accordion__content facet starSelection js_loadUrlOnClick"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;stars&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;bewertung&quot;}">
            <div class="stars p_rating100">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9iZXdlcnR1bmc9YWItZWluZW0tc3Rlcm4=" data-merge-tracking='{"san_FilterValue":"ab-einem-stern"}'> <span class="star" data-starrating="1" data-message="1 Stern & mehr" data-count="318"> o </span> </span><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9iZXdlcnR1bmc9YWItMi1zdGVybmVu" data-merge-tracking='{"san_FilterValue":"ab-2-sternen"}'> <span class="star" data-starrating="2" data-message="2 Sterne & mehr" data-count="317"> o </span> </span><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9iZXdlcnR1bmc9YWItMy1zdGVybmVu" data-merge-tracking='{"san_FilterValue":"ab-3-sternen"}'> <span class="star" data-starrating="3" data-message="3 Sterne & mehr" data-count="317"> o </span> </span><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9iZXdlcnR1bmc9YWItNC1zdGVybmVu" data-merge-tracking='{"san_FilterValue":"ab-4-sternen"}'> <span class="star" data-starrating="4" data-message="4 Sterne & mehr" data-count="313"> o </span> </span><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9iZXdlcnR1bmc9NS1zdGVybmU=" data-merge-tracking='{"san_FilterValue":"5-sterne"}'> <span class="star" data-starrating="5" data-message="5 Sterne" data-count="253"> o </span> </span>            </div>
            <span id="countField">
            </span>
            <div class='message'>
                
            </div>
        </dd>
        <dt class="p_accordion__header"
            id="facet-sonderprogramme"
            data-title="sonderprogramme"
            data-shortname="sonderprogramme">
                Sonderprogramme
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;sonderprogramme&quot;}">
        <div class="san-scroll-wrapper">
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9d29sbGU=" data-merge-tracking='{"san_FilterValue":"wolle"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="159">Wolle</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9a3VyenByb2dyYW1t" data-merge-tracking='{"san_FilterValue":"kurzprogramm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="151">Kurzprogramm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9YWxsZXJnaWtlci1oeWdpZW5lLXByb2dyYW1t" data-merge-tracking='{"san_FilterValue":"allergiker-hygiene-programm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="84">Allergiker/Hygiene-Programm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9c3BvcnQ=" data-merge-tracking='{"san_FilterValue":"sport"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="80">Sport</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9aGVtZGVu" data-merge-tracking='{"san_FilterValue":"hemden"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="70">Hemden</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9ZW5lcmdpZXNwYXJwcm9ncmFtbQ==" data-merge-tracking='{"san_FilterValue":"energiesparprogramm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="69">Energiesparprogramm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9ZmxlY2tlbnByb2dyYW1t" data-merge-tracking='{"san_FilterValue":"fleckenprogramm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="53">Fleckenprogramm</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9ZGF1bmVu" data-merge-tracking='{"san_FilterValue":"daunen"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="34">Daunen</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9c2VsYnN0cmVpbmlndW5n" data-merge-tracking='{"san_FilterValue":"selbstreinigung"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="23">Selbstreinigung</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9a25pdHRlcnNjaHV0eg==" data-merge-tracking='{"san_FilterValue":"knitterschutz"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="17">Knitterschutz</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9zb25kZXJwcm9ncmFtbWU9bmFjaHRwcm9ncmFtbQ==" data-merge-tracking='{"san_FilterValue":"nachtprogramm"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="3">Nachtprogramm</span> </span> </span>    </li>
            </ul>
        </div>
        </dd>
        <dt class="p_accordion__header"
            id="facet-nachhaltige-produkte"
            data-title="nachhaltige-produkte"
            data-shortname="nachhaltige-produkte">
                Nachhaltige Produkte
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;nachhaltige-produkte&quot;}">
        
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9uYWNoaGFsdGlnZS1wcm9kdWt0ZT1jb3R0b24tbWFkZS1pbi1hZnJpY2E=" data-merge-tracking='{"san_FilterValue":"cotton-made-in-africa"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="2">Cotton made in Africa</span> </span> </span>    </li>
            </ul>
        
        </dd>
        <dt class="p_accordion__header"
            id="facet-testsiegel"
            data-title="testsiegel"
            data-shortname="testsiegel">
                Testsiegel
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;testsiegel&quot;}">
        
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90ZXN0c2llZ2VsPXN0aWZ0dW5nLXdhcmVudGVzdA==" data-merge-tracking='{"san_FilterValue":"stiftung-warentest"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="6">Stiftung Warentest</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz90ZXN0c2llZ2VsPXRlc3RtYWdhemlu" data-merge-tracking='{"san_FilterValue":"testmagazin"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="4">Testmagazin</span> </span> </span>    </li>
            </ul>
        
        </dd>
        <dt class="p_accordion__header"
            id="facet-rabatt"
            data-title="rabatt"
            data-shortname="rabatt">
                Reduzierung
        </dt>
        <dd class="p_accordion__content facet reductionClass"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;radio&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;rabatt&quot;}">
                <ul class="san_facetValues san_facetValues--radio">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9yYWJhdHQ9MzA=" data-merge-tracking='{"san_FilterValue":"mind-30"}'> <span class="san_facetValues__link"> <span class="san-count" data-count="153">mind. 30%</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9yYWJhdHQ9NDA=" data-merge-tracking='{"san_FilterValue":"mind-40"}'> <span class="san_facetValues__link"> <span class="san-count" data-count="49">mind. 40%</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9yYWJhdHQ9NTA=" data-merge-tracking='{"san_FilterValue":"mind-50"}'> <span class="san_facetValues__link"> <span class="san-count" data-count="12">mind. 50%</span> </span> </span>    </li>
                </ul>
        </dd>
        <dt class="p_accordion__header"
            id="facet-aktion"
            data-title="aktion"
            data-shortname="aktion">
                Aktion
        </dt>
        <dd class="p_accordion__content facet generic"
            
            data-basis-tracking="{&quot;san_FilterActivity&quot;:&quot;add&quot;,&quot;san_FilterMethod&quot;:&quot;checkbox&quot;,&quot;san_FilterPosition&quot;:&quot;local&quot;,&quot;san_FilterType&quot;:&quot;aktion&quot;}">
        
            <ul class="san_facetValues san_facetValues--checkbox">
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ha3Rpb249a2xpbWFwcmFlbWll" data-merge-tracking='{"san_FilterValue":"klimapraemie"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="187">Klimaprämie</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ha3Rpb249Y2FzaGJhY2s=" data-merge-tracking='{"san_FilterValue":"cashback"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="20">Cashback</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ha3Rpb249dW5zZXItdGlwcA==" data-merge-tracking='{"san_FilterValue":"unser-tipp"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="5">Unser Tipp!</span> </span> </span>    </li>
    <li class="san_facetValues__item">
<span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9ha3Rpb249YW5nZWJvdGUtZGVyLXdvY2hl" data-merge-tracking='{"san_FilterValue":"angebote-der-woche"}'> <span class="text san_facetValues__link ellipsis" > <span class="san-count" data-count="1">Angebote der Woche</span> </span> </span>    </li>
            </ul>
        
        </dd>
                        </dl>
                        
                    </div>
                <!-- ensure that filterWrapper exists -->
                <script type="text/javascript">o_san && o_san.Sidebar && o_san.Sidebar.init();</script>
            </dd>
        </dl>
    </div>


<div class="san_sorting__wrapper">
<form class="p_form">
<label for="sortingSelect" class="p_form__label">Sortieren nach</label>
<select id="sortingSelect" class="p_form__select">
<option data-basis-tracking='{&quot;san_SortingChange&quot;:&quot;topseller&quot;}'
class="p_form__option"
value="/suche/waschmaschiene/"
selected
>Topseller
</option>
<option data-basis-tracking='{&quot;san_SortingChange&quot;:&quot;preis-aufsteigend&quot;}'
class="p_form__option priceSortedKey"
value="/suche/waschmaschiene/?sortiertnach=preis-aufsteigend"
>Preis: niedrigster zuerst
</option>
<option data-basis-tracking='{&quot;san_SortingChange&quot;:&quot;preis-absteigend&quot;}'
class="p_form__option priceSortedKey"
value="/suche/waschmaschiene/?sortiertnach=preis-absteigend"
>Preis: höchster zuerst
</option>
<option data-basis-tracking='{&quot;san_SortingChange&quot;:&quot;neuheiten&quot;}'
class="p_form__option"
value="/suche/waschmaschiene/?sortiertnach=neuheiten"
>Neuheiten
</option>
<option data-basis-tracking='{&quot;san_SortingChange&quot;:&quot;bewertung&quot;}'
class="p_form__option"
value="/suche/waschmaschiene/?sortiertnach=bewertung"
>Bewertungen
</option>
</select>
</form>
</div>
<div id="avContent" class="js_av_searchResult" data-pt="Suchergebnisseite"></div>

    <link rel="stylesheet" href="/reco-core/assets/ft6.reco-core-aws.reco_assets.2cb41799.css" integrity="sha256-WxqNoq4E6bkAmmEmsOt632EgPyRNEawxNGNiyhw8nqk=" crossorigin="anonymous" />
<link rel="preload" crossorigin="anonymous" href="/reco-core/assets/ft6.reco-core-aws.reco_assets.779cb402.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)" />
<script>
  window.o_reco = window.o_reco || {};
  window.o_reco.api = window.o_reco.api || {};
  window.o_reco.private = window.o_reco.private || {};

  window.o_reco.private.delegate = function (originalFunction) {
    return function () {
      var passedArguments = arguments;
      return new Promise(function (resolve, reject) {
        o_global.eventLoader.onAllScriptsExecuted(10, function () {
          window.o_reco.private[originalFunction].apply(null, passedArguments).then(resolve).catch(reject);
        });
      });
    };
  };

  window.o_reco.private.delegateAndWrap = function (originalFunction) {
    return function () {
      var passedArguments = arguments;
      return new Promise(function (resolve) {
        o_global.eventLoader.onAllScriptsExecuted(10, function () {
          var result = window.o_reco.private[originalFunction].apply(null, passedArguments);
          resolve(result);
        });
      });
    };
  };

  window.o_reco.api.initializeCinema = window.o_reco.private.delegateAndWrap('initializeCinema');
  window.o_reco.api.initializeExpandableCinema = window.o_reco.private.delegateAndWrap('initializeExpandableCinema');

  window.o_reco.api.loadRecommendationsForDetailView = window.o_reco.private.delegate('loadRecommendationsForDetailView');
  window.o_reco.api.loadRecommendationsForProductLine = window.o_reco.private.delegate('loadRecommendationsForProductLine');
  window.o_reco.api.loadRecommendationsForOrderOverview = window.o_reco.private.delegate('loadRecommendationsForOrderOverview');
  window.o_reco.api.loadRecommendationsForWishlist = window.o_reco.private.delegate('loadRecommendationsForWishlist');
  window.o_reco.api.loadRecommendationsForEntryPage = window.o_reco.private.delegate('loadRecommendationsForEntryPage');
  window.o_reco.api.loadBounceLayerCinema = window.o_reco.private.delegateAndWrap('loadBounceLayerCinema');
</script>

<div class="exactag san" data-trigger="sanProductList"></div>
<div id="san_searchResult"
     class="searchResult san_productLinkQP dominant-assortment-technik"
     data-dominant-assortment="TECHNIK"
     data-userquery="waschmaschiene"
     data-itemcount="432">
    <span id="san_advice"></span>
    <section id="san_resultSection">
        <article class="product large clearfix"
                 data-productid="308075513"
                 data-variationid="667359840-M24"
                 data-articlenumber="16434749"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;1&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/308075513"
                 data-listposition="1"
                 data-facetted-variation-ids="667359840-M24,667359840-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/lg-waschmaschine-f-1496-qd3ht-7-kg-1400-u-min-unterbaufaehig-308075513/#variationId=667359840-M24">
            LG Waschmaschine F 1496 QD3HT, 7 kg, 1400 U/Min, unterbaufähig
            </a>
                <script id="j_308075513" type="text/html">
                    {"product":{"brand":{"imageUrl":"https://i.otto.de/i/otto/3ef8444476f42313ad64fba0781203dd?$ov_brandlogo$","value":"LG","code":"LG"},"variationMap":{"667359840-M24":{"id":"667359840-M24","name":"LG Waschmaschine F 1496 QD3HT, 7 kg, 1400 U/Min, unterbaufähig","retailPrice":39900,"link":{"href":"/p/lg-waschmaschine-f-1496-qd3ht-7-kg-1400-u-min-unterbaufaehig-308075513/#variationId=667359840-M24","encodedHref":"L3AvbGctd2FzY2htYXNjaGluZS1mLTE0OTYtcWQzaHQtNy1rZy0xNDAwLXUtbWluLXVudGVyYmF1ZmFlaGlnLTMwODA3NTUxMy8jdmFyaWF0aW9uSWQ9NjY3MzU5ODQwLU0yNA=="},"imageUrl":"https://i.otto.de/i/otto/27162361?$responsive_ft2$","availability":{"value":"AVAILABLE"},"sale":true,"articleNumber":"16434749","hasSuggestedRetailPrice":true,"hasMoreExpensiveAvailableVariation":true,"formattedRetailPrice":"399,00","formattedComparativePrice":"749,00","isDealDateToday":false,"standardDescription":"<ul><li>mindestens 30% sparsamer als der Grenzwert zu A+++</li><li>7 kg Fassungsvermögen</li><li>1400 Touren</li><li>Betriebsgeräusch Waschen 52 dB</li><li>Inverter Direct Drive™</li></ul>","hasSellingPoints":true,"is24h":false,"expertReviews":{"reviewsCount":0,"hasExpertReview":false},"energyEfficiencyCategoryClass":{"category":"A","categorySuffix":"+++","color":"darkGreen","images":["https://i.otto.de/i/otto/25279522?$dv_energieetikett_s$"]}},"667359840-M48":{"id":"667359840-M48","name":"LG Waschmaschine F 1496 QD3HT, 7 kg, 1400 U/Min, unterbaufähig","retailPrice":44900,"link":{"href":"/p/lg-waschmaschine-f-1496-qd3ht-7-kg-1400-u-min-unterbaufaehig-308075513/#variationId=667359840-M48","encodedHref":"L3AvbGctd2FzY2htYXNjaGluZS1mLTE0OTYtcWQzaHQtNy1rZy0xNDAwLXUtbWluLXVudGVyYmF1ZmFlaGlnLTMwODA3NTUxMy8jdmFyaWF0aW9uSWQ9NjY3MzU5ODQwLU00OA=="},"imageUrl":"https://i.otto.de/i/otto/27162361?$responsive_ft2$","availability":{"value":"AVAILABLE"},"sale":true,"articleNumber":"16434749","hasSuggestedRetailPrice":true,"hasMoreExpensiveAvailableVariation":false,"formattedRetailPrice":"449,00","formattedComparativePrice":"749,00","isDealDateToday":false,"standardDescription":"<ul><li>mindestens 30% sparsamer als der Grenzwert zu A+++</li><li>7 kg Fassungsvermögen</li><li>1400 Touren</li><li>Betriebsgeräusch Waschen 52 dB</li><li>Inverter Direct Drive™</li></ul>","hasSellingPoints":true,"is24h":false,"expertReviews":{"reviewsCount":0,"hasExpertReview":false},"energyEfficiencyCategoryClass":{"category":"A","categorySuffix":"+++","color":"darkGreen","images":["https://i.otto.de/i/otto/25279522?$dv_energieetikett_s$"]}}},"differentVariationPrices":false,"customerReviews":{"numberOfReviews":2145,"averageRatingAsIconFont":"* * * * *"},"hasSeries":false,"sizesOverlayDisplayStatus":"NOT_DISPLAYABLE_NO_FASHION","largeSizePlus":false},"colorDetails":[{"name":"weiß","colorHexCode":"ededf0","baseColor":"weiss","variationIdsSortedByCheapestAvailability":["667359840-M24","667359840-M48"],"availableVariationIds":[]}],"variationIdsSortedByCheapestAvailability":["667359840-M24","667359840-M48"],"availableVariationIds":[],"placeholderUrl":"https://i.otto.de/i/otto/lh_platzhalter_ohne_abbildung"}
                </script>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="459358199"
                 data-variationid="567436358-M24"
                 data-articlenumber="800824"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;2&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/459358199"
                 data-listposition="2"
                 data-facetted-variation-ids="567436358-M24,567436358-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/siemens-waschmaschine-iq700-wm16w540-8-kg-1600-u-min-459358199/#variationId=567436358-M24">
            SIEMENS Waschmaschine iQ700 WM16W540, 8 kg, 1600 U/Min
            </a>
                <script id="j_459358199" type="text/html">
                    {"product":{"brand":{"imageUrl":"https://i.otto.de/i/otto/e17346633b2bfc0449165c48aec51ec1?$ov_brandlogo$","value":"SIEMENS","code":"SIEMENS"},"variationMap":{"567436358-M24":{"id":"567436358-M24","name":"SIEMENS Waschmaschine iQ700 WM16W540, 8 kg, 1600 U/Min","retailPrice":61900,"link":{"href":"/p/siemens-waschmaschine-iq700-wm16w540-8-kg-1600-u-min-459358199/#variationId=567436358-M24","encodedHref":"L3Avc2llbWVucy13YXNjaG1hc2NoaW5lLWlxNzAwLXdtMTZ3NTQwLTgta2ctMTYwMC11LW1pbi00NTkzNTgxOTkvI3ZhcmlhdGlvbklkPTU2NzQzNjM1OC1NMjQ="},"imageUrl":"https://i.otto.de/i/otto/14406566?$responsive_ft2$","availability":{"value":"AVAILABLE"},"sale":true,"articleNumber":"800824","hasSuggestedRetailPrice":true,"hasMoreExpensiveAvailableVariation":true,"formattedRetailPrice":"619,00","formattedComparativePrice":"1.079,00","isDealDateToday":false,"standardDescription":"<ul><li>mindestens 30% sparsamer als der Grenzwert zu A+++</li><li>8 kg Fassungsvermögen</li><li>1600 Touren</li><li>Betriebsgeräusch Waschen 48 dB</li><li>Intelligenter, langlebiger und leiser iQdrive-Motor für besonders wirkungsvolle und effiziente Wäschepflege</li></ul>","hasSellingPoints":true,"is24h":false,"expertReviews":{"hasExpertReview":false,"reviewsCount":0},"deal":{"iconUrl":"https://i.otto.de/i/otto/001_2018_23_haus_klimapraemie_flag_flag?$ov_promoflag$"},"energyEfficiencyCategoryClass":{"category":"A","categorySuffix":"+++","color":"darkGreen","images":["https://i.otto.de/i/otto/18131329?$dv_energieetikett_s$"]}},"567436358-M48":{"id":"567436358-M48","name":"SIEMENS Waschmaschine iQ700 WM16W540, 8 kg, 1600 U/Min","retailPrice":66900,"link":{"href":"/p/siemens-waschmaschine-iq700-wm16w540-8-kg-1600-u-min-459358199/#variationId=567436358-M48","encodedHref":"L3Avc2llbWVucy13YXNjaG1hc2NoaW5lLWlxNzAwLXdtMTZ3NTQwLTgta2ctMTYwMC11LW1pbi00NTkzNTgxOTkvI3ZhcmlhdGlvbklkPTU2NzQzNjM1OC1NNDg="},"imageUrl":"https://i.otto.de/i/otto/14406566?$responsive_ft2$","availability":{"value":"AVAILABLE"},"sale":true,"articleNumber":"800824","hasSuggestedRetailPrice":true,"hasMoreExpensiveAvailableVariation":false,"formattedRetailPrice":"669,00","formattedComparativePrice":"1.079,00","isDealDateToday":false,"standardDescription":"<ul><li>mindestens 30% sparsamer als der Grenzwert zu A+++</li><li>8 kg Fassungsvermögen</li><li>1600 Touren</li><li>Betriebsgeräusch Waschen 48 dB</li><li>Intelligenter, langlebiger und leiser iQdrive-Motor für besonders wirkungsvolle und effiziente Wäschepflege</li></ul>","hasSellingPoints":true,"is24h":false,"expertReviews":{"hasExpertReview":false,"reviewsCount":0},"deal":{"iconUrl":"https://i.otto.de/i/otto/001_2018_23_haus_klimapraemie_flag_flag?$ov_promoflag$"},"energyEfficiencyCategoryClass":{"category":"A","categorySuffix":"+++","color":"darkGreen","images":["https://i.otto.de/i/otto/18131329?$dv_energieetikett_s$"]}}},"differentVariationPrices":false,"customerReviews":{"numberOfReviews":561,"averageRatingAsIconFont":"* * * * *"},"hasSeries":false,"sizesOverlayDisplayStatus":"NOT_DISPLAYABLE_NO_FASHION","largeSizePlus":false},"colorDetails":[{"name":"weiß","colorHexCode":"dcdedd","baseColor":"weiss","variationIdsSortedByCheapestAvailability":["567436358-M24","567436358-M48"],"availableVariationIds":[]}],"variationIdsSortedByCheapestAvailability":["567436358-M24","567436358-M48"],"availableVariationIds":[],"placeholderUrl":"https://i.otto.de/i/otto/lh_platzhalter_ohne_abbildung"}
                </script>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="538742847"
                 data-variationid="538742878-M24"
                 data-articlenumber="836814"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;3&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/538742847"
                 data-listposition="3"
                 data-facetted-variation-ids="538742878-M24,538742878-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/samsung-waschmaschine-ww6500-addwash-ww80k6404qx-eg-8-kg-1400-u-min-538742847/#variationId=538742878-M24">
            Samsung Waschmaschine WW6500 AddWash WW80K6404QX/EG, 8 kg, 1400 U/Min
            </a>
                <script id="j_538742847" type="text/html">
                    {"product":{"brand":{"imageUrl":"https://i.otto.de/i/otto/1c317aaf7751f8b652912956fea2ae60?$ov_brandlogo$","value":"Samsung","code":"SAMSUNG"},"variationMap":{"538742878-M48":{"id":"538742878-M48","name":"Samsung Waschmaschine WW6500 AddWash WW80K6404QX/EG, 8 kg, 1400 U/Min","retailPrice":73900,"link":{"href":"/p/samsung-waschmaschine-ww6500-addwash-ww80k6404qx-eg-8-kg-1400-u-min-538742847/#variationId=538742878-M48","encodedHref":"L3Avc2Ftc3VuZy13YXNjaG1hc2NoaW5lLXd3NjUwMC1hZGR3YXNoLXd3ODBrNjQwNHF4LWVnLTgta2ctMTQwMC11LW1pbi01Mzg3NDI4NDcvI3ZhcmlhdGlvbklkPTUzODc0Mjg3OC1NNDg="},"imageUrl":"https://i.otto.de/i/otto/16517211?$responsive_ft2$","availability":{"value":"AVAILABLE"},"sale":true,"articleNumber":"836814","hasSuggestedRetailPrice":true,"hasMoreExpensiveAvailableVariation":false,"formattedRetailPrice":"739,00","formattedComparativePrice":"999,00","isDealDateToday":false,"standardDescription":"<ul><li>mindestens 40% sparsamer als der Grenzwert zu A+++</li><li>8 kg Fassungsvermögen</li><li>1400 Touren</li><li>Betriebsgeräusch Waschen 53 dB</li><li>AddWash – Nachlegen leicht gemacht zu jeder Zeit</li></ul>","hasSellingPoints":true,"is24h":false,"expertReviews":{"hasExpertReview":false,"reviewsCount":0},"deal":{"iconUrl":"https://i.otto.de/i/otto/001_2018_23_haus_klimapraemie_flag_flag?$ov_promoflag$"},"energyEfficiencyCategoryClass":{"category":"A","categorySuffix":"+++","color":"darkGreen","images":["https://i.otto.de/i/otto/16339991?$dv_energieetikett_s$"]}},"538742878-M24":{"id":"538742878-M24","name":"Samsung Waschmaschine WW6500 AddWash WW80K6404QX/EG, 8 kg, 1400 U/Min","retailPrice":68900,"link":{"href":"/p/samsung-waschmaschine-ww6500-addwash-ww80k6404qx-eg-8-kg-1400-u-min-538742847/#variationId=538742878-M24","encodedHref":"L3Avc2Ftc3VuZy13YXNjaG1hc2NoaW5lLXd3NjUwMC1hZGR3YXNoLXd3ODBrNjQwNHF4LWVnLTgta2ctMTQwMC11LW1pbi01Mzg3NDI4NDcvI3ZhcmlhdGlvbklkPTUzODc0Mjg3OC1NMjQ="},"imageUrl":"https://i.otto.de/i/otto/16517211?$responsive_ft2$","availability":{"value":"AVAILABLE"},"sale":true,"articleNumber":"836814","hasSuggestedRetailPrice":true,"hasMoreExpensiveAvailableVariation":true,"formattedRetailPrice":"689,00","formattedComparativePrice":"999,00","isDealDateToday":false,"standardDescription":"<ul><li>mindestens 40% sparsamer als der Grenzwert zu A+++</li><li>8 kg Fassungsvermögen</li><li>1400 Touren</li><li>Betriebsgeräusch Waschen 53 dB</li><li>AddWash – Nachlegen leicht gemacht zu jeder Zeit</li></ul>","hasSellingPoints":true,"is24h":false,"expertReviews":{"hasExpertReview":false,"reviewsCount":0},"deal":{"iconUrl":"https://i.otto.de/i/otto/001_2018_23_haus_klimapraemie_flag_flag?$ov_promoflag$"},"energyEfficiencyCategoryClass":{"category":"A","categorySuffix":"+++","color":"darkGreen","images":["https://i.otto.de/i/otto/16339991?$dv_energieetikett_s$"]}}},"differentVariationPrices":false,"customerReviews":{"numberOfReviews":139,"averageRatingAsIconFont":"* * * * *"},"hasSeries":false,"sizesOverlayDisplayStatus":"NOT_DISPLAYABLE_NO_FASHION","largeSizePlus":false},"colorDetails":[{"name":"grau","colorHexCode":"4d484e","baseColor":"grau","variationIdsSortedByCheapestAvailability":["538742878-M24","538742878-M48"],"availableVariationIds":[]}],"variationIdsSortedByCheapestAvailability":["538742878-M24","538742878-M48"],"availableVariationIds":[],"placeholderUrl":"https://i.otto.de/i/otto/lh_platzhalter_ohne_abbildung"}
                </script>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="527967277"
                 data-variationid="528038095-M24"
                 data-articlenumber="434054"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;4&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/527967277"
                 data-listposition="4"
                 data-facetted-variation-ids="528038095-M24,528038095-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/siemens-waschmaschine-iq300-wm14n0eco-6-kg-1400-u-min-527967277/#variationId=528038095-M24">
            SIEMENS Waschmaschine iQ300 WM14N0ECO, 6 kg, 1400 U/Min
            </a>
                <script id="j_527967277" type="text/html">
                    {"product":{"brand":{"imageUrl":"https://i.otto.de/i/otto/e17346633b2bfc0449165c48aec51ec1?$ov_brandlogo$","value":"SIEMENS","code":"SIEMENS"},"variationMap":{"528038095-M48":{"id":"528038095-M48","name":"SIEMENS Waschmaschine iQ300 WM14N0ECO, 6 kg, 1400 U/Min","retailPrice":44900,"link":{"href":"/p/siemens-waschmaschine-iq300-wm14n0eco-6-kg-1400-u-min-527967277/#variationId=528038095-M48","encodedHref":"L3Avc2llbWVucy13YXNjaG1hc2NoaW5lLWlxMzAwLXdtMTRuMGVjby02LWtnLTE0MDAtdS1taW4tNTI3OTY3Mjc3LyN2YXJpYXRpb25JZD01MjgwMzgwOTUtTTQ4"},"imageUrl":"https://i.otto.de/i/otto/15931046?$responsive_ft2$","availability":{"value":"AVAILABLE"},"sale":true,"articleNumber":"434054","hasSuggestedRetailPrice":true,"hasMoreExpensiveAvailableVariation":false,"formattedRetailPrice":"449,00","formattedComparativePrice":"649,00","isDealDateToday":false,"standardDescription":"<ul><li>mindestens 10% sparsamer als der Grenzwert zu A+++</li><li>6 kg Fassungsvermögen</li><li>1400 Touren</li><li>Betriebsgeräusch Waschen 54 dB</li><li>Intelligenter, langlebiger und leiser iQdrive-Motor</li></ul>","hasSellingPoints":true,"is24h":false,"expertReviews":{"reviewsCount":0,"hasExpertReview":false},"energyEfficiencyCategoryClass":{"category":"A","categorySuffix":"+++","color":"darkGreen","images":["https://i.otto.de/i/otto/16062657?$dv_energieetikett_s$"]}},"528038095-M24":{"id":"528038095-M24","name":"SIEMENS Waschmaschine iQ300 WM14N0ECO, 6 kg, 1400 U/Min","retailPrice":39900,"link":{"href":"/p/siemens-waschmaschine-iq300-wm14n0eco-6-kg-1400-u-min-527967277/#variationId=528038095-M24","encodedHref":"L3Avc2llbWVucy13YXNjaG1hc2NoaW5lLWlxMzAwLXdtMTRuMGVjby02LWtnLTE0MDAtdS1taW4tNTI3OTY3Mjc3LyN2YXJpYXRpb25JZD01MjgwMzgwOTUtTTI0"},"imageUrl":"https://i.otto.de/i/otto/15931046?$responsive_ft2$","availability":{"value":"AVAILABLE"},"sale":true,"articleNumber":"434054","hasSuggestedRetailPrice":true,"hasMoreExpensiveAvailableVariation":true,"formattedRetailPrice":"399,00","formattedComparativePrice":"649,00","isDealDateToday":false,"standardDescription":"<ul><li>mindestens 10% sparsamer als der Grenzwert zu A+++</li><li>6 kg Fassungsvermögen</li><li>1400 Touren</li><li>Betriebsgeräusch Waschen 54 dB</li><li>Intelligenter, langlebiger und leiser iQdrive-Motor</li></ul>","hasSellingPoints":true,"is24h":false,"expertReviews":{"reviewsCount":0,"hasExpertReview":false},"energyEfficiencyCategoryClass":{"category":"A","categorySuffix":"+++","color":"darkGreen","images":["https://i.otto.de/i/otto/16062657?$dv_energieetikett_s$"]}}},"differentVariationPrices":false,"customerReviews":{"numberOfReviews":353,"averageRatingAsIconFont":"* * * * *"},"hasSeries":false,"sizesOverlayDisplayStatus":"NOT_DISPLAYABLE_NO_FASHION","largeSizePlus":false},"colorDetails":[{"name":"weiß","colorHexCode":"e5e6e9","baseColor":"weiss","variationIdsSortedByCheapestAvailability":["528038095-M24","528038095-M48"],"availableVariationIds":[]}],"variationIdsSortedByCheapestAvailability":["528038095-M24","528038095-M48"],"availableVariationIds":[],"placeholderUrl":"https://i.otto.de/i/otto/lh_platzhalter_ohne_abbildung"}
                </script>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="685862423"
                 data-variationid="685863023-M48"
                 data-articlenumber="26146742"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;5&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/685862423"
                 data-listposition="5"
                 data-facetted-variation-ids="685863023-M48,685863440-M48,685863464-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/bauknecht-waschmaschine-super-eco-7418-9-kg-1400-u-min-inkl-4-jahre-herstellergarantie-685862423/#variationId=685863023-M48">
            BAUKNECHT Waschmaschine Super Eco 7418, 9 kg, 1400 U/Min, inkl. 4 Jahre Herstellergarantie
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="622071982"
                 data-variationid="622074434-M24"
                 data-articlenumber="48643664"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;6&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/622071982"
                 data-listposition="6"
                 data-facetted-variation-ids="622074434-M24,622074434-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/miele-waschmaschine-wkf-311-wps-speedcare-aplusplusplus-8-kg-1400-u-min-622071982/#variationId=622074434-M24">
            MIELE Waschmaschine WKF 311 WPS SpeedCare, A+++, 8 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="581412279"
                 data-variationid="581412515-M24"
                 data-articlenumber="573022"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;7&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/581412279"
                 data-listposition="7"
                 data-facetted-variation-ids="581412515-M24,581412515-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/lg-waschmaschine-f-14wm-9en0-9-kg-1400-u-min-581412279/#variationId=581412515-M24">
            LG Waschmaschine F 14WM 9EN0, 9 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="537219709"
                 data-variationid="537222960-M24"
                 data-articlenumber="156261"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;8&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/537219709"
                 data-listposition="8"
                 data-facetted-variation-ids="537222960-M24,537222960-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/bosch-waschmaschine-serie-4-doreen-wan282v8-7-kg-1400-u-min-537219709/#variationId=537222960-M24">
            BOSCH Waschmaschine Serie 4 Doreen WAN282V8, 7 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="424309145"
                 data-variationid="552339124-M24"
                 data-articlenumber="763513"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;9&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/424309145"
                 data-listposition="9"
                 data-facetted-variation-ids="552339124-M24,552339124-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/siemens-waschmaschine-iq100-wm14b2eco-6-kg-1400-u-min-424309145/#variationId=552339124-M24">
            SIEMENS Waschmaschine iQ100 WM14B2ECO, 6 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="571465955"
                 data-variationid="552338277-M24"
                 data-articlenumber="793487"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;10&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/571465955"
                 data-listposition="10"
                 data-facetted-variation-ids="552338277-M24,552338277-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/bosch-waschmaschine-serie-8-waw32541-8-kg-1600-u-min-571465955/#variationId=552338277-M24">
            BOSCH Waschmaschine Serie 8 WAW32541, 8 kg, 1600 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="701225697"
                 data-variationid="701225866-M24"
                 data-articlenumber="13289223"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;11&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/701225697"
                 data-listposition="11"
                 data-facetted-variation-ids="701225866-M24,701225866-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/siemens-waschmaschine-iq500-wm14t421-7-kg-1400-u-min-701225697/#variationId=701225866-M24">
            SIEMENS Waschmaschine iQ500 WM14T421, 7 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="624440607"
                 data-variationid="492371700-M48"
                 data-articlenumber="442093"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;12&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/624440607"
                 data-listposition="12"
                 data-facetted-variation-ids="492371700-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/siemens-waschmaschine-iq-500-wm14t641-8-kg-1400-u-min-i-dos-dosierautomatik-624440607/#variationId=492371700-M48">
            SIEMENS Waschmaschine iQ 500 WM14T641, 8 kg, 1400 U/Min, i-Dos Dosierautomatik
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="594023330"
                 data-variationid="594023494-M24"
                 data-articlenumber="300737"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;13&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/594023330"
                 data-listposition="13"
                 data-facetted-variation-ids="594023494-M24,594023494-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/beko-waschmaschine-wmo-626-6-kg-1600-u-min-594023330/#variationId=594023494-M24">
            BEKO Waschmaschine WMO 626, 6 kg, 1600 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="534028361"
                 data-variationid="534028622-M24"
                 data-articlenumber="240294"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;14&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/534028361"
                 data-listposition="14"
                 data-facetted-variation-ids="534028622-M24,534028622-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/bosch-waschmaschine-serie-8-waw325v0-9-kg-1600-u-min-534028361/#variationId=534028622-M24">
            BOSCH Waschmaschine Serie 8 WAW325V0, 9 kg, 1600 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="571296502"
                 data-variationid="570447057-M24"
                 data-articlenumber="622261"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;15&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/571296502"
                 data-listposition="15"
                 data-facetted-variation-ids="570447057-M24,570447057-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/miele-waschmaschine-wdb030wcs-d-lw-eco-7-kg-1400-u-min-571296502/#variationId=570447057-M24">
            Miele Waschmaschine WDB030WCS D LW Eco, 7 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="570591140"
                 data-variationid="643062149-M24"
                 data-articlenumber="97092609"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;16&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/570591140"
                 data-listposition="16"
                 data-facetted-variation-ids="643062149-M24,643062149-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/miele-waschmaschine-wkh132wps-d-lw-pwash-2-0-tdos-xl-9-kg-1600-u-min-570591140/#variationId=643062149-M24">
            Miele Waschmaschine WKH132WPS D LW PWash 2.0 &amp; TDos XL, 9 kg, 1600 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="570368839"
                 data-variationid="552339487-M24"
                 data-articlenumber="425832"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;17&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/570368839"
                 data-listposition="17"
                 data-facetted-variation-ids="552339487-M24,552339487-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/bauknecht-waschmaschine-toplader-wmt-ecostar-732-di-7-kg-1200-u-min-570368839/#variationId=552339487-M24">
            BAUKNECHT Waschmaschine Toplader WMT EcoStar 732 Di, 7 kg, 1200 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="566081524"
                 data-variationid="563540852-M24"
                 data-articlenumber="691776"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;18&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/566081524"
                 data-listposition="18"
                 data-facetted-variation-ids="563540852-M24,563540852-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/privileg-waschmaschine-pwf-m-643-6-kg-1400-u-min-566081524/#variationId=563540852-M24">
            Privileg Waschmaschine PWF M 643, 6 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="644018080"
                 data-variationid="690089463-M24"
                 data-articlenumber="24077134"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;19&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/644018080"
                 data-listposition="19"
                 data-facetted-variation-ids="690089463-M24,690089463-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/samsung-waschmaschine-ww4500-ww7ek44205w-eg-7-kg-1400-u-min-addwash-644018080/#variationId=690089463-M24">
            Samsung Waschmaschine WW4500 WW7EK44205W/EG, 7 kg, 1400 U/Min, AddWash
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="623340487"
                 data-variationid="623342642-M24"
                 data-articlenumber="91673735"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;20&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/623340487"
                 data-listposition="20"
                 data-facetted-variation-ids="623342642-M24,623342642-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/privileg-waschmaschine-toplader-pwt-a51052-5-kg-1000-u-min-623340487/#variationId=623342642-M24">
            Privileg Waschmaschine Toplader PWT A51052, 5 kg, 1000 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="595277928"
                 data-variationid="595278964-M48"
                 data-articlenumber="348514"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;21&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/595277928"
                 data-listposition="21"
                 data-facetted-variation-ids="595278964-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/aeg-waschmaschine-l6fb48fl-8-kg-1400-u-min-595277928/#variationId=595278964-M48">
            AEG Waschmaschine L6FB48FL, 8 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="639596120"
                 data-variationid="639596625-M24"
                 data-articlenumber="85178333"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;22&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/639596120"
                 data-listposition="22"
                 data-facetted-variation-ids="639596625-M24,639596625-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/bosch-waschmaschine-homeprofessional-way287w5-8-kg-1400-u-min-639596120/#variationId=639596625-M24">
            BOSCH Waschmaschine HomeProfessional WAY287W5, 8 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="479421262"
                 data-variationid="567440422-M24"
                 data-articlenumber="519823"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;23&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/479421262"
                 data-listposition="23"
                 data-facetted-variation-ids="567440422-M24,567440422-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/bosch-waschmaschine-serie-4-wae282v7-7-kg-1400-u-min-479421262/#variationId=567440422-M24">
            BOSCH Waschmaschine Serie 4 WAE282V7, 7 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="537226649"
                 data-variationid="537231915-M24"
                 data-articlenumber="143446"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;24&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/537226649"
                 data-listposition="24"
                 data-facetted-variation-ids="537231915-M24,537231915-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/bosch-waschmaschine-serie-8-waw285v0-9-kg-1400-u-min-537226649/#variationId=537231915-M24">
            BOSCH Waschmaschine Serie 8 WAW285V0, 9 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="581410552"
                 data-variationid="581410763-M24"
                 data-articlenumber="887875"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;25&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/581410552"
                 data-listposition="25"
                 data-facetted-variation-ids="581410763-M24,581410763-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/lg-waschmaschine-f-14wm-7en0-7-kg-1400-u-min-581410552/#variationId=581410763-M24">
            LG Waschmaschine F 14WM 7EN0, 7 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="699673726"
                 data-variationid="684332407-M24"
                 data-articlenumber="23851617"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;26&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/699673726"
                 data-listposition="26"
                 data-facetted-variation-ids="684332407-M24,684332407-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/gorenje-waschmaschine-wave-p-62s3-p-6-kg-1200-u-min-699673726/#variationId=684332407-M24">
            GORENJE Waschmaschine Wave P 62S3 P, 6 kg, 1200 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="448878977"
                 data-variationid="448880304-M24"
                 data-articlenumber="369652"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;27&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/448878977"
                 data-listposition="27"
                 data-facetted-variation-ids="448880304-M24,448880304-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/siemens-waschmaschine-iq700-wm14w5eco-8-kg-1400-u-min-448878977/#variationId=448880304-M24">
            SIEMENS Waschmaschine iQ700 WM14W5ECO, 8 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="699673725"
                 data-variationid="684334045-M24"
                 data-articlenumber="35152965"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;28&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/699673725"
                 data-listposition="28"
                 data-facetted-variation-ids="684334045-M24,684334045-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/gorenje-waschmaschine-wave-e-74s3-p-7-kg-1400-u-min-699673725/#variationId=684334045-M24">
            GORENJE Waschmaschine Wave E 74S3 P, 7 kg, 1400 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="460469386"
                 data-variationid="581198609-M24"
                 data-articlenumber="220978"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;29&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/460469386"
                 data-listposition="29"
                 data-facetted-variation-ids="581198609-M24,581198609-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/hanseatic-waschmaschine-hwm510a1-5-kg-1000-u-min-460469386/#variationId=581198609-M24">
            Hanseatic Waschmaschine HWM510A1, 5 kg, 1000 U/Min
            </a>
        </article>
        <hr class="p_line100"/>
        <article class="product large clearfix"
                 data-productid="466495389"
                 data-variationid="466498581-M24"
                 data-articlenumber="359008"
                 data-producttile-tracking="{&quot;wk.san_ListPosition&quot;:&quot;30&quot;,&quot;san_ProductLinkType&quot;:&quot;#linkType#&quot;}"
                 data-json-target="/tiles/products/views/tile/466495389"
                 data-listposition="30"
                 data-facetted-variation-ids="466498581-M24,466498581-M48">
            <a class="name lazy productLink facettedAjax"
               href="/p/bosch-waschmaschine-serie-8-waw28500-9-kg-1400-u-min-466495389/#variationId=466498581-M24">
            BOSCH Waschmaschine Serie 8 WAW28500, 9 kg, 1400 U/Min
            </a>
        </article>
        <div class="san_productListTeaser san_productListTeaser__target js_san_productListTeaser__target">
            <hr class="p_line100"/>
        </div>
    </section>
    <div class="clear"></div>
    <div class="san_paging__bottomWrapper"><ul class="san_paging_block">
<li id="san_pagingCurrentPage" class="san_paging__item san_paging__item--current"> <span class="p_btn50--4th san_paging__btn">1</span> </li><li id="san_pagingJumpTo2" class="san_paging__item"><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wPTImcHM9MzA=" data-tracking='{"san_NaviPaging":"2"}'> <p class="p_btn50--4th san_paging__btn">2</p> </span></li><li id="san_pagingJumpTo3" class="san_paging__item"><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wPTMmcHM9MzA=" data-tracking='{"san_NaviPaging":"3"}'> <p class="p_btn50--4th san_paging__btn">3</p> </span></li><li id="san_pagingJumpTo4" class="san_paging__item"><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wPTQmcHM9MzA=" data-tracking='{"san_NaviPaging":"4"}'> <p class="p_btn50--4th san_paging__btn">4</p> </span></li><li id="san_pagingJumpTo5" class="san_paging__item"><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wPTUmcHM9MzA=" data-tracking='{"san_NaviPaging":"5"}'> <p class="p_btn50--4th san_paging__btn">5</p> </span></li> <li id="san_pagingDots" class="san_paging__item"><span class="san_paging__dots">...</span></li>
<li id="san_pagingJumpTo15" class="san_paging__item"><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wPTE1JnBzPTMw" data-tracking='{"san_NaviPaging":"15"}'> <p class="p_btn50--4th san_paging__btn">15</p> </span></li>
<li id="san_pagingBottomNext" class="san_paging__item"><span class="ub64e" data-ub64e="L3N1Y2hlL3dhc2NobWFzY2hpZW5lLz9wPTImcHM9MzA=" data-tracking='{"san_NaviPaging":"2"}'> <p class="p_btn50--1st san_paging__btn" title="weiter"><i>&gt;</i></p> </span></li> </ul></div>
    
    <div id="san_av_bottom"></div>
    <link href="/wato-onsite/ft6.wato.onsite.814f440c.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)" /><noscript><link rel="stylesheet" href="/wato-onsite/ft6.wato.onsite.814f440c.css" crossorigin="anonymous" /></noscript>
<link rel="preload" crossorigin="anonymous" href="/wato-onsite/ft6.wato.onsite.049fc5ae.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)" />

    <div class="clear"></div>
    
</div>


            <div class="js_tracking" data-track="true"></div>
                    
                    

<!-- P13N -->
<div class="js_p13nShoppingHistory"></div>

        <div class="clear"></div>
    </div>
        <div class="clear"></div>
        <footer>
            <link href="/shoppages/static/assets/sho.shoppages.assets-public.fb0bb0b7.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)"/><noscript><link rel="stylesheet" href="/shoppages/static/assets/sho.shoppages.assets-public.fb0bb0b7.css" crossorigin="anonymous"/></noscript>
<link rel="preload" crossorigin="anonymous" href="/shoppages/static/assets/sho.shoppages.assets-public.811e7ed1.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/><div class="shoppages sp_footerNormal">
            <div class="user_system_rwd">
<link href="/user/assets/ft4.user.newsletter-snippet.6d514ffc.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)"/><noscript><link rel="stylesheet" href="/user/assets/ft4.user.newsletter-snippet.6d514ffc.css" crossorigin="anonymous"/></noscript>
<link rel="preload" crossorigin="anonymous" href="/user/assets/ft4.user.newsletter-snippet.d387de48.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/>        <div id="us_id_newsletterSnippet" class="us_newsletterSnippet">
        </div>
    </div>

    <div class="sp_topBorderWrapper">
        <div class="sp_topBorder"></div>
    </div>

    <!-- ########################### First Row ###################################################### -->
    <div class="sp_footerColumn sp_accountColumn sp_firstrow">
    <div class="sp_loginFooterArea">
        <div class="sp_loginAreaContainerWithName">
            <a class="sp_loginFooterLink sp_loginIcon ts-link"
               data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_MeinKonto&quot;}" href="">
                <div class="sp_loginArea">
                    <span class="sp_loginAreaIcon p_icons"></span>
                    <span class="sp_loginAreaIconBadge p_badge100 p_badge100--icon" style="display:none;"></span>
                </div>
            </a>

            <div class="sp_loginAreaNameContainer">
                <a class="sp_loginFooterLink ts-link" data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_MeinKonto&quot;}"
                   href="">
                    <div class="sp_loginAreaMyAccountHeadline p_link100--2nd">Mein Konto</div>
                </a>
                <a class="sp_loginFooterLink sp_loginAreaMyNameOrLoginLink p_link100--2nd sp_loginAreaMyNameOrLogin ts-link"
                   data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_MeinKonto&quot;}"
                   href="">
                </a>
                <a class="sp_loginAreaNotMyName p_link100--1st ts-link"
                   data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_ MeinKonto_notMyName&quot;}"
                   style="display:none;"></a>
            </div>
        </div>
    </div>
    </div>
    <div class="sp_footerColumn sp_paymentColumn sp_firstrow">
        <ul>
            <li>
                <div class="sp_columnEntryWrapper">
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd sp_iconLink"
                          data-ub64e="L3Avb3R0by1nZXNjaGVua2d1dHNjaGVpbi12b24tMTAtMjUwLWV1cm8tMTAwNjM0MTQz"
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Geschenk-Gutscheine&quot;}"><span
                            class="sp_ColumnIcon p_icons">g</span></span>
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd"
                          data-ub64e="L3Avb3R0by1nZXNjaGVua2d1dHNjaGVpbi12b24tMTAtMjUwLWV1cm8tMTAwNjM0MTQz"
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Geschenk-Gutscheine&quot;}">Geschenk-Gutscheine</span>
                </div>
            </li>
            <li>
                <div class="sp_columnEntryWrapper">
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd sp_iconLink"
                          data-ub64e="L2Jlc3RlbGxzY2hlaW4="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Online-Bestellschein&quot;}"><span
                            class="sp_ColumnIcon p_icons">b</span></span>
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd"
                          data-ub64e="L2Jlc3RlbGxzY2hlaW4="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Online-Bestellschein&quot;}">Online-Bestellschein</span>
                </div>
            </li>
        </ul>
    </div>
    <div class="sp_footerColumn sp_serviceColumn sp_firstrow">
        <ul>
            <li>
                <div class="sp_columnEntryWrapper">
                    <span class="sp_columnEntry ub64e ts-link p_link100--2nd sp_iconLink"
                          data-ub64e="L29yZGVyL2NhdGFsb2dz"
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_OTTO-Kataloge&quot;}"><span
                            class="sp_ColumnIcon p_icons">k</span></span>
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd"
                          data-ub64e="L29yZGVyL2NhdGFsb2dz"
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_OTTO-Kataloge&quot;}">OTTO-Kataloge</span>
                </div>
            </li>
            <li>
                <div class="sp_columnEntryWrapper sp_js_trackFeedback">
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd sp_iconLink"
                          data-ub64e="IGphdmFzY3JpcHQ6dm9pZCgwKTs="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_LobUndKritikBig&quot;}"><span
                            class="sp_ColumnIcon p_icons">L</span></span>
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd"
                          data-ub64e="IGphdmFzY3JpcHQ6dm9pZCgwKTs="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_LobUndKritikBig&quot;}">Lob &amp; Kritik</span>
                </div>
            </li>
        </ul>
    </div>
    <div class="sp_footerColumn sp_aboutColumn sp_firstrow sp_hideColumnWhenNotExtraLarge">
        <ul>
            <li>
                <div class="sp_iconContainer">
                    <span class="ub64e sp_link ts-link p_link100--2nd sp_ehi sp_icon"
                          data-ub64e="aHR0cHM6Ly9laGktc2llZ2VsLmRlL3ZlcmJyYXVjaGVyL3Nob3BzLW1pdC1zaWVnZWwvemVydGlmaXppZXJ0ZS1zaG9wcy96ZXJ0aWZpa2F0L2M3MWUyNGM5ZmRmNjgxMGQxYjY5NDc5YzcxMDk5YjJmLw=="
                          data-target="_blank"
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Shop&quot;}">&nbsp;</span>
                </div>
            </li>
        </ul>
    </div>

    <div class="sp_leftTopBorder">
        <div class="sp_topBorder"></div>
    </div>
    <div class="sp_rightTopBorder">
        <div class="sp_topBorder"></div>
    </div>


    <!-- ########################### Second Row ###################################################### -->

    <div class="sp_accountColumn sp_footerColumn sp_secondrow">
        <ul>
            <li>
                <div class="sp_columnEntryWrapper sp_columnHeadline">
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd sp_iconLink"
                          data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlLw=="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Otto-Service&quot;}"><span
                            class="sp_ColumnIcon p_icons">s</span></span>
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd"
                          data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlLw=="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Otto-Service&quot;}">Service</span>
                </div>
            </li>
            <li class="sp_columnEntry">
                <span class="p_copy100--2nd">Haben Sie Fragen?</span>

                <div class="sp_phoneNumber">
                    <span class="ub64e sp_link sp_hideWhenLarge ts-link p_headline75" data-ub64e="dGVsOis0OTE4MDYzMDMwMzA="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_ServiceHotline&quot;}">01806 - 30 30 30</span>
                    <span class="sp_hideWhenNotLarge p_headline75">01806 - 30 30 30</span>
                </div>
                <div class="p_copy75--2nd">
                    <div>Festnetz 20 Cent/Anruf,</div>
                    <div>Mobilfunk max.&nbsp;60&nbsp;Cent/Anruf</div>
                </div>
            </li>
            <li class="sp_columnEntry">
                <a class="js_openInPopup sp_link ts-link p_link100--2nd"
                   href="/user/callbackPopup"
                   data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_RueckrufService&quot;}">Kostenloser Rückruf-Service</a>
            </li>
            <li class="sp_lastColumnEntry">
                <a class="js_openInPopup sp_link ts-link p_link100--2nd" href="/user/contactForm"
                   data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Kontakt&quot;}"
                   target="_blank">service@otto.de</a>
            </li>
            <li class="sp_columnEntry sp_hideWhenNotExtraLarge sp_appStores">
                <a class="sp_link ts-link sp_apple-app-store-badge-de sp_icon"
                   href="https://app.adjust.com/ig5rmn?fallback=https%3A%2F%2Fitunes.apple.com%2Fde%2Fapp%2Fotto-shopping-fur-mode-wohnen%2Fid404844644%3Fmt%3D8"
                   data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_AppStoreApple&quot;}"
                   target="_blank"></a>
                <a class="sp_link ts-link sp_google-app-store-badge-de sp_icon"
                   href="https://app.adjust.com/ne03zg"
                   data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_AppStoreAndroid&quot;}"
                   target="_blank"></a>
            </li>
        </ul>
    </div>
    <div class="sp_footerColumn sp_mobilerow">
        <div class="sp_topBorder"></div>
        <div class="sp_mobilColumnWrapper">
            <div class="sp_mobilColumn">
                <ul>
                    <li>
                        <div class="sp_columnEntryWrapper">
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd sp_iconLink"
                          data-ub64e="L3Avb3R0by1nZXNjaGVua2d1dHNjaGVpbi12b24tMTAtMjUwLWV1cm8tMTAwNjM0MTQz"
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Geschenk-Gutscheine&quot;}"><span
                            class="sp_ColumnIcon p_icons">g</span></span>
                            <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd"
                                  data-ub64e="L3Avb3R0by1nZXNjaGVua2d1dHNjaGVpbi12b24tMTAtMjUwLWV1cm8tMTAwNjM0MTQz"
                                  data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Geschenk-Gutscheine&quot;}">Geschenk-Gutscheine</span>
                        </div>
                    </li>
                    <li>
                        <div class="sp_columnEntryWrapper">
                    <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd sp_iconLink"
                          data-ub64e="L2Jlc3RlbGxzY2hlaW4="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Online-Bestellschein&quot;}"><span
                            class="sp_ColumnIcon p_icons">b</span></span>
                            <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd"
                                  data-ub64e="L2Jlc3RlbGxzY2hlaW4="
                                  data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Online-Bestellschein&quot;}">Online-Bestellschein</span>
                        </div>
                    </li>
                </ul>
            </div>
            <div class="sp_mobilColumn">
                <ul>
                    <li>
                        <div class="sp_columnEntryWrapper">
                    <span class="sp_columnEntry ub64e ts-link p_link100--2nd sp_iconLink"
                          data-ub64e="L29yZGVyL2NhdGFsb2dz"
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_OTTO-Kataloge&quot;}"><span
                            class="sp_ColumnIcon p_icons">k</span></span>
                            <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd"
                                  data-ub64e="L29yZGVyL2NhdGFsb2dz"
                                  data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_OTTO-Kataloge&quot;}">OTTO-Kataloge</span>
                        </div>
                    </li>
                    <li>
                        <div class="sp_columnEntryWrapper sp_js_trackFeedback">
                            <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd sp_iconLink"
                                  data-ub64e="IGphdmFzY3JpcHQ6dm9pZCgwKTs="
                                  data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_LobUndKritikSmall&quot;}"><span
                                    class="sp_ColumnIcon p_icons">L</span></span>
                            <span class="sp_columnEntry ub64e sp_link ts-link p_link100--2nd"
                                  data-ub64e="IGphdmFzY3JpcHQ6dm9pZCgwKTs="
                                  data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_LobUndKritikSmall&quot;}">Lob &amp; Kritik</span>
                        </div>
                    </li>
                </ul>
            </div>
        </div>
    </div>
    <div class="sp_paymentColumn sp_secondrow sp_footerColumn sp_footerAccordion">
        <div class="sp_topBorder"></div>
        <ul>
            <li class="sp_columnHeadline p_headline25 sp_footerAccordionHeader js_footerAccordionHeader">Zahlungsarten</li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link sp_logo_payment_credit sp_icon"
                      data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL3BheW1lbnQjNTEyMjAxNGJlNGIwNThlZWRlNzMxNTY5"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Kreditkarten&quot;}">&nbsp;</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link sp_logo_paypal sp_icon"
                      data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL3BheW1lbnQjNTgyZGJlZGMzMjIzZTg1OTMyYTNiMjUy"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_PayPal&quot;}">&nbsp;</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link sp_logo_paydirekt sp_icon"
                      data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL3BheW1lbnQjNTlkYzg2M2YxNjgwMDUwMDAxYjU5NGZm"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_PayDirekt&quot;}">&nbsp;</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd"
                      data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL3BheW1lbnQjNTExZTYwZGJlNGIwMzkyOGFlOGMyNmVm"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Rechnung&quot;}">Rechnung</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd"
                      data-ub64e="L3Nob3BwYWdlcy9scF9tb25hdHNyYXRlbg=="
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Ratenzahlung&quot;}">Ratenzahlung*</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd"
                      data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL3BheW1lbnQjNTExZTU1ZjRlNGIwYmQ4ODRjNTgyZjFl"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Zahlpause&quot;}">Zahlpause*</span>
            </li>
            <li class="sp_lastColumnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd"
                      data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL3BheW1lbnQjNTEyMjA2ZWZlNGIwMzkyOGFlOGMyNmYw"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Vorkasse&quot;}">Vorkasse</span>
            </li>
        </ul>
    </div>

    <div class="sp_aboutColumn sp_footerColumn sp_footerAccordion sp_secondrow">
        <div class="sp_topBorder"></div>
        <ul>
            <li class="sp_columnHeadline p_headline25 sp_footerAccordionHeader js_footerAccordionHeader">Über uns</li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd" data-ub64e="L3Nob3BwYWdlcy9iZWdvb2Q="
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Nachhaltigkeit&quot;}">Nachhaltigkeit</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd" data-ub64e="L3VudGVybmVobWVu"
                      data-target="_blank"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Unternehmen&quot;}">Unternehmen</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent sp_lastelement">
                <span class="ub64e sp_link ts-link p_link100--2nd"
                      data-ub64e="L3VudGVybmVobWVuL2pvYnM="
                      data-target="_blank"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Jobs&quot;}">Jobs</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent sp_hideColumnEntryWhenNotLarge">
                        <span class="ub64e sp_link ts-link p_link100--2nd footerAgb"
                              data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL2FnYg=="
                              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_AGB&quot;}">AGB</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent sp_hideColumnEntryWhenNotLarge">
                        <span class="ub64e sp_link ts-link p_link100--2nd footerPrivacy"
                              data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL2RhdGVuc2NodXR6"
                              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Datenschutz&quot;}">Datenschutz</span>
            </li>
            <li class="sp_lastColumnEntry sp_footerAccordionContent sp_hideColumnEntryWhenNotLarge">
                <a class="sp_link ts-link p_link100--2nd footerImprint" href="/shoppages/service/impressum"
                   data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Impressum&quot;}">Impressum</a>
            </li>
        </ul>
    </div>

    <div class="sp_serviceColumn sp_footerColumn sp_footerAccordion sp_secondrow">
        <div class="sp_topBorder"></div>
        <ul>
            <li class="sp_columnHeadline p_headline25 sp_footerAccordionHeader js_footerAccordionHeader">OTTO Partner</li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link sp_logo_ottomarket sp_icon ts-link sp_Icon"
                      data-ub64e="aHR0cHM6Ly93d3cub3R0by5tYXJrZXQv"
                      data-target="_blank"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_OttoMarket&quot;}">
                </span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd" data-ub64e="aHR0cDovL3d3dy5zaG9wcGluZzI0LmRlLw=="
                      data-target="_blank"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Shopping24&quot;}">Shopping24</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd"
                      data-ub64e="L3Nob3BwYWdlcy92ZXJzaWNoZXJ1bmdlbg=="
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Versicherungen&quot;}">Versicherungen</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd"
                      data-ub64e="L3Nob3BwYWdlcy9zaG9wcGluZy1tb3Jl"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_ShoppingMore&quot;}">Shopping&amp;more</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
            <span class="ub64e sp_link ts-link p_link100--2nd" data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL3Zvcm9ydA=="
                  data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_OttoInIhrerNaehe&quot;}">OTTO in Ihrer Nähe</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd" data-ub64e="aHR0cDovL2FmZmlsaWF0ZS5vdHRvLmRlLw=="
                      data-target="_blank"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Affiliate&quot;}">OTTO Affiliate</span>
            </li>
            <li class="sp_lastColumnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd"
                      data-ub64e="L3Nob3BwYWdlcy93ZXJidW5nLW90dG8="
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_WerbungAtOtto&quot;}">Werbung@OTTO</span>
            </li>
        </ul>
    </div>

    <!-- ### THIRD ROW SHOW ONLY IN L AND BELOW ############################################## -->
    <div class="sp_accountColumn sp_footerColumn sp_thirdrow">

        <div class="sp_iconContainer">
                    <span class="ub64e sp_link ts-link p_link100--2nd sp_ehi sp_icon"
                          data-ub64e="aHR0cHM6Ly9laGktc2llZ2VsLmRlL3ZlcmJyYXVjaGVyL3Nob3BzLW1pdC1zaWVnZWwvemVydGlmaXppZXJ0ZS1zaG9wcy96ZXJ0aWZpa2F0L2M3MWUyNGM5ZmRmNjgxMGQxYjY5NDc5YzcxMDk5YjJmLw=="
                          data-target="_blank"
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Shop&quot;}">&nbsp;</span>
        </div>
        <div class="sp_centered sp_appStores">
            <a class="sp_link ts-link sp_apple-app-store-badge-de sp_icon"
               href="https://app.adjust.com/ig5rmn?fallback=https%3A%2F%2Fitunes.apple.com%2Fde%2Fapp%2Fotto-shopping-fur-mode-wohnen%2Fid404844644%3Fmt%3D8"
               data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_AppStoreApple&quot;}"
               target="_blank"></a>
            <a class="sp_link ts-link sp_google-app-store-badge-de sp_icon"
               href="https://app.adjust.com/ne03zg"
               data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_AppStoreAndroid&quot;}"
               target="_blank"></a>
        </div>

    </div>

    <div class="sp_paymentColumn sp_footerColumn sp_thirdrow">
        <ul>
            <li class="sp_columnHeadline p_headline25 sp_footerAccordionHeader js_footerAccordionHeader">OTTO Partner</li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link sp_logo_ottomarket sp_icon ts-link sp_Icon"
                      data-ub64e="aHR0cHM6Ly93d3cub3R0by5tYXJrZXQv"
                      data-target="_blank"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_OttoMarket&quot;}">
                </span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd" data-ub64e="aHR0cDovL3d3dy5zaG9wcGluZzI0LmRlLw=="
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Shopping24&quot;}">Shopping24</span>
            </li>
            <li class="sp_columnEntry sp_footerAccordionContent">
                <span class="ub64e sp_link ts-link p_link100--2nd" data-ub64e="L3Nob3BwYWdlcy92ZXJzaWNoZXJ1bmdlbg=="
                      data-target="_blank"
                      data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Versicherungen&quot;}">Versicherungen</span>
            </li>
        </ul>
    </div>

    <div class="sp_serviceColumn sp_footerColumn sp_thirdrow">
        <ul>
            <li class="sp_columnEntry">
                        <span class="ub64e sp_link ts-link p_link100--2nd"
                              data-ub64e="L3Nob3BwYWdlcy9zaG9wcGluZy1tb3Jl"
                              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_ShoppingMore&quot;}">Shopping&amp;more</span>
            </li>
            <li class="sp_columnEntry">
                        <span class="ub64e sp_link ts-link p_link100--2nd"
                              data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL3Zvcm9ydA=="
                              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_OttoInIhrerNaehe&quot;}">OTTO in Ihrer Nähe</span>
            </li>
            <li class="sp_columnEntry">
                    <span class="ub64e sp_link ts-link p_link100--2nd"
                          data-ub64e="aHR0cDovL2FmZmlsaWF0ZS5vdHRvLmRlLw=="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Affiliate&quot;}">OTTO Affiliate</span>
            </li>
            <li class="sp_columnEntry">
                    <span class="ub64e sp_link ts-link p_link100--2nd"
                          data-ub64e="L3Nob3BwYWdlcy93ZXJidW5nLW90dG8="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_WerbungAtOtto&quot;}">Werbung@OTTO</span>
            </li>
        </ul>
    </div>

    <!-- BREADCRUMB only show in Mobile Version ###################### -->

    <div class="sp_footerColumn sp_mobilerow">
        <div class="sp_topBorder"></div>
        <ul class="sp_footer_mobile_breadcrumb">
            <li class="">
                        <span class="ub64e sp_link ts-link p_link100--2nd footerAgb"
                              data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL2FnYg=="
                              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_AGB&quot;}">AGB</span>
            </li>
            <li class="">|</li>
            <li class="">
                        <span class="ub64e sp_link ts-link p_link100--2nd footerPrivacy"
                              data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL2RhdGVuc2NodXR6"
                              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Datenschutz&quot;}">Datenschutz</span>
            </li>
            <li class="">|</li>
            <li class="">
                <a class="sp_link ts-link p_link100--2nd footerImprint" href="/shoppages/service/impressum"
                   data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Impressum&quot;}">Impressum</a>
            </li>
        </ul>
    </div>

    <div class="sp_footerColumn sp_mobilerow">
        <div class="sp_iconContainer sp_centered">
            <span class="ub64e sp_link ts-link p_link100--2nd sp_ehi sp_icon"
                  data-ub64e="aHR0cHM6Ly9laGktc2llZ2VsLmRlL3ZlcmJyYXVjaGVyL3Nob3BzLW1pdC1zaWVnZWwvemVydGlmaXppZXJ0ZS1zaG9wcy96ZXJ0aWZpa2F0L2M3MWUyNGM5ZmRmNjgxMGQxYjY5NDc5YzcxMDk5YjJmLw=="
                  data-target="_blank"
                  data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Shop&quot;}">&nbsp;</span>
        </div>
    </div>

    <div class="sp_footerColumn sp_mobilerow">
        <div class="sp_centered sp_appStores">
            <a class="sp_link ts-link sp_apple-app-store-badge-de sp_icon"
               href="https://app.adjust.com/ig5rmn?fallback=https%3A%2F%2Fitunes.apple.com%2Fde%2Fapp%2Fotto-shopping-fur-mode-wohnen%2Fid404844644%3Fmt%3D8"
               data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_AppStoreApple&quot;}"
               target="_blank"></a>
            <a class="sp_link ts-link sp_google-app-store-badge-de sp_icon"
               href="https://app.adjust.com/ne03zg"
               data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_AppStoreAndroid&quot;}"
               target="_blank"></a>
        </div>
    </div>

    <div class="sp_ottoUspRowWrapper">
        <div class="sp_ottoUspRow">
            <div class="sp_ottoUspCenter">
                <div class="sp_ottoUsp">
                    <span class="sp_checkMark p_icons">!</span>
                    <span class="ub64e sp_link sp_footerUspDescr ts-link p_link100--2nd"
                          data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL3BheW1lbnQjNTExZTYwZGJlNGIwMzkyOGFlOGMyNmVm"
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_SichererKaufAufRechnung&quot;}">Sicherer Kauf auf Rechnung</span>
                </div>
                <div class="sp_ottoUsp">
                    <span class="sp_checkMark p_icons">!</span>
                    <span class="ub64e sp_link sp_footerUspDescr ts-link p_link100--2nd"
                          data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL3BheW1lbnQjNTExZTViOTZlNGIwMzkyOGFlOGMyNmVl"
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_EinfacheRatenzahlung&quot;}">Einfache Ratenzahlung*</span>
                </div>
                <div class="sp_ottoUsp">
                    <span class="sp_checkMark p_icons">!</span>
                    <span class="ub64e sp_link sp_footerUspDescr ts-link p_link100--2nd"
                          data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL2xpZWZlcnVuZyM1MTFlNDQwYWU0YjA5MjRjZTQ2ZDU0MDA="
                          data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_30TageRueckgabegarantie&quot;}">30 Tage Rückgabegarantie</span>
                </div>
            </div>
        </div>
    </div>
    <div class="sp_footerSocialMedia">
        <span class="ub64e sp_footer_social_pinterest sp_icon ts-link sp_footerSocialMediaIcon"
              data-ub64e="aHR0cHM6Ly9kZS5waW50ZXJlc3QuY29tL290dG9kZS8=" data-target="_blank"
              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Pinterest&quot;}"></span>
        <span class="ub64e sp_footer_social_facebook sp_icon ts-link sp_footerSocialMediaIcon"
              data-ub64e="aHR0cHM6Ly93d3cuZmFjZWJvb2suY29tL090dG8=" data-target="_blank"
              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Facebook&quot;}"></span>
        <span class="ub64e sp_link sp_footer_social_twitter sp_icon ts-link sp_footerSocialMediaIcon"
              data-ub64e="aHR0cHM6Ly90d2l0dGVyLmNvbS9vdHRvX2Rl"
              data-target="_blank"
              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Twitter&quot;}"></span>
        <span class="ub64e sp_link sp_footer_social_youtube sp_icon ts-link sp_footerSocialMediaIcon"
              data-ub64e="aHR0cDovL3d3dy55b3V0dWJlLmNvbS9vdHRv"
              data-target="_blank"
              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_YoutTube&quot;}">
        </span>
        <span class="ub64e sp_link sp_footer_social_instagram sp_icon ts-link sp_footerSocialMediaIcon"
              data-ub64e="aHR0cHM6Ly93d3cuaW5zdGFncmFtLmNvbS9vdHRvX2RlLw=="
              data-target="_blank"
              data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Instagram&quot;}">
        </span>
            <span class="ub64e sp_link sp_footer_social_whatsapp sp_icon ts-link sp_footerSocialMediaIcon"
                  data-ub64e="L3Nob3BwYWdlcy9XaGF0c2FwcC10aWNrZXI="
                  data-target="_blank"
                  data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_WhatsApp&quot;}">
            </span>
    </div>

    <div class="sp_footerBlogs">
        <a class="sp_link sp_footer_blogs_reblog sp_icon ts-link sp_footerBlogIcon" href="/reblog" target="_blank"
           data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_reBlog&quot;}">&nbsp;</a>
        <a class="sp_link sp_footer_blogs_platzschaffen sp_icon ts-link sp_footerBlogIcon"
           href="http://www.platzschaffenmitherz.de/?refid=OFoot1" target="_blank"
           data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_platzschaffen&quot;}">&nbsp;</a>
        <a class="sp_link sp_footer_blogs_roombeez sp_icon ts-link sp_footerBlogIcon" href="/roombeez" target="_blank"
           data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_roombeez&quot;}">&nbsp;</a>
        <a class="sp_link sp_footer_blogs_twoforfashion sp_icon ts-link sp_footerBlogIcon" href="/twoforfashion" target="_blank"
           data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_twoForFashion&quot;}">&nbsp;</a>
        <a class="sp_link sp_footer_blogs_soulfully sp_icon ts-link sp_footerBlogIcon" href="/soulfully" target="_blank"
           data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_Soulfully&quot;}">&nbsp;</a>
        <a class="sp_link sp_footer_blogs_updated sp_icon ts-link sp_footerBlogIcon" href="/updated"
           target="_blank" data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_UpdatedBlog&quot;}">&nbsp;</a>
        <a class="sp_link sp_osf_short sp_icon ts-link sp_festivalWeekBlockS" href="/inspiration/weltmeisterschaft"
           target="_blank" data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_FestivalWeek&quot;}">&nbsp;</a>
    </div>

    <div class="sp_festivalWeekBlockL">
        <a class="sp_link sp_osf_long sp_icon ts-link sp_footerBlogIcon" href="/inspiration/weltmeisterschaft"
           target="_blank" data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_FestivalWeek&quot;}">&nbsp;</a>
    </div>

    <div class="sp_veryFinePrint">
        <p class="p_copy75--2nd">Preisangaben inkl. gesetzl. MwSt. und zzgl. <span class="p_link--1st ub64e ts-link"
                                                                                   data-ub64e="L3Nob3BwYWdlcy9zZXJ2aWNlL2xpZWZlcnVuZw=="
                                                                                   data-ts-link="{&quot;ot_Origin&quot; : &quot;footer_ServiceUndVersandkosten&quot;}">Service- und Versandkosten</span>
        </p>

        <p class="p_copy75--2nd">* Bonität vorausgesetzt, gegen Aufpreis</p>
    </div>

    <div class="sp_footerOttoClaim">
        <span class="sp_ofig34x128">&nbsp;</span>
    </div>
</div>

        </footer>
    </div>
    
    <div class="infoContainer av_anchor"></div>
    <link href="/wato-onsite/ft6.wato.onsite.814f440c.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)" /><noscript><link rel="stylesheet" href="/wato-onsite/ft6.wato.onsite.814f440c.css" crossorigin="anonymous" /></noscript>
<link rel="preload" crossorigin="anonymous" href="/wato-onsite/ft6.wato.onsite.049fc5ae.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)" />

        

<link href="/shoppages/static/assets/sho.shoppages.assets-public.fb0bb0b7.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)"/><noscript><link rel="stylesheet" href="/shoppages/static/assets/sho.shoppages.assets-public.fb0bb0b7.css" crossorigin="anonymous"/></noscript>
<link rel="preload" crossorigin="anonymous" href="/shoppages/static/assets/sho.shoppages.assets-public.811e7ed1.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/><script src="/shoppages/static/assets/sho.shoppages.assets-synchronous.039983bb.js" integrity="sha256-cefznIO/arKwbUQmFkKsltHO6Tvifs6CYRGKvC1JqV4=" crossorigin="anonymous"></script>    <div data-title="" data-callback="o_shoppages.layer.init">
        <div class="shoppages shoppagesFragment">
    <script type="text/javascript">
/*jslint browser: true */
/*global o_global: true */
(function (loaderScript) {
  "use strict";
 
var script, firstScript, parts, len, i, mpathyVersion;
 
  function cookie(name){
    parts = document.cookie.split(/;\s*/);
    len = parts.length;
    name += '=';
    for (i = 0; i < len; i+=1) {
      if (parts[i].indexOf(name) !== -1) {
        return parts[i].replace(name, '');
      }
    }
  }
  o_global.eventLoader.onLoad(99, function() {
    if (o_global.mpathy.isEnabled()) {
      //if (cookie('mpathyEnabled') === "true") {
        if (cookie('trackingDisabled') !== "true") {
          mpathyVersion = cookie("mpathyVersion");
          if (mpathyVersion !== undefined && /^a2995\.[0-9]+\.js$/.test(mpathyVersion)) {
            loaderScript = "//cdn.m-pathy.com/js/otto/" + mpathyVersion;
          }
          
          
            script = document.createElement('script');
            script.src = loaderScript;
            script.type = 'text/javascript';
            script.async = true;
            firstScript = document.getElementsByTagName('script')[0];
            firstScript.parentNode.insertBefore(script, firstScript);
          
        }
     //}
   }
 });
}('//cdn.m-pathy.com/js/otto/a2995.18070216.js'));
</script>
        </div>
    </div>
    


    </div>
            <div data-name="body_static_files" style="display: none"></div>

<script src="/assets-static/global-resources/assets.global-resources.body.7aefd79f.js" integrity="sha256-hfWGd3kDFRQg50U8Hrff7O2Ocw5gaM36zTGc++EHct4=" crossorigin="anonymous"></script><link rel="preload" crossorigin="anonymous" href="/assets-static/global-resources/assets.global-resources.non-critical-body.0dbc0c63.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/>
<script src="/tracking-bct/tracking.bct.bct.476988fd.js" integrity="sha256-l0AjgxUqTTCcV7/Jw1BBg898oIDrdEdBArtMa+bpAok=" crossorigin="anonymous"></script>
<link rel="preload" crossorigin="anonymous" href="/scale-shop-assets/scale.shop-assets.scale_beacon.2cee3772.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/>
<link href="/apps-assets/apps.global.assets.93b8c639.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)"/><noscript><link rel="stylesheet" href="/apps-assets/apps.global.assets.93b8c639.css" crossorigin="anonymous"/></noscript>
<link rel="preload" crossorigin="anonymous" href="/apps-assets/apps.global.assets.062ffed8.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/>

<script type="text/javascript">o_global.assets={"js":{"all_72d7dcff5f9b5932":"https://www.otto.de/static/all/js/72d7dcff5f9b5932/public_non-critical_min.js","san_d4669455c49e9557":"https://www.otto.de/static/san/js/d4669455c49e9557/private_san_non-critical_min.js"},"css":{"all_f3a30817e4ab5efc":"https://www.otto.de/static/all/css/f3a30817e4ab5efc/files/public_non-critical_min.css","san_53978d4ac43edf7f":"https://www.otto.de/static/san/css/53978d4ac43edf7f/files/private_non-critical_min.css"}}</script>
  <script defer="defer" type="text/javascript" src="https://www.otto.de/static/all/js/1573a9aac30427f5/public_critical_min.js" crossorigin></script>
  <script defer="defer" type="text/javascript" src="https://www.otto.de/static/san/js/ddd8f69445adbfa5/private_san_critical_min.js" crossorigin></script>
  <noscript><img src="https://www.otto.de/static/all/img/global-resources/9cc2d231044d4951/beacons/ottonojs.gif" /></noscript>

    <div class="us_js_esiAjax"
     data-url="/user/sessionMaintenance"
     data-eventtype="load"
     data-priority="24"
     data-qa="user_session_maintenance">
</div>
<link href="/user/assets/ft4.user.public.a86e06cb.css" rel="preload" crossorigin="anonymous" as="style" onload="invokePreload.onStyleLoad(this)"/><noscript><link rel="stylesheet" href="/user/assets/ft4.user.public.a86e06cb.css" crossorigin="anonymous"/></noscript>
<link rel="preload" crossorigin="anonymous" href="/user/assets/ft4.user.public.b95c09c8.js" as="script" onload="invokePreload.onScriptLoad(this)" onerror="invokePreload.onScriptError(this)"/>
<div id="us_js_id_loginAreaFooter">
    <div id="us_js_id_loginAreaContentToRemove" class="user_system_rwd">
        <div class="us_hide">
            <div id="us_js_id_loginAreaContentToReplaceHeader">
                <div class="us_loginAreaContainerWithName" data-qa="user_login_area_header_container"
                     id="us_id_loginAreaContainerWithName">
                    <div class="us_loginArea us_js_loginAreaMenuHandle ub64e" id="us_id_loginArea" data-qa="user_login_area_header" data-ub64e="amF2YXNjcmlwdDp2b2lkKDApOw==">
<span class="us_loginAreaIcon p_icons" data-qa="user_login_area_icon">
Θ
</span>
                            <!--[if gt IE 8]><!-->
                            <span id="us_id_customerServiceWidgetBubble" class="us_loginAreaIconServiceBubble p_icons" data-qa="user_login_area_service_bubble">&#8264;</span><!--<![endif]-->
                        <span class="us_iconSubtitle us_js_loginAreaIconSubtitle"
                              data-qa="user_login_area_icon_subtitle"
                              data-loggedin="false"
                              data-login-state="UNKNOWN_VISITOR" >
                                        <span class="us_js_loginAreaIconSubtitleMyAccount"
                                              data-qa="user_loginAreaIconSubtitleMyAccount">Mein Konto</span>
                                    </span>
                    </div>
                    <div class="clear"></div>
                </div>

                <div class="us_loginMenu" id="us_id_loginAreaMenu" data-qa="user_login_area_menu">
<ul>
    <div class="us_loginAreaMenuItems ">
        <button id="us_id_closeLoginMenuButton" class="us_closeLoginMenuButton p_symbolBtn100--4th" type="button"
                data-qa="user_login_area_menu_close">
            <i>x</i>
        </button>
        <li class="login_area_overview_link">
            <span class="ub64e p_link100--2nd" data-ub64e="L3VzZXIvYWNjb3VudE92ZXJ2aWV3P2VudHJ5UG9pbnQ9bG9naW5BcmVh"
                  data-qa="user_login_area_menu_myaccount">
Mein Konto            </span>
        </li>
        <li class="login_area_menu_link">
            <span class="ub64e p_link100--2nd" data-ub64e="L3VzZXIvb3JkZXJzP2VudHJ5UG9pbnQ9bG9naW5BcmVh" data-qa="user_login_area_menu_myorders">
Meine Bestellungen            </span>
        </li>
        <li class="login_area_menu_link">
            <span class="ub64e p_link100--2nd" data-ub64e="L3VzZXIvaW52b2ljZXM/ZW50cnlQb2ludD1sb2dpbkFyZWE=" data-qa="user_login_area_menu_myinvoices">
Meine Rechnungen            </span>
        </li>
        <li class="login_area_menu_link">
            <span class="user_js_show_more_login_area_links p_link100--2nd" data-qa="user_login_area_menu_show_more_links">
mehr...            </span>
        </li>

        <li class="login_area_menu_link login_area_menu_hide_link">
            <span class="user_js_more_login_area_links_follow p_link100--2nd" data-ub64e="L3VzZXIvbXlib29raW5ncz9lbnRyeVBvaW50PWxvZ2luQXJlYQ==" data-qa="user_login_area_menu_mybookings">
Meine Konto-Buchungen            </span>
        </li>
        <li class="login_area_menu_link login_area_menu_hide_link">
            <span class="user_js_more_login_area_links_follow p_link100--2nd" data-ub64e="L3VzZXIvbXlkYXRhP2VudHJ5UG9pbnQ9bG9naW5BcmVh" data-qa="user_login_area_menu_mydata">
Meine persönlichen Daten            </span>
        </li>
        <li class="login_area_menu_link login_area_menu_hide_link">
            <span class="user_js_more_login_area_links_follow p_link100--2nd" data-ub64e="L3VzZXIvYWRkcmVzc2VzP2VudHJ5UG9pbnQ9bG9naW5BcmVh" data-qa="user_login_area_menu_myaddresses">
Meine Anschriften            </span>
        </li>
        <li class="login_area_menu_link login_area_menu_hide_link">
            <span class="user_js_more_login_area_links_follow p_link100--2nd" data-ub64e="L3VzZXIvbXlzZXR0aW5ncz9lbnRyeVBvaW50PWxvZ2luQXJlYQ==" data-qa="user_login_area_menu_mysettings">
Meine Einstellungen            </span>
        </li>
    </div>
        <li class="login_area_menu_button">
            <span class="ub64e p_btn100--1st" data-ub64e="L3VzZXIvbG9naW4/ZW50cnlQb2ludD1sb2dpbkFyZWE=" data-qa="user_login_area_login">
Anmelden            </span>
        </li>
            <li class="login_area_register_link">
                <span class="p_link100--1st ub64e" data-ub64e="L3VzZXIvcmVnaXN0ZXI/ZW50cnlQb2ludD1sb2dpbkFyZWE=" data-qa="user_login_area_register">
Neu bei OTTO? Jetzt registrieren                </span>
            </li>
</ul>
<!--[if gt IE 8]><!-->
<div class="us_customerServiceWidgetAdvertising us_js_customerServiceWidgetLoginAreaLink">
    <div class="us_customerServiceWidgetIconContainer">
        <div class="us_customerServiceWidgetIcon">&#8264;</div>
    </div>
    <div class="us_customerServiceWidgetLinkContainer">
        <div class="p_link100--1st" data-qa="user_login_area_customer_service_widget_questions">Haben Sie Fragen oder Kritik?</div>
    </div>
    <div class="clear"></div>
</div><!--<![endif]-->
                </div>
    <div class="us_customerServiceWidgetSMNotification" id="us_id_customerServiceWidgetSMNotification">
        <div class="us_customerServiceWidgetSMNotificationContentContainer">
            <button id="us_id_customerServiceWidgetSMNotificationCloseButton" type="button"
                    class="us_customerServiceWidgetPopupMenuCloseButton p_symbolBtn50--4th"
                    data-qa="user_customer_service_widget_sm_notification_close_button">
                <i>x</i>
            </button>
            <div class="us_customerServiceWidgetSMNotificationContent">
                <div class="us_customerServiceWidgetSMNotificationIconContainer">
                    <div class="us_customerServiceWidgetIcon">&#8264;</div>
                </div>
                <div class="us_customerServiceWidgetSMNotificationText">
                    <div class="p_copy75 p_copyBold">Haben Sie Fragen oder Kritik?</div>
                </div>
                <div class="clear"></div>
            </div>
        </div>
    </div>
            </div>
        </div>
    </div>
</div>




<!-- temporary -->


    </body>
    </html>
`

func TestListArticles(t *testing.T) {
	raw := []byte(searchResult)
	_, err := listArticlesTok(raw)

	if err != nil {
		fmt.Println(err)
	}



}

func TestGetImgUrl(t *testing.T) {
	img, err := getImageUrl(testHtmlMay2018)

	assert.Nil(t, err)
	assert.Equal(t, "https://i.otto.de/i/otto/26390776/red-dead-redemption-2-xbox-one.jpg?$formatz$", img)
}

func TestGetPrice(t *testing.T) {
	price, err := getPrice(testHtmlMay2018)

	assert.Nil(t, err)
	assert.Equal(t, "69.99", price)
}

func TestGetName(t *testing.T) {
	name, err := getName(testHtmlMay2018)

	assert.Nil(t, err)
	assert.Equal(t, "Red Dead Redemption 2 Xbox One", name)
}

func TestCustomizeImgUrl(t *testing.T) {
	oldUrl := "https://i.otto.de/i/otto/26390776/red-dead-redemption-2-xbox-one.jpg?$formatz$"

	newUrl := customizeImgUrl(oldUrl)

	assert.Equal(t, "https://i.otto.de/i/otto/26390776?h=520&amp;w=384&amp;sm=clamp", newUrl)
}


