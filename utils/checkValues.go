package utils

func CheckValue(condition bool, value, defaultValue string) string {
    if condition == false {
        return defaultValue
    }

    return value
}
