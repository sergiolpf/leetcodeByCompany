package stockpricefluctuation

import (
	"fmt"
	"testing"
)

// func TestConstructor_expectedValueForMap(t *testing.T) {

// 	got := Constructor()
// 	want := StockPrice{
// 		ListOfStockPrice: map[int]*Item{},
// 		current:          0,
// 	}

// 	assert.Equal(t, want, got, "Constructor is correct")

// }

// func TestUpdate_addingOneNewTimestamp(t *testing.T) {

// 	returnedStockPricetest := Constructor()
// 	returnedStockPricetest.Update(1, 123)

// 	myTestMap := make(map[int]int, 1)
// 	myTestMap[1] = 123

// 	expectedStockPrice := StockPrice{
// 		ListOfStockPrice: myTestMap,
// 		current:          1,
// 	}

// 	assert.Equal(t, expectedStockPrice, returnedStockPricetest)
// }

func TestUpdate_addingMultiNewTimestamps(t *testing.T) {

	returnedStockPricetest := Constructor()
	returnedStockPricetest.Update(765, 3848)
	returnedStockPricetest.Update(765, 8511)
	returnedStockPricetest.Update(853, 6124)
	returnedStockPricetest.Update(765, 5733)
	maxx := returnedStockPricetest.Maximum()
	minn := returnedStockPricetest.Minimum()

	fmt.Println("maxx: ", maxx)
	fmt.Println("minn: ", minn)

	//assert.Equal(t, expectedStockPrice, returnedStockPricetest)
}

// func TestUpdate_updatingExistingTimestamp(t *testing.T) {

// 	returnedStockPricetest := Constructor()
// 	returnedStockPricetest.Update(1, 123)
// 	returnedStockPricetest.Update(2, 321)
// 	returnedStockPricetest.Update(1, 200)

// 	myTestMap := make(map[int]int)
// 	myTestMap[1] = 200
// 	myTestMap[2] = 321

// 	expectedStockPrice := StockPrice{
// 		ListOfStockPrice: myTestMap,
// 		current:          2,
// 	}

// 	assert.Equal(t, expectedStockPrice, returnedStockPricetest)
// }

// func TestCurrent_whenThereIsOnlyOneNew(t *testing.T) {
// 	returnedStockPricetest := Constructor()
// 	returnedStockPricetest.Update(1, 123)

// 	expectedCurrent := 1

// 	assert.Equal(t, expectedCurrent, returnedStockPricetest.current)
// }

// func TestCurrent_whenThereAreMultipleUpdatesOnExistingTimestamp(t *testing.T) {
// 	returnedStockPricetest := Constructor()
// 	returnedStockPricetest.Update(1, 123)
// 	returnedStockPricetest.Update(1, 321)
// 	returnedStockPricetest.Update(1, 200)

// 	expectedCurrent := 1
// 	assert.Equal(t, expectedCurrent, returnedStockPricetest.current)

// }

// func TestCurrent_whenThereMultipleNewExistingUpdates(t *testing.T) {
// 	returnedStockPricetest := Constructor()
// 	returnedStockPricetest.Update(1, 123)
// 	returnedStockPricetest.Update(2, 321)
// 	returnedStockPricetest.Update(3, 200)
// 	returnedStockPricetest.Update(4, 200)
// 	returnedStockPricetest.Update(1, 200)
// 	returnedStockPricetest.Update(3, 200)

// 	expectedCurrent := 4
// 	assert.Equal(t, expectedCurrent, returnedStockPricetest.current)

// }
