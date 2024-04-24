package models

var Queries = Data{
	"LIST": Object{
		Type: SELECT,
	},
	"INSERT": Object{
		Type:  INSERT,
	},
	"DELETE": Object{
		Type:  DELETE,
	},
	"UPDATE_ACCOUNT": Object{
		Type:   UPDATE,
		Params: []string{"name", "provider", "credentials"},
	},
	"UPDATE_ALERT": Object{
		Type: UPDATE,
		Params: []string{"name", "type", "budget", "usage", "endpoint", "secret"},
	},
	"UPDATE_VIEW": Object{
		Type: UPDATE,
		Params: []string{"name", "filters", "exclude"},
	},
	"UPDATE_VIEW_EXCLUDE": Object{
		Type: UPDATE,
		Params: []string{"exclude"},
	},
	"RE_SCAN_ACCOUNT": Object{
		Type: UPDATE,
		Params: []string{"status"},
	},
	"RESOURCE_COUNT": Object{
		Query: "SELECT COUNT(*) as total FROM resources",
		Type: RAW,
	},
	"UPDATE_TAGS": Object{
		Type: UPDATE,
		Params: []string{"tags"},
	},
}