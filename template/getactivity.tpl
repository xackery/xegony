{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-12">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f"><a href="/task">Task</a> > <a href="/task/{{.Task.Id}}">{{.Task.Title}}</a> > Activity {{.Activity.Activityid}}</span>
			</div>

			<div class="panel-body">
				
						{{.Activity.Text3}}


			</div>
			<div class="panel-footer">
				Table - 6 rows
			</div>
		</div>
	</div>
</div>
	<!-- Footer-->
	<footer class="footer">
		<span class="pull-right">
			Example text
		</span>
		Company 2015-2020
	</footer>

</div>