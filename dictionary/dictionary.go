// © 2018 Imhotep Software LLC. All rights reserved.

// Package dictionary represent for a corpus of words
// that can be randomly picked.
package dictionary

// YOUR_CODE...

// Load loads a new dictionary given a base path and a dictionary name
func Load(dir, name string) (wl WordList, e error) {
	// YOUR_CODE...
}

// LoadDefault loads a dictionary by name from the default assets dir
func LoadDefault(name string) (wl WordList, e error) {
	// YOUR_CODE...
}

// Pick select a word at random from a list of words
func Pick(l WordList) string {
	// YOUR_CODE...
}
