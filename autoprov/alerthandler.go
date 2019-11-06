package main

import (
	"context"

	"github.com/mobiledgex/edge-cloud/cloudcommon"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/log"
	"github.com/opentracing/opentracing-go"
)

func alertChanged(ctx context.Context, old *edgeproto.Alert, new *edgeproto.Alert) {
	if new == nil {
		return
	}
	name, ok := new.Labels["alertname"]
	if !ok {
		log.SpanLog(ctx, log.DebugLevelApi, "alertname not found", "alert", new)
		return
	}

	var handler func(ctx context.Context, name string, alert *edgeproto.Alert) error
	switch name {
	case cloudcommon.AlertAutoScaleUp:
		fallthrough
	case cloudcommon.AlertAutoScaleDown:
		handler = autoScale
	}

	if handler == nil {
		return
	}
	// make a copy since we spawn a thread to deal with it.
	alert := alertCopy(new)
	go func() {
		cspan := log.StartSpan(log.DebugLevelApi, "auto scale", opentracing.ChildOf(log.SpanFromContext(ctx).Context()))
		cctx := log.ContextWithSpan(context.Background(), cspan)
		defer cspan.Finish()
		err := handler(cctx, name, alert)
		log.SpanLog(cctx, log.DebugLevelApi, "handled alert", "alert", alert, "err", err)
	}()
}

func alertCopy(a *edgeproto.Alert) *edgeproto.Alert {
	alert := *a
	for k, v := range a.Labels {
		alert.Labels[k] = v
	}
	for k, v := range a.Annotations {
		alert.Annotations[k] = v
	}
	return &alert
}