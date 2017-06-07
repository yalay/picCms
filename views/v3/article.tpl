<!DOCTYPE html>
<html lang="{{.lang}}">
<head>
    {{template "top"}}
    <title>{{.title}} - {{.keywords}}</title>
    <meta name="keywords" content="{{.keywords}}">
    <meta name="description" content="{{.title}} - {{.keywords}}">
</head>
<body class="home blog body_top">
    {{template "header" .}}
    <div class="main">
        <div class="main_inner">
            <div class="main_left" style="width:100%">
                <div class="item_title">
                    <h1> {{.title}} (<span>{{.pageId}} / {{.attachNum}}</span>)</h1>
                    <div class="single-cat"><span>{{func_lang "分类" .lang}}: </span> <a href="{{.cUrl}}" rel="category tag">{{.cName}}</a> / <span>{{func_lang "发布于" .lang}} </span>{{.pubDate}}</div>
                </div>
                <div class="item_info cl">
                    <div style="float:left;">
                        <i class="fa fa-eye"></i> <span class="cx-views">{{.hits}}</span> {{func_lang "人气" .lang}} /
                        <i class="fa fa-comment"></i> <span><a href="#respond">{{func_lang "参与" .lang}}</a></span> {{func_lang "评论" .lang}}
                    </div>
                    <div class="post_au">
                        <a style="margin-right:15px;color: #2CCBE6;" class="ajax ajax_dl_attachs" href="#"><i class="fa fa-download" style="margin-right:3px;"></i>{{func_lang "免费下载高清原图" .lang}}</a>
                    </div>
                </div>
                <div class="content" id="content">
                    <div class="content_left">
                        <a href="{{.preUrl}}" title='{{func_lang "上一页" .lang}}' class="pre-cat"><i class="fa fa-chevron-left"></i></a>
                        <a href="{{.nextUrl}}" title='{{func_lang "下一页" .lang}}' class="next-cat"><i class="fa fa-chevron-right"></i></a>
                        <div class="image_div" id="image_div">
                            <p><a href="{{.nextUrl}}"><img src="{{.file}}" alt="{{.title}}" title='{{func_lang "点击图片查看下一张" .lang}}'></a></p>
                            <div class="nav-links page_imges">{{if .pagination}}{{str2html .pagination}}{{end}}</div>
                        </div>

                        <div class="tag cl" style="margin-top:30px;">
                            <span class="dtpost-like cl">
                                <a class="favorite ajax_up" href="javascript:;">
                                    <i class="fa fa-thumbs-up"></i>
                                    <span class="count"><em class="ct_ding" style="color: #F58282;">{{.up}}</em>{{func_lang "个赞" .lang}}</span>
                                </a>
                                <a class="share-btn" href="javascript:;" onclick="javascript:userAddFavorite()" title='{{func_lang "收藏" .lang}}'>
                                    <i class="fa fa-star"></i>
                                    <span class="count">{{func_lang "收藏" .lang}}</span>
                                </a>
                                <a class="share-down ajax_dl_attachs" href="#"><i class="fa fa-download"></i><span class="count">{{func_lang "下载" .lang}}</span></a>
                            </span>
                        </div>
                    </div>
                </div>
                <div class="content_right_title">{{func_lang "相关资源" .lang}}：
                    <span class="single-tags">
                    {{range .tags}}
                    <a href="{{func_tagurl .}}">{{func_lang . $.lang}} </a>
                    {{end}}
                    </span>
                </div>
                <ul class="xg_content">
                    {{str2html (func_adsense "list-native")}}
                    {{range .relates}}
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
                <section class="single-post-comment">
                    <div class="single-post-comment-reply" id="respond" >
                        {{str2html (func_adsense "livere")}}
                    </div>
                </section>
            </div>
        </div>
    </div>
    <script type="text/javascript">
        var articleId = "{{.id}}";
    </script>
    {{template "footer" .}}
    <script type="text/javascript" src="/js/ajax.js"></script>
    <!-- JiaThis Button BEGIN -->
    <div class="jiathis_share_slide jiathis_share_32x32" id="jiathis_share_slide">
    <div class="jiathis_share_slide_top" id="jiathis_share_title"></div>
    <div class="jiathis_share_slide_inner">
    <div class="jiathis_style_32x32">
    <a class="jiathis_button_weixin"></a>
    <a class="jiathis_button_tsina"></a>
    <a class="jiathis_button_tqq"></a>
    <a class="jiathis_button_qzone"></a>
    <a href="http://www.jiathis.com/share" class="jiathis jiathis_txt jtico jtico_jiathis" target="_blank"></a>
    <script type="text/javascript">
    var jiathis_config = {data_track_clickback:'true'
        ,slide:{
            divid:'content',
            pos:'left'
        }
    };
    </script>
    <script type="text/javascript" src="http://v3.jiathis.com/code/jia.js?uid=2099909" charset="utf-8"></script>
    <script type="text/javascript" src="http://v3.jiathis.com/code/jiathis_slide.js" charset="utf-8"></script>
    </div></div></div>
    <!-- JiaThis Button END -->
</body>
</html>
