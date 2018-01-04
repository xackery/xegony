{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Spells</span>
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table id="items" data-paging="true" data-sorting="false" data-filtering="false" cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th width="20px"><i class="xa xa-lightning-bolt"></i></th>
                        <th width="20px">Lvl</th>
                        <th>Name</th>
                        <th>Description</th>
                        <th>Skill</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Spells}}
                    <tr>
                        <td><span class="item icon-{{$value.Icon}}-sm"></span></td>
                        <td>{{$value.LowestLevel}}</td>
                        <td><a item={{$value.ID}} href="/spell/{{$value.ID}}">{{$value.Name.String}}</a></td>
                        <td>{{$value.DescriptionName}}</td>
                        <td>{{$value.SkillName}}</td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>
                {{.Site.PageList}}
            </div>
            <div class="panel-footer">
                {{.Site.ResultCount}} total spells
            </div>
        </div>
    </div>
</div>
</div>

