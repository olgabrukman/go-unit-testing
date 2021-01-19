package example1_basic

import "testing"

func TestHelloSimple(t *testing.T) {
	emptyResult := hello("")
	if emptyResult != GenericGreeting {
		t.Errorf("failed, expected %v, got %v", GenericGreeting, emptyResult)
	} else {
		t.Log("generic greeting case succeeded")
	}

	result := hello("Olga")
	expected := "Hello Olga!"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("hello(\"Olga\") succeeded")
	}
}

func TestHelloFailWithError(t *testing.T) {
	emptyResult := hello("")
	if emptyResult != GenericGreeting {
		t.Errorf("failed, expected %v, got [%v]", GenericGreeting, emptyResult)
	}

	result := hello("Olga")
	expected := "Hello Olga!"
	if result != expected {
		t.Errorf("failed, expected \"%v\", got \"%v\"", expected, result)
	}
}

func TestHelloFailWithFatal(t *testing.T) {
	emptyResult := hello("")
	if emptyResult != GenericGreeting {
		t.Fatalf("failed, expected \"%v\", got \"%v\"", GenericGreeting, emptyResult)
	}

	result := hello("Olga")
	expected := "Hello Olga!"
	if result != expected {
		t.Fatalf("failed, expected \"%v\", got \"%v\"", expected, result)
	}
}
