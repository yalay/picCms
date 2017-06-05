{{define "list.en"}}
<li class="i_list list_n2">
    <a target="_blank" href="{{func_articleurl .Id}}" title='{{func_lang .Title "en"}}'>
        <img class="waitpic" src="/img/loading.gif" data-original="{{.Cover}}?s=270x370" width="270" height="370" alt='{{func_lang .Remark "en"}}' style="display: inline;">
    </a>
    <div class="case_info">
        <div class="meta-title"> {{func_lang .Title "en"}} </div>
        <div class="meta-post"><i class="fa fa-clock-o"></i> {{func_time2 .Addtime}} <span class="cx_like"><i class="fa fa-heart"></i> {{.Up}} </span></div>
    </div>
    <div class="meta_zan xl_1"><i class="fa fa-eye"></i> {{.Hits}} </div>
</li>
{{end}}
