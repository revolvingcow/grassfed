package app

import (
    "strings"
    "github.com/revel/revel"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
        ContentTypeFilter,
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB())
	// revel.OnAppStart(FillCache())
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
    // This was commented out so I could mask the origins with DNS
    //c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")

    // Add some common security headers
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

var ContentTypeFilter = func(c *revel.Controller, fc []revel.Filter) {
    path := c.Request.Request.URL.Path
    formats := []string{"json", "xml"}

    for _, format := range formats {
        if strings.HasSuffix(path, "." + format) {
            trimmed := strings.TrimSuffix(path, "." + format)
            c.Request.Request.URL.Path = trimmed
            c.Request.Request.RequestURI = trimmed
            c.Request.Format = format
            break
        }
    }

    fc[0](c, fc[1:])
}
