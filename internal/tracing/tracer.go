package tracing

import (
	settings "bbone/internal/config"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"net"
)

func InitTracing() io.Closer {
	jaegerConfig := settings.GetConfig().Jaeger

	cfgMetrics := &config.Configuration{
		ServiceName: jaegerConfig.Name,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: net.JoinHostPort(jaegerConfig.Host, jaegerConfig.Port),
		},
	}

	tracer, closer, err := cfgMetrics.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Err(err).Msgf("failed init jaeger: %v", err)
	}
	opentracing.SetGlobalTracer(tracer)

	return closer
}
