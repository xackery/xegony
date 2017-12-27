{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">



<div class="row" >
<div class="col-lg-12">

    <div class="hpanel">
        <div class="panel-body">
            <div class="text-muted small pull-right text-right">            
            </div>
            <div class="form-inline">
                Search: <input class="form-control" type="text">
            </div>

        </div>

    </div>


    <div class="hpanel forum-box">

    <div class="panel-heading">
        <span class="f">Forum</span>
    </div>

    {{range $key, $value := .Forums}}
    <div class="panel-body">
        <div class="row">

            <div class="col-md-10 forum-heading">
                <a href="/forum/{{$value.Id}}"><h4> {{$value.Name}}</h4></a>
                <a href="/forum/{{$value.Id}}"><div class="desc">{{$value.Description}}</div></a>
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
