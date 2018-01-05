{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Variables</span>
            </div>

            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>                        
                        <th>Name</th>   
                        <th>Value</th>
                        <th>Raw Value</th>
                        <th>Information</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Variables}}
                    <tr>
                        <td><a href="/variable/{{$value.Name}}">{{$value.Name}}</a></td>
                        <td><a href="/variable/{{$value.Name}}">{{$value.ValueParse}}</a></td>
                        <td><a href="/variable/{{$value.Name}}">{{$value.Value}}</a></td>
                        <td><a href="/variable/{{$value.Name}}">{{$value.Information}}</a></td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>

            </div>
            <div class="panel-footer">
                {{len .Variables}} total variables
            </div>
        </div>
    </div>
</div>
    <!-- Footer-->
    <footer class="footer">
        <span class="pull-right">
            Example text
        </span>
        Company 2015-2020
    </footer>

</div>