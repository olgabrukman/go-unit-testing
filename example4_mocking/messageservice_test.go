package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/mock"
)

type smsServiceMock struct {
	mock.Mock
}

// Our mocked smsService method
func (m smsServiceMock) SendChargeNotification(value int) error {
	log.Println("Mocked charge notification function")
	log.Printf("Value passed in: %d\n", value)
	// this records that the method was called and passes in the value
	// it was called with
	m.Called(value)
	// it then returns whatever we tell it to return
	// in this case error=nil to simulate an SMS Service Notification
	// sent out
	return nil
}

// TestChargeCustomer is where the magic happens
// here we create our SMSService mock
func TestChargeCustomer(t *testing.T) {
	var smsService smsServiceMock

	// we then define what should be returned from SendChargeNotification
	// when we pass in the value 100 to it. In this case, we want to return
	// true as it was successful in sending a notification
	smsService.On("SendChargeNotification", 100).Return(true)

	// next we want to define the service we wish to test
	myService := MyService{smsService}
	// and call said method
	myService.ChargeCustomer(20)

	// at the end, we verify that our myService.ChargeCustomer
	// method called our mocked SendChargeNotification method
	// AssertExpectations asserts that everything specified with On and Return was
	// in fact called as expected.  Calls may have occurred in any order.
	smsService.AssertExpectations(t)
}
