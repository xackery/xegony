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
					<tr><td>ID</td><td>{{.Spawn.ID}}</td></tr>
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
		{{if .SpawnNpcs}}
		<div class="col-lg-8">
			<div id="row">
				<div class="hpanel">
					<div class="panel-heading hbuilt">
						<div class="panel-tools">
							<a class="showhide"><i class="fa fa-chevron-up"></i></a>
							<a class="closebox"><i class="fa fa-times"></i></a>
						</div>
						Spawn Npcs
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
							{{range $key, $value := .SpawnNpcs}}
							<tr>								
								<td><a href="/spawn/{{$value.SpawnID}}/npc/{{$value.NpcID}}">{{$value.SpawnID}}</a></td>
								<td><a href="/npc/{{$value.NpcID}}">{{$value.NpcID}}</a></td>
								<td><a href="/spawn/{{$value.SpawnID}}/npc/{{$value.NpcID}}">{{$value.Chance}}</a></td>
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