/**
 * GMtool
 * @author Jerry
 * @since 2022/11/04
 */
layui.use(['func', 'admin', 'common', "layer", "laydate"], function () {

    //声明变量
    var func = layui.func
        , layer = layui.layer
        , common = layui.common
        , laydate = layui.laydate
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            // {type: 'checkbox', fixed: 'left'}
            {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'time', width: 200, title: '设置时间', align: 'center', templet: function (d) {
                    return "<div>" + layui.util.toDateString(d.time, 'yyyy-MM-dd HH:mm:ss') + "</div>"
                }
            }
            , {field: 'create_user', width: 80, title: '创建人', align: 'center'}
            , {field: 'create_time', width: 180, title: '添加时间', align: 'center', templet: function (d) {
                    return "<div>" + layui.util.toDateString(d.create_time, 'yyyy-MM-dd HH:mm:ss') + "</div>"
                }
            }
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList", function (layEvent, data) {

        });

        //【设置弹框】
        func.setWin("设置时间", 500, 350);

        $("#settime").click((function () {

            layer.open({
                type:2
                ,area: ['500px', '450px']
                ,title:'设置时间'
                ,content: "/settime/set"
                ,maxmin: true
                ,btn: []
                ,end: function(index, layero){
                    location.reload()
                }
            })
        }))
    }

    if (A == "set") {
        //日期时间选择器
        laydate.render({
            elem: '#time'
            ,type: 'datetime'
        });
    }
});
