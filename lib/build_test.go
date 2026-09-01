package lib

import "testing"

// TestBuildValidates is the headless smoke test: it composes the full
// game and asks the library to cross-check every name reference
// between managers. No window opens.
func TestBuildValidates(t *testing.T) {
	g := Build()
	if err := g.Validate(); err != nil {
		t.Fatalf("validate: %v", err)
	}
	for _, name := range []string{"bedroom", "kitchen"} {
		if !g.SceneManager.Has(name) {
			t.Errorf("missing scene %q", name)
		}
	}
	for _, name := range []string{"key", "beans", "mug"} {
		if !g.ItemManager.Has(name) {
			t.Errorf("missing item %q", name)
		}
	}
	if !g.DialogManager.Has("cat_morning") {
		t.Errorf("missing dialog cat_morning")
	}
	if !g.StageManager.Has("game") {
		t.Errorf("missing stage game")
	}
	stage, ok := g.StageManager.Get(g.StartStage)
	if !ok || !stage.World {
		t.Errorf("StartStage %q is not a registered world stage", g.StartStage)
	}
	if !g.ScriptManager.Has("intro") || !g.ScriptManager.Has("victory") {
		t.Errorf("missing intro/victory script")
	}
}
