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
			<div class="panel-footer">
			</div>
		</div>
	</div>
	{{if .Map}}
	<div class="col-lg-4">
		<div class="hpanel forum-box">
			<div class="panel-heading">
			   Map
			</div>
			<div class="panel-body">
				<span class="text-center">
				<svg width="300px" height="300px" version="1.1"
				     xmlns="http://www.w3.org/2000/svg" xmlns:xlink= "http://www.w3.org/1999/xlink">
					<image xlink:href="/images/maps/{{.Npc.ZoneName}}.png" x="0" y="0" height="300px" width="300px"/>				
					{{if .Spawns}}
					{{range $key, $value := .Spawns}}
					<circle cx="{{$value.XScaled}}" cy="{{$value.YScaled}}" r="3" fill="red"/>
					{{end}}
					{{end}}
				</svg>
				</span>
				{{if .Spawns}}					
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table table-striped">

						<thead>
						<tr>							
							<th>Name</th>
							<th>Location</th>
						</tr>
						</thead>
						<tbody>
						{{range $key, $value := .Spawns}}
						<tr>
							<td><a href="/spawn/{{$value.SpawngroupID}}">{{$value.Name}}</a></td>
							<td>{{$value.X}}, {{$value.Y}}, {{$value.Z}}</td>							
						</tr>
						{{end}}                
						</tbody>
					
				</table>
				</div>
				{{end}}
			</div>
			{{if .Spawns}}
			<div class="panel-footer">
				{{len .Spawns}} spawn locations
			</div>
			{{end}}
		</div>
	</div>
	{{end}}
	{{if .Items}}
	<div class="col-lg-4">
		<div class="hpanel forum-box">
			<div class="panel-heading">
			   Items
			</div>
			<div class="panel-body">
				
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table table-striped">

						<thead>
						<tr>
							<th>Source</th>							
							<th width="20px"><i class="xa xa-sword"></i></th>
							<th>Name</th>
							<th>Slot</th>
							<th><i title="Category" class="xa xa-vest"></i></th>

						</tr>
						</thead>
						<tbody>
						{{range $key, $value := .Items}}
						<tr>
							<td>{{$value.Reference}}</td>
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
				{{len .Items}} items drop from {{.Npc.CleanName}}
			</div>
		</div>

	</div>
	{{end}}	
	</div>
</div>
</div>

{{if .Map}}
<script>

	$(function () {

	//foo = Snap("#map");
	//foo.circle(300, 300, 100);
	/*s.line(500,500,300,100).attr({strokeWidth:1, stroke:"green"});
var cir = s.circle(300, 300, 2).attr({strokeWidth: 1, stroke:"red", fill:"maroon"});
console.log((parseFloat(cir.attr("cx"))+parseFloat(100)))
cir.animate({cx: 445.5260009765625, cy: 301.7739990234375}, 5000);*/
});
</script>
{{end}}