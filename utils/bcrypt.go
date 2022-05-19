package utils

import "golang.org/x/crypto/bcrypt"

// 密码加密: pwdHash  同PHP函数 password_hash()
func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// 密码验证: pwdVerify  同PHP函数 password_verify()
func CheckPasswd(pwd, hash string) bool {
	//第一个是数据库密码 第二个是用户输入密码
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))

	return err == nil
}
