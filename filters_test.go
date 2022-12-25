package degussa

import (
	"testing"
)

func checkMaterial(t *testing.T, m Material) {
	for _, l := range ByMaterial(m) {
		t.Run(l.Name, func(t *testing.T) {
			if l.Material != m {
				t.Errorf("material is %d, want %d", l.Material, l.Material)
			}
		})
	}
}

func TestMaterials(t *testing.T) {
	checkMaterial(t, Gold)
	checkMaterial(t, Silver)
	checkMaterial(t, Platinum)
	checkMaterial(t, Palladium)
}
