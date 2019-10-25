<!DOCTYPE html>

<html>
<head>
  <title>Beego | Duckula</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="_xsrf" content="{{.xsrf_token}}" />

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    .logo {
      width: 120px;
      height: 120px;
      margin: 40px auto;
      border-radius: 10px;
      background-image: url('https://blog.duckula.net:88/static/img/logo.png');
      background-color: #000;
      background-repeat: no-repeat;
      background-size: 80%;
      background-position: center center;
    }

    h1 {
      text-align: center;
      font-size: 42px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
      margin: 40px;
    }

    header {
      padding: 80px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
      margin: 20px auto;
    }

    a {
      color: #069;
      text-decoration: none;
      margin: 0 10px;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
  </style>
</head>

<body>
  <header>
    <h1 class="logo"></h1>
    <h1>Duckula + Beego</h1>
    <!--
    <div class="description">
      Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
    </div>
    -->
    <div class="description">
      <a href="{{urlfor "UserController.ListUser"}}">{{urlfor "UserController.ListUser"}}</a>
      <a href="/api/user/123">/api/user/123</a>
      <a href="/api/get">/api/get</a>
      <a href="/api/any?name=Tiger">/api/any?name=Tiger</a>
      <a href="/download/test.txt">/download/test.txt</a>
    </div>
    <div class="description">
      DB:{{.DB}}
      XSRF TOKEN:{{.xsrf_token}}<br>
      <!--<code>XSRF DATA:{{.xsrf_data}}</code><br>-->
      <form action="/api/post" method="POST">
        <input type="text" name="username" value="Tiger" />
        <button type="submit">Submit</button>
      </form>
      <form action="/api/user/create" method="POST">
        {{.xsrf_data}}
        <input type="text" name="username" value="Tiger" />
        <button type="submit">Submit</button>
      </form>
    </div>
  </header>
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
  <script>
    console.log(document.querySelector('meta[name=_xsrf]'));
  </script>
</body>
</html>
