package MyWeb

type RouterGroup struct {
	prefix string  // 前缀
	router *Router // 子路由
}
