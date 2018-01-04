{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Merchants</span>
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table id="items" data-paging="true" data-sorting="false" data-filtering="false" cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th>ID</th>
                        <th>Items</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Merchants}}
                    <tr>
                        <td><a href="/merchant/{{$value.MerchantID}}">{{$value.MerchantID}}</a></td>
                        <td>{{$value.ItemIconList}}</td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>
                {{.Site.PageList}}
            </div>
            <div class="panel-footer">
                {{.Site.ResultCount}} total merchants
            </div>
        </div>
    </div>
</div>
</div>

