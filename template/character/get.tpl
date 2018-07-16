{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-4">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f"><a href="/character">Character</a> > {{.Character.Name}}</span>
            </div>
            <div class="panel-body">
                <div class="pull-right text-right">
                    <div class="btn-group">
                        <i class="fa fa-facebook btn btn-default btn-xs"></i>
                        <i class="fa fa-twitter btn btn-default btn-xs"></i>
                        <i class="fa fa-linkedin btn btn-default btn-xs"></i>
                    </div>
                </div>
                <img alt="logo" class="img-circle m-b m-t-md" src="/images/profile.jpg">
                <h3><a href="/profile.html">{{.Character.Title}} {{.Character.Name}} {{.Character.LastName}}</a></h3>
                <div class="text-muted font-bold m-b-xs">{{.Character.ZoneID}}</div>
                <p>
                   Level {{.Character.Level}} <i class="xa {{.Character.RaceIcon}}"></i> {{.Character.RaceName}} <i class="xa {{.Character.ClassIcon}}"></i> {{.Character.ClassName}} 
                </p>
                <div class="progress m-t-xs full progress-small">
                    <div style="width: 65%" aria-valuemax="100" aria-valuemin="0" aria-valuenow="65" role="progressbar" class=" progress-bar progress-bar-success">
                        <span class="sr-only">35% Complete (success)</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</div>
</div>
