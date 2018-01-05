{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Hackers</span>
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table id="errors" data-paging="true" data-sorting="false" data-filtering="false" cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th>AccountID</th>
                        <th>Name</th>
                        <th>Message</th>
                        <th>Zone</th>
                        <th>Date</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Hackers}}
                    <tr>
                        <td>{{$value.AccountID}}</td>
                        <td>{{$value.Name}}</td>
                        <td>{{$value.Hacked}}</td>
                        <td>{{$value.ZoneID}}</td>
                        <td>{{$value.Date}}</td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>
                {{.HackerPage.PageList}}
            </div>
            <div class="panel-footer">
                {{.HackerPage.Total}} total hack entries
            </div>
        </div>
    </div>
</div>
</div>

