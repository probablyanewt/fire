package main

import (
	"regexp"
	"testing"
)

func createTestPage() *Page {
	aboutSomethingPage := Page{
		uri: "/about/something",
	}

	aboutPage := Page{
		uri: "/about",
	}
	aboutPage.AddChild(&aboutSomethingPage)

	contactPage := Page{
		uri: "/contact",
	}

	rootPage := Page{
		uri: "/",
	}
	rootPage.AddChild(&aboutPage)
	rootPage.AddChild(&contactPage)

	return &rootPage
}

// TestGetByUriPathShallow calls Page.GetByUriPath with a shallow match
func TestGetByUriPathShallow(t *testing.T) {
	uri := "/contact"
	want := regexp.MustCompile(uri)
	result, _ := createTestPage().GetByUri(uri)

	if !want.MatchString(result.uri) {
		t.Errorf("got %q, wanted %q", result.uri, want)
	}
}

// TestGetByUriPathDeep calls Page.GetByUriPath with a deep match
func TestGetByUriPathDeep(t *testing.T) {
	uri := "/about/something"
	want := regexp.MustCompile(uri)
	result, _ := createTestPage().GetByUri(uri)
	println(result.uri)
	if !want.MatchString(result.uri) {
		t.Errorf("got %q, wanted %q", result.uri, want)
	}
}
