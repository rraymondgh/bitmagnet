package batch

import (
	"time"

	"github.com/bitmagnet-io/bitmagnet/internal/classifier"
	"github.com/bitmagnet-io/bitmagnet/internal/model"
	"github.com/bitmagnet-io/bitmagnet/internal/processor"
	"github.com/bitmagnet-io/bitmagnet/internal/protocol"
)

const MessageName = "process_torrent_batch"

type MessageParams struct {
	InfoHashGreaterThan protocol.ID             `json:"InfoHashGreaterThan,omitempty"`
	UpdatedBefore       time.Time               `json:"UpdatedBefore,omitempty"`
	ClassifyMode        processor.ClassifyMode  `json:"ClassifyMode,omitempty"`
	ClassifierWorkflow  string                  `json:"ClassifierWorkflow,omitempty"`
	ClassifierFlags     classifier.Flags        `json:"ClassifierFlags,omitempty"`
	ChunkSize           uint                    `json:"ChunkSize,omitempty"`
	BatchSize           uint                    `json:"BatchSize,omitempty"`
	ContentTypes        []model.NullContentType `json:"ContentTypes,omitempty"`
	Orphans             bool                    `json:"Orphans,omitempty"`
}

func (p MessageParams) ApisDisabled() bool {
	if p.ClassifierFlags == nil {
		return false
	}

	enabledAny, ok := p.ClassifierFlags["apis_enabled"]
	if !ok {
		return false
	}

	enabled, ok := enabledAny.(bool)

	return ok && !enabled
}

func NewQueueJob(msg MessageParams, options ...model.QueueJobOption) (model.QueueJob, error) {
	if msg.BatchSize == 0 {
		msg.BatchSize = 100
	}

	if msg.ChunkSize == 0 {
		msg.ChunkSize = 10_000
	}

	return model.NewQueueJob(
		MessageName,
		msg,
		append([]model.QueueJobOption{model.QueueJobMaxRetries(2)}, options...)...,
	)
}
