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
                                <th width="10px"><i title="race" class="xa xa-bear"></i></th>
                                <th>Name</th>
                                <th>Level</th>
                                <th>Zone</th>  
                            </tr>
                            </thead>
                            <tbody>
                            {{range $key, $value := .Characters}}
                            <tr>            
                                <td><i title="{{$value.RaceName}}" class="xa {{$value.RaceIcon}}"></i></td>
                                <td><a href="/character/{{$value.ID}}">{{$value.CleanName}}</a></td>
                                <td>{{$value.Level}}</td>
                                <td>{{if $value.Zone}}<a href="/zone/{{$value.Zone.ID}}">{{$value.Zone.Name}}</a>{{end}}</td>
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