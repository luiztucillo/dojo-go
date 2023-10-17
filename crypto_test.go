package main

import "testing"

func TestEncrypt(t *testing.T) {
	value := "meli"
	expect := "phol"

	actual := Encrypt(value)

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}

func TestEcryptWithLastLetters(t *testing.T) {
	value := "xyz"
	expect := "abc"

	actual := Encrypt(value)

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}

func TestEncryptWithNumber(t *testing.T) {
	value := "meli123"
	expect := "phol123"

	actual := Encrypt(value)

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}

func TestDecrypt(t *testing.T) {
	value := "phol"
	expect := "meli"

	actual := Decrypt(value)

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}

func TestDecryptWithNumber(t *testing.T) {
	value := "phol123"
	expect := "meli123"

	actual := Decrypt(value)

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}

func TestEncryptDecrypt(t *testing.T) {
	value := "meli"
	expect := "meli"

	actual := Decrypt(Encrypt(value))

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}

func TestEncryptDecryptWithNumber(t *testing.T) {
	value := "meli123"
	expect := "meli123"

	actual := Decrypt(Encrypt(value))

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}

func TestEncryptDecryptWithSpace(t *testing.T) {
	value := "meli 123"
	expect := "meli 123"

	actual := Decrypt(Encrypt(value))

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}

func TestEncryptDecryptWithSpecialCharacter(t *testing.T) {
	value := "meli!@#$%^&*()_+"
	expect := "meli!@#$%^&*()_+"

	actual := Decrypt(Encrypt(value))

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}

func TestEncryptDecryptWithCapitalLetter(t *testing.T) {
	value := "Meli"
	expect := "Meli"

	actual := Decrypt(Encrypt(value))

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}

func TestEncryptWithUppercase(t *testing.T) {
	value := "Meli"
	expect := "Mhol"

	actual := Encrypt(value)

	if actual != expect {
		t.Errorf("Encrypt(%s): expected \"%s\", actual \"%s\"", value, expect, actual)
	}
}
