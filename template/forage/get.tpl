{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-4">
        <div class="hpanel forum-box">            
            <div class="panel-heading">
                <span class="f"><a href="/item">Forage</a> > {{.Forage.Item.Name}}</span>
            </div>

            <div class="panel-body">                
                <span class="slot slotdrop"><span alt="{{.Forage.Item.Name}}" class="item icon-{{.Forage.Item.Icon}}"></span></span>
            
            <br/>
            <br/>
        
                <h3>{{.Forage.Item.Name}}</h3>
            </div>
            <div class="panel-body">
                <table class="table table-striped">
                    <tbody>
                    <td>ID</td><td><a href="/item/{{.Forage.Item.ID}}">{{.Forage.Item.ID}}</a></td>
                    </tbody>
                </table>
            </div>
            <div class="panel-footer contact-footer">
                <div class="row">
                    <div class="col-md-4 border-right">
                        <div class="contact-stat"><span>Projects: </span> <strong>200</strong></div>
                    </div>
                    <div class="col-md-4 border-right">
                        <div class="contact-stat"><span>Messages: </span> <strong>300</strong></div>
                    </div>
                    <div class="col-md-4">
                        <div class="contact-stat"><span>Views: </span> <strong>400</strong></div>
                    </div>
                </div>
            </div>

        </div>
    </div>    
</div>

</div>

</div>