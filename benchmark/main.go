package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Yob      int    `json:"yob"`
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
	//Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func ReadJson() (result Users) {

	byteValue, err := ioutil.ReadFile("users.json")

	if err != nil {
		log.Fatal(err)
	}

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	/*
		for _, user := range users.Users {
			result = result + user.Name
		}

	*/

	return users //result
}

func ReadJsonOne() (result User) {

	byteValue, err := ioutil.ReadFile("user.json")

	if err != nil {
		log.Fatal(err)
	}

	var u User

	json.Unmarshal(byteValue, &u)

	return u //result

}

//type Charge []struct {
type Charge struct {
	TYPE     string `json:"TYPE"`
	ID       string `json:"ID"`
	FIELD1   string `json:"FIELD_1"`
	FIELD2   string `json:"FIELD_2"`
	FIELD3   string `json:"FIELD_3"`
	FIELD4   string `json:"FIELD_4"`
	FIELD5   string `json:"FIELD_5"`
	FIELD6   string `json:"FIELD_6"`
	FIELD7   string `json:"FIELD_7"`
	FIELD8   string `json:"FIELD_8"`
	FIELD9   int    `json:"FIELD_9"`
	FIELD10  string `json:"FIELD_10"`
	FIELD11  string `json:"FIELD_11"`
	FIELD12  string `json:"FIELD_12"`
	FIELD13  string `json:"FIELD_13"`
	FIELD14  string `json:"FIELD_14"`
	FIELD15  string `json:"FIELD_15"`
	FIELD16  string `json:"FIELD_16"`
	FIELD17  string `json:"FIELD_17"`
	FIELD18  int    `json:"FIELD_18"`
	FIELD19  string `json:"FIELD_19"`
	FIELD20  string `json:"FIELD_20"`
	FIELD21  string `json:"FIELD_21"`
	FIELD22  string `json:"FIELD_22"`
	FIELD23  string `json:"FIELD_23"`
	FIELD24  string `json:"FIELD_24"`
	FIELD25  string `json:"FIELD_25"`
	FIELD26  string `json:"FIELD_26"`
	FIELD27  int    `json:"FIELD_27"`
	FIELD28  string `json:"FIELD_28"`
	FIELD29  string `json:"FIELD_29"`
	FIELD30  string `json:"FIELD_30"`
	FIELD31  string `json:"FIELD_31"`
	FIELD32  string `json:"FIELD_32"`
	FIELD33  string `json:"FIELD_33"`
	FIELD34  string `json:"FIELD_34"`
	FIELD35  int    `json:"FIELD_35"`
	FIELD36  string `json:"FIELD_36"`
	FIELD37  string `json:"FIELD_37"`
	FIELD38  string `json:"FIELD_38"`
	FIELD39  string `json:"FIELD_39"`
	FIELD40  string `json:"FIELD_40"`
	FIELD41  string `json:"FIELD_41"`
	FIELD42  string `json:"FIELD_42"`
	FIELD43  string `json:"FIELD_43"`
	FIELD44  string `json:"FIELD_44"`
	FIELD45  string `json:"FIELD_45"`
	FIELD46  string `json:"FIELD_46"`
	FIELD47  int    `json:"FIELD_47"`
	FIELD48  string `json:"FIELD_48"`
	FIELD49  string `json:"FIELD_49"`
	FIELD50  string `json:"FIELD_50"`
	FIELD51  string `json:"FIELD_51"`
	FIELD52  string `json:"FIELD_52"`
	FIELD53  string `json:"FIELD_53"`
	FIELD54  string `json:"FIELD_54"`
	FIELD55  string `json:"FIELD_55"`
	FIELD56  string `json:"FIELD_56"`
	FIELD57  string `json:"FIELD_57"`
	FIELD58  string `json:"FIELD_58"`
	FIELD59  string `json:"FIELD_59"`
	FIELD60  string `json:"FIELD_60"`
	FIELD61  string `json:"FIELD_61"`
	FIELD62  string `json:"FIELD_62"`
	FIELD63  string `json:"FIELD_63"`
	FIELD64  string `json:"FIELD_64"`
	FIELD65  string `json:"FIELD_65"`
	FIELD66  string `json:"FIELD_66"`
	FIELD67  string `json:"FIELD_67"`
	FIELD68  string `json:"FIELD_68"`
	FIELD69  string `json:"FIELD_69"`
	FIELD70  string `json:"FIELD_70"`
	FIELD71  string `json:"FIELD_71"`
	FIELD72  string `json:"FIELD_72"`
	FIELD73  string `json:"FIELD_73"`
	FIELD74  string `json:"FIELD_74"`
	FIELD75  string `json:"FIELD_75"`
	FIELD76  string `json:"FIELD_76"`
	FIELD77  string `json:"FIELD_77"`
	FIELD78  string `json:"FIELD_78"`
	FIELD79  string `json:"FIELD_79"`
	FIELD80  string `json:"FIELD_80"`
	FIELD81  string `json:"FIELD_81"`
	FIELD82  string `json:"FIELD_82"`
	FIELD83  string `json:"FIELD_83"`
	FIELD84  string `json:"FIELD_84"`
	FIELD85  string `json:"FIELD_85"`
	FIELD86  string `json:"FIELD_86"`
	FIELD87  string `json:"FIELD_87"`
	FIELD88  string `json:"FIELD_88"`
	FIELD89  string `json:"FIELD_89"`
	FIELD90  int    `json:"FIELD_90"`
	FIELD91  int    `json:"FIELD_91"`
	FIELD92  string `json:"FIELD_92"`
	FIELD93  string `json:"FIELD_93"`
	FIELD94  string `json:"FIELD_94"`
	FIELD95  string `json:"FIELD_95"`
	FIELD96  string `json:"FIELD_96"`
	FIELD97  string `json:"FIELD_97"`
	FIELD98  int    `json:"FIELD_98"`
	FIELD99  int    `json:"FIELD_99"`
	FIELD100 string `json:"FIELD_100"`
	FIELD101 string `json:"FIELD_101"`
	FIELD102 string `json:"FIELD_102"`
	FIELD103 string `json:"FIELD_103"`
	FIELD104 string `json:"FIELD_104"`
	FIELD105 string `json:"FIELD_105"`
	FIELD106 string `json:"FIELD_106"`
	FIELD107 string `json:"FIELD_107"`
	FIELD108 string `json:"FIELD_108"`
	FIELD109 string `json:"FIELD_109"`
	FIELD110 string `json:"FIELD_110"`
	FIELD111 string `json:"FIELD_111"`
	FIELD112 string `json:"FIELD_112"`
	FIELD113 string `json:"FIELD_113"`
	FIELD114 string `json:"FIELD_114"`
	FIELD115 string `json:"FIELD_115"`
	FIELD116 string `json:"FIELD_116"`
	FIELD117 int    `json:"FIELD_117"`
	FIELD118 string `json:"FIELD_118"`
	FIELD119 string `json:"FIELD_119"`
	FIELD120 string `json:"FIELD_120"`
	FIELD121 string `json:"FIELD_121"`
	FIELD122 string `json:"FIELD_122"`
	FIELD123 string `json:"FIELD_123"`
	FIELD124 string `json:"FIELD_124"`
	FIELD125 string `json:"FIELD_125"`
	FIELD126 string `json:"FIELD_126"`
	FIELD127 string `json:"FIELD_127"`
	FIELD128 string `json:"FIELD_128"`
	FIELD129 string `json:"FIELD_129"`
	FIELD130 string `json:"FIELD_130"`
	FIELD131 string `json:"FIELD_131"`
	FIELD132 string `json:"FIELD_132"`
	FIELD133 string `json:"FIELD_133"`
	FIELD134 string `json:"FIELD_134"`
	FIELD135 string `json:"FIELD_135"`
	FIELD136 string `json:"FIELD_136"`
	FIELD137 string `json:"FIELD_137"`
	FIELD138 string `json:"FIELD_138"`
	FIELD139 string `json:"FIELD_139"`
	FIELD140 string `json:"FIELD_140"`
	FIELD141 string `json:"FIELD_141"`
	FIELD142 string `json:"FIELD_142"`
	FIELD143 string `json:"FIELD_143"`
}

func ReadJsonBig() (result Charge) {

	byteValue, err := ioutil.ReadFile("charge.json")

	if err != nil {
		log.Fatal(err)
	}

	var c Charge

	json.Unmarshal(byteValue, &c)

	return c //result

}

type Charge_num struct {
	TYPE   string `json:"TYPE"`
	ID     string `json:"ID"`
	Num1   string `json:"1"`
	Num2   string `json:"2"`
	Num3   string `json:"3"`
	Num4   string `json:"4"`
	Num5   string `json:"5"`
	Num6   string `json:"6"`
	Num7   string `json:"7"`
	Num8   string `json:"8"`
	Num9   int    `json:"9"`
	Num10  string `json:"10"`
	Num11  string `json:"11"`
	Num12  string `json:"12"`
	Num13  string `json:"13"`
	Num14  string `json:"14"`
	Num15  string `json:"15"`
	Num16  string `json:"16"`
	Num17  string `json:"17"`
	Num18  int    `json:"18"`
	Num19  string `json:"19"`
	Num20  string `json:"20"`
	Num21  string `json:"21"`
	Num22  string `json:"22"`
	Num23  string `json:"23"`
	Num24  string `json:"24"`
	Num25  string `json:"25"`
	Num26  string `json:"26"`
	Num27  int    `json:"27"`
	Num28  string `json:"28"`
	Num29  string `json:"29"`
	Num30  string `json:"30"`
	Num31  string `json:"31"`
	Num32  string `json:"32"`
	Num33  string `json:"33"`
	Num34  string `json:"34"`
	Num35  int    `json:"35"`
	Num36  string `json:"36"`
	Num37  string `json:"37"`
	Num38  string `json:"38"`
	Num39  string `json:"39"`
	Num40  string `json:"40"`
	Num41  string `json:"41"`
	Num42  string `json:"42"`
	Num43  string `json:"43"`
	Num44  string `json:"44"`
	Num45  string `json:"45"`
	Num46  string `json:"46"`
	Num47  int    `json:"47"`
	Num48  string `json:"48"`
	Num49  string `json:"49"`
	Num50  string `json:"50"`
	Num51  string `json:"51"`
	Num52  string `json:"52"`
	Num53  string `json:"53"`
	Num54  string `json:"54"`
	Num55  string `json:"55"`
	Num56  string `json:"56"`
	Num57  string `json:"57"`
	Num58  string `json:"58"`
	Num59  string `json:"59"`
	Num60  string `json:"60"`
	Num61  string `json:"61"`
	Num62  string `json:"62"`
	Num63  string `json:"63"`
	Num64  string `json:"64"`
	Num65  string `json:"65"`
	Num66  string `json:"66"`
	Num67  string `json:"67"`
	Num68  string `json:"68"`
	Num69  string `json:"69"`
	Num70  string `json:"70"`
	Num71  string `json:"71"`
	Num72  string `json:"72"`
	Num73  string `json:"73"`
	Num74  string `json:"74"`
	Num75  string `json:"75"`
	Num76  string `json:"76"`
	Num77  string `json:"77"`
	Num78  string `json:"78"`
	Num79  string `json:"79"`
	Num80  string `json:"80"`
	Num81  string `json:"81"`
	Num82  string `json:"82"`
	Num83  string `json:"83"`
	Num84  string `json:"84"`
	Num85  string `json:"85"`
	Num86  string `json:"86"`
	Num87  string `json:"87"`
	Num88  string `json:"88"`
	Num89  string `json:"89"`
	Num90  int    `json:"90"`
	Num91  int    `json:"91"`
	Num92  string `json:"92"`
	Num93  string `json:"93"`
	Num94  string `json:"94"`
	Num95  string `json:"95"`
	Num96  string `json:"96"`
	Num97  string `json:"97"`
	Num98  int    `json:"98"`
	Num99  int    `json:"99"`
	Num100 string `json:"100"`
	Num101 string `json:"101"`
	Num102 string `json:"102"`
	Num103 string `json:"103"`
	Num104 string `json:"104"`
	Num105 string `json:"105"`
	Num106 string `json:"106"`
	Num107 string `json:"107"`
	Num108 string `json:"108"`
	Num109 string `json:"109"`
	Num110 string `json:"110"`
	Num111 string `json:"111"`
	Num112 string `json:"112"`
	Num113 string `json:"113"`
	Num114 string `json:"114"`
	Num115 string `json:"115"`
	Num116 string `json:"116"`
	Num117 int    `json:"117"`
	Num118 string `json:"118"`
	Num119 string `json:"119"`
	Num120 string `json:"120"`
	Num121 string `json:"121"`
	Num122 string `json:"122"`
	Num123 string `json:"123"`
	Num124 string `json:"124"`
	Num125 string `json:"125"`
	Num126 string `json:"126"`
	Num127 string `json:"127"`
	Num128 string `json:"128"`
	Num129 string `json:"129"`
	Num130 string `json:"130"`
	Num131 string `json:"131"`
	Num132 string `json:"132"`
	Num133 string `json:"133"`
	Num134 string `json:"134"`
	Num135 string `json:"135"`
	Num136 string `json:"136"`
	Num137 string `json:"137"`
	Num138 string `json:"138"`
	Num139 string `json:"139"`
	Num140 string `json:"140"`
	Num141 string `json:"141"`
	Num142 string `json:"142"`
	Num143 string `json:"143"`
}

func ReadJsonBigProto() (result Charge_num) {

	byteValue, err := ioutil.ReadFile("charge_num.json")

	if err != nil {
		log.Fatal(err)
	}

	var c Charge_num

	json.Unmarshal(byteValue, &c)

	return c //result
}

func ReadCsv() [][]string {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	return records
}

func ReadCsvOne() [][]string {
	f, err := os.Open("user.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	return records
}

func ReadCsvBig() [][]string {
	f, err := os.Open("charge.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	return records
}

func main() {
	//fmt.Println(ReadJson())
	fmt.Println(ReadJsonBig())
	fmt.Println(ReadJsonBigProto())
	//fmt.Println()
	//fmt.Println(ReadCsv())
	//fmt.Println(ReadCsvBig())
}
