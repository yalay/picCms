<!DOCTYPE html>
<html>
<head>
	{{template "top"}}
	<title>{{.title}} - {{.keywords}}</title>
	<meta name="keywords" content="{{.keywords}}">
</head>
<body class="home blog body_top" youdao="bind">
    {{template "header" .}}
	<div class="main">
		<div class="main_inner">
			<div class="main_left" style="width:100%">
				<div class="item_title">
					<h1> {{.title}}(缩略图<span>{{.pageId}} / {{.attachNum}}</span>)</h1>
					<div class="single-cat"><span>分类:</span> <a href="{{.cUrl}}" rel="category tag">{{.cName}}</a> / <span>发布于</span>{{.pubDate}}</div>
				</div>
				<div class="item_info cl">
					<div style="float:left;">
						<i class="fa fa-eye"></i> <span class="cx-views">{{.hits}}</span> 人气 /
						<i class="fa fa-comment"></i> <span><a href="#respond">参与</a></span> 评论
					</div>
					<div class="post_au">
						<a style="margin-right:15px;color: #2CCBE6;" class="ajax ajax_dl_attachs" href="#"><i class="fa fa-download" style="margin-right:3px;"></i>免费下载高清原图</a>
					</div>
				</div>
				<div class="content" id="content">
					<div class="content_left">
						<a href="{{.preUrl}}" title="上一页" class="pre-cat"><i class="fa fa-chevron-left"></i></a>
						<a href="{{.nextUrl}}" title="下一页" class="next-cat"><i class="fa fa-chevron-right"></i></a>
						<div class="image_div" id="image_div">
							<p><a href="{{.nextUrl}}"><img src="{{.file}}" alt="{{.title}}" title="点击图片查看下一张"></a></p>
							<div class="nav-links page_imges">{{str2html .pagination}}</div>
						</div>

						<div class="tag cl" style="margin-top:30px;">
							<span class="dtpost-like cl">
								<!-- 点赞功能
								<a href="javascript:;" data-action="ding" data-id="{$article.id}" class="favorite">
									<i class="fa fa-thumbs-up"></i>
									<span class="count"><em class="ct_ding" style="color: #F58282;">{$article.up}</em>个赞</span>
								</a>
								-->
								<a class="share-btn" href="javascript:;" onclick="javascript:userAddFavorite()" title="收藏">
									<i class="fa fa-star"></i>
									<span class="count">收藏</span>
								</a>
								<a class="share-down ajax_dl_attachs" href="#"><i class="fa fa-download"></i><span class="count">下载原图</span></a>
							</span>
						</div>
					</div>
				</div>
				<div class="content_right_title">相关资源：
					<span class="single-tags">
					<!-- 待补充 -->
					</span>
				</div>
				<ul class="xg_content">
					<li class="i_list_frame list_n2">
					    <script type="text/javascript" data-idzone="2581393" src="https://ads.exoclick.com/nativeads.js"></script>
						<div class="meta_zan xl_1"><i class="fa fa-eye"></i> 99+ </div>
					</li>
					<!-- 待补充 -->
				</ul>
				<section class="single-post-comment">
                    <div class="single-post-comment-reply" id="respond" >
                        <!-- UY BEGIN -->
                        <div id="uyan_frame"></div>
                        <script type="text/javascript" src="http://v2.uyan.cc/code/uyan.js?uid=2099909"></script>
                        <!-- UY END -->
                    </div>
                </section>
			</div>
		</div>
	</div>
	{{template "footer"}}
	<script type="text/javascript">
		var articleId = "{{.id}}";
        var web_script = "/";
	</script>
	<script type="text/javascript" src="/js/ajax.js"></script>
    <!-- JiaThis Button BEGIN -->
    <div class="jiathis_share_slide jiathis_share_32x32" id="jiathis_share_slide">
    <div class="jiathis_share_slide_top" id="jiathis_share_title"></div>
    <div class="jiathis_share_slide_inner">
    <div class="jiathis_style_32x32">
    <a class="jiathis_button_qzone"></a>
    <a class="jiathis_button_tsina"></a>
    <a class="jiathis_button_tqq"></a>
    <a class="jiathis_button_weixin"></a>
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