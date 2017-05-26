<!DOCTYPE html>
<html>
<head>
    {{template "top"}}
    <title>关于{{.tag}}的图片 - {{.webName}}</title>
    <meta name="keywords" content="{{.tag}}">
    <meta name="description" content="关于{{.tag}}的图片">
</head>
<body class="home blog body_top" youdao="bind">
    {{template "header" .}}
    <div class="cat_bg">
        <div class="cat_bg_img" style="background-image:url(/img/tags.png);">
            <div><span style="font-size: 18px;color: #F14141;font-weight: 600;">{{.tag}}</span><br>关于{{.tag}}的图片</div>
        </div>
    </div>
    <!--分类导航-->
    <div class="fl flbg">
        <div class="fl_title"><div class="fl01">{{.tag}}</div></div>
    </div>
    <div class="update_area">
        <div class="update_area_content">
            <ul class="update_area_lists cl">
                {{range .tArticles}}
                {{template "list" .}}
                {{end}}
            </ul>
            <nav class="navigation pagination" role="navigation">
                <h2 class="screen-reader-text">文章导航</h2>
                <div class="nav-links">{{if .pagination}}{{str2html .pagination}}{{end}}</div>
            </nav>
        </div>
    </div>
    {{template "footer" .}}
</body>
</html>
