{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

	<div class="row">
		<div class="col-lg-4">
			<div class="hpanel forum-box">
				<div class="panel-heading">
					<span class="f"><a href="/spawn">Spawn</a> > {{.Spawn.Name}}</span>
				</div>
				<div class="panel-body">
					<h3>{{.Spawn.Name}}</h3>
				</div>           
				<div class="panel-body">
				<table class="table table-striped">
					<tbody>
					<tr><td>SpawngroupID</td><td>{{.Spawn.SpawngroupID}}</td></tr>
					<tr><td>Zone</td><td>{{.Spawn.Zone.String}}</td></tr>
					{{if .Spawn.Version}}<tr><td>Version</td><td>{{.Spawn.Version}}</td></tr>{{end}}
					<tr><td>X</td><td>{{.Spawn.X}}</td></tr>
					<tr><td>Y</td><td>{{.Spawn.Y}}</td></tr>
					<tr><td>Z</td><td>{{.Spawn.Z}}</td></tr>
					<tr><td>Heading</td><td>{{.Spawn.Heading}}</td></tr>
					<tr><td>RespawnTime</td><td>{{.Spawn.RespawnTime}}</td></tr>
					<tr><td>Variance</td><td>{{.Spawn.Variance}}</td></tr>
					<tr><td>Pathgrid</td><td>{{.Spawn.Pathgrid}}</td></tr>
					<tr><td>Condition</td><td>{{.Spawn.Condition}}</td></tr>
					<tr><td>CondValue</td><td>{{.Spawn.CondValue}}</td></tr>
					<tr><td>Enabled</td><td>{{.Spawn.Enabled}}</td></tr>
					<tr><td>Animation</td><td>{{.Spawn.Animation}}</td></tr>
					<tr><td>SpawnLimit</td><td>{{.Spawn.SpawnLimit}}</td></tr>
					<tr><td>Dist</td><td>{{.Spawn.Dist}}</td></tr>
					<tr><td>MaxX</td><td>{{.Spawn.MaxX}}</td></tr>
					<tr><td>MinX</td><td>{{.Spawn.MinX}}</td></tr>
					<tr><td>MaxY</td><td>{{.Spawn.MaxY}}</td></tr>
					<tr><td>MinY</td><td>{{.Spawn.MinY}}</td></tr>
					<tr><td>Delay</td><td>{{.Spawn.Delay}}</td></tr>
					<tr><td>Mindelay</td><td>{{.Spawn.Mindelay}}</td></tr>
					<tr><td>Despawn</td><td>{{.Spawn.Despawn}}</td></tr>
					<tr><td>DespawnTimer</td><td>{{.Spawn.DespawnTimer}}</td></tr>
					</tbody>
				</table>
			</div>

			</div>
		</div>
		{{if .SpawnEntrys}}
		<div class="col-lg-8">
			<div id="row">
				<div class="hpanel">
					<div class="panel-heading hbuilt">
						<div class="panel-tools">
							<a class="showhide"><i class="fa fa-chevron-up"></i></a>
							<a class="closebox"><i class="fa fa-times"></i></a>
						</div>
						Loot Table Entries
					</div>
					<div class="panel-body">
						<div class="table-responsive">
						<table cellpadding="1" cellspacing="1" class="table">
							<thead>
							<tr>
								<th>SpawngroupID ID</th>
								<th>NpcID</th>
								<th>Chance</th>
							</tr>
							</thead>
							<tbody>
							{{range $key, $value := .SpawnEntrys}}
							<tr>								
								<td><a href="/spawn/{{$value.SpawngroupID}}/{{$value.NpcID}}">{{$value.SpawngroupID}}</a></td>
								<td><a href="/spawn/{{$value.SpawngroupID}}/{{$value.NpcID}}">{{$value.NpcID}}</a></td>
								<td><a href="/spawn/{{$value.SpawngroupID}}/{{$value.NpcID}}">{{$value.Chance}}</a></td>
								<td><a href="/spawn/{{$value.SpawngroupID}}/{{$value.NpcID}}">{{$value.Mindrop}}</a></td>
								<td><a href="/spawn/{{$value.SpawngroupID}}/{{$value.NpcID}}">{{$value.Probability}}</a></td>
							</tr>
							{{end}}
							</tbody>
						</table>
						</div>

					</div>
				</div>
			</div>
		</div>
		{{end}}
		{{if .Spawn.Npcs}}
		<div class="col-lg-8">
			<div id="row">
				<div class="hpanel">
					<div class="panel-heading hbuilt">
						<div class="panel-tools">
							<a class="showhide"><i class="fa fa-chevron-up"></i></a>
							<a class="closebox"><i class="fa fa-times"></i></a>
						</div>
						Bestiary Using This Table
					</div>
					<div class="panel-body">
						<div class="table-responsive">
						<table cellpadding="1" cellspacing="1" class="table">
							<thead>
							<tr>
								<th>ID</th>
								<th>Name</th>
							</tr>
							</thead>
							<tbody>
							{{range $key, $value := .Spawn.Npcs}}
							<tr>
								<td><a href="/npc/{{$value.Id}}">{{$value.Id}}</a></td>
								<td><a href="/npc/{{$value.Id}}">{{$value.Name}}</a></td>								
							</tr>
							{{end}}
							</tbody>
						</table>
						</div>

					</div>
				</div>
			</div>
		</div>
		{{end}}
	</div>
</div>
</div>