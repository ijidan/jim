package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"jim/internal/global"
	"jim/pkg"
	"time"
)

func Jaeger() gin.HandlerFunc {
	return func(context *gin.Context) {
		var httpSpan opentracing.Span

		tracer, closer := pkg.NewJaeger(global.Config, "jim_api")
		defer func() {
			_ = closer.Close()
		}()
		spanContext, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(context.Request.Header))
		if err != nil {
			httpSpan = opentracing.StartSpan(context.Request.URL.Path)
			defer httpSpan.Finish()
		} else {
			httpSpan = opentracing.StartSpan(context.Request.URL.Path,
				opentracing.ChildOf(spanContext),
				opentracing.Tag{Key: "user_agent", Value: context.Request.UserAgent},
				opentracing.Tag{Key: string(ext.SpanKind), Value: "HTTP"},
				opentracing.StartTime(time.Now()))
			defer httpSpan.Finish()
		}
		context.Set("tracer_ctx", opentracing.ContextWithSpan(context, httpSpan))
		context.Set("tracer", tracer)

		context.Next()
	}
}
