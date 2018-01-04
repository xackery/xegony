{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">
<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Errors</span>
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table id="errors" data-paging="true" data-sorting="false" data-filtering="false" cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th>Scope</th>
                        <th>Message</th>
                        <th>Date</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Errors}}
                    <tr>
                        <td>{{$value.Scope}}
                        <td>{{$value.Message}}</td>
                        <td>{{$value.CreateDate}}</td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>
                {{.ErrorPage.PageList}}
            </div>
            <div class="panel-footer">
                {{.ErrorPage.Total}} total errors
            </div>
        </div>
    </div>
</div>
</div>

