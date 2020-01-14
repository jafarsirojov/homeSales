package main

import "fmt"

var homesSale = []home{
	{
		id:                  1,
		name:                "1 комнатная квартира, 99 кв.м. 9/25 эт.",
		price:               600_000,
		area:                "Душанбе",
		addressHouse:        "р.Сино, Борбад 33",
		distanceOfTheCenter: 10,
		postTime:            1_578_232_244,
		seller:              "Вася",
	},
	{
		id:                  2,
		name:                "2 комнатная квартира, 110 кв.м. 3/25 эт.",
		price:               650_000,
		area:                "Душанбе",
		addressHouse:        "р.Фирдовси, Дусти 15",
		distanceOfTheCenter: 13,
		postTime:            1_578_233_112,
		seller:              "Петя",
	},
	{
		id:                  3,
		name:                "1 комнатная квартира, 50 кв.м. 3/9 эт.",
		price:               200_000,
		area:                "Худжанд",
		addressHouse:        "р.Спитамен, Кахрамон 1",
		distanceOfTheCenter: 250,
		postTime:            1_578_235_602,
		seller:              "Олег",
	},
	{
		id:                  4,
		name:                "4 комнатная квартира, 120 кв.м. 3/4 эт.",
		price:               280_000,
		area:                "Куляб",
		addressHouse:        "р.Абвгд, Дониш 19",
		distanceOfTheCenter: 180,
		postTime:            1_578_235_694,
		seller:              "Игор",
	},
}

func ExampleSortByPriceForDescending() {
	result := sortByPriceForDescending(homesSale)
	fmt.Println(result)
	// Output: [{2 2 комнатная квартира, 110 кв.м. 3/25 эт. 650000 Душанбе р.Фирдовси, Дусти 15 13 1578233112 Петя} {1 1 комнатная квартира, 99 кв.м. 9/25 эт. 600000 Душанбе р.Сино, Борбад 33 10 1578232244 Вася} {4 4 комнатная квартира, 120 кв.м. 3/4 эт. 280000 Куляб р., Дусти 1 180 1578235694 Игор} {3 1 комнатная квартира, 50 кв.м. 3/9 эт. 200000 Худжанд р.Спитамен, Кахрамон 1 250 1578235602 Олег}]
}

func ExampleSortByPriceForAscending() {
	result := sortByPriceForAscending(homesSale)
	fmt.Println(result)
	// Output: [{3 1 комнатная квартира, 50 кв.м. 3/9 эт. 200000 Худжанд р.Спитамен, Кахрамон 1 250 1578235602 Олег} {4 4 комнатная квартира, 120 кв.м. 3/4 эт. 280000 Куляб р., Дусти 1 180 1578235694 Игор} {1 1 комнатная квартира, 99 кв.м. 9/25 эт. 600000 Душанбе р.Сино, Борбад 33 10 1578232244 Вася} {2 2 комнатная квартира, 110 кв.м. 3/25 эт. 650000 Душанбе р.Фирдовси, Дусти 15 13 1578233112 Петя}]
}

func ExampleSortByDistanceOfCenterClose_LongAway() {
	result := sortByDistanceOfCenterClose_LongAway(homesSale)
	fmt.Println(result)
	// Output: [{1 1 комнатная квартира, 99 кв.м. 9/25 эт. 600000 Душанбе р.Сино, Борбад 33 10 1578232244 Вася} {2 2 комнатная квартира, 110 кв.м. 3/25 эт. 650000 Душанбе р.Фирдовси, Дусти 15 13 1578233112 Петя} {4 4 комнатная квартира, 120 кв.м. 3/4 эт. 280000 Куляб р., Дусти 1 180 1578235694 Игор} {3 1 комнатная квартира, 50 кв.м. 3/9 эт. 200000 Худжанд р.Спитамен, Кахрамон 1 250 1578235602 Олег}]
}

func ExampleSortByDistanceOfCenterLongAway_Close() {
	result := sortByDistanceOfCenterLongAway_Close(homesSale)
	fmt.Println(result)
	// Output: [{3 1 комнатная квартира, 50 кв.м. 3/9 эт. 200000 Худжанд р.Спитамен, Кахрамон 1 250 1578235602 Олег} {4 4 комнатная квартира, 120 кв.м. 3/4 эт. 280000 Куляб р., Дусти 1 180 1578235694 Игор} {2 2 комнатная квартира, 110 кв.м. 3/25 эт. 650000 Душанбе р.Фирдовси, Дусти 15 13 1578233112 Петя} {1 1 комнатная квартира, 99 кв.м. 9/25 эт. 600000 Душанбе р.Сино, Борбад 33 10 1578232244 Вася}]
}

func ExampleSearchByMaxPrice_NoResults() {
	result := searchByMaxPrice(homesSale, 100_000)
	fmt.Println(result)
	// Output: []
}

func ExampleSearchByMaxPrice_OneResults() {
	result := searchByMaxPrice(homesSale, 220_000)
	fmt.Println(result)
	// Output: [{3 1 комнатная квартира, 50 кв.м. 3/9 эт. 200000 Худжанд р.Спитамен, Кахрамон 1 250 1578235602 Олег}]
}

func ExampleSearchByMaxPrice() {
	result := searchByMaxPrice(homesSale, 700_000)
	fmt.Println(result)
	// Output: [{3 1 комнатная квартира, 50 кв.м. 3/9 эт. 200000 Худжанд р.Спитамен, Кахрамон 1 250 1578235602 Олег} {4 4 комнатная квартира, 120 кв.м. 3/4 эт. 280000 Куляб р., Дусти 1 180 1578235694 Игор}]
}

func ExampleSearchByPrice_NoResults()  {
	result := searchByPrice(homesSale, 300_000)
	fmt.Println(result)
	// Output: []
}

func ExampleSearchByPrice_HaveResult()  {
	result := searchByPrice(homesSale, 600_000)
	fmt.Println(result)
	// Output: [{1 1 комнатная квартира, 99 кв.м. 9/25 эт. 600000 Душанбе р.Сино, Борбад 33 10 1578232244 Вася}]
}

func ExampleSearchByArea_NoResults()  {
	result := searchByArea(homesSale,"Москва")
	fmt.Println(result)
	// Output: []
}

func ExampleSearchByArea_HaveResults()  {
	result := searchByArea(homesSale,"Куляб")
	fmt.Println(result)
	// Output: [{4 4 комнатная квартира, 120 кв.м. 3/4 эт. 280000 Куляб р., Дусти 1 180 1578235694 Игор}]
}

func ExampleSearchByArea_TwoResults()  {
	result := searchByArea(homesSale,"Душанбе")
	fmt.Println(result)
	// Output: [{1 1 комнатная квартира, 99 кв.м. 9/25 эт. 600000 Душанбе р.Сино, Борбад 33 10 1578232244 Вася} {2 2 комнатная квартира, 110 кв.м. 3/25 эт. 650000 Душанбе р.Фирдовси, Дусти 15 13 1578233112 Петя}]
}

func ExampleSortByPostTime()  {
	result := sortByPostTime(homesSale)
	fmt.Println(result)
	// Output: [{4 4 комнатная квартира, 120 кв.м. 3/4 эт. 280000 Куляб р., Дусти 1 180 1578235694 Игор} {3 1 комнатная квартира, 50 кв.м. 3/9 эт. 200000 Худжанд р.Спитамен, Кахрамон 1 250 1578235602 Олег} {2 2 комнатная квартира, 110 кв.м. 3/25 эт. 650000 Душанбе р.Фирдовси, Дусти 15 13 1578233112 Петя} {1 1 комнатная квартира, 99 кв.м. 9/25 эт. 600000 Душанбе р.Сино, Борбад 33 10 1578232244 Вася}]
}
