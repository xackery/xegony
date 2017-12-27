{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">



<div class="row" >
<div class="col-lg-12">

    <div class="hpanel forum-box">

    <div class="panel-heading">
        <span class="f">Forum {{.Forum.Name}} > <span class="text-success"> Topics</span></span>
    </div>

    {{range $key, $value := .Topics}}
    <div class="panel-body">
        <div class="row">

            <div class="col-md-10 forum-heading">
                <a href="/topic/{{$value.Id}}"><h4> {{$value.Title}}</h4></a>
                <a href="/topic/{{$value.Id}}"><div class="desc">{{$value.Body}}</div></a>
            </div>
            <div class="col-md-1 forum-info">
                <span class="number"> 4780 </span>
                <small>Views</small>
            </div>
            <div class="col-md-1 forum-info">
                <span class="number"> 150 </span>
                <small>Posts</small>
            </div>
        </div>
    </div>
    {{end}}

    </div>
</div>

</div>
</div>
</div>
