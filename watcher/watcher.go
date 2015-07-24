// Package watcher is used for watching files and directories
// for automatic recompilation and restart of app on change
// when in development mode.
package watcher

import (
	"path/filepath"
	"sync"

	"github.com/anonx/sunplate/log"

	"gopkg.in/fsnotify.v1"
)

// Type is a watcher type that allows registering new
// pattern - actions pairs.
type Type struct {
	mu       sync.Mutex
	watchers []*fsnotify.Watcher
}

// NewType allocates and returns a new instance of watcher Type.
func NewType() *Type {
	return &Type{}
}

// Listen gets a pattern and a function. The function will be executed
// when files matching the pattern will be modified.
func (t *Type) Listen(pattern string, fn func()) {
	// Create a new watcher.
	w, err := fsnotify.NewWatcher()
	log.AssertNil(err)

	// Replace the unbuffered Event channel with a buffered one.
	// Otherwise multiple change events only come out one at a time.
	w.Events = make(chan fsnotify.Event, 100)
	w.Errors = make(chan error, 10)

	// Find files matching the pattern.
	ms, err := filepath.Glob(pattern)
	if err != nil {
		log.Error.Panicf("Pattern `%s` is malformed. Error: %v.", pattern, err)
	}

	// Add the files to the watcher.
	for i := range ms {
		log.Trace.Printf(`Adding "%s" to the list of watched files & directories...`, ms[i])
		err := w.Add(ms[i])
		if err != nil {
			log.Warn.Println(err)
		}
	}

	// Add watcher to the list.
	i := len(t.watchers)
	t.watchers = append(t.watchers, w)

	// Start watching process.
	go t.NotifyOnUpdate(i, fn)
}

// NotifyOnUpdate starts the function every time a file change
// event is received. Start it as a goroutine.
func (t *Type) NotifyOnUpdate(watcherIndex int, fn func()) {
	for {
		select {
		case ev := <-t.watchers[watcherIndex].Events:
			log.Trace.Println(ev.String())
			if restartRequired(ev) {
				t.mu.Lock()
				fn()
				t.mu.Unlock()
			}
		case err := <-t.watchers[watcherIndex].Errors:
			log.Warn.Println(err)
			continue
		}
	}
}

// restartRequired checks whether event indicates a file
// has been modified. If so, it returns true.
func restartRequired(event fsnotify.Event) bool {
	if event.Op&fsnotify.Chmod == fsnotify.Chmod {
		return true
	}
	return false
}
