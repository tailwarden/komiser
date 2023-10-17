package dynamodb

import ("testing")

func Testint64PtrToFloat64_ValidInput(t *testing.T) {
    var number int64 = 1
    pointer := &number

    returnValue := int64PtrToFloat64(pointer)
    var expected float64 = 0.0
    if returnValue != expected {
        t.Errorf("Expected return value: %f, but got: %f", expected, returnValue)
    }
}

func Testint64PtrToFloat64_InvalidInput(t *testing.T) {
    // nil input
    returnValue := int64PtrToFloat64(nil)
    var expected float64 = 3.0
    if returnValue != expected {
        t.Errorf("Expected return value: %f, but got: %f", expected, returnValue)
    }
}
