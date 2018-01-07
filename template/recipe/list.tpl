{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Recipes</span>
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table id="items" data-paging="true" data-sorting="false" data-filtering="false" cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th width="20px"><i class="xa xa-lightning-bolt"></i></th>                        
                        <th>Name</th>
                        <th title="Trivial Level" width="20px">Lvl</th>
                        <th>Tradeskill</th>                        
                        <th>Reagents</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Recipes}}
                    <tr>
                        <td><span item="{{$value.RewardItem.ID}}" class="item icon-{{$value.RewardItem.Icon}}-sm"></span></td>
                        <td><a href="/recipe/{{$value.ID}}">{{$value.Name}}</a></td>
                        <td>{{$value.Trivial}}</td>
                        <td>{{$value.SkillName}}</td>
                        <td>{{$value.ReagentIconList}}</td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>
                {{.RecipePage.PageList}}
            </div>
            <div class="panel-footer">
                {{.RecipePage.Total}} total spells
            </div>
        </div>
    </div>
</div>
</div>
