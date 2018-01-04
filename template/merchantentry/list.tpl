{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
	<div class="col-lg-4">
		<div class="hpanel forum-box">
			<div class="panel-heading">
				<span class="f"><a href="/merchant">Merchant</a> > {{.Merchant.MerchantID}}</span>
			</div>
			<div class="panel-body">
				<span class="text-center">                    
					<h3> {{.Merchant.MerchantID}}</h3>                
				</span>
			</div>           
			<div class="panel-body">
				<table class="table table-striped">
					<tbody>
					<tr><td>ID</td><td>{{.Merchant.MerchantID}}</td></tr>
					</tbody>
				</table>
			</div>
			<div class="panel-footer">
			</div>
		</div>
	</div>
	{{if .Merchant.Entrys}}
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
							<th width="20px"><i class="xa xa-sword"></i></th>
							<th>Name</th>
							<th>Slot</th>
							<th><i title="Category" class="xa xa-vest"></i></th>
							<th width="20px"><span title="Probability">%</span></th>							
						</tr>
						</thead>
						<tbody>
						{{range $key, $value := .Merchant.Entrys}}
						<tr>
							<td><span class="item icon-{{$value.Item.Icon}}-sm"></span></td>
							<td><a item={{$value.Item.ID}} href="/item/{{$value.Item.ID}}">{{$value.Item.Name}}</a></td>
							<td><span title="{{$value.Item.SlotList}}">{{$value.Item.SlotsFirstName}}</span></td>
							<td><i title="{{$value.Item.ItemtypeName}}" class="xa {{$value.Item.ItemtypeIcon}}"></i></td>
							<td>{{$value.Probability}}</td>
						</tr>
						{{end}}                
						</tbody>
				</table>
				</div>
			</div>
			<div class="panel-footer">
				{{len .Merchant.Entrys}} merchant entries
			</div>
		</div>
	</div>
	{{end}}	
	{{if .Merchant.Npcs}}
    <div class="col-lg-4">
        <div class="hpanel">
            <div class="panel-heading hbuilt">
                <div class="panel-tools">
                    <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                    <a class="closebox"><i class="fa fa-times"></i></a>
                </div>
                This merchant is found:
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table table-striped">
                    <thead>
                    <tr>
                        <th width="10px"><i title="Race" class="xa xa-bear"></i></th>
                        <th width="10px"><i title="Class" class="xa xa-all-for-one"></i></th>
                        <th>Name</th>
                        <th>Zone</th>
                        
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Merchant.Npcs}}
                    <tr>
                        <td><i title="{{$value.RaceName}}" class="xa {{$value.RaceIcon}}"></i></td>
                                <td><i title="{{$value.ClassName}}" class="xa {{$value.ClassIcon}}"></i></td>
                        <td><a href="/npc/{{$value.ID}}">{{$value.Name}}</a></td>
                        <td><a href="/zone/{{$value.ZoneName}}">{{$value.ZoneName}}</a></td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>

            </div>
            <div class="panel-footer">
               {{len .Merchant.Npcs}} total merchants
            </div>
        </div>
    </div>
    {{end}}
	
</div>
</div>
</div>