package tools

func Unique(slice []int) []int {
    uniques := []int{}

    for _, v := range slice {
        if !Contains(uniques[:], v) {
            uniques = append(uniques, v)
        }
    }
    return uniques
}

func Contains(slice []int, c int) bool {
    for _, v := range slice {
        if v == c {
            return true
        }
    }
    return false
}