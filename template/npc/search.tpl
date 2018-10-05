
<article class="content">
	<div class="title-search-block">
		<div class="title-block">
			<div class="row">
				<div class="col-md-6">
					<h3 class="title"> {{.Site.Page}} <a href="item-editor.html" class="btn btn-primary btn-sm rounded-s"> Add New </a>
						<div class="action dropdown">
							<button class="btn  btn-sm rounded-s btn-secondary dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> More actions... </button>
							<div class="dropdown-menu" aria-labelledby="dropdownMenu1">
								<a class="dropdown-item" href="#"><i class="fa fa-pencil-square-o icon"></i>Mark as a draft</a>
								<a class="dropdown-item" href="#" data-toggle="modal" data-target="#confirm-modal"><i class="fa fa-close icon"></i>Delete</a>
							</div>
						</div>
					</h3>
					<p class="title-description"> {{.Site.PageSummary}} </p>
				</div>
			</div>
		</div>
		<div class="items-search">
			<form class="form-inline">
				<div class="input-group">
					<input type="text" class="form-control boxed rounded-s" placeholder="Search for...">
					<span class="input-group-btn">
						<button class="btn btn-secondary rounded-s" type="button">
							<i class="fa fa-search"></i>
						</button>
					</span>
				</div>
			</form>
		</div>
	</div>
	<div class="card items">
		<ul class="item-list striped">
			<li class="item item-list-header">
				<div class="item-row">
					{{/*<div class="item-col fixed item-col-check">
						<label class="item-check" id="select-all-items">
							<input type="checkbox" class="checkbox">
							<span></span>
						</label>
					</div>*/}}
					<div class="item-col item-col-header fixed item-col-img md"><div><span>Icon</span></div></div>
					<div class="item-col item-col-header item-col-title"><div><span>Name</span></div></div>
					<div class="item-col item-col-header item-col-sales"><div><span>Level</span></div></div>
					<div class="item-col item-col-header item-col-sales"><div><span>Zone</span></div></div>
				</div>
			</li>
			{{range $key, $npc := .Resp.Npcs}}
			<li class="item">
				<div class="item-row">
					{{/*<div class="item-col fixed item-col-check">
						<label class="item-check" id="select-all-items">
							<input type="checkbox" class="checkbox">
							<span></span>
						</label>
					</div>*/}}
					<div class="item-col fixed item-col-img md"><a href="/npc?id={{$npc.ID}}"><div class="item-img rounded" style="background-image: url(https://s3.amazonaws.com/uifaces/faces/twitter/brad_frost/128.jpg)"></div></a></div>
					<div class="item-col item-col-title"><div class="item-heading">Name</div><div><a href="/npc?id={{$npc.ID}}" class=""><h4 class="item-title">{{$npc.CleanName}}</h4></a></div></div>
					<div class="item-col item-col-sales"><div class="item-heading">Level</div><div>{{$npc.Level}}</div></div>
					<div class="item-col item-col-stats no-overflow"><div class="item-heading">Zone</div><div class="no-overflow"><a href="">The Overthere</a></div></div>
				</div>
			</li>
			{{end}}
		</ul>
	</div>
	<nav class="text-right">
		<ul class="pagination">
			<li class="page-item">{{if eq .Resp.Offset 0}}{{else}}<a class="page-link" href="{{.Req.Path}}&offset=0">{{end}} Prev {{if eq .Resp.Offset 0}}{{else}}</a>{{end}}</li>
			<li class="page-item active">
				<a class="page-link" href="{{.Path}}&offset={{add .Resp.Offset 1}}"> {{add .Resp.Offset 1}} </a>
			</li>
			<li class="page-item">
				<a class="page-link" href=""> 2 </a>
			</li>
			<li class="page-item">
				<a class="page-link" href=""> 3 </a>
			</li>
			<li class="page-item">
				<a class="page-link" href=""> 4 </a>
			</li>
			<li class="page-item">
				<a class="page-link" href=""> 5 </a>
			</li>
			<li class="page-item">{{if eq .Resp.Offset 0}}{{else}}<a class="page-link" href="{{.Req.Path}}&offset=0">{{end}} Next {{if eq .Resp.Offset 0}}{{else}}</a>{{end}}</li>
		</ul>
	</nav>
</article>