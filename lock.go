package dynalock

import (
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	// DefaultLockBackOff if locing is unsuccessful then this backoff will be used
	DefaultLockBackOff = 3 * time.Second
)

type dynamodbLock struct {
	ddb      *Dynalock
	last     *KVPair
	renewCh  chan struct{}
	unlockCh chan struct{}

	key   string
	value *dynamodb.AttributeValue
	ttl   time.Duration
}

// Lock attempt to lock the DynamoDB record, this will BLOCK and retry at a rate of once every 3 seconds
func (l *dynamodbLock) Lock(stopChan chan struct{}) (<-chan struct{}, error) {
	lockHeld := make(chan struct{})

	success, err := l.tryLock(lockHeld, stopChan)
	if err != nil {
		return nil, err
	}
	if success {
		return lockHeld, nil
	}

	// FIXME: This really needs a jitter for backoff
	ticker := time.NewTicker(DefaultLockBackOff)

	for {
		select {
		case <-ticker.C:
			success, err := l.tryLock(lockHeld, stopChan)
			if err != nil {
				return nil, err
			}
			if success {
				return lockHeld, nil
			}
		case <-stopChan:
			return nil, ErrLockAcquireCancelled
		}
	}

}

// Unlock this will unlock and perfom a DELETE to remove the DynamoDB record
func (l *dynamodbLock) Unlock() error {
	l.unlockCh <- struct{}{}

	_, err := l.ddb.AtomicDelete(l.key, l.last)
	if err != nil {
		return err
	}
	l.last = nil

	return err
}

func (l *dynamodbLock) tryLock(lockHeld chan struct{}, stopChan chan struct{}) (bool, error) {
	success, new, err := l.ddb.AtomicPut(
		l.key,
		WriteWithPreviousKV(l.last),
		WriteWithAttributeValue(l.value),
		WriteWithTTL(l.ttl),
	)
	if err != nil {
		if err == ErrKeyNotFound || err == ErrKeyModified || err == ErrKeyExists {
			return false, nil
		}
		return false, err
	}
	if success {
		l.last = new
		// keep holding
		go l.holdLock(lockHeld, stopChan)
		return true, nil
	}

	return false, err
}

func (l *dynamodbLock) holdLock(lockHeld, stopChan chan struct{}) {
	defer close(lockHeld)

	hold := func() error {
		_, new, err := l.ddb.AtomicPut(
			l.key,
			WriteWithPreviousKV(l.last),
			WriteWithAttributeValue(l.value),
			WriteWithTTL(l.ttl),
		)
		if err == nil {
			l.last = new
		}
		return err
	}

	// may need a floor of 1 second set
	heartbeat := time.NewTicker(l.ttl / 3)
	defer heartbeat.Stop()

	for {
		select {
		case <-heartbeat.C:
			if err := hold(); err != nil {
				return
			}
		case <-l.renewCh:
			return
		case <-l.unlockCh:
			return
		case <-stopChan:
			return
		}
	}
}
