{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-12">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f">Create Forum</span>
            </div>

            <div class="panel-body">
                <form method="get" class="form-horizontal">
                <div class="form-group"><label class="col-sm-2 control-label">Normal</label>
                    <div class="col-sm-10"><input type="text" class="form-control"></div>
                </div>
                <div class="hr-line-dashed"></div>
                <div class="form-group"><label class="col-sm-2 control-label">Help text</label>

                    <div class="col-sm-10"><input type="text" class="form-control"> <span class="help-block m-b-none">A block of help text that breaks onto a new line and may extend beyond one line.</span>
                    </div>
                </div>
                <div class="hr-line-dashed"></div>
                <div class="form-group"><label class="col-sm-2 control-label">Password</label>

                    <div class="col-sm-10"><input type="password" class="form-control" name="password"></div>
                </div>
                <div class="hr-line-dashed"></div>
                <div class="form-group"><label class="col-sm-2 control-label">Placeholder</label>

                    <div class="col-sm-10"><input type="text" placeholder="placeholder" class="form-control"></div>
                </div>
                <div class="hr-line-dashed"></div>
                <div class="form-group"><label class="col-lg-2 control-label">Disabled</label>

                    <div class="col-lg-10"><input type="text" disabled="" placeholder="Disabled input here..." class="form-control"></div>
                </div>
                <div class="hr-line-dashed"></div>
                <div class="form-group"><label class="col-lg-2 control-label">Static control</label>

                    <div class="col-lg-10"><p class="form-control-static">email@example.com</p></div>
                </div>
                <div class="hr-line-dashed"></div>
                <div class="form-group">
                    <div class="col-sm-8 col-sm-offset-2">
                        <button class="btn btn-default" type="submit">Cancel</button>
                        <button class="btn btn-primary" type="submit">Save changes</button>
                    </div>
                </div>
                </form>

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