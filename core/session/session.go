package session

type Session struct {}

var session = make(map[string]interface{})

// Get data from the session
// 从会话中获取数据
func Get(key string) interface{} {
	return session[key]
}

// Set data to the session
// 为会话设置数据
func Set(key string, value interface{}) {
	session[key] = value
	return
}

// Deletes data from the session
// 删除会话中的数据
func Delete(key string) {
	delete(session, key)
	return
}