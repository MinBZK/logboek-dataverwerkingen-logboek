package http

import (
	"net/http"

	"github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go"
)

var propegator = logboek.TraceContextPropegator{}

func NewHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := propegator.Extract(r.Context(), r.Header)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

type Transport struct {
	rt http.RoundTripper
}

func NewTransport(rt http.RoundTripper) *Transport {
	if rt == nil {
		rt = http.DefaultTransport
	}

	return &Transport{
		rt: rt,
	}
}

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	ctx := r.Context()
	r = r.Clone(ctx)

	propegator.Inject(ctx, r.Header)

	return t.rt.RoundTrip(r)
}
