{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

	<div class="row">
		<div class="col-lg-12">

			<div class="hpanel">
				<div class="panel-body">
					<form method="GET" action="/spell/search">
						<div class="text-muted small pull-right text-right">            
						</div>
						<div class="form-inline">
							Search: <input class="form-control" name="search" {{if .Search}}value="{{.Search}}"{{end}} type="text">
						</div>
					</form>

				</div>

			</div>

			{{if .Search}}
			<div class="hpanel forum-box">

			<div class="panel-heading">
				<span class="f">Search results for '{{.Search}}'</span>
			</div>


				<div class="panel-body">
					{{if .Spells}}
					<div class="table-responsive">
					<table cellpadding="1" cellspacing="1" class="table table-striped">

							<thead>
							<tr>
								<th width="20px"><i class="xa xa-lightning-bolt"></i></th>
		                        <th width="20px">Lvl</th>
		                        <th>Name</th>
		                        <th>Description</th>
		                        <th>Skill</th>
							</tr>
							</thead>
							<tbody>
							{{range $key, $value := .Spells}}
		                    <tr>
		                        <td><span class="item icon-{{$value.Icon}}-sm"></span></td>
		                        <td>{{$value.LowestLevel}}</td>
		                        <td><a item={{$value.ID}} href="/spell/{{$value.ID}}">{{$value.Name.String}}</a></td>
		                        <td>{{$value.DescriptionName}}</td>
		                        <td>{{$value.SkillName}}</td>
		                    </tr>
		                    {{end}}                    
							</tbody>
						
					</table>
					</div>
					{{else}}
					{{if .Search}}
					No results for {{.Search}} were found
					{{end}}
					{{end}}

				</div>
				<div class="panel-footer">
					Table - 6 rows
				</div>
			</div>
			{{end}}

		</div>
	</div>
</div>
