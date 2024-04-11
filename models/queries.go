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
	"RE_SCAN_ACCOUNT": Object{
		Type: UPDATE,
		Params: []string{"status"},
	},
	"RESOURCE_COUNT": Object{
		Query: "SELECT COUNT(*) as total FROM resources",
		Type: RAW,
	},
}