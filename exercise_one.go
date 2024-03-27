package go_unit_test_bootcamp

func FindMissingDrone(droneIds []int) int {
	missingDrone := 0

	for _, id := range droneIds {
		missingDrone ^= id
	}

	return missingDrone
}
