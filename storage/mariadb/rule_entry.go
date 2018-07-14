package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	ruleEntryTable  = "rule_values"
	ruleEntryFields = "ruleset_id, rule_value, notes"
	ruleEntryBinds  = ":ruleset_id, :rule_value, :notes"
)

//GetRuleEntry will grab data from storage
func (s *Storage) GetRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	query := fmt.Sprintf("SELECT rule_name, %s FROM %s WHERE ruleset_id = ? AND rule_name = ?", ruleEntryFields, ruleEntryTable)
	err = s.db.Get(ruleEntry, query, rule.ID, ruleEntry.Name)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateRuleEntry will grab data from storage
func (s *Storage) CreateRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(rule_name, %s) VALUES (:rule_name, %s)", ruleEntryTable, ruleEntryFields, ruleEntryBinds)
	_, err = s.db.NamedExec(query, ruleEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListRuleEntry will grab data from storage
func (s *Storage) ListRuleEntry(page *model.Page, rule *model.Rule) (ruleEntrys []*model.RuleEntry, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "rule_name"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT rule_name, %s FROM %s WHERE ruleset_id = ? ORDER BY %s LIMIT %d OFFSET %d", ruleEntryFields, ruleEntryTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query, rule.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		ruleEntry := model.RuleEntry{}
		if err = rows.StructScan(&ruleEntry); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		ruleEntrys = append(ruleEntrys, &ruleEntry)
	}
	return
}

//ListRuleEntryTotalCount will grab data from storage
func (s *Storage) ListRuleEntryTotalCount(rule *model.Rule) (count int64, err error) {
	query := fmt.Sprintf("SELECT count(ruleset_id) FROM %s WHERE ruleset_id = ?", ruleEntryTable)
	err = s.db.Get(&count, query, rule.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListRuleEntryBySearch will grab data from storage
func (s *Storage) ListRuleEntryBySearch(page *model.Page, rule *model.Rule, ruleEntry *model.RuleEntry) (ruleEntrys []*model.RuleEntry, err error) {

	field := ""

	if len(ruleEntry.Name) > 0 {
		field += `rule_name LIKE :rule_name OR`
		ruleEntry.Name = fmt.Sprintf("%%%s%%", ruleEntry.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]
	ruleEntry.RuleID = rule.ID

	query := fmt.Sprintf("SELECT rule_name, %s FROM %s WHERE %s AND ruleset_id = :ruleset_id LIMIT %d OFFSET %d", ruleEntryFields, ruleEntryTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, ruleEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		ruleEntry := model.RuleEntry{}
		if err = rows.StructScan(&ruleEntry); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		ruleEntrys = append(ruleEntrys, &ruleEntry)
	}
	return
}

//ListRuleEntryBySearchTotalCount will grab data from storage
func (s *Storage) ListRuleEntryBySearchTotalCount(rule *model.Rule, ruleEntry *model.RuleEntry) (count int64, err error) {
	field := ""
	if len(ruleEntry.Name) > 0 {
		field += `rule_name LIKE :rule_name OR`
		ruleEntry.Name = fmt.Sprintf("%%%s%%", ruleEntry.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	ruleEntry.RuleID = rule.ID
	query := fmt.Sprintf("SELECT count(ruleset_id) FROM %s WHERE %s AND ruleset_id = :ruleset_id", ruleEntryTable, field)

	rows, err := s.db.NamedQuery(query, ruleEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
	}
	return
}

//EditRuleEntry will grab data from storage
func (s *Storage) EditRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {

	prevRuleEntry := &model.RuleEntry{
		Name: ruleEntry.Name,
	}
	err = s.GetRuleEntry(rule, prevRuleEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous ruleEntry")
		return
	}

	field := ""
	if len(ruleEntry.Value) > 0 && prevRuleEntry.Value != ruleEntry.Value {
		field += "value = :value, "
	}
	if len(ruleEntry.Description.String) > 0 && prevRuleEntry.Description.String != ruleEntry.Description.String {
		field += "notes = :notes, "
	}
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE ruleset_id = :ruleset_id AND rule_name = :rule_name", ruleEntryTable, field)
	result, err := s.db.NamedExec(query, ruleEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//DeleteRuleEntry will grab data from storage
func (s *Storage) DeleteRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE ruleset_id = ? AND rule_name = ?", ruleEntryTable)
	result, err := s.db.Exec(query, rule.ID, ruleEntry.Name)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

func (s *Storage) insertTestRuleEntry() (err error) {

	_, err = s.db.Exec(`INSERT INTO rule_values (ruleset_id, rule_name, rule_value, notes)
VALUES
	(0, 'Character:MarqueeHPUpdates', 'false', 'Will show Health % in center of screen < 100%'),
	(0, 'Character:RestrictSpellScribing', 'false', 'Restricts spell scribing to allowable races/classes of spell scroll, if true'),
	(1, 'Skills:MaxTrainSpecializations', '50', NULL),
	(1, 'Skills:SwimmingStartValue', '100', NULL),
	(1, 'World:IsGMPetitionWindowEnabled', 'false', NULL),
	(1, 'Spells:ReflectType', '1', NULL),
	(1, 'Spells:AI_EngagedNoSpellMaxRecast', '1000', NULL),
	(1, 'Spells:AI_EngagedBeneficialSelfChance', '100', NULL),
	(1, 'Character:UseNewStatsWindow', 'true', NULL),
	(1, 'Character:ItemDSMitigationCap', '50', NULL),
	(1, 'Character:ItemSpellDmgCap', '250', NULL),
	(1, 'Character:TradeskillUpBlacksmithing', '2', NULL),
	(1, 'Character:ItemHealAmtCap', '250', NULL),
	(1, 'Character:TradeskillUpBrewing', '3', NULL),
	(1, 'Character:ItemClairvoyanceCap', '250', NULL),
	(1, 'Character:TradeskillUpBaking', '2', NULL),
	(1, 'Spells:EnableBlockedBuffs', 'true', NULL),
	(1, 'Combat:NPCBashKickStunChance', '15', NULL),
	(1, 'Combat:MeleeCritDifficulty', '8900', NULL),
	(1, 'Character:ItemEnduranceRegenCap', '15', NULL),
	(1, 'Character:RespawnFromHoverTimer', '300', NULL),
	(1, 'Character:RespawnFromHover', 'false', NULL),
	(1, 'Character:MarqueeHPUpdates', 'false', NULL),
	(1, 'Combat:AdjustSpecialProcPerMinute', 'false', NULL),
	(1, 'Combat:UseOldDamageIntervalRules', 'false', NULL),
	(1, 'Spells:MaxTotalSlotsNPC', '66', NULL),
	(1, 'Spells:MaxSongSlotsNPC', '10', NULL),
	(1, 'Spells:MaxDiscSlotsNPC', '1', NULL),
	(1, 'Spells:MaxBuffSlotsNPC', '55', NULL),
	(1, 'Character:SoDClientUseSoDHPManaEnd', 'true', NULL),
	(1, 'Character:UseOldClassExpPenalties', 'false', NULL),
	(1, 'Skills:MaxTradeskillSearchSkillDiff', '50', NULL),
	(1, 'Skills:SenseHeadingStartValue', '200', NULL),
	(1, 'GM:MinStatusToSummonItem', '250', NULL),
	(1, 'Combat:HitCapPre10', '20', NULL),
	(1, 'Combat:MinHastedDelay', '400', NULL),
	(1, 'Combat:ArcheryBonusChance', '50', NULL),
	(1, 'Skills:UseLimitTradeskillSearchSkillDiff', 'true', NULL),
	(1, 'Skills:TrainSenseHeading', 'false', NULL),
	(1, 'Spells:SwarmPetTargetLock', 'false', NULL),
	(1, 'Character:DeathExpLossMaxLevel', '255', NULL),
	(1, 'Character:GreenModifier', '20', NULL),
	(1, 'World:GuildBankZoneID', '345', NULL),
	(1, 'Combat:DefaultRampageTargets', '1', NULL),
	(1, 'NPC:ReturnNonQuestNoDropItems', 'true', NULL),
	(1, 'Aggro:AllowTickPulling', 'false', NULL),
	(1, 'Spells:EnableSpellGlobals', 'false', NULL),
	(1, 'Spells:ReflectMessagesClose', 'true', NULL),
	(1, 'Spells:UseCHAScribeHack', 'false', NULL),
	(1, 'Zone:MarkMQWarpLT', 'false', NULL),
	(1, 'Zone:HotZoneBonus', '0.5000000000000', NULL),
	(1, 'Chat:GlobalChatLevelLimit', '1', NULL),
	(1, 'Chat:KarmaGlobalChatLimit', '72', NULL),
	(1, 'Combat:FleeSnareHPRatio', '21', 'HP at which snare will halt movement of a fleeing NPC.'),
	(1, 'World:MaxClientsSetByStatus', 'false', NULL),
	(1, 'Map:FixZWhenMoving', 'true', NULL),
	(1, 'Map:MobZVisualDebug', 'false', NULL),
	(1, 'AA:ExpPerPoint', '23976503', NULL),
	(1, 'Character:FactionLossMultiplier', '1.0000000000000', NULL),
	(1, 'Character:MaxFearDurationForPlayerCharacter', '1', NULL),
	(1, 'Character:MaxCharmDurationForPlayerCharacter', '15', NULL),
	(1, 'Character:BaseHPRegenBonusRaces', '4352', NULL),
	(1, 'AA:MaxEffectSlots', '7', 'the highest slot # used in the aa_effects table. have to use MAX_AA_EFFECT_SLOTS for now'),
	(1, 'Combat:RoundKickBonus', '5', NULL),
	(1, 'Combat:BackstabBonus', '0', NULL),
	(1, 'Combat:TigerClawBonus', '10', NULL),
	(1, 'Combat:ClothACSoftcap', '75', NULL),
	(1, 'Combat:EagleStrikeBonus', '15', NULL),
	(1, 'Combat:LeatherACSoftcap', '100', NULL),
	(1, 'Combat:DragonPunchBonus', '20', NULL),
	(1, 'Combat:MonkACSoftcap', '120', NULL),
	(1, 'Combat:FlyingKickBonus', '25', NULL),
	(1, 'Combat:ChainACSoftcap', '200', NULL),
	(1, 'Combat:MonkDamageTableBonus', '10', NULL),
	(1, 'Combat:PlateACSoftcap', '300', NULL),
	(1, 'Combat:MaxFlurryHits', '2', NULL),
	(1, 'Combat:MaxRampageTargets', '3', NULL),
	(1, 'Spells:SacrificeItemID', '9963', NULL),
	(1, 'Spells:MaxTotalSlotsPET', '30', NULL),
	(1, 'Spells:SacrificeMaxLevel', '69', NULL),
	(1, 'Spells:SacrificeMinLevel', '46', NULL),
	(1, 'Map:UseClosestZ', 'true', ''),
	(1, 'Map:FindBestZHeightAdjust', '1', NULL),
	(1, 'Pathing:ZDiffThreshold', '10.0000000000000', NULL),
	(1, 'Pathing:RouteUpdateFrequencyShort', '1000', NULL),
	(1, 'Pathing:RouteUpdateFrequencyNodeCount', '5', NULL),
	(1, 'Pathing:MinDistanceForLOSCheckShort', '40000.0000000000000', NULL),
	(1, 'Pathing:MinNodesTraversedForLOSCheck', '3', NULL),
	(1, 'Pathing:RouteUpdateFrequencyLong', '5000', NULL),
	(1, 'Pathing:LOSCheckFrequency', '1000', NULL),
	(1, 'Pathing:MaxNodesLeftForLOSCheck', '4', ''),
	(1, 'Pathing:MinDistanceForLOSCheckLong', '1000000.0000000000000', NULL),
	(1, 'Pathing:Guard', 'true', NULL),
	(1, 'Combat:ArcheryBaseDamageBonus', '1.0000000000000', NULL),
	(1, 'Pathing:Fear', 'true', NULL),
	(1, 'Pathing:CullNodesFromStart', '1', NULL),
	(1, 'Pathing:CullNodesFromEnd', '1', NULL),
	(1, 'Pathing:CandidateNodeRangeZ', '50.0000000000000', NULL),
	(1, 'NPC:SmartLastFightingDelayMoving', 'true', NULL),
	(1, 'Aggro:UseLevelAggro', 'true', NULL),
	(1, 'Chat:SuppressCommandErrors', 'false', NULL),
	(1, 'Combat:ArcheryBonusRequiresStationary', 'true', NULL),
	(1, 'Combat:OldACSoftcapRules', 'false', NULL),
	(1, 'NPC:LastFightingDelayMovingMax', '30000', NULL),
	(1, 'NPC:LastFightingDelayMovingMin', '5000', NULL),
	(1, 'NPC:NPCToNPCAggroTimerMax', '6000', NULL),
	(1, 'Zone:PEQZoneDebuff2', '2209', NULL),
	(1, 'Zone:ReservedInstances', '30', NULL),
	(1, 'Zone:PEQZoneDebuff1', '4454', NULL),
	(1, 'Zone:EbonCrystalItemID', '40902', NULL),
	(1, 'Zone:PEQZoneReuseTime', '300', NULL),
	(1, 'Zone:RadiantCrystalItemID', '40903', NULL),
	(1, 'Pathing:AggroReturnToGrid', 'true', NULL),
	(1, 'Pathing:CandidateNodeRangeXY', '400.0000000000000', NULL),
	(1, 'Pathing:Aggro', 'true', NULL),
	(1, 'Adventure:DistanceForRescueComplete', '2500.0000000000000', NULL),
	(1, 'Adventure:NumberKillsForBossSpawn', '40', NULL),
	(1, 'Adventure:DistanceForRescueAccept', '10000.0000000000000', NULL),
	(1, 'Adventure:MaxNumberForRaid', '36', NULL),
	(1, 'Adventure:MaxLevelRange', '9', NULL),
	(1, 'Adventure:MaxNumberForGroup', '6', NULL),
	(1, 'Adventure:MinNumberForRaid', '18', NULL),
	(1, 'Adventure:MinNumberForGroup', '2', NULL),
	(1, 'Adventure:ItemIDToEnablePorts', '41000', NULL),
	(1, 'Character:KillsPerGroupLeadershipAA', '250', NULL),
	(1, 'Character:KillsPerRaidLeadershipAA', '250', NULL),
	(1, 'Character:YellowModifier', '125', NULL),
	(1, 'Character:RedModifier', '150', NULL),
	(1, 'Character:WhiteModifier', '100', NULL),
	(1, 'Character:LightBlueModifier', '40', NULL),
	(1, 'Character:BlueModifier', '90', NULL),
	(1, 'Combat:NPCBonusHitChance', '26.0000000000000', NULL),
	(1, 'Character:UseXPConScaling', 'true', NULL),
	(1, 'Zone:UsePEQZoneDebuffs', 'false', NULL),
	(1, 'Zone:EnableLoggedOffReplenishments', 'true', NULL),
	(1, 'Character:RestRegenTimeToActivate', '30', NULL),
	(1, 'Character:RestRegenRaidTimeToActivate', '300', NULL),
	(1, 'Character:RestRegenPercent', '2', NULL),
	(1, 'Character:RaidExpMultiplier', '0.2000000029802', NULL),
	(1, 'World:MinGMAntiHackStatus', '11', NULL),
	(1, 'World:SoFStartZoneID', '-1', NULL),
	(1, 'Chat:KarmaUpdateIntervalMS', '1200000', NULL),
	(1, 'Character:GroupExpMultiplier', '0.6499999761581', NULL),
	(1, 'Chat:MaxMessagesBeforeKick', '20', NULL),
	(1, 'Chat:IntervalDurationMS', '60000', NULL),
	(1, 'Chat:MinimumMessagesPerInterval', '4', NULL),
	(1, 'Chat:MaximumMessagesPerInterval', '12', NULL),
	(1, 'Chat:MinStatusToBypassAntiSpam', '80', NULL),
	(1, 'Chat:EnableMailKeyIPVerification', 'true', NULL),
	(1, 'Chat:EnableAntiSpam', 'true', NULL),
	(1, 'Combat:WeaponSkillFalloff', '0.3300000131130', NULL),
	(1, 'Combat:ArcheryHitPenalty', '0.4499999880791', NULL),
	(1, 'Combat:MinChancetoHit', '5.0000000000000', NULL),
	(1, 'Combat:HitBonusPerLevel', '1.2000000476837', NULL),
	(1, 'Combat:HitFalloffMajor', '50.0000000000000', NULL),
	(1, 'Combat:HitFalloffModerate', '7.0000000000000', NULL),
	(1, 'Combat:HitFalloffMinor', '5.0000000000000', NULL),
	(1, 'Merchant:ChaPenaltyMod', '1.5199999809265', NULL),
	(1, 'Merchant:ChaBonusMod', '3.4500000476837', NULL),
	(1, 'Adventure:LDoNBaseTrapDifficulty', '15.0000000000000', NULL),
	(1, 'Merchant:PricePenaltyPct', '4', NULL),
	(1, 'Pets:PetPowerLevelCap', '10.0000000000000', NULL),
	(1, 'Merchant:PriceBonusPct', '4', NULL),
	(1, 'Merchant:BuyCostMod', '0.6499999761581', NULL),
	(1, 'Merchant:SellCostMod', '1.0499999523163', NULL),
	(1, 'Merchant:UsePriceMod', 'true', NULL),
	(1, 'World:GMAccountIPList', 'false', NULL),
	(1, 'Character:AAExpMultiplier', '0.6499999761581', NULL),
	(1, 'EventLog:RecordBuyFromMerchant', 'true', NULL),
	(1, 'EventLog:RecordSellToMerchant', 'true', NULL),
	(1, 'QueryServ:PlayerLogDropItem', 'false', NULL),
	(1, 'Chat:EnableVoiceMacros', 'false', NULL),
	(1, 'Spells:TranslocateTimeLimit', '0', NULL),
	(1, 'Channels:DeleteTimer', '1440', NULL),
	(1, 'Adventure:LDoNTrapDistanceUse', '625', NULL),
	(1, 'Channels:RequiredStatusListAll', '251', NULL),
	(1, 'Channels:RequiredStatusAdmin', '251', NULL),
	(1, 'Mail:ExpireUnread', '31536000', NULL),
	(1, 'Mail:ExpireRead', '31536000', NULL),
	(1, 'Mail:ExpireTrash', '0', NULL),
	(1, 'Adventure:LDoNAdventureExpireTime', '1800', NULL),
	(1, 'Mail:EnableMailSystem', 'false', NULL),
	(1, 'QueryServ:PlayerLogZone', 'false', NULL),
	(1, 'Bazaar:MaxBarterSearchResults', '200', NULL),
	(1, 'Bazaar:EnableWarpToTrader', 'true', NULL),
	(1, 'World:TutorialZoneID', '189', NULL),
	(1, 'Bazaar:MaxSearchResults', '200', NULL),
	(1, 'Bazaar:AuditTrail', 'true', NULL),
	(1, 'World:ExemptAccountLimitStatus', '200', NULL),
	(1, 'World:TitaniumStartZoneID', '-1', NULL),
	(1, 'World:AccountSessionLimit', '1', NULL),
	(1, 'Spells:AI_PursueNoSpellMaxRecast', '2000', NULL),
	(1, 'Character:BindAnywhere', 'false', NULL),
	(1, 'Character:UseDeathExpLossMult', 'false', NULL),
	(1, 'Character:DeathExpLossMultiplier', '2', NULL),
	(1, 'Combat:BaseHitChance', '69.0000000000000', NULL),
	(1, 'Combat:AgiHitFactor', '0.0099999997765', NULL),
	(1, 'Spells:ResistPerLevelDiff', '85', NULL),
	(1, 'Character:SharedBankPlat', 'true', NULL),
	(1, 'Spells:WizCritChance', '5', NULL),
	(1, 'Spells:WizCritRatio', '0', NULL),
	(1, 'Spells:WizCritLevel', '30', NULL),
	(1, 'Spells:BaseCritRatio', '100', NULL),
	(1, 'Spells:AI_EngagedBeneficialOtherChance', '25', NULL),
	(1, 'Spells:BaseCritChance', '0', NULL),
	(1, 'Spells:AI_EngagedDetrimentalChance', '20', NULL),
	(1, 'Combat:BaseProcChance', '0.0350000001490', NULL),
	(1, 'Combat:ProcDexDivideBy', '11000.0000000000000', NULL),
	(1, 'Combat:AvgProcsPerMinute', '2.0000000000000', NULL),
	(1, 'Combat:ProcPerMinDexContrib', '0.0750000029802', NULL),
	(1, 'Chat:ServerWideAuction', 'true', NULL),
	(1, 'Combat:AdjustProcPerMinute', 'true', NULL),
	(1, 'Chat:ServerWideOOC', 'true', NULL),
	(1, 'Character:MaxExpLevel', '60', NULL),
	(1, 'Character:ItemATKCap', '250', NULL),
	(1, 'Zone:AutoShutdownDelay', '5000', NULL),
	(1, 'World:ClearTempMerchantlist', 'false', NULL),
	(1, 'World:AddMaxClientsPerIP', '5', NULL),
	(1, 'Spells:AI_IdleNoSpellMaxRecast', '60000', NULL),
	(1, 'Spells:AI_IdleBeneficialChance', '100', NULL),
	(1, 'World:EnableReturnHomeButton', 'true', NULL),
	(1, 'World:MaxLevelForTutorial', '1', NULL),
	(1, 'World:MinOfflineTimeToReturnHome', '21600', NULL),
	(1, 'World:AddMaxClientsStatus', '25', NULL),
	(1, 'Spells:AI_PursueDetrimentalChance', '90', NULL),
	(1, 'Spells:AI_IdleNoSpellMinRecast', '6000', NULL),
	(1, 'TaskSystem:KeepOneRecordPerCompletedTask', 'true', NULL),
	(1, 'TaskSystem:EnableTaskProximity', 'true', NULL),
	(1, 'World:EnableTutorialButton', 'false', NULL),
	(1, 'TaskSystem:RecordCompletedOptionalActivities', 'true', NULL),
	(1, 'TaskSystem:PeriodicCheckTimer', '5', NULL),
	(1, 'TaskSystem:RecordCompletedTasks', 'true', NULL),
	(1, 'TaskSystem:EnableTaskSystem', 'true', NULL),
	(1, 'Character:SkillUpModifier', '200', NULL),
	(1, 'NPC:EnableNPCQuestJournal', 'false', NULL),
	(1, 'NPC:CorpseUnlockTimer', '150000', NULL),
	(1, 'NPC:EmptyNPCCorpseDecayTimeMS', '6000', NULL),
	(1, 'Spells:PartialHitChanceFear', '0.2500000000000', NULL),
	(1, 'World:UseBannedIPsTable', 'true', NULL),
	(1, 'Pathing:Find', 'true', NULL),
	(1, 'Spells:PreNerfBardAEDoT', 'false', NULL),
	(1, 'Spells:Jun182014HundredHandsRevamp', 'false', NULL),
	(1, 'Zone:MQWarpExemptStatus', '150', NULL),
	(1, 'Character:ItemStrikethroughCap', '35', NULL),
	(1, 'Zone:MQWarpDetectionDistanceFactor', '9.0000000000000', NULL),
	(1, 'Character:ItemDoTShieldingCap', '35', NULL),
	(1, 'Character:ItemStunResistCap', '35', NULL),
	(1, 'Character:ItemSpellShieldingCap', '35', NULL),
	(1, 'Character:ItemCombatEffectsCap', '100', NULL),
	(1, 'Character:ItemShieldingCap', '35', NULL),
	(1, 'Character:ItemAccuracyCap', '150', NULL),
	(1, 'Character:ItemAvoidanceCap', '100', NULL),
	(1, 'Zone:MQZoneExemptStatus', '150', NULL),
	(1, 'Character:ItemDamageShieldCap', '30', NULL),
	(1, 'Zone:MQGhostExemptStatus', '150', NULL),
	(1, 'Zone:MQGateExemptStatus', '150', NULL),
	(1, 'Zone:EnableMQWarpDetector', 'true', NULL),
	(1, 'Combat:FleeIfNotAlone', 'false', NULL),
	(1, 'Zone:EnableMQZoneDetector', 'true', NULL),
	(1, 'Zone:EnableMQGateDetector', 'true', NULL),
	(1, 'World:MaxClientsPerIP', '5', NULL),
	(1, 'Combat:ThrowingCritDifficulty', '1100', NULL),
	(1, 'Combat:MinRangedAttackDist', '25', NULL),
	(1, 'Aggro:PetSpellAggroMod', '10', NULL),
	(1, 'Character:DeathItemLossLevel', '40', NULL),
	(1, 'Zone:EnableMQGhostDetector', 'true', NULL),
	(1, 'NPC:BuffFriends', 'true', NULL),
	(1, 'Aggro:SpellAggroMod', '100', NULL),
	(1, 'Aggro:SongAggroMod', '33', NULL),
	(1, 'Combat:AssistNoTargetSelf', 'false', NULL),
	(1, 'Combat:RampageHitsTarget', 'false', NULL),
	(1, 'Combat:ProcTargetOnly', 'true', NULL),
	(1, 'Aggro:CriticallyWoundedAggroMod', '100', NULL),
	(1, 'Aggro:CurrentTargetAggroMod', '0', NULL),
	(1, 'Aggro:MaxScalingProcAggro', '400', NULL),
	(1, 'Aggro:IntAggroThreshold', '75', NULL),
	(1, 'Aggro:SittingAggroMod', '35', NULL),
	(1, 'Aggro:MeleeRangeAggroMod', '20', NULL),
	(1, 'Aggro:SmartAggroList', 'true', NULL),
	(1, 'Watermap:FishingLineLength', '100.0000000000000', NULL),
	(1, 'Watermap:FishingLineStepSize', '1.0000000000000', NULL),
	(1, 'Watermap:FishingRodLength', '30.0000000000000', NULL),
	(1, 'Watermap:CheckForWaterWhenFishing', 'true', NULL),
	(1, 'Spells:UseLiveSpellProjectileGFX', 'false', NULL),
	(1, 'Watermap:CheckForWaterOnSendTo', 'false', NULL),
	(1, 'Spells:FocusCombatProcs', 'false', NULL),
	(1, 'Watermap:CheckForWaterAtWaypoints', 'true', NULL),
	(1, 'Watermap:CheckForWaterWhenMoving', 'true', NULL),
	(1, 'NPC:OOCRegen', '1', NULL),
	(1, 'Watermap:CheckWaypointsInWaterWhenLoading', 'true', NULL),
	(1, 'NPC:SayPauseTimeInSec', '3', NULL),
	(1, 'Combat:UseIntervalAC', 'true', NULL),
	(1, 'Combat:PetAttackMagicLevel', '30', NULL),
	(1, 'Character:ItemManaRegenCap', '25', NULL),
	(1, 'Character:ItemHealthRegenCap', '25', NULL),
	(1, 'Character:HealOnLevel', 'true', NULL),
	(1, 'Character:FeignKillsPet', 'false', NULL),
	(1, 'Map:FixPathingZMaxDeltaWaypoint', '20.0000000000000', NULL),
	(1, 'Map:FixPathingZMaxDeltaLoading', '20.0000000000000', NULL),
	(1, 'Map:FixPathingZMaxDeltaMoving', '20.0000000000000', NULL),
	(1, 'Zone:ClientLinkdeadMS', '180000', NULL),
	(1, 'Map:FixPathingZMaxDeltaSendTo', '20.0000000000000', NULL),
	(1, 'Map:FixPathingZAtWaypoints', 'true', NULL),
	(1, 'Map:FixPathingZWhenMoving', 'true', NULL),
	(1, 'Map:FixPathingZOnSendTo', 'false', NULL),
	(1, 'Zone:GraveyardTimeMS', '1200000', NULL),
	(1, 'Map:FixPathingZWhenLoading', 'true', NULL),
	(1, 'Zone:EnableShadowrest', 'false', NULL),
	(1, 'NPC:UseItemBonusesForNonPets', 'true', NULL),
	(1, 'Character:ExpMultiplier', '0.6499999761581', NULL),
	(1, 'NPC:MinorNPCCorpseDecayTimeMS', '600000', NULL),
	(1, 'NPC:MajorNPCCorpseDecayTimeMS', '1800000', NULL),
	(1, 'Zone:NPCPositonUpdateTicCount', '32', NULL),
	(1, 'Zone:MinOfflineTimeToReplenishments', '21600', NULL),
	(1, 'Combat:ClientBaseCritChance', '0', ''),
	(1, 'Spells:PartialHitChance', '0.6999999880791', NULL),
	(1, 'Combat:MaxChancetoHit', '95.0000000000000', NULL),
	(1, 'Spells:ResistMod', '0.4000000059605', NULL),
	(1, 'Spells:AutoResistDiff', '15', NULL),
	(1, 'Spells:ResistChance', '2.0000000000000', NULL),
	(1, 'Character:ConsumptionMultiplier', '100', NULL),
	(1, 'Combat:BerserkBaseCritChance', '6', ''),
	(1, 'Combat:NPCBashKickLevel', '6', NULL),
	(1, 'Combat:WarBerBaseCritChance', '3', 'The base crit chance for warriors and berserkers:only applies to clients'),
	(1, 'Combat:MeleeBaseCritChance', '0', 'The base crit chance for non warriors:NOTE: This will apply to NPCs as well'),
	(1, 'Combat:EnableFearPathing', 'true', NULL),
	(1, 'Combat:FleeHPRatio', '21', NULL),
	(1, 'World:ExemptMaxClientsStatus', '-1', NULL),
	(1, 'Combat:PetBaseCritChance', '0', NULL),
	(1, 'Combat:ArcheryCritDifficulty', '3400', NULL),
	(1, 'World:ZoneAutobootTimeoutMS', '120000', NULL),
	(1, 'World:ClientKeepaliveTimeoutMS', '95000', NULL),
	(1, 'Pets:AttackCommandRange', '200.0000000000000', NULL),
	(1, 'Skills:MaxTrainTradeskills', '21', NULL),
	(1, 'Combat:BerserkerFrenzyStart', '35', NULL),
	(1, 'Guild:MaxMembers', '2048', NULL),
	(1, 'Character:EnduranceRegenMultiplier', '100', NULL),
	(1, 'Character:ManaRegenMultiplier', '100', NULL),
	(1, 'Character:HPRegenMultiplier', '100', NULL),
	(1, 'Character:AutosaveIntervalS', '300', NULL),
	(1, 'Character:LeaveNakedCorpses', 'true', NULL),
	(1, 'Character:DeathExpLossLevel', '10', NULL),
	(1, 'Character:CorpseDecayTimeMS', '604800000', NULL),
	(1, 'GM:MinStatusToZoneAnywhere', '250', NULL),
	(1, 'Combat:HitCapPre20', '40', NULL),
	(1, 'Character:MaxLevel', '60', NULL),
	(1, 'Character:LeaveCorpses', 'true', NULL),
	(1, 'Character:HasteCap', '100', NULL),
	(1, 'Character:TradeskillUpAlchemy', '2', NULL),
	(1, 'Character:MaxDraggedCorpses', '2', NULL),
	(1, 'Character:ShowExpValues', '0', NULL),
	(1, 'Character:DragCorpseDistance', '400.0000000000000', NULL),
	(1, 'Character:EnvironmentDamageMulipliter', '1.0000000000000', NULL),
	(1, 'Character:RestRegenEndurance', 'true', NULL),
	(1, 'Character:UseOldBindWound', 'false', NULL),
	(1, 'Spells:VirusSpreadDistance', '30', NULL),
	(1, 'Spells:BaseImmunityLevel', '55', NULL),
	(1, 'Combat:ArcheryNPCMultiplier', '1.0000000000000', NULL),
	(1, 'Combat:NPCACFactor', '2.2500000000000', NULL),
	(1, 'Character:ItemCastsUseFocus', 'false', NULL),
	(1, 'Character:SpamHPUpdates', 'false', NULL),
	(1, 'World:ExpansionSettings', '16383', NULL),
	(1, 'World:PVPSettings', '0', NULL),
	(1, 'Spells:AI_PursueNoSpellMinRecast', '500', NULL),
	(1, 'World:FVNoDropFlag', '0', NULL),
	(1, 'World:TellQueueSize', '20', NULL),
	(1, 'Zone:GlobalLootMultiplier', '1', NULL),
	(1, 'Pathing:MinNodesLeftForLOSCheck', '4', NULL),
	(1, 'Character:MinStatusForNoDropExemptions', '80', NULL),
	(1, 'Character:OrnamentationAugmentType', '20', NULL),
	(1, 'Character:SkillCapMaxLevel', '-1', NULL),
	(1, 'Character:BaseInstrumentSoftCap', '36', NULL),
	(1, 'Character:BaseRunSpeedCap', '158', NULL),
	(1, 'Console:SessionTimeOut', '600000', NULL),
	(1, 'Character:FactionGainMultiplier', '1.0000000000000', NULL),
	(1, 'Guild:PlayerCreationAllowed', 'true', NULL),
	(1, 'Spells:NPC_UseFocusFromSpells', 'true', NULL),
	(1, 'Guild:PlayerCreationLimit', '1', NULL),
	(1, 'Combat:LevelToStopDamageCaps', '0', NULL),
	(1, 'Guild:PlayerCreationRequiredLevel', '0', NULL),
	(1, 'Combat:NPCAssistCap', '5', NULL),
	(1, 'Guild:PlayerCreationRequiredStatus', '0', NULL),
	(1, 'Combat:NPCAssistCapTimer', '6000', NULL),
	(1, 'Guild:PlayerCreationRequiredTime', '0', NULL),
	(1, 'Combat:BerserkerFrenzyEnd', '45', NULL),
	(1, 'Spells:LiveLikeFocusEffects', 'true', NULL),
	(1, 'Spells:NPCIgnoreBaseImmunity', 'true', NULL),
	(1, 'NPC:StartEnrageValue', '9', NULL),
	(1, 'NPC:NPCToNPCAggroTimerMin', '500', NULL),
	(1, 'NPC:LiveLikeEnrage', 'false', NULL),
	(1, 'NPC:EnableMeritBasedFaction', 'false', NULL),
	(1, 'Character:ItemExtraDmgCap', '150', NULL),
	(1, 'Character:CheckCursorEmptyWhenLooting', 'true', NULL),
	(1, 'Character:UseSpellFileSongCap', 'true', NULL),
	(1, 'Character:MaintainIntoxicationAcrossZones', 'true', NULL),
	(1, 'Character:KeepLevelOverMax', 'false', NULL),
	(1, 'Bots:BotAAExpansion', '8', 'The expansion through which bots will obtain AAs'),
	(1, 'Character:UseRaceClassExpBonuses', 'false', NULL),
	(1, 'Character:UseOldRaceExpPenalties', 'false', NULL),
	(1, 'Character:EnableDiscoveredItems', 'true', NULL),
	(1, 'Character:EnableAggroMeter', 'true', NULL),
	(1, 'Combat:WarriorThreatBonus', '0', ''),
	(1, 'Character:CorpseResTimeMS', '10800000', NULL),
	(1, 'Character:StatCap', '0', NULL),
	(1, 'Character:FoodLossPerUpdate', '35', NULL),
	(1, 'Character:SumCorpseDecayTimeMS', '604800000', ''),
	(1, 'World:DeleteStaleCorpeBackups', 'true', NULL),
	(1, 'Zone:EnableZoneControllerGlobals', 'false', NULL),
	(1, 'Zone:UsePlayerCorpseBackups', 'true', NULL),
	(1, 'Character:EnableXTargetting', 'true', NULL),
	(1, 'Spells:AvgSpellProcsPerMinute', '6.0000000000000', NULL),
	(1, 'Combat:FleeMultiplier', '2.0000000000000', NULL),
	(1, 'Spells:ResistFalloff', '67', NULL),
	(1, 'Spells:AI_EngagedNoSpellMinRecast', '500', NULL),
	(1, 'Spells:CharismaEffectiveness', '10', NULL),
	(1, 'Spells:CharismaEffectivenessCap', '255', NULL),
	(1, 'Spells:AI_SpellCastFinishedFailRecast', '800', NULL),
	(1, 'Spells:CharmBreakCheckChance', '25', NULL),
	(1, 'Spells:FRProjectileItem_NPC', '80684', NULL),
	(1, 'Combat:AvgDefProcsPerMinute', '2.0000000000000', NULL),
	(1, 'Combat:DefProcPerMinAgiContrib', '0.0750000029802', NULL),
	(1, 'Adventure:LDoNCriticalFailTrapThreshold', '10.0000000000000', NULL),
	(1, 'Combat:SpecialAttackACBonus', '15', NULL),
	(1, 'Combat:ClientStunLevel', '55', NULL),
	(1, 'Combat:QuiverHasteCap', '1000', NULL),
	(1, 'Combat:FrenzyBonus', '0', NULL),
	(1, 'Combat:TauntSkillFalloff', '0.3300000131130', NULL),
	(1, 'Aggro:TunnelVisionAggroMod', '0.7500000000000', NULL),
	(1, 'Spells:MaxCastTimeReduction', '50', NULL),
	(1, 'Spells:RootBreakCheckChance', '70', NULL),
	(1, 'Spells:RootBreakFromSpells', '55', NULL),
	(1, 'Spells:DeathSaveCharismaMod', '3', NULL),
	(1, 'Spells:AdditiveBonusWornType', '0', NULL),
	(1, 'Spells:DivineInterventionHeal', '8000', NULL),
	(1, 'Combat:NPCFlurryChance', '20', NULL),
	(1, 'Combat:MonkACBonusWeight', '15', NULL),
	(1, 'AA:Stacking', 'true', NULL),
	(1, 'QueryServ:PlayerLogMerchantTransactions', 'false', NULL),
	(1, 'Spells:AdditiveBonusValues', 'false', 'Allow certain bonuses to be calculated by adding together the value from each item, instead of taking the highest value. (ie Add together all Cleave Effects)'),
	(1, 'Character:PerCharacterQglobalMaxLevel', 'false', NULL),
	(1, 'Character:GrantHoTTOnCreate', 'false', NULL),
	(1, 'QueryServ:PlayerLogNPCKills', 'false', NULL),
	(1, 'QueryServ:PlayerLogTrades', 'false', NULL),
	(1, 'QueryServ:PlayerLogChat', 'true', NULL),
	(1, 'QueryServ:PlayerLogMoneyTrades', 'false', ''),
	(1, 'QueryServ:PlayerLogPCCoordinates', 'false', NULL),
	(1, 'Chat:FlowCommandstoPerl_EVENT_SAY', 'true', ''),
	(1, 'World:IPLimitDisconnectAll', 'false', NULL),
	(1, 'World:MaxClientsSimplifiedLogic', 'false', NULL),
	(1, 'World:StartZoneSameAsBindOnCreation', 'true', NULL),
	(1, 'Pets:UnTargetableSwarmPet', 'false', NULL),
	(1, 'QueryServ:MerchantLogTransactions', 'true', ''),
	(1, 'QueryServ:PlayerLogDeletes', 'false', NULL),
	(1, 'QueryServ:PlayerLogHandins', 'true', NULL),
	(1, 'QueryServ:PlayerLogMoves', 'false', NULL),
	(1, 'Mercs:AllowMercs', 'false', NULL),
	(1, 'Mercs:MercsUsePathing', 'true', NULL),
	(1, 'Mercs:SuspendIntervalMS', '10000', NULL),
	(1, 'Mercs:UpkeepIntervalMS', '180000', NULL),
	(1, 'Range:ClientNPCScan', '300', NULL),
	(1, 'Mercs:SuspendIntervalS', '10', NULL),
	(1, 'Range:CriticalDamage', '80', NULL),
	(1, 'Mercs:UpkeepIntervalS', '180', NULL),
	(1, 'Mercs:ScaleRate', '100', NULL),
	(1, 'Aggro:ClientAggroCheckInterval', '6', NULL),
	(1, 'Mercs:AggroRadius', '100', NULL),
	(1, 'Range:SpellMessages', '75', NULL),
	(1, 'Range:SongMessages', '75', NULL),
	(1, 'Range:MobPositionUpdates', '600', NULL),
	(1, 'Mercs:AggroRadiusPuller', '25', NULL),
	(1, 'Range:Emote', '135', NULL),
	(1, 'Range:BeginCast', '200', NULL),
	(1, 'Range:Anims', '135', NULL),
	(1, 'Range:SpellParticles', '135', NULL),
	(1, 'Range:DamageMessages', '50', NULL),
	(1, 'Zone:LevelBasedEXPMods', 'true', NULL),
	(1, 'Mercs:ResurrectRadius', '50', NULL),
	(1, 'Range:Say', '135', NULL),
	(1, 'Mercs:ChargeMercPurchaseCost', 'false', NULL),
	(1, 'Mercs:ChargeMercUpkeepCost', 'false', NULL),
	(1, 'Spells:BuffLevelRestrictions', 'true', NULL),
	(1, 'Zone:WeatherTimer', '600', NULL),
	(1, 'Spells:CharismaCharmDuration', 'false', NULL),
	(1, 'Spells:CharismaResistCap', '255', 'Maximium amount of CHA that will effect charm resist rate.'),
	(1, 'Inventory:EnforceAugmentRestriction', 'false', NULL),
	(1, 'Inventory:EnforceAugmentUsability', 'false', NULL),
	(1, 'Inventory:EnforceAugmentWear', 'false', NULL),
	(1, 'Spells:FearBreakCheckChance', '70', NULL),
	(1, 'Spells:FRProjectileItem_SOF', '80684', NULL),
	(1, 'Spells:SuccorFailChance', '1', NULL),
	(1, 'Spells:FRProjectileItem_Titanium', '1113', NULL),
	(1, 'Combat:AvgSpecialProcsPerMinute', '2.0000000000000', NULL),
	(1, 'Client:UseLiveFactionMessage', 'true', NULL),
	(1, 'Client:UseLiveBlockedMessage', 'false', NULL),
	(1, 'Character:ActiveInvSnapshots', 'false', NULL),
	(1, 'Character:AllowMQTarget', 'false', NULL),
	(1, 'Character:InvSnapshotMinIntervalM', '180', NULL),
	(1, 'Character:InvSnapshotMinRetryM', '30', NULL),
	(1, 'Character:AvoidanceCap', '750', NULL),
	(1, 'Character:InvSnapshotHistoryD', '30', NULL),
	(1, 'Spells:SHDProcIDOffByOne', 'false', NULL),
	(1, 'Character:RestrictSpellScribing', 'false', NULL),
	(1, 'Character:EnableAvoidanceCap', 'false', NULL),
	(1, 'Combat:MeleePush', 'false', NULL),
	(1, 'Combat:MeleePushChance', '0', NULL),
	(1, 'Character:UseStackablePickPocketing', 'true', NULL),
	(1, 'World:UseClientBasedExpansionSettings', 'false', NULL),
	(1, 'Character:UnmemSpellsOnDeath', 'false', NULL),
	(1, 'Zone:UseZoneController', 'false', NULL),
	(1, 'Character:IksarCommonTongue', '100', NULL),
	(1, 'Character:TrollCommonTongue', '100', NULL),
	(1, 'Character:OgreCommonTongue', '100', NULL),
	(1, 'NPC:UseClassAsLastName', 'true', NULL),
	(1, 'NPC:NewLevelScaling', 'true', NULL),
	(1, 'World:EnableIPExemptions', 'false', NULL),
	(1, 'Character:TradeskillUpFletching', '2', NULL),
	(1, 'Character:TradeskillUpJewelcrafting', '2', NULL),
	(1, 'Character:TradeskillUpMakePoison', '2', NULL),
	(1, 'Character:TradeskillUpPottery', '4', NULL),
	(1, 'Character:TradeskillUpResearch', '1', NULL),
	(1, 'Character:TradeskillUpTinkering', '2', NULL),
	(1, 'Combat:AAMitigationACFactor', '3.0000000000000', NULL),
	(1, 'Combat:WarriorACSoftcapReturn', '0.4499999880791', NULL),
	(1, 'Combat:KnightACSoftcapReturn', '0.3300000131130', NULL),
	(1, 'Combat:LowPlateChainACSoftcapReturn', '0.2300000041723', NULL),
	(1, 'Combat:LowChainLeatherACSoftcapReturn', '0.1700000017881', NULL),
	(1, 'Combat:CasterACSoftcapReturn', '0.0599999986589', NULL),
	(1, 'Combat:MiscACSoftcapReturn', '0.3000000119209', NULL),
	(1, 'Combat:WarACSoftcapReturn', '0.3447999954224', NULL),
	(1, 'Combat:ClrRngMnkBrdACSoftcapReturn', '0.3030000030994', NULL),
	(1, 'Combat:PalShdACSoftcapReturn', '0.3226000070572', NULL),
	(1, 'Combat:DruNecWizEncMagACSoftcapReturn', '0.2000000029802', NULL),
	(1, 'Combat:RogShmBstBerACSoftcapReturn', '0.2500000000000', NULL),
	(1, 'Combat:SoftcapFactor', '1.8799999952316', NULL),
	(1, 'Combat:ACthac0Factor', '0.5500000119209', NULL),
	(1, 'Combat:ACthac20Factor', '0.5500000119209', NULL),
	(1, 'Character:UseOldConSystem', 'false', NULL),
	(1, 'Mercs:AllowMercSuspendInCombat', 'true', NULL),
	(1, 'Spells:NPC_UseFocusFromItems', 'false', NULL),
	(1, 'Spells:UseAdditiveFocusFromWornSlot', 'false', NULL),
	(1, 'Spells:AlwaysSendTargetsBuffs', 'false', NULL),
	(1, 'Spells:FlatItemExtraSpellAmt', 'false', NULL),
	(1, 'Spells:IgnoreSpellDmgLvlRestriction', 'false', NULL),
	(1, 'Spells:AllowItemTGB', 'false', NULL),
	(1, 'Spells:NPCInnateProcOverride', 'true', NULL),
	(1, 'Spells:OldRainTargets', 'false', NULL),
	(1, 'Combat:NPCCanCrit', 'false', NULL),
	(1, 'Combat:TauntOverLevel', 'true', NULL),
	(1, 'Combat:EXPFromDmgShield', 'false', NULL),
	(1, 'Combat:UseArcheryBonusRoll', 'false', NULL),
	(1, 'Combat:OneProcPerWeapon', 'true', NULL),
	(1, 'Combat:ProjectileDmgOnImpact', 'true', NULL),
	(1, 'Combat:UseLiveCombatRounds', 'true', NULL),
	(1, 'Combat:UseRevampHandToHand', 'false', NULL),
	(1, 'Combat:ClassicMasterWu', 'false', NULL),
	(1, 'Merchant:EnableAltCurrencySell', 'true', NULL),
	(1, 'QueryServ:PlayerLogDeaths', 'false', NULL),
	(1, 'QueryServ:PlayerLogConnectDisconnect', 'false', NULL),
	(1, 'QueryServ:PlayerLogLevels', 'false', NULL),
	(1, 'QueryServ:PlayerLogAARate', 'false', NULL),
	(1, 'QueryServ:PlayerLogQGlobalUpdate', 'false', NULL),
	(1, 'QueryServ:PlayerLogTaskUpdates', 'false', NULL),
	(1, 'QueryServ:PlayerLogKeyringAddition', 'false', NULL),
	(1, 'QueryServ:PlayerLogAAPurchases', 'false', NULL),
	(1, 'QueryServ:PlayerLogTradeSkillEvents', 'false', NULL),
	(1, 'QueryServ:PlayerLogIssuedCommandes', 'false', NULL),
	(1, 'QueryServ:PlayerLogMoneyTransactions', 'false', NULL),
	(1, 'QueryServ:PlayerLogAlternateCurrencyTransactions', 'false', NULL),
	(1, 'Inventory:DeleteTransformationMold', 'true', NULL),
	(1, 'Inventory:AllowAnyWeaponTransformation', 'false', NULL),
	(1, 'Inventory:TransformSummonedBags', 'false', NULL);`)
	if err != nil {
		err = errors.Wrap(err, "failed to insert npc data")
		return
	}
	return
}

//createTableRuleEntry will grab data from storage
func (s *Storage) createTableRuleEntry() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE rule_values (
  ruleset_id tinyint(3) unsigned NOT NULL DEFAULT '0',
  rule_name varchar(64) NOT NULL DEFAULT '',
  rule_value varchar(30) NOT NULL DEFAULT '',
  notes text,
  PRIMARY KEY (ruleset_id,rule_name),
  KEY ruleset_id (ruleset_id)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
