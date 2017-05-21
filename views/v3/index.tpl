<!DOCTYPE html>
<html>
<head>
	{{template "top"}}
	<title>{{.webName}} - {{.webKeywords}}</title>
	<meta name="keywords" content="{{.webKeywords}}">
	<meta name="description" content="{{.webDesc}}">
	<link rel="stylesheet" href="//cdn.bootcss.com/bxslider/4.2.12/jquery.bxslider.min.css" type="text/css" media="all">
	<link rel="stylesheet" href="/css/backtotop.css" type="text/css" media="all">
</head>
<body class="home blog body_top" youdao="bind">
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
                <ul class="h-screen"><li class="current-menu-item"><a href="{{func_cateurl .Cid}}"> {{.Name}} </a></li></ul>
            </div>
            <ul class="h-soup cl">
                <li class="open"><i class="fa fa-coffee"></i><a href="{{func_cateurl .Cid}}" title="{{.Name}}">  查看更多 </a></li>
            </ul>
        </div>
        <div class="update_area">
            <div class="update_area_content">
                <ul class="update_area_lists cl">
                    {{range (func_articles .Cid 0 5)}}
                    {{template "list" .}}
                    {{end}}
                </ul>
            </div>
        </div>
	{{end}}
	{{template "footer"}}
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
