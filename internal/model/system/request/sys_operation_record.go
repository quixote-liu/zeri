package request

import (
	"zeri/internal/model/common/request"
	"zeri/internal/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
