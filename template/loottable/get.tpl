{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

	<div class="row">
		<div class="col-lg-4">
			<div class="hpanel forum-box">
				<div class="panel-heading">
					<span class="f"><a href="/table">Table</a> > {{.LootTable.Name}}</span>
				</div>
				<div class="panel-body">
					<h3>{{.LootTable.Name}}</h3>
					<div class="text-muted font-bold m-b-xs"></div>
					<p>
					</p>
					<div class="progress m-t-xs full progress-small">
						<div style="width: 65%" aria-valuemax="100" aria-valuemin="0" aria-valuenow="65" role="progressbar" class=" progress-bar progress-bar-success">
							<span class="sr-only">35% Complete (success)</span>
						</div>
					</div>
				</div>           
				<div class="panel-body">
				<table class="table table-striped">
					<tbody>
					<tr><td>Minimum Cash</td><td>{{.LootTable.Mincash}}</a></td></tr>
					<tr><td>Maximum Cash</td><td>{{.LootTable.Maxcash}}</a></td></tr>
					<tr><td>Average Coin</td><td>{{.LootTable.Avgcoin}}</a></td></tr>
					</tbody>
				</table>
			</div>

			</div>
		</div>
		{{if .LootTable.Entries}}
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
								<th>LootDrop ID</th>
								<th>Multiplier</th>
								<th>DropLimit</th>
								<th>MinDrop</th>
								<th>Probability</th>
							</tr>
							</thead>
							<tbody>
							{{range $key, $value := .LootTable.Entries}}
							<tr>
								<td><a href="/lootdrop/{{$value.LootdropID}}">{{$value.LootdropID}}</a></td>
								<td><a href="/lootdrop/{{$value.LootdropID}}">{{$value.Multiplier}}</a></td>
								<td><a href="/lootdrop/{{$value.LootdropID}}">{{$value.Droplimit}}</a></td>
								<td><a href="/lootdrop/{{$value.LootdropID}}">{{$value.Mindrop}}</a></td>
								<td><a href="/lootdrop/{{$value.LootdropID}}">{{$value.Probability}}</a></td>
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
		{{if .LootTable.Npcs}}
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
							{{range $key, $value := .LootTable.Npcs}}
							<tr>
								<td><a href="/npc/{{$value.ID}}">{{$value.ID}}</a></td>
								<td><a href="/npc/{{$value.ID}}">{{$value.Name}}</a></td>								
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