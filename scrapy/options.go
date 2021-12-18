package scrapy

type Option interface {
	apply(n *NewCoder)
}

type function func(n *NewCoder)

func (f function) apply(n *NewCoder) {
	f(n)
}

func WithQuery(query string) Option {
	return function(func(n *NewCoder) {

		n.query = query
	})
}
