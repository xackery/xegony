{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Zone List</span>
            </div>

            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>                        
                        <th>Name</th>   
                        <th>Expansion</th>         
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Zones}}
                    <tr>
                        <td><a href="/npc/zone/{{$value.ID}}">{{$value.LongName}}</a></td>
                        <td><a href="/npc/zone/{{$value.ID}}">{{$value.Expansion.Name}}</a></td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>

            </div>
            <div class="panel-footer">
                {{.Page.Total}} total zones
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