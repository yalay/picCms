{{define "list.zh"}}
<li class="i_list list_n2">
    <a target="_blank" href="{{func_articleurl .Id}}" title="{{.Title}}">
        <img class="waitpic" src="/img/loading.gif" data-original="{{.Cover}}?s=270x370" width="270" height="370" alt="{{.Remark}}" style="display: inline;">
    </a>
    <div class="case_info">
        <div class="meta-title"> {{.Title}} </div>
        <div class="meta-post"><i class="fa fa-clock-o"></i> {{func_time2 .Addtime}} <span class="cx_like"><i class="fa fa-heart"></i> {{.Up}} </span></div>
    </div>
    <div class="meta_zan xl_1"><i class="fa fa-eye"></i> {{.Hits}} </div>
</li>
{{end}}
