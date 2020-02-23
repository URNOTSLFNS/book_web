<!DOCTYPE html>
<html lang="en">
<head>

    <meta charset="UTF-8">
    <title>小说网</title>
    <link rel="stylesheet" href="static/css/css.css">
    <script src="static/js/jquery-3.4.1.js"></script>
    <script src="static/js/js.js"></script>
    <script>
        var index=0;
        function Change() {
            index++;
            var a=document.getElementsByClassName("picture");
            if(index==a.length) index=0;
            for(var i=0;i<a.length;i++){
                a[i].style.display='none';
            }
            a[index].style.display='block';
        }
        setInterval(Change,2000);
    </script>
</head>
<body>
<div class="header">
    <div class="bt">
        小说网
    </div>
   <div class="search" method="post">
       <form action="/search">
           <input type="text" name="bookId" class="seText"><input type="submit" class="stn" value="搜索">
       </form>
   </div>
<!--<button class="clear">退出登录</button>-->
</div>
<div class="content1">
    <img src="static/img/img1.jpg" alt="1" class="picture p1">
    <img src="static/img/img2.jpg" alt="2" class="picture p2">
    <img src="static/img/img3.jpg" alt="3" class="picture p3">
</div>
<div class="content2">
    <div class="phb">
        <div class="line">
            排行榜
        </div>
        <div class="pai">
    
        </div>
        <div class="page">
            <a href="javascript:;">2</a>
            <a href="javascript:;" class="cur">1</a>
        </div>
    </div>
    <div class="dl">
        <ul>
           <form  class="login_form"  name = "login" action="/login" method="post">
        <h1 class="login_title">用户登录</h1>
      <li>
      账号  <input type="text"  class="input_txt" name = "userName" value="{{.userName}}">
      </li>
      <li>
      密码 <input type="password" name = "passWord"  class="input_txt">
      </li>

     <li>
                账号类型：<lable><input type="radio" value="1"  name="select" checked="checked">用户
                <input type="radio" value="2"  name="select">管理员
                <input type="radio" value="3"  name="select">作家</lable>
      </li>
      <li>
       <input type="submit" class="de" value="登录">
      </li>
       </ul>
        <span>{{.errmsg}}</span>
    </form>
 

        <div class="zc">
            <h2>
                <a href="http://localhost:8080/register">注册</a>
            </h2>
        </div>
    </div>
    <div class="shu">
       <a href="http://localhost:8080/shelf">我的书架</a>
  </div>
    <div class="sort">
        <ul>
           <form action="http://localhost:8080/classify" method="post">
            <li class="left b1"><input type="submit" name="book_type" value="玄幻"> </li>
            <li class="right b1"><input type="submit" name="book_type" value="历史"></li>
            <li class="left b2"><input type="submit" name="book_type" value="仙侠"></li>
            <li class="right b2"><input type="submit" name="book_type" value="都市"></li>
            <li class="left b3"><input type="submit" name="book_type" value="体育"></li>
            <li class="right b3"><input type="submit" name="book_type" value="灵异"></li>
            </form>
        </ul>
    </div>
    <div class="tu">
        <a href="javascript: ;"><img src="static/img/jubao.png" alt=""></a>

    </div>
    <div class="none">
        <div class="face_image"><a href="javascript: ;"><img src="static/img/90.jpg" alt=""></a></div>
        <div class="face_info">
            <h3>第一序列</h3>
            <span>作者：<a href="javascript: ;">会说话的肘子</a></span>
            <p>有人说，当灾难降临时，精神意志才是人类面对危险的第一序列武器。</p>
            </div>
    </div>
    <div class="snr">
<!--        <ul><li><div class=list><div class='list_name'><span>霸道总裁爱上我</span></div>-->
<!--            <div class='list_del'><i><img src='static/img/shanchu.jpg'alt=''></i></div>-->
<!--            <div class='list_start'>-->
<!--                <form action="http://localhost:8080/book/intro" method="get"><button>查看详情</button></form></div>-->
<!--        </div></li>-->
        </ul>
    </div>
</div>
<div class="footer">
    <a href="javascript: ;">|关于网站|</a>
    <a href="javascript: ;">|法律声明|</a>
    <p>作者发布小说作品时，请遵守国家互联网信息管理办法规定。
        本站所收录小说作品、社区话题、书库评论均属其个人行为，不代表本站立场。</p>
</div>
</body>
</html>