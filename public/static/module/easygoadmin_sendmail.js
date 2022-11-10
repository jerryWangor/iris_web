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
            {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'role_list', width: 300, title: '角色ID列表', align: 'center'}
            , {field: 'item_list', width: 500, title: '道具列表', align: 'center', templet: function (d) {
                    return '<span>'+JSON.stringify(d.item_list)+'</span>';
                }
            }
            , {field: 'status', width: 120, title: '状态', align: 'center', templet(d) {
                    if (d.status == 0) {
                        return '<span class="layui-bg-cyan">未发送</span>';
                    } else if (d.status == 1) {
                        return '<span class="layui-bg-green">发送成功</span>';
                    } else if (d.status == 2) {
                        return '<span class="layui-bg-red">发送失败</span>';
                    }
                }
            }
            , {field: 'return_info', width: 300, title: '接口返回信息', align: 'center'}
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
        // func.setWin("设置时间", 500, 350);

        $("#sendmail").click((function () {
            layer.open({
                type:2
                ,area: ['700px', '550px']
                ,title:'发送奖励'
                ,content: "/sendmail/mail"
                ,maxmin: true
                ,btn: []
                ,end: function(index, layero){
                    location.reload()
                }
            })
        }));
    }

});