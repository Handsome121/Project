// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version control.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation

import (
	"regexp"
	"testing"
	"time"
)

func TestRequired(t *testing.T) {
	valid := Validation{}

	if valid.Required(nil, "nil").Ok {
		t.Error("nil object should be false")
	}
	if !valid.Required(true, "bool").Ok {
		t.Error("Bool value should always return true")
	}
	if !valid.Required(false, "bool").Ok {
		t.Error("Bool value should always return true")
	}
	if valid.Required("", "string").Ok {
		t.Error("\"'\" string should be false")
	}
	if valid.Required(" ", "string").Ok {
		t.Error("\" \" string should be false") // For #2361
	}
	if valid.Required("\n", "string").Ok {
		t.Error("new line string should be false") // For #2361
	}
	if !valid.Required("astaxie", "string").Ok {
		t.Error("string should be true")
	}
	if valid.Required(0, "zero").Ok {
		t.Error("Integer should not be equal 0")
	}
	if !valid.Required(1, "int").Ok {
		t.Error("Integer except 0 should be true")
	}
	if !valid.Required(time.Now(), "time").Ok {
		t.Error("time should be true")
	}
	if valid.Required([]string{}, "emptySlice").Ok {
		t.Error("empty slice should be false")
	}
	if !valid.Required([]interface{}{"ok"}, "slice").Ok {
		t.Error("slice should be true")
	}
}

func TestMin(t *testing.T) {
	valid := Validation{}

	if valid.Min(-1, 0, "min0").Ok {
		t.Error("-1 is less than the minimum value of 0 should be false")
	}
	if !valid.Min(1, 0, "min0").Ok {
		t.Error("1 is greater or equal than the minimum value of 0 should be true")
	}
}

func TestMax(t *testing.T) {
	valid := Validation{}

	if valid.Max(1, 0, "max0").Ok {
		t.Error("1 is greater than the minimum value of 0 should be false")
	}
	if !valid.Max(-1, 0, "max0").Ok {
		t.Error("-1 is less or equal than the maximum value of 0 should be true")
	}
}

func TestRange(t *testing.T) {
	valid := Validation{}

	if valid.Range(-1, 0, 1, "range0_1").Ok {
		t.Error("-1 is between 0 and 1 should be false")
	}
	if !valid.Range(1, 0, 1, "range0_1").Ok {
		t.Error("1 is between 0 and 1 should be true")
	}
}

func TestMinSize(t *testing.T) {
	valid := Validation{}

	if valid.MinSize("", 1, "minSize1").Ok {
		t.Error("the length of \"\" is less than the minimum value of 1 should be false")
	}
	if !valid.MinSize("ok", 1, "minSize1").Ok {
		t.Error("the length of \"ok\" is greater or equal than the minimum value of 1 should be true")
	}
	if valid.MinSize([]string{}, 1, "minSize1").Ok {
		t.Error("the length of empty slice is less than the minimum value of 1 should be false")
	}
	if !valid.MinSize([]interface{}{"ok"}, 1, "minSize1").Ok {
		t.Error("the length of [\"ok\"] is greater or equal than the minimum value of 1 should be true")
	}
}

func TestMaxSize(t *testing.T) {
	valid := Validation{}

	if valid.MaxSize("ok", 1, "maxSize1").Ok {
		t.Error("the length of \"ok\" is greater than the maximum value of 1 should be false")
	}
	if !valid.MaxSize("", 1, "maxSize1").Ok {
		t.Error("the length of \"\" is less or equal than the maximum value of 1 should be true")
	}
	if valid.MaxSize([]interface{}{"ok", false}, 1, "maxSize1").Ok {
		t.Error("the length of [\"ok\", false] is greater than the maximum value of 1 should be false")
	}
	if !valid.MaxSize([]string{}, 1, "maxSize1").Ok {
		t.Error("the length of empty slice is less or equal than the maximum value of 1 should be true")
	}
}

func TestLength(t *testing.T) {
	valid := Validation{}

	if valid.Length("", 1, "length1").Ok {
		t.Error("the length of \"\" must equal 1 should be false")
	}
	if !valid.Length("1", 1, "length1").Ok {
		t.Error("the length of \"1\" must equal 1 should be true")
	}
	if valid.Length([]string{}, 1, "length1").Ok {
		t.Error("the length of empty slice must equal 1 should be false")
	}
	if !valid.Length([]interface{}{"ok"}, 1, "length1").Ok {
		t.Error("the length of [\"ok\"] must equal 1 should be true")
	}
}

func TestAlpha(t *testing.T) {
	valid := Validation{}

	if valid.Alpha("a,1-@ $", "alpha").Ok {
		t.Error("\"a,1-@ $\" are valid alpha characters should be false")
	}
	if !valid.Alpha("abCD", "alpha").Ok {
		t.Error("\"abCD\" are valid alpha characters should be true")
	}
}

func TestNumeric(t *testing.T) {
	valid := Validation{}

	if valid.Numeric("a,1-@ $", "numeric").Ok {
		t.Error("\"a,1-@ $\" are valid numeric characters should be false")
	}
	if !valid.Numeric("1234", "numeric").Ok {
		t.Error("\"1234\" are valid numeric characters should be true")
	}
}

func TestAlphaNumeric(t *testing.T) {
	valid := Validation{}

	if valid.AlphaNumeric("a,1-@ $", "alphaNumeric").Ok {
		t.Error("\"a,1-@ $\" are valid alpha or numeric characters should be false")
	}
	if !valid.AlphaNumeric("1234aB", "alphaNumeric").Ok {
		t.Error("\"1234aB\" are valid alpha or numeric characters should be true")
	}
}

func TestMatch(t *testing.T) {
	valid := Validation{}

	if valid.Match("suchuangji@gmail", regexp.MustCompile(`^\w+@\w+\.\w+$`), "match").Ok {
		t.Error("\"suchuangji@gmail\" match \"^\\w+@\\w+\\.\\w+$\"  should be false")
	}
	if !valid.Match("suchuangji@gmail.com", regexp.MustCompile(`^\w+@\w+\.\w+$`), "match").Ok {
		t.Error("\"suchuangji@gmail\" match \"^\\w+@\\w+\\.\\w+$\"  should be true")
	}
}

func TestNoMatch(t *testing.T) {
	valid := Validation{}

	if valid.NoMatch("123@gmail", regexp.MustCompile(`[^\w\d]`), "nomatch").Ok {
		t.Error("\"123@gmail\" not match \"[^\\w\\d]\"  should be false")
	}
	if !valid.NoMatch("123gmail", regexp.MustCompile(`[^\w\d]`), "match").Ok {
		t.Error("\"123@gmail\" not match \"[^\\w\\d@]\"  should be true")
	}
}

func TestAlphaDash(t *testing.T) {
	valid := Validation{}

	if valid.AlphaDash("a,1-@ $", "alphaDash").Ok {
		t.Error("\"a,1-@ $\" are valid alpha or numeric or dash(-_) characters should be false")
	}
	if !valid.AlphaDash("1234aB-_", "alphaDash").Ok {
		t.Error("\"1234aB\" are valid alpha or numeric or dash(-_) characters should be true")
	}
}

func TestEmail(t *testing.T) {
	valid := Validation{}

	if valid.Email("not@a email", "email").Ok {
		t.Error("\"not@a email\" is a valid email address should be false")
	}
	if !valid.Email("suchuangji@gmail.com", "email").Ok {
		t.Error("\"suchuangji@gmail.com\" is a valid email address should be true")
	}
	if valid.Email("@suchuangji@gmail.com", "email").Ok {
		t.Error("\"@suchuangji@gmail.com\" is a valid email address should be false")
	}
	if valid.Email("suchuangji@gmail.com ok", "email").Ok {
		t.Error("\"suchuangji@gmail.com ok\" is a valid email address should be false")
	}
}

func TestIP(t *testing.T) {
	valid := Validation{}

	if valid.IP("11.255.255.256", "IP").Ok {
		t.Error("\"11.255.255.256\" is a valid ip address should be false")
	}
	if !valid.IP("01.11.11.11", "IP").Ok {
		t.Error("\"suchuangji@gmail.com\" is a valid ip address should be true")
	}
}

func TestBase64(t *testing.T) {
	valid := Validation{}

	if valid.Base64("suchuangji@gmail.com", "base64").Ok {
		t.Error("\"suchuangji@gmail.com\" are a valid base64 characters should be false")
	}
	if !valid.Base64("c3VjaHVhbmdqaUBnbWFpbC5jb20=", "base64").Ok {
		t.Error("\"c3VjaHVhbmdqaUBnbWFpbC5jb20=\" are a valid base64 characters should be true")
	}
}

func TestMobile(t *testing.T) {
	valid := Validation{}

	validMobiles := []string{
		"19800008888",
		"18800008888",
		"18000008888",
		"8618300008888",
		"+8614700008888",
		"17300008888",
		"+8617100008888",
		"8617500008888",
		"8617400008888",
		"16200008888",
		"16500008888",
		"16600008888",
		"16700008888",
		"13300008888",
		"14900008888",
		"15300008888",
		"17300008888",
		"17700008888",
		"18000008888",
		"18900008888",
		"19100008888",
		"19900008888",
		"19300008888",
		"13000008888",
		"13100008888",
		"13200008888",
		"14500008888",
		"15500008888",
		"15600008888",
		"16600008888",
		"17100008888",
		"17500008888",
		"17600008888",
		"18500008888",
		"18600008888",
		"13400008888",
		"13500008888",
		"13600008888",
		"13700008888",
		"13800008888",
		"13900008888",
		"14700008888",
		"15000008888",
		"15100008888",
		"15200008888",
		"15800008888",
		"15900008888",
		"17200008888",
		"17800008888",
		"18200008888",
		"18300008888",
		"18400008888",
		"18700008888",
		"18800008888",
		"19800008888",
	}

	for _, m := range validMobiles {
		if !valid.Mobile(m, "mobile").Ok {
			t.Error(m + " is a valid mobile phone number should be true")
		}
	}
}

func TestTel(t *testing.T) {
	valid := Validation{}

	if valid.Tel("222-00008888", "telephone").Ok {
		t.Error("\"222-00008888\" is a valid telephone number should be false")
	}
	if !valid.Tel("022-70008888", "telephone").Ok {
		t.Error("\"022-70008888\" is a valid telephone number should be true")
	}
	if !valid.Tel("02270008888", "telephone").Ok {
		t.Error("\"02270008888\" is a valid telephone number should be true")
	}
	if !valid.Tel("70008888", "telephone").Ok {
		t.Error("\"70008888\" is a valid telephone number should be true")
	}
}

func TestPhone(t *testing.T) {
	valid := Validation{}

	if valid.Phone("222-00008888", "phone").Ok {
		t.Error("\"222-00008888\" is a valid phone number should be false")
	}
	if !valid.Mobile("+8614700008888", "phone").Ok {
		t.Error("\"+8614700008888\" is a valid phone number should be true")
	}
	if !valid.Tel("02270008888", "phone").Ok {
		t.Error("\"02270008888\" is a valid phone number should be true")
	}
}

func TestZipCode(t *testing.T) {
	valid := Validation{}

	if valid.ZipCode("", "zipcode").Ok {
		t.Error("\"00008888\" is a valid zipcode should be false")
	}
	if !valid.ZipCode("536000", "zipcode").Ok {
		t.Error("\"536000\" is a valid zipcode should be true")
	}
}

func TestValid(t *testing.T) {
	type user struct {
		ID   int
		Name string `valid:"Required;Match(/^(test)?\\w*@(/test/);com$/)"`
		Age  int    `valid:"Required;Range(1, 140)"`
	}
	valid := Validation{}

	u := user{Name: "test@/test/;com", Age: 40}
	b, err := valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if !b {
		t.Error("validation should be passed")
	}

	uptr := &user{Name: "test", Age: 40}
	valid.Clear()
	b, err = valid.Valid(uptr)
	if err != nil {
		t.Fatal(err)
	}
	if b {
		t.Error("validation should not be passed")
	}
	if len(valid.Errors) != 1 {
		t.Fatalf("valid errors len should be 1 but got %d", len(valid.Errors))
	}
	if valid.Errors[0].Key != "Name.Match" {
		t.Errorf("Message key should be `Name.Match` but got %s", valid.Errors[0].Key)
	}

	u = user{Name: "test@/test/;com", Age: 180}
	valid.Clear()
	b, err = valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if b {
		t.Error("validation should not be passed")
	}
	if len(valid.Errors) != 1 {
		t.Fatalf("valid errors len should be 1 but got %d", len(valid.Errors))
	}
	if valid.Errors[0].Key != "Age.Range." {
		t.Errorf("Message key should be `Age.Range` but got %s", valid.Errors[0].Key)
	}
}

func TestRecursiveValid(t *testing.T) {
	type User struct {
		ID   int
		Name string `valid:"Required;Match(/^(test)?\\w*@(/test/);com$/)"`
		Age  int    `valid:"Required;Range(1, 140)"`
	}

	type AnonymouseUser struct {
		ID2   int
		Name2 string `valid:"Required;Match(/^(test)?\\w*@(/test/);com$/)"`
		Age2  int    `valid:"Required;Range(1, 140)"`
	}

	type Account struct {
		Password string `valid:"Required"`
		U        User
		AnonymouseUser
	}
	valid := Validation{}

	u := Account{Password: "abc123_", U: User{}}
	b, err := valid.RecursiveValid(u)
	if err != nil {
		t.Fatal(err)
	}
	if b {
		t.Error("validation should not be passed")
	}
}

func TestSkipValid(t *testing.T) {
	type User struct {
		ID int

		Email    string `valid:"Email"`
		ReqEmail string `valid:"Required;Email"`

		IP    string `valid:"IP"`
		ReqIP string `valid:"Required;IP"`

		Mobile    string `valid:"Mobile"`
		ReqMobile string `valid:"Required;Mobile"`

		Tel    string `valid:"Tel"`
		ReqTel string `valid:"Required;Tel"`

		Phone    string `valid:"Phone"`
		ReqPhone string `valid:"Required;Phone"`

		ZipCode    string `valid:"ZipCode"`
		ReqZipCode string `valid:"Required;ZipCode"`
	}

	u := User{
		ReqEmail:   "a@a.com",
		ReqIP:      "127.0.0.1",
		ReqMobile:  "18888888888",
		ReqTel:     "02088888888",
		ReqPhone:   "02088888888",
		ReqZipCode: "510000",
	}

	valid := Validation{}
	b, err := valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if b {
		t.Fatal("validation should not be passed")
	}

	valid = Validation{RequiredFirst: true}
	b, err = valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if !b {
		t.Fatal("validation should be passed")
	}
}

func TestPointer(t *testing.T) {
	type User struct {
		ID int

		Email    *string `valid:"Email"`
		ReqEmail *string `valid:"Required;Email"`
	}

	u := User{
		ReqEmail: nil,
		Email:    nil,
	}

	valid := Validation{}
	b, err := valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if b {
		t.Fatal("validation should not be passed")
	}

	validEmail := "a@a.com"
	u = User{
		ReqEmail: &validEmail,
		Email:    nil,
	}

	valid = Validation{RequiredFirst: true}
	b, err = valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if !b {
		t.Fatal("validation should be passed")
	}

	u = User{
		ReqEmail: &validEmail,
		Email:    nil,
	}

	valid = Validation{}
	b, err = valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if b {
		t.Fatal("validation should not be passed")
	}

	invalidEmail := "a@a"
	u = User{
		ReqEmail: &validEmail,
		Email:    &invalidEmail,
	}

	valid = Validation{RequiredFirst: true}
	b, err = valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if b {
		t.Fatal("validation should not be passed")
	}

	u = User{
		ReqEmail: &validEmail,
		Email:    &invalidEmail,
	}

	valid = Validation{}
	b, err = valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if b {
		t.Fatal("validation should not be passed")
	}
}

func TestCanSkipAlso(t *testing.T) {
	type User struct {
		ID int

		Email      string `valid:"Email"`
		ReqEmail   string `valid:"Required;Email"`
		MatchRange int    `valid:"Range(10, 20)"`
	}

	u := User{
		ReqEmail:   "a@a.com",
		Email:      "",
		MatchRange: 0,
	}

	valid := Validation{RequiredFirst: true}
	b, err := valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if b {
		t.Fatal("validation should not be passed")
	}

	valid = Validation{RequiredFirst: true}
	valid.CanSkipAlso("Range")
	b, err = valid.Valid(u)
	if err != nil {
		t.Fatal(err)
	}
	if !b {
		t.Fatal("validation should be passed")
	}

}
