{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-4">
        <div class="hpanel forum-box">            
            <div class="panel-heading">
                <span class="f"><a href="/item">Item</a> > {{.Item.Name}}</span>
            </div>

            <div class="panel-body">                
                <span class="slot slotdrop"><span alt="{{.Item.Name}}" class="item icon-{{.Item.Icon}}"></span></span>
            
            <br/>
            <br/>
        
                <h3>{{.Item.Name}}</h3>
            </div>
            <div class="panel-body">
                <table class="table table-striped">
                    <tbody>
                    {{if .Item.ID}}<tr><td>ID</td><td>{{.Item.ID}}</td></tr>{{end}}
                    {{if .Item.Name}}<tr><td>Name</td><td>{{.Item.Name}}</td></tr>{{end}}
                    {{if .Item.Aagi}}<tr><td>Aagi</td><td>{{.Item.Aagi}}</td></tr>{{end}}
                    {{if .Item.Ac}}<tr><td>Ac</td><td>{{.Item.Ac}}</td></tr>{{end}}
                    {{if .Item.Accuracy}}<tr><td>Accuracy</td><td>{{.Item.Accuracy}}</td></tr>{{end}}
                    {{if .Item.Acha}}<tr><td>Acha</td><td>{{.Item.Acha}}</td></tr>{{end}}
                    {{if .Item.Adex}}<tr><td>Adex</td><td>{{.Item.Adex}}</td></tr>{{end}}
                    {{if .Item.Aint}}<tr><td>Aint</td><td>{{.Item.Aint}}</td></tr>{{end}}
                    {{if .Item.Artifactflag}}<tr><td>Artifactflag</td><td>{{.Item.Artifactflag}}</td></tr>{{end}}
                    {{if .Item.Asta}}<tr><td>Asta</td><td>{{.Item.Asta}}</td></tr>{{end}}
                    {{if .Item.Astr}}<tr><td>Astr</td><td>{{.Item.Astr}}</td></tr>{{end}}
                    {{if .Item.Attack}}<tr><td>Attack</td><td>{{.Item.Attack}}</td></tr>{{end}}
                    {{if .Item.Augrestrict}}<tr><td>Augrestrict</td><td>{{.Item.Augrestrict}}</td></tr>{{end}}
                    {{if .Item.Augslot1type}}{{if .Item.Augslot1visible.Value}}
                    <tr><td>Augslot1type</td><td>{{.Item.Augslot1type}}</td></tr>
                    <tr><td>Augslot1visible</td><td>{{.Item.Augslot1visible.Value}}</td></tr>
                    {{end}}{{end}}
                    {{if .Item.Augslot2type}}{{if .Item.Augslot2visible.Value}}
                    <tr><td>Augslot2type</td><td>{{.Item.Augslot2type}}</td></tr>
                    <tr><td>Augslot2visible</td><td>{{.Item.Augslot2visible.Value}}</td></tr>
                    {{end}}{{end}}
                    {{if .Item.Augslot3type}}{{if .Item.Augslot3visible.Value}}
                    <tr><td>Augslot3type</td><td>{{.Item.Augslot3type}}</td></tr>
                    <tr><td>Augslot3visible</td><td>{{.Item.Augslot3visible.Value}}</td></tr>
                    {{end}}{{end}}
                    {{if .Item.Augslot4type}}{{if .Item.Augslot4visible.Value}}
                    <tr><td>Augslot4type</td><td>{{.Item.Augslot4type}}</td></tr>
                    <tr><td>Augslot4visible</td><td>{{.Item.Augslot4visible.Value}}</td></tr>
                    {{end}}{{end}}
                    {{if .Item.Augslot5type}}{{if .Item.Augslot5visible.Value}}
                    <tr><td>Augslot5type</td><td>{{.Item.Augslot5type}}</td></tr>
                    <tr><td>Augslot5visible</td><td>{{.Item.Augslot5visible.Value}}</td></tr>
                    {{end}}{{end}}
                    {{if .Item.Augslot6type}}{{if .Item.Augslot6visible.Value}}
                    <tr><td>Augslot6type</td><td>{{.Item.Augslot6type}}</td></tr>
                    <tr><td>Augslot6visible</td><td>{{.Item.Augslot6visible.Value}}</td></tr>
                    {{end}}{{end}}
                    {{if .Item.Augtype}}<tr><td>Augtype</td><td>{{.Item.Augtype}}</td></tr>{{end}}
                    {{if .Item.Avoidance}}<tr><td>Avoidance</td><td>{{.Item.Avoidance}}</td></tr>{{end}}
                    {{if .Item.Awis}}<tr><td>Awis</td><td>{{.Item.Awis}}</td></tr>{{end}}
                    {{if .Item.Bagsize}}<tr><td>Bagsize</td><td>{{.Item.Bagsize}}</td></tr>{{end}}
                    {{if .Item.Bagslots}}<tr><td>Bagslots</td><td>{{.Item.Bagslots}}</td></tr>{{end}}
                    {{if .Item.Bagtype}}<tr><td>Bagtype</td><td>{{.Item.Bagtype}}</td></tr>{{end}}
                    {{if .Item.Bagwr}}<tr><td>Bagwr</td><td>{{.Item.Bagwr}}</td></tr>{{end}}
                    {{if .Item.Banedmgamt}}<tr><td>Banedmgamt</td><td>{{.Item.Banedmgamt}}</td></tr>{{end}}
                    {{if .Item.Banedmgraceamt}}<tr><td>Banedmgraceamt</td><td>{{.Item.Banedmgraceamt}}</td></tr>{{end}}
                    {{if .Item.Banedmgbody}}<tr><td>Banedmgbody</td><td>{{.Item.Banedmgbody}}</td></tr>{{end}}
                    {{if .Item.Banedmgrace}}<tr><td>Banedmgrace</td><td>{{.Item.Banedmgrace}}</td></tr>{{end}}
                    {{if .Item.Bardtype}}<tr><td>Bardtype</td><td>{{.Item.Bardtype}}</td></tr>{{end}}
                    {{if .Item.Bardvalue}}<tr><td>Bardvalue</td><td>{{.Item.Bardvalue}}</td></tr>{{end}}
                    {{if .Item.Book}}<tr><td>Book</td><td>{{.Item.Book}}</td></tr>{{end}}
                    {{if .Item.Casttime}}<tr><td>Casttime</td><td>{{.Item.Casttime}}</td></tr>{{end}}
                    {{if .Item.Casttime2}}<tr><td>Casttime2</td><td>{{.Item.Casttime2}}</td></tr>{{end}}
                    {{if .Item.Charmfile}}<tr><td>Charmfile</td><td>{{.Item.Charmfile}}</td></tr>{{end}}
                    {{if ne .Item.Charmfileid "0"}}<tr><td>Charmfileid</td><td>{{.Item.Charmfileid}}</td></tr>{{end}}
                    {{if .Item.Classes}}<tr><td>Classes</td><td>{{.Item.ClassList}}</td></tr>{{end}}
                    {{if .Item.Color}}<tr><td>Color</td><td>{{.Item.Color}} <span style="color: rgba({{.Item.StyleColor}})">(Sample)</span> {{.Item.StyleColor}}</td></tr>{{end}}
                    {{if ne .Item.Combateffects "0"}}<tr><td>Combateffects</td><td>{{.Item.Combateffects}}</td></tr>{{end}}
                    {{if .Item.Extradmgskill}}<tr><td>Extradmgskill</td><td>{{.Item.Extradmgskill}}</td></tr>{{end}}
                    {{if .Item.Extradmgamt}}<tr><td>Extradmgamt</td><td>{{.Item.Extradmgamt}}</td></tr>{{end}}
                    {{if .Item.PriceName}}<tr><td>Price</td><td>{{.Item.PriceName}}</td></tr>{{end}}
                    {{if .Item.Cr}}<tr><td>Cr</td><td>{{.Item.Cr}}</td></tr>{{end}}
                    {{if .Item.Damage}}<tr><td>Damage</td><td>{{.Item.Damage}}</td></tr>{{end}}
                    {{if .Item.Damageshield}}<tr><td>Damageshield</td><td>{{.Item.Damageshield}}</td></tr>{{end}}
                    {{if .Item.Deity}}<tr><td>Deity</td><td>{{.Item.Deity}}</td></tr>{{end}}
                    {{if .Item.Delay}}<tr><td>Delay</td><td>{{.Item.Delay}}</td></tr>{{end}}
                    {{if .Item.Augdistiller}}<tr><td>Augdistiller</td><td>{{.Item.Augdistiller}}</td></tr>{{end}}
                    {{if .Item.Dotshielding}}<tr><td>Dotshielding</td><td>{{.Item.Dotshielding}}</td></tr>{{end}}
                    {{if .Item.Dr}}<tr><td>Dr</td><td>{{.Item.Dr}}</td></tr>{{end}}
                    {{if .Item.Clicktype}}<tr><td>Clicktype</td><td>{{.Item.Clicktype}}</td></tr>{{end}}
                    {{if .Item.Clicklevel2}}<tr><td>Clicklevel2</td><td>{{.Item.Clicklevel2}}</td></tr>{{end}}
                    {{if .Item.Elemdmgtype}}<tr><td>Elemdmgtype</td><td>{{.Item.Elemdmgtype}}</td></tr>{{end}}
                    {{if .Item.Elemdmgamt}}<tr><td>Elemdmgamt</td><td>{{.Item.Elemdmgamt}}</td></tr>{{end}}
                    {{if .Item.Endur}}<tr><td>Endur</td><td>{{.Item.Endur}}</td></tr>{{end}}
                    {{if .Item.Factionamt1}}<tr><td>Factionamt1</td><td>{{.Item.Factionamt1}}</td></tr>{{end}}
                    {{if .Item.Factionamt2}}<tr><td>Factionamt2</td><td>{{.Item.Factionamt2}}</td></tr>{{end}}
                    {{if .Item.Factionamt3}}<tr><td>Factionamt3</td><td>{{.Item.Factionamt3}}</td></tr>{{end}}
                    {{if .Item.Factionamt4}}<tr><td>Factionamt4</td><td>{{.Item.Factionamt4}}</td></tr>{{end}}
                    {{if .Item.Factionmod1}}<tr><td>Factionmod1</td><td>{{.Item.Factionmod1}}</td></tr>{{end}}
                    {{if .Item.Factionmod2}}<tr><td>Factionmod2</td><td>{{.Item.Factionmod2}}</td></tr>{{end}}
                    {{if .Item.Factionmod3}}<tr><td>Factionmod3</td><td>{{.Item.Factionmod3}}</td></tr>{{end}}
                    {{if .Item.Factionmod4}}<tr><td>Factionmod4</td><td>{{.Item.Factionmod4}}</td></tr>{{end}}
                    {{if .Item.Filename}}<tr><td>Filename</td><td>{{.Item.Filename}}</td></tr>{{end}}
                    {{if gt .Item.Focuseffect 0}}<tr><td>Focuseffect</td><td>{{.Item.Focuseffect}}</td></tr>{{end}}
                    {{if .Item.Fr}}<tr><td>Fr</td><td>{{.Item.Fr}}</td></tr>{{end}}
                    {{if .Item.Fvnodrop}}<tr><td>Fvnodrop</td><td>{{.Item.Fvnodrop}}</td></tr>{{end}}
                    {{if .Item.Haste}}<tr><td>Haste</td><td>{{.Item.Haste}}</td></tr>{{end}}
                    {{if .Item.Clicklevel}}<tr><td>Clicklevel</td><td>{{.Item.Clicklevel}}</td></tr>{{end}}
                    {{if .Item.Hp}}<tr><td>Hp</td><td>{{.Item.Hp}}</td></tr>{{end}}
                    {{if .Item.Regen}}<tr><td>Regen</td><td>{{.Item.Regen}}</td></tr>{{end}}
                    {{if .Item.Icon}}<tr><td>Icon</td><td>{{.Item.Icon}}</td></tr>{{end}}
                    {{if .Item.Idfile}}<tr><td>Idfile</td><td>{{.Item.Idfile}}</td></tr>{{end}}
                    {{if .Item.Itemclass}}<tr><td>Itemclass</td><td>{{.Item.Itemclass}}</td></tr>{{end}}
                    {{if .Item.Itemtype}}<tr><td>Itemtype</td><td>{{.Item.Itemtype}}</td></tr>{{end}}
                    {{if .Item.Ldonprice}}<tr><td>Ldonprice</td><td>{{.Item.Ldonprice}}</td></tr>{{end}}
                    {{if .Item.Ldontheme}}<tr><td>Ldontheme</td><td>{{.Item.Ldontheme}}</td></tr>{{end}}
                    {{if .Item.Ldonsold}}<tr><td>Ldonsold</td><td>{{.Item.Ldonsold}}</td></tr>{{end}}
                    {{if .Item.Light}}<tr><td>Light</td><td>{{.Item.Light}}</td></tr>{{end}}
                    {{if .Item.Lore}}<tr><td>Lore</td><td>{{.Item.Lore}}</td></tr>{{end}}
                    {{if gt .Item.Loregroup 0}}<tr><td>Loregroup</td><td>{{.Item.Loregroup}}</td></tr>{{end}}
                    {{if .Item.Magic}}<tr><td>Magic</td><td>{{.Item.Magic}}</td></tr>{{end}}
                    {{if .Item.Mana}}<tr><td>Mana</td><td>{{.Item.Mana}}</td></tr>{{end}}
                    {{if .Item.Manaregen}}<tr><td>Manaregen</td><td>{{.Item.Manaregen}}</td></tr>{{end}}
                    {{if .Item.Enduranceregen}}<tr><td>Enduranceregen</td><td>{{.Item.Enduranceregen}}</td></tr>{{end}}
                    {{if .Item.Material}}<tr><td>Material</td><td>{{.Item.Material}}</td></tr>{{end}}
                    {{if .Item.Herosforgemodel}}<tr><td>Herosforgemodel</td><td>{{.Item.Herosforgemodel}}</td></tr>{{end}}
                    {{if .Item.Maxcharges}}<tr><td>Maxcharges</td><td>{{.Item.Maxcharges}}</td></tr>{{end}}
                    {{if .Item.Mr}}<tr><td>Mr</td><td>{{.Item.Mr}}</td></tr>{{end}}
                    {{if .Item.Nodrop}}<tr><td>Nodrop</td><td>{{.Item.Nodrop}}</td></tr>{{end}}
                    {{if .Item.Norent}}<tr><td>Norent</td><td>{{.Item.Norent}}</td></tr>{{end}}
                    {{if .Item.Pendingloreflag}}<tr><td>Pendingloreflag</td><td>{{.Item.Pendingloreflag}}</td></tr>{{end}}
                    {{if .Item.Pr}}<tr><td>Pr</td><td>{{.Item.Pr}}</td></tr>{{end}}
                    {{if .Item.Procrate}}<tr><td>Procrate</td><td>{{.Item.Procrate}}</td></tr>{{end}}
                    {{if .Item.Races}}<tr><td>Races</td><td>{{.Item.RaceList}}</td></tr>{{end}}
                    {{if .Item.Range}}<tr><td>Range</td><td>{{.Item.Range}}</td></tr>{{end}}
                    {{if .Item.Reclevel}}<tr><td>Reclevel</td><td>{{.Item.Reclevel}}</td></tr>{{end}}
                    {{if .Item.Recskill}}<tr><td>Recskill</td><td>{{.Item.Recskill}}</td></tr>{{end}}
                    {{if .Item.Reqlevel}}<tr><td>Reqlevel</td><td>{{.Item.Reqlevel}}</td></tr>{{end}}
                    {{if .Item.SellrateName}}<tr><td>Sellrate</td><td>{{.Item.SellrateName}} ({{.Item.Sellrate}}x)</td></tr>{{end}}
                    {{if .Item.Shielding}}<tr><td>Shielding</td><td>{{.Item.Shielding}}</td></tr>{{end}}
                    {{if .Item.Size}}<tr><td>Size</td><td>{{.Item.Size}}</td></tr>{{end}}
                    {{if gt .Item.Skillmodtype 0}}<tr><td>Skillmodtype</td><td>{{.Item.Skillmodtype}}</td></tr>{{end}}
                    {{if .Item.Skillmodvalue}}<tr><td>Skillmodvalue</td><td>{{.Item.Skillmodvalue}}</td></tr>{{end}}
                    {{if .Item.SlotList}}<tr><td>Slots</td><td>{{.Item.SlotList}}</td></tr>{{end}}
                    {{if gt .Item.Clickeffect 0}}<tr><td>Clickeffect</td><td>{{.Item.Clickeffect}}</td></tr>{{end}}
                    {{if .Item.Spellshield}}<tr><td>Spellshield</td><td>{{.Item.Spellshield}}</td></tr>{{end}}
                    {{if .Item.Strikethrough}}<tr><td>Strikethrough</td><td>{{.Item.Strikethrough}}</td></tr>{{end}}
                    {{if .Item.Stunresist}}<tr><td>Stunresist</td><td>{{.Item.Stunresist}}</td></tr>{{end}}
                    {{if .Item.Summonedflag}}<tr><td>Summonedflag</td><td>{{.Item.Summonedflag}}</td></tr>{{end}}
                    {{if .Item.Tradeskills}}<tr><td>Tradeskills</td><td>{{.Item.Tradeskills}}</td></tr>{{end}}
                    {{if .Item.Favor}}<tr><td>Favor</td><td>{{.Item.Favor}}</td></tr>{{end}}
                    {{if .Item.Weight}}<tr><td>Weight</td><td>{{.Item.Weight}}</td></tr>{{end}}
                    {{if .Item.Unk012}}<tr><td>Unk012</td><td>{{.Item.Unk012}}</td></tr>{{end}}
                    {{if .Item.Unk013}}<tr><td>Unk013</td><td>{{.Item.Unk013}}</td></tr>{{end}}
                    {{if .Item.Benefitflag}}<tr><td>Benefitflag</td><td>{{.Item.Benefitflag}}</td></tr>{{end}}
                    {{if .Item.Unk054}}<tr><td>Unk054</td><td>{{.Item.Unk054}}</td></tr>{{end}}
                    {{if .Item.Unk059}}<tr><td>Unk059</td><td>{{.Item.Unk059}}</td></tr>{{end}}
                    {{if .Item.Booktype}}<tr><td>Booktype</td><td>{{.Item.Booktype}}</td></tr>{{end}}
                    {{if .Item.Recastdelay}}<tr><td>Recastdelay</td><td>{{.Item.Recastdelay}}</td></tr>{{end}}
                    {{if .Item.Recasttype}}<tr><td>Recasttype</td><td>{{.Item.Recasttype}}</td></tr>{{end}}
                    {{if .Item.Guildfavor}}<tr><td>Guildfavor</td><td>{{.Item.Guildfavor}}</td></tr>{{end}}
                    {{if .Item.Unk123}}<tr><td>Unk123</td><td>{{.Item.Unk123}}</td></tr>{{end}}
                    {{if .Item.Unk124}}<tr><td>Unk124</td><td>{{.Item.Unk124}}</td></tr>{{end}}
                    {{if .Item.Attuneable}}<tr><td>Attuneable</td><td>{{.Item.Attuneable}}</td></tr>{{end}}
                    {{if .Item.Nopet}}<tr><td>Nopet</td><td>{{.Item.Nopet}}</td></tr>{{end}}
                    {{if .Item.Updated.Time}}<tr><td>Updated</td><td>{{.Item.Updated.Time}}</td></tr>{{end}}
                    {{if .Item.Comment}}<tr><td>Comment</td><td>{{.Item.Comment}}</td></tr>{{end}}
                    {{if .Item.Unk127}}<tr><td>Unk127</td><td>{{.Item.Unk127}}</td></tr>{{end}}
                    {{if .Item.Pointtype}}<tr><td>Pointtype</td><td>{{.Item.Pointtype}}</td></tr>{{end}}
                    {{if .Item.Potionbelt}}<tr><td>Potionbelt</td><td>{{.Item.Potionbelt}}</td></tr>{{end}}
                    {{if .Item.Potionbeltslots}}<tr><td>Potionbeltslots</td><td>{{.Item.Potionbeltslots}}</td></tr>{{end}}
                    {{if ne .Item.Stacksize -1}}<tr><td>Stacksize</td><td>{{.Item.Stacksize}}</td></tr>{{end}}
                    {{if .Item.Notransfer}}<tr><td>Notransfer</td><td>{{.Item.Notransfer}}</td></tr>{{end}}
                    {{if .Item.Stackable}}<tr><td>Stackable</td><td>{{.Item.Stackable}}</td></tr>{{end}}
                    {{if .Item.Unk134}}<tr><td>Unk134</td><td>{{.Item.Unk134}}</td></tr>{{end}}
                    {{if .Item.Unk137}}<tr><td>Unk137</td><td>{{.Item.Unk137}}</td></tr>{{end}}
                    {{if ne .Item.Proceffect -1}}<tr><td>Proceffect</td><td>{{.Item.Proceffect}}</td></tr>{{end}}
                    {{if .Item.Proctype}}<tr><td>Proctype</td><td>{{.Item.Proctype}}</td></tr>{{end}}
                    {{if .Item.Proclevel2}}<tr><td>Proclevel2</td><td>{{.Item.Proclevel2}}</td></tr>{{end}}
                    {{if .Item.Proclevel}}<tr><td>Proclevel</td><td>{{.Item.Proclevel}}</td></tr>{{end}}
                    {{if .Item.Unk142}}<tr><td>Unk142</td><td>{{.Item.Unk142}}</td></tr>{{end}}
                    {{if ne .Item.Worneffect -1}}<tr><td>Worneffect</td><td>{{.Item.Worneffect}}</td></tr>{{end}}
                    {{if .Item.Worntype}}<tr><td>Worntype</td><td>{{.Item.Worntype}}</td></tr>{{end}}
                    {{if .Item.Wornlevel2}}<tr><td>Wornlevel2</td><td>{{.Item.Wornlevel2}}</td></tr>{{end}}
                    {{if .Item.Wornlevel}}<tr><td>Wornlevel</td><td>{{.Item.Wornlevel}}</td></tr>{{end}}
                    {{if .Item.Unk147}}<tr><td>Unk147</td><td>{{.Item.Unk147}}</td></tr>{{end}}
                    {{if .Item.Focustype}}<tr><td>Focustype</td><td>{{.Item.Focustype}}</td></tr>{{end}}
                    {{if .Item.Focuslevel2}}<tr><td>Focuslevel2</td><td>{{.Item.Focuslevel2}}</td></tr>{{end}}
                    {{if .Item.Focuslevel}}<tr><td>Focuslevel</td><td>{{.Item.Focuslevel}}</td></tr>{{end}}
                    {{if .Item.Unk152}}<tr><td>Unk152</td><td>{{.Item.Unk152}}</td></tr>{{end}}
                    {{if ne .Item.Scrolleffect -1}}<tr><td>Scrolleffect</td><td>{{.Item.Scrolleffect}}</td></tr>{{end}}
                    {{if .Item.Scrolltype}}<tr><td>Scrolltype</td><td>{{.Item.Scrolltype}}</td></tr>{{end}}
                    {{if .Item.Scrolllevel2}}<tr><td>Scrolllevel2</td><td>{{.Item.Scrolllevel2}}</td></tr>{{end}}
                    {{if .Item.Scrolllevel}}<tr><td>Scrolllevel</td><td>{{.Item.Scrolllevel}}</td></tr>{{end}}
                    {{if .Item.Unk157}}<tr><td>Unk157</td><td>{{.Item.Unk157}}</td></tr>{{end}}
                    {{if .Item.Serialized.Time}}<tr><td>Serialized</td><td>{{.Item.Serialized.Time}}</td></tr>{{end}}
                    {{if .Item.Verified.Time}}<tr><td>Verified</td><td>{{.Item.Verified.Time}}</td></tr>{{end}}
                    {{if .Item.Serialization.String}}<tr><td>Serialization</td><td>{{.Item.Serialization.String}}</td></tr>{{end}}
                    {{if .Item.Source}}<tr><td>Source</td><td>{{.Item.Source}}</td></tr>{{end}}
                    {{if .Item.Unk033}}<tr><td>Unk033</td><td>{{.Item.Unk033}}</td></tr>{{end}}
                    {{if .Item.Lorefile}}<tr><td>Lorefile</td><td>{{.Item.Lorefile}}</td></tr>{{end}}
                    {{if .Item.Unk014}}<tr><td>Unk014</td><td>{{.Item.Unk014}}</td></tr>{{end}}
                    {{if .Item.Svcorruption}}<tr><td>Svcorruption</td><td>{{.Item.Svcorruption}}</td></tr>{{end}}
                    {{if .Item.Skillmodmax}}<tr><td>Skillmodmax</td><td>{{.Item.Skillmodmax}}</td></tr>{{end}}
                    {{if .Item.Unk060}}<tr><td>Unk060</td><td>{{.Item.Unk060}}</td></tr>{{end}}
                    {{if .Item.Augslot1unk2}}<tr><td>Augslot1unk2</td><td>{{.Item.Augslot1unk2}}</td></tr>{{end}}
                    {{if .Item.Augslot2unk2}}<tr><td>Augslot2unk2</td><td>{{.Item.Augslot2unk2}}</td></tr>{{end}}
                    {{if .Item.Augslot3unk2}}<tr><td>Augslot3unk2</td><td>{{.Item.Augslot3unk2}}</td></tr>{{end}}
                    {{if .Item.Augslot4unk2}}<tr><td>Augslot4unk2</td><td>{{.Item.Augslot4unk2}}</td></tr>{{end}}
                    {{if .Item.Augslot5unk2}}<tr><td>Augslot5unk2</td><td>{{.Item.Augslot5unk2}}</td></tr>{{end}}
                    {{if .Item.Augslot6unk2}}<tr><td>Augslot6unk2</td><td>{{.Item.Augslot6unk2}}</td></tr>{{end}}
                    {{if gt .Item.Unk120 0 }}<tr><td>Unk120</td><td>{{.Item.Unk120}}</td></tr>{{end}}
                    {{if .Item.Unk121}}<tr><td>Unk121</td><td>{{.Item.Unk121}}</td></tr>{{end}}
                    {{if .Item.Questitemflag}}<tr><td>Questitemflag</td><td>{{.Item.Questitemflag}}</td></tr>{{end}}
                    
                    {{if gt .Item.Clickunk5 0}}<tr><td>Clickunk5</td><td>{{.Item.Clickunk5}}</td></tr>{{end}}
                    {{if ne .Item.Clickunk6 ""}}<tr><td>Clickunk6</td><td>{{.Item.Clickunk6}}</td></tr>{{end}}
                    {{if ne .Item.Clickunk7 -1}}<tr><td>Clickunk7</td><td>{{.Item.Clickunk7}}</td></tr>{{end}}
                    {{if .Item.Procunk1}}<tr><td>Procunk1</td><td>{{.Item.Procunk1}}</td></tr>{{end}}
                    {{if .Item.Procunk2}}<tr><td>Procunk2</td><td>{{.Item.Procunk2}}</td></tr>{{end}}
                    {{if .Item.Procunk3}}<tr><td>Procunk3</td><td>{{.Item.Procunk3}}</td></tr>{{end}}
                    {{if .Item.Procunk4}}<tr><td>Procunk4</td><td>{{.Item.Procunk4}}</td></tr>{{end}}
                    {{if .Item.Procunk6}}<tr><td>Procunk6</td><td>{{.Item.Procunk6}}</td></tr>{{end}}
                    {{if .Item.Procunk7}}<tr><td>Procunk7</td><td>{{.Item.Procunk7}}</td></tr>{{end}}
                    {{if .Item.Wornunk1}}<tr><td>Wornunk1</td><td>{{.Item.Wornunk1}}</td></tr>{{end}}
                    {{if .Item.Wornunk2}}<tr><td>Wornunk2</td><td>{{.Item.Wornunk2}}</td></tr>{{end}}
                    {{if .Item.Wornunk3}}<tr><td>Wornunk3</td><td>{{.Item.Wornunk3}}</td></tr>{{end}}
                    {{if .Item.Wornunk4}}<tr><td>Wornunk4</td><td>{{.Item.Wornunk4}}</td></tr>{{end}}
                    {{if .Item.Wornunk5}}<tr><td>Wornunk5</td><td>{{.Item.Wornunk5}}</td></tr>{{end}}
                    {{if .Item.Wornunk6}}<tr><td>Wornunk6</td><td>{{.Item.Wornunk6}}</td></tr>{{end}}
                    {{if .Item.Wornunk7}}<tr><td>Wornunk7</td><td>{{.Item.Wornunk7}}</td></tr>{{end}}
                    {{if .Item.Focusunk1}}<tr><td>Focusunk1</td><td>{{.Item.Focusunk1}}</td></tr>{{end}}
                    {{if .Item.Focusunk2}}<tr><td>Focusunk2</td><td>{{.Item.Focusunk2}}</td></tr>{{end}}
                    {{if .Item.Focusunk3}}<tr><td>Focusunk3</td><td>{{.Item.Focusunk3}}</td></tr>{{end}}
                    {{if .Item.Focusunk4}}<tr><td>Focusunk4</td><td>{{.Item.Focusunk4}}</td></tr>{{end}}
                    {{if .Item.Focusunk5}}<tr><td>Focusunk5</td><td>{{.Item.Focusunk5}}</td></tr>{{end}}
                    {{if .Item.Focusunk6}}<tr><td>Focusunk6</td><td>{{.Item.Focusunk6}}</td></tr>{{end}}
                    {{if .Item.Focusunk7}}<tr><td>Focusunk7</td><td>{{.Item.Focusunk7}}</td></tr>{{end}}
                    {{if .Item.Scrollunk1}}<tr><td>Scrollunk1</td><td>{{.Item.Scrollunk1}}</td></tr>{{end}}
                    {{if .Item.Scrollunk2}}<tr><td>Scrollunk2</td><td>{{.Item.Scrollunk2}}</td></tr>{{end}}
                    {{if .Item.Scrollunk3}}<tr><td>Scrollunk3</td><td>{{.Item.Scrollunk3}}</td></tr>{{end}}
                    {{if .Item.Scrollunk4}}<tr><td>Scrollunk4</td><td>{{.Item.Scrollunk4}}</td></tr>{{end}}
                    {{if .Item.Scrollunk5}}<tr><td>Scrollunk5</td><td>{{.Item.Scrollunk5}}</td></tr>{{end}}
                    {{if .Item.Scrollunk6}}<tr><td>Scrollunk6</td><td>{{.Item.Scrollunk6}}</td></tr>{{end}}
                    {{if .Item.Scrollunk7}}<tr><td>Scrollunk7</td><td>{{.Item.Scrollunk7}}</td></tr>{{end}}
                    {{if .Item.Unk193}}<tr><td>Unk193</td><td>{{.Item.Unk193}}</td></tr>{{end}}
                    {{if .Item.Purity}}<tr><td>Purity</td><td>{{.Item.Purity}}</td></tr>{{end}}
                    {{if .Item.Evoitem}}<tr><td>Evoitem</td><td>{{.Item.Evoitem}}</td></tr>{{end}}
                    {{if .Item.Evoid}}<tr><td>Evoid</td><td>{{.Item.Evoid}}</td></tr>{{end}}
                    {{if .Item.Evolvinglevel}}<tr><td>Evolvinglevel</td><td>{{.Item.Evolvinglevel}}</td></tr>{{end}}
                    {{if .Item.Evomax}}<tr><td>Evomax</td><td>{{.Item.Evomax}}</td></tr>{{end}}
                    {{if .Item.Clickname}}<tr><td>Clickname</td><td>{{.Item.Clickname}}</td></tr>{{end}}
                    {{if .Item.Procname}}<tr><td>Procname</td><td>{{.Item.Procname}}</td></tr>{{end}}
                    {{if .Item.Wornname}}<tr><td>Wornname</td><td>{{.Item.Wornname}}</td></tr>{{end}}
                    {{if .Item.Focusname}}<tr><td>Focusname</td><td>{{.Item.Focusname}}</td></tr>{{end}}
                    {{if .Item.Scrollname}}<tr><td>Scrollname</td><td>{{.Item.Scrollname}}</td></tr>{{end}}
                    {{if .Item.Dsmitigation}}<tr><td>Dsmitigation</td><td>{{.Item.Dsmitigation}}</td></tr>{{end}}
                    {{if .Item.HeroicStr}}<tr><td>HeroicStr</td><td>{{.Item.HeroicStr}}</td></tr>{{end}}
                    {{if .Item.HeroicInt}}<tr><td>HeroicInt</td><td>{{.Item.HeroicInt}}</td></tr>{{end}}
                    {{if .Item.HeroicWis}}<tr><td>HeroicWis</td><td>{{.Item.HeroicWis}}</td></tr>{{end}}
                    {{if .Item.HeroicAgi}}<tr><td>HeroicAgi</td><td>{{.Item.HeroicAgi}}</td></tr>{{end}}
                    {{if .Item.HeroicDex}}<tr><td>HeroicDex</td><td>{{.Item.HeroicDex}}</td></tr>{{end}}
                    {{if .Item.HeroicSta}}<tr><td>HeroicSta</td><td>{{.Item.HeroicSta}}</td></tr>{{end}}
                    {{if .Item.HeroicCha}}<tr><td>HeroicCha</td><td>{{.Item.HeroicCha}}</td></tr>{{end}}
                    {{if .Item.HeroicPr}}<tr><td>HeroicPr</td><td>{{.Item.HeroicPr}}</td></tr>{{end}}
                    {{if .Item.HeroicDr}}<tr><td>HeroicDr</td><td>{{.Item.HeroicDr}}</td></tr>{{end}}
                    {{if .Item.HeroicFr}}<tr><td>HeroicFr</td><td>{{.Item.HeroicFr}}</td></tr>{{end}}
                    {{if .Item.HeroicCr}}<tr><td>HeroicCr</td><td>{{.Item.HeroicCr}}</td></tr>{{end}}
                    {{if .Item.HeroicMr}}<tr><td>HeroicMr</td><td>{{.Item.HeroicMr}}</td></tr>{{end}}
                    {{if .Item.HeroicSvcorrup}}<tr><td>HeroicSvcorrup</td><td>{{.Item.HeroicSvcorrup}}</td></tr>{{end}}
                    {{if .Item.Healamt}}<tr><td>Healamt</td><td>{{.Item.Healamt}}</td></tr>{{end}}
                    {{if .Item.Spelldmg}}<tr><td>Spelldmg</td><td>{{.Item.Spelldmg}}</td></tr>{{end}}
                    {{if .Item.Clairvoyance}}<tr><td>Clairvoyance</td><td>{{.Item.Clairvoyance}}</td></tr>{{end}}
                    {{if .Item.Backstabdmg}}<tr><td>Backstabdmg</td><td>{{.Item.Backstabdmg}}</td></tr>{{end}}
                    {{if .Item.Created}}<tr><td>Created</td><td>{{.Item.Created}}</td></tr>{{end}}
                    {{if .Item.Elitematerial}}<tr><td>Elitematerial</td><td>{{.Item.Elitematerial}}</td></tr>{{end}}
                    {{if .Item.Ldonsellbackrate}}<tr><td>Ldonsellbackrate</td><td>{{.Item.Ldonsellbackrate}}</td></tr>{{end}}
                    {{if .Item.Scriptfileid}}<tr><td>Scriptfileid</td><td>{{.Item.Scriptfileid}}</td></tr>{{end}}
                    {{if .Item.Expendablearrow}}<tr><td>Expendablearrow</td><td>{{.Item.Expendablearrow}}</td></tr>{{end}}
                    {{if .Item.Powersourcecapacity}}<tr><td>Powersourcecapacity</td><td>{{.Item.Powersourcecapacity}}</{{end}}</td></tr>
                    {{if .Item.Bardeffect}}<tr><td>Bardeffect</td><td>{{.Item.Bardeffect}}</td></tr>{{end}}
                    {{if .Item.Bardeffecttype}}<tr><td>Bardeffecttype</td><td>{{.Item.Bardeffecttype}}</td></tr>{{end}}
                    {{if .Item.Bardlevel2}}<tr><td>Bardlevel2</td><td>{{.Item.Bardlevel2}}</td></tr>{{end}}
                    {{if .Item.Bardlevel}}<tr><td>Bardlevel</td><td>{{.Item.Bardlevel}}</td></tr>{{end}}
                    {{if .Item.Bardunk1}}<tr><td>Bardunk1</td><td>{{.Item.Bardunk1}}</td></tr>{{end}}
                    {{if .Item.Bardunk2}}<tr><td>Bardunk2</td><td>{{.Item.Bardunk2}}</td></tr>{{end}}
                    {{if .Item.Bardunk3}}<tr><td>Bardunk3</td><td>{{.Item.Bardunk3}}</td></tr>{{end}}
                    {{if .Item.Bardunk4}}<tr><td>Bardunk4</td><td>{{.Item.Bardunk4}}</td></tr>{{end}}
                    {{if .Item.Bardunk5}}<tr><td>Bardunk5</td><td>{{.Item.Bardunk5}}</td></tr>{{end}}
                    {{if .Item.Bardname}}<tr><td>Bardname</td><td>{{.Item.Bardname}}</td></tr>{{end}}
                    {{if .Item.Bardunk7}}<tr><td>Bardunk7</td><td>{{.Item.Bardunk7}}</td></tr>{{end}}
                    {{if .Item.Unk214}}<tr><td>Unk214</td><td>{{.Item.Unk214}}</td></tr>{{end}}
                    {{if .Item.Unk219}}<tr><td>Unk219</td><td>{{.Item.Unk219}}</td></tr>{{end}}
                    {{if .Item.Unk220}}<tr><td>Unk220</td><td>{{.Item.Unk220}}</td></tr>{{end}}
                    {{if .Item.Unk221}}<tr><td>Unk221</td><td>{{.Item.Unk221}}</td></tr>{{end}}
                    {{if .Item.Heirloom}}<tr><td>Heirloom</td><td>{{.Item.Heirloom}}</td></tr>{{end}}
                    {{if .Item.Unk223}}<tr><td>Unk223</td><td>{{.Item.Unk223}}</td></tr>{{end}}
                    {{if .Item.Unk224}}<tr><td>Unk224</td><td>{{.Item.Unk224}}</td></tr>{{end}}
                    {{if .Item.Unk225}}<tr><td>Unk225</td><td>{{.Item.Unk225}}</td></tr>{{end}}
                    {{if .Item.Unk226}}<tr><td>Unk226</td><td>{{.Item.Unk226}}</td></tr>{{end}}
                    {{if .Item.Unk227}}<tr><td>Unk227</td><td>{{.Item.Unk227}}</td></tr>{{end}}
                    {{if .Item.Unk228}}<tr><td>Unk228</td><td>{{.Item.Unk228}}</td></tr>{{end}}
                    {{if .Item.Unk229}}<tr><td>Unk229</td><td>{{.Item.Unk229}}</td></tr>{{end}}
                    {{if .Item.Unk230}}<tr><td>Unk230</td><td>{{.Item.Unk230}}</td></tr>{{end}}
                    {{if .Item.Unk231}}<tr><td>Unk231</td><td>{{.Item.Unk231}}</td></tr>{{end}}
                    {{if .Item.Unk232}}<tr><td>Unk232</td><td>{{.Item.Unk232}}</td></tr>{{end}}
                    {{if .Item.Unk233}}<tr><td>Unk233</td><td>{{.Item.Unk233}}</td></tr>{{end}}
                    {{if .Item.Unk234}}<tr><td>Unk234</td><td>{{.Item.Unk234}}</td></tr>{{end}}
                    {{if .Item.Placeable}}<tr><td>Placeable</td><td>{{.Item.Placeable}}</td></tr>{{end}}
                    {{if .Item.Unk236}}<tr><td>Unk236</td><td>{{.Item.Unk236}}</td></tr>{{end}}
                    {{if .Item.Unk237}}<tr><td>Unk237</td><td>{{.Item.Unk237}}</td></tr>{{end}}
                    {{if .Item.Unk238}}<tr><td>Unk238</td><td>{{.Item.Unk238}}</td></tr>{{end}}
                    {{if .Item.Unk239}}<tr><td>Unk239</td><td>{{.Item.Unk239}}</td></tr>{{end}}
                    {{if .Item.Unk240}}<tr><td>Unk240</td><td>{{.Item.Unk240}}</td></tr>{{end}}
                    {{if .Item.Unk241}}<tr><td>Unk241</td><td>{{.Item.Unk241}}</td></tr>{{end}}
                    {{if .Item.Epicitem}}<tr><td>Epicitem</td><td>{{.Item.Epicitem}}</td></tr>{{end}}
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
    {{if .Npcs}}
    <div class="col-lg-8">
        <div class="hpanel">
            <div class="panel-heading hbuilt">
                <div class="panel-tools">
                    <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                    <a class="closebox"><i class="fa fa-times"></i></a>
                </div>
                This item is found off mobs:
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table table-striped">
                    <thead>
                    <tr>
                        <th width="10px"><i title="Race" class="xa xa-bear"></i></th>
                        <th width="10px"><i title="Class" class="xa xa-all-for-one"></i></th>
                        <th>Name</th>
                        <th>Zone</th>
                        
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Npcs}}
                    <tr>
                        <td><i title="{{$value.RaceName}}" class="xa {{$value.RaceIcon}}"></i></td>
                                <td><i title="{{$value.ClassName}}" class="xa {{$value.ClassIcon}}"></i></td>
                        <td><a href="/npc/{{$value.ID}}">{{$value.Name}}</a></td>
                        <td><a href="/zone/{{$value.ZoneName}}">{{$value.ZoneName}}</a></td>
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>

            </div>
            <div class="panel-footer">
               {{len .Npcs}} total creatures
            </div>
        </div>
    </div>
    {{end}}
    {{if .Fishings}}
    <div class="col-lg-8">
        <div class="hpanel">
            <div class="panel-heading hbuilt">
                <div class="panel-tools">
                    <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                    <a class="closebox"><i class="fa fa-times"></i></a>
                </div>
                This item is found by fishing
            </div>
            <div class="panel-body">
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table table-striped">
                    <thead>
                    <tr>
                        <th>Zone</th>                        
                    </tr>
                    </thead>
                    <tbody>
                    {{range $key, $value := .Fishings}}
                    <tr>
                       {{if $value.ZoneID}}<td><a href="/fishing/{{$value.ID}}">{{$value.Zone.Name}}</a></td><td>Unknown Zone for {{$value.ID}}</td>{{else}}{{end}}
                    </tr>
                    {{end}}                
                    </tbody>
                </table>
                </div>

            </div>
            <div class="panel-footer">
               {{len .Npcs}} total creatures
            </div>
        </div>
    </div>
    {{end}}
</div>

</div>

</div>