## Publish

```go
type JobMessageType struct {
	ID                  uint64 `redis:"ID" json:"ID"`
	UID                 uint64 `redis:"UID" json:"UID"`
	ActionID            uint64 `redis:"actionID" json:"actionID"`
	FileID              string `redis:"fileID" json:"fileID"`
	Params              string `redis:"params" json:"params"`
}

func PublishJobMessage(message JobMessageType) (err error) {
	messageByte, err := json.Marshal(message)
	if err != nil {
		return
	}

	err = rabbitMQ.Publish(
		envVariable.MessageQueueConfig.JobExchangeName,
		envVariable.MessageQueueConfig.JobExchangeType,
		"",
		messageByte,
	)
	return
}
```

## Consume

```go
// event_audit_consumer.go
func eventAuditConsumer() {
	// subscribe event_audit direct
	rabbitMQ.Consumer(
		envVariable.MessageQueueConfig.EventAuditExchangeName,
		envVariable.MessageQueueConfig.EventAuditExchangeType,
		envVariable.MessageQueueConfig.EventAuditQueue,
		"event_audit",
		eventAuditService.MessageHandler,
	)
}

// consumer.go
func Consumer() {
	go eventAuditConsumer()
}

// main.go
func main() {
	// other code...
	
	go Consumer()

	// Run gin server
	_ = r.Run("0.0.0.0:" + port)
}
```