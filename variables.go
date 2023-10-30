/*
 * Copyright (c) Brandon Jordan
 */

package main

var variables map[string]variableValue

type variableValue struct {
	variableType string
	valueType    tokenType
	value        any
	getAs        string
	coerce       string
	constant     bool
	repeatItem   bool
}

var globals map[string]variableValue

func makeGlobals() {
	if len(globals) != 0 {
		return
	}
	globals = make(map[string]variableValue)
	globals["ShortcutInput"] = variableValue{
		variableType: "ExtensionInput",
		valueType:    String,
		value:        "ShortcutInput",
	}
	globals["CurrentDate"] = variableValue{
		variableType: "CurrentDate",
		valueType:    Date,
		value:        "CurrentDate",
	}
	globals["Clipboard"] = variableValue{
		variableType: "Clipboard",
		valueType:    String,
		value:        "Clipboard",
	}
	globals["Device"] = variableValue{
		variableType: "DeviceDetails",
		valueType:    String,
		value:        "DeviceDetails",
	}
	globals["Ask"] = variableValue{
		variableType: "Ask",
		valueType:    String,
		value:        "Ask",
	}
	globals["RepeatItem"] = variableValue{
		variableType: "Variable",
		valueType:    String,
		value:        "Repeat Item",
	}
	globals["RepeatIndex"] = variableValue{
		variableType: "Variable",
		valueType:    String,
		value:        "Repeat Index",
	}
}

func validReference(identifier string) bool {
	if _, found := globals[identifier]; found {
		isInputVariable(identifier)
		return true
	}
	if _, found := variables[identifier]; found {
		return true
	}

	return false
}