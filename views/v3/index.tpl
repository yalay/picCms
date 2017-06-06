<!DOCTYPE html>
<html lang="{{.lang}}">
<head>
    {{template "top"}}
    <title>{{.webName}} - {{.webKeywords}}</title>
    <meta name="keywords" content="{{.webKeywords}}">
    <meta name="description" content="{{.webDesc}}">
    <link rel="stylesheet" href="//cdn.bootcss.com/bxslider/4.2.12/jquery.bxslider.min.css" type="text/css" media="all">
    <link rel="stylesheet" href="/css/backtotop.css" type="text/css" media="all">
</head>
<body class="home blog body_top">
    {{template "header" .}}
    <!--效果html开始-->
    <div class="site-wrap hide">
        <ul class="bxslider">
            {{range (func_articles 0 5 6)}}
            <li><a target="_blank" href="{{func_articleurl .Id}}"><img src="{{.Cover}}?s=590x394" title="{{.Title}}"></a></li>
            {{end}}
        </ul>
    </div>
    {{range func_cates}}
        <div class="home-filter">
            <div class="h-screen-wrap">
                <ul class="h-screen"><li class="current-menu-item"><a href="{{func_cateurl .EngName}}"> {{func_lang .Name $.lang}} </a></li></ul>
            </div>
            <ul class="h-soup cl">
                <li class="open"><i class="fa fa-coffee"></i><a href="{{func_cateurl .EngName}}" title="{{func_lang .Name $.lang}}">  {{func_lang "查看更多" $.lang}} </a></li>
            </ul>
        </div>
        <div class="update_area">
            <div class="update_area_content">
                <ul class="update_area_lists cl">
                    {{range (func_articles .Cid 0 5)}}
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
            </div>
        </div>
    {{end}}
    {{template "footer" .}}
    <script type="text/javascript" src="//cdn.bootcss.com/bxslider/4.2.12/jquery.bxslider.min.js"></script>
    <script type="text/javascript">
        $(document).ready(function(){
            $('.site-wrap').removeClass('hide');
            $('.bxslider').bxSlider({
                moveSlides: 1,
                slideMargin: 5,
                infiniteLoop: true,
                slideWidth: 590,
                minSlides: 1,
                maxSlides: 6,
                pager: false,
                controls: true,
                auto: true,
            });
        });
    </script>
</body>
</html>
