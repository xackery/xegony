{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

<div class="row">
	<div class="col-lg-6">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f"><a href="/zone">Zone</a> > {{.Zone.LongName}}</span>
			</div>

			<div class="panel-body">     	
				<p>		
					Zone Description Goes Here.
				</p>
				<table class="table table-striped">
					<tbody>
						<tr><td>Expansion</td><td>{{.Zone.ExpansionName}}</td></tr>
						<tr><td>Short Name</td><td>{{.Zone.ShortName.String}}</td></tr>
						<tr><td>Safe Point</td><td>{{.Zone.SafeX}}, {{.Zone.SafeY}}, {{.Zone.SafeZ}}</td></tr>
						{{if .Zone.MinLevel}}<tr><td>Minimum Level</td><td>{{.Zone.MinLevel}}</td></tr>{{end}}
						<tr><td>Exp Multiplier</td><td>{{.Zone.ZoneExpMultiplier}}</td></tr>
						{{if .Zone.Canbind}}<tr><td>Can Bind Here?</td><td>Yes</td></tr>{{end}}
						{{if ne .Zone.Cancombat 1}}<tr><td>Is Combat Enabled?</td><td>No</td></tr>{{end}}
						{{if ne .Zone.Canlevitate 1}}<tr><td>Is Levitation Enabled?</td><td>No</td></tr>{{end}}
						<tr><td>Is Outdoors?</td><td>{{if .Zone.Castoutdoor}}Yes{{else}}No{{end}}</td></tr>
						{{if .Zone.Hotzone}}<tr><td>Is Hot Zone</td><td>Yes</tr>{{end}}
						{{if .Zone.Suspendbuffs}}<tr><td>Are Buffs Suspended?</td><td>Yes</tr>{{end}}
					</tbody>
				</table>
			</div>           
		</div>
	</div>
</div>

</div>
</div>