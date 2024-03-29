(() => {
    var t = {
        974: () => {
            layui.define(["form", "layer", "table", "common", "treeTable"], (function (t) {
                "use strict";
                var e, a, n, i, o, l = layui.form, r = layui.table, c = layui.layer, u = layui.common,
                    d = layui.treeTable, s = layui.$, f = 0, h = 0, m = !1, p = {
                        tableIns: function (t, l, d = null, p = "", b = !1) {
                            a = l, n = d, p && "" != p || (p = cUrl + "/list");
                            var y = s("#param").val();
                            if (y && (y = JSON.parse(y), s.isArray(y))) for (var v in y) p.indexOf("?") >= 0 ? p += "&" + y[v] : p += "?" + y[v];
                            return e = r.render({
                                elem: "#" + a,
                                url: p,
                                method: "post",
                                cellMinWidth: 150,
                                page: {
                                    layout: ["refresh", "prev", "page", "next", "skip", "count", "limit"],
                                    curr: 1,
                                    groups: 10,
                                    first: "首页",
                                    last: "尾页"
                                },
                                height: "full-100",
                                limit: 20,
                                limits: [20, 30, 40, 50, 60, 70, 80, 90, 100, 150, 200, 1e3],
                                even: !0,
                                cols: [t],
                                loading: !0,
                                done: function (t, e, a) {
                                    if (o) {
                                        var n = s(".layui-table-body").find("table").find("tbody");
                                        n.children("tr").on("dblclick", (function () {
                                            var e = n.find(".layui-table-hover").data("index"), a = t.data[e];
                                            u.edit(i, a.id, f, h)
                                        }))
                                    }
                                }
                            }), r.on("toolbar(" + a + ")", (function (t) {
                                var e = r.checkStatus(t.config.id);
                                switch (t.event) {
                                    case"getCheckData":
                                        var a = e.data;
                                        c.alert(JSON.stringify(a));
                                        break;
                                    case"getCheckLength":
                                        a = e.data, c.msg("选中了：" + a.length + " 个");
                                        break;
                                    case"isAll":
                                        c.msg(e.isAll ? "全选" : "未全选")
                                }
                            })), r.on("tool(" + a + ")", (function (t) {
                                var e = t.data, a = t.event;
                                "edit" === a ? u.edit(i, e.id, f, h, [], (function (t, e) {
                                    2 == e && s(".layui-laypage-btn").click()
                                }), m) : "detail" === a ? u.detail(i, e.id, f, h, m) : "del" === a ? u.delete(e.id, (function (e, a) {
                                    a && t.del()
                                })) : "cache" === a ? u.cache(e.id) : "copy" === a ? u.copy(i, e.id, f, h) : n && n(a, e)
                            })), r.on("checkbox(" + a + ")", (function (t) {
                            })), r.on("edit(" + a + ")", (function (t) {
                                var e = t.value, a = t.data, n = t.field, i = {};
                                i.id = a.id, i[n] = e;
                                var o = JSON.stringify(i), l = JSON.parse(o), r = cUrl + "/update";
                                u.ajaxPost(r, l, (function (t, e) {
                                }), "更新中...")
                            })), r.on("row(" + a + ")", (function (t) {
                                t.tr.addClass("layui-table-click").siblings().removeClass("layui-table-click"), t.data
                            })), b && r.on("sort(" + a + ")", (function (t) {
                                r.reload(a, {initSort: t, where: {field: t.field, order: t.type}})
                            })), this
                        }, treetable: function (t = [], e, n = !0, o = 0, l = "", r = null, m = "") {
                            a = e, m || (m = cUrl + "/list");
                            var p = d.render({
                                elem: "#" + e,
                                url: m,
                                method: "POST",
                                height: "full-50",
                                cellMinWidth: 80,
                                tree: {iconIndex: 1, idName: "id", pidName: l || "pid", isPidData: !0},
                                cols: [t],
                                done: function (t, e, a) {
                                    c.closeAll("loading")
                                },
                                style: "margin-top:0;"
                            });
                            d.on("tool(" + e + ")", (function (t) {
                                var e = t.data, a = t.event, n = e.id;
                                "add" === a ? u.edit(i, 0, f, h, ["pid=" + n], (function (t, e) {
                                    2 == e && location.reload()
                                })) : "edit" === a ? u.edit(i, n, f, h, [], (function (t, e) {
                                    2 == e && location.reload()
                                })) : "addz" === a ? u.edit(i, 0, f, h, ["pid=" + n], (function (t, e) {
                                    2 == e && location.reload()
                                })) : "del" === a ? u.delete(n, (function (e, a) {
                                    a && t.del()
                                })) : r && r(a, n, 0)
                            })), s("#collapse").on("click", (function () {
                                return p.foldAll(), !1
                            })), s("#expand").on("click", (function () {
                                return p.expandAll(), !1
                            })), s("#refresh").on("click", (function () {
                                return p.refresh(), !1
                            })), s("#search").click((function () {
                                var t = s("#keywords").val();
                                return t ? p.filterData(t) : p.clearFilter(), !1
                            }))
                        }, setWin: function (t, e = 0, a = 0, n = !1) {
                            return i = t, f = e, h = a, m = n, this
                        }, setDbclick: function (t) {
                            return o = t || !0, this
                        }, searchForm: function (t, e) {
                            l.on("submit(" + t + ")", (function (t) {
                                return u.searchForm(r, t, e), !1
                            }))
                        }, getCheckData: function (t) {
                            return t || (t = a), r.checkStatus(t).data
                        }, initDate: function (t, e = null) {
                            u.initDate(t, (function (t, a) {
                                e && e(t, a)
                            }))
                        }, showWin: function (t, e, a = 0, n = 0, i = [], o = 2, l = [], r = null, c = !1) {
                            u.showWin(t, e, a, n, i, o, l, (function (t, e) {
                                r && r(t, e)
                            }), c)
                        }, ajaxPost: function (t, e, a = null, n = "处理中...") {
                            u.ajaxPost(t, e, a, n)
                        }, ajaxGet: function (t, e, a = null, n = "处理中...") {
                            u.ajaxGet(t, e, a, n)
                        }, formSwitch: function (t, e = "", a = null) {
                            u.formSwitch(t, e, (function (t, e) {
                                a && a(t, e)
                            }))
                        }, uploadFile: function (t, e = null, a = "", n = "xls|xlsx", i = 10240, o = {}) {
                            u.uploadFile(t, (function (t, a) {
                                e && e(t, a)
                            }), a, n, i, o)
                        }
                    };
                u.verify(), l.on("submit(submitForm)", (function (t) {
                    return u.submitForm(t.field, null, (function (t, e) {
                        console.log("保存成功回调")
                    })), !1
                })), l.on("submit(searchForm)", (function (t) {
                    return u.searchForm(r, t), !1
                })), s(".btnOption").click((function () {
                    null != (n = s(this).attr("data-param")) && (console.log(n), n = JSON.parse(n), console.log(n));
                    var t = p.getCheckData(a);
                    switch (s(this).attr("lay-event")) {
                        case"add":
                            u.edit(i, 0, f, h, n, (function (t, e) {
                                2 == e && location.reload()
                            }), m);
                            break;
                        case"dall":
                            (o = {title: "批量删除"}).url = cUrl + "/delete", o.data = t, o.confirm = !0, o.type = "POST", u.batchFunc(o, (function () {
                                e.reload()
                            }));
                            break;
                        case"batchCache":
                            (o = {title: "批量重置缓存"}).url = cUrl + "/batchCache", o.data = t, o.confirm = !0, o.type = "GET", u.batchFunc(o, (function () {
                                e.reload()
                            }));
                            break;
                        case"batchEnable":
                            (o = {title: "批量启用状态"}).url = cUrl + "/batchStatus", o.param = n, o.data = t, o.form = "submitForm", o.confirm = !0, o.show_tips = "处理中...", o.type = "POST", u.batchFunc(o, (function () {
                                e.reload()
                            }));
                            break;
                        case"batchDisable":
                            (o = {title: "批量禁用状态"}).url = cUrl + "/batchStatus", o.param = n, o.data = t, o.confirm = !0, o.show_tips = "处理中...", o.type = "POST", u.batchFunc(o, (function () {
                                e.reload()
                            }));
                            break;
                        case"export":
                            var n = [], o = s(".layui-form-item [name]").serializeArray();
                            s.each(o, (function () {
                                n.push(this.name + "=" + this.value)
                            })), (o = {title: "导出数据"}).url = cUrl + "/export", o.data = t, o.confirm = !0, o.type = "POST", o.show_tips = "数据准备中...", o.param = n, u.batchFunc(o, (function (t, e) {
                                window.location.href = "/common/download?fileName=" + encodeURI(t.data) + "&isDelete=" + !0
                            }));
                            break;
                        case"import":
                            u.uploadFile("import", (function (t, e) {
                            }))
                    }
                })), window.formClose = function () {
                    var t = parent.layer.getFrameIndex(window.name);
                    parent.layer.close(t)
                }, t("func", p)
            }))
        }
    }, e = {};
    !function a(n) {
        var i = e[n];
        if (void 0 !== i) return i.exports;
        var o = e[n] = {exports: {}};
        return t[n](o, o.exports, a), o.exports
    }(974)
})();