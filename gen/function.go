package gen

import (
	"fmt"
	"github.com/state-alchemists/ayanami/generator"
	"log"
	"strings"
)

// ExposedFunction exposed ready function definition
type ExposedFunction struct {
	FunctionName       string
	FunctionReturn     string
	FunctionPackage    string
	FunctionAssignment string
	InputDeclaration   string
	OutputDeclaration  string
	Inputs             []string
	Outputs            []string
	JoinedInputs       string
	JoinedOutputs      string
}

// Function a definition of function
type Function struct {
	Inputs               []string
	Outputs              []string
	FunctionName         string
	FunctionPackage      string
	FunctionDependencies []string // should belong to the same package
	generator.StringHelper
}

// GetFileName get name of function file
func (f *Function) GetFileName() string {
	return fmt.Sprintf("%s.go", strings.ToLower(f.FunctionName))
}

// ToExposed change function to it exposed counterpart
func (f *Function) ToExposed() ExposedFunction {
	// get declaration
	inputDeclaration := fmt.Sprintf("%s interface{}", strings.Join(f.Inputs, ", "))
	var outputTypes []string
	var functionReturns []string
	for range f.Outputs {
		outputTypes = append(outputTypes, "interface{}")
		functionReturns = append(functionReturns, "nil")
	}
	outputDeclaration := strings.Join(outputTypes, ", ")
	functionReturn := strings.Join(functionReturns, ", ")
	// get assignment
	outputAssignment := strings.Join(f.Outputs, ", ")
	inputAssignment := strings.Join(f.Inputs, ", ")
	functionAssignment := fmt.Sprintf("%s := %s(%s)", outputAssignment, f.FunctionName, inputAssignment)
	return ExposedFunction{
		FunctionName:       f.FunctionName,
		FunctionReturn:     functionReturn,
		FunctionPackage:    f.FunctionPackage,
		FunctionAssignment: functionAssignment,
		InputDeclaration:   inputDeclaration,
		OutputDeclaration:  outputDeclaration,
		Inputs:             f.Inputs,
		Outputs:            f.Outputs,
		JoinedInputs:       f.QuoteArrayAndJoin(f.Inputs, ", "),
		JoinedOutputs:      f.QuoteArrayAndJoin(f.Outputs, ", "),
	}
}

// Validate validating an event
func (f *Function) Validate() bool {
	for _, input := range f.Inputs {
		if !f.IsMatch(input, "^[A-Za-z][a-zA-Z0-9]*$") {
			log.Printf("[ERROR] Invalid input `%s`", input)
			return false
		}
	}
	for _, output := range f.Outputs {
		if !f.IsMatch(output, "^[A-Za-z][a-zA-Z0-9]*$") {
			log.Printf("[ERROR] Invalid output `%s`", output)
			return false
		}
	}
	if !f.IsMatch(f.FunctionName, "^[A-Z][a-zA-Z0-9]*$") && f.FunctionPackage == "" {
		log.Println("[ERROR] Function name should not be alphanumeric and function package should not be empty")
		return false
	}
	return true
}

// GetFunctionFileName get name of function file
func (f *Function) GetFunctionFileName() string {
	return fmt.Sprintf("%s.go", strings.ToLower(f.FunctionName))
}

// NewFunction create new function
func NewFunction(functionPackage, functionName string, inputs, outputs, dependencies []string) Function {
	return Function{
		FunctionPackage:      functionPackage,
		FunctionName:         functionName,
		FunctionDependencies: dependencies,
		Inputs:               inputs,
		Outputs:              outputs,
	}
}
