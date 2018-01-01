{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

	<div class="row">
		<div class="col-lg-4">
			<div class="hpanel forum-box">
				<div class="panel-heading">
					<span class="f"><a href="/spawn">Spawn</a> > {{.SpawnEntry.NpcID}}</span>
				</div>
				<div class="panel-body">
					<h3>{{.SpawnEntry.NpcID}}</h3>
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
	</div>
</div>
</div>