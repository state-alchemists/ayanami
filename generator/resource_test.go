package generator

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestResource(t *testing.T) {
	_, currentDirPath, _, ok := runtime.Caller(0)
	if !ok {
		t.Error("fail to locate directory")
	}
	ayanamiDirPath := filepath.Dir(filepath.Dir(currentDirPath))
	dummyProjectPath := filepath.Join(ayanamiDirPath, "generator", "dummyProject")
	// resource
	resource, err := NewResourceByProjectPath(dummyProjectPath)
	if err != nil {
		t.Errorf("Get error: %s", err)
	}

	// check template
	if resource.GetTemplate() == nil {
		t.Errorf("template should not be nil")
	}
	// check sourcePath
	expectedSourceCodePath := filepath.Join(dummyProjectPath, "sourcecode")
	actualSourceCodePath := resource.GetSourcePath()
	if expectedSourceCodePath != actualSourceCodePath {
		t.Errorf("expected `%s`, get `%s`", expectedSourceCodePath, actualSourceCodePath)
	}
	// check depPath
	expectedDeployablePath := filepath.Join(dummyProjectPath, "deployable")
	actualDeployablePath := resource.GetDepPath()
	if expectedDeployablePath != actualDeployablePath {
		t.Errorf("expected `%s`, get `%s`", expectedDeployablePath, actualDeployablePath)
	}

	// check existance
	expectedExistance := true
	actualExistance := resource.IsSourceExists("exists.py")
	if expectedExistance != actualExistance {
		t.Errorf("expected `%t`, get `%t`", expectedExistance, actualExistance)
	}
	expectedExistance = true
	actualExistance = resource.IsDepExists("exists.py")
	if expectedExistance != actualExistance {
		t.Errorf("expected `%t`, get `%t`", expectedExistance, actualExistance)
	}

	// check non-existance
	expectedExistance = false
	actualExistance = resource.IsSourceExists("not-exists.py")
	if expectedExistance != actualExistance {
		t.Errorf("expected `%t`, get `%t`", expectedExistance, actualExistance)
	}
	expectedExistance = false
	actualExistance = resource.IsDepExists("not-exists.py")
	if expectedExistance != actualExistance {
		t.Errorf("expected `%t`, get `%t`", expectedExistance, actualExistance)
	}

	// writeSource (scaffold)
	err = resource.WriteSource("hello.py", "hello.py", "world")
	if err != nil {
		t.Errorf("Get error: %s", err)
	}
	// check content of sourcecode/hello.py
	expectedSourceHelloPyPath := filepath.Join(dummyProjectPath, "sourcecode", "hello.py")
	defer os.RemoveAll(expectedSourceHelloPyPath)
	sourceHelloPyByte, err := ioutil.ReadFile(expectedSourceHelloPyPath)
	if err != nil {
		t.Errorf("Get error: %s", err)
	}
	actualSourceHelloPyContent := string(sourceHelloPyByte)
	expectedSourceHelloPyContent := "print(\"hello world\")"
	if expectedSourceHelloPyContent != actualSourceHelloPyContent {
		t.Errorf("expected `%s`, get `%s`", expectedSourceHelloPyContent, actualSourceHelloPyContent)
	}

	// copySourceToDeployable (build)
	err = resource.CopySourceToDep("hello.py", "hello.py")
	if err != nil {
		t.Errorf("Get error: %s", err)
	}
	// check content of deployable/hello.py
	expectedDeployedHelloPyPath := filepath.Join(dummyProjectPath, "deployable", "hello.py")
	defer os.RemoveAll(expectedDeployedHelloPyPath)
	deployedHelloPyByte, err := ioutil.ReadFile(expectedDeployedHelloPyPath)
	if err != nil {
		t.Errorf("Get error: %s", err)
	}
	actualDeployedHelloPyContent := string(deployedHelloPyByte)
	expectedDeployedHelloPyContent := "print(\"hello world\")"
	if expectedDeployedHelloPyContent != actualDeployedHelloPyContent {
		t.Errorf("expected `%s`, get `%s`", expectedDeployedHelloPyContent, actualDeployedHelloPyContent)
	}

	// writeDep (build)
	resource.WriteDep("fruit.py", "fruit.py", []string{"orange", "grape", "strawberry"})
	// check content of deployable/fruit.py
	expectedDeployedFruitPyPath := filepath.Join(dummyProjectPath, "deployable", "fruit.py")
	defer os.RemoveAll(expectedDeployedFruitPyPath)
	deployedFruitPyByte, err := ioutil.ReadFile(expectedDeployedFruitPyPath)
	if err != nil {
		t.Errorf("Get error: %s", err)
	}
	actualDeployedFruitPyContent := string(deployedFruitPyByte)
	expectedDeployedFruitPyContent := "print(\"orange grape strawberry \")"
	if expectedDeployedFruitPyContent != actualDeployedFruitPyContent {
		t.Errorf("expected `%s`, get `%s`", expectedDeployedFruitPyContent, actualDeployedFruitPyContent)
	}

}
