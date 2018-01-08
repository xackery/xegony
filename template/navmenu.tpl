<li class="{{if eq .Site.Page "forum"}}active{{end}}"><a href="/"> <span class="nav-label"><i class="xa xa-coffee-mug"></i> Forum</span></a></li>
<li class="{{if eq .Site.Page "feature"}}active{{end}}"><a href="/feature"> <span class="nav-label"><i class="xa xa-bomb-explosion"></i> Features</span></a></li>
<li class="{{if eq .Site.Page "guide"}}active{{end}}"><a href="/guide"> <span class="nav-label"><i class="xa xa-trefoil-lily"></i> Guides</span></a></li>


{{if .Site.User}}
<div class="profile-picture">
    
    <div class="stats-label text-color">
     <a href="/dashboard">   <span class="font-extra-bold font-uppercase">{{.Site.User.Name}}</span></a>
        <div id="expLine" class="small-chart m-t-sm"></div>
        <div>
            <h4 class="font-extra-bold m-b-xs">
                1,123,123,123
            </h4>
            <small class="text-muted">Experience last 24 hours</small>
        </div>
    </div>
</div>
{{end}}