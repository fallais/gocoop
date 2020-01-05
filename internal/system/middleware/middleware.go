package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

//------------------------------------------------------------------------------
// Middleware
//------------------------------------------------------------------------------

// Cors is a Goji middleware to handle CORS (from the great Zenithar)
func Cors() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logrus.WithFields(logrus.Fields{
				"middleware": "CORS",
			}).Debugln("Entering middleware")

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			// Stop here if its Preflighted OPTIONS request
			if r.Method == "OPTIONS" {
				return
			}

			// Serve
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

// XRequestID is a goji middleware to track requestID (from the great Zenithar)
func XRequestID() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logrus.WithFields(logrus.Fields{
				"middleware": "XRequestID",
			}).Debugln("Entering middleware")

			if len(r.Header.Get("X-Request-ID")) > 0 {
				// Prepare the context
				ctx := context.WithValue(r.Context(), "reqID", r.Header.Get("X-Request-ID"))

				// Serve
				h.ServeHTTP(w, r.WithContext(ctx))
			} else {
				// Serve
				h.ServeHTTP(w, r)
			}
		}

		return http.HandlerFunc(fn)
	}
}

// Metrics is a GOJI middleware to provide Prometheus metrics
func Metrics() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logrus.WithFields(logrus.Fields{
				"middleware": "Metrics",
			}).Debugln("Entering middleware")

			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
