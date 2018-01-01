{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-12">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f">Loot Table List</span>
			</div>

			<div class="panel-body">
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table">
					<thead>
					<tr>
						<th>Name</th>
						<th>MinCash</th>
						<th>MaxCash</th>
						<th>AvgCoin</th>
					</tr>
					</thead>
					<tbody>
					{{range $key, $value := .LootTables}}
					<tr>
						<td><a href="/loottable/{{$value.ID}}">{{$value.Name}}</a></td>
						<td>{{$value.MinCashName}}</td>
						<td>{{$value.MaxCashName}}</td>
						<td>{{$value.AvgCoinName}}</td>
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