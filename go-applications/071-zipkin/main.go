package main

import (
	"log"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/trace/zipkin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel@v1.15.0"
)

func main() {
	// Create and register a Zipkin exporter
	exporter, err := zipkin.NewExportPipeline(
		"your-zipkin-server:9411",
		"go-trace-app",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer exporter.Shutdown()

	// Create a tracer provider with the exporter
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)
	otel.SetTracerProvider(tracerProvider)

	// Retrieve a tracer from the tracer provider
	tracer := otel.Tracer("go-trace-app")

	// Create a simple HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a new context with a span
		ctx, span := tracer.Start(r.Context(), "http_request")
		defer span.End()

		// Perform some actions
		span.AddEvent("handling_request")
		span.SetAttributes(
			attribute.String("http.method", r.Method),
			attribute.String("http.url", r.URL.String()),
		)

		// Simulate some work
		for i := 0; i < 10000000; i++ {
			// Do some computation
		}

		// Send the response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	})

	// Start the HTTP server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
