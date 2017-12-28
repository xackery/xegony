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
</div>
</div>


{{if .Site.User}}
<div class="row" >
    <div class="col-lg-12">
    <div class="hpanel forum-box">

        <div class="panel-heading">
            Reply to {{.Topic.Title}}
        </div>

        <div class="panel-body">
            <div class="reply">
            </div>
            <div class="btn-group">
                    <button data-toggle="dropdown" class="btn btn-info btn-xs dropdown-toggle"><i class="ra ra-key"></i> Icon <span class="caret"></span></button>
                    <div class="dropdown-menu hdropdown">
                        <table>
                            <tbody>
                            <tr>
<td><a href="#"><i class="ra-tiny ra ra-aura"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-aware"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-double-team"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-falling"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-monster-skull"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-muscle-fat"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-muscle-up"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-player"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-player-despair"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-player-dodge"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-player-king"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-player-lift"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-player-pain"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-player-pyromaniac"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-player-shot"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-player-teleport"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-player-thunder-struck"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bottle-vapors"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bottled-bolt"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-broken-bottle"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-bubbling-potion"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-heart-bottle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-corked-tube"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fizzing-flask"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-flask"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-round-bottom-flask"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-vail"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-vase"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-ammo-bag"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-alligator-clip"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-ball"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-book"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-candle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-castle-flag"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-compass"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crown"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crown-of-thorns"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-horseshoe"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hourglass"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-jigsaw-piece"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-kettlebell"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-key"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-key-basic"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lantern-flame"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-wrench"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lit-candelabra"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-match"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-medical-pack"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-mirror"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-moon-sun"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-nails"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-noose"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-ocarina"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-pawn"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-pill"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-pills"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-ping-pong"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-potion"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-quill-ink"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-ringing-bell"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-rune-stone"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sherif"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-ship-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-shotgun-shell"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-slash-ring"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-snorkel"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-soccer-ball"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-spray"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-stopwatch"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-syringe"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-three-keys"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-torch"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-trophy"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-wooden-sign"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-beetle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bird-claw"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-butterfly"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cat"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dinosaur"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dragon"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-dragonfly"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-eye-monster"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fairy"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fish"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fox"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gecko"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hydra"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-insect-jaws"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lion"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-love-howl"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-maggot"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-octopus"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-rabbit"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-raven"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sea-serpent"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-seagull"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-shark"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sheep"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-snail"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-snake"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-spider-face"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-spiked-tentacle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-spiral-shell"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-suckered-tentacle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-tentacle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-two-dragons"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-venomous-snake"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-wyvern"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-wolf-head"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-wolf-howl"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-alien-fire"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-batteries"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-battery-0"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-battery-25"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-battery-50"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-battery-75"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-battery-100"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-battery-black"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-battery-negative"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-battery-positive"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-battery-white"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-clockwork"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cog"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cog-wheel"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-defibrilate"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-energise"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fast-ship"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gamepad-cross"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gear-heart"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gears"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-laser-site"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lever"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-light-bulb"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lighthouse"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-load"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-magnet"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-microphone"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-nuclear"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-radar-dish"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-radioactive"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-reactor"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-recycle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-regeneration"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-repair"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-robot-arm"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-rss"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sattelite"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-save"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-speech-bubble"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-speech-bubbles"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-surveillance-camera"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-telescope"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-tesla"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-unplugged"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-wifi"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-wireless-signal"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bottom-right"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cancel"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-clovers"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-clovers-card"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-diamonds"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-diamonds-card"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hearts"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hearts-card"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-spades"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-spades-card"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-suits"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-chessboard"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dice-one"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dice-two"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dice-three"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dice-four"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dice-five"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dice-six"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-perspective-dice-one"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-perspective-dice-two"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-perspective-dice-three"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-perspective-dice-four"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-perspective-dice-five"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-perspective-dice-six"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-perspective-dice-random"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-acorn"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-apple"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-beer"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-brandy-bottle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-carrot"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cheese"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-chicken-leg"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-coffee-mug"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crab-claw"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-egg"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-egg-pod"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-eggplant"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-honeycomb"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-ice-cube"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-knife-fork"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-meat"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-roast-chicken"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-toast"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-clover"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-daisy"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dead-tree"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-flower"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-flowers"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-grass"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-grass-patch"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-leaf"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-palm-tree"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-pine-tree"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sprout"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sprout-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-super-mushroom"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-trefoil-lily"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-zigzag-leaf"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-aquarius"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-aries"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cancer"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-capricorn"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gemini"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-libra"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-leo"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-pisces"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sagittarius"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-scorpio"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-taurus"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-virgo"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-acid"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-arson"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-biohazard"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-blade-bite"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-blast"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-blaster"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bleeding-eye"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bleeding-hearts"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bone-bite"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-burning-meteor"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crush"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-decapitation"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fall-down"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fire"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-food-chain"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-focused-lightning"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-heartburn"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-poison-cloud"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-tombstone"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-brain-freeze"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-burning-book"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-burning-embers"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-burning-eye"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-burst-blob"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cold-heart"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crystal-ball"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crystal-cluster"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crystal-wand"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crystals"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-diamond"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-divert"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-doubled"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dragon-breath"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-droplet"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-droplet-splash"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-emerald"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-eyeball"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fairy-wand"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fire-breath"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fire-ring"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fire-symbol"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-flame-symbol"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-frost-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-frostfire"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gem"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gem-pendant"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gloop"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gold-bar"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-health"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-health-decrease"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-health-increase"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hospital-cross"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hot-surface"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hydra-shot"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-incense"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-kaleidoscope"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lava"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-level-four"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-level-four-advanced"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-level-three"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-level-three-advanced"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-level-two"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-level-two-advanced"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lightning"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lightning-bolt"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lightning-storm"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lightning-trio"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sapphire"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-small-fire"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-snowflake"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sun"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sun-symbol"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sunbeams"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-triforce"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-two-hearts"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-water-drop"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-arcane-mask"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-all-for-one"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-anvil"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-archer"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-archery-target"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-arena"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-arrow-cluster"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-arrow-flights"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-axe"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-axe-swing"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-barbed-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-barrier"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bat-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-battered-axe"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-beam-wake"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bear-trap"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bolt-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bomb-explosion"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-bombs"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bone-knife"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-boomerang"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-boot-stomp"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bowie-knife"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-broadhead-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-broken-bone"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-broken-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bullets"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cannon-shot"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-chemical-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-chain"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-circular-saw"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-circular-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cluster-bomb"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cracked-helm"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cracked-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-croc-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crossbow"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crossed-axes"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-crossed-bones"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crossed-pistols"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crossed-sabers"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crossed-swords"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-daggers"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dervish-swords"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-diving-dagger"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-drill"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dripping-blade"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dripping-knife"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dripping-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-duel"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-explosion"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-explosive-materials"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-eye-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fire-bomb"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fire-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fireball-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-flaming-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-flaming-claw"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-flaming-trident"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-flat-hammer"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-frozen-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gavel"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-gear-hammer"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-grappling-hook"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-grenade"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-guillotine"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-halberd"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hammer"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hammer-drop"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hand-saw"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-harpoon-trident"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-helmet"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-horns"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-heavy-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-implosion"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-jetpack"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-kitchen-knives"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-knife"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-knight-helmet"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-kunai"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-large-hammer"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-laser-blast"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-lightning-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-mass-driver"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-mp5"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-musket"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-plain-dagger"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-relic-blade"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-revolver"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-rifle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-round-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-scythe"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-shuriken"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sickle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-spear-head"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-spikeball"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-spiked-mace"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-spinning-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-supersonic-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-target-arrows"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-target-laser"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-thorn-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-thorny-vine"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-trident"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-vest"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-vine-whip"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-zebra-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-ankh"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-angel-wings"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-anchor"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-batwings"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bell"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bird-mask"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bowling-pin"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-bridge"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-broken-skull"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-campfire"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-candle-fire"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-capitol"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-castle-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-circle-of-circles"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-crowned-heart"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cubes"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cut-palm"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-cycle"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-death-skull"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-demolish"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-desert-skull"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-dragon-wing"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-feather-wing"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-feathered-wing"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fedora"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-fluffy-swirl"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-footprint"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-forging"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-forward"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-glass-heart"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-groundbreaker"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-guarded-tower"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hand"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hand-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-heart-tower"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-heavy-fall"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-help"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hive-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hole-ladder"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-hood"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-horn-call"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-interdiction"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-locked-fortress"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-meat-hook"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-metal-gate"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-mine-wagon"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-mining-diamonds"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-mountains"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-nodular"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-ocean-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-omega"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-on-target"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-ophiuchus"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-overhead"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-overmind"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-pawprint"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-podium"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-pyramids"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-queen-crown"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-radial-balance"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-reverse"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-scroll-unfurled"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-shoe-prints"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-shot-through-the-heart"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-shovel"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-sideswipe"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-site"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-skull"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-skull-trophy"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny ra ra-spawn-node"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-splash"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-targeted"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-tic-tac-toe"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-tooth"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-tower"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-trail"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-turd"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-uncertainty"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-underhand"></i></a></td>
<td><a href="#"><i class="ra-tiny ra ra-x-mark"></i></a></td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            <span class="pull-right"><button type="button" class="btn w-xs btn-success">Reply</button></span>
        </div>



    </div>
    </div>
</div>
{{end}}
</div>
</div>

<script src="/vendor/summernote/dist/summernote.min.js"></script>
<script>
    $(document).ready(function() {

    $('.reply').summernote({
            height: 100,
            toolbar: [
                ['headline', ['style']],
                ['style', ['bold', 'italic', 'underline', 'superscript', 'subscript', 'strikethrough', 'clear']],
                ['textsize', ['fontsize']],
                ['alignment', ['ul', 'ol', 'paragraph', 'lineheight']],
            ]
        });
    });
    </script>