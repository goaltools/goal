// Package watcher is used for watching files and directories
// for automatic recompilation and restart of app on change
// when in development mode.
package watcher

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/colegion/goal/utils/log"

	"gopkg.in/fsnotify.v1"
)

// Type is a watcher type that allows registering new
// pattern - actions pairs.
type Type struct {
	mu    sync.Mutex
	files map[string]bool
}

// NewType allocates and returns a new instance of watcher Type.
func NewType() *Type {
	return &Type{
		files: map[string]bool{},
	}
}

// Listen gets a pattern and a function. The function will be executed
// when files matching the pattern will be modified.
func (t *Type) Listen(pattern string, fn func()) *fsnotify.Watcher {
	// Create a new watcher.
	w, err := fsnotify.NewWatcher()
	log.AssertNil(err)

	// Find directories matching the pattern.
	ds := glob(pattern)

	// Add the files to the watcher.
	for i := range ds {
		log.Trace.Printf(`Adding "%s" to the list of watched directories...`, ds[i])
		err := w.Add(ds[i])
		if err != nil {
			log.Warn.Println(err)
		}
	}

	// Start watching process.
	go t.NotifyOnUpdate(filepath.ToSlash(pattern), w, fn)
	return w
}

// ListenFile is equivalent of Listen but for files.
// If file is added using ListenFile and the same file
// is withing a pattern of Listen, only the first one
// will trigger restarts.
// I.e. we have the following calls:
//	w.Listen("./", fn1)
//	w.ListenFile("./goal.yml", fn2)
// If "goal.yml" file is modified fn2 will be triggered.
// fn1 may be triggered by changes in any file inside
// "./" directory except "goal.yml".
func (t *Type) ListenFile(path string, fn func()) *fsnotify.Watcher {
	// Create a new watcher.
	w, err := fsnotify.NewWatcher()
	log.AssertNil(err)

	// Watch a directory instead of file.
	// See issue #17 of fsnotify to find out more
	// why we do this.
	dir := filepath.Dir(path)
	w.Add(dir)

	// Clean path and replace back slashes
	// to the normal ones.
	path = filepath.ToSlash(path)

	// Start watching process.
	t.files[path] = true
	go t.NotifyOnUpdate(path, w, fn)
	return w
}

// NotifyOnUpdate starts the function every time a file change
// event is received. Start it as a goroutine.
func (t *Type) NotifyOnUpdate(pattern string, watcher *fsnotify.Watcher, fn func()) {
	for {
		select {
		case ev := <-watcher.Events:
			// Convert path to the Linux format.
			name := filepath.ToSlash(ev.Name)

			// Make sure this is the exact event type that
			// requires a restart.
			if !restartRequired(ev) {
				continue
			}

			// If this is a directory watcher, but a file that was registered
			// with ListenFile has been modified,
			// ignore this event.
			if !t.files[pattern] && t.files[name] {
				continue
			}

			// If this is a single file watcher, make sure this is
			// exactly the file that should be watched, not
			// some other.
			if t.files[pattern] && name != pattern {
				continue
			}

			// Trigger the registered functions.
			t.mu.Lock()
			fn()
			t.mu.Unlock()
		case <-watcher.Errors:
			return
		}
	}
}

// restartRequired checks whether event indicates a file
// has been modified. If so, it returns true.
func restartRequired(event fsnotify.Event) bool {
	// Do not restart if "./bin" directory is modified.
	// TODO: make this configurable.
	d := filepath.ToSlash(event.Name)
	if d == "./bin" || d == "bin" {
		return false
	}

	if event.Op&fsnotify.Chmod == fsnotify.Chmod {
		return false
	}

	log.Trace.Printf(`FS object "%s" has been modified, restarting...`, event.Name)
	return true
}

// glob returns names of all directories matching pattern or nil.
// The only supported special character is an asterisk at the end.
// It means that the directory is expected to be scanned recursively.
// There is no way for fsnotify to watch individual files (see #17),
// so we support only directories.
// File system errors such as I/O reading are ignored.
func glob(pattern string) (ds []string) {
	// Make sure pattern is not empty.
	l := len(pattern)
	if l == 0 {
		return
	}

	// Check whether we should scan the directory recursively.
	recurs := pattern[l-1] == '*'
	if recurs {
		// Trim the asterisk at the end.
		pattern = pattern[:l-1]
	}

	// Make sure such path exists and it is a directory rather than a file.
	info, err := os.Stat(pattern)
	if err != nil {
		return
	}
	if !info.IsDir() {
		log.Warn.Printf(`"%s" is not a directory, skipping it.`, pattern)
		return
	}

	// If not recursive scan was expected, return the path as is.
	if !recurs {
		ds = append(ds, pattern)
		return // Return as is.
	}

	// Start searching directories recursively.
	filepath.Walk(pattern, func(path string, info os.FileInfo, err error) error {
		// Make sure there are no any errors.
		if err != nil {
			return err
		}

		// Make sure the path represents a directory.
		if !info.IsDir() {
			return nil
		}

		// Add the directory path to the list.
		ds = append(ds, path)
		return nil
	})
	return
}
