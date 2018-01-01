{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f"><a href="/item/byslot">Item By Slot</a> > {{.SlotID}}</span>
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th width="10px"><i class="xa xa-sword"></i></th>
                        <th>Name</th>
                        <th>Slot</th>
                        <th width="20px"><i title="Category" class="xa xa-anvil"></i></th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Items}}
                    <tr>
                        <td><span class="item icon-{{$value.Icon}}-sm"></span></td>
                        <td><a item={{$value.ID}} href="/item/{{$value.ID}}">{{$value.Name}}</a></td>
                        <td>{{$value.SlotsFirstName}}</td>
                        <td><i title="{{$value.ItemtypeName}}" class="xa {{$value.ItemtypeIcon}}"></i></td>
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