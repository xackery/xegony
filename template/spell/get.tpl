{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-4">
        <div class="hpanel forum-box">            
            <div class="panel-heading">
                <span class="f"><a href="/spell">Spell</a> > {{.Spell.Name.String}}</span>
            </div>

            <div class="panel-body">                
                <span class="slot slotdrop"><span alt="{{.Spell.Name.String}}" class="item icon-{{.Spell.Icon}}"></span></span>            
            <br/>
            <br/>
        
                <h3>{{.Spell.Name.String}}</h3>
            </div>
            <div class="panel-body">
                <table class="table table-striped">
                    <tbody>
                    <tr><td>Key</td><td>Value</td></tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
    {{if .Items}}
    <div class="col-lg-4">
        <div class="hpanel forum-box">
            <div class="panel-heading">
               Items with {{.Spell.Name.String}}
            </div>
            <div class="panel-body">
                
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table table-striped">

                        <thead>
                        <tr>                            
                            <th width="20px"><i class="xa xa-sword"></i></th>
                            <th>Name</th>
                            <th>Slot</th>
                            <th><i title="Category" class="xa xa-vest"></i></th>
                        </tr>
                        </thead>
                        <tbody>
                        {{range $key, $value := .Items}}
                        <tr>
                            <td><span class="item icon-{{$value.Icon}}-sm"></span></td>
                                <td><a item={{$value.ID}} href="/item/{{$value.ID}}">{{$value.Name}}</a></td>
                                <td><span title="{{$value.SlotList}}">{{$value.SlotsFirstName}}</span></td>
                                <td><i title="{{$value.ItemtypeName}}" class="xa {{$value.ItemtypeIcon}}"></i></td></td>
                        </tr>
                        {{end}}                
                        </tbody>
                    
                </table>
                </div>
            </div>
            <div class="panel-footer">
                {{len .Items}} items have the {{.Spell.Name.String}} effect
            </div>
        </div>

    </div>
    {{end}} 
    {{if .Npcs}}
    <div class="col-lg-4">
        <div class="hpanel">
            <div class="panel-heading hbuilt">
                <div class="panel-tools">
                    <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                    <a class="closebox"><i class="fa fa-times"></i></a>
                </div>
                This spell is used by these creatures:
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table table-striped">
                    <thead>
                    <tr>
                        <th width="10px"><i title="Race" class="xa xa-bear"></i></th>
                        <th width="10px"><i title="Class" class="xa xa-all-for-one"></i></th>
                        <th>Name</th>
                        <th>Zone</th>
                        
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Npcs}}
                    <tr>
                        <td><i title="{{$value.RaceName}}" class="xa {{$value.RaceIcon}}"></i></td>
                                <td><i title="{{$value.ClassName}}" class="xa {{$value.ClassIcon}}"></i></td>
                        <td><a href="/npc/{{$value.ID}}">{{$value.Name}}</a></td>
                        <td><a href="/zone/{{$value.ZoneName}}">{{$value.ZoneName}}</a></td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>

            </div>
            <div class="panel-footer">
                {{len .Npcs}} creatures
            </div>
        </div>
    </div>
    {{end}}
</div>
</div>
</div>