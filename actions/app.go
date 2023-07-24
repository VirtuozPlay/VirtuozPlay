package actions

import (
	csrf "github.com/gobuffalo/mw-csrf"
	"github.com/rs/cors"
	"net/http"
	"strconv"
	"virtuozplay/locales"
	"virtuozplay/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	i18n "github.com/gobuffalo/mw-i18n/v2"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var ForceSSL = loadForceSSL()

var (
	app *buffalo.App
	T   *i18n.Translator
)

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		options := buffalo.Options{
			Env:         ENV,
			SessionName: "_virtuozplay_session",
			// Setup CORS to allow all localhost:* origins in dev mode
			PreWares: []buffalo.PreWare{corsConfiguration().Handler},
		}
		if ENV != "production" {
			// Allow incoming LAN connections
			options.Addr = ":3000"
		}
		app = buffalo.New(options)

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Wraps each request in a transaction.
		//   c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))
		// Setup and use translations:
		app.Use(translations())

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		homeHandler := csrf.New(HomeHandler)
		app.GET("/", homeHandler)
		app.GET("/about", homeHandler)
		app.GET("/checkup", homeHandler)
		app.GET("/collection", homeHandler)
		app.GET("/profile", homeHandler)
		app.GET("/stats", homeHandler)
		app.GET("/play", homeHandler)
		app.GET("/play/{performanceId}", homeHandler)

		// GraphQL endpoints
		// TODO add CSRF protection
		app.ANY("/graphql", GraphQLHandler)
		if ENV != "production" {
			app.ANY("/graphql/playground", GraphQLPlaygroundHandler)
		}

		app.ServeFiles("/", http.FS(DistFS())) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(locales.FS(), "en-US"); err != nil {
		err := app.Stop(err)
		if err != nil {
			panic(err)
		}
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ForceSSL && ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

func loadForceSSL() bool {
	raw := envy.Get("GO_FORCE_SSL", "1")
	value := true

	if raw == "" {
		value = false
	} else {
		asInt, err := strconv.Atoi(raw)
		if err == nil {
			value = asInt != 0
		}
	}

	return value
}

// corsConfiguration returns the CORS configuration for VirtuozPlay
// All other origins are blocked by default unless in development mode where
// localhost:* is allowed.
func corsConfiguration() *cors.Cors {
	if ENV == "production" {
		return cors.Default()
	}
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:*",
			"https://localhost:*",
		},
	})
}
