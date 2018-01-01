{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Hot Zones, EXP zones List</span>
            </div>

            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th>Name</th>   
                        <th><span title="ExpBonus uses the formula: 1+ZEM*HOTZONE">ExpBonus</span></th>
                        <th>Hotzone</th>

                        <th><span title="Zone Experience Multiplier">ZEM</span></th>
                        <th>Expansion</th>         
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Zones}}
                    <tr>
                        <td><a href="/zone/{{$value.ZoneIDNumber}}">{{$value.LongName}}</a></td>
                        <td><a title="ExpBonus is based on: 1+(ZEM={{$value.ZoneExpMultiplier}}){{if $value.Hotzone}}*HOTZONE{{end}}" href="/zone/{{$value.ZoneIDNumber}}">{{$value.Modifier}}x</a></td>
                        <td><a href="/zone/{{$value.ZoneIDNumber}}">{{$value.Hotzone}}</a></td>
                        <td><a href="/zone/{{$value.ZoneIDNumber}}">{{$value.ZoneExpMultiplier}}</a></td>
                        <td><a href="/zone/{{$value.ZoneIDNumber}}">{{$value.ExpansionName}}</a></td>
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