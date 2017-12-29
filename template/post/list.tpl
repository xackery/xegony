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
                    <button data-toggle="dropdown" class="btn btn-info btn-xs dropdown-toggle"><i class="xa xa-key"></i> Icon <span class="caret"></span></button>
                    <div class="dropdown-menu hdropdown">
                        <table>
                            <tbody>
                            <tr>
<td><a href="#"><i class="ra-tiny xa xa-aura"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-aware"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-double-team"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-falling"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-monster-skull"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-muscle-fat"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-muscle-up"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-player"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-player-despair"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-player-dodge"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-player-king"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-player-lift"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-player-pain"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-player-pyromaniac"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-player-shot"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-player-teleport"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-player-thunder-struck"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bottle-vapors"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bottled-bolt"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-broken-bottle"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-bubbling-potion"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-heart-bottle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-corked-tube"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fizzing-flask"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-flask"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-round-bottom-flask"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-vail"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-vase"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-ammo-bag"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-alligator-clip"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-ball"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-book"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-candle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-castle-flag"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-compass"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crown"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crown-of-thorns"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-horseshoe"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hourglass"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-jigsaw-piece"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-kettlebell"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-key"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-key-basic"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lantern-flame"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-wrench"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lit-candelabra"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-match"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-medical-pack"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-mirror"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-moon-sun"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-nails"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-noose"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-ocarina"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-pawn"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-pill"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-pills"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-ping-pong"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-potion"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-quill-ink"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-ringing-bell"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-rune-stone"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sherif"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-ship-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-shotgun-shell"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-slash-ring"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-snorkel"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-soccer-ball"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-spray"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-stopwatch"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-syringe"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-three-keys"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-torch"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-trophy"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-wooden-sign"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-beetle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bird-claw"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-butterfly"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cat"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dinosaur"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dragon"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-dragonfly"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-eye-monster"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fairy"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fish"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fox"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gecko"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hydra"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-insect-jaws"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lion"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-love-howl"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-maggot"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-octopus"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-rabbit"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-raven"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sea-serpent"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-seagull"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-shark"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sheep"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-snail"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-snake"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-spider-face"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-spiked-tentacle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-spiral-shell"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-suckered-tentacle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-tentacle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-two-dragons"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-venomous-snake"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-wyvern"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-wolf-head"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-wolf-howl"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-alien-fire"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-batteries"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-battery-0"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-battery-25"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-battery-50"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-battery-75"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-battery-100"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-battery-black"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-battery-negative"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-battery-positive"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-battery-white"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-clockwork"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cog"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cog-wheel"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-defibrilate"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-energise"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fast-ship"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gamepad-cross"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gear-heart"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gears"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-laser-site"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lever"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-light-bulb"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lighthouse"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-load"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-magnet"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-microphone"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-nuclear"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-radar-dish"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-radioactive"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-reactor"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-recycle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-regeneration"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-repair"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-robot-arm"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-rss"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sattelite"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-save"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-speech-bubble"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-speech-bubbles"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-surveillance-camera"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-telescope"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-tesla"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-unplugged"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-wifi"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-wireless-signal"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bottom-right"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cancel"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-clovers"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-clovers-card"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-diamonds"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-diamonds-card"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hearts"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hearts-card"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-spades"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-spades-card"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-suits"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-chessboard"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dice-one"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dice-two"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dice-three"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dice-four"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dice-five"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dice-six"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-perspective-dice-one"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-perspective-dice-two"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-perspective-dice-three"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-perspective-dice-four"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-perspective-dice-five"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-perspective-dice-six"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-perspective-dice-random"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-acorn"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-apple"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-beer"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-brandy-bottle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-carrot"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cheese"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-chicken-leg"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-coffee-mug"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crab-claw"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-egg"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-egg-pod"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-eggplant"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-honeycomb"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-ice-cube"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-knife-fork"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-meat"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-roast-chicken"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-toast"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-clover"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-daisy"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dead-tree"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-flower"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-flowers"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-grass"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-grass-patch"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-leaf"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-palm-tree"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-pine-tree"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sprout"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sprout-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-super-mushroom"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-trefoil-lily"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-zigzag-leaf"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-aquarius"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-aries"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cancer"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-capricorn"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gemini"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-libra"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-leo"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-pisces"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sagittarius"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-scorpio"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-taurus"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-virgo"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-acid"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-arson"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-biohazard"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-blade-bite"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-blast"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-blaster"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bleeding-eye"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bleeding-hearts"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bone-bite"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-burning-meteor"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crush"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-decapitation"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fall-down"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fire"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-food-chain"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-focused-lightning"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-heartburn"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-poison-cloud"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-tombstone"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-brain-freeze"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-burning-book"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-burning-embers"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-burning-eye"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-burst-blob"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cold-heart"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crystal-ball"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crystal-cluster"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crystal-wand"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crystals"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-diamond"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-divert"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-doubled"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dragon-breath"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-droplet"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-droplet-splash"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-emerald"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-eyeball"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fairy-wand"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fire-breath"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fire-ring"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fire-symbol"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-flame-symbol"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-frost-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-frostfire"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gem"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gem-pendant"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gloop"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gold-bar"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-health"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-health-decrease"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-health-increase"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hospital-cross"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hot-surface"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hydra-shot"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-incense"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-kaleidoscope"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lava"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-level-four"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-level-four-advanced"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-level-three"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-level-three-advanced"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-level-two"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-level-two-advanced"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lightning"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lightning-bolt"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lightning-storm"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lightning-trio"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sapphire"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-small-fire"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-snowflake"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sun"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sun-symbol"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sunbeams"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-triforce"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-two-hearts"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-water-drop"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-arcane-mask"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-all-for-one"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-anvil"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-archer"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-archery-target"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-arena"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-arrow-cluster"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-arrow-flights"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-axe"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-axe-swing"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-barbed-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-barrier"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bat-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-battered-axe"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-beam-wake"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bear-trap"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bolt-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bomb-explosion"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-bombs"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bone-knife"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-boomerang"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-boot-stomp"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bowie-knife"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-broadhead-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-broken-bone"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-broken-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bullets"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cannon-shot"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-chemical-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-chain"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-circular-saw"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-circular-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cluster-bomb"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cracked-helm"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cracked-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-croc-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crossbow"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crossed-axes"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-crossed-bones"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crossed-pistols"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crossed-sabers"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crossed-swords"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-daggers"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dervish-swords"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-diving-dagger"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-drill"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dripping-blade"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dripping-knife"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dripping-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-duel"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-explosion"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-explosive-materials"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-eye-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fire-bomb"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fire-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fireball-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-flaming-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-flaming-claw"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-flaming-trident"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-flat-hammer"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-frozen-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gavel"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-gear-hammer"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-grappling-hook"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-grenade"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-guillotine"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-halberd"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hammer"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hammer-drop"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hand-saw"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-harpoon-trident"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-helmet"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-horns"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-heavy-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-implosion"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-jetpack"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-kitchen-knives"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-knife"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-knight-helmet"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-kunai"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-large-hammer"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-laser-blast"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-lightning-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-mass-driver"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-mp5"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-musket"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-plain-dagger"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-relic-blade"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-revolver"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-rifle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-round-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-scythe"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-shuriken"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sickle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-spear-head"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-spikeball"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-spiked-mace"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-spinning-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-supersonic-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sword"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-target-arrows"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-target-laser"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-thorn-arrow"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-thorny-vine"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-trident"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-vest"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-vine-whip"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-zebra-shield"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-ankh"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-angel-wings"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-anchor"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-batwings"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bell"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bird-mask"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bowling-pin"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-bridge"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-broken-skull"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-campfire"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-candle-fire"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-capitol"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-castle-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-circle-of-circles"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-crowned-heart"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cubes"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cut-palm"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-cycle"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-death-skull"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-demolish"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-desert-skull"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-dragon-wing"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-feather-wing"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-feathered-wing"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fedora"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-fluffy-swirl"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-footprint"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-forging"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-forward"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-glass-heart"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-groundbreaker"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-guarded-tower"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hand"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hand-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-heart-tower"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-heavy-fall"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-help"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hive-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hole-ladder"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-hood"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-horn-call"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-interdiction"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-locked-fortress"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-meat-hook"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-metal-gate"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-mine-wagon"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-mining-diamonds"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-mountains"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-nodular"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-ocean-emblem"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-omega"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-on-target"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-ophiuchus"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-overhead"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-overmind"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-pawprint"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-podium"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-pyramids"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-queen-crown"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-radial-balance"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-reverse"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-scroll-unfurled"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-shoe-prints"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-shot-through-the-heart"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-shovel"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-sideswipe"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-site"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-skull"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-skull-trophy"></i></a></td>
</tr><tr>
<td><a href="#"><i class="ra-tiny xa xa-spawn-node"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-splash"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-targeted"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-tic-tac-toe"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-tooth"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-tower"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-trail"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-turd"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-uncertainty"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-underhand"></i></a></td>
<td><a href="#"><i class="ra-tiny xa xa-x-mark"></i></a></td>
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