package sharedObjects

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type TimeStampAsH1 struct {
	app.Compo
	manager       *Manager
	timeStampAsH1 string
}

// Manager is the main controller of this application, also the root Body
type Manager struct {
	app.Compo
	searchBar *SearchBar
	//instanceTable *InstanceTable
	TimeStampAsH1 *TimeStampAsH1
}

type SearchBar struct {
	app.Compo
	manager      *Manager
	searchString string
}
