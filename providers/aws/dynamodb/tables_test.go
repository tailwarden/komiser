package dynamodb

import ("testing")

func TestInt64PtrToFloat64_ValidInput(t *testing.T) {
    var number int64 = 1
    pointer := &number

    returnValue := Int64PtrToFloat64(pointer)
    var expected float64 = 1.0
    if returnValue != expected {
        t.Errorf("Expected return value: %f, but got: %f", expected, returnValue)
    }
}

func TestInt64PtrToFloat64_NilInput(t *testing.T) {
    // nil input
    returnValue := Int64PtrToFloat64(nil)
    var expected float64 = 0.0
    if returnValue != expected {
        t.Errorf("Expected return value: %f, but got: %f", expected, returnValue)
    }
}
