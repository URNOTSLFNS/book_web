<body>

    <form  class="login_form"  name = "admin_delete" action="/admin/delete" method="post">
        <h1 class="login_title">管理员删除</h1>
        删除书本的ID <input type="text"  class="input_txt" name = "bookId" value="{{.bookId}}">
        <input type="submit" value="删 除" class="input_sub">
        <span>{{.errmsg}}</span>
    </form>

</body>
