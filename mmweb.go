package MyWeb

import "net/http"

type Mmweb struct {
	Router      *Router
	RouterGroup *RouterGroup
}

// Mmweb实例
func Instance() *Mmweb {
	return &Mmweb{Router: NewRouter(), RouterGroup: new(RouterGroup)}
}

// 启动Mmweb框架
func (m Mmweb) ListenAndServer(addr string) {
	http.ListenAndServe(addr, m.Router)
}

// 路由组功能
func (r *Router) Group(prefix string) *RouterGroup {
	return &RouterGroup{prefix: prefix, router: r}
}

func (group *RouterGroup) HandleGroup(method, path string, handler HandlerFunc) {
	fullpath := group.prefix + path
	group.router.Handle(method, fullpath, handler)
}
