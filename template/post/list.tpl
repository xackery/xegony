{{template "header" .}}

{{template "navigation" .}}



<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">



<div class="row" >
<div class="col-lg-12">

<div class="hpanel forum-box">

<div class="panel-heading">
    <span class="f"> Forum {{.Forum.Name}} > <span class="text-success">{{.Topic.Title}}</span> </span>
</div>

{{range $key, $value := .Posts}}
<div class="panel-body">

    <div class="media">
        <div class="media-image pull-left">
            <img src="images/a4.jpg" alt="profile-picture">
            <div class="author-info">
                <strong>Anna Smith</strong><br>
                April 11.2015
                <div class="badges">
                    <i class="fa fa-star text-warning"></i>
                    <i class="fa fa-shield text-success"></i>

                </div>
            </div>
        </div>
        <div class="media-body">
            {{$value.Body}}
        </div>
    </div>
</div>
{{end}}
</div>