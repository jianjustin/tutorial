package metric

import (
	"fmt"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/metrics/teststat"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"testing"
)

func TestCounter(t *testing.T) {
	s := httptest.NewServer(promhttp.HandlerFor(stdprometheus.DefaultGatherer, promhttp.HandlerOpts{}))
	defer s.Close()

	scrape := func() string {
		resp, _ := http.Get(s.URL)
		buf, _ := ioutil.ReadAll(resp.Body)
		return string(buf)
	}

	namespace, subsystem, name := "ns", "ss", "foo"
	re := regexp.MustCompile(namespace + `_` + subsystem + `_` + name + `{alpha="alpha-value",beta="beta-value"} ([0-9\.]+)`)

	counter := prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      name,
		Help:      "This is the help string.",
	}, []string{"alpha", "beta"}).With("beta", "beta-value", "alpha", "alpha-value") // order shouldn't matter

	value := func() float64 {
		matches := re.FindStringSubmatch(scrape())
		f, _ := strconv.ParseFloat(matches[1], 64)
		return f
	}
	counter.Add(1)
	if err := teststat.TestCounter(counter, value); err != nil {
		fmt.Println(err)
	}
}
