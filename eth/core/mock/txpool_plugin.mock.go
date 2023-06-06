// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/ethereum/go-ethereum/core/types"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/txpool"
	"sync"
)

// Ensure, that TxPoolPluginMock does implement core.TxPoolPlugin.
// If this is not the case, regenerate this file with moq.
var _ core.TxPoolPlugin = &TxPoolPluginMock{}

// TxPoolPluginMock is a mock implementation of core.TxPoolPlugin.
//
//	func TestSomethingThatUsesTxPoolPlugin(t *testing.T) {
//
//		// make and configure a mocked core.TxPoolPlugin
//		mockedTxPoolPlugin := &TxPoolPluginMock{
//			GetHandlerFunc: func() txpool.Handler {
//				panic("mock out the GetHandler method")
//			},
//			PrepareFunc: func(header *types.Header)  {
//				panic("mock out the Prepare method")
//			},
//		}
//
//		// use mockedTxPoolPlugin in code that requires core.TxPoolPlugin
//		// and then make assertions.
//
//	}
type TxPoolPluginMock struct {
	// GetHandlerFunc mocks the GetHandler method.
	GetHandlerFunc func() txpool.Handler

	// PrepareFunc mocks the Prepare method.
	PrepareFunc func(header *types.Header)

	// calls tracks calls to the methods.
	calls struct {
		// GetHandler holds details about calls to the GetHandler method.
		GetHandler []struct {
		}
		// Prepare holds details about calls to the Prepare method.
		Prepare []struct {
			// Header is the header argument value.
			Header *types.Header
		}
	}
	lockGetHandler sync.RWMutex
	lockPrepare       sync.RWMutex
}

// GetHandler calls GetHandlerFunc.
func (mock *TxPoolPluginMock) GetHandler() txpool.Handler {
	if mock.GetHandlerFunc == nil {
		panic("TxPoolPluginMock.GetHandlerFunc: method is nil but TxPoolPlugin.GetHandler was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetHandler.Lock()
	mock.calls.GetHandler = append(mock.calls.GetHandler, callInfo)
	mock.lockGetHandler.Unlock()
	return mock.GetHandlerFunc()
}

// GetHandlerCalls gets all the calls that were made to GetHandler.
// Check the length with:
//
//	len(mockedTxPoolPlugin.GetHandlerCalls())
func (mock *TxPoolPluginMock) GetHandlerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetHandler.RLock()
	calls = mock.calls.GetHandler
	mock.lockGetHandler.RUnlock()
	return calls
}

// Prepare calls PrepareFunc.
func (mock *TxPoolPluginMock) Prepare(header *types.Header) {
	if mock.PrepareFunc == nil {
		panic("TxPoolPluginMock.PrepareFunc: method is nil but TxPoolPlugin.Prepare was just called")
	}
	callInfo := struct {
		Header *types.Header
	}{
		Header: header,
	}
	mock.lockPrepare.Lock()
	mock.calls.Prepare = append(mock.calls.Prepare, callInfo)
	mock.lockPrepare.Unlock()
	mock.PrepareFunc(header)
}

// PrepareCalls gets all the calls that were made to Prepare.
// Check the length with:
//
//	len(mockedTxPoolPlugin.PrepareCalls())
func (mock *TxPoolPluginMock) PrepareCalls() []struct {
	Header *types.Header
} {
	var calls []struct {
		Header *types.Header
	}
	mock.lockPrepare.RLock()
	calls = mock.calls.Prepare
	mock.lockPrepare.RUnlock()
	return calls
}
