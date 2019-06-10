package integrationtest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	expectedBanner := fmt.Sprintf("%s%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
		"<pre>",
		` _____________________________________________________________________________ `,
		`/  _   _      _ _         _   _                       ___    _______   _   _  \`,
		`| | | | | ___| | | ___   | |_| |__   ___ _ __ ___    |_ _|  / /___ /  | | | | |`,
		`| | |_| |/ _ \ | |/ _ \  | __| '_ \ / _ \ '__/ _ \    | |  / /  |_ \  | | | | |`,
		`| |  _  |  __/ | | (_) | | |_| | | |  __/ | |  __/_   | |  \ \ ___) | | |_| | |`,
		`| |_| |_|\___|_|_|\___/   \__|_| |_|\___|_|  \___( ) |___|  \_\____/   \___/  |`,
		`|                                                |/                           |`,
		`\                                                                             /`,
		` ----------------------------------------------------------------------------- `,
		`        \   ^__^`,
		`         \  (oo)\_______`,
		`            (__)\       )\/\`,
		`                ||----w |`,
		`                ||     ||`,
		"</pre>",
	)

	// run services
	go MainGateway()
	go MainFlowBanner()
	go MainFlowRoot()
	go MainServiceCmd()
	go MainServiceHTML()

	// wait for a while
	time.Sleep(100 * time.Millisecond)

	// emulate root request
	responseRoot, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Errorf("Get error %s", err)
	}
	defer responseRoot.Body.Close()
	bodyRoot, err := ioutil.ReadAll(responseRoot.Body)
	if err != nil {
		t.Errorf("Get error %s", err)
	}
	expectedRootResponse := "hello world"
	actualRootResponse := string(bodyRoot)
	if actualRootResponse != expectedRootResponse {
		t.Errorf("expected :\n%s, get :\n%s", expectedRootResponse, actualRootResponse)
	}

	// emulate banner request
	responseBanner, err := http.Get(fmt.Sprintf("http://localhost:8080/banner?text=%s", url.QueryEscape("Hello there, I <3 U")))
	if err != nil {
		t.Errorf("Get error %s", err)
	}
	defer responseBanner.Body.Close()
	bodyBanner, err := ioutil.ReadAll(responseBanner.Body)
	if err != nil {
		t.Errorf("Get error %s", err)
	}
	actualBanner := string(bodyBanner)
	if actualBanner != expectedBanner {
		t.Errorf("expected :\n%s, get :\n%s", expectedBanner, actualBanner)
	}

}
