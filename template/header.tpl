
<!--[if lt IE 7]>
<p class="alert alert-danger">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your experience.</p>
<![endif]-->

<!-- Header -->
<div id="header">
	<div class="color-line"></div>
	<div id="logo" class="light-version">
		<span>
			<img class="img-fluid" src="/images/logo.png" alt="{{.Site.Name}}" style="width: 20px"/>{{.Site.Name}}
		</span>
	</div>
	<nav role="navigation">
		<div class="small-logo">
			<span class="text-primary"><img class="img-fluid" src="/images/logo.png" alt="{{.Site.Name}}" style="width: 20px"/>{{.Site.Name}}</span>
		</div>

		<div class="navbar-left collapse navbar-collapse" id="left-nav">
		<ul class="nav navbar-nav no-borders">
			<li class="dropdown">
				<a class="dropdown-toggle label-menu-corner" href="#" data-toggle="dropdown">
					<i class="xa xa-monster-skull"></i>
				</a>
				<ul class="dropdown-menu hdropdown">						
					<li><a href="/character">Character List</a></li>
					<li><a href="/character/byonline">Who Is Online?</a></li>
					<li><a href="/character/byranking">Ranking</a></li>
				</ul>
			</li>


			<li class="dropdown">
				<a class="dropdown-toggle label-menu-corner" href="#" data-toggle="dropdown">
					<i class="xa xa-sword"></i>
				</a>
				<ul class="dropdown-menu hdropdown">
					<li><a href="/item">Item List</a></li>
					<li><a href="/spell">Spells</a></li>
				</ul>
			</li>

			<li class="dropdown">
				<a class="dropdown-toggle label-menu-corner" href="#" data-toggle="dropdown">
					<i class="xa xa-octopus"></i>
				</a>
				<ul class="dropdown-menu hdropdown">
					<li class="{{if eq .Site.Page "npcbyzone"}}active{{end}}"><a href="/npc/byzone"> <span class="nav-label">NPC By Zone</span></a></li>
					<li class="{{if eq .Site.Page "npcbyfaction"}}active{{end}}"><a href="/npc/byfaction"> <span class="nav-label">NPC By Faction</span></a></li>						
					<li class="{{if eq .Site.Page "npclist"}}active{{end}}"><a href="/npc"> <span class="nav-label">NPC List</span></a></li>
					<li class="{{if eq .Site.Page "npclist"}}active{{end}}"><a href="/merchant"> <span class="nav-label">Merchant List</span></a></li>
					<li class="{{if eq .Site.Page "npclist"}}active{{end}}"><a href="/spawn"> <span class="nav-label">Spawn List</span></a></li>					
					<li><a href="/loottable">Loot Tables</a></li>
				</ul>
			</li>

			<li class="dropdown">
				<a class="dropdown-toggle label-menu-corner" href="#" data-toggle="dropdown">
					<i class="xa xa-leaf"></i>
				</a>
				<ul class="dropdown-menu hdropdown">
					<li><a href="/fishing">Fishing</a></li>
					<li><a href="/forage">Forage</a></li>
					<li><a href="/recipe">Recipe List</a></li>
					<li><a href="/recipe/bytradeskill">Recipe By Tradeskill</a></li>
				</ul>
			</li>

			<li class="dropdown">
				<a class="dropdown-toggle label-menu-corner" href="#" data-toggle="dropdown">
					<i class="xa xa-scroll-unfurled"></i>
				</a>
				<ul class="dropdown-menu hdropdown flipInX">
					<li><a href="/task">Tasks</a></li>					
				</ul>
			</li>

			<li class="dropdown">
				<a class="dropdown-toggle label-menu-corner" href="#" data-toggle="dropdown">
					<i class="xa xa-palm-tree"></i>
				</a>
				<ul class="dropdown-menu hdropdown">
					<li><a href="/zone/bylevels">Leveling Chart</a></li>
					<li><a href="/zone">Zone List</a></li>
				</ul>
			</li>
		</ul>
		</div>

		<div class="mobile-menu">
			<button type="button" class="navbar-toggle mobile-menu-toggle" data-toggle="collapse" data-target="#mobile-collapse">
				<i class="fa fa-chevron-down"></i>
			</button>
			<button type="button" class="navbar-toggle mobile-menu-toggle" data-toggle="collapse" data-target="#left-nav" aria-expanded="false">							        
			</button>
			<div class="collapse mobile-navbar" id="mobile-collapse">
				<ul class="nav" id="top-menu">
					{{template "navmenu" .}}
				</ul>
			</div>
			
		</div>
		
		<div class="navbar-right">
			<ul class="nav navbar-nav no-borders">
				{{if .Site.User}}				
				<li class="dropdown">
					<a class="dropdown-toggle" href="#" data-toggle="dropdown">
						<i class="pe-7s-speaker"></i>
					</a>
					<ul class="dropdown-menu hdropdown notification flipInX">
						<li>
							<a>
								<span class="label label-success">NEW</span> It is a long established.
							</a>
						</li>
						<li>
							<a>
								<span class="label label-warning">WAR</span> There are many variations.
							</a>
						</li>
						<li>
							<a>
								<span class="label label-danger">ERR</span> Contrary to popular belief.
							</a>
						</li>
						<li class="summary"><a href="#">See all notifications</a></li>
					</ul>
				</li>
				
				<li class="dropdown">
					<a class="dropdown-toggle label-menu-corner" href="#" data-toggle="dropdown">
						<i class="pe-7s-mail"></i>
						<span class="label label-success">4</span>
					</a>
					<ul class="dropdown-menu hdropdown flipInX">
						<div class="title">
							You have 4 new messages
						</div>
						<li>
							<a>
								It is a long established.
							</a>
						</li>
						<li>
							<a>
								There are many variations.
							</a>
						</li>
						<li>
							<a>
								Lorem Ipsum is simply dummy.
							</a>
						</li>
						<li>
							<a>
								Contrary to popular belief.
							</a>
						</li>
						<li class="summary"><a href="#">See All Messages</a></li>
					</ul>
				</li>
				<li class="dropdown">
					<a class="dropdown-toggle label-menu-corner" href="#" data-toggle="dropdown">
						<i class="pe-7s-user"></i>
					</a>
					<ul class="dropdown-menu hdropdown">                        
						<li><a href="/dashboard">Dashboard</a></li>
						<li><a href="/item/search/byaccount">Inventory Search</a></li>
						{{if .Site.User.IsAdmin}}
						<li><a href="/variable"><i class="xa xa-three-keys"></i> Variable</a></li>
						<li><a href="/rule"><i class="xa xa-wireless-signal"></i> Rule</a></li>
						<li><a href="/error"><i class="xa xa-fast-ship"></i> Error</a></li>
						<li><a href="/hacker"><i class="xa xa-radioactive"></i> Hacker</a></li>
						{{end}}
						<li><a href="/logout">Logout</a></li>						
					</ul>
				</li>
	   
				{{else}}
				<li>
					<a href="/login">
						<i class="pe-7s-key"></i>
					</a>
				</li>
				{{end}}
			</ul>

		</div>
	</nav>
</div>