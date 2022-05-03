package sopt_test

import (
	"fmt"
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

func TestLongShort(t *testing.T) {
	err := sopt.New().SetOption("", "verbose", "", "", false, false, sopt.VarTypeBool, nil)
	if err == nil {
		t.Errorf("Expected error, but long short worked.")
		t.Fail()
	} else {
		t.Log("Long short failed as expected.")
	}
}

func TestShortLong(t *testing.T) {
	err := sopt.New().SetOption("", "", "v", "", false, false, sopt.VarTypeBool, nil)
	if err == nil {
		t.Errorf("Expected error, but short long worked.")
		t.Fail()
	} else {
		t.Log("Short long failed as expected.")
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
	err := opt.SetOption("", "p", "port", "Port number.", 3000, false, sopt.VarTypeInt, nil)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err.Error())
		t.FailNow()
	}

	args := []string{"-p", "4000"}
	err = opt.ParseArgs(args)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err.Error())
		t.FailNow()
	}

	if opt.GetInt("p") != 4000 {
		t.Errorf("Expected -p=4000, but got %d", opt.GetInt("p"))
		opt.ShowOptions()
		t.Fail()
	}
}

const moo = `                 (__)
                 (oo)
           /------\/
          / |    ||
         *  /\---/\
            ~~   ~~
..."Have you mooed today?"...
`

func moocmd(args []string) error {
	println(moo)
	fmt.Printf("Args: %+v\n", args)
	return nil
}

func TestCommand(t *testing.T) {
	opt := sopt.New()
	opt.SetCommand("moo", "Have you mooed today?", "", moocmd, nil)
	args := []string{"moo", "--help"}
	err := opt.ParseArgs(args)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err.Error())
		t.Fail()
	}
}
