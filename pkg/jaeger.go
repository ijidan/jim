package pkg

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	jConfig "jim/config"
	"sync"
)

const LocalAgentHostPort = "192.168.33.10:6831"

type jer struct {
	Tracer opentracing.Tracer
	Closer io.Closer
	Span   opentracing.Span
}

//初始化
func (j *jer) init(conf *jConfig.Config, service string) (opentracing.Tracer, io.Closer) {
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
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Error: connot init Jaeger: %v\n", err))
	}
	j.Tracer = tracer
	j.Closer = closer
	return tracer, closer
}

func (j *jer) CreateSpan(spanName string) opentracing.Span {
	j.Span = j.Tracer.StartSpan(spanName)
	return j.Span
}

var (
	onceJer     sync.Once
	instanceJer *jer
)

func GetJaegerInstance(conf *jConfig.Config, service string, spanNameRoot string) *jer {
	onceJer.Do(func() {
		instanceJer = &jer{}
		instanceJer.init(conf, service)
		instanceJer.CreateSpan(spanNameRoot)
	})
	return instanceJer
}
