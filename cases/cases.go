package cases

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/oauth"
	"github.com/xackery/xegony/storage"
)

var (
	readers      = make(map[string]storage.Reader)
	readLock     = sync.RWMutex{}
	writers      = make(map[string]storage.Writer)
	writeLock    = sync.RWMutex{}
	initializers = make(map[string]storage.Initializer)
	initLock     = sync.RWMutex{}
	oauths       = make(map[string]oauth.Wrapper)
	oauthLock    = sync.RWMutex{}
)

//Initialize a specific scope
func Initialize(scope string, sr storage.Reader, sw storage.Writer, si storage.Initializer) (err error) {
	SetReader(scope, sr)
	SetWriter(scope, sw)
	SetInitializer(scope, si)
	return
}

//FlushStorage resets all storage devices
func FlushStorage() (err error) {
	err = flushAllReader()
	if err != nil {
		return
	}
	err = flushAllWriter()
	if err != nil {
		return
	}
	err = flushAllInitializer()
	if err != nil {
		return
	}
	err = flushAllOauth()
	if err != nil {
		return
	}
	return
}

//InitializeAllMemoryStorage will load all memory based points
func InitializeAllMemoryStorage() (err error) {

	err = LoadClassFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zone to memory")
		return
	}

	err = LoadDeityFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zone to memory")
		return
	}

	err = LoadOauthTypeFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zone to memory")
		return
	}

	err = LoadRaceFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zone to memory")
		return
	}

	err = LoadRuleFromDBToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load rule to memory")
		return
	}

	err = LoadRuleEntryFromDBToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load ruleEntry to memory")
		return
	}

	err = LoadSpellAnimationFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellAnimation to memory")
		return
	}

	err = LoadSpellAnimationTypeFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellAnimationType to memory")
		return
	}

	err = LoadSpellDurationFormulaFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellDurationFormula to memory")
		return
	}

	err = LoadSpellEffectFormulaFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellEffectFormula to memory")
		return
	}

	err = LoadSpellEffectTypeFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellEffectType to memory")
		return
	}

	err = LoadSpellTargetTypeFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellTargetType to memory")
		return
	}

	err = LoadSpellTravelTypeFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellTravelType to memory")
		return
	}

	err = LoadVariableFromDBToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load variable to memory")
		return
	}

	err = LoadZoneFromDBToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zone to memory")
		return
	}

	err = LoadZoneExpansionFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zoneExpansion to memory")
		return
	}
	return
}

// InitializeAllDatabaseStorage readers, writers, and initializers
func InitializeAllDatabaseStorage(sr storage.Reader, sw storage.Writer, si storage.Initializer) (err error) {

	scopes := []string{
		"account",
		"character",
		"forum",
		"item",
		"npc",
		"rule",
		"ruleEntry",
		"spawn",
		"spawnEntry",
		"spawnNpc",
		"spell",
		"user",
		"userAccount",
		"userLink",
		"variable",
		"zone",
	}

	for _, scope := range scopes {
		err = Initialize(scope, sr, sw, si)
		if err != nil {
			return
		}
	}
	return
}

func flushAllReader() (err error) {
	readLock.Lock()
	readers = map[string]storage.Reader{}
	readLock.Unlock()
	return
}

func getReader(scope string) (sr storage.Reader, err error) {
	readLock.RLock()
	sr, ok := readers[scope]
	if !ok {
		err = fmt.Errorf("Not initialized")
	}
	readLock.RUnlock()
	return
}

// SetReader sets a reader with scope
func SetReader(scope string, sr storage.Reader) {
	readLock.Lock()
	readers[scope] = sr
	readLock.Unlock()
	return
}

func flushAllWriter() (err error) {
	writeLock.Lock()
	writers = map[string]storage.Writer{}
	writeLock.Unlock()
	return
}

func getWriter(scope string) (sw storage.Writer, err error) {
	writeLock.RLock()
	sw, ok := writers[scope]
	if !ok {
		err = fmt.Errorf("Not initialized")
	}
	writeLock.RUnlock()
	return
}

// SetWriter sets a writer with scope
func SetWriter(scope string, sw storage.Writer) {
	writeLock.Lock()
	writers[scope] = sw
	writeLock.Unlock()
	return
}

func flushAllInitializer() (err error) {
	initLock.Lock()
	initializers = map[string]storage.Initializer{}
	initLock.Unlock()
	return
}

func getInitializer(scope string) (si storage.Initializer, err error) {
	initLock.RLock()
	si, ok := initializers[scope]
	if !ok {
		err = fmt.Errorf("Not initialized")
	}
	initLock.RUnlock()
	return
}

// SetInitializer sets an initializer with scope
func SetInitializer(scope string, si storage.Initializer) {
	initLock.Lock()
	initializers[scope] = si
	initLock.Unlock()
	return
}

func flushAllOauth() (err error) {
	oauthLock.Lock()
	oauths = map[string]oauth.Wrapper{}
	oauthLock.Unlock()
	return
}

func getOauth(scope string) (ow oauth.Wrapper, err error) {
	oauthLock.RLock()
	ow, ok := oauths[scope]
	if !ok {
		err = fmt.Errorf("Not initialized")
	}
	oauthLock.RUnlock()
	return
}

// SetOauth sets an oauth with scope
func SetOauth(scope string, ow oauth.Wrapper) {
	oauthLock.Lock()
	oauths[scope] = ow
	oauthLock.Unlock()
	return
}
