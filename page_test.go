package main

import (
	"regexp"
	"testing"
)

// createTestPageTree creates a test page tree.
// It returns the root page, and a deep child page.
func createTestPageTree() (*page, *page) {
	rootPage := newPageTree()
	deepChild := rootPage.addChild("about").addChild("something").addChild("cool")
	rootPage.addChild("contact")
	blogPage := rootPage.addChild("blog")
	blogPage.addChild("post1")
	blogPage.addChild("post2")

	return rootPage, deepChild
}

// TestGetRootpage calls page.getRootPage to test that the rootPage is returned
func TestGetRootPage(t *testing.T) {
	want := regexp.MustCompile("/$")
	_, deepChild := createTestPageTree()
	result := deepChild.getRootPage()

	if !want.MatchString(result.name) {
		t.Errorf("got %q, wanted %q", result.name, want)
	}
}

// TestBuildFullUri calls page.buildFullUri to test that a full uri is built for that page node
func TestBuildFullUri(t *testing.T) {
	want := regexp.MustCompile("^about/something/cool$")
	_, deepChild := createTestPageTree()
	result := deepChild.buildFullUri()

	if !want.MatchString(result) {
		t.Errorf("Wanted full Uri %q, received %q", want, result)
	}
}

// TestGetChildByName calls page.getChildByName to test that a child with a given name can be returned
func TestGetChildByName(t *testing.T) {
	root, _ := createTestPageTree()
	result, _ := root.getChildByName("about")

	if result == nil {
		t.Errorf("Wanted about child page, received nil")
	}
}

// TestGetNoChildByName calls page.getChildByName to test that if no child with that name exists, an error is returned
func TestGetNoChildByName(t *testing.T) {
	root, _ := createTestPageTree()
	_, err := root.getChildByName("missing")

	if err == nil {
		t.Errorf("Wanted an error, received nil")
	}
}

// TestGetByUriShallow calls page.getByUri with a shallow match
func TestGetByUriShallow(t *testing.T) {
	uri := "contact"
	want := regexp.MustCompile("^contact$")
	root, _ := createTestPageTree()
	result, _ := root.getByUri(uri)

	if !want.MatchString(result.name) {
		t.Errorf("got %q, wanted %q", result.name, want)
	}
}

// TestGetByUriDeep calls page.GetByUriPath with a deep match
func TestGetByUriDeep(t *testing.T) {
	uri := "about/something/cool"
	want := regexp.MustCompile("^cool$")
	root, _ := createTestPageTree()
	result, _ := root.getByUri(uri)

	if !want.MatchString(result.name) {
		t.Errorf("got %q, wanted %q", result.name, want)
	}
}

// TestCrawlTreeByUriShallow calls page.crawlTreeByUri with a shallow match
func TestCrawlTreeByUriShall(t *testing.T) {
	uri := "contact"
	want := regexp.MustCompile("^contact$")
	root, _ := createTestPageTree()
	result, remainder, _ := root.crawlTreeByUri(uri)

	if !want.MatchString(result.name) {
		t.Errorf("Wanted uri %q, received %q", want, result.name)
	}

	if remainder != nil {
		t.Errorf("Wanted no remainder, received %q", *remainder)
	}
}

// TestCrawlTreeByUriDeep calls page.crawlTreeByUri with a deep match
func TestCrawlTreeByUriDeep(t *testing.T) {
	uri := "about/something/cool"
	want := regexp.MustCompile("^cool$")
	root, _ := createTestPageTree()
	result, remainder, _ := root.crawlTreeByUri(uri)

	if !want.MatchString(result.name) {
		t.Errorf("Wanted uri %q, received %q", want, result.name)
	}

	if remainder != nil {
		t.Errorf("Wanted no remainder, received %q", *remainder)
	}
}

// TestCrawlTreeByUriIncomplete calls page.crawlTreeByUri with a incomplete match
func TestCrawlTreeByUriIncomplete(t *testing.T) {
	uri := "about/something/nonsense"
	wantName := regexp.MustCompile("^something$")
	wantRemainder := regexp.MustCompile("^nonsense$")
	root, _ := createTestPageTree()
	result, remainder, _ := root.crawlTreeByUri(uri)

	if !wantName.MatchString(result.name) {
		t.Errorf("Wanted uri %q, received %q", wantName, result.name)
	}

	if !wantRemainder.MatchString(*remainder) {
		t.Errorf("Wanted remainder %q, received %q", wantRemainder, *remainder)
	}
}

//TODO:
// TestAddToTreeFromUrI calls page.addToTreeFromUri to ensure valid nodes are added to the tree

// TestFilePathToUri calls filePathToUri to ensure a valid uri is created from a filePath
func TestFilePathToUri(t *testing.T) {
	filePath := "pages/blog/category/post1.gohtml"
	want := regexp.MustCompile("^blog/category/post1$")
	result := filePathToUri(filePath)

	if !want.MatchString(result) {
		t.Errorf("got %q, wanted %q", result, want)
	}
}

// TestIndexFilePathToUri calls filePathToUri to ensure a valid uri is created from an index filePath
func TestIndexFilePathToUri(t *testing.T) {
	filePath := "pages/blog/category/index.gohtml"
	want := regexp.MustCompile("^blog/category$")
	result := filePathToUri(filePath)

	if !want.MatchString(result) {
		t.Errorf("got %q, wanted %q", result, want)
	}
}
