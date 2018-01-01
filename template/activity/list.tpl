{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-4">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f"><a href="/task/">Task</a> > <a href="/task/{{.Task.ID}}">{{.Task.Title}}</a></span>
			</div>

			<div class="panel-body">
					Task Details Go here				
			</div>
		</div>
	</div>
	<div class="col-lg-8">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f">Activities</span>
			</div>

			<div class="panel-body">
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table">
					<thead>
					<tr>
						<th>ID</th>
						<th>Name</th>
						<th>Description</th>
					</tr>
					</thead>
					<tbody>
					{{range $key, $value := .Activitys}}
					<tr>
						<td><a href="/task/{{$value.Taskid}}/activity/{{$value.Activityid}}">{{$value.Activityid}}</a></td>
						<td><a href="/task/{{$value.Taskid}}/activity/{{$value.Activityid}}">{{$value.Activitytype}}</a></td>
						<td><a href="/task/{{$value.Taskid}}/activity/{{$value.Activityid}}">{{$value.Text3}}</a></td>
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