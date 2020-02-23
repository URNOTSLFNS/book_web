<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>注册</title>
    <link rel="stylesheet" href="static/css/recss.css">
    <script src="static/js/jquery-3.4.1.js"></script>
    <script src="static/js/js.js"></script>
</head>
<body>
<div class="box1">
    <h1>用户注册</h1>
</div>
     <div class="box2">
         <div class="box3">
         <form action="/register" method="post">
              用户名：<input type="text" name="userName" class="user">
             <br><br><br>
             密   码： <input type="password" name="passWord" class="re_password">
             <input type="submit" value="注册" class="btn" name="register">
         </form>
         </div>
     </div>
</body>
</html>