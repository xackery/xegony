{{template "header" .}}

{{template "navigation" .}}

<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-4">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f"><a href="/npc">NPC</a> > {{.Npc.CleanName}}</span>
            </div>
            <div class="panel-body">
                <img alt="logo" class="img-circle m-b m-t-md" src="/images/profile.jpg">
                <h3>{{.Npc.CleanName}}</h3>
                <div class="text-muted font-bold m-b-xs"></div>
                <p>
                    {{.Npc.Id}}
                </p>
                <div class="progress m-t-xs full progress-small">
                    <div style="width: 65%" aria-valuemax="100" aria-valuemin="0" aria-valuenow="65" role="progressbar" class=" progress-bar progress-bar-success">
                        <span class="sr-only">35% Complete (success)</span>
                    </div>
                </div>
            </div>           
            <div class="panel-body">
                <dl>
                    <dt>Description lists</dt>
                    <dd>A description list is perfect for defining terms.</dd>
                    <dt>Euismod</dt>
                    <dd>Vestibulum id ligula porta felis euismod semper eget lacinia odio sem nec elit.</dd>
                    <dd>Donec id elit non mi porta gravida at eget metus.</dd>
                    <dt>Malesuada porta</dt>
                    <dd>Etiam porta sem malesuada magna mollis euismod.</dd>
                </dl>
            </div>

        </div>
    </div>
    <div class="col-lg-8">
        <div id="row">
            <div class="hpanel">
                <div class="panel-heading hbuilt">
                    <div class="panel-tools">
                        <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                        <a class="closebox"><i class="fa fa-times"></i></a>
                    </div>
                    Spawns In
                </div>
                <div class="panel-body">
                    <div class="table-responsive">
                    <table cellpadding="1" cellspacing="1" class="table">
                        <thead>
                        <tr>
                            <th>ZoneID</th>                        
                        </tr>
                        </thead>
                        <tbody>
                        {{range $key, $value := .Npc.SpawnsIn}}
                        <tr>
                            <td><a href="/zone/{{$value.ZoneId}}">{{$value.ZoneShortName}}</a></td>         
                        </tr>
                        {{end}}
                        </tbody>
                    </table>
                    </div>

                </div>
            </div>
            <div class="hpanel">
                <div class="panel-heading hbuilt">
                    <div class="panel-tools">
                        <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                        <a class="closebox"><i class="fa fa-times"></i></a>
                    </div>
                   Drops
                </div>
                <div class="panel-body">
                    <div class="table-responsive">
                    <table cellpadding="1" cellspacing="1" class="table">
                        <thead>
                        <tr>
                            <th>Name</th>
                            <th>Charges</th>
                            <th>Chance</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{range $key, $value := .Npc.Drops.LootDrops}}
                        <tr>
                            <td class="text-left"><a href="/item/{{$value.ItemId}}"><span class="item icon-{{$value.ItemIcon}}"></span></a>
                            <a href="/item/{{$value.ItemId}}">{{$value.ItemName}}</a></td>
                            <td><a href="/item/{{$value.ItemId}}">{{$value.ItemCharges}}</a></td>
                            <td><a href="/item/{{$value.ItemId}}">{{$value.Chance}}</a></td>
                        </tr>
                        {{end}}
                        </tbody>
                    </table>
                    </div>

                </div>
            </div>
        </div>
    </div>
    </div>
</div>
</div>