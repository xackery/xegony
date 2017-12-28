{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="normalheader small-header">
    <div class="hpanel">
        <div class="panel-body">
           

            <div id="hbreadcrumb" class="pull-right m-t-lg">
                <ol class="hbreadcrumb breadcrumb">
                    <li class="active">
                        <span>Bestiary</span>
                    </li>
                </ol>
            </div>
            <h2 class="font-light m-b-xs">
               Bestiary List
            </h2>
            <small>There are 215 characters on the server.</small>
        </div>
    </div>
</div>

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel">
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table">
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
                        <td><span class="slot-sm"><span class="item icon-{{$value.Icon}}-sm"></span></span></td>
                        <td><a item={{$value.Id}} href="/npc/{{$value.Id}}">{{$value.Name}}</a></td>
                        <td>{{$value.SlotsFirstName}}</td>
                        <td>{{$value.ItemtypeName}}</td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>

            </div>
            <div class="panel-footer">
                Table - 6 rows
            </div>
        </div>
    </div>
</div>
</div>