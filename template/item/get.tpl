{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-4">
        <div class="hpanel forum-box">            
            <div class="panel-heading">
                <span class="f"><a href="/item">Item</a> > {{.Item.Name}}</span>
            </div>

            <div class="panel-body">                
                <span class="slot slotdrop"><span alt="{{.Item.Name}}" class="item icon-{{.Item.Icon}}"></span></span>
            
            <br/>
            <br/>
        
                <h3>{{.Item.Name}}</h3>
            </div>
            <div class="panel-body">
                <table class="table table-striped">
                    <tbody>
                    {{if .Item.ID}}<tr><td>ID</td><td>{{.Item.ID}}</td></tr>{{end}}
                    {{if .Item.Name}}<tr><td>Name</td><td>{{.Item.Name}}</td></tr>{{end}}                    
                    </tbody>
                </table>
            </div>
            <div class="panel-footer contact-footer">
                <div class="row">
                    <div class="col-md-4 border-right">
                        <div class="contact-stat"><span>Projects: </span> <strong>200</strong></div>
                    </div>
                    <div class="col-md-4 border-right">
                        <div class="contact-stat"><span>Messages: </span> <strong>300</strong></div>
                    </div>
                    <div class="col-md-4">
                        <div class="contact-stat"><span>Views: </span> <strong>400</strong></div>
                    </div>
                </div>
            </div>

        </div>
    </div>
    {{if .Npcs}}
    <div class="col-lg-8">
        <div class="hpanel">
            <div class="panel-heading hbuilt">
                <div class="panel-tools">
                    <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                    <a class="closebox"><i class="fa fa-times"></i></a>
                </div>
                This item is found off mobs:
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
               {{len .Npcs}} total creatures
            </div>
        </div>
    </div>
    {{end}}
    {{if .Merchants}}
    <div class="col-lg-8">
        <div class="hpanel">
            <div class="panel-heading hbuilt">
                <div class="panel-tools">
                    <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                    <a class="closebox"><i class="fa fa-times"></i></a>
                </div>
                This item is found off merchants:
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
                    {{range $merchant := .Merchants}}
                         {{if $merchant}}
                         {{range $value := $merchant.Merchant.Npcs}}
                    
                    <tr>
                        <td><i title="{{$value.RaceName}}" class="xa {{$value.RaceIcon}}"></i></td>
                                <td><i title="{{$value.ClassName}}" class="xa {{$value.ClassIcon}}"></i></td>
                        <td><a href="/npc/{{$value.ID}}">{{$value.Name}}</a></td>
                        <td><a href="/zone/{{$value.ZoneName}}">{{$value.ZoneName}}</a></td>
                    </tr>          
                    {{end}}
                    {{end}}
                    {{end}}                
                    </tbody>
                </table>
                </div>

            </div>
            <div class="panel-footer">
               {{len .Npcs}} total creatures
            </div>
        </div>
    </div>
    {{end}}
    {{if .Recipes}}
    <div class="col-lg-8">
        <div class="hpanel">
            <div class="panel-heading hbuilt">
                <div class="panel-tools">
                    <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                    <a class="closebox"><i class="fa fa-times"></i></a>
                </div>
                This item is used in recipes
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table table-striped">
                    <thead>
                    <tr>
                        <th>Name</th>                        
                        <th>Trivial</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Recipes}}
                    <tr>
                       {{if $value.Recipe}}<td><a href="/recipe/{{$value.Recipe.ID}}">{{$value.Recipe.Name}}</a>{{end}}
                       <td>{{$value.Recipe.Trivial}}</td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>

            </div>
            <div class="panel-footer">
               {{len .Npcs}} total creatures
            </div>
        </div>
    </div>
    {{end}}
    {{if .Fishings}}
    <div class="col-lg-8">
        <div class="hpanel">
            <div class="panel-heading hbuilt">
                <div class="panel-tools">
                    <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                    <a class="closebox"><i class="fa fa-times"></i></a>
                </div>
                This item is found by fishing
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table table-striped">
                    <thead>
                    <tr>
                        <th>Zone</th>                        
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Fishings}}
                    <tr>
                       {{if $value.ZoneID}}<td><a href="/fishing/{{$value.ID}}">{{$value.Zone.Name}}</a></td><td>Unknown Zone for {{$value.ID}}</td>{{else}}{{end}}
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>

            </div>
            <div class="panel-footer">
               {{len .Npcs}} total creatures
            </div>
        </div>
    </div>
    {{end}}
</div>

</div>

</div>