{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Forages</span>
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table id="forages" data-paging="true" data-sorting="false" data-filtering="false" cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th width="20px"><i class="xa xa-sword"></i></th>
                        <th>Name</th>
                        <th>Price</th>
                        <th>Chance</th>
                        <th width="20px"><i title="Category" class="xa xa-anvil"></i></th>
                        <th>Zone</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Forages}}
                    <tr>
                        <td><span class="item icon-{{$value.Item.Icon}}-sm"></span></td>
                        <td><a item={{$value.ID}} href="/forage/{{$value.ID}}">{{$value.Item.Name}}</a></td>
                        <td>{{$value.Item.PriceName}}</td>
                        <td>{{$value.Chance}}</td>
                        <td><i title="{{$value.Item.ItemtypeName}}" class="xa {{$value.Item.ItemtypeIcon}}"></i></td>
                        <td>{{if $value.Zone}}{{$value.Zone.ShortName.String}}{{end}}</td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>
                {{.ForagePage.PageList}}
            </div>
            <div class="panel-footer">
                {{.ForagePage.Total}} total forage items
            </div>
        </div>
    </div>
</div>
</div>

