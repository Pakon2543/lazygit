// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/jesseduffield/lazygit/pkg/commands"
	"github.com/jesseduffield/lazygit/pkg/commands/types"
)

type FakeIBranchesMgr struct {
	AllBranchesCmdObjStub        func() types.ICmdObj
	allBranchesCmdObjMutex       sync.RWMutex
	allBranchesCmdObjArgsForCall []struct {
	}
	allBranchesCmdObjReturns struct {
		result1 types.ICmdObj
	}
	allBranchesCmdObjReturnsOnCall map[int]struct {
		result1 types.ICmdObj
	}
	CheckoutStub        func(string, commands.CheckoutOpts) error
	checkoutMutex       sync.RWMutex
	checkoutArgsForCall []struct {
		arg1 string
		arg2 commands.CheckoutOpts
	}
	checkoutReturns struct {
		result1 error
	}
	checkoutReturnsOnCall map[int]struct {
		result1 error
	}
	CurrentBranchNameStub        func() (string, string, error)
	currentBranchNameMutex       sync.RWMutex
	currentBranchNameArgsForCall []struct {
	}
	currentBranchNameReturns struct {
		result1 string
		result2 string
		result3 error
	}
	currentBranchNameReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	DeleteStub        func(string, bool) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	GetBranchGraphCmdObjStub        func(string) types.ICmdObj
	getBranchGraphCmdObjMutex       sync.RWMutex
	getBranchGraphCmdObjArgsForCall []struct {
		arg1 string
	}
	getBranchGraphCmdObjReturns struct {
		result1 types.ICmdObj
	}
	getBranchGraphCmdObjReturnsOnCall map[int]struct {
		result1 types.ICmdObj
	}
	GetUpstreamStub        func(string) (string, error)
	getUpstreamMutex       sync.RWMutex
	getUpstreamArgsForCall []struct {
		arg1 string
	}
	getUpstreamReturns struct {
		result1 string
		result2 error
	}
	getUpstreamReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	MergeStub        func(string, commands.MergeOpts) error
	mergeMutex       sync.RWMutex
	mergeArgsForCall []struct {
		arg1 string
		arg2 commands.MergeOpts
	}
	mergeReturns struct {
		result1 error
	}
	mergeReturnsOnCall map[int]struct {
		result1 error
	}
	NewBranchStub        func(string, string) error
	newBranchMutex       sync.RWMutex
	newBranchArgsForCall []struct {
		arg1 string
		arg2 string
	}
	newBranchReturns struct {
		result1 error
	}
	newBranchReturnsOnCall map[int]struct {
		result1 error
	}
	SetUpstreamStub        func(string, string) error
	setUpstreamMutex       sync.RWMutex
	setUpstreamArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setUpstreamReturns struct {
		result1 error
	}
	setUpstreamReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIBranchesMgr) AllBranchesCmdObj() types.ICmdObj {
	fake.allBranchesCmdObjMutex.Lock()
	ret, specificReturn := fake.allBranchesCmdObjReturnsOnCall[len(fake.allBranchesCmdObjArgsForCall)]
	fake.allBranchesCmdObjArgsForCall = append(fake.allBranchesCmdObjArgsForCall, struct {
	}{})
	stub := fake.AllBranchesCmdObjStub
	fakeReturns := fake.allBranchesCmdObjReturns
	fake.recordInvocation("AllBranchesCmdObj", []interface{}{})
	fake.allBranchesCmdObjMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBranchesMgr) AllBranchesCmdObjCallCount() int {
	fake.allBranchesCmdObjMutex.RLock()
	defer fake.allBranchesCmdObjMutex.RUnlock()
	return len(fake.allBranchesCmdObjArgsForCall)
}

func (fake *FakeIBranchesMgr) AllBranchesCmdObjCalls(stub func() types.ICmdObj) {
	fake.allBranchesCmdObjMutex.Lock()
	defer fake.allBranchesCmdObjMutex.Unlock()
	fake.AllBranchesCmdObjStub = stub
}

func (fake *FakeIBranchesMgr) AllBranchesCmdObjReturns(result1 types.ICmdObj) {
	fake.allBranchesCmdObjMutex.Lock()
	defer fake.allBranchesCmdObjMutex.Unlock()
	fake.AllBranchesCmdObjStub = nil
	fake.allBranchesCmdObjReturns = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeIBranchesMgr) AllBranchesCmdObjReturnsOnCall(i int, result1 types.ICmdObj) {
	fake.allBranchesCmdObjMutex.Lock()
	defer fake.allBranchesCmdObjMutex.Unlock()
	fake.AllBranchesCmdObjStub = nil
	if fake.allBranchesCmdObjReturnsOnCall == nil {
		fake.allBranchesCmdObjReturnsOnCall = make(map[int]struct {
			result1 types.ICmdObj
		})
	}
	fake.allBranchesCmdObjReturnsOnCall[i] = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeIBranchesMgr) Checkout(arg1 string, arg2 commands.CheckoutOpts) error {
	fake.checkoutMutex.Lock()
	ret, specificReturn := fake.checkoutReturnsOnCall[len(fake.checkoutArgsForCall)]
	fake.checkoutArgsForCall = append(fake.checkoutArgsForCall, struct {
		arg1 string
		arg2 commands.CheckoutOpts
	}{arg1, arg2})
	stub := fake.CheckoutStub
	fakeReturns := fake.checkoutReturns
	fake.recordInvocation("Checkout", []interface{}{arg1, arg2})
	fake.checkoutMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBranchesMgr) CheckoutCallCount() int {
	fake.checkoutMutex.RLock()
	defer fake.checkoutMutex.RUnlock()
	return len(fake.checkoutArgsForCall)
}

func (fake *FakeIBranchesMgr) CheckoutCalls(stub func(string, commands.CheckoutOpts) error) {
	fake.checkoutMutex.Lock()
	defer fake.checkoutMutex.Unlock()
	fake.CheckoutStub = stub
}

func (fake *FakeIBranchesMgr) CheckoutArgsForCall(i int) (string, commands.CheckoutOpts) {
	fake.checkoutMutex.RLock()
	defer fake.checkoutMutex.RUnlock()
	argsForCall := fake.checkoutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBranchesMgr) CheckoutReturns(result1 error) {
	fake.checkoutMutex.Lock()
	defer fake.checkoutMutex.Unlock()
	fake.CheckoutStub = nil
	fake.checkoutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBranchesMgr) CheckoutReturnsOnCall(i int, result1 error) {
	fake.checkoutMutex.Lock()
	defer fake.checkoutMutex.Unlock()
	fake.CheckoutStub = nil
	if fake.checkoutReturnsOnCall == nil {
		fake.checkoutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkoutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBranchesMgr) CurrentBranchName() (string, string, error) {
	fake.currentBranchNameMutex.Lock()
	ret, specificReturn := fake.currentBranchNameReturnsOnCall[len(fake.currentBranchNameArgsForCall)]
	fake.currentBranchNameArgsForCall = append(fake.currentBranchNameArgsForCall, struct {
	}{})
	stub := fake.CurrentBranchNameStub
	fakeReturns := fake.currentBranchNameReturns
	fake.recordInvocation("CurrentBranchName", []interface{}{})
	fake.currentBranchNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeIBranchesMgr) CurrentBranchNameCallCount() int {
	fake.currentBranchNameMutex.RLock()
	defer fake.currentBranchNameMutex.RUnlock()
	return len(fake.currentBranchNameArgsForCall)
}

func (fake *FakeIBranchesMgr) CurrentBranchNameCalls(stub func() (string, string, error)) {
	fake.currentBranchNameMutex.Lock()
	defer fake.currentBranchNameMutex.Unlock()
	fake.CurrentBranchNameStub = stub
}

func (fake *FakeIBranchesMgr) CurrentBranchNameReturns(result1 string, result2 string, result3 error) {
	fake.currentBranchNameMutex.Lock()
	defer fake.currentBranchNameMutex.Unlock()
	fake.CurrentBranchNameStub = nil
	fake.currentBranchNameReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIBranchesMgr) CurrentBranchNameReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.currentBranchNameMutex.Lock()
	defer fake.currentBranchNameMutex.Unlock()
	fake.CurrentBranchNameStub = nil
	if fake.currentBranchNameReturnsOnCall == nil {
		fake.currentBranchNameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.currentBranchNameReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIBranchesMgr) Delete(arg1 string, arg2 bool) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBranchesMgr) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeIBranchesMgr) DeleteCalls(stub func(string, bool) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeIBranchesMgr) DeleteArgsForCall(i int) (string, bool) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBranchesMgr) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBranchesMgr) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBranchesMgr) GetBranchGraphCmdObj(arg1 string) types.ICmdObj {
	fake.getBranchGraphCmdObjMutex.Lock()
	ret, specificReturn := fake.getBranchGraphCmdObjReturnsOnCall[len(fake.getBranchGraphCmdObjArgsForCall)]
	fake.getBranchGraphCmdObjArgsForCall = append(fake.getBranchGraphCmdObjArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetBranchGraphCmdObjStub
	fakeReturns := fake.getBranchGraphCmdObjReturns
	fake.recordInvocation("GetBranchGraphCmdObj", []interface{}{arg1})
	fake.getBranchGraphCmdObjMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBranchesMgr) GetBranchGraphCmdObjCallCount() int {
	fake.getBranchGraphCmdObjMutex.RLock()
	defer fake.getBranchGraphCmdObjMutex.RUnlock()
	return len(fake.getBranchGraphCmdObjArgsForCall)
}

func (fake *FakeIBranchesMgr) GetBranchGraphCmdObjCalls(stub func(string) types.ICmdObj) {
	fake.getBranchGraphCmdObjMutex.Lock()
	defer fake.getBranchGraphCmdObjMutex.Unlock()
	fake.GetBranchGraphCmdObjStub = stub
}

func (fake *FakeIBranchesMgr) GetBranchGraphCmdObjArgsForCall(i int) string {
	fake.getBranchGraphCmdObjMutex.RLock()
	defer fake.getBranchGraphCmdObjMutex.RUnlock()
	argsForCall := fake.getBranchGraphCmdObjArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIBranchesMgr) GetBranchGraphCmdObjReturns(result1 types.ICmdObj) {
	fake.getBranchGraphCmdObjMutex.Lock()
	defer fake.getBranchGraphCmdObjMutex.Unlock()
	fake.GetBranchGraphCmdObjStub = nil
	fake.getBranchGraphCmdObjReturns = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeIBranchesMgr) GetBranchGraphCmdObjReturnsOnCall(i int, result1 types.ICmdObj) {
	fake.getBranchGraphCmdObjMutex.Lock()
	defer fake.getBranchGraphCmdObjMutex.Unlock()
	fake.GetBranchGraphCmdObjStub = nil
	if fake.getBranchGraphCmdObjReturnsOnCall == nil {
		fake.getBranchGraphCmdObjReturnsOnCall = make(map[int]struct {
			result1 types.ICmdObj
		})
	}
	fake.getBranchGraphCmdObjReturnsOnCall[i] = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeIBranchesMgr) GetUpstream(arg1 string) (string, error) {
	fake.getUpstreamMutex.Lock()
	ret, specificReturn := fake.getUpstreamReturnsOnCall[len(fake.getUpstreamArgsForCall)]
	fake.getUpstreamArgsForCall = append(fake.getUpstreamArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetUpstreamStub
	fakeReturns := fake.getUpstreamReturns
	fake.recordInvocation("GetUpstream", []interface{}{arg1})
	fake.getUpstreamMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIBranchesMgr) GetUpstreamCallCount() int {
	fake.getUpstreamMutex.RLock()
	defer fake.getUpstreamMutex.RUnlock()
	return len(fake.getUpstreamArgsForCall)
}

func (fake *FakeIBranchesMgr) GetUpstreamCalls(stub func(string) (string, error)) {
	fake.getUpstreamMutex.Lock()
	defer fake.getUpstreamMutex.Unlock()
	fake.GetUpstreamStub = stub
}

func (fake *FakeIBranchesMgr) GetUpstreamArgsForCall(i int) string {
	fake.getUpstreamMutex.RLock()
	defer fake.getUpstreamMutex.RUnlock()
	argsForCall := fake.getUpstreamArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIBranchesMgr) GetUpstreamReturns(result1 string, result2 error) {
	fake.getUpstreamMutex.Lock()
	defer fake.getUpstreamMutex.Unlock()
	fake.GetUpstreamStub = nil
	fake.getUpstreamReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeIBranchesMgr) GetUpstreamReturnsOnCall(i int, result1 string, result2 error) {
	fake.getUpstreamMutex.Lock()
	defer fake.getUpstreamMutex.Unlock()
	fake.GetUpstreamStub = nil
	if fake.getUpstreamReturnsOnCall == nil {
		fake.getUpstreamReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getUpstreamReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeIBranchesMgr) Merge(arg1 string, arg2 commands.MergeOpts) error {
	fake.mergeMutex.Lock()
	ret, specificReturn := fake.mergeReturnsOnCall[len(fake.mergeArgsForCall)]
	fake.mergeArgsForCall = append(fake.mergeArgsForCall, struct {
		arg1 string
		arg2 commands.MergeOpts
	}{arg1, arg2})
	stub := fake.MergeStub
	fakeReturns := fake.mergeReturns
	fake.recordInvocation("Merge", []interface{}{arg1, arg2})
	fake.mergeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBranchesMgr) MergeCallCount() int {
	fake.mergeMutex.RLock()
	defer fake.mergeMutex.RUnlock()
	return len(fake.mergeArgsForCall)
}

func (fake *FakeIBranchesMgr) MergeCalls(stub func(string, commands.MergeOpts) error) {
	fake.mergeMutex.Lock()
	defer fake.mergeMutex.Unlock()
	fake.MergeStub = stub
}

func (fake *FakeIBranchesMgr) MergeArgsForCall(i int) (string, commands.MergeOpts) {
	fake.mergeMutex.RLock()
	defer fake.mergeMutex.RUnlock()
	argsForCall := fake.mergeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBranchesMgr) MergeReturns(result1 error) {
	fake.mergeMutex.Lock()
	defer fake.mergeMutex.Unlock()
	fake.MergeStub = nil
	fake.mergeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBranchesMgr) MergeReturnsOnCall(i int, result1 error) {
	fake.mergeMutex.Lock()
	defer fake.mergeMutex.Unlock()
	fake.MergeStub = nil
	if fake.mergeReturnsOnCall == nil {
		fake.mergeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mergeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBranchesMgr) NewBranch(arg1 string, arg2 string) error {
	fake.newBranchMutex.Lock()
	ret, specificReturn := fake.newBranchReturnsOnCall[len(fake.newBranchArgsForCall)]
	fake.newBranchArgsForCall = append(fake.newBranchArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.NewBranchStub
	fakeReturns := fake.newBranchReturns
	fake.recordInvocation("NewBranch", []interface{}{arg1, arg2})
	fake.newBranchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBranchesMgr) NewBranchCallCount() int {
	fake.newBranchMutex.RLock()
	defer fake.newBranchMutex.RUnlock()
	return len(fake.newBranchArgsForCall)
}

func (fake *FakeIBranchesMgr) NewBranchCalls(stub func(string, string) error) {
	fake.newBranchMutex.Lock()
	defer fake.newBranchMutex.Unlock()
	fake.NewBranchStub = stub
}

func (fake *FakeIBranchesMgr) NewBranchArgsForCall(i int) (string, string) {
	fake.newBranchMutex.RLock()
	defer fake.newBranchMutex.RUnlock()
	argsForCall := fake.newBranchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBranchesMgr) NewBranchReturns(result1 error) {
	fake.newBranchMutex.Lock()
	defer fake.newBranchMutex.Unlock()
	fake.NewBranchStub = nil
	fake.newBranchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBranchesMgr) NewBranchReturnsOnCall(i int, result1 error) {
	fake.newBranchMutex.Lock()
	defer fake.newBranchMutex.Unlock()
	fake.NewBranchStub = nil
	if fake.newBranchReturnsOnCall == nil {
		fake.newBranchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.newBranchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBranchesMgr) SetUpstream(arg1 string, arg2 string) error {
	fake.setUpstreamMutex.Lock()
	ret, specificReturn := fake.setUpstreamReturnsOnCall[len(fake.setUpstreamArgsForCall)]
	fake.setUpstreamArgsForCall = append(fake.setUpstreamArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.SetUpstreamStub
	fakeReturns := fake.setUpstreamReturns
	fake.recordInvocation("SetUpstream", []interface{}{arg1, arg2})
	fake.setUpstreamMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBranchesMgr) SetUpstreamCallCount() int {
	fake.setUpstreamMutex.RLock()
	defer fake.setUpstreamMutex.RUnlock()
	return len(fake.setUpstreamArgsForCall)
}

func (fake *FakeIBranchesMgr) SetUpstreamCalls(stub func(string, string) error) {
	fake.setUpstreamMutex.Lock()
	defer fake.setUpstreamMutex.Unlock()
	fake.SetUpstreamStub = stub
}

func (fake *FakeIBranchesMgr) SetUpstreamArgsForCall(i int) (string, string) {
	fake.setUpstreamMutex.RLock()
	defer fake.setUpstreamMutex.RUnlock()
	argsForCall := fake.setUpstreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBranchesMgr) SetUpstreamReturns(result1 error) {
	fake.setUpstreamMutex.Lock()
	defer fake.setUpstreamMutex.Unlock()
	fake.SetUpstreamStub = nil
	fake.setUpstreamReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBranchesMgr) SetUpstreamReturnsOnCall(i int, result1 error) {
	fake.setUpstreamMutex.Lock()
	defer fake.setUpstreamMutex.Unlock()
	fake.SetUpstreamStub = nil
	if fake.setUpstreamReturnsOnCall == nil {
		fake.setUpstreamReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setUpstreamReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBranchesMgr) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allBranchesCmdObjMutex.RLock()
	defer fake.allBranchesCmdObjMutex.RUnlock()
	fake.checkoutMutex.RLock()
	defer fake.checkoutMutex.RUnlock()
	fake.currentBranchNameMutex.RLock()
	defer fake.currentBranchNameMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getBranchGraphCmdObjMutex.RLock()
	defer fake.getBranchGraphCmdObjMutex.RUnlock()
	fake.getUpstreamMutex.RLock()
	defer fake.getUpstreamMutex.RUnlock()
	fake.mergeMutex.RLock()
	defer fake.mergeMutex.RUnlock()
	fake.newBranchMutex.RLock()
	defer fake.newBranchMutex.RUnlock()
	fake.setUpstreamMutex.RLock()
	defer fake.setUpstreamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIBranchesMgr) recordInvocation(key string, args []interface{}) {
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

var _ commands.IBranchesMgr = new(FakeIBranchesMgr)