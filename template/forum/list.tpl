{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">


<div class="row" >
<div class="col-lg-12">

    <div class="hpanel forum-box">

        <div class="panel-heading hbuilt">
            <div class="row">
                <div class="col-md-4">
                Forum
                </div>
                <div class="col-md-6">            
                    <form action="panels.html">
                        <div class="input-group">
                            <input type="text" class="form-control input-xs" placeholder="Search something..." name="s">
                            <span class="input-group-btn">
                                <button class="btn btn-xs btn-default" type="submit">
                                    <span class="fa fa-search"></span>
                                </button>
                            </span>
                        </div>
                    </form>
                    
                </div>
                <div class="col-md-2 pull-right">
                    <a id="add" class="btn btn-default btn-xs">Add</a>
                    <a id="edit" class="btn btn-default btn-xs">Edit</a>
                </div>
            </div>
        </div>

        <div class="panel-body">
            
            {{range $key, $value := .Forums}}
            <div class="row">
                <div class="col-md-10 forum-heading">

                    <div class="dataForum" data-forum-id="{{$value.ID}}">                        
                        <a href="/forum/{{$value.ID}}">                            
                                <h4><i class="{{$value.Icon}}"></i></h4> <h4 id="name" data-type="text" data-pk="{{$value.ID}}" data-url="/api/forum/{{$value.ID}}" data-title="Enter forum name">{{$value.Name}}</h4>                            
                        </a>
                        <a href="/forum/{{$value.ID}}"><div class="desc">{{$value.Description}}</div></a>
                    </div>
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
            {{end}}
        </div>
    </div>
</div>

</div>
</div>
</div>

<script>
$(document).ready(function() {
    $isEditable = false
    $.fn.editable.defaults.mode = 'inline';
    
    $( "#edit" ).click(function() {
        if (!$isEditable) {   
            $('#name').editable({
                ajaxOptions: {
                    type: "PUT",
                    dataType: "json"
                },
                params: function(params) {
                    return JSON.stringify({
                        ID: params.pk,
                        "name": params.value,			            
                    })
                }
            });
        } else {
            $('#name').editable({
                disabled: true
            });
        }
    });
    
});
</script>
