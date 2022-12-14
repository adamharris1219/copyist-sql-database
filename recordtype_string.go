// Code generated by "stringer -type=recordType"; DO NOT EDIT.

package copyist

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DriverOpen-1]
	_ = x[ConnExec-2]
	_ = x[ConnPrepare-3]
	_ = x[ConnQuery-4]
	_ = x[ConnBegin-5]
	_ = x[StmtNumInput-6]
	_ = x[StmtExec-7]
	_ = x[StmtQuery-8]
	_ = x[TxCommit-9]
	_ = x[TxRollback-10]
	_ = x[ResultLastInsertId-11]
	_ = x[ResultRowsAffected-12]
	_ = x[RowsColumns-13]
	_ = x[RowsNext-14]
}

const _recordType_name = "DriverOpenConnExecConnPrepareConnQueryConnBeginStmtNumInputStmtExecStmtQueryTxCommitTxRollbackResultLastInsertIdResultRowsAffectedRowsColumnsRowsNext"

var _recordType_index = [...]uint8{0, 10, 18, 29, 38, 47, 59, 67, 76, 84, 94, 112, 130, 141, 149}

func (i recordType) String() string {
	i -= 1
	if i < 0 || i >= recordType(len(_recordType_index)-1) {
		return "recordType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _recordType_name[_recordType_index[i]:_recordType_index[i+1]]
}
