package sopt_test

import (
	"testing"

	"github.com/grimdork/sopt"
)

func TestRemoveFirstGroup(t *testing.T) {
	opt := sopt.New()
	opt.AddGroup("Two")
	opt.AddGroup("Three")
	t.Logf("%v", opt)
	opt.RemoveGroup("default")
	t.Logf("%v", opt)
	if opt.GroupCount() != 2 {
		t.Errorf("Expected 2 groups, but got %d", opt.GroupCount())
		t.Fail()
	}

	g := opt.GetGroup("default")
	if g != nil {
		t.Errorf("Group 'default' should not exist, but does.")
		t.Fail()
	}
}

func TestRemoveMiddleGroup(t *testing.T) {
	opt := sopt.New()
	opt.AddGroup("Two")
	opt.AddGroup("Three")
	t.Logf("%v", opt)
	opt.RemoveGroup("Two")
	t.Logf("%v", opt)
	if opt.GroupCount() != 2 {
		t.Errorf("Expected 2 groups, but got %d", opt.GroupCount())
		t.Fail()
	}

	g := opt.GetGroup("Two")
	if g != nil {
		t.Errorf("Group 'Two' should not exist, but does.")
		t.Fail()
	}
}

func TestRemoveLastGroup(t *testing.T) {
	opt := sopt.New()
	opt.AddGroup("Two")
	opt.AddGroup("Three")
	t.Logf("%v", opt)
	opt.RemoveGroup("Three")
	t.Logf("%v", opt)
	if opt.GroupCount() != 2 {
		t.Errorf("Expected 2 groups, but got %d", opt.GroupCount())
		t.Fail()
	}

	g := opt.GetGroup("Three")
	if g != nil {
		t.Errorf("Group 'Three' should not exist, but does.")
		t.Fail()
	}
}

func TestBool(t *testing.T) {
	opt := sopt.New()
	opt.SetOption("", "v", "verbose", "Show more details in output.", false, false, sopt.VarTypeBool, nil)
}

func TestString(t *testing.T) {
	opt := sopt.New()
	opt.SetOption("", "f", "file", "Full file path.", "", false, sopt.VarTypeString, nil)
}

func TestInt(t *testing.T) {
	opt := sopt.New()
	opt.SetOption("", "p", "port", "Port number.", 3000, false, sopt.VarTypeInt, nil)
}

func TestSortGroup(t *testing.T) {
	opt := sopt.New()
	opt.SetOption("", "v", "verbose", "Show more details in output.", false, false, sopt.VarTypeBool, nil)
	opt.SetOption("", "f", "file", "Full file path.", "", false, sopt.VarTypeString, nil)
	opt.SetOption("", "p", "port", "Port number.", 0, false, sopt.VarTypeInt, nil)
	opt.SetDefaultHelp()
	g := opt.GetGroup("default")
	if g == nil {
		t.Errorf("Group 'default' should exist, but does not.")
		t.FailNow()
	}

	t.Log("Unsorted:")
	for _, o := range g.GetOptions() {
		t.Logf("ShortName: %s LongName: %s", o.ShortName, o.LongName)
	}

	g.Sort()
	t.Log("Sorted:")
	list := g.GetOptions()
	for _, o := range list {
		t.Logf("ShortName: %s LongName: %s", o.ShortName, o.LongName)
	}

	if list[0].ShortName != "f" || list[1].ShortName != "h" || list[2].ShortName != "p" || list[3].ShortName != "v" {
		t.Log("Sort order not the same as expected.")
		t.FailNow()
	}
}

func TestAutoGroup(t *testing.T) {
	opt := sopt.New()
	opt.SetOption("General", "v", "verbose", "Show more details in output.", false, false, sopt.VarTypeBool, nil)
	list := opt.GetGroups()
	if list[1].Name != "General" {
		t.Errorf("Expected 'General' group, but got %s", list[1].Name)
		t.Fail()
	}
}
