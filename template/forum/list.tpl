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
        <div class="table-responsive">
        <table cellpadding="1" cellspacing="1" class="table table-condensed table-borderless">
            <col width="40px">
            <tbody>
                <tr class="newRow">
                    <td rowspan="2"><a id="newIcon" class="newField" data-type="select2" data-title="Select icon"> <h1><i class="xa xa-monster-skull"></i></h1></a></td>
                    <td colspan="2"><a id="newName" class="newField" data-type="text" data-title="Enter forum name"><h4>Name</h4></a></td>    
                    <td><span class="number"> 4780 </span><small>Views</small></td>
                </tr>
                <tr class="newRow">
                    <td colspan="2"><div id="newDescription" class="newField" class="desc" data-type="text" data-title="Edit description">Description</div></td>
                    <td><span class="number"> 150 </span><small>Posts</small></td>
                </tr>    
            {{range $key, $value := .Forums}}                            
               
                <tr>
                    <td rowspan="2"><a class="editLink" href="/forum/{{$value.ID}}"> <h1><div class="editIcon" data-type="select2" data-pk="{{$value.ID}}" data-url="/api/forum/{{$value.ID}}" data-title="Select icon"> <i class="{{$value.Icon}}" ></i></div></h1></a></td>
                    <td colspan="2"><a class="editLink" href="/forum/{{$value.ID}}"> <h4 class="editName" data-type="text" data-pk="{{$value.ID}}" data-url="/api/forum/{{$value.ID}}" data-title="Enter forum name">{{$value.Name}}</h4></a></td>
                    <td><span class="number"> 4780 </span><small>Views</small></td>
                    </a>
                </tr>
                <tr>
                    <td colspan="2"><a class="editLink" href="/forum/{{$value.ID}}"> <div class="desc editDescription" data-pk="{{$value.ID}}" data-type="text" data-url="/api/forum/{{$value.ID}}" data-title="Edit description">{{$value.Description}}</div></a></td>
                    <td><span class="number"> 150 </span><small>Posts</small></td>
                </tr>
            {{end}}                                      
            </tbody>
        </table>
        </div>
    </div>
</div>

</div>
</div>
</div>

<script>
$(document).ready(function() {
	$isEditMode = false;
    $isAddable = false;
    $.fn.editable.defaults.mode = 'inline';

    function ToggleEditMode() {
        $.each([ ".editName", ".editDescription", ".editIcon" ], function( index, value ) {
            $(value).editable('toggleDisabled');
        });
        $isEditMode = !$isEditMode;
        if ($isEditMode) {
            $(".editLink").attr("onclick", "");
        } else {
            $(".editLink").attr("onclick", "return false;");
        }
    }

    $(".newRow").hide();
    $('#newName').editable({
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
    $("#add").click(function() {
        $isAddable = !$isAddable;
         if ($isAddable) {
             if (!$isEditMode) {
                ToggleEditMode();
             }
             $(".newRow").show()
            console.log("Enabling add mode")
        } else {
             $(".newRow").hide()
        } 
    });

    $.each([ ".newName", ".newDescription", ".newIcon" ], function( index, value ) {
        var valField = value.substring(5).toLowerCase()
        var source = "";
        if (valField == "icon") {
            source = [
                {id: 'xa-monster-skull', text: 'skull'},
                {id: 'xa-foo', text: 'foo'},
            ]
        }
        $(value).editable({
            source: source,
            ajaxOptions: {
                type: "POST",
                dataType: "json"
            },
            params: function(params) {
                var paramData = {
                    ID: params.pk
                }
                paramData[valField] = params.value
                return JSON.stringify(paramData)
            }
        });
        $(value).editable('toggleDisabled');        
    });


    
    $.each([ ".editName", ".editDescription", ".editIcon" ], function( index, value ) {
        var valField = value.substring(5).toLowerCase()
        var source = "";
        if (valField == "icon") {
            source = [
                {id: 'xa-monster-skull', text: 'skull'},
                {id: 'xa-foo', text: 'foo'},
            ]
        }
        $(value).editable({
            source: source,
            ajaxOptions: {
                type: "PUT",
                dataType: "json"
            },
            params: function(params) {
                var paramData = {
                    ID: params.pk
                }
                paramData[valField] = params.value
                return JSON.stringify(paramData)
            }
        });
        $(value).editable('toggleDisabled');        
    });

    $("#edit").click(function() {
        ToggleEditMode();        

    });
    
});
</script>
