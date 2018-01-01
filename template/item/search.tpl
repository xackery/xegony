{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

	<div class="row">
		<div class="col-lg-12">

			<div class="hpanel">
				<div class="panel-body">
					<form method="GET" action="/item/search">
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
					{{if .Items}}
					<div class="table-responsive">
					<table cellpadding="1" cellspacing="1" class="table table-striped">

							<thead>
							<tr>
								<th width="20px"><i class="xa xa-sword"></i></th>
								<th>Name</th>
								<th>Slot</th>
								<th width="20px"><i title="Category" class="xa xa-anvil"></i></th>
							</tr>
							</thead>
							<tbody>
							{{range $key, $value := .Items}}
							<tr>
								<td><span class="item icon-{{$value.Icon}}-sm"></span></td>
								<td><a item={{$value.ID}} href="/item/{{$value.ID}}">{{$value.Name}}</a></td>
								<td><span title="{{$value.SlotList}}">{{$value.SlotsFirstName}}</span></td>
								<td><i title="{{$value.ItemtypeName}}" class="xa {{$value.ItemtypeIcon}}"></i></td></td>
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
