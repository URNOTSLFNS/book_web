<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>我的书架</title>
    <script src="static/js/jquery-3.4.1.js"></script>
    <script src="static/js/js.js"></script>
    <style>
        .box1{
            height: 80px;
            width: 30%;
            padding-left: 50px;
        }
        .snr{
            height: 700px;
            width: 387px;
            background-color: white;
            margin-left: 313px;
            margin-top: 50px;
            position: absolute;
            border: 1px solid #999999;
        }
        .snr .list{
            height: 60px;
            width: 100%;
        }
        .list_name{
            display: inline-block;
        }
        .list_name span{
            line-height: 60px;
            margin-left: 30px;
            user-select: none;
        }
        .list_start{
            float: right;
            margin-top: 21px;
        }
        .list_del{
            width: 100px;
            height: 100%;
            float: right;
            text-align: center;
            line-height: 45px;
        }

        .list_start button{
            background-color: #FFFFFF;
            outline: none;
            border: none;
            opacity: 0.5;
        }
        .list_start button:hover{
            opacity: 1;
        }
        .list_del i{
            display: inline-block;
            width: 18px;
            height: 18px;
            margin-right: 10px;
            margin-top: 12px;
            /*vertical-align: -3px;*/
        }
        .list_del i img{
            width: 18px;
            height: 18px;
            opacity: 0.5;
        }
        .list_del i img:hover{
            opacity: 1;
        }
        .snr>ul>li{
            border-bottom: 1px solid #CCCCCC;
            display: inline-block;
        }
    </style>
</head>
<body>
<div class="box1">
    <h1>我的书架</h1>
</div>
<div class="snr">
<ul>
    <li>
        <div class=list>
            <div class='list_name'><span>霸道总裁爱上我</span></div>
            <div class='list_del'><i><img src='static/img/shanchu.jpg' alt=''></i></div>
            <div class='list_start'>
                <form action="http://localhost:8080/shelf" method="get"><button>查看详情</button></form>
            </div>
        </div>
    </li>
    <li>
        <div class=list>
            <div class='list_name'><span>霸道总裁爱上我</span></div>
            <div class='list_del'><i><img src='static/img/shanchu.jpg' alt=''></i></div>
            <div class='list_start'>
                <form action="http://localhost:8080/shelf" method="get"><button>查看详情</button></form>
            </div>
        </div>
    </li>  
       <span>{{.errmsg}}</span>
</ul>
</div>
</body>
</html>