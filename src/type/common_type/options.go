package common_type

type Options struct {
	InferOpt                    InferOpt                    `json:"infer_opt,omitempty"`
	TraceInferenceOpt           TraceInferenceOpt           `json:"trace_inference_opt,omitempty"`
	ExplainOpt                  ExplainOpt                  `json:"explain_opt,omitempty"`
	ParallelOpt                 ParallelOpt                 `json:"parallel_opt,omitempty"`
	PrefetchSizeOpt             PrefetchSizeOpt             `json:"prefetch_size_opt,omitempty"`
	PrefetchOpt                 PrefetchOpt                 `json:"prefetch_opt,omitempty"`
	SessionIdleTimeoutOpt       SessionIdleTimeoutOpt       `json:"session_idle_timeout_opt,omitempty"`
	TransactionTimeoutOpt       TransactionTimeoutOpt       `json:"transaction_timeout_opt,omitempty"`
	SchemaLockAcquireTimeoutOpt SchemaLockAcquireTimeoutOpt `json:"schema_lock_acquire_timeout_opt,omitempty"`
	ReadAnyReplicaOpt           ReadAnyReplicaOpt           `json:"read_any_replica_opt,omitempty"`
}

type InferOpt struct {
	Infer bool `json:"infer,omitempty"`
}

type TraceInferenceOpt struct {
	TraceInference bool `json:"trace_inference,omitempty"`
}

type ExplainOpt struct {
	Explain bool `json:"explain,omitempty"`
}

type ParallelOpt struct {
	Parallel bool `json:"parallel,omitempty"`
}

type PrefetchSizeOpt struct {
	PrefetchSize bool `json:"prefetch_size,omitempty"`
}

type PrefetchOpt struct {
	Prefetch bool `json:"prefetch,omitempty"`
}

type SessionIdleTimeoutOpt struct {
	SessionIdleTimeoutMillis bool `json:"session_idle_timeout_millis,omitempty"`
}

type TransactionTimeoutOpt struct {
	TransactionTimeoutMillis bool `json:"transaction_timeout_millis,omitempty"`
}

type SchemaLockAcquireTimeoutOpt struct {
	SchemaLockAcquireTimeoutMillis bool `json:"schema_lock_acquire_timeout_millis,omitempty"`
}

type ReadAnyReplicaOpt struct {
	ReadAnyReplica bool `json:"read_any_replica,omitempty"`
}
