package library

const (
	// 公共错误码 https://wiki.medlinker.com/pages/viewpage.action?pageId=17598142
	Success = 0

	//  500
	ErrInternal = 500
)

func EqualByCode(err error, code int) bool {
	return true
}
