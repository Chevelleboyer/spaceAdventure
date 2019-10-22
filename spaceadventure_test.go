package main

import (
	"testing"
	//"github.com/stretchr/testify/assert"
	)

func TestChooseRandPlanet(t *testing.T) {
	seedJson := map[string]interface{} {
  		"mars": "Planet red",
  		"pluto": "Planet small",
  		"neptune": "Veronica mars",
	}

	randPlanet := ChooseRandPlanet(seedJson)
	if randPlanet != "Veronica mars" {
		t.Errorf("ChooseRandPlanet(seedJson) = %d; wanted Veronica mars", randPlanet)
	}
}

//Not sure how to accomadate user input with unit tests ??
// func TestLookupPlanet(t *testing.T) {
// 	seedJson := map[string]interface{} {
//   		"mars": "Planet red",
//   		"pluto": "Planet small",
//   		"neptune": "Veronica mars",
// 	}

// 	planet := LookupPlanet(seedJson)
// 	assert.Equal(t, planet, "Planet red")
// }