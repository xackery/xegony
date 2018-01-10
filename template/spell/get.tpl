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
        
                <h3>Spell: {{.Spell.Name.String}}</h3>
            </div>
            <div class="panel-body">
                <table class="table table-striped">
                    <tbody>
                    {{if .Spell.ID}}<tr><td>ID</td><td>{{.Spell.ID}}</td>{{end}}
                    <tr><td>Classes</td><td>{{.Spell.ClassesList}}</td>
                    {{if .Spell.Mana}}<tr><td>Mana</td><td>{{.Spell.Mana}}</td>{{end}}
                    <tr><td>Skill</td><td>{{.Spell.Skill.Name}}</td>
                    <tr><td>CastTime</td><td>{{.Spell.CastTimeName}}</td>
                    {{if .Spell.RecoveryTimeName}}<tr><td>Recovery Time</td><td>{{.Spell.RecoveryTimeName}}</td>{{end}}
                    {{if .Spell.RecastTimeName}}<tr><td>RecastTime</td><td>{{.Spell.RecastTimeName}}</td>{{end}}
                    {{if .Spell.Range}}<tr><td>Range</td><td>{{.Spell.Range}}</td>{{end}}
                    <tr><td>Target</td><td>{{.Spell.TargetTypeName}}</td>
                    <tr><td>Resist</td><td>{{.Spell.ResistTypeName}}</td>
                    <tr><td>Duration</td><td>{{.Spell.BuffDurationName}}</td>
                    {{range $key, $value := .Spell.Reagents}}
                    <tr><td>Reagent</td><td>{{$value.Name}}</td>
                    {{end}}


                    {{if .Spell.TeleportZone.String}}<tr><td>TeleportZone</td><td>{{.Spell.TeleportZone.String}}</td>{{end}}
                    {{if .Spell.YouCast.String}}<tr><td>YouCast</td><td>{{.Spell.YouCast.String}}</td>{{end}}
                    {{if .Spell.OtherCasts.String}}<tr><td>OtherCasts</td><td>{{.Spell.OtherCasts.String}}</td>{{end}}
                    {{if .Spell.CastOnYou.String}}<tr><td>CastOnYou</td><td>{{.Spell.CastOnYou.String}}</td>{{end}}
                    {{if .Spell.CastOnOther.String}}<tr><td>CastOnOther</td><td>{{.Spell.CastOnOther.String}}</td>{{end}}
                    {{if .Spell.SpellFades.String}}<tr><td>SpellFades</td><td>{{.Spell.SpellFades.String}}</td>{{end}}
                    
                    {{if .Spell.Aoerange}}<tr><td>Aoerange</td><td>{{.Spell.Aoerange}}</td>{{end}}
                    {{if .Spell.Pushback}}<tr><td>Pushback</td><td>{{.Spell.Pushback}}</td>{{end}}
                    {{if .Spell.Pushup}}<tr><td>Pushup</td><td>{{.Spell.Pushup}}</td>{{end}}
                    </tbody>
                </table>
            </div>
        </div>
     </div>
     <div class="col-lg-4">
          <div class="hpanel forum-box">            
            <div class="panel-heading">
                Spell Effects
            </div>
               
            <div class="panel-body">
                <table class="table table-striped">
                    <tbody>
                    {{range $key, $value := .Spell.Effects}}
                    <tr><td>{{$value.ID}}</td><td>{{$value.Name}}</td><td>{{$value.FormulaName}}</td><td>{{$value.BaseValue}}</td></tr>
                    {{end}}
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