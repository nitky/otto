// Code generated by "stringer -type=ExecuteTask execute.go"; DO NOT EDIT

package otto

import "fmt"

const _ExecuteTask_name = "ExecuteTaskInvalid"

var _ExecuteTask_index = [...]uint8{0, 18}

func (i ExecuteTask) String() string {
	if i >= ExecuteTask(len(_ExecuteTask_index)-1) {
		return fmt.Sprintf("ExecuteTask(%d)", i)
	}
	return _ExecuteTask_name[_ExecuteTask_index[i]:_ExecuteTask_index[i+1]]
}
