package notify

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Notify struct {
	events map[string][]chan interface{}
	lock sync.RWMutex
}

func NewNotify() *Notify {
	n := new(Notify)
	n.events = make(map[string][]chan interface{})
	return n
}

func (n *Notify) start(key string, outputChan chan interface{}) {
	n.lock.Lock()
	defer n.lock.Unlock()
	n.events[key] = append(n.events[key], outputChan)
}

func (n *Notify) Stop(key string, outputChan chan interface{}) error {
	n.lock.Lock()
	defer n.lock.Unlock()

	newArray := make([]chan interface{}, 0)
	outChans, ok := n.events[key]
	if !ok {
		return fmt.Errorf("not found")
	}

	for _, ch := range outChans {
		if ch != outputChan {
			newArray = append(newArray, ch)
		} else {
			close(ch)
		}
	}
	n.events[key] = newArray

	return nil

}

func (n *Notify) StopAll(key string) error {
	n.lock.Lock()
	defer n.lock.Unlock()

	outChans, ok := n.events[key]
	if !ok {
		return errors.New("not found")
	}
	for _, ch := range outChans {
		close(ch)
	}
	delete(n.events, key)

	return nil

}

func (n *Notify) Post(key string, data interface{}) error {
	n.lock.RLock()
	defer n.lock.RUnlock()

	outChans, ok := n.events[key]
	if !ok {
		return fmt.Errorf("not found")
	}

	for _, outChan := range outChans {
		outChan <- data
	}
	return nil
}

func (n *Notify) PostTimeout(key string, data interface{}, timeout time.Duration) error {
	n.lock.RLock()
	defer n.lock.RUnlock()

	outChans, ok :=n.events[key]
	if !ok {
		return fmt.Errorf("not found")
	}
	for _, outputChan := range outChans {
		select {
		case outputChan <- data:
		case <-time.After(timeout):
		}
	}

	return nil
}
