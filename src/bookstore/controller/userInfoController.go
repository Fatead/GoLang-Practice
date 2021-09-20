package controller

import (
	"awesomeProject/src/bookstore/dao"
	"awesomeProject/src/bookstore/model"
	"awesomeProject/src/bookstore/utils"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	success, uid, err := dao.Login(username, password)
	if err != nil {
		fmt.Println("登录过程中出现错误")
	}
	if success {
		//用户登录成功，为该用户创建session
		uuid := utils.CreateUUID()
		session := &model.Session{
			SessionId: uuid,
			Username:  username,
			UserId:    uid,
		}
		err := dao.AddSession(session)
		if err != nil {
			fmt.Println("添加session失败：", err)
		}
		//创建一个cookie，并将cookie写入到浏览器中
		cookie := &http.Cookie{
			HttpOnly: true,
			Name:     "user",
			Value:    uuid, //没有设置cookie过期时间、最长时间。则该cookie为会话cookie，关闭浏览器后失效
		}
		http.SetCookie(w, cookie)
		t := template.Must(template.ParseFiles("src/bookstore/views/pages/user/login_success.html"))
		err = t.Execute(w, username)
		if err != nil {
			fmt.Println("解析模板过程中出现异常", err)
		}
	} else {
		//登录失败处理逻辑，跳转到登录失败页面并给出相应的提示信息
		t := template.Must(template.ParseFiles("src/bookstore/views/pages/user/logining.html"))
		err := t.Execute(w, "登录失败，请检查输入的用户名和密码。")
		if err != nil {
			fmt.Fprintln(w, "解析模板出现异常 ，err:", err)
		}
	}
}

//注册账号
func Register(w http.ResponseWriter, r *http.Request) {
	//从HTTP报文体中解析得到用户的注册信息
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user := &model.UserInfo{
		Username: username,
		Password: password,
		Email:    email,
	}
	//先检查数据库中该用户名是否已经存在
	row, err := dao.FindUserByName(username)
	//该用户名还未被使用过
	if row == nil && err != nil {
		err := dao.AddUser(user)
		if err != nil {
			fmt.Println("注解处理过程中出现错误", err)
		}
		//解析页面模板
		t := template.Must(template.ParseFiles("src/bookstore/views/pages/user/logining.html"))
		errExe := t.Execute(w, "")
		if errExe != nil {
			fmt.Fprintln(w, "解析模板出现异常 ，err:", errExe)
		}
	}
}

//FindUserByName 通过AJAX验证用户名是否重复
func FindUserByName(w http.ResponseWriter, r *http.Request) {
	row, err := dao.FindUserByName(r.PostFormValue("username"))
	if err == nil && row != nil {
		w.Write([]byte("用户名已存在！请重新输入。"))
	} else {
		w.Write([]byte("<font style='color:blue'>用户名可用。</font>"))
	}
}

//Logout 用户登出系统，销毁cookie信息
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie == nil {
		fmt.Println("获取Cookie失败，Cookie可能不存在，用户还未登录")
	} else {
		//根据session_id从数据库中把session数据删除
		session, errFindSession := dao.DeleteSessionById(cookie.Value)
		if session == 0 && errFindSession != nil {
			fmt.Println("数据库中没有找到该session相关的记录信息,err", errFindSession)
		} else {
			//设置http报文中的cookie MaxAge属性为-1，即立即销毁
			cookie.MaxAge = -1
			http.SetCookie(w, cookie)
			fmt.Println("删除相关登录的session成功，重定向到主页")
			http.Redirect(w, r, "/main", 302)
		}
	}
}

//ToLogin 登录前预处理，先检查该用户是否已经登录过系统
func ToLogin(w http.ResponseWriter, r *http.Request) {
	// 获取用户的cookie，检查是否已经登录过
	isLogin, _, _ := dao.IsLogin(r)
	if isLogin {
		//该用户已经登录过，只需要重定向页面到主页即可
		http.Redirect(w, r, "/main", 302)
	} else {
		t := template.Must(template.ParseFiles("src/bookstore/views/pages/user/logining.html"))
		t.Execute(w, "")
	}
}

//IndexHandler 查找所有的图书
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageNo := r.PostFormValue("PageNo")
	if pageNo == "" {
		//如果该属性为空，则定位到首页
		pageNo = "1"
	}
	//将string转换为int
	indexPage, _ := strconv.ParseInt(pageNo, 10, 64)
	//根据当前页数获取该页数对应的图书信息
	pages, err := dao.GetPageBooks(indexPage)
	if err != nil {
		fmt.Println("分页查询图书出现异常", err)
	}
	isLogin, session, err := dao.IsLogin(r)
	if !isLogin || session == nil {
		fmt.Println("数据库中没有查找到该Session相关的内容，err", err)
		pages.IsLogin = false
	} else {
		pages.IsLogin = true
		pages.UserName = session.Username
	}
	t := template.Must(template.ParseFiles("src/bookstore/views/index.html"))
	t.Execute(w, pages)
}
