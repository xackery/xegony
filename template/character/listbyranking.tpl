{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">

            <div class="panel-heading">
                <span class="f">Character List By Ranking</span>
            </div>

            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th width="10px">R</th>
                        <th width="10px">C</th>
                        <th>Name</th>
                        <th>Level</th>
                        <th>AA</th>
                        <th>HP</th>
                        <th>Mana</th>
                        <th>AC</th>
                        <th>ATK</th>
                        <th>HP+</th>
                        <th>Mana+</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Characters}}
                    <tr>
                        <td><i title="{{$value.RaceName}}" class="ra {{$value.RaceIcon}}"></i></td>                        
                        <td><i title="{{$value.ClassName}}" class="ra {{$value.ClassIcon}}"></i></td>                        
                        <td><a href="/character/{{$value.Id}}">{{$value.Name}}</a></td>
                        <th>{{$value.Level}}</th>
                        <th>{{$value.AA}}</th>
                        <th>{{$value.TotalHP}}</th>
                        <th>{{$value.TotalMana}}</th>
                        <th>{{$value.AC}}</th>
                        <th>{{$value.ATK}}</th>
                        <th>{{$value.HPRegen}}</th>
                        <th>{{$value.ManaRegen}}</th>
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
</div>