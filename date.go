package calender
import 
"errors"

type Date struct {
	year int
	month int
	day int
}

func(d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid year")
	}
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int) {
	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) {
	
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}
	d.day = day
	return nil
}
