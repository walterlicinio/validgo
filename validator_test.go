package validgo

import "testing"

func TestValidator_AllFieldsCorrect(t *testing.T) {
	type User struct {
		Name  string
		Email string
		Age   int
	}

	user := &User{
		Name:  "John Doe",
		Email: "newmail@mail.com",
		Age:   30,
	}

	v := New(user)
	v.Field("Name").IsString().MinLength(3).MaxLength(50)
	v.Field("Email").IsString().IsEmail()
	v.Field("Age").IsInt().Min(18).Max(60)

	if v.HasErrors() {
		t.Errorf("Validation failed: %v", v.Errors())
	}
}

func TestValidator_InvalidName(t *testing.T) {
	type User struct {
		Name  string
		Email string
		Age   int
	}

	user := &User{
		Name:  "J",
		Email: "newmail@mail.com",
		Age:   30,
	}

	v := New(user)
	v.Field("Name").IsString().MinLength(3).MaxLength(50)
	v.Field("Email").IsString().IsEmail()
	v.Field("Age").IsInt().Min(18).Max(60)

	if !v.HasErrors() {
		t.Errorf("Validation should have failed")
	}
}

func TestValidator_InvalidEmail(t *testing.T) {
	type User struct {
		Name  string
		Email string
		Age   int
	}

	user := &User{
		Name:  "John Doe",
		Email: "newmail",
		Age:   30,
	}

	v := New(user)
	v.Field("Name").IsString().MinLength(3).MaxLength(50)
	v.Field("Email").IsString().IsEmail()
	v.Field("Age").IsInt().Min(18).Max(60)

	if !v.HasErrors() {
		t.Errorf("Validation should have failed")
	}
}

func TestValidator_InvalidAge(t *testing.T) {
	type User struct {
		Name  string
		Email string
		Age   int
	}

	user := &User{
		Name:  "John Doe",
		Email: "newmail@mail.com",
		Age:   10,
	}

	v := New(user)
	v.Field("Name").IsString().MinLength(3).MaxLength(50)
	v.Field("Email").IsString().IsEmail()
	v.Field("Age").IsInt().Min(18).Max(60)

	if !v.HasErrors() {
		t.Errorf("Validation should have failed")
	}
}

func TestValidator_InvalidField(t *testing.T) {
	type User struct {
		Name  string
		Email string
		Age   int
	}

	user := &User{
		Name:  "John Doe",
		Email: "newmail@mail.com",
		Age:   30,
	}

	v := New(user)
	v.Field("Name").IsString().MinLength(3).MaxLength(50)
	v.Field("Email").IsString().IsEmail()
	v.Field("Age").IsInt().Min(18).Max(60)
	v.Field("InvalidField").IsString()

	if !v.HasErrors() {
		t.Errorf("Validation should have failed")
	}
}

func TestValidator_CustomValidation(t *testing.T) {
	type User struct {
		Name  string
		Email string
		Age   int
	}

	user := &User{
		Name:  "John Doe",
		Email: "newmail@mail.com",
		Age:   33,
	}

	v := New(user)
	v.Field("Name").IsString().MinLength(3).MaxLength(50)
	v.Field("Email").IsString().IsEmail()
	v.Field("Age").IsInt().Min(18).Max(60)
	v.Field("Age").Custom(func(value any) bool {
		return value.(int)%2 == 0
	})

	if !v.HasErrors() {
		t.Errorf("Validation should have failed")
	}
}

func TestValidator_Transform(t *testing.T) {
	type User struct {
		Name  string
		Email string
		Age   int
	}

	user := &User{
		Name:  "John Doe",
		Email: "newmail@mail.com",
		Age:   33,
	}

	v := New(user)
	v.Field("Name").IsString().MinLength(3).MaxLength(50)
	v.Field("Email").IsString().IsEmail()
	v.Field("Age").IsInt().Min(18).Max(60)

	v.Field("Age").Transform(func(value any) any {
		return value.(int) * 2
	})

	if v.HasErrors() {
		t.Errorf("Validation failed: %v", v.Errors())
	}

	if user.Age != 66 {
		t.Errorf("Transformation failed")
	}
}
