let Toast;
let pageNum = 1;
const pageSize = 8;
//弹窗全局设置
Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
});
$(function () {
    //分页
    page()
    //查询
    $("#form-search").bind("click", function () {
        page()
    })
    //新增
    $("#form-save").on("submit", function (ev) {
        ev.preventDefault();
        $.ajax("/api/record/record", {
            type: "POST",
            dataType: 'json',
            data: $('#form-save').serialize()
        }).done(function (e) {
            if (e.success) {
                alter2(1, "成功")
                $("#btn-to-save").click()
                $('#form-save')[0].reset()
                page()
            } else {
                alter2(3, e.msg)
            }
        }).fail(function (err) {

        })
    })
//    修改
    $("#form-update").on("submit", function (ev) {
        ev.preventDefault();
        $.ajax("/api/record/record/" + $("#uuid2").val(), {
            type: "PUT",
            dataType: 'json',
            data: $('#form-update').serialize()
        }).done(function (e) {
            if (e.success) {
                alter2(1, "成功")
                $("#btn-to-update").click()
                $('#form-update')[0].reset()
                page()
            } else {
                alter2(3, e.msg)
            }
        }).fail(function (err) {

        })
    })
})

//分页标签
function pageLabel(o) {
    pageNum = o
    page()
}

//分页查询
function page() {
    $("#table-content").find("tr").remove()
    $("#table-page").find("li").remove()
    $.ajax("/api/record/records", {
        type: "GET",
        dataType: 'json',
        data: {
            "pageNum": pageNum,
            "pageSize": pageSize,
            "dates": $("#table_search").val()
        }
    }).done(function (e) {
        if (e.success) {
            $(e.data).each(function (i, o) {
                $("#table-content").append(trs(i, o))
            })
            //    分页
            $(e.page.pageData).each(function (i, o) {
                let a = (o === pageNum ? 'active' : '')
                let l = '<li class="page-item ' + a + '" onclick="pageLabel(\'' + o + '\')"><a class="page-link" href="#">' + o + '</a></li>'
                $("#table-page").append(l)
            })
        } else {

        }
    }).fail(function (err) {

    })
}

//tr模板
function trs(i, e) {
    return '<tr>'
        + '<td>' + (i + 1) + '</td>'
        + '<td>' + (e.dates) + '</td>'
        + '<td>' + (e.money) + '</td>'
        + '<td>' + (e.remarks) + '</td>'
        + '<td>'
        + '<button type="button" class="btn btn-danger btn-xs" onclick="del(\'' + e.uuid + '\')">删除</button>'
        + '&nbsp;&nbsp;'
        + '<button type="button" class="btn btn-warning btn-xs" onclick="one(\'' + e.uuid + '\')"' +
        ' data-toggle="modal" data-target="#modal-update">修改/查看</button>'
        + '</td>'
        + '</tr>'
}

//删除
function del(o) {
    alter2IsOk("是否确定删除？").then(function (e) {
        if (e.value) {
            $.ajax("/api/record/record/" + o, {
                type: "DELETE",
                dataType: 'json'
            }).done(function (e) {
                if (e.success) {
                    page()
                } else {
                    alter2(4, e.msg)
                }
            }).fail(function (err) {

            })
        }
    })
}

//根据id获取
function one(e) {
    $.ajax("/api/record/record/" + e, {
        type: "GET",
        dataType: 'json'
    }).done(function (e) {
        if (e.success) {
            $("#uuid2").val(e.data.uuid)
            $("#dates2").val(e.data.dates)
            $("#money2").val(e.data.money)
            $("#remarks2").val(e.data.remarks)
        } else {
            alter2(3, e.msg)
        }
    }).fail(function (err) {

    })
}

//弹窗提示
function alter2(icon, title) {
    switch (icon) {
        case 1:
            icon = 'success'
            break
        case 2:
            icon = 'info'
            break
        case 3:
            icon = 'warning'
            break
        case 4:
            icon = 'error'
            break
    }
    Toast.fire({
        icon: icon,
        title: title,
        showConfirmButton: false,
        timer: 2000
    })
}

//弹窗提示
function alter2IsOk(butText) {
    return Swal.fire({
        icon: 'warning',
        showConfirmButton: true,
        confirmButtonColor: '#3085d6',
        confirmButtonText: butText,
        showCancelButton: true,
        cancelButtonColor: '#d33',
        cancelButtonText: "取消"
    })
}