// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 用户管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func', 'laydate', 'form'], function () {

    //声明变量
    var func = layui.func
        , $ = layui.$
        , u = layui.common
        , form = layui.form
        , laydate = layui.laydate;

    // 自己做监听提交
    form.on('submit(submitForm)', function (data) {
        if ($("input:checkbox[name='roleIds']:checked").length == 0) {
            // return;
        }
        //获取checkbox[name='roleIds']的值，获取所有选中的复选框，并将其值放入数组中
        var arr = new Array();
        $("input:checkbox[name='roleIds']:checked").each(function(i){
            arr[i] = $(this).val();
        });
        //  替换 data.field.level的数据为拼接后的字符串
        data.field.roleIds = arr.join(",");//将数组合并成字符串
        console.log("表单数据", data.field)

        return u.submitForm(data.field, null, (function (t, e) {
            console.log("保存成功回调")
        })), !1;
    });

    //执行一个laydate实例
    laydate.render({
        elem: '#birthday' //指定元素
    });

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'realname', width: 120, title: '用户姓名', align: 'center'}
            // , {field: 'gender', width: 60, title: '性别', align: 'center', templet(d) {
            //         var cls = "";
            //         if (d.gender == 1) {
            //             // 男
            //             cls = "layui-btn-normal";
            //         } else if (d.gender == 2) {
            //             // 女
            //             cls = "layui-btn-danger";
            //         } else if (d.gender == 3) {
            //             // 保密
            //             cls = "layui-btn-warm";
            //         }
            //         return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.genderName + '</span>';
            //     }
            // }
            , {field: 'username', width: 100, title: '登录名', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet: function (d) {
                    return '<input type="checkbox" name="status" value="' + d.id + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="status" ' + (d.status == 1 ? 'checked' : '') + '>';
                }
            }
            , {field: 'deptName', width: 150, title: '所属部门', align: 'center'}
            , {field: 'levelName', width: 120, title: '职级名称', align: 'center'}
            , {field: 'positionName', width: 120, title: '岗位名称', align: 'center'}
            , {field: 'sort', width: 100, title: '排序号', align: 'center'}
            , {field: 'create_time', width: 180, title: '添加时间', align: 'center', templet: function (d) {
                    return "<div>" + layui.util.toDateString(d.create_time, 'yyyy-MM-dd HH:mm:ss') + "</div>"
                }
            }
            , {field: 'update_time', width: 180, title: '更新时间', align: 'center', templet: function (d) {
                    return "<div>" + layui.util.toDateString(d.update_time, 'yyyy-MM-dd HH:mm:ss') + "</div>"
                }
            }
            , {fixed: 'right', width: 250, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList", function (layEvent, data) {
            if (layEvent === 'resetPwd') {
                layer.confirm('您确定要初始化当前用户的密码吗？', {
                    icon: 3,
                    skin: 'layer-ext-moon',
                    btn: ['确认', '取消'] //按钮
                }, function (index) {
                    //初始化密码
                    var url = cUrl + "/resetPwd";
                    // 切记采用FormData表单提交
                    var formData = new FormData();
                    formData.append("id", data.id);
                    func.ajaxPost(url, formData, function (data, success) {
                        console.log("重置密码：" + (success ? "成功" : "失败"));
                        // 关闭弹窗
                        layer.close(index);
                    })
                });
            }
        });

        //【设置弹框】
        func.setWin("用户");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});