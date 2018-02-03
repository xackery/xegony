package cases

import (
	"fmt"
	"sync"

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

// InitializeAll readers, writers, and initializers
func InitializeAll(sr storage.Reader, sw storage.Writer, si storage.Initializer) (err error) {
	scopes := []string{
		"account",
		"character",
		"spell",
		"rule",
		"ruleEntry",
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

func getOauth(scope string) (ow oauth.Wrapper, err error) {
	writeLock.RLock()
	ow, ok := oauths[scope]
	if !ok {
		err = fmt.Errorf("Not initialized")
	}
	writeLock.RUnlock()
	return
}

// SetOauth sets an oauth with scope
func SetOauth(scope string, ow oauth.Wrapper) {
	writeLock.Lock()
	oauths[scope] = ow
	writeLock.Unlock()
	return
}
