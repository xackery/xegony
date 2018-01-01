{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">

            <div class="panel-heading">
                <span class="f">Characters Online</span>
            </div>

            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th width="10px"><i class="xa xa-double-team" title="Race"></i></th>
                        <th width="10px"><i class="xa xa-pawn" title="Class"></i></th>
                        <th>Name</th>
                        <th>ZoneID</th>                        
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Characters}}
                    <tr>
                        <td><i class="xa {{$value.RaceIcon}}"></i></td>
                        <td><i class="xa {{$value.ClassIcon}}"></i></td>                        
                        <td><a href="/character/{{$value.ID}}">{{$value.Name}}</a></td>
                        
                        <td><a href="/zone/{{$value.ZoneID}}">{{$value.ZoneID}}</a></td>         
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
    <!-- Footer-->
    <footer class="footer">
        <span class="pull-right">
            Example text
        </span>
        Company 2015-2020
    </footer>

</div>