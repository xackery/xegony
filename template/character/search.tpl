{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

    <div class="row">
        <div class="col-lg-12">

            <div class="hpanel">
                <div class="panel-body">
                    <form method="GET" action="/character/search">
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
                <span class="f">Search results for '{{.Search}}'</span>
            </div>


                <div class="panel-body">
                    {{if .Characters}}
                    <div class="table-responsive">
                    <table cellpadding="1" cellspacing="1" class="table">

                            <thead>
                            <tr>
                                <th width="10px"><i class="ra ra-double-team" title="Race"></i></th>
                        <th width="10px"><i class="ra ra-pawn" title="Class"></i></th>
                                <th>Name</th>
                                <th>ZoneID</th>                        
                            </tr>
                            </thead>
                            <tbody>
                            {{range $key, $value := .Characters}}
                            <tr>
                                <td><i class="ra {{$value.RaceIcon}}"></i></td>
                                <td><i class="ra {{$value.ClassIcon}}"></i></td                                
                                <td><a href="/character/{{$value.Id}}">{{$value.Name}}</a></td>
                                
                                <td><a href="/zone/{{$value.ZoneId}}">{{$value.ZoneId}}</a></td>         
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