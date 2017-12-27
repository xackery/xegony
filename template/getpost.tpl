{{template "header" .}}

{{template "navigation" .}}



<!-- Main Wrapper -->
<div id="wrapper">

    <div class="normalheader ">
        <div class="hpanel">
            <div class="panel-body">
                <a class="small-header-action" href="forum_details.html">
                    <div class="clip-header">
                        <i class="fa fa-arrow-up"></i>
                    </div>
                </a>

                <div id="hbreadcrumb" class="pull-right m-t-lg">
                    <ol class="hbreadcrumb breadcrumb">
                        <li><a href="index.html">Dashboard</a></li>
                        <li>
                            <span>App views</span>
                        </li>
                        <li class="active">
                            <span>Topic</span>
                        </li>
                    </ol>
                </div>
                <h2 class="font-light m-b-xs">
                    Topic
                </h2>
                <small>Topic for forum page.</small>
            </div>
        </div>
    </div>

<div class="content">



<div class="row" >
<div class="col-lg-12">

<div class="hpanel forum-box">

<div class="panel-heading">
                <span class="pull-right">
                    <i class="fa fa-clock-o"> </i> Last modification: 10.12.2015, 10:22 am
                </span>
    <span class="f"> General topics > Announcements > <span class="text-success">Free talks</span> </span>
</div>

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
            {{.Post.Body}}
        </div>
    </div>
</div>
</div>