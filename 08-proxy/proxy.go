package proxy

type Subject interface {
	Do() string
}

type RealSubject struct {
}

func (r *RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real *RealSubject
}

func (p *Proxy) Do() string {
	var res string

	// 调用真实对象之前做一些前置工作，例如：权限验证之类的
	res += "pre:"

	res += p.real.Do()

	// 调用真实对象之后做一些后置工作，例如：缓存结果等
	res += ":after"

	return res
}
