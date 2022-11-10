/**
 * 道具管理
 * @author Jerry
 * @since 2021/7/26
 */
layui.use(['func', 'common', 'form', 'transfer'], function () {

    //声明变量
    var func = layui.func
        , common = layui.common
        , form = layui.form
        , transfer = layui.transfer
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {field: 'id', width: 80, title: 'ID', align: 'center', sort: true}
            , {field: 'type_id', width: 120, title: '道具ID', align: 'center', sort: true}
            , {field: 'name', width: 250, title: '道具名称', align: 'left'}
            , {field: 'type', width: 80, title: '类型', align: 'center', templet(d) {
                    if (d.type == 1) {
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">角色</span>';
                    } else if (d.type == 2) {
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">道具</span>';
                    }
                }
            }
            // , { field: 'icon', width: 80, title: '图标', align: 'center', templet: '<p><i class="layui-icon {{d.icon}}"></i></p>'}
            , {fixed: 'right', width: 250, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("道具", 500, 350);

    } else if (A == "edit") {
        // /**
        //  * 提交表单
        //  */
        // form.on('submit(submitForm2)', function (data) {
        //     // 提交表单
        //     common.submitForm(data.field, null, function (res, success) {
        //         console.log("保存成功回调");
        //     });
        //     return false;
        // });
    }
});
