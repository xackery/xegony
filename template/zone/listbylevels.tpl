{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Hot Zones List</span>
            </div>

            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table table-bordered table-striped">
                    <thead>
                    <tr>
                        <th>Name</th>   
                        <th><span title="Experience Bonus takes Hotzone and zone multipliers into account">ExpBonus</span></th>
                        <th>Expansion</th>    
                        <th>1</th>
                        <th>5</th>
                        <th>10</th>
                        <th>15</th>
                        <th>20</th>
                        <th>25</th>
                        <th>30</th>
                        <th>35</th>
                        <th>40</th>
                        <th>45</th>
                        <th>50</th>
                        <th>55</th>
                        <th>60</th>     
                        <th>65</th>     
                        <th>70</th>     
                        <th>75</th>     
                        <th>80</th>     

                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Zones}}
                    <tr>
                        <td><a href="/zone/{{$value.ZoneIDNumber}}">{{$value.LongName}}</a></td>
                        <td><a href="/zone/{{$value.ZoneIDNumber}}">{{$value.Modifier}}x</a></td>
                        <td><a href="/zone/{{$value.ZoneIDNumber}}">{{$value.ExpansionName}}</a></td>
                        <td>{{if iszonelevel 1 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 1">X</span>{{end}}</td>
                        <td>{{if iszonelevel 5 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 5">X</span>{{end}}</td>
                        <td>{{if iszonelevel 10 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 10">X</span>{{end}}</td>
                        <td>{{if iszonelevel 15 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 15">X</span>{{end}}</td>
                        <td>{{if iszonelevel 20 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 20">X</span>{{end}}</td>
                        <td>{{if iszonelevel 25 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 25">X</span>{{end}}</td>
                        <td>{{if iszonelevel 30 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 30">X</span>{{end}}</td>
                        <td>{{if iszonelevel 35 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 35">X</span>{{end}}</td>
                        <td>{{if iszonelevel 40 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 40">X</span>{{end}}</td>
                        <td>{{if iszonelevel 45 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 45">X</span>{{end}}</td>
                        <td>{{if iszonelevel 50 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 50">X</span>{{end}}</td>
                        <td>{{if iszonelevel 55 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 55">X</span>{{end}}</td>
                        <td>{{if iszonelevel 60 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 60">X</span>{{end}}</td>
                        <td>{{if iszonelevel 65 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 65">X</span>{{end}}</td>
                        <td>{{if iszonelevel 70 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 70">X</span>{{end}}</td>                        
                        <td>{{if iszonelevel 75 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 75">X</span>{{end}}</td>
                        <td>{{if iszonelevel 80 $value.Levels}}<span title="{{$value.LongName}} is viable to hunt at, at level 80">X</span>{{end}}</td>
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