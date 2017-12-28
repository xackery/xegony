{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-12">
		<div class="hpanel forum-box">

			<div class="panel-heading">
				<span class="f">Character List</span>
			</div>

			<div class="panel-body">
				<div class="table-responsive" class="tooltip-info">
				<table cellpadding="1" cellspacing="1" class="table  table-striped">
					<thead>
					<tr>
						<th width="10px">R</th>
						<th width="10px">C</th>
						<th>Name</th>	
						<th>Level</th>					
						<th>ZoneID</th>
					</tr>
					</thead>
					<tbody>
					{{range $key, $value := .Characters}}
					<tr>
						<td>><i title="{{$value.RaceName}}" class="ra {{$value.RaceIcon}}""></i></td>
						<td><i class="ra {{$value.ClassIcon}}" title="{{$value.ClassName}}"></i></td>						
						<td><a href="/character/{{$value.Id}}">{{$value.Name}}</a></td>
						<td>{{$value.Level}}</td>
						<td><a href="/zone/{{$value.ZoneId}}">{{$value.ZoneId}}</a></td>         
					</tr>
					{{end}}                
					</tbody>
				</table>
				</div>

			</div>
			<div class="panel-footer">
				Table - 6 rows
			</div>
		</div>
	</div>
</div>
</div>