// Code generated by counterfeiter. DO NOT EDIT.
package auctioncellrepfakes

import (
	"sync"

	"code.cloudfoundry.org/executor/containermetrics"
	"code.cloudfoundry.org/rep"
)

type FakeContainerMetricsProvider struct {
	MetricsStub        func() map[string]*containermetrics.CachedContainerMetrics
	metricsMutex       sync.RWMutex
	metricsArgsForCall []struct{}
	metricsReturns     struct {
		result1 map[string]*containermetrics.CachedContainerMetrics
	}
	metricsReturnsOnCall map[int]struct {
		result1 map[string]*containermetrics.CachedContainerMetrics
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerMetricsProvider) Metrics() map[string]*containermetrics.CachedContainerMetrics {
	fake.metricsMutex.Lock()
	ret, specificReturn := fake.metricsReturnsOnCall[len(fake.metricsArgsForCall)]
	fake.metricsArgsForCall = append(fake.metricsArgsForCall, struct{}{})
	fake.recordInvocation("Metrics", []interface{}{})
	fake.metricsMutex.Unlock()
	if fake.MetricsStub != nil {
		return fake.MetricsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.metricsReturns.result1
}

func (fake *FakeContainerMetricsProvider) MetricsCallCount() int {
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	return len(fake.metricsArgsForCall)
}

func (fake *FakeContainerMetricsProvider) MetricsReturns(result1 map[string]*containermetrics.CachedContainerMetrics) {
	fake.MetricsStub = nil
	fake.metricsReturns = struct {
		result1 map[string]*containermetrics.CachedContainerMetrics
	}{result1}
}

func (fake *FakeContainerMetricsProvider) MetricsReturnsOnCall(i int, result1 map[string]*containermetrics.CachedContainerMetrics) {
	fake.MetricsStub = nil
	if fake.metricsReturnsOnCall == nil {
		fake.metricsReturnsOnCall = make(map[int]struct {
			result1 map[string]*containermetrics.CachedContainerMetrics
		})
	}
	fake.metricsReturnsOnCall[i] = struct {
		result1 map[string]*containermetrics.CachedContainerMetrics
	}{result1}
}

func (fake *FakeContainerMetricsProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerMetricsProvider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ rep.ContainerMetricsProvider = new(FakeContainerMetricsProvider)