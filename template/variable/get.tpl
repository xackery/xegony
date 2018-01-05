{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

<div class="row">
	<div class="col-lg-6">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f"><a href="/variable">Variable</a> > {{.Variable.Name}}</span>
			</div>

			<div class="panel-body">     	
				<table class="table table-striped">
					<tbody>
						<tr><td>Name</td><td>{{.Variable.Name}}</td></tr>
						<tr><td>Value</td><td>{{.Variable.ValueParse}}</td></tr>
						<tr><td>Raw Value</td><td>{{.Variable.Value}}</td></tr>
						<tr><td>Information</td><td>{{.Variable.Information}}</td></tr>
					</tbody>
				</table>
			</div>           
		</div>
	</div>
</div>

</div>
</div>