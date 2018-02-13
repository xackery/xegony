package api

const tooltipTemplate = `
<div class="eqitem">
	
    <span class="slot"><span class="item icon-{{.Item.Icon}}"></span></span><br>{{.Item.Name}}<br>
    {{if .Item.Magic}}Magic{{end}}{{if .Item.NoTransfer}}No Trade{{end}}<br>
    Class: {{.Item.ClassList}}<br>
    Race: {{.Item.RaceList}}<br>
    {{.Item.SlotList}}<br>
    <br>
    <table cellpadding="0" cellspacing="0" border="0">
    <tbody>
        <tr><td colspan="2">Size:<span style="float:right; padding-left:8px; white-space:nowrap">{{.Item.Size.Name}}</span></td><td style="padding-right:8px;"></td>{{if .Item.Hitpoint}}<td style="padding-right:4px" nowrap="">HP:</td><td align="right">{{.Item.Hp}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr><td style="padding-right:4px" nowrap="">Weight:</td><td align="right">{{.Item.Weight}}</td><td class="hvalue" style="padding-right:8px;"></td>{{if .Item.Mana}}<td style="padding-right:4px" nowrap="">Mana:</td><td align="right">{{.Item.Mana}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr>{{if .Item.RecommendedLevel}}<td style="padding-right:4px" nowrap="">Rec Level:</td><td align="right">{{.Item.RecommendedLevel}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td>{{if .Item.Endurance}}<td style="padding-right:4px" nowrap="">Endur:</td><td align="right">{{.Item.Endurance}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr>{{if .Item.RequiredLevel}}<td style="padding-right:4px" nowrap="">Req Level:</td><td align="right">{{.Item.RequiredLevel}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr><td style="height:4px;font-size:1px">&nbsp;</td></tr>
        <tr>{{if .Item.Strength}}<td style="padding-right:4px" nowrap="">Strength:</td><td align="right">{{.Item.Strength}}</td>{{end}}{{if .Item.HeroicStrength}}<td class="hvalue" style="padding-right:8px;">{{.Item.HeroicStrength}}</td>{{end}}{{if .Item.MagicResistance}}<td style="padding-right:4px" nowrap="">Magic:</td><td align="right">{{.Item.MagicResistance}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td>{{if .Item.Attack}}<td style="padding-right:4px" nowrap="">Attack:</td><td align="right">{{.Item.Attack}}</td><td class="hvalue" style="padding-right:8px;"></td>{{end}}</tr>
        <tr>{{if .Item.Stamina}}<td style="padding-right:4px" nowrap="">Stamina:</td><td align="right">{{.Item.Stamina}}</td>{{end}}{{if .Item.HeroicStamina}}<td class="hvalue" style="padding-right:8px;">+{{.Item.HeroicStamina}}</td>{{end}}{{if .Item.FireResistance}}<td style="padding-right:4px" nowrap="">Fire:</td><td align="right">{{.Item.FireResistance}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td>{{if .Item.Regen}}<td style="padding-right:4px" nowrap="">HP Regen:</td><td align="right">{{.Item.Regen}}</td><td class="hvalue" style="padding-right:8px;"></td>{{end}}</tr>
        <tr>{{if .Item.Intelligence}}<td style="padding-right:4px" nowrap="">Intelligence:</td><td align="right">{{.Item.Intelligence}}</td>{{end}}{{if .Item.HeroicIntelligence}}<td class="hvalue" style="padding-right:8px;">+{{.Item.HeroicIntelligence}}</td>{{end}}{{if .Item.ColdResistance}}<td style="padding-right:4px" nowrap="">Cold:</td><td align="right">{{.Item.ColdResistance}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td>{{if .Item.ManaRegen}}<td style="padding-right:4px" nowrap="">Mana Regen:</td><td align="right">{{.Item.ManaRegen}}</td><td class="hvalue" style="padding-right:8px;"></td>{{end}}</tr>
        <tr>{{if .Item.Wisdom}}<td style="padding-right:4px" nowrap="">Wisdom:</td><td align="right">{{.Item.Wisdom}}</td>{{end}}{{if .Item.HeroicWisdom}}<td class="hvalue" style="padding-right:8px;">+{{.Item.HeroicWisdom}}</td>{{end}}{{if .Item.DiseaseResistance}}<td style="padding-right:4px" nowrap="">Disease:</td><td align="right">{{.Item.DiseaseResistance}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr>{{if .Item.Agility}}<td style="padding-right:4px" nowrap="">Agility:</td><td align="right">{{.Item.Agility}}</td>{{end}}{{if .Item.HeroicAgility}}<td class="hvalue" style="padding-right:8px;">+{{.Item.HeroicAgility}}</td>{{end}}{{if .Item.PoisonResistance}}<td style="padding-right:4px" nowrap="">Poison:</td><td align="right">{{.Item.PoisonResistance}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr>{{if .Item.Dexterity}}<td style="padding-right:4px" nowrap="">Dexterity:</td><td align="right">{{.Item.Dexterity}}</td>{{end}}{{if .Item.HeroicDexterity}}<td class="hvalue" style="padding-right:8px;">+{{.Item.HeroicDexterity}}</td>{{end}}</tr>
        <tr>{{if .Item.Charisma}}<td style="padding-right:4px" nowrap="">Charisma:</td><td align="right">{{.Item.Charisma}}</td>{{end}}{{if .Item.HeroicCharisma}}<td class="hvalue" style="padding-right:8px;">+{{.Item.HeroicCharisma}}</td>{{end}}</tr>
    </tbody>
    </table>
    {{/*<div class="augments">
        Slot 1, type 3 (General: Spell Effect): empty<br>
        Slot 2, type 5 (Weapon: Elem Damage): empty<br>
        Slot 3, type 7 (General: Group): empty<br>
        Slot 4, type 9 (General: Dragons Points): empty<br>
    </div>*/}}
    {{/*<div class="effects">
        Effect: <a rel="eq:spell:9616" href="/spell/9616" target="_blank">Sharpshooting VII</a> (Worn)<br>
        Focus Effect: <a rel="eq:spell:42971" href="/spell/42971" target="_blank">Detrimental Duration 26 L110</a><br>
    </div>*/}}

</div>`
