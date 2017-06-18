{{define "header"}}
<!-- 头部代码begin -->
<div class="index_header nav_headertop">
    <div class="header_inner">
        <div class="logo"><a href="{{.webUrl}}"><img src="/img/logo.png" alt="{{.webName}}"></a></div>
        <div class="header_menu">
            <ul>
                <li {{if eq $.cid 0}} class="current-menu-item"{{end}}><a href="/">{{func_lang "网站首页" .lang}}</a></li>
                {{range func_cates}}
                <li {{if eq $.cid .Cid}} class="current-menu-item"{{end}}><a href="{{func_cateurl .EngName}}">{{func_lang .Name $.lang}}</a></li>
                {{end}}
                <li class="megamenu toke {{if eq $.cid 90}} current-menu-ancestor{{end}}"><a>{{func_lang "精品专题" .lang}}</a>
                    <ul class="sub-menu">
                        <li><a>{{func_lang "人气模特" .lang}}</a>
                            <ul class="sub-menu">
                                {{range func_topics}}
                                <li><a href="{{func_topicurl .EngName}}">{{func_lang .Name $.lang}}</a></li>
                                {{end}}
                            </ul>
                        </li>
                        <li><a>{{func_lang "热门标签" .lang}}</a>
                            <ul class="sub-menu">
                                {{range (func_hottags 10)}}
                                <li><a href="{{func_tagurl .}}">{{func_lang . $.lang}}</a></li>
                                {{end}}
                            </ul>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
        <div class="login_text hide">
            <ul id="mobile_menu">
            <li {{if eq $.cid 0}} class="current-menu-item"{{end}}><a href="{{.webUrl}}">{{func_lang "网站首页" .lang}}</a></li>
            {{range func_cates}}
            <li {{if eq $.cid .Cid}} class="current-menu-item"{{end}}><a href="{{func_cateurl .EngName}}">{{func_lang .Name $.lang}}</a></li>
            {{end}}
        </div>
    </div>
</div>
<!-- 头部代码end -->
{{end}}
