<!DOCTYPE html>
<html>
<head>
    <title>HTTP-proxy - {{ .SiteDomain }}</title>
    <link rel="stylesheet" href="/static/css/application.css"/>
    <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css"/>
    <link rel="stylesheet" href="/static/css/bootstrap.min.css"/>
    <link rel="stylesheet" href="/static/css/sites.css"/>
</head>
<body role="document">


<!-- Fixed navbar -->
<nav class="navbar navbar-inverse navbar-fixed-top">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#"> 统一线上HTTP代理管理</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
                <li class="active"><a href="/sites">Home</a></li>
                <li><a href="/about">About</a></li>
            </ul>
        </div><!--/.nav-collapse -->
    </div>
</nav>


<div class="container theme-showcase" role="main">

{{.LayoutContent}}

</div>
</body>
</html>
