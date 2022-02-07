(function (a) {
    window.fsFontLoader = {};
    var b = !1,
        c = {
            load: function (a) {
                typeof a.webfonts == "object" && typeof a.webfonts.length == "undefined" && (a.webfonts = [a.webfonts]); var b = a.webfonts.length > a.batchThreshold; return d(a, b), this
            }
        }, d = function (c, d) {
            var j = c.webfonts, k = 0; a.each(j, function (l, m) {
                c.loadingClasses && h(m.css_class), e(c, m.css_class) ? (c.loadingClasses && i(m.css_class), typeof c.onSuccess == "function" && c.onSuccess(), typeof c.onComplete == "function" && c.onComplete()) : a.ajax(m.url + "?token=" + c.securityToken, {
                    success: function (a) { k++, g(c, m.css_class), typeof c.onSuccess == "function" && c.onSuccess() }, error: function (a) {
                        typeof c.onError == "function" && c.onError()
                    }, complete: function (a) {
                        var e = a.responseText; c.loadingClasses && i(m.css_class), typeof c.onComplete == "function" && c.onComplete(), !b && d ? l + 1 >= j.length && c.stylesHTMLContent ? (b = !0, f(e, !0, c)) : f(e, k % c.batchThreshold === 0, c) : f(e, !0, c)
                    }
                })
            })
        }, e = function (b, c) {
            return a.inArray(c, b.loadedTypefaces) > -1
        }, f = function (b, c, d) {
            d.stylesHTMLContent += b, c && (a("<style type='text/css' class='font-loader-injection'>" + d.stylesHTMLContent + "</style>").appendTo(a("head")), d.stylesHTMLContent = "")
        }, g = function (a, b) {
            a.loadedTypefaces.push(b)
        }, h = function (b) {
            a("." + b).addClass("wf-is-loading")
        }, i = function (b) { a("." + b).removeClass("wf-is-loading").addClass("wf-has-loaded") }; a.fontLoader = function (b, d) {
            var e = a.extend({}, a.fontLoader.defaults, d); return typeof window.fsFontLoader[e.nameSpace] == "undefined" ? e.loadedTypefaces = window.fsFontLoader[e.nameSpace] = [] : e.loadedTypefaces = window.fsFontLoader[e.nameSpace], c[b] ? c[b].apply(this, [e]) : typeof b == "object" || !b ? c.load.apply(this, [e]) : (a.error("Method " + b + " does not exist on jquery.font-loader.js"), a(this))
        }, a.fontLoader.defaults = {
            nameSpace: "loadedTypefaces", loadedTypefaces: [], loadingClasses: !0, batchThreshold: 10, stylesHTMLContent: ""
        }
})(jQuery)


  // Load the cover first,------ ffd07a20-19cd-4bf5-83fe-16eda96730d6
  //fast.fonts.net/dv2/3/1ffbded4-7bb0-4dae-95e1-fcb78b1041e2.woff?d44f19a684109620e4841471ae90e8186df205ea3b7e1c0d250a32d125adc27a69c5a6864d6ab62481f78c18b7a346bdbe432e4146d4f31e17e2d0f3f127fdd2eb43afc4cbcc0f9cb967ecc47dd3d385feb313630217444a5d
  //var primary_typefaces = {"url":"/webfonts/355f4a802e2bfc89f66e44415cccd6af.css","css_class":"webfont-214a31815ebf5bcac19b4b98831a54df","css_font_family":"214a31815ebf5bcac19b4b98831a5"};
  var primary_typefaces = {"url":"/webfonts/214a31815ebf5bcac19b4b98831a54df.css","css_class":"webfont-214a31815ebf5bcac19b4b98831a54df","css_font_family":"214a31815ebf5bcac19b4b98831a5"};
  $.fontLoader("load", {
    webfonts: primary_typefaces
  });
  // Then the rest of typefaces
  var secondary_typefaces = [{"url":"/webfonts/c0c06cf7cd84c8c568e3c0ad65871055.css","css_class":"webfont-c0c06cf7cd84c8c568e3c0ad65871055","css_font_family":"c0c06cf7cd84c8c568e3c0ad65871"},
  {"url":"/webfonts/c5633a2fce72f9376dd44f5f87aa4de5.css","css_class":"webfont-c5633a2fce72f9376dd44f5f87aa4de5","css_font_family":"c5633a2fce72f9376dd44f5f87aa4"},
  {"url":"/webfonts/60c47f5db1e31edb93d28d2d54ef39e3.css","css_class":"webfont-60c47f5db1e31edb93d28d2d54ef39e3","css_font_family":"60c47f5db1e31edb93d28d2d54ef3"}];
  $.fontLoader("load", {
    webfonts: secondary_typefaces
  });
