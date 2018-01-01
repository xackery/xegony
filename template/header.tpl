
<!--[if lt IE 7]>
<p class="alert alert-danger">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your experience.</p>
<![endif]-->

<!-- Header -->
<div id="header">
    <div class="color-line">
    </div>
    <div id="logo" class="light-version">
        <span>
            <img class="img-fluid" src="/images/logo.png" alt="RebuildEQ" style="width: 20px"/>{{.Site.Name}}
        </span>
    </div>
    <nav role="navigation">
        <div class="small-logo">
            <span class="text-primary">{{.Site.Name}}</span>
        </div>
        
        <div class="mobile-menu">
            <button type="button" class="navbar-toggle mobile-menu-toggle" data-toggle="collapse" data-target="#mobile-collapse">
                <i class="fa fa-chevron-down"></i>
            </button>
            <div class="collapse mobile-navbar" id="mobile-collapse">
                <ul class="nav" id="top-menu">
                    {{template "navmenu" .}}
                </ul>
            </div>
        </div>
        
        <div class="navbar-right">
            <ul class="nav navbar-nav no-borders">
                {{if .Site.User}}
                <li class="dropdown">
                    <a href="/dashboard">
                        <i class="pe-7s-graph2"></i>
                    </a>
                </li>
                <li class="dropdown">
                    <a class="dropdown-toggle" href="#" data-toggle="dropdown">
                        <i class="pe-7s-speaker"></i>
                    </a>
                    <ul class="dropdown-menu hdropdown notification flipInX">
                        <li>
                            <a>
                                <span class="label label-success">NEW</span> It is a long established.
                            </a>
                        </li>
                        <li>
                            <a>
                                <span class="label label-warning">WAR</span> There are many variations.
                            </a>
                        </li>
                        <li>
                            <a>
                                <span class="label label-danger">ERR</span> Contrary to popular belief.
                            </a>
                        </li>
                        <li class="summary"><a href="#">See all notifications</a></li>
                    </ul>
                </li>
                
                <li class="dropdown">
                    <a class="dropdown-toggle label-menu-corner" href="#" data-toggle="dropdown">
                        <i class="pe-7s-mail"></i>
                        <span class="label label-success">4</span>
                    </a>
                    <ul class="dropdown-menu hdropdown flipInX">
                        <div class="title">
                            You have 4 new messages
                        </div>
                        <li>
                            <a>
                                It is a long established.
                            </a>
                        </li>
                        <li>
                            <a>
                                There are many variations.
                            </a>
                        </li>
                        <li>
                            <a>
                                Lorem Ipsum is simply dummy.
                            </a>
                        </li>
                        <li>
                            <a>
                                Contrary to popular belief.
                            </a>
                        </li>
                        <li class="summary"><a href="#">See All Messages</a></li>
                    </ul>
                </li>
                <li class="dropdown">
                    <a class="dropdown-toggle label-menu-corner" href="#" data-toggle="dropdown">
                        <i class="pe-7s-user"></i>
                    </a>
                    <ul class="dropdown-menu hdropdown flipInX">                        
                        <li><a href="/dashboard">Dashboard</a></li>
                        <li><a href="/item/search/byaccount">Inventory Search</a></li>
                        <li><a href="/logout">Logout</a></li>
                    </ul>
                </li>
       
                {{else}}
                <li>
                    <a href="/login">
                        <i class="pe-7s-key"></i>
                    </a>
                </li>
                {{end}}
            </ul>

        </div>
    </nav>
</div>