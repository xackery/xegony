{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                Bestiary
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th width="10px"><i title="Race" class="xa xa-bear"></i></th>
                        <th width="10px"><i title="Class" class="xa xa-all-for-one"></i></th>
                        <th>Name</th>
                        <th>CharacterID</th>                   
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Npcs}}
                    <tr>
                        <td><i title="{{$value.Race.Name}}" class="xa {{$value.Race.Icon}}"></i></td>
                        <td><i title="{{$value.Class.Name}}" class="xa {{$value.Class.Icon}}"></i></td>
                        <td><a href="/npc/{{$value.ID}}">{{$value.CleanName}}</a></td>
                        <td>{{$value.ID}}</td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>
                
            </div>
            <div class="panel-footer">
                {{.Page.Total}} total creatures
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