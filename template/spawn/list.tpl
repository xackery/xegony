{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-12">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f">Spawn List</span>
			</div>

			<div class="panel-body">
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table">
					<thead>
					<tr>
						<th>Name</th>
						<th>SpawnLimit</th>
						<th>Delay</th>
						<th>MinDelay</th>
					</tr>
					</thead>
					<tbody>
					{{range $key, $value := .Spawns}}
					<tr>
						<td><a href="/spawn/{{$value.ID}}/entry">{{$value.Name}}</a></td>
						<td>{{$value.SpawnLimit}}</td>
						<td>{{$value.Delay}}</td>
						<td>{{$value.MinDelay}}</td>
					</tr>
					{{end}}                
					</tbody>
				</table>
				</div>

			</div>
		</div>
	</div>
</div>

</div>