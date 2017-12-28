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
								<th width="20px"><i class="ra ra-sword"></i></th>
								<th>Name</th>
								<th>Slot</th>
								<th>Category</th>
							</tr>
							</thead>
							<tbody>
							{{range $key, $value := .Items}}
							<tr>
								<td><span title="{{$value.Name}}" class="item icon-{{$value.Icon}}-sm"></span></td>
								<td><a item={{$value.Id}} href="/item/{{$value.Id}}">{{$value.Name}}</a></td>
								<td><span title="{{$value.SlotList}}">{{$value.SlotsFirstName}}</span></td>
								<td><i title="{{$value.ItemtypeName}}" class="ra {{$value.ItemtypeIcon}}"></i></td>
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

<div id="eqitem"><h1>Abettor's Earring</h1>Magic, No Trade<br>Class: ALL<br>Race: ALL<br>Ear<br><br><table cellpadding="0" cellspacing="0" border="0"><tbody><tr><td colspan="2">Size:<span style="float:right; padding-left:8px; white-space:nowrap">TINY</span></td><td style="padding-right:8px;"></td><td style="padding-right:4px" nowrap="">HP:</td><td align="right">274</td><td class="hvalue" style="padding-right:8px;"></td></tr><tr><td style="padding-right:4px" nowrap="">Weight:</td><td align="right">0.3</td><td class="hvalue" style="padding-right:8px;"></td><td style="padding-right:4px" nowrap="">Mana:</td><td align="right">267</td><td class="hvalue" style="padding-right:8px;"></td></tr><tr><td style="padding-right:4px" nowrap="">Rec Level:</td><td align="right">75</td><td class="hvalue" style="padding-right:8px;"></td><td style="padding-right:4px" nowrap="">Endur:</td><td align="right">267</td><td class="hvalue" style="padding-right:8px;"></td></tr><tr><td style="padding-right:4px" nowrap="">Req Level:</td><td align="right">70</td><td class="hvalue" style="padding-right:8px;"></td></tr><tr><td style="height:4px;font-size:1px">&nbsp;</td></tr><tr><td style="padding-right:4px" nowrap="">Strength:</td><td align="right">19</td><td class="hvalue" style="padding-right:8px;">+1</td><td style="padding-right:4px" nowrap="">Magic:</td><td align="right">15</td><td class="hvalue" style="padding-right:8px;"></td><td style="padding-right:4px" nowrap="">Attack:</td><td align="right">12</td><td class="hvalue" style="padding-right:8px;"></td></tr><tr><td style="padding-right:4px" nowrap="">Stamina:</td><td align="right">14</td><td class="hvalue" style="padding-right:8px;">+1</td><td style="padding-right:4px" nowrap="">Fire:</td><td align="right">15</td><td class="hvalue" style="padding-right:8px;"></td><td style="padding-right:4px" nowrap="">HP Regen:</td><td align="right">3</td><td class="hvalue" style="padding-right:8px;"></td></tr><tr><td style="padding-right:4px" nowrap="">Intelligence:</td><td align="right">19</td><td class="hvalue" style="padding-right:8px;">+1</td><td style="padding-right:4px" nowrap="">Cold:</td><td align="right">15</td><td class="hvalue" style="padding-right:8px;"></td><td style="padding-right:4px" nowrap="">Mana Regen:</td><td align="right">2</td><td class="hvalue" style="padding-right:8px;"></td></tr><tr><td style="padding-right:4px" nowrap="">Wisdom:</td><td align="right">11</td><td class="hvalue" style="padding-right:8px;">+1</td><td style="padding-right:4px" nowrap="">Disease:</td><td align="right">15</td><td class="hvalue" style="padding-right:8px;"></td></tr><tr><td style="padding-right:4px" nowrap="">Agility:</td><td align="right">15</td><td class="hvalue" style="padding-right:8px;">+2</td><td style="padding-right:4px" nowrap="">Poison:</td><td align="right">15</td><td class="hvalue" style="padding-right:8px;"></td></tr><tr><td style="padding-right:4px" nowrap="">Dexterity:</td><td align="right">11</td><td class="hvalue" style="padding-right:8px;">+1</td></tr><tr><td style="padding-right:4px" nowrap="">Charisma:</td><td align="right">16</td><td class="hvalue" style="padding-right:8px;"></td></tr></tbody></table><div class="augments">Slot 1, type 3 (General: Spell Effect): empty<br>Slot 2, type 5 (Weapon: Elem Damage): empty<br>Slot 3, type 7 (General: Group): empty<br>Slot 4, type 9 (General: Dragons Points): empty<br></div><div class="effects">Effect: <a rel="eq:spell:9616" href="/spell/9616" target="_blank">Sharpshooting VII</a> (Worn)<br>Focus Effect: <a rel="eq:spell:42971" href="/spell/42971" target="_blank">Detrimental Duration 26 L110</a><br></div></div>

<script>
$(document).ready(function(){
	
    

	showTooltip();
	function showTooltip() {
		console.log("Initialized");
		$("[item]").each(function() {
			$(this).tooltip({
				track: true,
				open: function(event, ui) {
					console.log("Trigger");
					var id = this.id;
					var split_id = id.split('_');
					var userid = split_id[1];
					console.log(ui.attr("item"))
					$.ajax({
						url:'/api/item/'+ui.attr("item")+"/tooltip",
						type:'get',
						success: function(response){
							console.log(response);
							$(ui).tooltip('option','content',response.Data);     
						}
		  			});
				}
			});
		});
	}

	$("[item]").mouseout(function(event, ui){
		console.log($(this).attr("item"));
		// re-initializing tooltip
		$(this).attr('title','Please wait...');
		$(this).tooltip();
		$('.eqitem').hide();
	});
});
</script>