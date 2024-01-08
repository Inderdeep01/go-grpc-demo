package main

import (
	"context"
	"log"

	pb "github.com/inderdeep01/go-grpc-demo-yt/proto"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var tracer trace.Tracer

const (
	port = ":8080"
)

type textMapCarrier struct {
	propagation.TextMapCarrier
}

func newExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
}

func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("ExampleService"),
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
	ctxParent := context.Background()

	exporter, err := newExporter(ctxParent)
	if err != nil {
		log.Fatalf("Failed to initalize exporter: %v", err)
	}

	// Trace provider
	tp := newTraceProvider(exporter)

	// Handle shutdown
	defer func() {
		_ = tp.Shutdown(ctxParent)
	}()

	otel.SetTracerProvider(tp)

	// tracer to be used
	tracer = tp.Tracer("demo-service")
	/* traceContext := propagation.TraceContext{}
	TextMapCarrier := textMapCarrier{}
	TextMapCarrier.Set("key", "value")
	traceContext.Inject(ctxParent, TextMapCarrier) */

	conn, err := grpc.DialContext(ctxParent, "localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	names := &pb.NameList{
		Names: []string{"Inderdeep", "Singh", "Sidhu", "Golang", "Alice", "Bob"},
	}
	if client == nil {
		log.Fatalf("client is nil")
	} else {
		ctxParentClient, span := tracer.Start(ctxParent, "Client Initialized")
		defer span.End()
		log.Printf("client %v\n", client)
		ctx, span := tracer.Start(ctxParentClient, "Sending Unary Request")
		callSayHello(ctx, client) //-> call to unary function
		span.End()
		//callSayHelloServerStream(client, names) //-> call to server stream function
		//callSayHelloClientStream(client, names) //-> call to client stream function
		ctx, span = tracer.Start(ctxParentClient, "Sending Bidirectional Request")
		callSayHelloBidirectionalStream(ctx, client, names) //-> call to bidirectional stream function
		span.End()
	}
}
