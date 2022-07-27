package main

import "testing"

func Test_findRepeat(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{arr: []string{
				"2278189940000010213",
				"2273069410000010911",
				"2275277760000013050",
				"2265405500000011799",
				"2282256530000015386",
				"2265120890000010213",
				"2273120950000019307",
				"2280727440000010213",
				"2264998350000018427",
				"2282249630000012282",
				"2277111830000019717",
				"2266822870000018820",
				"2262239830000019460",
				"2279085230000012642",
				"2282223610000012865",
				"2277795830000019979",
				"2276955300000011959",
				"2270176470000018274",
				"2268672570000010213",
				"2282213280000019958",
				"2264000090000014216",
				"2277268350000015323",
				"2273141890000016430",
				"2266708330000015323",
				"2282209640000017205",
				"2263483140000017331",
				"2263172020000017407",
				"2281819370000017695",
				"2279368010000011801",
				"2282206890000018421",
				"2278787660000012910",
				"2273943620000015281",
				"2273097870000011610",
				"2265033780000012551",
				"2282205370000010796",
				"2262667850000010213",
				"2280878800000012080",
				"2276104390000012593",
				"2274675080000010213",
				"2282200290000018300",
				"2271195280000010213",
				"2267127070000019007",
				"2266145850000011504",
				"2265971570000018483",
				"2282151520000010749",
				"2265002620000012203",
				"2264986040000016265",
				"2263105860000012401",
				"2262654890000012355",
				"2281871000000013638",
				"2280837750000012642",
				"2280633770000017230",
				"2280461200000010144",
				"2281843870000010296",
				"2280459540000010144",
				"2280107050000016952",
				"2280045710000019717",
				"2280031850000012642",
				"2281840920000017331",
				"2279801870000016715",
				"2279248520000017239",
				"2278062330000015397",
				"2278036330000013554",
				"2277344370000012642",
				"2276485590000019310",
				"2274723290000010562",
				"2273272160000013624",
				"2281807730000017317",
				"2272805440000011751",
				"2270972620000017940",
				"2270064520000019460",
				"2268746900000019613",
				"2281770920000012949",
				"2267739450000010213",
				"2267477670000019015",
				"2267138940000019014",
				"2267135300000018981",
				"2281750080000013935",
				"2267120770000019004",
				"2267105820000010607",
				"2267090390000018981",
				"2266875730000018847",
				"2281712090000019726",
				"2266090130000018559",
				"2266025560000016278",
				"2265438680000016270",
				"2265223810000010213",
				"2281706530000018015",
				"2264981640000015707",
				"2262729880000017323",
				"2262723850000017318",
				"2281662160000018572",
				"2280771950000017104",
				"2281640210000015062",
				"2280770490000018971",
				"2280760050000010213",
				"2280718310000013946",
				"2280686020000019589",
				"2281592590000019890",
				"2280670980000010854",
				"2280662290000019750",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findRepeat(tt.args.arr)
		})
	}
}
