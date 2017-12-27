{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f"><a href="/item/byslot">Item By Slot</a> > {{.SlotId}}</span>
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table">
                    <thead>
                    <tr>
                        <th>Name</th>
                        <th>CharacterID</th>                   
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Items}}
                    <tr>
                        
                        <td><a href="/item/{{$value.Id}}">{{$value.Name}}</a></td>
                        <td>{{$value.Id}}</td>
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
    <!-- Footer-->
    <footer class="footer">
        <span class="pull-right">
            Example text
        </span>
        Company 2015-2020
    </footer>

</div>