package consts

type StatusCode int

const (
	// 业务成功状态码（抖音方案里面固定了，表示成功）
	StatusOk StatusCode = 0
	// 业务失败状态码（0以外的都能表示失败）
	StatusError StatusCode = -1
)
