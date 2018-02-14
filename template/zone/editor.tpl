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
				<div class="col-md-4">
					<table class="table table-striped">
					<tbody>
					<tr><td>Name:</td><td>Orc</td></tr>
					<tr><td>Name:</td><td>Orc</td></tr>
					</tbody>
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

	function onMapClick(e) {
		console.log("Orc click");
		console.log(e);
	}

	L.marker([250, 250], {icon: modelIcon}).on('click', onMapClick).addTo(map);
	
	{{range $spawnKey, $spawnValue := .Spawns}}
		{{range $key, $value := $spawnValue.Entrys}}
			L.marker([{{$value.X}}, {{$value.Y}}], {icon: modelIcon}).on('click', onMapClick).addTo(map); //{{range $npcKey, $npcValue := $spawnValue.Npcs}}{{$npcValue.Npc.ID}} {{end}}
		{{end}}
	{{end}}
</script>