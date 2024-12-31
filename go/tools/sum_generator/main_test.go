package main

import "testing"

func TestSpec(t *testing.T) {
	type Case struct {
		in    string
		specs []Spec
	}
	spec := NewSpec(" GlobalDef : _sum: *A | *B")
	if spec.exportedName != "GlobalDef" {
		t.Errorf("expected exportedName to be GlobalDef, got %s", spec.exportedName)
	}
	if spec.internalName != "_sum" {
		t.Errorf("expected internalName to be _sum, got %s", spec.internalName)
	}
	if len(spec.types) != 2 {
		t.Errorf("expected len(spec.types) to be 2, got %d", len(spec.types))
	}
	if spec.types[0] != "*A" {
		t.Errorf("expected spec.types[0] to be *A, got %s", spec.types[0])
	}
	if spec.types[1] != "*B" {
		t.Errorf("expected spec.types[1] to be *B, got %s", spec.types[1])
	}
}
