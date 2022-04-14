package messages_test

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/utils/messages"
)

var message = &messages.Message{
	Session: &discordgo.Session{},
	MC:      &discordgo.MessageCreate{},
	Content: "prefix!testCommand argument1 argument2 argument3",
	Prefix:  "prefix!",
}

func Test_GetCommandWithPrefix(t *testing.T) {
	expected := "prefix!testCommand"
	val := message.GetCommandWithPrefix()

	if expected != val {
		t.Fatalf("Test_GetCommandWithPrefix failed. Expected: %s - Value received: %s", expected, val)
		t.Fail()
	} else {
		t.Log("Test_GetCommandWithPrefix passed the test succesfully.")
	}
}

func Test_GetCommand(t *testing.T) {
	expected := "testCommand"
	val := message.GetCommand()

	if expected != val {
		t.Fatalf("Test_GetCommand failed. Expected: %s - Value received: %s", expected, val)
		t.Fail()
	} else {
		t.Log("Test_GetCommand passed the test succesfully.")
	}
}

func Test_HasCommand(t *testing.T) {
	expected := true
	val := message.HasCommand()

	if expected != val {
		t.Fatalf("Test_HasCommand failed. Expected: %t - Value received: %t", expected, val)
		t.FailNow()
	} else {
		t.Log("Test_HasCommand passed the test succesfully.")
	}
}

func Test_GetArguments(t *testing.T) {
	expected := []string{"argument1", "argument2", "argument3"}
	val := message.GetArguments()
	err := false

	for i := range val {
		if expected[i] != val[i] {
			err = true
		}
	}

	if err {
		t.Fatalf("Test_GetArguments failed. Arguments are not the same.")
		t.Fail()
	} else {
		t.Log("Test_GetArguments passed the test succesfully.")
	}
}

func Test_GetPlainContent(t *testing.T) {
	expected := "argument1 argument2 argument3"
	val := message.GetPlainContent()

	if expected != val {
		t.Fatalf("Test_GetPlainContent failed. Expected: %s - Value received: %s", expected, val)
		t.Fail()
	} else {
		t.Log("Test_GetPlainContent passed the test succesfully.")
	}
}

func Test_HasHelpPetition(t *testing.T) {
	expected := false
	val := message.HasHelpPetition()

	if expected != val {
		t.Fatalf("Test_HasHelPetition failed.")
		t.Fail()
	} else {
		t.Log("Test_HasHelpPetition passed the test succesfully.")
	}
}
