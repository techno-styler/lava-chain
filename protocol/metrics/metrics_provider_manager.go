package metrics

import (
	"net/http"
	"sync"

	"github.com/lavanet/lava/utils"
	spectypes "github.com/lavanet/lava/x/spec/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	MetricsListenFlagName = "metrics-listen-address"
	DisabledFlagOption    = "disabled"
)

type ProviderMetricsManager struct {
	providerMetrics           map[string]*ProviderMetrics
	lock                      sync.RWMutex
	totalCUServicedMetric     *prometheus.GaugeVec
	totalCUPaidMetric         *prometheus.GaugeVec
	totalRelaysServicedMetric *prometheus.GaugeVec
	totalErroredMetric        *prometheus.GaugeVec
	consumerQoSMetric         *prometheus.GaugeVec
}

func NewProviderMetricsManager(networkAddress string) *ProviderMetricsManager {
	if networkAddress == DisabledFlagOption {
		utils.LavaFormatWarning("prometheus endpoint inactive, option is disabled", nil)
		return nil
	}
	totalCUServicedMetric := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "provider_total_cu_serviced",
		Help: "The total number of CUs serviced by the provider over time.",
	}, []string{"spec", "apiInterface"})

	// Create a new GaugeVec metric to represent the TotalCUPaid over time.
	totalCUPaidMetric := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "provider_total_cu_paid",
		Help: "The total number of CUs paid to the provider over time.",
	}, []string{"spec", "apiInterface"})

	// Create a new GaugeVec metric to represent the TotalRelaysServiced over time.
	totalRelaysServicedMetric := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "provider_total_relays_serviced",
		Help: "The total number of relays serviced by the provider over time.",
	}, []string{"spec", "apiInterface"})

	// Create a new GaugeVec metric to represent the TotalErrored over time.
	totalErroredMetric := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "provider_total_errored",
		Help: "The total number of errors encountered by the provider over time.",
	}, []string{"spec", "apiInterface"})

	consumerQoSMetric := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "Consumer_QoS",
		Help: "The latest QoS score from a consumer",
	}, []string{"spec", "consumer_address", "qos_metric"})
	// Register the metrics with the Prometheus registry.
	prometheus.MustRegister(totalCUServicedMetric)
	prometheus.MustRegister(totalCUPaidMetric)
	prometheus.MustRegister(totalRelaysServicedMetric)
	prometheus.MustRegister(totalErroredMetric)
	prometheus.MustRegister(consumerQoSMetric)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		utils.LavaFormatInfo("prometheus endpoint listening", utils.Attribute{Key: "Listen Address", Value: networkAddress})
		http.ListenAndServe(networkAddress, nil)
	}()
	return &ProviderMetricsManager{
		providerMetrics:           map[string]*ProviderMetrics{},
		totalCUServicedMetric:     totalCUServicedMetric,
		totalCUPaidMetric:         totalCUPaidMetric,
		totalRelaysServicedMetric: totalRelaysServicedMetric,
		totalErroredMetric:        totalErroredMetric,
		consumerQoSMetric:         consumerQoSMetric,
	}
}
func (pme *ProviderMetricsManager) getProviderMetric(specID string, apiInterface string) *ProviderMetrics {
	return pme.providerMetrics[specID+apiInterface]
}

func (pme *ProviderMetricsManager) setProviderMetric(metrics *ProviderMetrics) {
	specID := metrics.specID
	apiInterface := metrics.apiInterface
	pme.providerMetrics[specID+apiInterface] = metrics
}

func (pme *ProviderMetricsManager) AddProviderMetrics(specID string, apiInterface string) *ProviderMetrics {
	if pme == nil {
		return nil
	}
	pme.lock.Lock()
	defer pme.lock.Unlock()
	if pme.getProviderMetric(specID, apiInterface) == nil {
		providerMetric := NewProviderMetrics(specID, apiInterface, pme.totalCUServicedMetric, pme.totalCUPaidMetric, pme.totalRelaysServicedMetric, pme.totalErroredMetric, pme.consumerQoSMetric)
		pme.setProviderMetric(providerMetric)
	}
	return pme.getProviderMetric(specID, apiInterface)
}

func (pme *ProviderMetricsManager) AddPayment(specID string, cu uint64) {
	if pme == nil {
		return
	}
	availableAPIInterface := []string{
		spectypes.APIInterfaceJsonRPC,
		spectypes.APIInterfaceTendermintRPC,
		spectypes.APIInterfaceRest,
		spectypes.APIInterfaceGrpc,
	}
	for _, apiInterface := range availableAPIInterface {
		providerMetrics := pme.getProviderMetric(specID, apiInterface)
		if providerMetrics != nil {
			providerMetrics.AddPayment(cu)
		}
	}
}
