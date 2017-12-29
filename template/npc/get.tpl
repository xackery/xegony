{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-4">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f"><a href="/npc">NPC</a> > {{.Npc.CleanName}}</span>
            </div>
            <div class="panel-body">
                <img alt="logo" class="img-circle m-b m-t-md" src="/images/profile.jpg">
                <h3>{{.Npc.CleanName}}</h3>
                <div class="text-muted font-bold m-b-xs"></div>
                <p>
                    {{.Npc.Id}}
                </p>
                <div class="progress m-t-xs full progress-small">
                    <div style="width: 65%" aria-valuemax="100" aria-valuemin="0" aria-valuenow="65" role="progressbar" class=" progress-bar progress-bar-success">
                        <span class="sr-only">35% Complete (success)</span>
                    </div>
                </div>
            </div>           
            <div class="panel-body">
                <table class="table table-striped">
                    <tbody>
                    <tr><td>Spawns In</td><td>{{.Npc.ZoneName}}</td></tr>
                    <tr><td>HP</td><td>{{comma .Npc.Hp}}</td></tr>
                    <tr><td>Class</td><td>{{.Npc.ClassName}}</td></tr>
                    <tr><td>Race</td><td>{{.Npc.RaceName}}</td></tr>
                    <tr><td>Damage</td><td>{{.Npc.Mindmg}}-{{.Npc.Maxdmg}}</td></tr>
                    </tbody>
                </table>
            </div>

        </div>
    </div>
    <div class="col-lg-8">
            <div class="hpanel">
                <div class="panel-heading hbuilt">
                   Drops
                </div>
                <div class="panel-body">
                    {{if .Items}}
                    <div class="table-responsive">
                    <table cellpadding="1" cellspacing="1" class="table table-striped">

                            <thead>
                            <tr>
                                <th width="20px"><i class="ra ra-sword"></i></th>
                                <th>Name</th>
                                <th>Slot</th>
                                <th><i title="Category" class="ra ra-vest"></i></th>
                            </tr>
                            </thead>
                            <tbody>
                            {{range $key, $value := .Items}}
                            <tr>
                                <td><span class="item icon-{{$value.Icon}}-sm"></span></td>
                                <td><a item={{$value.ItemId}} href="/item/{{$value.ItemId}}">{{$value.Name}}</a></td>
                                <td><span title="{{$value.SlotList}}">{{$value.SlotsFirstName}}</span></td>
                                <td><i title="{{$value.ItemtypeName}}" class="ra {{$value.ItemtypeIcon}}"></i></td>
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
            </div>
        </div>
    </div>
    </div>
</div>
</div>