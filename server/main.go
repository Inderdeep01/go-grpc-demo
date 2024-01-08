package main

import (
	"context"
	"log"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

var tracer trace.Tracer

const (
	port = ":8080"
)

func newExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
}

func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("ExampleService-Server"),
		),
	)

	if err != nil {
		panic(err)
	}

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}

func main() {
	ctx := context.Background()

	exporter, err := newExporter(ctx)
	if err != nil {
		log.Fatalf("Failed to initalize exporter: %v", err)
	}
	log.Println("Connected to exporter")

	// Trace provider
	tp := newTraceProvider(exporter) //otel.GetTracerProvider()
	// Handle shutdown
	defer func() {
		_ = tp.Shutdown(ctx)
	}()

	otel.SetTracerProvider(tp)

	// tracer to be used
	tracer = tp.Tracer("demo-service")

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Started server at %v\n", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
