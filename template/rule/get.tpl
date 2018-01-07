{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

<div class="row">
	<div class="col-lg-6">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f"><a href="/rule">Rule</a> > {{.Rule.Name}}</span>
			</div>

			<div class="panel-body">     	
				<table class="table table-striped">
					<tbody>
						<tr><td>Name</td><td>{{.Rule.Name}}</td></tr>
						<tr><td>Value</td><td>{{.Rule.ValueParse}}</td></tr>
						<tr><td>Raw Value</td><td>{{.Rule.Value}}</td></tr>
						<tr><td>Information</td><td>{{.Rule.Notes.String}}</td></tr>
					</tbody>
				</table>
			</div>           
		</div>
	</div>
</div>

</div>
</div>