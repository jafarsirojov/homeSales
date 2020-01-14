package main

import (
	"sort"
)

type home struct {
	id int64
	name string
	price int64
	area string
	addressHouse string
	distanceOfTheCenter int
	postTime int64
	seller string
}


func sortByPriceForDescending(homes []home) []home {
	result := make([]home, len(homes))
	copy(result, homes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price > result[j].price
	})
	return result
}

func sortByPriceForAscending(homes []home) []home {
	result := make([]home, len(homes))
	copy(result, homes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price < result[j].price
	})
	return result
}



func sortByDistanceOfCenterClose_LongAway(homes []home) []home {
	result := make([]home, len(homes))
	copy(result,homes)
	sort.Slice(result, func( i, j int) bool{
		return result[i].distanceOfTheCenter < result[j].distanceOfTheCenter
	})
	return result
}

func sortByDistanceOfCenterLongAway_Close(homes []home) []home {
	result := make([]home, len(homes))
	copy(result,homes)
	sort.Slice(result, func( i, j int) bool{
		return result[i].distanceOfTheCenter > result[j].distanceOfTheCenter
	})
	return result
}


func searchByMaxPrice(homes []home, maxPrice int64)  []home {
	result := make([]home, 0)
	for _, home := range homes{
		if home.price <= maxPrice{
			result = append(result,home)
		}
	}
	return result
}


func searchByPrice(homes []home, price int64) []home {
	result := make([]home,0)
	for _, home :=range homes{
		if home.price == price{
			result = append(result,home)
		}
	}
	return result
}


func searchByArea(homes []home, area string) []home {
	result := make([]home, 0)
	for _, home := range homes{
		if home.area == area{
			result = append(result,home)
		}
	}
	return result
}

func sortByPostTime(homes []home) []home {
	result := make([]home, len(homes))
	copy(result,homes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].postTime > result[j].postTime
	})
	return result
}


func main() {}
