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

        var time = $("#time").val();
        time++
        var timer2 = setInterval(function () {
            time++
            let timeStr = timestampToTime(time)
            // console.log("time", timeStr)
            $("#server_time").text(timeStr)
        }, 1000);
    }

    if (A == "set") {
        //日期时间选择器
        laydate.render({
            elem: '#time'
            ,type: 'datetime'
        });
    }
});

function timestampToTime(timestamp) {
    var date = new Date(timestamp * 1000);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
    var Y = date.getFullYear() + '-';
    var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
    var D = (date.getDate() < 10) ? '0' + date.getDate() + ' ' : date.getDate() + ' ';
    var h = (date.getHours() < 10) ? '0' + date.getHours() + ':' : date.getHours() + ':';
    var m = (date.getMinutes() < 10) ? '0' + date.getMinutes() + ':' : date.getMinutes() + ':';
    var s = (date.getSeconds() < 10) ? '0' + date.getSeconds() : date.getSeconds();
    return Y+M+D+h+m+s;
}