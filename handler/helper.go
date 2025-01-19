package handler

import "math/rand"

func isSeatIsAvailable(seat string, availableSeats []string) bool {
	for _, v := range availableSeats {
		if seat == v {
			return true
		}

	}
	return false

}

func removeUserPickSeat(seats []string, element string) []string {
	var result []string
	for _, seat := range seats {
		if seat != element {
			result = append(result, seat)
		}
	}
	return result
}

func getRandomSeat(arr []string) (string, []string) {
	slice := arr[:]
	if len(slice) == 0 {
		return "", slice
	}
	randomIndex := rand.Intn(len(slice))
	randomValue := slice[randomIndex]
	slice = append(slice[:randomIndex], slice[randomIndex+1:]...)

	return randomValue, slice
}

func StoreUserBookingDeatils(selectedSeat, bookingDate, firstName, lastName, emailId string,
	bookingDetails map[string][]BookingSeat) map[string][]BookingSeat {
	if bookingDetails == nil {
		bookingDetails = make(map[string][]BookingSeat)
	}
	userName := firstName + lastName
	if val, ok := bookingDetails[emailId]; ok {
		bookingSeat := BookingSeat{
			SeatNumber: selectedSeat,
			Date:       bookingDate,
			Name:       userName,
			Staus:      bookingActiveStatus,
		}
		updatedSeat := append(val, bookingSeat)
		bookingDetails[emailId] = updatedSeat
	} else {
		bookingSeat := []BookingSeat{
			{
				SeatNumber: selectedSeat,
				Date:       bookingDate,
				Name:       userName,
				Staus:      bookingActiveStatus,
			},
		}
		bookingDetails[emailId] = bookingSeat
	}
	return bookingDetails
}
