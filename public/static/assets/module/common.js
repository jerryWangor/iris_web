(() => {
    var n = {
        351: () => {
            layui.define(["form", "layer", "laydate", "upload", "element", "base"], (function (n) {
                "use strict";
                var e = layui.form, a = void 0 === parent.layer ? layui.layer : top.layer, t = layui.laydate,
                    o = layui.upload, i = (layui.element, layui.base), r = layui.$, l = {
                        edit: function (n, e = 0, a = 0, t = 0, o = [], r = null, c = !1) {
                            var s = e > 0 ? "修改" : "新增";
                            i.isEmpty(n) ? s += "内容" : s += n;
                            var u = cUrl + "/edit/" + e;
                            if (Array.isArray(o)) for (var f in o) u += "/" + o[f];
                            l.showWin(s, u, a, t, o, 2, [], (function (n, e) {
                                r && r(n, e)
                            }), c)
                        }, detail: function (n, e, a = 0, t = 0, o = !1) {
                            var i = cUrl + "/detail/" + e;
                            l.showWin(n + "详情", i, a, t, [], 2, [], null, o)
                        }, cache: function (n) {
                            var e = cUrl + "/cache";
                            l.ajaxPost(e, {id: n}, (function (n, e) {
                            }))
                        }, copy: function (n, e, a = 0, t = 0) {
                            var o = cUrl + "/copy/" + e;
                            l.showWin(n + "复制", o, a, t)
                        }, delete: function (n, e = null) {
                            a.confirm("您确定要删除吗？删除后将无法恢复！", {
                                icon: 3,
                                skin: "layer-ext-moon",
                                btn: ["确认", "取消"]
                            }, (function (t) {
                                var o = cUrl + "/delete/" + n;
                                console.log(o), l.ajaxPost(o, {}, (function (n, o) {
                                    e && (a.close(t), e(n, o))
                                }), "正在删除。。。")
                            }))
                        }, batchFunc: function (n, e = null) {
                            var t = n.url, o = n.title, i = (n.form, n.confirm || !1), r = n.show_tips || "处理中...",
                                c = n.data || [], s = n.param || [], u = n.type || "POST";
                            if ("导出数据" != o && 0 == c.length) return a.msg("请选择数据", {icon: 5}), !1;
                            var f = [];
                            for (var d in c) f.push(c[d].id);
                            var m = f.join(","), p = {};
                            if (p.id = m, Array.isArray(s)) for (var d in s) {
                                var y = s[d].split("=");
                                p[y[0]] = y[1]
                            }
                            console.log(p), i ? a.confirm("您确定要【" + o + "】选中的数据吗？", {
                                icon: 3,
                                title: "提示信息"
                            }, (function (n) {
                                "POST" == u ? t.indexOf("/delete") >= 0 ? l.ajaxPost(t + "/" + m, {}, e, r) : l.ajaxPost(t, p, e, r) : l.ajaxGet(t + "/" + m, {}, e, r)
                            })) : "POST" == u ? l.ajaxPost(t, p, e, r) : l.ajaxGet(t + "/" + m, {}, e, r)
                        }, verify: function () {
                            e.verify({
                                number: [/^[0-9]*$/, "请输入数字"], username: function (n, e) {
                                    return new RegExp("^[a-zA-Z0-9_一-龥\\s·]+$").test(n) ? /(^\_)|(\__)|(\_+$)/.test(n) ? title + "首尾不能出现下划线'_'" : /^\d+\d+\d$/.test(n) ? title + "不能全为数字" : void 0 : title + "不能含有特殊字符"
                                }, pass: [/^[\S]{6,12}$/, "密码必须6到12位，且不能出现空格"]
                            })
                        }, submitForm: function (n, e = null, a = null, t = !0) {
                            var o = [], c = [], s = n;
                            if (r.each(s, (function (n, e) {
                                if (console.log(n + ":" + e), /\[|\]|【|】/g.test(n)) {
                                    var a = n.match(/\[(.+?)\]/g);
                                    e = n.match("\\[(.+?)\\]")[1];
                                    var t = n.replace(a, "");
                                    r.inArray(t, o) < 0 && o.push(t), c[t] || (c[t] = []), c[t].push(e)
                                }
                            })), console.log(s), console.log(o), console.log(c), r.each(o, (function (n, e) {
                                var a = [];
                                r.each(c[e], (function (n, t) {
                                    a.push(t), delete s[e + "[" + t + "]"]
                                })), s[e] = a.join(",")
                            })), null == e) {
                                e = cUrl;
                                var u = r("form").attr("action");
                                i.isEmpty(u) ? null != n.id && (0 == n.id ? e += "/add" : n.id > 0 && (e += "/update")) : e = u
                            }
                            // console.log(s);
                            var f = new FormData;
                            r.each(s, (function (n, e) {
                                f.append(n, e), console.log(n + "," + e)
                            })), console.log(f), l.ajaxPost(e, f, (function (n, e) {
                                if (e) return t && setTimeout((function () {
                                    var n = parent.layer.getFrameIndex(window.name);
                                    parent.layer.close(n)
                                }), 500), a && a(n, e), !1
                            }))
                        }, searchForm: function (n, e, a = "tableList") {
                            n.reload(a, {page: {curr: 1}, where: e.field})
                        }, initDate: function (n, e = null) {
                            if (Array.isArray(n)) for (var a in n) {
                                var o = n[a].split("|");
                                if (o[2]) var i = o[2].split(",");
                                var r = {};
                                if (r.elem = "#" + o[0], r.type = o[1], r.theme = "molv", r.range = "true" === o[3] || o[3], r.calendar = !0, r.show = !1, r.position = "absolute", r.trigger = "click", r.btns = ["clear", "now", "confirm"], r.mark = {
                                    "0-06-25": "生日",
                                    "0-12-31": "跨年"
                                }, r.ready = function (n) {
                                }, r.change = function (n, e, a) {
                                }, r.done = function (n, a, t) {
                                    e && e(n, a)
                                }, i) {
                                    var l = i[0];
                                    if (l) {
                                        var c = !isNaN(l);
                                        r.min = c ? parseInt(l) : l
                                    }
                                    var s = i[1];
                                    if (s) {
                                        var u = !isNaN(s);
                                        r.max = u ? parseInt(s) : s
                                    }
                                }
                                t.render(r)
                            }
                        }, showWin: function (n, e, a = 0, t = 0, o = [], i = 2, l = [], c = null, s = !1) {
                            var u = layui.layer.open({
                                title: n,
                                type: i,
                                area: [a + "px", t + "px"],
                                content: e,
                                shadeClose: s,
                                shade: .4,
                                skin: "layui-layer-admin",
                                success: function (n, e) {
                                    if (Array.isArray(o)) for (var a in o) {
                                        var t = o[a].split("=");
                                        layui.layer.getChildFrame("body", e).find("#" + t[0]).val(t[1])
                                    }
                                    c && c(e, 1)
                                },
                                end: function () {
                                    c(u, 2)
                                }
                            });
                            0 == a && (layui.layer.full(u), r(window).on("resize", (function () {
                                layui.layer.full(u)
                            })))
                        }, ajaxPost: function (n, e, t = null, o = "处理中,请稍后...") {
                            var i = null;
                            r.ajax({
                                type: "POST",
                                url: n,
                                data: e,
                                dataType: "json",
                                contentType: !1,
                                processData: !1,
                                beforeSend: function () {
                                    i = a.msg(o, {icon: 16, shade: .01, time: 0})
                                },
                                success: function (n) {
                                    if (0 != n.code) return a.close(i), a.msg(n.msg, {icon: 5}), !1;
                                    a.msg(n.msg, {icon: 1, time: 500}, (function () {
                                        a.close(i), t && t(n, !0)
                                    }))
                                },
                                error: function () {
                                    a.close(i), a.msg("AJAX请求异常"), t && t(null, !1)
                                }
                            })
                        }, ajaxGet: function (n, e, t = null, o = "处理中,请稍后...") {
                            var i = null;
                            r.ajax({
                                type: "GET",
                                url: n,
                                data: e,
                                contentType: "application/json",
                                dataType: "json",
                                beforeSend: function () {
                                    i = a.msg(o, {icon: 16, shade: .01, time: 0})
                                },
                                success: function (n) {
                                    if (0 != n.code) return a.msg(n.msg, {icon: 5}), !1;
                                    a.msg(n.msg, {icon: 1, time: 500}, (function () {
                                        a.close(i), t && t(n, !0)
                                    }))
                                },
                                error: function () {
                                    a.msg("AJAX请求异常"), t && t(null, !1)
                                }
                            })
                        }, formSwitch: function (n, a = "", t = null) {
                            e.on("switch(" + n + ")", (function (e) {
                                var o = this.checked ? "1" : "2";
                                i.isEmpty(a) && (a = cUrl + "/set" + n.substring(0, 1).toUpperCase() + n.substring(1));
                                var r = new FormData;
                                r.append("id", this.value), r.append(n, o), console.log(r), l.ajaxPost(a, r, (function (n, e) {
                                    t && t(n, e)
                                }))
                            }))
                        }, uploadFile: function (n, e = null, t = "", r = "xls|xlsx", l = 10240, c = {}) {
                            i.isEmpty(t) && (t = cUrl + "/uploadFile"), o.render({
                                elem: "#" + n,
                                url: t,
                                auto: !1,
                                exts: r,
                                accept: "file",
                                size: l,
                                method: "post",
                                data: c,
                                before: function (n) {
                                    a.msg("上传并处理中。。。", {icon: 16, shade: .01, time: 0})
                                },
                                done: function (n) {
                                    return a.closeAll(), 0 == n.code ? a.alert(n.msg, {
                                        title: "上传反馈",
                                        skin: "layui-layer-molv",
                                        closeBtn: 1,
                                        anim: 0,
                                        btn: ["确定", "取消"],
                                        icon: 6,
                                        yes: function () {
                                            e && e(n, !0)
                                        },
                                        btn2: function () {
                                        }
                                    }) : a.msg(n.msg, {icon: 5}), !1
                                },
                                error: function () {
                                    return a.msg("数据请求异常")
                                }
                            })
                        }
                    };
                n("common", l)
            }))
        }
    }, e = {};
    !function a(t) {
        var o = e[t];
        if (void 0 !== o) return o.exports;
        var i = e[t] = {exports: {}};
        return n[t](i, i.exports, a), i.exports
    }(351)
})();