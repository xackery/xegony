{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-4">
        <div class="hpanel forum-box">            
            <div class="panel-heading">
                <span class="f"><a href="/spell">Spell</a> > {{.Spell.Name.String}}</span>
            </div>

            <div class="panel-body">                
                <span class="slot slotdrop"><span alt="{{.Spell.Name.String}}" class="item icon-{{.Spell.Icon}}"></span></span>            
            <br/>
            <br/>
        
                <h3>{{.Spell.Name.String}}</h3>
            </div>
            <div class="panel-body">
                <table class="table table-striped">
                    <tbody>
                    {{if .Spell.ID}}<tr><td>ID</td><td>{{.Spell.ID}}</td>{{end}}
                    {{if .Spell.Name}}<tr><td>Name</td><td>{{.Spell.Name.String}}</td>{{end}}
                    {{if .Spell.Player1}}<tr><td>Player1</td><td>{{.Spell.Player1.String}}</td>{{end}}
                    {{if .Spell.TeleportZone.String}}<tr><td>TeleportZone</td><td>{{.Spell.TeleportZone.String}}</td>{{end}}
                    {{if .Spell.YouCast.String}}<tr><td>YouCast</td><td>{{.Spell.YouCast.String}}</td>{{end}}
                    {{if .Spell.OtherCasts.String}}<tr><td>OtherCasts</td><td>{{.Spell.OtherCasts.String}}</td>{{end}}
                    {{if .Spell.CastOnYou.String}}<tr><td>CastOnYou</td><td>{{.Spell.CastOnYou.String}}</td>{{end}}
                    {{if .Spell.CastOnOther.String}}<tr><td>CastOnOther</td><td>{{.Spell.CastOnOther.String}}</td>{{end}}
                    {{if .Spell.SpellFades.String}}<tr><td>SpellFades</td><td>{{.Spell.SpellFades.String}}</td>{{end}}
                    {{if .Spell.Range}}<tr><td>Range</td><td>{{.Spell.Range}}</td>{{end}}
                    {{if .Spell.Aoerange}}<tr><td>Aoerange</td><td>{{.Spell.Aoerange}}</td>{{end}}
                    {{if .Spell.Pushback}}<tr><td>Pushback</td><td>{{.Spell.Pushback}}</td>{{end}}
                    {{if .Spell.Pushup}}<tr><td>Pushup</td><td>{{.Spell.Pushup}}</td>{{end}}
                    {{if .Spell.CastTime}}<tr><td>CastTime</td><td>{{.Spell.CastTime}}</td>{{end}}
                    {{if .Spell.RecoveryTime}}<tr><td>RecoveryTime</td><td>{{.Spell.RecoveryTime}}</td>{{end}}
                    {{if .Spell.RecastTime}}<tr><td>RecastTime</td><td>{{.Spell.RecastTime}}</td>{{end}}
                    {{if .Spell.Buffdurationformula}}<tr><td>Buffdurationformula</td><td>{{.Spell.Buffdurationformula}}</td>{{end}}
                    {{if .Spell.Buffduration}}<tr><td>Buffduration</td><td>{{.Spell.Buffduration}}</td>{{end}}
                    {{if .Spell.Aeduration}}<tr><td>Aeduration</td><td>{{.Spell.Aeduration}}</td>{{end}}
                    {{if .Spell.Mana}}<tr><td>Mana</td><td>{{.Spell.Mana}}</td>{{end}}
                    {{if .Spell.EffectBaseValue1}}<tr><td>EffectBaseValue1</td><td>{{.Spell.EffectBaseValue1}}</td>{{end}}
                    {{if .Spell.EffectBaseValue2}}<tr><td>EffectBaseValue2</td><td>{{.Spell.EffectBaseValue2}}</td>{{end}}
                    {{if .Spell.EffectBaseValue3}}<tr><td>EffectBaseValue3</td><td>{{.Spell.EffectBaseValue3}}</td>{{end}}
                    {{if .Spell.EffectBaseValue4}}<tr><td>EffectBaseValue4</td><td>{{.Spell.EffectBaseValue4}}</td>{{end}}
                    {{if .Spell.EffectBaseValue5}}<tr><td>EffectBaseValue5</td><td>{{.Spell.EffectBaseValue5}}</td>{{end}}
                    {{if .Spell.EffectBaseValue6}}<tr><td>EffectBaseValue6</td><td>{{.Spell.EffectBaseValue6}}</td>{{end}}
                    {{if .Spell.EffectBaseValue7}}<tr><td>EffectBaseValue7</td><td>{{.Spell.EffectBaseValue7}}</td>{{end}}
                    {{if .Spell.EffectBaseValue8}}<tr><td>EffectBaseValue8</td><td>{{.Spell.EffectBaseValue8}}</td>{{end}}
                    {{if .Spell.EffectBaseValue9}}<tr><td>EffectBaseValue9</td><td>{{.Spell.EffectBaseValue9}}</td>{{end}}
                    {{if .Spell.EffectBaseValue10}}<tr><td>EffectBaseValue10</td><td>{{.Spell.EffectBaseValue10}}</td>{{end}}
                    {{if .Spell.EffectBaseValue11}}<tr><td>EffectBaseValue11</td><td>{{.Spell.EffectBaseValue11}}</td>{{end}}
                    {{if .Spell.EffectBaseValue12}}<tr><td>EffectBaseValue12</td><td>{{.Spell.EffectBaseValue12}}</td>{{end}}
                    {{if .Spell.EffectLimitValue1}}<tr><td>EffectLimitValue1</td><td>{{.Spell.EffectLimitValue1}}</td>{{end}}
                    {{if .Spell.EffectLimitValue2}}<tr><td>EffectLimitValue2</td><td>{{.Spell.EffectLimitValue2}}</td>{{end}}
                    {{if .Spell.EffectLimitValue3}}<tr><td>EffectLimitValue3</td><td>{{.Spell.EffectLimitValue3}}</td>{{end}}
                    {{if .Spell.EffectLimitValue4}}<tr><td>EffectLimitValue4</td><td>{{.Spell.EffectLimitValue4}}</td>{{end}}
                    {{if .Spell.EffectLimitValue5}}<tr><td>EffectLimitValue5</td><td>{{.Spell.EffectLimitValue5}}</td>{{end}}
                    {{if .Spell.EffectLimitValue6}}<tr><td>EffectLimitValue6</td><td>{{.Spell.EffectLimitValue6}}</td>{{end}}
                    {{if .Spell.EffectLimitValue7}}<tr><td>EffectLimitValue7</td><td>{{.Spell.EffectLimitValue7}}</td>{{end}}
                    {{if .Spell.EffectLimitValue8}}<tr><td>EffectLimitValue8</td><td>{{.Spell.EffectLimitValue8}}</td>{{end}}
                    {{if .Spell.EffectLimitValue9}}<tr><td>EffectLimitValue9</td><td>{{.Spell.EffectLimitValue9}}</td>{{end}}
                    {{if .Spell.EffectLimitValue10}}<tr><td>EffectLimitValue10</td><td>{{.Spell.EffectLimitValue10}}</td>{{end}}
                    {{if .Spell.EffectLimitValue11}}<tr><td>EffectLimitValue11</td><td>{{.Spell.EffectLimitValue11}}</td>{{end}}
                    {{if .Spell.EffectLimitValue12}}<tr><td>EffectLimitValue12</td><td>{{.Spell.EffectLimitValue12}}</td>{{end}}
                    {{if .Spell.Max1}}<tr><td>Max1</td><td>{{.Spell.Max1}}</td>{{end}}
                    {{if .Spell.Max2}}<tr><td>Max2</td><td>{{.Spell.Max2}}</td>{{end}}
                    {{if .Spell.Max3}}<tr><td>Max3</td><td>{{.Spell.Max3}}</td>{{end}}
                    {{if .Spell.Max4}}<tr><td>Max4</td><td>{{.Spell.Max4}}</td>{{end}}
                    {{if .Spell.Max5}}<tr><td>Max5</td><td>{{.Spell.Max5}}</td>{{end}}
                    {{if .Spell.Max6}}<tr><td>Max6</td><td>{{.Spell.Max6}}</td>{{end}}
                    {{if .Spell.Max7}}<tr><td>Max7</td><td>{{.Spell.Max7}}</td>{{end}}
                    {{if .Spell.Max8}}<tr><td>Max8</td><td>{{.Spell.Max8}}</td>{{end}}
                    {{if .Spell.Max9}}<tr><td>Max9</td><td>{{.Spell.Max9}}</td>{{end}}
                    {{if .Spell.Max10}}<tr><td>Max10</td><td>{{.Spell.Max10}}</td>{{end}}
                    {{if .Spell.Max11}}<tr><td>Max11</td><td>{{.Spell.Max11}}</td>{{end}}
                    {{if .Spell.Max12}}<tr><td>Max12</td><td>{{.Spell.Max12}}</td>{{end}}
                    {{if .Spell.Icon}}<tr><td>Icon</td><td>{{.Spell.Icon}}</td>{{end}}
                    {{if .Spell.Memicon}}<tr><td>Memicon</td><td>{{.Spell.Memicon}}</td>{{end}}
                    {{if gt .Spell.Components1 0}}<tr><td>Components1</td><td>{{.Spell.Components1}}</td>{{end}}
                    {{if gt .Spell.Components2 0}}<tr><td>Components2</td><td>{{.Spell.Components2}}</td>{{end}}
                    {{if gt .Spell.Components3 0}}<tr><td>Components3</td><td>{{.Spell.Components3}}</td>{{end}}
                    {{if gt .Spell.Components4 0}}<tr><td>Components4</td><td>{{.Spell.Components4}}</td>{{end}}
                    {{if gt .Spell.Components1 0}}{{if .Spell.ComponentCounts1}}<tr><td>ComponentCounts1</td><td>{{.Spell.ComponentCounts1}}</td>{{end}}{{end}}
                    {{if gt .Spell.Components2 0}}{{if .Spell.ComponentCounts2}}<tr><td>ComponentCounts2</td><td>{{.Spell.ComponentCounts2}}</td>{{end}}{{end}}
                    {{if gt .Spell.Components3 0}}{{if .Spell.ComponentCounts3}}<tr><td>ComponentCounts3</td><td>{{.Spell.ComponentCounts3}}</td>{{end}}{{end}}
                    {{if gt .Spell.Components4 0}}{{if .Spell.ComponentCounts4}}<tr><td>ComponentCounts4</td><td>{{.Spell.ComponentCounts4}}</td>{{end}}{{end}}
                    {{if gt .Spell.Noexpendreagent1 0}}<tr><td>Noexpendreagent1</td><td>{{.Spell.Noexpendreagent1}}</td>{{end}}
                    {{if gt .Spell.Noexpendreagent2 0}}<tr><td>Noexpendreagent2</td><td>{{.Spell.Noexpendreagent2}}</td>{{end}}
                    {{if gt .Spell.Noexpendreagent3 0}}<tr><td>Noexpendreagent3</td><td>{{.Spell.Noexpendreagent3}}</td>{{end}}
                    {{if gt .Spell.Noexpendreagent4 0}}<tr><td>Noexpendreagent4</td><td>{{.Spell.Noexpendreagent4}}</td>{{end}}
                    {{if .Spell.Formula1}}<tr><td>Formula1</td><td>{{.Spell.Formula1}}</td>{{end}}
                    {{if .Spell.Formula2}}<tr><td>Formula2</td><td>{{.Spell.Formula2}}</td>{{end}}
                    {{if .Spell.Formula3}}<tr><td>Formula3</td><td>{{.Spell.Formula3}}</td>{{end}}
                    {{if .Spell.Formula4}}<tr><td>Formula4</td><td>{{.Spell.Formula4}}</td>{{end}}
                    {{if .Spell.Formula5}}<tr><td>Formula5</td><td>{{.Spell.Formula5}}</td>{{end}}
                    {{if .Spell.Formula6}}<tr><td>Formula6</td><td>{{.Spell.Formula6}}</td>{{end}}
                    {{if .Spell.Formula7}}<tr><td>Formula7</td><td>{{.Spell.Formula7}}</td>{{end}}
                    {{if .Spell.Formula8}}<tr><td>Formula8</td><td>{{.Spell.Formula8}}</td>{{end}}
                    {{if .Spell.Formula9}}<tr><td>Formula9</td><td>{{.Spell.Formula9}}</td>{{end}}
                    {{if .Spell.Formula10}}<tr><td>Formula10</td><td>{{.Spell.Formula10}}</td>{{end}}
                    {{if .Spell.Formula11}}<tr><td>Formula11</td><td>{{.Spell.Formula11}}</td>{{end}}
                    {{if .Spell.Formula12}}<tr><td>Formula12</td><td>{{.Spell.Formula12}}</td>{{end}}
                    {{if .Spell.Lighttype}}<tr><td>Lighttype</td><td>{{.Spell.Lighttype}}</td>{{end}}
                    {{if gt .Spell.Goodeffect 0}}<tr><td>Goodeffect</td><td>{{.Spell.Goodeffect}}</td>{{end}}
                    {{if .Spell.Activated}}<tr><td>Activated</td><td>{{.Spell.Activated}}</td>{{end}}
                    {{if .Spell.Resisttype}}<tr><td>Resisttype</td><td>{{.Spell.Resisttype}}</td>{{end}}
                    {{if .Spell.Effectid1}}<tr><td>Effectid1</td><td>{{.Spell.Effectid1}}</td>{{end}}
                    {{if .Spell.Effectid2}}<tr><td>Effectid2</td><td>{{.Spell.Effectid2}}</td>{{end}}
                    {{if .Spell.Effectid3}}<tr><td>Effectid3</td><td>{{.Spell.Effectid3}}</td>{{end}}
                    {{if .Spell.Effectid4}}<tr><td>Effectid4</td><td>{{.Spell.Effectid4}}</td>{{end}}
                    {{if .Spell.Effectid5}}<tr><td>Effectid5</td><td>{{.Spell.Effectid5}}</td>{{end}}
                    {{if .Spell.Effectid6}}<tr><td>Effectid6</td><td>{{.Spell.Effectid6}}</td>{{end}}
                    {{if .Spell.Effectid7}}<tr><td>Effectid7</td><td>{{.Spell.Effectid7}}</td>{{end}}
                    {{if .Spell.Effectid8}}<tr><td>Effectid8</td><td>{{.Spell.Effectid8}}</td>{{end}}
                    {{if .Spell.Effectid9}}<tr><td>Effectid9</td><td>{{.Spell.Effectid9}}</td>{{end}}
                    {{if .Spell.Effectid10}}<tr><td>Effectid10</td><td>{{.Spell.Effectid10}}</td>{{end}}
                    {{if .Spell.Effectid11}}<tr><td>Effectid11</td><td>{{.Spell.Effectid11}}</td>{{end}}
                    {{if .Spell.Effectid12}}<tr><td>Effectid12</td><td>{{.Spell.Effectid12}}</td>{{end}}
                    {{if .Spell.Targettype}}<tr><td>Targettype</td><td>{{.Spell.Targettype}}</td>{{end}}
                    {{if .Spell.Basediff}}<tr><td>Basediff</td><td>{{.Spell.Basediff}}</td>{{end}}
                    {{if .Spell.Skill}}<tr><td>Skill</td><td>{{.Spell.Skill}}</td>{{end}}
                    {{if .Spell.Zonetype}}<tr><td>Zonetype</td><td>{{.Spell.Zonetype}}</td>{{end}}
                    {{if .Spell.Environmenttype}}<tr><td>Environmenttype</td><td>{{.Spell.Environmenttype}}</td>{{end}}
                    {{if .Spell.Timeofday}}<tr><td>Timeofday</td><td>{{.Spell.Timeofday}}</td>{{end}}
                    <tr><td>Classes</td><td>{{.Spell.ClassesList}}</td>
                    {{if .Spell.Targetanim}}<tr><td>Targetanim</td><td>{{.Spell.Targetanim}}</td>{{end}}
                    {{if .Spell.Traveltype}}<tr><td>Traveltype</td><td>{{.Spell.Traveltype}}</td>{{end}}
                    {{if .Spell.Spellaffectindex}}<tr><td>Spellaffectindex</td><td>{{.Spell.Spellaffectindex}}</td>{{end}}
                    {{if .Spell.DisallowSit}}<tr><td>DisallowSit</td><td>{{.Spell.DisallowSit}}</td>{{end}}
                    {{if .Spell.Deities0}}<tr><td>Deities0</td><td>{{.Spell.Deities0}}</td>{{end}}
                    {{if .Spell.Deities1}}<tr><td>Deities1</td><td>{{.Spell.Deities1}}</td>{{end}}
                    {{if .Spell.Deities2}}<tr><td>Deities2</td><td>{{.Spell.Deities2}}</td>{{end}}
                    {{if .Spell.Deities3}}<tr><td>Deities3</td><td>{{.Spell.Deities3}}</td>{{end}}
                    {{if .Spell.Deities4}}<tr><td>Deities4</td><td>{{.Spell.Deities4}}</td>{{end}}
                    {{if .Spell.Deities5}}<tr><td>Deities5</td><td>{{.Spell.Deities5}}</td>{{end}}
                    {{if .Spell.Deities6}}<tr><td>Deities6</td><td>{{.Spell.Deities6}}</td>{{end}}
                    {{if .Spell.Deities7}}<tr><td>Deities7</td><td>{{.Spell.Deities7}}</td>{{end}}
                    {{if .Spell.Deities8}}<tr><td>Deities8</td><td>{{.Spell.Deities8}}</td>{{end}}
                    {{if .Spell.Deities9}}<tr><td>Deities9</td><td>{{.Spell.Deities9}}</td>{{end}}
                    {{if .Spell.Deities10}}<tr><td>Deities10</td><td>{{.Spell.Deities10}}</td>{{end}}
                    {{if .Spell.Deities11}}<tr><td>Deities11</td><td>{{.Spell.Deities11}}</td>{{end}}
                    {{if .Spell.Deities12}}<tr><td>Deities12</td><td>{{.Spell.Deities12}}</td>{{end}}
                    {{if .Spell.Deities13}}<tr><td>Deities13</td><td>{{.Spell.Deities13}}</td>{{end}}
                    {{if .Spell.Deities14}}<tr><td>Deities14</td><td>{{.Spell.Deities14}}</td>{{end}}
                    {{if .Spell.Deities15}}<tr><td>Deities15</td><td>{{.Spell.Deities15}}</td>{{end}}
                    {{if .Spell.Deities16}}<tr><td>Deities16</td><td>{{.Spell.Deities16}}</td>{{end}}
                    {{if .Spell.Field142}}<tr><td>Field142</td><td>{{.Spell.Field142}}</td>{{end}}
                    {{if .Spell.Field143}}<tr><td>Field143</td><td>{{.Spell.Field143}}</td>{{end}}
                    {{if .Spell.NewIcon}}<tr><td>NewIcon</td><td>{{.Spell.NewIcon}}</td>{{end}}
                    {{if .Spell.Spellanim}}<tr><td>Spellanim</td><td>{{.Spell.Spellanim}}</td>{{end}}
                    {{if .Spell.Uninterruptable}}<tr><td>Uninterruptable</td><td>{{.Spell.Uninterruptable}}</td>{{end}}
                    {{if .Spell.Resistdiff}}<tr><td>Resistdiff</td><td>{{.Spell.Resistdiff}}</td>{{end}}
                    {{if .Spell.DotStackingExempt}}<tr><td>DotStackingExempt</td><td>{{.Spell.DotStackingExempt}}</td>{{end}}
                    {{if .Spell.Deleteable}}<tr><td>Deleteable</td><td>{{.Spell.Deleteable}}</td>{{end}}
                    {{if .Spell.Recourselink}}<tr><td>Recourselink</td><td>{{.Spell.Recourselink}}</td>{{end}}
                    {{if .Spell.NoPartialResist}}<tr><td>NoPartialResist</td><td>{{.Spell.NoPartialResist}}</td>{{end}}
                    {{if .Spell.Field152}}<tr><td>Field152</td><td>{{.Spell.Field152}}</td>{{end}}
                    {{if .Spell.Field153}}<tr><td>Field153</td><td>{{.Spell.Field153}}</td>{{end}}
                    {{if .Spell.ShortBuffBox}}<tr><td>ShortBuffBox</td><td>{{.Spell.ShortBuffBox}}</td>{{end}}
                    {{if .Spell.Descnum}}<tr><td>Descnum</td><td>{{.Spell.Descnum}}</td>{{end}}
                    {{if .Spell.Typedescnum}}<tr><td>Typedescnum</td><td>{{.Spell.Typedescnum}}</td>{{end}}
                    {{if .Spell.Effectdescnum}}<tr><td>Effectdescnum</td><td>{{.Spell.Effectdescnum}}</td>{{end}}
                    {{if .Spell.Effectdescnum2}}<tr><td>Effectdescnum2</td><td>{{.Spell.Effectdescnum2}}</td>{{end}}
                    {{if .Spell.NpcNoLos}}<tr><td>NpcNoLos</td><td>{{.Spell.NpcNoLos}}</td>{{end}}
                    {{if .Spell.Field160}}<tr><td>Field160</td><td>{{.Spell.Field160}}</td>{{end}}
                    {{if .Spell.Reflectable}}<tr><td>Reflectable</td><td>{{.Spell.Reflectable}}</td>{{end}}
                    {{if .Spell.Bonushate}}<tr><td>Bonushate</td><td>{{.Spell.Bonushate}}</td>{{end}}
                    {{if .Spell.Field163}}<tr><td>Field163</td><td>{{.Spell.Field163}}</td>{{end}}
                    {{if .Spell.Field164}}<tr><td>Field164</td><td>{{.Spell.Field164}}</td>{{end}}
                    {{if .Spell.LdonTrap}}<tr><td>LdonTrap</td><td>{{.Spell.LdonTrap}}</td>{{end}}
                    {{if .Spell.Endurcost}}<tr><td>Endurcost</td><td>{{.Spell.Endurcost}}</td>{{end}}
                    {{if .Spell.Endurtimerindex}}<tr><td>Endurtimerindex</td><td>{{.Spell.Endurtimerindex}}</td>{{end}}
                    {{if .Spell.Isdiscipline}}<tr><td>Isdiscipline</td><td>{{.Spell.Isdiscipline}}</td>{{end}}
                    {{if .Spell.Field169}}<tr><td>Field169</td><td>{{.Spell.Field169}}</td>{{end}}
                    {{if .Spell.Field170}}<tr><td>Field170</td><td>{{.Spell.Field170}}</td>{{end}}
                    {{if .Spell.Field171}}<tr><td>Field171</td><td>{{.Spell.Field171}}</td>{{end}}
                    {{if .Spell.Field172}}<tr><td>Field172</td><td>{{.Spell.Field172}}</td>{{end}}
                    {{if .Spell.Hateadded}}<tr><td>Hateadded</td><td>{{.Spell.Hateadded}}</td>{{end}}
                    {{if .Spell.Endurupkeep}}<tr><td>Endurupkeep</td><td>{{.Spell.Endurupkeep}}</td>{{end}}
                    {{if .Spell.Numhitstype}}<tr><td>Numhitstype</td><td>{{.Spell.Numhitstype}}</td>{{end}}
                    {{if .Spell.Numhits}}<tr><td>Numhits</td><td>{{.Spell.Numhits}}</td>{{end}}
                    {{if .Spell.Pvpresistbase}}<tr><td>Pvpresistbase</td><td>{{.Spell.Pvpresistbase}}</td>{{end}}
                    {{if .Spell.Pvpresistcalc}}<tr><td>Pvpresistcalc</td><td>{{.Spell.Pvpresistcalc}}</td>{{end}}
                    {{if .Spell.Pvpresistcap}}<tr><td>Pvpresistcap</td><td>{{.Spell.Pvpresistcap}}</td>{{end}}
                    {{if .Spell.SpellCategory}}<tr><td>SpellCategory</td><td>{{.Spell.SpellCategory}}</td>{{end}}
                    {{if .Spell.Field181}}<tr><td>Field181</td><td>{{.Spell.Field181}}</td>{{end}}
                    {{if .Spell.Field182}}<tr><td>Field182</td><td>{{.Spell.Field182}}</td>{{end}}
                    {{if .Spell.PcnpcOnlyFlag}}<tr><td>PcnpcOnlyFlag</td><td>{{.Spell.PcnpcOnlyFlag}}</td>{{end}}
                    {{if .Spell.CastNotStanding}}<tr><td>CastNotStanding</td><td>{{.Spell.CastNotStanding}}</td>{{end}}
                    {{if .Spell.CanMgb}}<tr><td>CanMgb</td><td>{{.Spell.CanMgb}}</td>{{end}}
                    {{if .Spell.Nodispell}}<tr><td>Nodispell</td><td>{{.Spell.Nodispell}}</td>{{end}}
                    {{if .Spell.NpcCategory}}<tr><td>NpcCategory</td><td>{{.Spell.NpcCategory}}</td>{{end}}
                    {{if .Spell.NpcUsefulness}}<tr><td>NpcUsefulness</td><td>{{.Spell.NpcUsefulness}}</td>{{end}}
                    {{if .Spell.Minresist}}<tr><td>Minresist</td><td>{{.Spell.Minresist}}</td>{{end}}
                    {{if .Spell.Maxresist}}<tr><td>Maxresist</td><td>{{.Spell.Maxresist}}</td>{{end}}
                    {{if .Spell.ViralTargets}}<tr><td>ViralTargets</td><td>{{.Spell.ViralTargets}}</td>{{end}}
                    {{if .Spell.ViralTimer}}<tr><td>ViralTimer</td><td>{{.Spell.ViralTimer}}</td>{{end}}
                    {{if .Spell.Nimbuseffect}}<tr><td>Nimbuseffect</td><td>{{.Spell.Nimbuseffect}}</td>{{end}}
                    {{if .Spell.Conestartangle}}<tr><td>Conestartangle</td><td>{{.Spell.Conestartangle}}</td>{{end}}
                    {{if .Spell.Conestopangle}}<tr><td>Conestopangle</td><td>{{.Spell.Conestopangle}}</td>{{end}}
                    {{if .Spell.Sneaking}}<tr><td>Sneaking</td><td>{{.Spell.Sneaking}}</td>{{end}}
                    {{if .Spell.NotExtendable}}<tr><td>NotExtendable</td><td>{{.Spell.NotExtendable}}</td>{{end}}
                    {{if .Spell.Field198}}<tr><td>Field198</td><td>{{.Spell.Field198}}</td>{{end}}
                    {{if .Spell.Field199}}<tr><td>Field199</td><td>{{.Spell.Field199}}</td>{{end}}
                    {{if .Spell.Suspendable}}<tr><td>Suspendable</td><td>{{.Spell.Suspendable}}</td>{{end}}
                    {{if .Spell.ViralRange}}<tr><td>ViralRange</td><td>{{.Spell.ViralRange}}</td>{{end}}
                    {{if .Spell.Songcap}}<tr><td>Songcap</td><td>{{.Spell.Songcap}}</td>{{end}}
                    {{if .Spell.Field203}}<tr><td>Field203</td><td>{{.Spell.Field203}}</td>{{end}}
                    {{if .Spell.Field204}}<tr><td>Field204</td><td>{{.Spell.Field204}}</td>{{end}}
                    {{if .Spell.NoBlock}}<tr><td>NoBlock</td><td>{{.Spell.NoBlock}}</td>{{end}}
                    {{if .Spell.Field206}}<tr><td>Field206</td><td>{{.Spell.Field206}}</td>{{end}}
                    {{if .Spell.Spellgroup}}<tr><td>Spellgroup</td><td>{{.Spell.Spellgroup}}</td>{{end}}
                    {{if .Spell.Rank}}<tr><td>Rank</td><td>{{.Spell.Rank}}</td>{{end}}
                    {{if .Spell.Field209}}<tr><td>Field209</td><td>{{.Spell.Field209}}</td>{{end}}
                    {{if .Spell.Field210}}<tr><td>Field210</td><td>{{.Spell.Field210}}</td>{{end}}
                    {{if .Spell.Castrestriction}}<tr><td>Castrestriction</td><td>{{.Spell.Castrestriction}}</td>{{end}}
                    {{if .Spell.Allowrest}}<tr><td>Allowrest</td><td>{{.Spell.Allowrest}}</td>{{end}}
                    {{if .Spell.Incombat}}<tr><td>Incombat</td><td>{{.Spell.Incombat}}</td>{{end}}
                    {{if .Spell.Outofcombat}}<tr><td>Outofcombat</td><td>{{.Spell.Outofcombat}}</td>{{end}}
                    {{if .Spell.Field215}}<tr><td>Field215</td><td>{{.Spell.Field215}}</td>{{end}}
                    {{if .Spell.Field216}}<tr><td>Field216</td><td>{{.Spell.Field216}}</td>{{end}}
                    {{if .Spell.Field217}}<tr><td>Field217</td><td>{{.Spell.Field217}}</td>{{end}}
                    {{if .Spell.Aemaxtargets}}<tr><td>Aemaxtargets</td><td>{{.Spell.Aemaxtargets}}</td>{{end}}
                    {{if .Spell.Maxtargets}}<tr><td>Maxtargets</td><td>{{.Spell.Maxtargets}}</td>{{end}}
                    {{if .Spell.Field220}}<tr><td>Field220</td><td>{{.Spell.Field220}}</td>{{end}}
                    {{if .Spell.Field221}}<tr><td>Field221</td><td>{{.Spell.Field221}}</td>{{end}}
                    {{if .Spell.Field222}}<tr><td>Field222</td><td>{{.Spell.Field222}}</td>{{end}}
                    {{if .Spell.Field223}}<tr><td>Field223</td><td>{{.Spell.Field223}}</td>{{end}}
                    {{if .Spell.Persistdeath}}<tr><td>Persistdeath</td><td>{{.Spell.Persistdeath}}</td>{{end}}
                    {{if .Spell.Field225}}<tr><td>Field225</td><td>{{.Spell.Field225}}</td>{{end}}
                    {{if .Spell.Field226}}<tr><td>Field226</td><td>{{.Spell.Field226}}</td>{{end}}
                    {{if .Spell.MinDist}}<tr><td>MinDist</td><td>{{.Spell.MinDist}}</td>{{end}}
                    {{if .Spell.MinDistMod}}<tr><td>MinDistMod</td><td>{{.Spell.MinDistMod}}</td>{{end}}
                    {{if .Spell.MaxDist}}<tr><td>MaxDist</td><td>{{.Spell.MaxDist}}</td>{{end}}
                    {{if .Spell.MaxDistMod}}<tr><td>MaxDistMod</td><td>{{.Spell.MaxDistMod}}</td>{{end}}
                    {{if .Spell.MinRange}}<tr><td>MinRange</td><td>{{.Spell.MinRange}}</td>{{end}}
                    {{if .Spell.Field232}}<tr><td>Field232</td><td>{{.Spell.Field232}}</td>{{end}}
                    {{if .Spell.Field233}}<tr><td>Field233</td><td>{{.Spell.Field233}}</td>{{end}}
                    {{if .Spell.Field234}}<tr><td>Field234</td><td>{{.Spell.Field234}}</td>{{end}}
                    {{if .Spell.Field235}}<tr><td>Field235</td><td>{{.Spell.Field235}}</td>{{end}}
                    {{if .Spell.Field236}}<tr><td>Field236</td><td>{{.Spell.Field236}}</td>{{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
    {{if .Items}}
    <div class="col-lg-4">
        <div class="hpanel forum-box">
            <div class="panel-heading">
               Items with {{.Spell.Name.String}}
            </div>
            <div class="panel-body">
                
                <div class="table-responsive">
                <table cellpadding="1" cellspacing="1" class="table table-striped">

                        <thead>
                        <tr>                            
                            <th width="20px"><i class="xa xa-sword"></i></th>
                            <th>Name</th>
                            <th>Slot</th>
                            <th><i title="Category" class="xa xa-vest"></i></th>
                        </tr>
                        </thead>
                        <tbody>
                        {{range $key, $value := .Items}}
                        <tr>
                            <td><span class="item icon-{{$value.Icon}}-sm"></span></td>
                                <td><a item={{$value.ID}} href="/item/{{$value.ID}}">{{$value.Name}}</a></td>
                                <td><span title="{{$value.SlotList}}">{{$value.SlotsFirstName}}</span></td>
                                <td><i title="{{$value.ItemtypeName}}" class="xa {{$value.ItemtypeIcon}}"></i></td></td>
                        </tr>
                        {{end}}                
                        </tbody>
                    
                </table>
                </div>
            </div>
            <div class="panel-footer">
                {{len .Items}} items have the {{.Spell.Name.String}} effect
            </div>
        </div>

    </div>
    {{end}} 
    {{if .Npcs}}
    <div class="col-lg-4">
        <div class="hpanel">
            <div class="panel-heading hbuilt">
                <div class="panel-tools">
                    <a class="showhide"><i class="fa fa-chevron-up"></i></a>
                    <a class="closebox"><i class="fa fa-times"></i></a>
                </div>
                This spell is used by these creatures:
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
                {{len .Npcs}} creatures
            </div>
        </div>
    </div>
    {{end}}
</div>
</div>
</div>