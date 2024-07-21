package main

func hireEmployee(t int, deal int) string {
	res := "Không thuê"
	if t <= 0 || t > 150 {
		res = "Invalid"
	} else if t >= 16 && t <= 18 {
		if deal <= 500 {
			res = "Thuê bán thời gian"
		}
	} else if t >= 18 && t <= 55 {
		if deal <= 1000 {
			res = "Thuê toàn thời gian"
		}
	}
	return res
}
