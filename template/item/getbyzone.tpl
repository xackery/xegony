{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f"><a href="/item/byslot">Items By Zone</a> > {{.Zone.LongName}}</span>
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                      
                        <th width="20px"><i class="xa xa-sword"></i></th>
                        <th>Name</th>
                        <th>Slot</th>                        
                        <th width="20px"><i title="Category" class="xa xa-anvil"></i></th>
                        <th width="20px"><i title="Race" class="xa xa-octopus"></i></th>
                        <th>Drops From</th>
                        <th>Level</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .NpcLoots}}
                    <tr> 
                      
                        <td><span class="slot-sm"><span class="item icon-{{$value.Icon}}-sm"></span></span></td>
                        <td><a item={{$value.ItemID}} href="/npc/{{$value.Item.ID}}">{{$value.Npc.Name}}</a></td>                        
                        <td>{{$value.SlotsFirstName}}</td>                        
                        <td><i title="{{$value.ItemtypeName}}" class="xa {{$value.ItemtypeIcon}}"></i></td>
                          <td><i title="{{$value.Npc.RaceName}}" class="xa {{$value.Npc.RaceIcon}}"></i></td>
                        <td><a href="/npc/{{$value.NpcID}}">{{$value.NpcCleanName}}</a></td>
                        <td>{{$value.Npc.Level}}</td>                        
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