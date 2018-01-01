{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-12">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f">Loot Drop List</span>
			</div>

			<div class="panel-body">
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table">
					<thead>
					<tr>
						<th>ItemId</th>
						<th>Name</th>
					</tr>
					</thead>
					<tbody>
					{{range $key, $value := .LootDrops}}
					<tr>
						<td><a href="/lootdrop/{{$value.ID}}">{{$value.ID}}</a></td>
						<td>{{$value.Name}}</td>						
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