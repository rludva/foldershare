package foldershare

type FileTime struct {
	year, month, day              int
	hour, minute, second, msecond int
}

func (ft FileTime) Equal(n FileTime) bool {
	return ft.year == n.year && ft.month == n.month && ft.day == n.day &&
		ft.hour == n.hour && ft.minute == n.minute && ft.second == n.second && ft.msecond == n.msecond
}

func (a FileTime) IsNewer(b FileTime) bool {
	if a.year > b.year {
		return true
	}
	if a.year < b.year {
		return false
	}

	if a.month > b.month {
		return true
	}
	if a.month < b.month {
		return false
	}

	if a.day > b.day {
		return true
	}
	if a.day < b.day {
		return false
	}

	if a.hour > b.hour {
		return true
	}

	if a.hour < b.hour {
		return false
	}
	if a.minute > b.minute {
		return true
	}

	if a.minute < b.minute {
		return false
	}

	if a.second > b.second {
		return true
	}
	if a.second < b.second {
		return false
	}

	if a.msecond > b.msecond {
		return true
	}
	if a.msecond < b.msecond {
		return false
	}

	return false
}

