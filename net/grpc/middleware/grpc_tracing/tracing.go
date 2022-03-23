package grpc_tracing

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport/zipkin"
	"google.golang.org/grpc/metadata"
	"io"
	"time"
)

type TextMapWriter struct {
	metadata.MD
}

func (t TextMapWriter) Set(key, val string) {
	t.MD[key] = append(t.MD[key], val)
}

type TextMapReader struct {
	metadata.MD
}

func (t TextMapReader) ForeachKey(handler func(key, val string) error) error {
	for key, val := range t.MD {
		for _, v := range val {
			if err := handler(key, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// zipkinHost: http://hostname:9411/api/v1/spans
func InitJaeger(service string, zipkinHost string) (tracer opentracing.Tracer, closer io.Closer, err error) {
	//cfg := &config.Configuration{
	//	Sampler: &config.SamplerConfig{
	//		Type:  "const",
	//		Param: 1,
	//	},
	//	Reporter: &config.ReporterConfig{
	//		LogSpans:           true,
	//		LocalAgentHostPort: jaegerAgentHost,
	//	},
	//}
	sender, _ := zipkin.NewHTTPTransport(fmt.Sprintf("http://%s:9411/api/v1/spans", zipkinHost))
	//sender, _ := jaeger.NewUDPTransport("jaeger-agent.istio-system:5775", 0)
	tracer, closer = jaeger.NewTracer(
		service,
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(
			sender,
			jaeger.ReporterOptions.BufferFlushInterval(1*time.Second)),
	)
	//tracer, closer, err = cfg.New(service, config.Logger(jaeger.StdLogger))
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, err
}
