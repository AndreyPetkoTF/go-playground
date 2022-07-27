package main

func totalFruit(fruits []int) int {
	totalAmountForAllIndexes := 0
	for startIndex := 0; startIndex < len(fruits); startIndex++ {
		indexTotalAmount := calculateAmount(startIndex, fruits)

		if indexTotalAmount > totalAmountForAllIndexes {
			totalAmountForAllIndexes = indexTotalAmount
		}

		if totalAmountForAllIndexes > len(fruits)-startIndex+1 {
			break
		}
	}

	return totalAmountForAllIndexes
}

func calculateAmount(startIndex int, fruits []int) int {
	totalAmount := 1
	firstBucketType := fruits[startIndex]
	secondBucketType := -1

	for i := startIndex + 1; i < len(fruits); i++ {
		typeNextThree := fruits[i]
		if secondBucketType == -1 && firstBucketType != typeNextThree {
			secondBucketType = typeNextThree
			break
		}
	}

	for i := startIndex + 1; i < len(fruits); i++ {
		typeNextThree := fruits[i]

		if firstBucketType == typeNextThree || secondBucketType == typeNextThree {
			totalAmount++
		} else {
			break
		}
	}

	return totalAmount
}
