{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="normalheader small-header">
    <div class="hpanel">
        <div class="panel-body">
           

            <div id="hbreadcrumb" class="pull-right m-t-lg">
                <ol class="hbreadcrumb breadcrumb">
                    <li class="active">
                        <span>Character</span>
                    </li>
                </ol>
            </div>
            <h2 class="font-light m-b-xs">
               Search for {{.Search}}
            </h2>
            <small>There are 215 characters on the server.</small>
        </div>
    </div>
</div>

<div class="content">
    {{if .Characters}}
    <div class="row">
        <div class="col-lg-12">
            <div class="hpanel">
                <div class="panel-body">
                    <div class="table-responsive">
                    <table cellpadding="1" cellspacing="1" class="table">
                        <thead>
                        <tr>
                            <th>Name</th>
                            <th>CharacterID</th>
                            <th>ZoneID</th>                        
                        </tr>
                        </thead>
                        <tbody>
                        {{range $key, $value := .Characters}}
                        <tr>
                            
                            <td><a href="/character/{{$value.Id}}">{{$value.Name}}</a></td>
                            <td>{{$value.Id}}</td>
                            <td><a href="/zone/{{$value.ZoneId}}">{{$value.ZoneId}}</a></td>         
                        </tr>
                        {{end}}                
                        </tbody>
                    </table>
                    </div>

                </div>
                <div class="panel-footer">
                    Table - 6 rows
                </div>
            </div>
        </div>
    </div>
    {{else}}
    
    {{end}}
</div>