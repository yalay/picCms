<!DOCTYPE html>
<html lang="{{.lang}}">
<head>
    {{template "top"}}
    <title>{{.tTitle}}</title>
    <meta name="keywords" content="{{.tag}}">
    <meta name="description" content="{{.tTitle}}">
</head>
<body class="home blog body_top">
    {{template "header" .}}
    <div class="cat_bg">
        <div class="cat_bg_img" style="background-image:url(/img/tags.png);">
            <div><span style="font-size: 18px;color: #F14141;font-weight: 600;">{{.tag}}</span><br>{{.tTitle}}</div>
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
                <li class="i_list list_n2">
                    <a target="_blank" href="{{func_articleurl .Id}}" title='{{func_lang .Title $.lang}}'>
                        <img class="waitpic" src="/img/loading.gif" data-original="{{.Cover}}?s=270x370" width="270" height="370" alt='{{func_lang .Remark $.lang}}' style="display: inline;">
                    </a>
                    <div class="case_info">
                        <div class="meta-title"> {{func_lang .Title $.lang}} </div>
                        <div class="meta-post"><i class="fa fa-clock-o"></i> {{func_time2 .Addtime}} <span class="cx_like"><i class="fa fa-heart"></i> {{.Up}} </span></div>
                    </div>
                    <div class="meta_zan xl_1"><i class="fa fa-eye"></i> {{.Hits}} </div>
                </li>
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
