/*
依赖反转原则（DIP） 依赖反转原则指高级模块不应依赖于低级模块，而是两者都应该依赖于抽象。此外，抽象不应依赖于细节，但细节应该依赖于抽象。
*/

package solid

//type Logger struct{}
//
//func (l *Logger) Log(message string) {
//	// Write log to a file
//}
//
//type UserManager struct {
//	logger *Logger
//}
//
//func NewUserManager() *UserManager {
//	return &UserManager{
//		logger: &Logger{},
//	}
//}
//
//func (u *UserManager) CreateUser(user *User) {
//	// Create user
//	u.logger.Log("User created!")
//}

//在上面的例子中，UserManager 直接依赖于 Logger 结构体。我们使用特定类型的日志记录器构造 UserManager，这将使将来更改它和进行单元测试更加
//困难。这违反了 DIP 原则。我们可以重构代码以遵循它：

type Logger interface {
	Log(message string) error
}

type FileLogger struct{}

func (f *FileLogger) Log(message string) error {
	// Write log to a file
	return nil
}

type UserManager struct {
	logger Logger
}

func NewUserManager(logger Logger) *UserManager {
	return &UserManager{
		logger: logger,
	}
}

func (u *UserManager) CreateUser(user *User) {
	// Create user
	err := u.logger.Log("User created")
	if err != nil {
		return
	}
}

/*
我们添加了一个 Logger 接口并将其注入到 UserManager 中，构造函数 NewUserManager 现在接受一个接口并用它来初始化我们的 UserManager 结构体，
通过这样做，我们可以在将来更轻松地更改日志记录器并在测试中轻松地进行存根或模拟。我们 反转了 依赖关系，使代码更加灵活和易于维护。
*/
