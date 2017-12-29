<!-- Navigation -->
<aside id="menu">
	<div id="navigation">		
		<ul class="nav" id="side-menu">			
			<li class="{{if eq .Site.Page "forum"}}active{{end}}"><a href="/"> <span class="nav-label"><i class="xa xa-coffee-mug"></i> Forum</span></a></li>

			
			<li class="{{if eq .Site.Section "character"}}active{{end}}">
				<a href="#" aira-expanded="{{if eq .Site.Section "character"}}true{{else}}false{{end}}"> 
					<span class="nav-label"><i class="xa xa-monster-skull"></i> Character</span>
					<span class="fa arrow"/>
				</a>				
				<ul class="nav nav-second-level collapse {{if eq .Site.Section "character"}}in{{end}}" aria-expanded="{{if eq .Site.Section "character"}}true{{else}}false{{end}}">
					<li class="{{if eq .Site.Page "characterbyonline"}}active{{end}}"><a href="/character/byonline"> <span class="nav-label">Online</span></a></li>					
					<li class="{{if eq .Site.Page "characterbyranking"}}active{{end}}"><a href="/character/byranking"> <span class="nav-label">Ranking</span></a></li>					
					<li class="{{if eq .Site.Page "charactersearch"}}active{{end}}"><a href="/character/search"> <span class="nav-label">Search</span></a></li>
				</ul>
			</li>	

			<li class="{{if eq .Site.Section "npc"}}active{{end}}">
				<a href="#" aira-expanded="{{if eq .Site.Section "npc"}}true{{else}}false{{end}}"> 
					<span class="nav-label"><i class="xa xa-octopus"></i> Bestiary</span>
					<span class="fa arrow"/>
				</a>
				<ul class="nav nav-second-level collapse {{if eq .Site.Section "npc"}}in{{end}}" aria-expanded="{{if eq .Site.Section "npc"}}true{{else}}false{{end}}">
					<li class="{{if eq .Site.Page "npcbyzone"}}active{{end}}"><a href="/npc/byzone"> <span class="nav-label">By Zone</span></a></li>
					<li class="{{if eq .Site.Page "npcbyfaction"}}active{{end}}"><a href="/npc/byfaction"> <span class="nav-label">By Faction</span></a></li>
					<li class="{{if eq .Site.Page "npcsearch"}}active{{end}}"><a href="/npc/search"> <span class="nav-label">Search</span></a></li>
					<li class="{{if eq .Site.Page "npclist"}}active{{end}}"><a href="/npc"> <span class="nav-label">List</span></a></li>
				</ul>
			</li>


			<li class="{{if eq .Site.Section "zone"}}active{{end}}">
				<a href="#" aira-expanded="{{if eq .Site.Section "zone"}}true{{else}}false{{end}}"> 
					<span class="nav-label"><i class="xa xa-palm-tree"></i> Zone</span>
					<span class="fa arrow"/>
				</a>
				<ul class="nav nav-second-level collapse {{if eq .Site.Section "zone"}}in{{end}}" aria-expanded="{{if eq .Site.Section "zone"}}true{{else}}false{{end}}">					
					<li class="{{if eq .Site.Page "zonebyexpansion"}}active{{end}}"><a href="/zone/byexpansion"> <span class="nav-label">By Expansion</span></a></li>
					<li class="{{if eq .Site.Page "zonelevels"}}active{{end}}"><a href="/zone/levels"> <span class="nav-label">Levels</span></a></li>
					<li class="{{if eq .Site.Page "zonehotzone"}}active{{end}}"><a href="/zone/byhotzone"> <span class="nav-label">Hot Zone</span></a></li>
					<li class="{{if eq .Site.Page "listzone"}}active{{end}}"><a href="/zone"> <span class="nav-label">List</span></a></li>
				</ul>
			</li>

			<li class="{{if eq .Site.Section "item"}}active{{end}}">
				<a href="#" aira-expanded="{{if eq .Site.Section "item"}}true{{else}}false{{end}}"> 
					<span class="nav-label"><i class="xa xa-sword"></i> Item</span>
					<span class="fa arrow"/>
				</a>
				<ul class="nav nav-second-level collapse {{if eq .Site.Section "item"}}in{{end}}" aria-expanded="{{if eq .Site.Section "item"}}true{{else}}false{{end}}">
					<li class="{{if eq .Site.Page "itemsearch"}}active{{end}}"><a href="/item/search"> <span class="nav-label">Search</span></a></li>
					<li class="{{if eq .Site.Page "itembyslot"}}active{{end}}"><a href="/item/byslot"> <span class="nav-label">By Slot</span></a></li>
					<li class="{{if eq .Site.Page "itembyzone"}}active{{end}}"><a href="/item/byzone"> <span class="nav-label">By Zone</span></a></li>					
					<li class="{{if eq .Site.Page "itemlist"}}active{{end}}"><a href="/item"> <span class="nav-label">List</span></a></li>
				</ul>
			</li>

			<li class="{{if eq .Site.Section "task"}}active{{end}}">
				<a href="#" aira-expanded="{{if eq .Site.Section "task"}}true{{else}}false{{end}}"> 
					<span class="nav-label"><i class="xa xa-wooden-sign"></i> Task</span>
					<span class="fa arrow"/>
				</a>
				<ul class="nav nav-second-level collapse {{if eq .Site.Section "task"}}in{{end}}" aria-expanded="{{if eq .Site.Section "task"}}true{{else}}false{{end}}">
					<li class="{{if eq .Site.Page "tasklist"}}active{{end}}"><a href="/task"> <span class="nav-label">List</span></a></li>
					<li class="{{if eq .Site.Page "tasksearch"}}active{{end}}"><a href="/task/search"> <span class="nav-label">Search</span></a></li>					
				</ul>
			</li>

			<li class="{{if eq .Site.Section "loottable"}}active{{end}}">
				<a href="#" aira-expanded="{{if eq .Site.Section "loottable"}}true{{else}}false{{end}}"> 
					<span class="nav-label"><i class="fa fa-shopping-basket"></i> Loot</span>
					<span class="fa arrow"/>
				</a>
				<ul class="nav nav-second-level collapse {{if eq .Site.Section "loottable"}}in{{end}}" aria-expanded="{{if eq .Site.Section "loottable"}}true{{else}}false{{end}}">					
					<li class="{{if eq .Site.Page "loottablelist"}}active{{end}}"><a href="/loottable"> <span class="nav-label">Table List</span></a></li>
				</ul>
			</li>
			{{if .Site.User}}
			<div class="profile-picture">
	            
	            <div class="stats-label text-color">
	             <a href="/dashboard">   <span class="font-extra-bold font-uppercase">{{.Site.User.Name}}</span></a>

	                <div id="expLine" class="small-chart m-t-sm"></div>
	                <div>
	                    <h4 class="font-extra-bold m-b-xs">
	                        1,123,123,123
	                    </h4>
	                    <small class="text-muted">Experience last 24 hours</small>
	                </div>
	            </div>
	        </div>
	        {{end}}
		</ul>
	</div>
</aside>

{{if .Site.User}}
<script src="/vendor/sparkline/index.js"></script>
<script type="text/javascript">
    $("#expLine").sparkline([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24], {
        type: 'bar',
        barWidth: 4,
        height: '30px',
        barColor: '#62cb31',
        negBarColor: '#53ac2a'
    });
</script>
{{end}}