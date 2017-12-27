{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

    <div class="row">
        <div class="col-lg-12">

            <div class="hpanel">
                <div class="panel-body">
                    <form method="GET" action="/item/search/byaccount">
                        <div class="text-muted small pull-right text-right">            
                        </div>
                        <div class="form-inline">
                            Search: <input class="form-control" name="search" {{if .Search}}value="{{.Search}}"{{end}} type="text">
                        </div>
                    </form>

                </div>

            </div>

            {{if .Search}}
            <div class="hpanel forum-box">

            <div class="panel-heading">
                <span class="f">Inventory Search results for '{{.Search}}' across your characters</span>
            </div>


                <div class="panel-body">
                    {{if .Items}}
                    <div class="table-responsive">
                    <table cellpadding="1" cellspacing="1" class="table">

                            <thead>
                            <tr>
                                <th><i class="ra ra-sword"></i></th>
                                <th>Item Name</th>
                                <th>Character</th>
                                <th>Location</th>
                            </tr>
                            </thead>
                            <tbody>
                            {{range $key, $value := .Items}}
                            <tr>                               
                            <td><span class="slot-sm"><span title="{{$value.Name}}" class="item icon-{{$value.Icon}}-sm"></span></span></td> 
                                <td><a href="/item/{{$value.Id}}">{{$value.Name}}</a></td>
                                <td><a href="/character/{{$value.Character.Id}}">{{$value.Character.Name}}</a></td>
                                <td>{{$value.SlotName}}</td>
                            </tr>
                            {{end}}                
                            </tbody>
                        
                    </table>
                    </div>
                    {{else}}
                    {{if .Search}}
                    No results for {{.Search}} were found
                    {{end}}
                    {{end}}

                </div>
                <div class="panel-footer">
                    Table - 6 rows
                </div>
            </div>
            {{end}}

        </div>
    </div>
</div>