package infrastructure

var API = map[string]string{
	// user
	"reputation": "https://stackoverflow.com/users?tab=reputation",
	"newUser":    "https://stackoverflow.com/users?tab=newusers",
	"voters":     "https://stackoverflow.com/users?tab=voters",
	"editors":    "https://stackoverflow.com/users?tab=editors",
	"moderators": "https://stackoverflow.com/users?tab=moderators",

	// tag
	"tags": "https://stackoverflow.com/tags",

	// job
	"jobs": "https://stackoverflow.com/jobs?med=site-ui&ref=jobs-tab",

	// question
	"search": "https://stackoverflow.com/search?q=%s",
}
