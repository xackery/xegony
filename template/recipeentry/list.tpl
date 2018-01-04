{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-4">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f"><a href="/npc">NPC</a> > {{.Recipe.Name}}</span>
			</div>
			<div class="panel-body">
				<span class="text-center">                    
					<h3> {{.Recipe.Name}}</h3>                
				</span>
			</div>           
			<div class="panel-body">
				<table class="table table-striped">
					<tbody>
					<tr><td>ID</td><td>{{.Recipe.ID}}</td></tr>
					<tr><td>Tradeskill</td><td>{{.Recipe.SkillName}}</td></tr>
					<tr><td>Trivial</td><td>{{.Recipe.Trivial}}</td></tr>
					<tr><td>Tools</td><td>{{.Recipe.ToolIconList}}</td></tr>
					<tr><td>Ingredients</td><td>{{.Recipe.ReagentIconList}}</td></tr>
					</tbody>
				</table>
			</div>
			<div class="panel-footer">
			</div>
		</div>
	</div>
	{{if .Recipe.Entrys}}
	<div class="col-lg-4">
		<div class="hpanel forum-box">
			<div class="panel-heading">
			   Entries
			</div>
			<div class="panel-body">				
				<div class="table-responsive">
				<table cellpadding="1" cellspacing="1" class="table table-striped">
						<thead>
						<tr>
							<th>Name</th>							
						</tr>
						</thead>
						<tbody>
						{{range $key, $value := .Recipe.Entrys}}
						<tr>
							<td><a href="/recipe/{{$value.RecipeID}}/{{$value.Item.ID}}"><span item="{{$value.Item.ID}}" class="item icon-{{$value.Item.Icon}}-sm"></span>{{$value.Item.Name}}</a></td>
						</tr>
						{{end}}                
						</tbody>
					
				</table>
				</div>
			</div>
			<div class="panel-footer">
				{{len .Recipe.Entrys}} recipe entries
			</div>
		</div>
	</div>
	{{end}}
	
</div>
</div>
</div>