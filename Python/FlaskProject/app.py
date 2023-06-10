# 导入Flask扩展
import sys
from imp import reload

import pymysql
from flask import Flask, render_template, flash, request, redirect, url_for
from flask_migrate import Migrate, MigrateCommand
from flask_script import Manager
from flask_sqlalchemy import SQLAlchemy
from flask_wtf import FlaskForm, CSRFProtect
from wtforms import StringField, SubmitField
from wtforms.validators import DataRequired

reload(sys)
pymysql.install_as_MySQLdb()

# 创建Flask应用程序实例
# 需要传入__name__,作用是为了确定资源所在的路径
app = Flask(__name__)

app.config['SQLALCHEMY_DATABASE_URI'] = 'mysql://root:123456@localhost/flask_books'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
app.secret_key = 'itheima'

db = SQLAlchemy(app)
migrate = Migrate(app, db)

manager = Manager(app)
manager.add_command('db', MigrateCommand)


class Author(db.Model):
    __tablename__ = 'authors'
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(16), unique=True)
    books = db.relationship('Book', backref='author')

    def __repr__(self):
        return 'Author:%s' % self.name


class Book(db.Model):
    __tablename__ = 'books'
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(16), unique=True)
    author_id = db.Column(db.Integer, db.ForeignKey('authors.id'))


# 自定义表单类
class AuthorForm(FlaskForm):
    author = StringField(u'作者', validators=[DataRequired()])
    book = StringField(u'书籍', validators=[DataRequired()])
    submit = SubmitField(u'提交')
    csrf = CSRFProtect()


# 定义路由和视图函数
# 默认请求方式为GET,如果需要增加，需要自行指定
@app.route('/', methods=['GET', 'POST'])
def index():
    # 创建自定义表单类
    author_form = AuthorForm()
    print(author_form.validate_on_submit())
    if author_form.validate_on_submit():
        print("=========")
        author_name = author_form.author.data
        book_name = author_form.book.data
        author = Author.query.filter_by(name=author_name).first()
        if author:
            book = Book.query.filter_by(name=book_name).first()
            if book:
                flash('已存在同名书籍')
            else:
                try:
                    new_book = Book(name=book_name, author_id=author.id)
                    db.session.add(new_book)
                    db.session.commit()
                except Exception as e:
                    print(e)
                    flash('添加书籍失败')
                    db.session.rollback()
        else:
            try:
                new_author = Author(name=author_name)
                db.session.add(new_author)
                db.session.commit()
                new_book = Book(name=book_name, author_id=new_author.id)
                db.session.add(new_book)
                db.session.commit()
            except Exception as e:
                print(e)
                flash('作者和书籍添加失败')
                db.session.rollback()
    else:
        print("++++++++++++")
        if request.method == 'POST':
            flash('参数不全')

    # 查询所有的作者信息，让信息传递给模板
    authors = Author.query.all()

    return render_template('books.html', authors=authors, form=author_form)


@app.route('/delete_book/<book_id>')
def delete_book(book_id):
    book = Book.query.get(book_id)
    if book:
        try:
            db.session.delete(book_id)
            db.session.commit()
        except Exception as e:
            print(e)
            flash('删除书籍出错')
            db.session.rollback()
    else:
        flash('书籍找不到')

    return redirect(url_for('index'))


@app.route('/delete_author/<author_id>')
def delete_author(author_id):
    author = Author.query.get(author_id)
    if author:
        try:
            Book.query.filter_by(author_id=author_id).delete()
            db.session.delete(author_id)
            db.session.commit()
        except Exception as e:
            print(e)
            flash('删除书籍出错')
            db.session.rollback()
    else:
        flash('书籍找不到')

    return redirect(url_for('index'))


if __name__ == '__main__':
    # 执行了app.run，就会将Flask程序运行在一个简易的服务器
    db.drop_all()
    db.create_all()
    au1 = Author(name='老王')
    au2 = Author(name='老惠')
    au3 = Author(name='老刘')
    db.session.add_all([au1, au2, au3])
    db.session.commit()
    bk1 = Book(name='老王回忆录', author_id=au1.id)
    bk2 = Book(name='我读书少，你别骗我', author_id=au1.id)
    bk3 = Book(name='如何才能让自己更骚', author_id=au2.id)
    bk4 = Book(name='怎样才能征服美丽少女', author_id=au3.id)
    bk5 = Book(name='如何征服英俊少男', author_id=au3.id)
    db.session.add_all([bk1, bk2, bk3, bk4, bk5])
    db.session.commit()
    app.run(debug=True)
