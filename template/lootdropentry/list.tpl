{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">

	<div class="col-lg-12">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f">Loot Drop Entry List</span>
			</div>

			<div class="panel-body">
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table">
					<thead>
					<tr>
						<th>ItemId</th>
						<th>Charges</th>
						<th><i title="Show As Equipped" class="fa fa-eye"></i></th>
						<th>Multiplier</th>
						<th>Chance</th>
						<th>Min</th>
						<th>Max</th>
						<th><i title="Disabled Chance" class="fa fa-times"></i></th>
					</tr>
					</thead>
					<tbody>
					{{range $key, $value := .LootDropEntrys}}
					<tr>
						<td><a href="/lootdrop/{{$value.LootdropID}}/{{$value.ItemId}}">{{$value.ItemId}}</a></td>
						<td>{{$value.ItemCharges}}</td>
						<td>{{if $value.EquipItem}}<i title="Show As Equipped" class="fa fa-eye"></i>{{end}}</td>
						<td>{{$value.Multiplier}}</td>
						<td>{{$value.Chance}}</td>
						<td>{{$value.Minlevel}}</td>
						<td>{{$value.Maxlevel}}</td>
						<td>{{if $value.DisabledChance}}<i title="Disabled Chance" class="fa fa-times"></i>{{end}}</td>

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