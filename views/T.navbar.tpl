{{define "navbar"}}
<style>
  body {
    background: url("/static/img/44.jpg") no-repeat;
    background-size: cover;
  }

  .navbar {
    background-color: rgb(206, 241, 247);
  }

  .navbar-nav>.active>a,
  .navbar-nav>.active>a:hover,
  .navbar-nav>.active>a:focus {
    background-color: #aee6ef
  }
</style>
<div class="navbar navbar-default navbar-static-top">
  <div class="container">
    <div>
      <a class="navbar-brand" href="/">我的博客</a>
      <ul class="nav navbar-nav">
        <li {{if .IsHome}} class="active" {{end}}><a href="/">首页</a></li>
        <li {{if .IsCategory}} class="active" {{end}}><a href="/category">分类</a></li>
        <li {{if .IsTopic}} class="active" {{end}}><a href="/topic">文章</a></li>
      </ul>
    </div>

    <div class="pull-right">
      <ul class="nav navbar-nav">
        {{if .IsLogin}}
        <li><a href="/login?exit=true">退出</a></li>
        {{else}}
        <li><a href="/login">登录</a></li>
        {{end}}
      </ul>
    </div>

  </div>
</div>
{{end}}