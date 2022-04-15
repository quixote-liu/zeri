package request

import (
	"zeri/internal/model/common/request"
	"zeri/internal/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
