{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="normalheader small-header">
    <div class="hpanel">
        <div class="panel-body">           

            <div id="hbreadcrumb" class="pull-right m-t-lg">
                <ol class="hbreadcrumb breadcrumb">
                    <li><a href="index.html">Dashboard</a></li>
                    <li>
                        <span>Tables</span>
                    </li>
                    <li class="active">
                        <span>Tables design</span>
                    </li>
                </ol>
            </div>
            <h2 class="font-light m-b-xs">
                {{.Character.Name}}
            </h2>
            <small>Examples of various designs of tables.</small>
        </div>
    </div>
</div>

<div class="content">


<div class="row">   
    <div class="col-md-offset-3 col-md-5">
        <ul class="nav nav-tabs">
            <li class="active"><a href="#inventory" data-toggle="tab" aria-expanded="true"><i class="glyphicon glyphicon-th"></i> Inventory</a></li>
            <li class=""><a href="#bank" data-toggle="tab" aria-expanded="false"><i class="glyphicon glyphicon-piggy-bank"></i> Bank</a></li>
            <li class=""><a href="#statistics" data-toggle="tab" aria-expanded="false"><i class="glyphicon glyphicon-stats"></i> Statistics</a></li>
        </ul>
        <div class="tab-content">
            <div class="tab-pane fade active in" id="inventory">
                <div class="panel panel-default">
                    <div class="panel-body">
                        <div class="item-inventory col-md-6" style="padding: 0px 5px">
                            <div class="panel panel-default">
                            
                                <div class="panel-heading">Equipment</div>
                                <div class="panel-body slotarea equipment-border" unselectable="on">
                                    <span class="char-monogram" style="background:url('/images/monograms/{{.Character.Class}}.gif') no-repeat;background-size:80px 160px;"></span>
                                    <div class="row">
                                        <div class="col-md-12">
                                            <span class="slot1 slotdrop ui-droppable" slot-id="1" slot-name="Left Ear" data-mask="18"><span class="{{if index .Inventory 1}}{{with index .Inventory 1}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot2 slotdrop ui-droppable" slot-id="2" slot-name="Head" data-mask="4"><span class="{{if index .Inventory 2}}{{with index .Inventory 2}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot3 slotdrop ui-droppable" slot-id="3" slot-name="Mask" data-mask="8"><span class="{{if index .Inventory 3}}{{with index .Inventory 3}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot4 slotdrop ui-droppable" slot-id="4" slot-name="Right Ear" data-mask="18"><span class="{{if index .Inventory 4}}{{with index .Inventory 4}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col-md-12">
                                            <span class="slot17 slotdrop ui-droppable" slot-id="17" slot-name="Chest" data-mask="131072"><span class="{{if index .Inventory 17}}{{with index .Inventory 17}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot-spacer"></span>
                                            <span class="slot-spacer"></span>
                                            <span class="slot5 slotdrop ui-droppable" slot-id="5" slot-name="Neck" data-mask="32"><span class="{{if index .Inventory 5}}{{with index .Inventory 5}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col-md-12">
                                            <span class="slot7 slotdrop ui-droppable" slot-id="7" slot-name="Arms" data-mask="128"><span class="{{if index .Inventory 7}}{{with index .Inventory 7}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot-spacer"></span>
                                            <span class="slot-spacer"></span>
                                            <span class="slot8 slotdrop ui-droppable" slot-id="8" slot-name="Back" data-mask="256"><span class="{{if index .Inventory 8}}{{with index .Inventory 8}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col-md-12">
                                            <span class="slot20 slotdrop ui-droppable" slot-id="20" slot-name="Waist" data-mask="1048576"><span class="{{if index .Inventory 20}}{{with index .Inventory 20}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot-spacer"></span>
                                            <span class="slot-spacer"></span>
                                            <span class="slot6 slotdrop ui-droppable" slot-id="6" slot-name="Shoulder" data-mask="64"><span class="{{if index .Inventory 6}}{{with index .Inventory 6}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col-md-12">
                                            <span class="slot9"  slot-id="9" slot-name="Left Bracer" data-mask="1536"><span class="{{if index .Inventory 9}}{{with index .Inventory 9}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot-spacer"></span>
                                            <span class="slot-spacer"></span>
                                            <span class="slot10 slotdrop ui-droppable" slot-id="10" slot-name="Right Bracer" data-mask="1536"><span class="{{if index .Inventory 10}}{{with index .Inventory 10}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col-md-12">
                                            <span class="slot18 slotdrop ui-droppable" slot-id="18" slot-name="Legs" data-mask="262144"><span class="{{if index .Inventory 18}}{{with index .Inventory 18}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot12 slotdrop ui-droppable" slot-id="12" slot-name="Hands" data-mask="4096"><span class="{{if index .Inventory 12}}{{with index .Inventory 12}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot0 slotdrop ui-droppable" slot-id="0" slot-name="Charm" data-mask="1"><span class="{{if index .Inventory 0}}{{with index .Inventory 0}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot19 slotdrop ui-droppable" slot-id="19" slot-name="Feet" data-mask="524288"><span class="{{if index .Inventory 19}}{{with index .Inventory 19}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col-md-12">
                                            <span class="slot-spacer"></span>
                                            <span class="slot15 slotdrop ui-droppable" slot-id="15" slot-name="Left Ring" data-mask="98304"><span class="{{if index .Inventory 15}}{{with index .Inventory 15}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot16 slotdrop ui-droppable" slot-id="16" slot-name="Right Ring" data-mask="98304"><span class="{{if index .Inventory 16}}{{with index .Inventory 16}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot9999 slotdrop ui-droppable" slot-id="9999" slot-name="Power Source" data-mask="4194304"><span class="{{if index .Inventory 9999}}{{with index .Inventory 9999}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col-md-12">
                                            <span class="slot13 slotdrop ui-droppable" slot-id="13" slot-name="Primary" data-mask="8192"><span class="{{if index .Inventory 13}}{{with index .Inventory 13}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot14 slotdrop ui-droppable" slot-id="14" slot-name="Secondary" data-mask="16384"><span class="{{if index .Inventory 14}}{{with index .Inventory 14}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot11 slotdrop ui-droppable" slot-id="11" slot-name="Range" data-mask="2048"><span class="{{if index .Inventory 11}}{{with index .Inventory 11}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                            <span class="slot21 slotdrop ui-droppable" slot-id="21" slot-name="Ammo" data-mask="2097152"><span class="{{if index .Inventory 21}}{{with index .Inventory 21}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="item-bags col-md-3" style="padding: 0px">
                            <div class="panel panel-default">
                                <div class="panel-heading">Inventory</div>
                                <div class="panel-body slotarea" unselectable="on" style="padding: 10px">
                                    <span class="slot22 slotdrop ui-droppable" slot-id="22" slot-name="Inventory"><span class="{{if index .Inventory 22}}{{with index .Inventory 22}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot26 slotdrop ui-droppable" slot-id="26" slot-name="Inventory"><span class="{{if index .Inventory 26}}{{with index .Inventory 26}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <br><br>
                                    <span class="slot23 slotdrop ui-droppable" slot-id="23" slot-name="Inventory"><span class="{{if index .Inventory 23}}{{with index .Inventory 23}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot27 slotdrop ui-droppable" slot-id="27" slot-name="Inventory"><span class="{{if index .Inventory 27}}{{with index .Inventory 27}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <br><br>
                                    <span class="slot24 slotdrop ui-droppable" slot-id="24" slot-name="Inventory"><span class="{{if index .Inventory 24}}{{with index .Inventory 24}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot28 slotdrop ui-droppable" slot-id="28" slot-name="Inventory"><span class="{{if index .Inventory 28}}{{with index .Inventory 28}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <br><br>
                                    <span class="slot25 slotdrop ui-droppable" slot-id="25" slot-name="Inventory"><span class="{{if index .Inventory 25}}{{with index .Inventory 25}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot29 slotdrop ui-droppable" slot-id="29" slot-name="Inventory"><span class="{{if index .Inventory 29}}{{with index .Inventory 29}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <br><br>
                                </div>
                            </div>                            
                        </div>
                        <div class="item-bags col-md-3" style="padding: 0px 0px 0px 5px">
                            <div class="bag22 panel panel-default" style="display: none">
                                <div class="panel-heading">Bag 1</div>
                                <div class="panel-body" style="padding: 10px">
                                    <span class="slot251 slot slotdrop ui-droppable" slot-id="251" slot-name="Bag 1-1"><span class="{{if index .Inventory 251}}{{with index .Inventory 251}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot252 slot slotdrop ui-droppable" slot-id="252" slot-name="Bag 1-2"><span class="{{if index .Inventory 252}}{{with index .Inventory 252}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot253 slot slotdrop ui-droppable" slot-id="253" slot-name="Bag 1-3"><span class="{{if index .Inventory 253}}{{with index .Inventory 253}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot254 slot slotdrop ui-droppable" slot-id="254" slot-name="Bag 1-4"><span class="{{if index .Inventory 254}}{{with index .Inventory 254}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot255 slot slotdrop ui-droppable" slot-id="255" slot-name="Bag 1-5"><span class="{{if index .Inventory 255}}{{with index .Inventory 255}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot256 slot slotdrop ui-droppable" slot-id="256" slot-name="Bag 1-6"><span class="{{if index .Inventory 256}}{{with index .Inventory 256}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot257 slot slotdrop ui-droppable" slot-id="257" slot-name="Bag 1-7"><span class="{{if index .Inventory 257}}{{with index .Inventory 257}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot258 slot slotdrop ui-droppable" slot-id="258" slot-name="Bag 1-8"><span class="{{if index .Inventory 258}}{{with index .Inventory 258}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot259 slot slotdrop ui-droppable" slot-id="259" slot-name="Bag 1-9"><span class="{{if index .Inventory 259}}{{with index .Inventory 259}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot260 slot slotdrop ui-droppable" slot-id="260" slot-name="Bag 1-10"><span class="{{if index .Inventory 260}}{{with index .Inventory 260}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                </div>
                            </div>
                            <div class="bag23 panel panel-default" style="display: none">
                                <div class="panel-heading">Bag 2</div>
                                <div class="panel-body" style="padding: 10px">
                                    <span class="slot261 slot slotdrop ui-droppable" slot-id="261" slot-name="Bag 2-1"><span class="{{if index .Inventory 261}}{{with index .Inventory 261}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot262 slot slotdrop ui-droppable" slot-id="262" slot-name="Bag 2-2"><span class="{{if index .Inventory 262}}{{with index .Inventory 262}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot263 slot slotdrop ui-droppable" slot-id="263" slot-name="Bag 2-3"><span class="{{if index .Inventory 263}}{{with index .Inventory 263}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot264 slot slotdrop ui-droppable" slot-id="264" slot-name="Bag 2-4"><span class="{{if index .Inventory 264}}{{with index .Inventory 264}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot265 slot slotdrop ui-droppable" slot-id="265" slot-name="Bag 2-5"><span class="{{if index .Inventory 265}}{{with index .Inventory 265}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot266 slot slotdrop ui-droppable" slot-id="266" slot-name="Bag 2-6"><span class="{{if index .Inventory 266}}{{with index .Inventory 266}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot267 slot slotdrop ui-droppable" slot-id="267" slot-name="Bag 2-7"><span class="{{if index .Inventory 267}}{{with index .Inventory 267}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot268 slot slotdrop ui-droppable" slot-id="268" slot-name="Bag 2-8"><span class="{{if index .Inventory 268}}{{with index .Inventory 268}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot269 slot slotdrop ui-droppable" slot-id="269" slot-name="Bag 2-9"><span class="{{if index .Inventory 269}}{{with index .Inventory 269}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot270 slot slotdrop ui-droppable" slot-id="270" slot-name="Bag 2-10"><span class="{{if index .Inventory 270}}{{with index .Inventory 270}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                </div>
                            </div>
                            <div class="bag24 panel panel-default" style="display: none">
                                <div class="panel-heading">Bag 3</div>
                                <div class="panel-body" style="padding: 10px">
                                    <span class="slot271 slot slotdrop ui-droppable" slot-id="271" slot-name="Bag 3-1"><span class="{{if index .Inventory 271}}{{with index .Inventory 271}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot272 slot slotdrop ui-droppable" slot-id="272" slot-name="Bag 3-2"><span class="{{if index .Inventory 272}}{{with index .Inventory 272}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot273 slot slotdrop ui-droppable" slot-id="273" slot-name="Bag 3-3"><span class="{{if index .Inventory 273}}{{with index .Inventory 273}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot274 slot slotdrop ui-droppable" slot-id="274" slot-name="Bag 3-4"><span class="{{if index .Inventory 274}}{{with index .Inventory 274}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot275 slot slotdrop ui-droppable" slot-id="275" slot-name="Bag 3-5"><span class="{{if index .Inventory 275}}{{with index .Inventory 275}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot276 slot slotdrop ui-droppable" slot-id="276" slot-name="Bag 3-6"><span class="{{if index .Inventory 276}}{{with index .Inventory 276}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot277 slot slotdrop ui-droppable" slot-id="277" slot-name="Bag 3-7"><span class="{{if index .Inventory 277}}{{with index .Inventory 277}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot278 slot slotdrop ui-droppable" slot-id="278" slot-name="Bag 3-8"><span class="{{if index .Inventory 278}}{{with index .Inventory 278}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot279 slot slotdrop ui-droppable" slot-id="279" slot-name="Bag 3-9"><span class="{{if index .Inventory 279}}{{with index .Inventory 279}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot280 slot slotdrop ui-droppable" slot-id="280" slot-name="Bag 3-10"><span class="{{if index .Inventory 280}}{{with index .Inventory 280}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                </div>
                            </div>
                            <div class="bag25 panel panel-default" style="display: none">
                                <div class="panel-heading">Bag 4</div>
                                <div class="panel-body" style="padding: 10px">
                                    <span class="slot281 slot slotdrop ui-droppable" slot-id="281" slot-name="Bag 4-1"><span class="{{if index .Inventory 281}}{{with index .Inventory 281}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot282 slot slotdrop ui-droppable" slot-id="282" slot-name="Bag 4-2"><span class="{{if index .Inventory 282}}{{with index .Inventory 282}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot283 slot slotdrop ui-droppable" slot-id="283" slot-name="Bag 4-3"><span class="{{if index .Inventory 283}}{{with index .Inventory 283}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot284 slot slotdrop ui-droppable" slot-id="284" slot-name="Bag 4-4"><span class="{{if index .Inventory 284}}{{with index .Inventory 284}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot285 slot slotdrop ui-droppable" slot-id="285" slot-name="Bag 4-5"><span class="{{if index .Inventory 285}}{{with index .Inventory 285}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot286 slot slotdrop ui-droppable" slot-id="286" slot-name="Bag 4-6"><span class="{{if index .Inventory 286}}{{with index .Inventory 286}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot287 slot slotdrop ui-droppable" slot-id="287" slot-name="Bag 4-7"><span class="{{if index .Inventory 287}}{{with index .Inventory 287}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot288 slot slotdrop ui-droppable" slot-id="288" slot-name="Bag 4-8"><span class="{{if index .Inventory 288}}{{with index .Inventory 288}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot289 slot slotdrop ui-droppable" slot-id="289" slot-name="Bag 4-9"><span class="{{if index .Inventory 289}}{{with index .Inventory 289}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot290 slot slotdrop ui-droppable" slot-id="290" slot-name="Bag 4-10"><span class="{{if index .Inventory 290}}{{with index .Inventory 290}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                </div>
                            </div>
                            <div class="bag26 panel panel-default" style="display: none">
                                <div class="panel-heading">Bag 5</div>
                                <div class="panel-body" style="padding: 10px">
                                    <span class="slot291 slot slotdrop ui-droppable" slot-id="291" slot-name="Bag 5-1"><span class="{{if index .Inventory 291}}{{with index .Inventory 291}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot292 slot slotdrop ui-droppable" slot-id="292" slot-name="Bag 5-2"><span class="{{if index .Inventory 292}}{{with index .Inventory 292}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot293 slot slotdrop ui-droppable" slot-id="293" slot-name="Bag 5-3"><span class="{{if index .Inventory 293}}{{with index .Inventory 293}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot294 slot slotdrop ui-droppable" slot-id="294" slot-name="Bag 5-4"><span class="{{if index .Inventory 294}}{{with index .Inventory 294}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot295 slot slotdrop ui-droppable" slot-id="295" slot-name="Bag 5-5"><span class="{{if index .Inventory 295}}{{with index .Inventory 295}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot296 slot slotdrop ui-droppable" slot-id="296" slot-name="Bag 5-6"><span class="{{if index .Inventory 296}}{{with index .Inventory 296}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot297 slot slotdrop ui-droppable" slot-id="297" slot-name="Bag 5-7"><span class="{{if index .Inventory 297}}{{with index .Inventory 297}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot298 slot slotdrop ui-droppable" slot-id="298" slot-name="Bag 5-8"><span class="{{if index .Inventory 298}}{{with index .Inventory 298}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot299 slot slotdrop ui-droppable" slot-id="299" slot-name="Bag 5-9"><span class="{{if index .Inventory 299}}{{with index .Inventory 299}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot300 slot slotdrop ui-droppable" slot-id="300" slot-name="Bag 5-10"><span class="{{if index .Inventory 300}}{{with index .Inventory 300}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                </div>
                            </div>
                            <div class="bag27 panel panel-default" style="display: none">
                                <div class="panel-heading">Bag 6</div>
                                <div class="panel-body" style="padding: 10px">
                                    <span class="slot301 slot slotdrop ui-droppable" slot-id="301" slot-name="Bag 6-1"><span class="{{if index .Inventory 301}}{{with index .Inventory 301}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot302 slot slotdrop ui-droppable" slot-id="302" slot-name="Bag 6-2"><span class="{{if index .Inventory 302}}{{with index .Inventory 302}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot303 slot slotdrop ui-droppable" slot-id="303" slot-name="Bag 6-3"><span class="{{if index .Inventory 303}}{{with index .Inventory 303}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot304 slot slotdrop ui-droppable" slot-id="304" slot-name="Bag 6-4"><span class="{{if index .Inventory 304}}{{with index .Inventory 304}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot305 slot slotdrop ui-droppable" slot-id="305" slot-name="Bag 6-5"><span class="{{if index .Inventory 305}}{{with index .Inventory 305}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot306 slot slotdrop ui-droppable" slot-id="306" slot-name="Bag 6-6"><span class="{{if index .Inventory 306}}{{with index .Inventory 306}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot307 slot slotdrop ui-droppable" slot-id="307" slot-name="Bag 6-7"><span class="{{if index .Inventory 307}}{{with index .Inventory 307}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot308 slot slotdrop ui-droppable" slot-id="308" slot-name="Bag 6-8"><span class="{{if index .Inventory 308}}{{with index .Inventory 308}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot309 slot slotdrop ui-droppable" slot-id="309" slot-name="Bag 6-9"><span class="{{if index .Inventory 309}}{{with index .Inventory 309}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot310 slot slotdrop ui-droppable" slot-id="310" slot-name="Bag 6-10"><span class="{{if index .Inventory 310}}{{with index .Inventory 310}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                </div>
                            </div>
                            <div class="bag28 panel panel-default" style="display: none">
                                <div class="panel-heading">Bag 7</div>
                                <div class="panel-body" style="padding: 10px">
                                    <span class="slot311 slot slotdrop ui-droppable" slot-id="311" slot-name="Bag 7-1"><span class="{{if index .Inventory 311}}{{with index .Inventory 311}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot312 slot slotdrop ui-droppable" slot-id="312" slot-name="Bag 7-2"><span class="{{if index .Inventory 312}}{{with index .Inventory 312}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot313 slot slotdrop ui-droppable" slot-id="313" slot-name="Bag 7-3"><span class="{{if index .Inventory 313}}{{with index .Inventory 313}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot314 slot slotdrop ui-droppable" slot-id="314" slot-name="Bag 7-4"><span class="{{if index .Inventory 314}}{{with index .Inventory 314}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot315 slot slotdrop ui-droppable" slot-id="315" slot-name="Bag 7-5"><span class="{{if index .Inventory 315}}{{with index .Inventory 315}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot316 slot slotdrop ui-droppable" slot-id="316" slot-name="Bag 7-6"><span class="{{if index .Inventory 316}}{{with index .Inventory 316}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot317 slot slotdrop ui-droppable" slot-id="317" slot-name="Bag 7-7"><span class="{{if index .Inventory 317}}{{with index .Inventory 317}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot318 slot slotdrop ui-droppable" slot-id="318" slot-name="Bag 7-8"><span class="{{if index .Inventory 318}}{{with index .Inventory 318}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot319 slot slotdrop ui-droppable" slot-id="319" slot-name="Bag 7-9"><span class="{{if index .Inventory 319}}{{with index .Inventory 319}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot320 slot slotdrop ui-droppable" slot-id="320" slot-name="Bag 7-10"><span class="{{if index .Inventory 320}}{{with index .Inventory 320}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                </div>
                            </div>
                            <div class="bag29 panel panel-default" style="display: none">
                                <div class="panel-heading">Bag 8</div>
                                <div class="panel-body" style="padding: 10px">
                                    <span class="slot321 slot slotdrop ui-droppable" slot-id="321" slot-name="Bag 8-1"><span class="{{if index .Inventory 321}}{{with index .Inventory 321}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot322 slot slotdrop ui-droppable" slot-id="322" slot-name="Bag 8-2"><span class="{{if index .Inventory 322}}{{with index .Inventory 322}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot323 slot slotdrop ui-droppable" slot-id="323" slot-name="Bag 8-3"><span class="{{if index .Inventory 323}}{{with index .Inventory 323}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot324 slot slotdrop ui-droppable" slot-id="324" slot-name="Bag 8-4"><span class="{{if index .Inventory 324}}{{with index .Inventory 324}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot325 slot slotdrop ui-droppable" slot-id="325" slot-name="Bag 8-5"><span class="{{if index .Inventory 325}}{{with index .Inventory 325}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot326 slot slotdrop ui-droppable" slot-id="326" slot-name="Bag 8-6"><span class="{{if index .Inventory 326}}{{with index .Inventory 326}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot327 slot slotdrop ui-droppable" slot-id="327" slot-name="Bag 8-7"><span class="{{if index .Inventory 327}}{{with index .Inventory 327}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot328 slot slotdrop ui-droppable" slot-id="328" slot-name="Bag 8-8"><span class="{{if index .Inventory 328}}{{with index .Inventory 328}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot329 slot slotdrop ui-droppable" slot-id="329" slot-name="Bag 8-9"><span class="{{if index .Inventory 329}}{{with index .Inventory 329}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot330 slot slotdrop ui-droppable" slot-id="330" slot-name="Bag 8-10"><span class="{{if index .Inventory 330}}{{with index .Inventory 330}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="tab-pane fade" id="bank">
                <div class="panel panel-default">
                    <div class="panel-body">
                        <div class="col-md-9 bank-border">
                            <div class="panel panel-default">
                                <div class="panel-heading">Bank</div>
                                <div class="panel-body slotarea" unselectable="on">
                                    <div class="bank-slots col-md-3">
                                        <span class="slot2000 slot slotdrop ui-droppable" slot-id="2000" slot-name="Bank"><div class="item ui-draggable ui-draggable-handle" item-id="32601" item-name="Backpack*" item-icon="565" slot-id="2000" data-mask="0" slot-name="Bank" item-quantity="1" id="findslot-2000" style="position: relative; background: url(&quot;/images/items/item_565.gif&quot;);" is-bag="1" bag-slots="8" onclick="GetDetails(this)"></div><span class="{{if index .Inventory 2000}}{{with index .Inventory 2000}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2004 slot slotdrop ui-droppable" slot-id="2004" slot-name="Bank"><div class="item ui-draggable ui-draggable-handle" item-id="32601" item-name="Backpack*" item-icon="565" slot-id="2004" data-mask="0" slot-name="Bank" item-quantity="1" id="findslot-2004" style="position: relative; background: url(&quot;/images/items/item_565.gif&quot;);" is-bag="1" bag-slots="8" onclick="GetDetails(this)"></div><span class="{{if index .Inventory 2004}}{{with index .Inventory 2004}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <br><br>
                                        <span class="slot2001 slot slotdrop ui-droppable" slot-id="2001" slot-name="Bank"><span class="{{if index .Inventory 2001}}{{with index .Inventory 2001}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2005 slot slotdrop ui-droppable" slot-id="2005" slot-name="Bank"><span class="{{if index .Inventory 2005}}{{with index .Inventory 2005}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <br><br>
                                        <span class="slot2002 slot slotdrop ui-droppable" slot-id="2002" slot-name="Bank"><span class="{{if index .Inventory 2002}}{{with index .Inventory 2002}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2006 slot slotdrop ui-droppable" slot-id="2006" slot-name="Bank"><span class="{{if index .Inventory 2006}}{{with index .Inventory 2006}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <br><br>
                                        <span class="slot2003 slot slotdrop ui-droppable" slot-id="2003" slot-name="Bank"><span class="{{if index .Inventory 2003}}{{with index .Inventory 2003}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2007 slot slotdrop ui-droppable" slot-id="2007" slot-name="Bank"><span class="{{if index .Inventory 2007}}{{with index .Inventory 2007}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    </div>
                                    <div class="bank-slots col-md-3">
                                        <span class="slot2008 slot slotdrop ui-droppable" slot-id="2008" slot-name="Bank"><span class="{{if index .Inventory 2008}}{{with index .Inventory 2008}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2012 slot slotdrop ui-droppable" slot-id="2012" slot-name="Bank"><span class="{{if index .Inventory 2012}}{{with index .Inventory 2012}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <br><br>
                                        <span class="slot2009 slot slotdrop ui-droppable" slot-id="2009" slot-name="Bank"><span class="{{if index .Inventory 2009}}{{with index .Inventory 2009}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2013 slot slotdrop ui-droppable" slot-id="2013" slot-name="Bank"><span class="{{if index .Inventory 2013}}{{with index .Inventory 2013}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <br><br>
                                        <span class="slot2010 slot slotdrop ui-droppable" slot-id="2010" slot-name="Bank"><span class="{{if index .Inventory 2010}}{{with index .Inventory 2010}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2014 slot slotdrop ui-droppable" slot-id="2014" slot-name="Bank"><span class="{{if index .Inventory 2014}}{{with index .Inventory 2014}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <br><br>
                                        <span class="slot2011 slot slotdrop ui-droppable" slot-id="2011" slot-name="Bank"><span class="{{if index .Inventory 2011}}{{with index .Inventory 2011}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2015 slot slotdrop ui-droppable" slot-id="2015" slot-name="Bank"><div class="item ui-draggable ui-draggable-handle" item-id="32601" item-name="Backpack*" item-icon="565" slot-id="2015" data-mask="0" slot-name="Bank" item-quantity="1" id="findslot-2015" style="position: relative; background: url(&quot;/images/items/item_565.gif&quot;);" is-bag="1" bag-slots="8" onclick="GetDetails(this)"></div><span class="{{if index .Inventory 2015}}{{with index .Inventory 2015}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    </div>
                                    <div class="bank-slots col-md-3">
                                        <span class="slot2016 slot slotdrop ui-droppable" slot-id="2016" slot-name="Bank"><span class="{{if index .Inventory 2016}}{{with index .Inventory 2016}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2020 slot slotdrop ui-droppable" slot-id="2020" slot-name="Bank"><span class="{{if index .Inventory 2020}}{{with index .Inventory 2020}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <br><br>
                                        <span class="slot2017 slot slotdrop ui-droppable" slot-id="2017" slot-name="Bank"><span class="{{if index .Inventory 2017}}{{with index .Inventory 2017}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2021 slot slotdrop ui-droppable" slot-id="2021" slot-name="Bank"><span class="{{if index .Inventory 2021}}{{with index .Inventory 2021}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <br><br>
                                        <span class="slot2018 slot slotdrop ui-droppable" slot-id="2018" slot-name="Bank"><span class="{{if index .Inventory 2018}}{{with index .Inventory 2018}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2022 slot slotdrop ui-droppable" slot-id="2022" slot-name="Bank"><span class="{{if index .Inventory 2022}}{{with index .Inventory 2022}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <br><br>
                                        <span class="slot2019 slot slotdrop ui-droppable" slot-id="2019" slot-name="Bank"><span class="{{if index .Inventory 2019}}{{with index .Inventory 2019}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                        <span class="slot2023 slot slotdrop ui-droppable" slot-id="2023" slot-name="Bank"><span class="{{if index .Inventory 2023}}{{with index .Inventory 2023}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    </div>
                                    <div class="bank-slots col-md-3">
                                        <div class="panel panel-default item-destroy">
                                            <div class="panel-heading">Destroy</div>
                                            <div class="panel-body slotarea destroy-border" unselectable="on">
                                                <span class="slot slotdrop ui-droppable" slot-id="-1" slot-name="Destroy"></span>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="item-bags col-md-3" style="padding: 0px">
                            <div class="bag2000 panel panel-default" style="display: none">
                                <div class="panel-heading">Bag 2000</div>
                                <div class="panel-body" style="padding: 10px">
                                    <span class="slot2031 slot slotdrop ui-droppable" slot-id="2031" slot-name="Bank Bag 2000-1"><div class="item ui-draggable ui-draggable-handle" item-id="9991" item-name="Bread Cakes*" item-icon="1021" slot-id="2031" data-mask="0" slot-name="Bank Bag 2000-1" item-quantity="18" id="findslot-2031" style="position: relative; background: url(&quot;/images/items/item_1021.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"><div class="item-stack-border"><span class="item-stack-count">18</span></div></div><span class="{{if index .Inventory 2031}}{{with index .Inventory 2031}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot2032 slot slotdrop ui-droppable" slot-id="2032" slot-name="Bank Bag 2000-2"><div class="item ui-draggable ui-draggable-handle" item-id="1008" item-name="Cloth Sleeves" item-icon="670" slot-id="2032" data-mask="128" slot-name="Bank Bag 2000-2" item-quantity="1" id="findslot-2032" style="position: relative; background: url(&quot;/images/items/item_670.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"></div><span class="{{if index .Inventory 2032}}{{with index .Inventory 2032}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot2033 slot slotdrop ui-droppable" slot-id="2033" slot-name="Bank Bag 2000-3"><span class="{{if index .Inventory 2033}}{{with index .Inventory 2033}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot2034 slot slotdrop ui-droppable" slot-id="2034" slot-name="Bank Bag 2000-4"><div class="item ui-draggable ui-draggable-handle" item-id="9990" item-name="Skin of Milk" item-icon="717" slot-id="2034" data-mask="0" slot-name="Bank Bag 2000-4" item-quantity="18" id="findslot-2034" style="position: relative; background: url(&quot;/images/items/item_717.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"><div class="item-stack-border"><span class="item-stack-count">18</span></div></div><span class="{{if index .Inventory 2034}}{{with index .Inventory 2034}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot2035 slot slotdrop ui-droppable" slot-id="2035" slot-name="Bank Bag 2000-5"><div class="item ui-draggable ui-draggable-handle" item-id="21779" item-name="Bandages*" item-icon="812" slot-id="2035" data-mask="0" slot-name="Bank Bag 2000-5" item-quantity="20" id="findslot-2035" style="position: relative; background: url(&quot;/images/items/item_812.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"><div class="item-stack-border"><span class="item-stack-count">20</span></div></div><span class="{{if index .Inventory 2035}}{{with index .Inventory 2035}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot2036 slot slotdrop ui-droppable" slot-id="2036" slot-name="Bank Bag 2000-6"><span class="{{if index .Inventory 2036}}{{with index .Inventory 2036}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot2037 slot slotdrop ui-droppable" slot-id="2037" slot-name="Bank Bag 2000-7"><span class="{{if index .Inventory 2037}}{{with index .Inventory 2037}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot2038 slot slotdrop ui-droppable" slot-id="2038" slot-name="Bank Bag 2000-8"><span class="{{if index .Inventory 2038}}{{with index .Inventory 2038}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot2039 slot slotdrop ui-droppable" slot-id="2039" slot-name="Bank Bag 2000-9"><span class="{{if index .Inventory 2039}}{{with index .Inventory 2039}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <span class="slot2040 slot slotdrop ui-droppable" slot-id="2040" slot-name="Bank Bag 2000-10"><span class="{{if index .Inventory 2040}}{{with index .Inventory 2040}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                    <br><br>
                                </div>
                            </div>
                            <div class="bag2001 panel panel-default" style="display: none">
                                <div class="panel-heading">Bag 2001</div>
                                <div class="panel-body" style="padding: 10px"><span class="slot2042 slot slotdrop ui-droppable" slot-id="2042" slot-name="Bank Bag 2001-1"></span>
                                <span class="slot2043 slot slotdrop ui-droppable" slot-id="2043" slot-name="Bank Bag 2001-2"><span class="{{if index .Inventory 2043}}{{with index .Inventory 2043}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                <span class="slot2044 slot slotdrop ui-droppable" slot-id="2044" slot-name="Bank Bag 2001-3"><span class="{{if index .Inventory 2044}}{{with index .Inventory 2044}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                <span class="slot2045 slot slotdrop ui-droppable" slot-id="2045" slot-name="Bank Bag 2001-4"><span class="{{if index .Inventory 2045}}{{with index .Inventory 2045}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                <span class="slot2046 slot slotdrop ui-droppable" slot-id="2046" slot-name="Bank Bag 2001-5"><span class="{{if index .Inventory 2046}}{{with index .Inventory 2046}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                <span class="slot2047 slot slotdrop ui-droppable" slot-id="2047" slot-name="Bank Bag 2001-6"><span class="{{if index .Inventory 2047}}{{with index .Inventory 2047}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                <span class="slot2048 slot slotdrop ui-droppable" slot-id="2048" slot-name="Bank Bag 2001-7"><span class="{{if index .Inventory 2048}}{{with index .Inventory 2048}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                <span class="slot2049 slot slotdrop ui-droppable" slot-id="2049" slot-name="Bank Bag 2001-8"><span class="{{if index .Inventory 2049}}{{with index .Inventory 2049}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                <span class="slot2050 slot slotdrop ui-droppable" slot-id="2050" slot-name="Bank Bag 2001-9"><span class="{{if index .Inventory 2050}}{{with index .Inventory 2050}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                                <span class="slot2051 slot slotdrop ui-droppable" slot-id="2051" slot-name="Bank Bag 2001-10"><span class="{{if index .Inventory 2051}}{{with index .Inventory 2051}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                            </div>
                        </div>
                        <div class="bag2002 panel panel-default" style="display: none">
                            <div class="panel-heading">Bag 2002</div>
                            <div class="panel-body" style="padding: 10px"><span class="slot2053 slot slotdrop ui-droppable" slot-id="2053" slot-name="Bank Bag 2002-1"></span>
                            <span class="slot2054 slot slotdrop ui-droppable" slot-id="2054" slot-name="Bank Bag 2002-2"><span class="{{if index .Inventory 2054}}{{with index .Inventory 2054}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                            <span class="slot2055 slot slotdrop ui-droppable" slot-id="2055" slot-name="Bank Bag 2002-3"><span class="{{if index .Inventory 2055}}{{with index .Inventory 2055}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                            <span class="slot2056 slot slotdrop ui-droppable" slot-id="2056" slot-name="Bank Bag 2002-4"><span class="{{if index .Inventory 2056}}{{with index .Inventory 2056}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                            <span class="slot2057 slot slotdrop ui-droppable" slot-id="2057" slot-name="Bank Bag 2002-5"><span class="{{if index .Inventory 2057}}{{with index .Inventory 2057}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                            <span class="slot2058 slot slotdrop ui-droppable" slot-id="2058" slot-name="Bank Bag 2002-6"><span class="{{if index .Inventory 2058}}{{with index .Inventory 2058}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                            <span class="slot2059 slot slotdrop ui-droppable" slot-id="2059" slot-name="Bank Bag 2002-7"><span class="{{if index .Inventory 2059}}{{with index .Inventory 2059}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                            <span class="slot2060 slot slotdrop ui-droppable" slot-id="2060" slot-name="Bank Bag 2002-8"><span class="{{if index .Inventory 2060}}{{with index .Inventory 2060}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                            <span class="slot2062 slot slotdrop ui-droppable" slot-id="2062" slot-name="Bank Bag 2002-9"><span class="{{if index .Inventory 2062}}{{with index .Inventory 2062}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                            <span class="slot2063 slot slotdrop ui-droppable" slot-id="2063" slot-name="Bank Bag 2002-10"><span class="{{if index .Inventory 2063}}{{with index .Inventory 2063}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        </div>
                    </div>
                    <div class="bag2003 panel panel-default" style="display: none">
                        <div class="panel-heading">Bag 2003</div>
                        <div class="panel-body" style="padding: 10px"><span class="slot2064 slot slotdrop ui-droppable" slot-id="2064" slot-name="Bank Bag 2003-1"></span>
                        <span class="slot2065 slot slotdrop ui-droppable" slot-id="2065" slot-name="Bank Bag 2003-2"><span class="{{if index .Inventory 2065}}{{with index .Inventory 2065}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2066 slot slotdrop ui-droppable" slot-id="2066" slot-name="Bank Bag 2003-3"><span class="{{if index .Inventory 2066}}{{with index .Inventory 2066}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2067 slot slotdrop ui-droppable" slot-id="2067" slot-name="Bank Bag 2003-4"><span class="{{if index .Inventory 2067}}{{with index .Inventory 2067}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2068 slot slotdrop ui-droppable" slot-id="2068" slot-name="Bank Bag 2003-5"><span class="{{if index .Inventory 2068}}{{with index .Inventory 2068}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2069 slot slotdrop ui-droppable" slot-id="2069" slot-name="Bank Bag 2003-6"><span class="{{if index .Inventory 2069}}{{with index .Inventory 2069}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2070 slot slotdrop ui-droppable" slot-id="2070" slot-name="Bank Bag 2003-7"><span class="{{if index .Inventory 2070}}{{with index .Inventory 2070}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2071 slot slotdrop ui-droppable" slot-id="2071" slot-name="Bank Bag 2003-8"><span class="{{if index .Inventory 2071}}{{with index .Inventory 2071}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2072 slot slotdrop ui-droppable" slot-id="2072" slot-name="Bank Bag 2003-9"><span class="{{if index .Inventory 2072}}{{with index .Inventory 2072}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2073 slot slotdrop ui-droppable" slot-id="2073" slot-name="Bank Bag 2003-10"><span class="{{if index .Inventory 2073}}{{with index .Inventory 2073}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
                <div class="bag2004 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2004</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2075 slot slotdrop ui-droppable" slot-id="2075" slot-name="Bank Bag 2004-1"><div class="item ui-draggable ui-draggable-handle" item-id="9991" item-name="Bread Cakes*" item-icon="1021" slot-id="2075" data-mask="0" slot-name="Bank Bag 2004-1" item-quantity="18" id="findslot-2075" style="position: relative; background: url(&quot;/images/items/item_1021.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"><div class="item-stack-border"><span class="item-stack-count">18</span></div></div><span class="{{if index .Inventory 2075}}{{with index .Inventory 2075}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2076 slot slotdrop ui-droppable" slot-id="2076" slot-name="Bank Bag 2004-2"><div class="item ui-draggable ui-draggable-handle" item-id="1008" item-name="Cloth Sleeves" item-icon="670" slot-id="2076" data-mask="128" slot-name="Bank Bag 2004-2" item-quantity="1" id="findslot-2076" style="position: relative; background: url(&quot;/images/items/item_670.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"></div><span class="{{if index .Inventory 2076}}{{with index .Inventory 2076}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2077 slot slotdrop ui-droppable" slot-id="2077" slot-name="Bank Bag 2004-3"><span class="{{if index .Inventory 2077}}{{with index .Inventory 2077}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2078 slot slotdrop ui-droppable" slot-id="2078" slot-name="Bank Bag 2004-4"><div class="item ui-draggable ui-draggable-handle" item-id="9990" item-name="Skin of Milk" item-icon="717" slot-id="2078" data-mask="0" slot-name="Bank Bag 2004-4" item-quantity="18" id="findslot-2078" style="position: relative; background: url(&quot;/images/items/item_717.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"><div class="item-stack-border"><span class="item-stack-count">18</span></div></div><span class="{{if index .Inventory 2078}}{{with index .Inventory 2078}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2079 slot slotdrop ui-droppable" slot-id="2079" slot-name="Bank Bag 2004-5"><div class="item ui-draggable ui-draggable-handle" item-id="21779" item-name="Bandages*" item-icon="812" slot-id="2079" data-mask="0" slot-name="Bank Bag 2004-5" item-quantity="20" id="findslot-2079" style="position: relative; background: url(&quot;/images/items/item_812.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"><div class="item-stack-border"><span class="item-stack-count">20</span></div></div><span class="{{if index .Inventory 2079}}{{with index .Inventory 2079}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2080 slot slotdrop ui-droppable" slot-id="2080" slot-name="Bank Bag 2004-6"><span class="{{if index .Inventory 2080}}{{with index .Inventory 2080}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2081 slot slotdrop ui-droppable" slot-id="2081" slot-name="Bank Bag 2004-7"><span class="{{if index .Inventory 2081}}{{with index .Inventory 2081}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2082 slot slotdrop ui-droppable" slot-id="2082" slot-name="Bank Bag 2004-8"><span class="{{if index .Inventory 2082}}{{with index .Inventory 2082}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2083 slot slotdrop ui-droppable" slot-id="2083" slot-name="Bank Bag 2004-9"><span class="{{if index .Inventory 2083}}{{with index .Inventory 2083}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2084 slot slotdrop ui-droppable" slot-id="2084" slot-name="Bank Bag 2004-10"><span class="{{if index .Inventory 2084}}{{with index .Inventory 2084}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2005 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2005</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2086 slot slotdrop ui-droppable" slot-id="2086" slot-name="Bank Bag 2005-1"><span class="{{if index .Inventory 2086}}{{with index .Inventory 2086}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2087 slot slotdrop ui-droppable" slot-id="2087" slot-name="Bank Bag 2005-2"><span class="{{if index .Inventory 2087}}{{with index .Inventory 2087}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2088 slot slotdrop ui-droppable" slot-id="2088" slot-name="Bank Bag 2005-3"><span class="{{if index .Inventory 2088}}{{with index .Inventory 2088}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2089 slot slotdrop ui-droppable" slot-id="2089" slot-name="Bank Bag 2005-4"><span class="{{if index .Inventory 2089}}{{with index .Inventory 2089}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2090 slot slotdrop ui-droppable" slot-id="2090" slot-name="Bank Bag 2005-5"><span class="{{if index .Inventory 2090}}{{with index .Inventory 2090}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2091 slot slotdrop ui-droppable" slot-id="2091" slot-name="Bank Bag 2005-6"><span class="{{if index .Inventory 2091}}{{with index .Inventory 2091}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2092 slot slotdrop ui-droppable" slot-id="2092" slot-name="Bank Bag 2005-7"><span class="{{if index .Inventory 2092}}{{with index .Inventory 2092}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2093 slot slotdrop ui-droppable" slot-id="2093" slot-name="Bank Bag 2005-8"><span class="{{if index .Inventory 2093}}{{with index .Inventory 2093}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2094 slot slotdrop ui-droppable" slot-id="2094" slot-name="Bank Bag 2005-9"><span class="{{if index .Inventory 2094}}{{with index .Inventory 2094}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2095 slot slotdrop ui-droppable" slot-id="2095" slot-name="Bank Bag 2005-10"><span class="{{if index .Inventory 2095}}{{with index .Inventory 2095}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
                <div class="bag2006 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2006</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2097 slot slotdrop ui-droppable" slot-id="2097" slot-name="Bank Bag 2006-1"><span class="{{if index .Inventory 2097}}{{with index .Inventory 2097}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2098 slot slotdrop ui-droppable" slot-id="2098" slot-name="Bank Bag 2006-2"><span class="{{if index .Inventory 2098}}{{with index .Inventory 2098}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2099 slot slotdrop ui-droppable" slot-id="2099" slot-name="Bank Bag 2006-3"><span class="{{if index .Inventory 2099}}{{with index .Inventory 2099}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2100 slot slotdrop ui-droppable" slot-id="2100" slot-name="Bank Bag 2006-4"><span class="{{if index .Inventory 2100}}{{with index .Inventory 2100}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2101 slot slotdrop ui-droppable" slot-id="2101" slot-name="Bank Bag 2006-5"><span class="{{if index .Inventory 2101}}{{with index .Inventory 2101}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2102 slot slotdrop ui-droppable" slot-id="2102" slot-name="Bank Bag 2006-6"><span class="{{if index .Inventory 2102}}{{with index .Inventory 2102}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2103 slot slotdrop ui-droppable" slot-id="2103" slot-name="Bank Bag 2006-7"><span class="{{if index .Inventory 2103}}{{with index .Inventory 2103}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2104 slot slotdrop ui-droppable" slot-id="2104" slot-name="Bank Bag 2006-8"><span class="{{if index .Inventory 2104}}{{with index .Inventory 2104}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2105 slot slotdrop ui-droppable" slot-id="2105" slot-name="Bank Bag 2006-9"><span class="{{if index .Inventory 2105}}{{with index .Inventory 2105}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2106 slot slotdrop ui-droppable" slot-id="2106" slot-name="Bank Bag 2006-10"><span class="{{if index .Inventory 2106}}{{with index .Inventory 2106}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2007 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2007</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2108 slot slotdrop ui-droppable" slot-id="2108" slot-name="Bank Bag 2007-1"><span class="{{if index .Inventory 2108}}{{with index .Inventory 2108}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2109 slot slotdrop ui-droppable" slot-id="2109" slot-name="Bank Bag 2007-2"><span class="{{if index .Inventory 2109}}{{with index .Inventory 2109}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2110 slot slotdrop ui-droppable" slot-id="2110" slot-name="Bank Bag 2007-3"><span class="{{if index .Inventory 2110}}{{with index .Inventory 2110}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2111 slot slotdrop ui-droppable" slot-id="2111" slot-name="Bank Bag 2007-4"><span class="{{if index .Inventory 2111}}{{with index .Inventory 2111}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2112 slot slotdrop ui-droppable" slot-id="2112" slot-name="Bank Bag 2007-5"><span class="{{if index .Inventory 2112}}{{with index .Inventory 2112}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2113 slot slotdrop ui-droppable" slot-id="2113" slot-name="Bank Bag 2007-6"><span class="{{if index .Inventory 2113}}{{with index .Inventory 2113}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2114 slot slotdrop ui-droppable" slot-id="2114" slot-name="Bank Bag 2007-7"><span class="{{if index .Inventory 2114}}{{with index .Inventory 2114}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2115 slot slotdrop ui-droppable" slot-id="2115" slot-name="Bank Bag 2007-8"><span class="{{if index .Inventory 2115}}{{with index .Inventory 2115}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2116 slot slotdrop ui-droppable" slot-id="2116" slot-name="Bank Bag 2007-9"><span class="{{if index .Inventory 2116}}{{with index .Inventory 2116}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2117 slot slotdrop ui-droppable" slot-id="2117" slot-name="Bank Bag 2007-10"><span class="{{if index .Inventory 2117}}{{with index .Inventory 2117}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
                <div class="bag2008 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2008</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2119 slot slotdrop ui-droppable" slot-id="2119" slot-name="Bank Bag 2008-1"><span class="{{if index .Inventory 2119}}{{with index .Inventory 2119}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2120 slot slotdrop ui-droppable" slot-id="2120" slot-name="Bank Bag 2008-2"><span class="{{if index .Inventory 2120}}{{with index .Inventory 2120}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2121 slot slotdrop ui-droppable" slot-id="2121" slot-name="Bank Bag 2008-3"><span class="{{if index .Inventory 2121}}{{with index .Inventory 2121}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2122 slot slotdrop ui-droppable" slot-id="2122" slot-name="Bank Bag 2008-4"><span class="{{if index .Inventory 2122}}{{with index .Inventory 2122}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2123 slot slotdrop ui-droppable" slot-id="2123" slot-name="Bank Bag 2008-5"><span class="{{if index .Inventory 2123}}{{with index .Inventory 2123}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2124 slot slotdrop ui-droppable" slot-id="2124" slot-name="Bank Bag 2008-6"><span class="{{if index .Inventory 2124}}{{with index .Inventory 2124}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2125 slot slotdrop ui-droppable" slot-id="2125" slot-name="Bank Bag 2008-7"><span class="{{if index .Inventory 2125}}{{with index .Inventory 2125}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2126 slot slotdrop ui-droppable" slot-id="2126" slot-name="Bank Bag 2008-8"><span class="{{if index .Inventory 2126}}{{with index .Inventory 2126}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2127 slot slotdrop ui-droppable" slot-id="2127" slot-name="Bank Bag 2008-9"><span class="{{if index .Inventory 2127}}{{with index .Inventory 2127}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2128 slot slotdrop ui-droppable" slot-id="2128" slot-name="Bank Bag 2008-10"><span class="{{if index .Inventory 2128}}{{with index .Inventory 2128}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2009 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2009</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2129 slot slotdrop ui-droppable" slot-id="2129" slot-name="Bank Bag 2009-1"><span class="{{if index .Inventory 2129}}{{with index .Inventory 2129}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2130 slot slotdrop ui-droppable" slot-id="2130" slot-name="Bank Bag 2009-2"><span class="{{if index .Inventory 2130}}{{with index .Inventory 2130}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2131 slot slotdrop ui-droppable" slot-id="2131" slot-name="Bank Bag 2009-3"><span class="{{if index .Inventory 2131}}{{with index .Inventory 2131}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2132 slot slotdrop ui-droppable" slot-id="2132" slot-name="Bank Bag 2009-4"><span class="{{if index .Inventory 2132}}{{with index .Inventory 2132}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2133 slot slotdrop ui-droppable" slot-id="2133" slot-name="Bank Bag 2009-5"><span class="{{if index .Inventory 2133}}{{with index .Inventory 2133}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2134 slot slotdrop ui-droppable" slot-id="2134" slot-name="Bank Bag 2009-6"><span class="{{if index .Inventory 2134}}{{with index .Inventory 2134}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2135 slot slotdrop ui-droppable" slot-id="2135" slot-name="Bank Bag 2009-7"><span class="{{if index .Inventory 2135}}{{with index .Inventory 2135}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2136 slot slotdrop ui-droppable" slot-id="2136" slot-name="Bank Bag 2009-8"><span class="{{if index .Inventory 2136}}{{with index .Inventory 2136}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2137 slot slotdrop ui-droppable" slot-id="2137" slot-name="Bank Bag 2009-9"><span class="{{if index .Inventory 2137}}{{with index .Inventory 2137}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2138 slot slotdrop ui-droppable" slot-id="2138" slot-name="Bank Bag 2009-10"><span class="{{if index .Inventory 2138}}{{with index .Inventory 2138}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2010 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2010</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2139 slot slotdrop ui-droppable" slot-id="2139" slot-name="Bank Bag 2010-1"><span class="{{if index .Inventory 2139}}{{with index .Inventory 2139}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2141 slot slotdrop ui-droppable" slot-id="2141" slot-name="Bank Bag 2010-2"><span class="{{if index .Inventory 2141}}{{with index .Inventory 2141}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2142 slot slotdrop ui-droppable" slot-id="2142" slot-name="Bank Bag 2010-3"><span class="{{if index .Inventory 2142}}{{with index .Inventory 2142}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2143 slot slotdrop ui-droppable" slot-id="2143" slot-name="Bank Bag 2010-4"><span class="{{if index .Inventory 2143}}{{with index .Inventory 2143}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2144 slot slotdrop ui-droppable" slot-id="2144" slot-name="Bank Bag 2010-5"><span class="{{if index .Inventory 2144}}{{with index .Inventory 2144}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2145 slot slotdrop ui-droppable" slot-id="2145" slot-name="Bank Bag 2010-6"><span class="{{if index .Inventory 2145}}{{with index .Inventory 2145}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2146 slot slotdrop ui-droppable" slot-id="2146" slot-name="Bank Bag 2010-7"><span class="{{if index .Inventory 2146}}{{with index .Inventory 2146}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2147 slot slotdrop ui-droppable" slot-id="2147" slot-name="Bank Bag 2010-8"><span class="{{if index .Inventory 2147}}{{with index .Inventory 2147}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2148 slot slotdrop ui-droppable" slot-id="2148" slot-name="Bank Bag 2010-9"><span class="{{if index .Inventory 2148}}{{with index .Inventory 2148}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2149 slot slotdrop ui-droppable" slot-id="2149" slot-name="Bank Bag 2010-10"><span class="{{if index .Inventory 2149}}{{with index .Inventory 2149}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
                <div class="bag2011 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2011</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2150 slot slotdrop ui-droppable" slot-id="2150" slot-name="Bank Bag 2011-1"><span class="{{if index .Inventory 2150}}{{with index .Inventory 2150}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2151 slot slotdrop ui-droppable" slot-id="2151" slot-name="Bank Bag 2011-2"><span class="{{if index .Inventory 2151}}{{with index .Inventory 2151}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2153 slot slotdrop ui-droppable" slot-id="2153" slot-name="Bank Bag 2011-3"><span class="{{if index .Inventory 2153}}{{with index .Inventory 2153}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2154 slot slotdrop ui-droppable" slot-id="2154" slot-name="Bank Bag 2011-4"><span class="{{if index .Inventory 2154}}{{with index .Inventory 2154}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2155 slot slotdrop ui-droppable" slot-id="2155" slot-name="Bank Bag 2011-5"><span class="{{if index .Inventory 2155}}{{with index .Inventory 2155}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2156 slot slotdrop ui-droppable" slot-id="2156" slot-name="Bank Bag 2011-6"><span class="{{if index .Inventory 2156}}{{with index .Inventory 2156}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2157 slot slotdrop ui-droppable" slot-id="2157" slot-name="Bank Bag 2011-7"><span class="{{if index .Inventory 2157}}{{with index .Inventory 2157}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2158 slot slotdrop ui-droppable" slot-id="2158" slot-name="Bank Bag 2011-8"><span class="{{if index .Inventory 2158}}{{with index .Inventory 2158}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2159 slot slotdrop ui-droppable" slot-id="2159" slot-name="Bank Bag 2011-9"><span class="{{if index .Inventory 2159}}{{with index .Inventory 2159}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2160 slot slotdrop ui-droppable" slot-id="2160" slot-name="Bank Bag 2011-10"><span class="{{if index .Inventory 2160}}{{with index .Inventory 2160}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2012 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2012</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2161 slot slotdrop ui-droppable" slot-id="2161" slot-name="Bank Bag 2012-1"><span class="{{if index .Inventory 2161}}{{with index .Inventory 2161}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2163 slot slotdrop ui-droppable" slot-id="2163" slot-name="Bank Bag 2012-2"><span class="{{if index .Inventory 2163}}{{with index .Inventory 2163}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2164 slot slotdrop ui-droppable" slot-id="2164" slot-name="Bank Bag 2012-3"><span class="{{if index .Inventory 2164}}{{with index .Inventory 2164}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2165 slot slotdrop ui-droppable" slot-id="2165" slot-name="Bank Bag 2012-4"><span class="{{if index .Inventory 2165}}{{with index .Inventory 2165}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2166 slot slotdrop ui-droppable" slot-id="2166" slot-name="Bank Bag 2012-5"><span class="{{if index .Inventory 2166}}{{with index .Inventory 2166}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2167 slot slotdrop ui-droppable" slot-id="2167" slot-name="Bank Bag 2012-6"><span class="{{if index .Inventory 2167}}{{with index .Inventory 2167}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2168 slot slotdrop ui-droppable" slot-id="2168" slot-name="Bank Bag 2012-7"><span class="{{if index .Inventory 2168}}{{with index .Inventory 2168}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2169 slot slotdrop ui-droppable" slot-id="2169" slot-name="Bank Bag 2012-8"><span class="{{if index .Inventory 2169}}{{with index .Inventory 2169}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2170 slot slotdrop ui-droppable" slot-id="2170" slot-name="Bank Bag 2012-9"><span class="{{if index .Inventory 2170}}{{with index .Inventory 2170}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2171 slot slotdrop ui-droppable" slot-id="2171" slot-name="Bank Bag 2012-10"><span class="{{if index .Inventory 2171}}{{with index .Inventory 2171}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
                <div class="bag2013 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2013</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2172 slot slotdrop ui-droppable" slot-id="2172" slot-name="Bank Bag 2013-1"><span class="{{if index .Inventory 2172}}{{with index .Inventory 2172}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2174 slot slotdrop ui-droppable" slot-id="2174" slot-name="Bank Bag 2013-2"><span class="{{if index .Inventory 2174}}{{with index .Inventory 2174}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2175 slot slotdrop ui-droppable" slot-id="2175" slot-name="Bank Bag 2013-3"><span class="{{if index .Inventory 2175}}{{with index .Inventory 2175}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2176 slot slotdrop ui-droppable" slot-id="2176" slot-name="Bank Bag 2013-4"><span class="{{if index .Inventory 2176}}{{with index .Inventory 2176}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2177 slot slotdrop ui-droppable" slot-id="2177" slot-name="Bank Bag 2013-5"><span class="{{if index .Inventory 2177}}{{with index .Inventory 2177}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2178 slot slotdrop ui-droppable" slot-id="2178" slot-name="Bank Bag 2013-6"><span class="{{if index .Inventory 2178}}{{with index .Inventory 2178}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2179 slot slotdrop ui-droppable" slot-id="2179" slot-name="Bank Bag 2013-7"><span class="{{if index .Inventory 2179}}{{with index .Inventory 2179}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2180 slot slotdrop ui-droppable" slot-id="2180" slot-name="Bank Bag 2013-8"><span class="{{if index .Inventory 2180}}{{with index .Inventory 2180}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2181 slot slotdrop ui-droppable" slot-id="2181" slot-name="Bank Bag 2013-9"><span class="{{if index .Inventory 2181}}{{with index .Inventory 2181}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2182 slot slotdrop ui-droppable" slot-id="2182" slot-name="Bank Bag 2013-10"><span class="{{if index .Inventory 2182}}{{with index .Inventory 2182}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2014 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2014</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2183 slot slotdrop ui-droppable" slot-id="2183" slot-name="Bank Bag 2014-1"><span class="{{if index .Inventory 2183}}{{with index .Inventory 2183}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2185 slot slotdrop ui-droppable" slot-id="2185" slot-name="Bank Bag 2014-2"><span class="{{if index .Inventory 2185}}{{with index .Inventory 2185}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2186 slot slotdrop ui-droppable" slot-id="2186" slot-name="Bank Bag 2014-3"><span class="{{if index .Inventory 2186}}{{with index .Inventory 2186}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2187 slot slotdrop ui-droppable" slot-id="2187" slot-name="Bank Bag 2014-4"><span class="{{if index .Inventory 2187}}{{with index .Inventory 2187}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2188 slot slotdrop ui-droppable" slot-id="2188" slot-name="Bank Bag 2014-5"><span class="{{if index .Inventory 2188}}{{with index .Inventory 2188}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2189 slot slotdrop ui-droppable" slot-id="2189" slot-name="Bank Bag 2014-6"><span class="{{if index .Inventory 2189}}{{with index .Inventory 2189}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2190 slot slotdrop ui-droppable" slot-id="2190" slot-name="Bank Bag 2014-7"><span class="{{if index .Inventory 2190}}{{with index .Inventory 2190}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2191 slot slotdrop ui-droppable" slot-id="2191" slot-name="Bank Bag 2014-8"><span class="{{if index .Inventory 2191}}{{with index .Inventory 2191}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2192 slot slotdrop ui-droppable" slot-id="2192" slot-name="Bank Bag 2014-9"><span class="{{if index .Inventory 2192}}{{with index .Inventory 2192}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2193 slot slotdrop ui-droppable" slot-id="2193" slot-name="Bank Bag 2014-10"><span class="{{if index .Inventory 2193}}{{with index .Inventory 2193}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
                <div class="bag2015 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2015</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2194 slot slotdrop ui-droppable" slot-id="2194" slot-name="Bank Bag 2015-1"><span class="{{if index .Inventory 2194}}{{with index .Inventory 2194}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2196 slot slotdrop ui-droppable" slot-id="2196" slot-name="Bank Bag 2015-2"><div class="item ui-draggable ui-draggable-handle" item-id="9991" item-name="Bread Cakes*" item-icon="1021" slot-id="2196" data-mask="0" slot-name="Bank Bag 2015-2" item-quantity="18" id="findslot-2196" style="position: relative; background: url(&quot;/images/items/item_1021.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"><div class="item-stack-border"><span class="item-stack-count">18</span></div></div><span class="{{if index .Inventory 2196}}{{with index .Inventory 2196}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2197 slot slotdrop ui-droppable" slot-id="2197" slot-name="Bank Bag 2015-3"><div class="item ui-draggable ui-draggable-handle" item-id="1008" item-name="Cloth Sleeves" item-icon="670" slot-id="2197" data-mask="128" slot-name="Bank Bag 2015-3" item-quantity="1" id="findslot-2197" style="position: relative; background: url(&quot;/images/items/item_670.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"></div><span class="{{if index .Inventory 2197}}{{with index .Inventory 2197}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2198 slot slotdrop ui-droppable" slot-id="2198" slot-name="Bank Bag 2015-4"><span class="{{if index .Inventory 2198}}{{with index .Inventory 2198}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2199 slot slotdrop ui-droppable" slot-id="2199" slot-name="Bank Bag 2015-5"><div class="item ui-draggable ui-draggable-handle" item-id="9990" item-name="Skin of Milk" item-icon="717" slot-id="2199" data-mask="0" slot-name="Bank Bag 2015-5" item-quantity="18" id="findslot-2199" style="position: relative; background: url(&quot;/images/items/item_717.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"><div class="item-stack-border"><span class="item-stack-count">18</span></div></div><span class="{{if index .Inventory 2199}}{{with index .Inventory 2199}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2200 slot slotdrop ui-droppable" slot-id="2200" slot-name="Bank Bag 2015-6"><div class="item ui-draggable ui-draggable-handle" item-id="21779" item-name="Bandages*" item-icon="812" slot-id="2200" data-mask="0" slot-name="Bank Bag 2015-6" item-quantity="20" id="findslot-2200" style="position: relative; background: url(&quot;/images/items/item_812.gif&quot;);" is-bag="0" bag-slots="0" onclick="GetDetails(this)"><div class="item-stack-border"><span class="item-stack-count">20</span></div></div><span class="{{if index .Inventory 2200}}{{with index .Inventory 2200}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2201 slot slotdrop ui-droppable" slot-id="2201" slot-name="Bank Bag 2015-7"><span class="{{if index .Inventory 2201}}{{with index .Inventory 2201}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2202 slot slotdrop ui-droppable" slot-id="2202" slot-name="Bank Bag 2015-8"><span class="{{if index .Inventory 2202}}{{with index .Inventory 2202}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2203 slot slotdrop ui-droppable" slot-id="2203" slot-name="Bank Bag 2015-9"><span class="{{if index .Inventory 2203}}{{with index .Inventory 2203}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2204 slot slotdrop ui-droppable" slot-id="2204" slot-name="Bank Bag 2015-10"><span class="{{if index .Inventory 2204}}{{with index .Inventory 2204}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2016 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2016</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2205 slot slotdrop ui-droppable" slot-id="2205" slot-name="Bank Bag 2016-1"><span class="{{if index .Inventory 2205}}{{with index .Inventory 2205}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2207 slot slotdrop ui-droppable" slot-id="2207" slot-name="Bank Bag 2016-2"><span class="{{if index .Inventory 2207}}{{with index .Inventory 2207}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2208 slot slotdrop ui-droppable" slot-id="2208" slot-name="Bank Bag 2016-3"><span class="{{if index .Inventory 2208}}{{with index .Inventory 2208}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2209 slot slotdrop ui-droppable" slot-id="2209" slot-name="Bank Bag 2016-4"><span class="{{if index .Inventory 2209}}{{with index .Inventory 2209}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2210 slot slotdrop ui-droppable" slot-id="2210" slot-name="Bank Bag 2016-5"><span class="{{if index .Inventory 2210}}{{with index .Inventory 2210}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2211 slot slotdrop ui-droppable" slot-id="2211" slot-name="Bank Bag 2016-6"><span class="{{if index .Inventory 2211}}{{with index .Inventory 2211}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2212 slot slotdrop ui-droppable" slot-id="2212" slot-name="Bank Bag 2016-7"><span class="{{if index .Inventory 2212}}{{with index .Inventory 2212}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2213 slot slotdrop ui-droppable" slot-id="2213" slot-name="Bank Bag 2016-8"><span class="{{if index .Inventory 2213}}{{with index .Inventory 2213}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2214 slot slotdrop ui-droppable" slot-id="2214" slot-name="Bank Bag 2016-9"><span class="{{if index .Inventory 2214}}{{with index .Inventory 2214}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2215 slot slotdrop ui-droppable" slot-id="2215" slot-name="Bank Bag 2016-10"><span class="{{if index .Inventory 2215}}{{with index .Inventory 2215}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
                <div class="bag2017 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2017</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2216 slot slotdrop ui-droppable" slot-id="2216" slot-name="Bank Bag 2017-1"><span class="{{if index .Inventory 2216}}{{with index .Inventory 2216}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2218 slot slotdrop ui-droppable" slot-id="2218" slot-name="Bank Bag 2017-2"><span class="{{if index .Inventory 2218}}{{with index .Inventory 2218}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2219 slot slotdrop ui-droppable" slot-id="2219" slot-name="Bank Bag 2017-3"><span class="{{if index .Inventory 2219}}{{with index .Inventory 2219}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2220 slot slotdrop ui-droppable" slot-id="2220" slot-name="Bank Bag 2017-4"><span class="{{if index .Inventory 2220}}{{with index .Inventory 2220}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2221 slot slotdrop ui-droppable" slot-id="2221" slot-name="Bank Bag 2017-5"><span class="{{if index .Inventory 2221}}{{with index .Inventory 2221}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2222 slot slotdrop ui-droppable" slot-id="2222" slot-name="Bank Bag 2017-6"><span class="{{if index .Inventory 2222}}{{with index .Inventory 2222}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2223 slot slotdrop ui-droppable" slot-id="2223" slot-name="Bank Bag 2017-7"><span class="{{if index .Inventory 2223}}{{with index .Inventory 2223}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2224 slot slotdrop ui-droppable" slot-id="2224" slot-name="Bank Bag 2017-8"><span class="{{if index .Inventory 2224}}{{with index .Inventory 2224}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2225 slot slotdrop ui-droppable" slot-id="2225" slot-name="Bank Bag 2017-9"><span class="{{if index .Inventory 2225}}{{with index .Inventory 2225}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2226 slot slotdrop ui-droppable" slot-id="2226" slot-name="Bank Bag 2017-10"><span class="{{if index .Inventory 2226}}{{with index .Inventory 2226}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2018 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2018</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2227 slot slotdrop ui-droppable" slot-id="2227" slot-name="Bank Bag 2018-1"><span class="{{if index .Inventory 2227}}{{with index .Inventory 2227}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2229 slot slotdrop ui-droppable" slot-id="2229" slot-name="Bank Bag 2018-2"><span class="{{if index .Inventory 2229}}{{with index .Inventory 2229}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2230 slot slotdrop ui-droppable" slot-id="2230" slot-name="Bank Bag 2018-3"><span class="{{if index .Inventory 2230}}{{with index .Inventory 2230}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2231 slot slotdrop ui-droppable" slot-id="2231" slot-name="Bank Bag 2018-4"><span class="{{if index .Inventory 2231}}{{with index .Inventory 2231}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2232 slot slotdrop ui-droppable" slot-id="2232" slot-name="Bank Bag 2018-5"><span class="{{if index .Inventory 2232}}{{with index .Inventory 2232}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2233 slot slotdrop ui-droppable" slot-id="2233" slot-name="Bank Bag 2018-6"><span class="{{if index .Inventory 2233}}{{with index .Inventory 2233}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2234 slot slotdrop ui-droppable" slot-id="2234" slot-name="Bank Bag 2018-7"><span class="{{if index .Inventory 2234}}{{with index .Inventory 2234}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2235 slot slotdrop ui-droppable" slot-id="2235" slot-name="Bank Bag 2018-8"><span class="{{if index .Inventory 2235}}{{with index .Inventory 2235}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2236 slot slotdrop ui-droppable" slot-id="2236" slot-name="Bank Bag 2018-9"><span class="{{if index .Inventory 2236}}{{with index .Inventory 2236}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2237 slot slotdrop ui-droppable" slot-id="2237" slot-name="Bank Bag 2018-10"><span class="{{if index .Inventory 2237}}{{with index .Inventory 2237}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
                <div class="bag2019 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2019</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2238 slot slotdrop ui-droppable" slot-id="2238" slot-name="Bank Bag 2019-1"><span class="{{if index .Inventory 2238}}{{with index .Inventory 2238}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2240 slot slotdrop ui-droppable" slot-id="2240" slot-name="Bank Bag 2019-2"><span class="{{if index .Inventory 2240}}{{with index .Inventory 2240}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2241 slot slotdrop ui-droppable" slot-id="2241" slot-name="Bank Bag 2019-3"><span class="{{if index .Inventory 2241}}{{with index .Inventory 2241}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2242 slot slotdrop ui-droppable" slot-id="2242" slot-name="Bank Bag 2019-4"><span class="{{if index .Inventory 2242}}{{with index .Inventory 2242}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2243 slot slotdrop ui-droppable" slot-id="2243" slot-name="Bank Bag 2019-5"><span class="{{if index .Inventory 2243}}{{with index .Inventory 2243}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2244 slot slotdrop ui-droppable" slot-id="2244" slot-name="Bank Bag 2019-6"><span class="{{if index .Inventory 2244}}{{with index .Inventory 2244}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2245 slot slotdrop ui-droppable" slot-id="2245" slot-name="Bank Bag 2019-7"><span class="{{if index .Inventory 2245}}{{with index .Inventory 2245}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2246 slot slotdrop ui-droppable" slot-id="2246" slot-name="Bank Bag 2019-8"><span class="{{if index .Inventory 2246}}{{with index .Inventory 2246}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2247 slot slotdrop ui-droppable" slot-id="2247" slot-name="Bank Bag 2019-9"><span class="{{if index .Inventory 2247}}{{with index .Inventory 2247}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2248 slot slotdrop ui-droppable" slot-id="2248" slot-name="Bank Bag 2019-10"><span class="{{if index .Inventory 2248}}{{with index .Inventory 2248}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2020 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2020</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2249 slot slotdrop ui-droppable" slot-id="2249" slot-name="Bank Bag 2020-1"><span class="{{if index .Inventory 2249}}{{with index .Inventory 2249}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2251 slot slotdrop ui-droppable" slot-id="2251" slot-name="Bank Bag 2020-2"><span class="{{if index .Inventory 2251}}{{with index .Inventory 2251}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2252 slot slotdrop ui-droppable" slot-id="2252" slot-name="Bank Bag 2020-3"><span class="{{if index .Inventory 2252}}{{with index .Inventory 2252}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2253 slot slotdrop ui-droppable" slot-id="2253" slot-name="Bank Bag 2020-4"><span class="{{if index .Inventory 2253}}{{with index .Inventory 2253}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2254 slot slotdrop ui-droppable" slot-id="2254" slot-name="Bank Bag 2020-5"><span class="{{if index .Inventory 2254}}{{with index .Inventory 2254}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2255 slot slotdrop ui-droppable" slot-id="2255" slot-name="Bank Bag 2020-6"><span class="{{if index .Inventory 2255}}{{with index .Inventory 2255}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2256 slot slotdrop ui-droppable" slot-id="2256" slot-name="Bank Bag 2020-7"><span class="{{if index .Inventory 2256}}{{with index .Inventory 2256}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2257 slot slotdrop ui-droppable" slot-id="2257" slot-name="Bank Bag 2020-8"><span class="{{if index .Inventory 2257}}{{with index .Inventory 2257}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2258 slot slotdrop ui-droppable" slot-id="2258" slot-name="Bank Bag 2020-9"><span class="{{if index .Inventory 2258}}{{with index .Inventory 2258}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2259 slot slotdrop ui-droppable" slot-id="2259" slot-name="Bank Bag 2020-10"><span class="{{if index .Inventory 2259}}{{with index .Inventory 2259}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
                <div class="bag2021 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2021</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2260 slot slotdrop ui-droppable" slot-id="2260" slot-name="Bank Bag 2021-1"><span class="{{if index .Inventory 2260}}{{with index .Inventory 2260}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2262 slot slotdrop ui-droppable" slot-id="2262" slot-name="Bank Bag 2021-2"><span class="{{if index .Inventory 2262}}{{with index .Inventory 2262}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2263 slot slotdrop ui-droppable" slot-id="2263" slot-name="Bank Bag 2021-3"><span class="{{if index .Inventory 2263}}{{with index .Inventory 2263}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2264 slot slotdrop ui-droppable" slot-id="2264" slot-name="Bank Bag 2021-4"><span class="{{if index .Inventory 2264}}{{with index .Inventory 2264}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2265 slot slotdrop ui-droppable" slot-id="2265" slot-name="Bank Bag 2021-5"><span class="{{if index .Inventory 2265}}{{with index .Inventory 2265}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2266 slot slotdrop ui-droppable" slot-id="2266" slot-name="Bank Bag 2021-6"><span class="{{if index .Inventory 2266}}{{with index .Inventory 2266}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2267 slot slotdrop ui-droppable" slot-id="2267" slot-name="Bank Bag 2021-7"><span class="{{if index .Inventory 2267}}{{with index .Inventory 2267}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2268 slot slotdrop ui-droppable" slot-id="2268" slot-name="Bank Bag 2021-8"><span class="{{if index .Inventory 2268}}{{with index .Inventory 2268}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2269 slot slotdrop ui-droppable" slot-id="2269" slot-name="Bank Bag 2021-9"><span class="{{if index .Inventory 2269}}{{with index .Inventory 2269}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2270 slot slotdrop ui-droppable" slot-id="2270" slot-name="Bank Bag 2021-10"><span class="{{if index .Inventory 2270}}{{with index .Inventory 2270}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2022 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2022</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2271 slot slotdrop ui-droppable" slot-id="2271" slot-name="Bank Bag 2022-1"><span class="{{if index .Inventory 2271}}{{with index .Inventory 2271}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2273 slot slotdrop ui-droppable" slot-id="2273" slot-name="Bank Bag 2022-2"><span class="{{if index .Inventory 2273}}{{with index .Inventory 2273}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2274 slot slotdrop ui-droppable" slot-id="2274" slot-name="Bank Bag 2022-3"><span class="{{if index .Inventory 2274}}{{with index .Inventory 2274}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2275 slot slotdrop ui-droppable" slot-id="2275" slot-name="Bank Bag 2022-4"><span class="{{if index .Inventory 2275}}{{with index .Inventory 2275}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2276 slot slotdrop ui-droppable" slot-id="2276" slot-name="Bank Bag 2022-5"><span class="{{if index .Inventory 2276}}{{with index .Inventory 2276}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2277 slot slotdrop ui-droppable" slot-id="2277" slot-name="Bank Bag 2022-6"><span class="{{if index .Inventory 2277}}{{with index .Inventory 2277}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2278 slot slotdrop ui-droppable" slot-id="2278" slot-name="Bank Bag 2022-7"><span class="{{if index .Inventory 2278}}{{with index .Inventory 2278}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2279 slot slotdrop ui-droppable" slot-id="2279" slot-name="Bank Bag 2022-8"><span class="{{if index .Inventory 2279}}{{with index .Inventory 2279}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2280 slot slotdrop ui-droppable" slot-id="2280" slot-name="Bank Bag 2022-9"><span class="{{if index .Inventory 2280}}{{with index .Inventory 2280}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2281 slot slotdrop ui-droppable" slot-id="2281" slot-name="Bank Bag 2022-10"><span class="{{if index .Inventory 2281}}{{with index .Inventory 2281}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
                <div class="bag2023 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2023</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2282 slot slotdrop ui-droppable" slot-id="2282" slot-name="Bank Bag 2023-1"><span class="{{if index .Inventory 2282}}{{with index .Inventory 2282}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2284 slot slotdrop ui-droppable" slot-id="2284" slot-name="Bank Bag 2023-2"><span class="{{if index .Inventory 2284}}{{with index .Inventory 2284}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2285 slot slotdrop ui-droppable" slot-id="2285" slot-name="Bank Bag 2023-3"><span class="{{if index .Inventory 2285}}{{with index .Inventory 2285}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2286 slot slotdrop ui-droppable" slot-id="2286" slot-name="Bank Bag 2023-4"><span class="{{if index .Inventory 2286}}{{with index .Inventory 2286}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2287 slot slotdrop ui-droppable" slot-id="2287" slot-name="Bank Bag 2023-5"><span class="{{if index .Inventory 2287}}{{with index .Inventory 2287}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2288 slot slotdrop ui-droppable" slot-id="2288" slot-name="Bank Bag 2023-6"><span class="{{if index .Inventory 2288}}{{with index .Inventory 2288}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2289 slot slotdrop ui-droppable" slot-id="2289" slot-name="Bank Bag 2023-7"><span class="{{if index .Inventory 2289}}{{with index .Inventory 2289}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2290 slot slotdrop ui-droppable" slot-id="2290" slot-name="Bank Bag 2023-8"><span class="{{if index .Inventory 2290}}{{with index .Inventory 2290}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2291 slot slotdrop ui-droppable" slot-id="2291" slot-name="Bank Bag 2023-9"><span class="{{if index .Inventory 2291}}{{with index .Inventory 2291}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <span class="slot2292 slot slotdrop ui-droppable" slot-id="2292" slot-name="Bank Bag 2023-10"><span class="{{if index .Inventory 2292}}{{with index .Inventory 2292}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                        <br><br>
                    </div>
                </div>
                <div class="bag2024 panel panel-default" style="display: none">
                    <div class="panel-heading">Bag 2024</div>
                    <div class="panel-body" style="padding: 10px">
                        <span class="slot2293 slot slotdrop ui-droppable" slot-id="2293" slot-name="Bank Bag 2024-1"><span class="{{if index .Inventory 2293}}{{with index .Inventory 2293}}item icon-{{.Icon}}{{end}}{{end}}"></span></span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<div class="tab-pane fade" id="statistics">
    <div class="panel panel-default">
        <div class="panel-body">
            <div class="col-md-12 statistics-border">
                <div class="item-stats col-md-4" style="padding: 0px 5px">
                    <div class="panel panel-default">
                        <div class="panel-heading">Player Information</div>
                        <div class="panel-body">Name: Shin <br>
                            Level: 70<br>
                            Race: 6<br>
                            Class: Bard<br>
                            Deity: 207<br>
                            Zone: 158<br>
                            X: 1,122.55<br>
                            Y: 592.26<br>
                            Z: 63.13<br>
                        Heading: 129.63</div>
                    </div>
                </div>
                <div class="item-stats col-md-4" style="padding: 0px 5px">
                    <div class="panel panel-default">
                        <div class="panel-heading">Main Statistics</div>
                        <div class="panel-body">Health: 2,681<br>
                        Mana: 2,062<br>                          Endurance: 2,100<br></div>
                    </div>
                </div>
                <div class="item-stats col-md-4" style="padding: 0px 5px">
                    <div class="panel panel-default">
                        <div class="panel-heading">Combat Statistics</div>
                        <div class="panel-body">Agility: 95<br>
                            Charisma: 110<br>
                            Dexterity: 90<br>
                            Intelligence: 75<br>
                            Stamina: 65<br>
                            Strength: 70<br>
                        Wisdom: 80</div>
                    </div>
                </div>
            </div>
        </div>
    </div>   
</div>

</div>