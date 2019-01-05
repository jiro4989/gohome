package gohome

import (
	"os"
	"testing"
)

func TestGetHomeDir(t *testing.T) {
	// normal
	os.Setenv("HOME", "/home/test")
	dir, err := GetHomeDir()
	if err != nil {
		t.Error("could not get HOME environment.")
	}
	if dir != "/home/test" {
		t.Error("could not get HOME environment.")
	}

	os.Setenv("HOME", "")
	os.Setenv("USERPROFILE", `C:\Users\test`)
	dir, err = GetHomeDir()
	if err != nil {
		t.Error("could not get USERPROFILE environment.")
	}
	if dir != `C:\Users\test` {
		t.Error("could not get USERPROFILE environment.")
	}

	// error
	os.Setenv("HOME", "")
	os.Setenv("USERPROFILE", "")
	dir, err = GetHomeDir()
	if err == nil {
		t.Error("error not occured.")
	}
	if dir != "" {
		t.Error("illegal environment value is got.")
	}

}

func TestGetConfigDir(t *testing.T) {
	os.Setenv("HOME", "/home/test")
	conf, err := GetConfigDir()
	if err != nil {
		t.Error("could not get config dir.")
	}
	if conf != "/home/test/.config" {
		t.Error("config dir path is illegal.")
	}

	os.Setenv("HOME", "")
	os.Setenv("USERPROFILE", `C:\Users\test`)
	conf, err = GetConfigDir()
	if err != nil {
		t.Error("could not get USERPROFILE environment.")
	}
	if conf != `C:\Users\test/.config` {
		t.Error("path is illegal.", conf)
	}

	os.Setenv("HOME", "")
	os.Setenv("USERPROFILE", "")
	conf, err = GetConfigDir()
	if err == nil {
		t.Error("error not occured.")
	}
	if conf != "" {
		t.Error("illegal environment value is got.")
	}
}
