{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

	<div class="row">
		<div class="col-lg-6">
			<div class="hpanel forum-box">
				<div class="panel-heading">
					<span class="f"><a href="/spawn">Spawn</a> > {{.SpawnEntry.NpcID}}</span>
				</div>
				<div class="panel-body">
					<h3>Spawn Entry {{.SpawnEntry.NpcID}}</h3>
				</div>           
				<div class="panel-body">
				<table class="table table-striped">
					<tbody>
					<tr><td>SpawngroupID</td><td>{{.SpawnEntry.SpawngroupID}}</td></tr>
					<tr><td>NpcID</td><td>{{.SpawnEntry.NpcID}}</td></tr>
					<tr><td>Chance</td><td>{{.SpawnEntry.Chance}}</td></tr>					
					</tbody>
				</table>
			</div>

			</div>
		</div>		
		<div class="col-lg-6">
			<div class="hpanel forum-box">
				<div class="panel-heading">
					Spawn Entry NPC Details
				</div>
				<div class="panel-body">
					<h3>{{.Npc.CleanName}}</h3>
				</div>           
				<div class="panel-body">
				<table class="table table-striped">
					<tbody>
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

	</div>
</div>
</div>