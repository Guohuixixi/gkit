package option

// Option 基于Option模式封装的函数类型
type Option[T any] func(t *T)

// Apply 应用Option模式
func Apply[T any](t *T, opts ...Option[T]) {
	for _, opt := range opts {
		opt(t)
	}
}

// OptionError 基于Option模式封装的函数类型,返回值为Error
type OptionError[T any] func(t *T) error

// ApplyError 应用Option模式,返回值为Error,用于处理一些会产生错误的函数
func ApplyError[T any](t *T, opts ...OptionError[T]) error {
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return err
		}
	}
	return nil
}
