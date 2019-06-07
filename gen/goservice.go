package gen

import (
	"fmt"
	"github.com/state-alchemists/ayanami/generator"
	"log"
	"path/filepath"
)

// ExposedGoServiceConfig exposed ready flowConfig
type ExposedGoServiceConfig struct {
	ServiceName string
	PackageName string
	Functions   map[string]ExposedFunction
}

// GoServiceConfig configuration to generate GoService
type GoServiceConfig struct {
	ServiceName string
	PackageName string
	Functions   map[string]Function
	*generator.IOHelper
	generator.StringHelper
}

// Set replace/add service's function
func (config *GoServiceConfig) Set(method string, function Function) {
	config.Functions[method] = function
}

// Validate validating config
func (config *GoServiceConfig) Validate() bool {
	log.Printf("[INFO] Validating %s", config.ServiceName)
	if config.IsAlphaNumeric(config.ServiceName) {
		log.Printf("[ERROR] Service name should be alphanumeric, but `%s` found", config.ServiceName)
		return false
	}
	if config.PackageName == "" {
		log.Println("[ERROR] Package name should not be empty")
		return false
	}
	for methodName, function := range config.Functions {
		if !config.IsAlphaNumeric(methodName) {
			log.Printf("[ERROR] method should be alphanumeric, but `%s` found", methodName)
			return false
		}
		if !function.Validate() {
			return false
		}
	}
	return true
}

// Scaffold scaffolding config
func (config *GoServiceConfig) Scaffold() error {
	for _, function := range config.Functions {
		data := function.ToExposed()
		packageSourcePath := function.FunctionPackage
		functionFileName := function.GetFunctionFileName()
		// write function
		functionSourcePath := filepath.Join(packageSourcePath, functionFileName)
		if !config.IsSourceExists(functionSourcePath) {
			log.Printf("[INFO] Create %s", functionFileName)
			err := config.WriteSource(functionSourcePath, "srvc.function.go", data)
			if err != nil {
				return err
			}
		} else {
			log.Printf("[INFO] %s exists", functionFileName)
		}
		// write dependencies
		for _, dependency := range function.FunctionDependencies {
			dependencySourcePath := filepath.Join(packageSourcePath, dependency)
			if !config.IsSourceExists(dependencySourcePath) {
				log.Printf("[INFO] Create %s", dependency)
				err := config.WriteSource(dependencySourcePath, "dependency.go", data)
				if err != nil {
					return err
				}
			} else {
				log.Printf("[INFO] %s exists", dependency)
			}
		}
	}
	return nil
}

// Build building config
func (config *GoServiceConfig) Build() error {
	log.Printf("[INFO] Building %s", config.ServiceName)
	depPath := fmt.Sprintf("srvc-%s", config.ServiceName)
	// write functions and dependencies
	for _, function := range config.Functions {
		packageSourcePath := function.FunctionPackage
		packageDepPath := filepath.Join(depPath, function.FunctionPackage)
		functionFileName := function.GetFunctionFileName()
		// copy function
		log.Printf("[INFO] Create %s", functionFileName)
		functionSourcePath := filepath.Join(packageSourcePath, functionFileName)
		functionDepPath := filepath.Join(packageDepPath, functionFileName)
		err := config.CopySourceToDep(functionSourcePath, functionDepPath)
		if err != nil {
			return err
		}
		// copy dependencies
		for _, dependency := range function.FunctionDependencies {
			log.Printf("[INFO] Create %s", dependency)
			dependencySourcePath := filepath.Join(packageSourcePath, dependency)
			dependencyDepPath := filepath.Join(packageDepPath, dependency)
			err := config.CopySourceToDep(dependencySourcePath, dependencyDepPath)
			if err != nil {
				return err
			}
		}
	}
	// write main.go
	log.Println("[INFO] Create main.go")
	mainPath := filepath.Join(depPath, "main.go")
	err := config.WriteDep(mainPath, "srvc.main.go", config.toExposed())
	if err != nil {
		return err
	}
	// write go.mod
	log.Println("[INFO] Create go.mod")
	goModPath := filepath.Join(depPath, "go.mod")
	err = config.WriteDep(goModPath, "go.mod", config.PackageName)
	return err
}

func (config *GoServiceConfig) toExposed() ExposedGoServiceConfig {
	exposedFunctions := make(map[string]ExposedFunction)
	for methodName, function := range config.Functions {
		exposedFunctions[methodName] = function.ToExposed()
	}
	return ExposedGoServiceConfig{
		PackageName: config.PackageName,
		ServiceName: config.ServiceName,
		Functions:   exposedFunctions,
	}
}

// NewGoService create new goservice
func NewGoService(ioHelper *generator.IOHelper, packageName, serviceName string, functions map[string]Function) GoServiceConfig {
	return GoServiceConfig{
		PackageName: packageName,
		ServiceName: serviceName,
		Functions:   functions,
		IOHelper:    ioHelper,
	}
}

// NewEmptyGoService create new empty service
func NewEmptyGoService(ioHelper *generator.IOHelper, packageName, serviceName string) GoServiceConfig {
	return NewGoService(ioHelper, packageName, serviceName, make(map[string]Function))
}
