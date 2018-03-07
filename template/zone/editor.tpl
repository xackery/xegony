{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-8">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f">{{.Zone.LongName}} Editor</span>
			</div>

			<div class="panel-body">
				<div class="col-md-8">
					<div id="map" style="height: 512px"></div>
				</div>
			</div>
		</div>
	</div>
	<div class="col-lg-4">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f">Info</span>
			</div>

			<div class="panel-body">
				<table class="table table-striped">
					<tbody>
					<tr><td>Spawn Name:</td><td id="infoName">Orc</td></tr>
					<tr><td>ID:</td><td id="infoID"></td></tr>	
					<tr><td>Limit:</td><td id="infoLimit"></td></tr>
					<tr><td>Distance:</td><td id="infoDistance"></td></tr>
					<tr><td>Minimum Delay:</td><td id="infoMinimumDelay"></td></tr>
					<tr><td>Despawn:</td><td id="infoDespawn"></td></tr>
					<tr><td>Despawn Timer:</td><td id="infoDespawnTimer"></td></tr>
					</tbody>
				</table>
				<div class="table-responsive">
                	<table cellpadding="1" cellspacing="1" class="table">
						<thead>
						<tr>
							<th width="10px"><i title="Race" class="xa xa-bear"></i></th>
							<th width="10px"><i title="Class" class="xa xa-all-for-one"></i></th>
							<th width="10px"><i title="Level">Lvl</i></th>
							<th>Name</th>
							<th>Chance</th>                   
						</tr>
						</thead>
						<tbody id="infoNpcs">						
						</tobdy>
					</table>
				</div>
			</div>
		</div>
	</div>
</div>

</div>

<script>
	
	var map = L.map('map', {
		crs: L.CRS.Simple
	});

	var bounds = [[0,0], [500,500]];
	var image = L.imageOverlay('/images/maps/everfrost.png', bounds).addTo(map);
	map.fitBounds(bounds);

	
	var LeafIcon = L.Icon.extend({
		options: {
			//shadowUrl: 'leaf-shadow.png',
			iconSize:     [30, 30],
			shadowSize:   [25, 25],
			iconAnchor:   [30, 30],
			shadowAnchor: [4, 62],
			popupAnchor:  [-3, -76]
		}
	});

	var modelIcon = new LeafIcon({iconUrl: '/images/npc/icon/54.png'})		

	{{range $spawnKey, $spawnValue := .Spawns}}
		{{range $key, $value := $spawnValue.Entrys}}
			function onMapClick{{ unescapeJS $spawnValue.ID }}(e) {
				$('#infoName').html('{{ unescapeJS $spawnValue.Name}}');
				$('#infoID').html('{{ unescapeJS $spawnValue.ID}}');
				$('#infoName').html('{{ unescapeJS $spawnValue.Name }}');
				$('#infoID').html('{{ unescapeJS $spawnValue.ID }}');
				$('#infoLimit').html('{{ unescapeJS $spawnValue.Limit }}');
				$('#infoDistance').html('{{ unescapeJS $spawnValue.Distance }}');
				$('#infoMinimumDelay').html('{{ unescapeJS $spawnValue.MinimumDelay }}');
				$('#infoDespawn').html('{{ unescapeJS $spawnValue.Despawn }}');
				$('#infoDespawnTimer').html('{{ unescapeJS $spawnValue.DespawnTimer }}');
				$('#infoNpcs').html('{{ unescapeJS $spawnValue.Npcs }}');
				var infoNpcs = ""
				{{range $npcKey, $npcValue := $spawnValue.Npcs}}
				infoNpcs += "<tr>\n";
				infoNpcs += '<td><i class="{{ unescapeJS $npcValue.Npc.Race.Icon }}"></i></td>';
				infoNpcs += '<td><i class="{{ unescapeJS $npcValue.Npc.Class.Icon }}"></i></td>';
				infoNpcs += "<td>{{ unescapeJS $npcValue.Npc.Level }}</td>";
				infoNpcs += "<td>{{ unescapeJS $npcValue.Npc.Name }}</td>";
				infoNpcs += "<td>{{ unescapeJS $npcValue.Chance }}</td>";
				$('#infoNpcs').html(infoNpcs);				
				{{end}}
			}
			var marker = L.marker([{{$value.YScaled}},{{$value.XScaled}}], {icon: modelIcon, spawnID:{{$spawnValue.ID}}});
			marker.bindPopup({{$spawnValue.Name}}+" "+{{$spawnValue.ID}});
			marker.on('click', onMapClick{{ unescapeJS $spawnValue.ID}});
			marker.addTo(map);
		{{end}}
	{{end}}
</script>