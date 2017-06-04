{{define "footer"}}
<div class="foot" id="footer">
    <div class="foot_list">
        <div class="foot_num"><div>{{func_lang "文章总数" .lang}}</div> <div>170+</div></div>
        <div class="foot_num"><div>{{func_lang "评论总数" .lang}}</div> <div>23+</div></div>
        <div class="foot_num"><div>{{func_lang "专题栏目" .lang}}</div> <div>17+</div></div>
        <div class="foot_num"><div>{{func_lang "运营天数" .lang}}</div> <div>137+</div></div>
    </div>
</div>
<!--footer-->
<footer class="w100 cl">
    <div class="w1080 fot cl">
        <p class="footer_menus">
            <a href="mailto:{{.email}}">{{func_lang "广告合作" .lang}}</a>
            <a href="{{.mapUrl}}">{{func_lang "网站地图" .lang}}</a>
            <a href="{{.rssurl}}">{{func_lang "RSS订阅" .lang}}</a>
        </p>
        <p>{{if .copyright}}{{str2html .copyright}}{{end}}{{if .tongji}}{{str2html .tongji}}{{end}}</p>
    </div>
    <div class="cbbfixed" style="bottom: -90px;"><a class="gotop cbbtn"><i class="fa fa-angle-up"></i></a></div>
</footer>
<script type="text/javascript" src="//cdn.bootcss.com/jquery_lazyload/1.9.7/jquery.lazyload.min.js"></script>
<script type="text/javascript" src="//cdn.bootcss.com/SlickNav/1.0.10/jquery.slicknav.min.js"></script>
<script type="text/javascript" src="/js/main.js"></script>
{{end}}
