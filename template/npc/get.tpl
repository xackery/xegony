{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-6">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f"><a href="/npc">NPC</a> > {{.Npc.CleanName}}</span>
			</div>
			<div class="panel-body">
				<span class="text-center">                    
					<h3> {{.Npc.CleanName}}</h3>                
				</span>
			</div>           
			<div class="panel-body">
				<table class="table table-striped">
					<tbody>
					<tr><td>Spawns In</td><td><a href="/zone/{{.Npc.ZoneID}}">{{.Npc.ZoneName}}</a></td></tr>
					<tr><td>Class</td><td><i title="{{.Npc.ClassName}}" class="xa {{.Npc.ClassIcon}}"></i> {{.Npc.ClassName}}</td></tr>
					<tr><td>Race</td><td><i title="{{.Npc.RaceName}}" class="xa {{.Npc.RaceIcon}}"></i> 
					{{.Npc.RaceName}}</td></tr>
					<tr><td>Level</td><td>{{.Npc.Level}}</td></tr>
					<tr><td>HP</td><td>{{comma .Npc.Hp}}</td></tr>
					<tr><td>Damage</td><td>{{.Npc.Mindmg}}-{{.Npc.Maxdmg}}</td></tr>
					<tr><td>Experience</td><td><span title="Experience is calculated based on zone experience modifiers and hotzone">{{comma .Npc.Experience}}</span></td></tr>
					{{range $field, $description := .Npc.SpecialAbilitiesList}}                    
						<tr><td>{{$field}}</td><td>{{$description}}</td>
					{{end}}
					</tbody>
				</table>
			</div>

		</div>
	</div>
	{{if .Items}}
	<div class="col-lg-6">
		<div class="hpanel">
			<div class="panel-heading hbuilt">
			   Drops
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
							<td><a item={{$value.ItemId}} href="/item/{{$value.ItemId}}">{{$value.Name}}</a></td>
							<td><span title="{{$value.SlotList}}">{{$value.SlotsFirstName}}</span></td>
							<td><i title="{{$value.ItemtypeName}}" class="xa {{$value.ItemtypeIcon}}"></i></td>
						</tr>
						{{end}}                
						</tbody>
					
				</table>
				</div>
			</div>
		</div>
	</div>
	{{end}}                    
	</div>
	</div>
</div>
</div>