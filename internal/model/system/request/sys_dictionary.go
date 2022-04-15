package request

import (
	"zeri/internal/model/common/request"
	"zeri/internal/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
