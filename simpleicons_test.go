// Copyright (c) 2026 the go-icons authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package simpleicons

import (
	"strings"
	"testing"
)

func TestIconKnown(t *testing.T) {
	for _, n := range []string{"reddit", "instagram", "x", "mastodon", "lemmy", "bluesky", "rss", "ycombinator", "tiktok"} {
		s := Icon(n)
		if s == "" || !strings.Contains(s, "<svg") {
			t.Errorf("Icon(%q) returned no SVG", n)
		}
		if !Has(n) {
			t.Errorf("Has(%q) = false", n)
		}
	}
}

func TestIconUnknown(t *testing.T) {
	if Icon("definitely-not-a-brand") != "" {
		t.Error("unknown icon should return an empty string")
	}
	if Has("definitely-not-a-brand") {
		t.Error("Has of an unknown icon should be false")
	}
}

func TestNames(t *testing.T) {
	names := Names()
	if len(names) < 1000 {
		t.Fatalf("expected the full Simple Icons set, got %d names", len(names))
	}
	for i, n := range names {
		if i > 0 && names[i-1] > n {
			t.Fatalf("Names() not sorted at %d", i)
		}
		if !Has(n) {
			t.Errorf("listed name %q does not resolve", n)
		}
	}
}
