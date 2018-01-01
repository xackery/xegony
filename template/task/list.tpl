{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-12">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f">Task List</span>
			</div>

			<div class="panel-body">
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table">
					<thead>
					<tr>
						<th>ID</th>
						<th>Title</th>
						<th>Type</th>
						<th>Reward</th>
						<th>Cash</th>
						<th>Xp</th>
					</tr>
					</thead>
					<tbody>
					{{range $key, $value := .Tasks}}
					<tr>
						<td><a href="/task/{{$value.ID}}">{{$value.ID}}</a></td>
						<td><a href="/task/{{$value.ID}}">{{$value.Title}}</a></td>
						<td>{{$value.RewardMethodName}}</a></td>
						<td>{{$value.RewardName}}</td>
						<td>{{$value.CashRewardName}}</td>
						<td>{{if gt $value.Xpreward 0}}{{$value.Xpreward}}xp{{end}}</td>
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