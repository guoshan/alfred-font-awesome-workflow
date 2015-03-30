package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCommands_Find(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &Command{outStream: outStream, errStream: errStream}
	terms := []string{"app"}

	status := cmd.Find(terms)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := `<?xml version="1.0" encoding="UTF-8"?><items><item valid="true" arg="apple" uid="f179" unicode="f179"><title>apple</title><subtitle>Paste class name: fa-apple</subtitle><icon>./icons/fa-apple.png</icon></item><item valid="true" arg="whatsapp" uid="f232" unicode="f232"><title>whatsapp</title><subtitle>Paste class name: fa-whatsapp</subtitle><icon>./icons/fa-whatsapp.png</icon></item></items>`
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommands_Put_NameFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &Command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"name": "apple"}

	status := cmd.Put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := "fa-apple"
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommands_Put_CodeFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &Command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"code": "apple"}

	status := cmd.Put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := "f179"
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommands_Put_RefFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &Command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"ref": "apple"}

	status := cmd.Put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := ""
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommands_Put_URLFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &Command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"url": "apple"}

	status := cmd.Put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := "http://fontawesome.io/icon/apple/"
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}