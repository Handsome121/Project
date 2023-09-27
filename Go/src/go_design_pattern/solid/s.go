/*
SRP
单一职责原则：对象应该仅具有一种单一功能
只能有一个引起它变化的原因
*/

package solid

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

func (u *User) Save() error {
	// Save user to database
	return nil
}

func (u *User) Notify() error {
	// Send notification email to user
	return nil
}

//在上面的例子中，User 结构体有两个职责：将用户保存到数据库并发送通知电子邮件。这违反了 SRP 原则。为了遵循 SRP 原则，我们可以重构代码：

type UserRepository interface {
	Save(user *User) error
}

type NotificationService interface {
	Notify(user *User) error
}
