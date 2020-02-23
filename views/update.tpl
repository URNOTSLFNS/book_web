<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>修改信息</title>
    <link rel="stylesheet" href="static/css/updata.css">
    <script src="static/js/jquery-3.4.1.js"></script>
    <script src="static/js/js.js"></script>
</head>
<body>
<div class="box5">
    <div class="boxmiddle">
        <div class="deta">
                     <div class="image"><img src="" alt=""></div>
                             <div class="info">
                                书本的ID <input type="text"  class="bookid" name = "Id" ><br>
                                 书名：<input type="text" class="bookname" name = "update_name"><br>
                                 类型:<input type="text" class="booktype" name = "update_category"><br>
                                 作者：<input type="text" class="writer" name = "update_writer"><br>
                                  <button type="submit" class="finish">确定</button>
                             </div>
                     </div>
    </div>
</div>
    <span>{{.errmsg}}</span>
</body>
</html>