package models

import (
	"github.com/adrianomr/investments/src/domain/models"
	"time"
)

var cdis_march = []models.Cdi{
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 1),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 4),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 5),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 6),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 7),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 8),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 11),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 12),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 13),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 14),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 15),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 18),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 19),
	},
	{
		Rate: 0.041957,
		Date: createDate(2024, 3, 20),
	},
	{
		Rate: 0.040168,
		Date: createDate(2024, 3, 21),
	},
	{
		Rate: 0.040168,
		Date: createDate(2024, 3, 22),
	},
	{
		Rate: 0.040168,
		Date: createDate(2024, 3, 25),
	},
	{
		Rate: 0.040168,
		Date: createDate(2024, 3, 26),
	},
	{
		Rate: 0.040168,
		Date: createDate(2024, 3, 27),
	},
	{
		Rate: 0.040168,
		Date: createDate(2024, 3, 28),
	},
}

var cdis_may = []models.Cdi{
	{
		Rate: 0.040168,
		Date: createDate(2024, 5, 2),
	},
	{
		Rate: 0.040168,
		Date: createDate(2024, 5, 3),
	},
	{
		Rate: 0.040168,
		Date: createDate(2024, 5, 6),
	},
	{
		Rate: 0.040168,
		Date: createDate(2024, 5, 7),
	},
	{
		Rate: 0.040168,
		Date: createDate(2024, 5, 8),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 9),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 10),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 13),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 14),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 15),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 16),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 17),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 20),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 21),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 22),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 23),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 24),
	},
	{
		Rate: 0.039270,
		Date: createDate(2024, 5, 27),
	},
}

func createDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 12, 12, 12, 0, time.UTC)
}
