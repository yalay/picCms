{{define "header"}}
<!-- 头部代码begin -->
<div class="index_header nav_headertop">
	<div class="header_inner">
		<div class="logo"><a href="{{.webUrl}}"><img src="/img/logo.png" alt="{{.webName}}"></a></div>
		<div class="header_menu">
			<ul>
				<li {{if eq $.cid 0}} class="current-menu-item"{{end}}><a href="{{.webUrl}}">网站首页</a></li>
				{{range func_cates}}
				<li {{if eq $.cid .Cid}} class="current-menu-item"{{end}}><a href="{{func_cateurl .EngName}}">{{.Name}}</a></li>
				{{end}}
			</ul>
		</div>
		<div class="login_text hide">
			<ul id="mobile_menu">
			<li {{if eq $.cid 0}} class="current-menu-item"{{end}}><a href="{{.webUrl}}">网站首页</a></li>
			{{range func_cates}}
			<li {{if eq $.cid .Cid}} class="current-menu-item"{{end}}><a href="{{func_cateurl .EngName}}">{{.Name}}</a></li>
			{{end}}
		</div>
	</div>
</div>
<!-- 头部代码end -->
{{end}}
