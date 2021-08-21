package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	http.Handle("/ping", WithLogging(pingHandler()))
	addr := "127.0.0.1:8080"
	logrus.WithField("addr", addr).Info("starting server")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		logrus.WithField("event", "start server").Fatal(err)
	}
}

func pingHandler() http.Handler {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		key, ok := r.URL.Query()["id"]
		if ok {
			rw.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintf(rw, "duy %v", key)
		}
	}
	return http.HandlerFunc(fn)
}

type (
	// struct for holding response details
	responseData struct {
		status int
		size   int
	}

	// our http.ResponseWriter implementation
	loggingResponseWriter struct {
		http.ResponseWriter // compose original http.ResponseWriter
		responseData        *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b) // write response using original http.ResponseWriter
	r.responseData.size += size            // capture size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode) // write status code using original http.ResponseWriter
	r.responseData.status = statusCode       // capture status code
}

func WithLogging(h http.Handler) http.Handler {
	loggingFn := func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lrw := loggingResponseWriter{
			ResponseWriter: rw, // compose original http.ResponseWriter
			responseData:   responseData,
		}
		h.ServeHTTP(&lrw, req) // inject our implementation of http.ResponseWriter
		query := req.ParseForm()
		duration := time.Since(start)

		logrus.WithFields(logrus.Fields{
			"uri":      req.RequestURI,
			"method":   req.Method,
			"query":    query,
			"status":   responseData.status,
			"duration": duration,
			"size":     responseData.size,
		}).Info("request completed")
	}
	return http.HandlerFunc(loggingFn)
}
