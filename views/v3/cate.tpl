<!DOCTYPE html>
<html lang="{{.lang}}">
<head>
    {{template "top"}}
    <title>{{.cName}} - {{.cKeywords}}</title>
    <meta name="keywords" content="{{.cKeywords}}">
    <meta name="description" content="{{.cDesc}}">
</head>
<body class="home blog body_top">
    {{template "header" .}}
    <div class="cat_bg">
        <div class="cat_bg_img" style="background-image:url(/img/cate_{{.cid}}.png);">
            <div><span style="font-size: 18px;color: #F14141;font-weight: 600;">{{.cName}}</span><br>{{.cDesc}}</div>
        </div>
    </div>
    <!--分类导航-->
    <div class="fl flbg">
        <div class="fl_title"><div class="fl01">{{.cName}}</div></div>
    </div>
    <div class="update_area">
        <div class="update_area_content">
            <ul class="update_area_lists cl">
                {{str2html (func_adsense "list-native")}}
                {{range .cArticles}}
                    {{if eq $.lang "zh"}}
                    {{template "list.zh" .}}
                    {{else}}
                    {{template "list.en" .}}
                    {{end}}
                {{end}}
            </ul>
            <nav class="navigation pagination" role="navigation">
                <h2 class="screen-reader-text">{{func_lang "文章导航" .lang}}</h2>
                <div class="nav-links">{{if .pagination}}{{str2html .pagination}}{{end}}</div>
            </nav>
        </div>
    </div>
    {{template "footer" .}}
</body>
</html>
