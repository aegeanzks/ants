package gnet

//全局方法和常量

const (
	//handle
	EVENT_CONN_READ      = 1
	EVENT_CONN_SEND      = 2
	EVENT_HEARTBEAT_PINT = 3 //心跳
	//close sign
	SIGN_SEND_ERROR       = -2
	SIGN_CLOSE_NULL       = -1
	SIGN_CLOSE_OK         = 0
	SIGN_CLOSE_DISTAL     = 1
	SIGN_CLOSE_ERROR      = 2
	SIGN_CLOSE_ERROR_READ = 3
	SIGN_CLOSE_ERROR_SEND = 4
	SIGN_CLOSE_HEARTBEAT  = 5
	//chan sign
	NET_CHAN_SIZE = 1000
	//服务器最大socket链接
	NET_SERVER_CONN_SIZE = 10000
	//pack
	HEAD_SIZE       = 2
	NET_BUFF_MINLEN = 1
	NET_BUFF_MAXLEN = 1024 * 10 //(10MB)
	//心跳
	PING_TIME = 1000 * 60 * 5 //心跳时间，秒(1000 * 60 * 5)
)

//转化字节
func ToBytes(data interface{}) []byte {
	switch data.(type) {
	case IBytes:
		return data.(IBytes).Bytes()
	case []byte:
		return data.([]byte)
	case string:
		return []byte(data.(string))
	}
	return nil
}

func check_size(size int) bool {
	return size >= NET_BUFF_MINLEN && size <= NET_BUFF_MAXLEN
}
