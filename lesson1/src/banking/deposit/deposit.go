package deposit

type Deposit struct {
	StartSum             float64
	Months               int
	Percent              float64
	Capitalization       bool
	CapitalizationPeriod int
	countResults         DepositResult
}

type DepositResult struct {
	DepositSum  float64
	FinalSum    float64
	PercentsSum float64
	MonthlyList []DepositResultMonthly
}

type DepositResultMonthly struct {
	Month       int
	DepositSum  float64
	FinalSum    float64
	PercentsSum float64
}

func (d *Deposit) GetCountResults() DepositResult {
	return d.countResults
}

func (d *Deposit) Count() {
	d.countResults = DepositResult{
		DepositSum:  d.StartSum,
		FinalSum:    d.StartSum,
		MonthlyList: make([]DepositResultMonthly, 0),
	}
	for month := 1; month < d.Months+1; month++ {
		monthPercents := d.countResults.DepositSum / 100 * 5 / 12
		d.countResults.PercentsSum += monthPercents
		d.countResults.FinalSum += monthPercents
		if d.Capitalization {
			if month%d.CapitalizationPeriod == 0 {
				d.countResults.DepositSum += monthPercents
			}
		}
		d.countResults.MonthlyList = append(d.countResults.MonthlyList, DepositResultMonthly{
			Month:       month,
			DepositSum:  d.countResults.DepositSum,
			FinalSum:    d.countResults.FinalSum,
			PercentsSum: monthPercents,
		})
	}
}
