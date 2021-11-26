package pkg

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	jConfig "jim/config"
)

func NewJaeger(conf *jConfig.Config, service string) (opentracing.Tracer, io.Closer) {
	localAgentHostPort := fmt.Sprintf("%s:%d", conf.Jaeger.Host, conf.Jaeger.Port)
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: localAgentHostPort,
		},
	}
	//tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
