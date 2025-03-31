package main

import (
	"encoding/json"
	"syscall/js"
)

func formatJSON(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return "Error: No input provided"
	}

	input := args[0].String()
	if input == "" {
		return "Error: Input is empty"
	}

	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return "Error: Invalid JSON - " + err.Error()
	}

	formatted, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "Error: Failed to format JSON - " + err.Error()
	}

	return string(formatted)
}

func main() {
	js.Global().Set("formatJSON", js.FuncOf(formatJSON))

	input := js.Global().Get("document").Call("getElementById", "input")
	output := js.Global().Get("document").Call("getElementById", "output")

	input.Call("addEventListener", "input", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		inputValue := input.Get("value").String()
		formatted := formatJSON(js.Value{}, []js.Value{js.ValueOf(inputValue)})
		output.Set("value", formatted)
		return nil
	}))

	select {}
}
