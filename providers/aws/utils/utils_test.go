package utils

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/aws/aws-sdk-go-v2/service/pricing"
// )

// func TestGetCost(t *testing.T) {
// 	// Single price dimension
// 	pd := PriceDimensions{
// 		BeginRange: 0,
// 		EndRange:   "Inf",
// 		PricePerUnit: struct {
// 			USD float64 `json:"USD,string"`
// 		}{USD: 0.1},
// 	}
// 	pds := []PriceDimensions{pd}
// 	cost := GetCost(pds, 10.0)
// 	expected := 1.0
// 	if cost != expected {
// 		t.Errorf("Expected cost: %f, but got: %f", expected, cost)
// 	}

// 	// Multiple price dimensions
// 	pd1 := PriceDimensions{
// 		BeginRange: 0,
// 		EndRange:   "10",
// 		PricePerUnit: struct {
// 			USD float64 `json:"USD,string"`
// 		}{USD: 0.2},
// 	}
// 	pd2 := PriceDimensions{
// 		BeginRange: 10,
// 		EndRange:   "Inf",
// 		PricePerUnit: struct {
// 			USD float64 `json:"USD,string"`
// 		}{USD: 0.1},
// 	}
// 	pds = []PriceDimensions{pd1, pd2}
// 	cost = GetCost(pds, 20)
// 	expected = 3.0
// 	if cost != expected {
// 		t.Errorf("Expected cost: %f, but got: %f", expected, cost)
// 	}
// }

// func TestGetPriceMap(t *testing.T) {
// 	testCases := []struct {
// 		inputPriceList       []string
// 		field                string
// 		expectedNumProducts  int
// 		expectedNumPriceDims map[string]int
// 	}{
// 		// Minimal valid JSON input with a single product and price dimension
// 		{
// 			inputPriceList: []string{`
// 				{
// 					"product": {
// 						"attributes": {
// 							"group": "TestGroup"
// 						}
// 					},
// 					"terms": {
// 						"OnDemand": {
// 							"test_term": {
// 								"priceDimensions": {
// 									"test_price_dimension": {
// 										"beginRange": "0",
// 										"endRange": "Inf",
// 										"pricePerUnit": {
// 											"USD": "0.1"
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}`},
// 			field:                "group",
// 			expectedNumProducts:  1,
// 			expectedNumPriceDims: map[string]int{"TestGroup": 1},
// 		},
// 		// Multiple products with different price dimensions
// 		{
// 			inputPriceList: []string{
// 				// Product 1 with 2 price dimensions
// 				`
// 				{
// 					"product": {
// 						"attributes": {
// 							"group": "TestGroup1"
// 						}
// 					},
// 					"terms": {
// 						"OnDemand": {
// 							"test_term": {
// 								"priceDimensions": {
// 									"test_price_dimension1": {
// 										"beginRange": "0",
// 										"endRange": "100",
// 										"pricePerUnit": {
// 											"USD": "0.2"
// 										}
// 									},
// 									"test_price_dimension2": {
// 										"beginRange": "100",
// 										"endRange": "Inf",
// 										"pricePerUnit": {
// 											"USD": "0.3"
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}`,
// 				// Product 2 with 3 price dimensions
// 				`
// 				{
// 					"product": {
// 						"attributes": {
// 							"group": "TestGroup2"
// 						}
// 					},
// 					"terms": {
// 						"OnDemand": {
// 							"test_term": {
// 								"priceDimensions": {
// 									"test_price_dimension1": {
// 										"beginRange": "0",
// 										"endRange": "50",
// 										"pricePerUnit": {
// 											"USD": "0.1"
// 										}
// 									},
// 									"test_price_dimension2": {
// 										"beginRange": "50",
// 										"endRange": "100",
// 										"pricePerUnit": {
// 											"USD": "0.15"
// 										}
// 									},
// 									"test_price_dimension3": {
// 										"beginRange": "100",
// 										"endRange": "Inf",
// 										"pricePerUnit": {
// 											"USD": "0.2"
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}`,
// 			},
// 			field:                "group",
// 			expectedNumProducts:  2,
// 			expectedNumPriceDims: map[string]int{"TestGroup1": 2, "TestGroup2": 3},
// 		},
// 		// Minimal valid JSON input with a single product, one price dimension & "instanceType" attribute
// 		{
// 			inputPriceList: []string{`
// 				{
// 					"product": {
// 						"attributes": {
// 							"instanceType": "TestInstanceType"
// 						}
// 					},
// 					"terms": {
// 						"OnDemand": {
// 							"test_term": {
// 								"priceDimensions": {
// 									"test_price_dimension": {
// 										"beginRange": "0",
// 										"endRange": "Inf",
// 										"pricePerUnit": {
// 											"USD": "0.1"
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}`},
// 			field:                "instanceType",
// 			expectedNumProducts:  1,
// 			expectedNumPriceDims: map[string]int{"TestInstanceType": 1},
// 		},
// 	}
// 	for i, testCase := range testCases {
// 		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
// 			output := pricing.GetProductsOutput{
// 				PriceList: testCase.inputPriceList,
// 			}
// 			priceMap, err := GetPriceMap(&output, "group")
// 			if err != nil {
// 				t.Errorf("Expected no error, but got: %v", err)
// 			}

// 			if len(priceMap) != testCase.expectedNumProducts {
// 				t.Errorf("Expected %d products in priceMap, but got %d", testCase.expectedNumProducts, len(priceMap))
// 			}

// 			for group, priceDims := range priceMap {
// 				if len(priceDims) != testCase.expectedNumPriceDims[group] {
// 					t.Errorf("Expected %d price dimensions for group %s, but got %d", testCase.expectedNumPriceDims[group], group, len(priceDims))
// 				}
// 			}
// 		})
// 	}
// }

// func TestGetPriceMap_InvalidJSON(t *testing.T) {
// 	// Invalid JSON input
// 	invalidJSON := "invalid JSON"
// 	output := pricing.GetProductsOutput{
// 		PriceList: []string{invalidJSON},
// 	}
// 	_, err := GetPriceMap(&output, "group")
// 	if err == nil {
// 		t.Error("Expected an error, but got nil")
// 	}
// }

// func TestGetPriceMap_NoPricingOutput(t *testing.T) {
// 	// PricingOutput is nil
// 	priceMap, err := GetPriceMap(nil, "group")
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}
// 	if len(priceMap) != 0 {
// 		t.Errorf("Expected an empty priceMap, but got %v", priceMap)
// 	}
// }

// func TestInt64PtrToFloat64_ValidInput(t *testing.T) {
//     var number int64 = 1
//     pointer := &number

//     returnValue := Int64PtrToFloat64(pointer)
//     var expected float64 = 1.0
//     if returnValue != expected {
//         t.Errorf("Expected return value: %f, but got: %f", expected, returnValue)
//     }
// }

// func TestInt64PtrToFloat64_NilInput(t *testing.T) {
//     // nil input
//     returnValue := Int64PtrToFloat64(nil)
//     var expected float64 = 0.0
//     if returnValue != expected {
//         t.Errorf("Expected return value: %f, but got: %f", expected, returnValue)
//     }
// }

