package iso3166

type Iso3166 struct {
	nameEn      string
	nameFr      string
	alpha2Code  string
	alpha3Code  string
	numericCode string
}

type Isos3166 map[string]*Iso3166

func (iso Isos3166) Iso3166ByNumericCode(code string) *Iso3166 {
	for _, sc := range iso {
		if sc.numericCode == code {
			return sc
		}
	}

	return nil
}

func (iso Isos3166) Iso3166ByAlpha2Code(code string) *Iso3166 {
	for _, sc := range iso {
		if sc.alpha2Code == code {
			return sc
		}
	}

	return nil
}

func (iso Isos3166) Iso3166ByAlpha3Code(code string) *Iso3166 {
	sc, ok := iso[code]
	if !ok {
		return nil
	}

	return sc
}

func ExistsIso3166ByAlpha3Code(code string) bool {
	return isos3166.Iso3166ByAlpha3Code(code) != nil
}

func ExistsIso3166ByAlpha2Code(code string) bool {
	return isos3166.Iso3166ByAlpha2Code(code) != nil
}

func ExistsIso3166ByNumericCode(code string) bool {
	return isos3166.Iso3166ByNumericCode(code) != nil
}

var isos3166 = Isos3166{
	"ABW": {nameEn: "Aruba", nameFr: "Aruba", alpha2Code: "AW", alpha3Code: "ABW", numericCode: "533"},
	"AFG": {nameEn: "Afghanistan", nameFr: "Afghanistan (l')", alpha2Code: "AF", alpha3Code: "AFG", numericCode: "004"},
	"AGO": {nameEn: "Angola", nameFr: "Angola (l')", alpha2Code: "AO", alpha3Code: "AGO", numericCode: "024"},
	"AIA": {nameEn: "Anguilla", nameFr: "Anguilla", alpha2Code: "AI", alpha3Code: "AIA", numericCode: "660"},
	"ALA": {nameEn: "Åland Islands", nameFr: "Åland(les Îles)", alpha2Code: "AX", alpha3Code: "ALA", numericCode: "248"},
	"ALB": {nameEn: "Albania", nameFr: "Albanie (l')", alpha2Code: "AL", alpha3Code: "ALB", numericCode: "008"},
	"AND": {nameEn: "Andorra", nameFr: "Andorre (l')", alpha2Code: "AD", alpha3Code: "AND", numericCode: "020"},
	"ARE": {nameEn: "United Arab Emirates (the)", nameFr: "Émirats arabes unis (les)", alpha2Code: "AE", alpha3Code: "ARE", numericCode: "784"},
	"ARG": {nameEn: "Argentina", nameFr: "Argentine (l')", alpha2Code: "AR", alpha3Code: "ARG", numericCode: "032"},
	"ARM": {nameEn: "Armenia", nameFr: "Arménie (l')", alpha2Code: "AM", alpha3Code: "ARM", numericCode: "051"},
	"ASM": {nameEn: "American Samoa", nameFr: "Samoa américaines (les)", alpha2Code: "AS", alpha3Code: "ASM", numericCode: "016"},
	"ATA": {nameEn: "Antarctica", nameFr: "Antarctique (l')", alpha2Code: "AQ", alpha3Code: "ATA", numericCode: "010"},
	"ATF": {nameEn: "French Southern Territories (the)", nameFr: "Terres australes françaises (les)", alpha2Code: "TF", alpha3Code: "ATF", numericCode: "260"},
	"ATG": {nameEn: "Antigua and Barbuda", nameFr: "Antigua-et-Barbuda", alpha2Code: "AG", alpha3Code: "ATG", numericCode: "028"},
	"AUS": {nameEn: "Australia", nameFr: "Australie (l')", alpha2Code: "AU", alpha3Code: "AUS", numericCode: "036"},
	"AUT": {nameEn: "Austria", nameFr: "Autriche (l')", alpha2Code: "AT", alpha3Code: "AUT", numericCode: "040"},
	"AZE": {nameEn: "Azerbaijan", nameFr: "Azerbaïdjan (l')", alpha2Code: "AZ", alpha3Code: "AZE", numericCode: "031"},
	"BDI": {nameEn: "Burundi", nameFr: "Burundi (le)", alpha2Code: "BI", alpha3Code: "BDI", numericCode: "108"},
	"BEL": {nameEn: "Belgium", nameFr: "Belgique (la)", alpha2Code: "BE", alpha3Code: "BEL", numericCode: "056"},
	"BEN": {nameEn: "Benin", nameFr: "Bénin (le)", alpha2Code: "BJ", alpha3Code: "BEN", numericCode: "204"},
	"BES": {nameEn: "Bonaire, Sint Eustatius and Saba", nameFr: "Bonaire, Saint-Eustache et Saba", alpha2Code: "BQ", alpha3Code: "BES", numericCode: "535"},
	"BFA": {nameEn: "Burkina Faso", nameFr: "Burkina Faso (le)", alpha2Code: "BF", alpha3Code: "BFA", numericCode: "854"},
	"BGD": {nameEn: "Bangladesh", nameFr: "Bangladesh (le)", alpha2Code: "BD", alpha3Code: "BGD", numericCode: "050"},
	"BGR": {nameEn: "Bulgaria", nameFr: "Bulgarie (la)", alpha2Code: "BG", alpha3Code: "BGR", numericCode: "100"},
	"BHR": {nameEn: "Bahrain", nameFr: "Bahreïn", alpha2Code: "BH", alpha3Code: "BHR", numericCode: "048"},
	"BHS": {nameEn: "Bahamas (the)", nameFr: "Bahamas (les)", alpha2Code: "BS", alpha3Code: "BHS", numericCode: "044"},
	"BIH": {nameEn: "Bosnia and Herzegovina", nameFr: "Bosnie-Herzégovine (la)", alpha2Code: "BA", alpha3Code: "BIH", numericCode: "070"},
	"BLM": {nameEn: "Saint Barthélemy", nameFr: "Saint-Barthélemy", alpha2Code: "BL", alpha3Code: "BLM", numericCode: "652"},
	"BLR": {nameEn: "Belarus", nameFr: "Bélarus (le)", alpha2Code: "BY", alpha3Code: "BLR", numericCode: "112"},
	"BLZ": {nameEn: "Belize", nameFr: "Belize (le)", alpha2Code: "BZ", alpha3Code: "BLZ", numericCode: "084"},
	"BMU": {nameEn: "Bermuda", nameFr: "Bermudes (les)", alpha2Code: "BM", alpha3Code: "BMU", numericCode: "060"},
	"BOL": {nameEn: "Bolivia (Plurinational State of)", nameFr: "Bolivie (État plurinational de)", alpha2Code: "BO", alpha3Code: "BOL", numericCode: "068"},
	"BRA": {nameEn: "Brazil", nameFr: "Brésil (le)", alpha2Code: "BR", alpha3Code: "BRA", numericCode: "076"},
	"BRB": {nameEn: "Barbados", nameFr: "Barbade (la)", alpha2Code: "BB", alpha3Code: "BRB", numericCode: "052"},
	"BRN": {nameEn: "Brunei Darussalam", nameFr: "Brunéi Darussalam (le)", alpha2Code: "BN", alpha3Code: "BRN", numericCode: "096"},
	"BTN": {nameEn: "Bhutan", nameFr: "Bhoutan (le)", alpha2Code: "BT", alpha3Code: "BTN", numericCode: "064"},
	"BVT": {nameEn: "Bouvet Island", nameFr: "Bouvet (l'Île)", alpha2Code: "BV", alpha3Code: "BVT", numericCode: "074"},
	"BWA": {nameEn: "Botswana", nameFr: "Botswana (le)", alpha2Code: "BW", alpha3Code: "BWA", numericCode: "072"},
	"CAF": {nameEn: "Central African Republic (the)", nameFr: "République centrafricaine (la)", alpha2Code: "CF", alpha3Code: "CAF", numericCode: "140"},
	"CAN": {nameEn: "Canada", nameFr: "Canada (le)", alpha2Code: "CA", alpha3Code: "CAN", numericCode: "124"},
	"CCK": {nameEn: "Cocos (Keeling) Islands (the)", nameFr: "Cocos (les Îles)/ Keeling (les Îles)", alpha2Code: "CC", alpha3Code: "CCK", numericCode: "166"},
	"CHE": {nameEn: "Switzerland", nameFr: "Suisse (la)", alpha2Code: "CH", alpha3Code: "CHE", numericCode: "756"},
	"CHL": {nameEn: "Chile", nameFr: "Chili (le)", alpha2Code: "CL", alpha3Code: "CHL", numericCode: "152"},
	"CHN": {nameEn: "China", nameFr: "Chine (la)", alpha2Code: "CN", alpha3Code: "CHN", numericCode: "156"},
	"CIV": {nameEn: "Côte d'Ivoire", nameFr: "Côte d'Ivoire (la)", alpha2Code: "CI", alpha3Code: "CIV", numericCode: "384"},
	"CMR": {nameEn: "Cameroon", nameFr: "Cameroun (le)", alpha2Code: "CM", alpha3Code: "CMR", numericCode: "120"},
	"COD": {nameEn: "Congo (the Democratic Republic of the)", nameFr: "Congo (la République démocratique du)", alpha2Code: "CD", alpha3Code: "COD", numericCode: "180"},
	"COG": {nameEn: "Congo (the)", nameFr: "Congo (le)", alpha2Code: "CG", alpha3Code: "COG", numericCode: "178"},
	"COK": {nameEn: "Cook Islands (the)", nameFr: "Cook (les Îles)", alpha2Code: "CK", alpha3Code: "COK", numericCode: "184"},
	"COL": {nameEn: "Colombia", nameFr: "Colombie (la)", alpha2Code: "CO", alpha3Code: "COL", numericCode: "170"},
	"COM": {nameEn: "Comoros (the)", nameFr: "Comores (les)", alpha2Code: "KM", alpha3Code: "COM", numericCode: "174"},
	"CPV": {nameEn: "Cabo Verde", nameFr: "Cabo Verde", alpha2Code: "CV", alpha3Code: "CPV", numericCode: "132"},
	"CRI": {nameEn: "Costa Rica", nameFr: "Costa Rica (le)", alpha2Code: "CR", alpha3Code: "CRI", numericCode: "188"},
	"CUB": {nameEn: "Cuba", nameFr: "Cuba", alpha2Code: "CU", alpha3Code: "CUB", numericCode: "192"},
	"CUW": {nameEn: "Curaçao", nameFr: "Curaçao", alpha2Code: "CW", alpha3Code: "CUW", numericCode: "531"},
	"CXR": {nameEn: "Christmas Island,Christmas (l'Île)", alpha2Code: "CX", alpha3Code: "CXR", numericCode: "162"},
	"CYM": {nameEn: "Cayman Islands (the)", nameFr: "Caïmans (les Îles)", alpha2Code: "KY", alpha3Code: "CYM", numericCode: "136"},
	"CYP": {nameEn: "Cyprus", nameFr: "Chypre", alpha2Code: "CY", alpha3Code: "CYP", numericCode: "196"},
	"CZE": {nameEn: "Czechia", nameFr: "Tchéquie (la)", alpha2Code: "CZ", alpha3Code: "CZE", numericCode: "203"},
	"DEU": {nameEn: "Germany", nameFr: "Allemagne (l')", alpha2Code: "DE", alpha3Code: "DEU", numericCode: "276"},
	"DJI": {nameEn: "Djibouti", nameFr: "Djibouti", alpha2Code: "DJ", alpha3Code: "DJI", numericCode: "262"},
	"DMA": {nameEn: "Dominica", nameFr: "Dominique (la)", alpha2Code: "DM", alpha3Code: "DMA", numericCode: "212"},
	"DNK": {nameEn: "Denmark", nameFr: "Danemark (le)", alpha2Code: "DK", alpha3Code: "DNK", numericCode: "208"},
	"DOM": {nameEn: "Dominican Republic (the)", nameFr: "dominicaine (la République)", alpha2Code: "DO", alpha3Code: "DOM", numericCode: "214"},
	"DZA": {nameEn: "Algeria", nameFr: "Algérie (l')", alpha2Code: "DZ", alpha3Code: "DZA", numericCode: "012"},
	"ECU": {nameEn: "Ecuador", nameFr: "Équateur (l')", alpha2Code: "EC", alpha3Code: "ECU", numericCode: "218"},
	"EGY": {nameEn: "Egypt", nameFr: "Égypte (l')", alpha2Code: "EG", alpha3Code: "EGY", numericCode: "818"},
	"ERI": {nameEn: "Eritrea", nameFr: "Érythrée (l')", alpha2Code: "ER", alpha3Code: "ERI", numericCode: "232"},
	"ESH": {nameEn: "Western Sahara*", nameFr: "Sahara occidental (le)*", alpha2Code: "EH", alpha3Code: "ESH", numericCode: "732"},
	"ESP": {nameEn: "Spain", nameFr: "Espagne (l')", alpha2Code: "ES", alpha3Code: "ESP", numericCode: "724"},
	"EST": {nameEn: "Estonia", nameFr: "Estonie (l')", alpha2Code: "EE", alpha3Code: "EST", numericCode: "233"},
	"ETH": {nameEn: "Ethiopia", nameFr: "Éthiopie (l')", alpha2Code: "ET", alpha3Code: "ETH", numericCode: "231"},
	"FIN": {nameEn: "Finland", nameFr: "Finlande (la)", alpha2Code: "FI", alpha3Code: "FIN", numericCode: "246"},
	"FJI": {nameEn: "Fiji", nameFr: "Fidji (les)", alpha2Code: "FJ", alpha3Code: "FJI", numericCode: "242"},
	"FLK": {nameEn: "Falkland Islands (the) [Malvinas]", nameFr: "Falkland (les Îles)/Malouines (les Îles)", alpha2Code: "FK", alpha3Code: "FLK", numericCode: "238"},
	"FRA": {nameEn: "France", nameFr: "France (la)", alpha2Code: "FR", alpha3Code: "FRA", numericCode: "250"},
	"FRO": {nameEn: "Faroe Islands (the)", nameFr: "Féroé (les Îles)", alpha2Code: "FO", alpha3Code: "FRO", numericCode: "234"},
	"FSM": {nameEn: "Micronesia (Federated States of)", nameFr: "Micronésie (États fédérés de)", alpha2Code: "FM", alpha3Code: "FSM", numericCode: "583"},
	"GAB": {nameEn: "Gabon", nameFr: "Gabon (le)", alpha2Code: "GA", alpha3Code: "GAB", numericCode: "266"},
	"GBR": {nameEn: "United Kingdom of Great Britain and Northern Ireland (the)", nameFr: "Royaume-Uni de Grande-Bretagne et d'Irlande du Nord (le)", alpha2Code: "GB", alpha3Code: "GBR", numericCode: "826"},
	"GEO": {nameEn: "Georgia", nameFr: "Géorgie (la)", alpha2Code: "GE", alpha3Code: "GEO", numericCode: "268"},
	"GGY": {nameEn: "Guernsey", nameFr: "Guernesey", alpha2Code: "GG", alpha3Code: "GGY", numericCode: "831"},
	"GHA": {nameEn: "Ghana", nameFr: "Ghana (le)", alpha2Code: "GH", alpha3Code: "GHA", numericCode: "288"},
	"GIB": {nameEn: "Gibraltar", nameFr: "Gibraltar", alpha2Code: "GI", alpha3Code: "GIB", numericCode: "292"},
	"GIN": {nameEn: "Guinea", nameFr: "Guinée (la)", alpha2Code: "GN", alpha3Code: "GIN", numericCode: "324"},
	"GLP": {nameEn: "Guadeloupe", nameFr: "Guadeloupe (la)", alpha2Code: "GP", alpha3Code: "GLP", numericCode: "312"},
	"GMB": {nameEn: "Gambia (the)", nameFr: "Gambie (la)", alpha2Code: "GM", alpha3Code: "GMB", numericCode: "270"},
	"GNB": {nameEn: "Guinea-Bissau", nameFr: "Guinée-Bissau (la)", alpha2Code: "GW", alpha3Code: "GNB", numericCode: "624"},
	"GNQ": {nameEn: "Equatorial Guinea", nameFr: "Guinée équatoriale (la)", alpha2Code: "GQ", alpha3Code: "GNQ", numericCode: "226"},
	"GRC": {nameEn: "Greece", nameFr: "Grèce (la)", alpha2Code: "GR", alpha3Code: "GRC", numericCode: "300"},
	"GRD": {nameEn: "Grenada", nameFr: "Grenade (la)", alpha2Code: "GD", alpha3Code: "GRD", numericCode: "308"},
	"GRL": {nameEn: "Greenland", nameFr: "Groenland (le)", alpha2Code: "GL", alpha3Code: "GRL", numericCode: "304"},
	"GTM": {nameEn: "Guatemala", nameFr: "Guatemala (le)", alpha2Code: "GT", alpha3Code: "GTM", numericCode: "320"},
	"GUF": {nameEn: "French Guiana", nameFr: "Guyane française (la )", alpha2Code: "GF", alpha3Code: "GUF", numericCode: "254"},
	"GUM": {nameEn: "Guam", nameFr: "Guam", alpha2Code: "GU", alpha3Code: "GUM", numericCode: "316"},
	"GUY": {nameEn: "Guyana", nameFr: "Guyana (le)", alpha2Code: "GY", alpha3Code: "GUY", numericCode: "328"},
	"HKG": {nameEn: "Hong Kong", nameFr: "Hong Kong", alpha2Code: "HK", alpha3Code: "HKG", numericCode: "344"},
	"HMD": {nameEn: "Heard Island and McDonald Islands", nameFr: "Heard-et-Îles MacDonald (l'Île)", alpha2Code: "HM", alpha3Code: "HMD", numericCode: "334"},
	"HND": {nameEn: "Honduras", nameFr: "Honduras (le)", alpha2Code: "HN", alpha3Code: "HND", numericCode: "340"},
	"HRV": {nameEn: "Croatia", nameFr: "Croatie (la)", alpha2Code: "HR", alpha3Code: "HRV", numericCode: "191"},
	"HTI": {nameEn: "Haiti", nameFr: "Haïti", alpha2Code: "HT", alpha3Code: "HTI", numericCode: "332"},
	"HUN": {nameEn: "Hungary", nameFr: "Hongrie (la)", alpha2Code: "HU", alpha3Code: "HUN", numericCode: "348"},
	"IDN": {nameEn: "Indonesia", nameFr: "Indonésie (l')", alpha2Code: "ID", alpha3Code: "IDN", numericCode: "360"},
	"IMN": {nameEn: "Isle of Man", nameFr: "Île de Man", alpha2Code: "IM", alpha3Code: "IMN", numericCode: "833"},
	"IND": {nameEn: "India", nameFr: "Inde (l')", alpha2Code: "IN", alpha3Code: "IND", numericCode: "356"},
	"IOT": {nameEn: "British Indian Ocean Territory (the)", nameFr: "Indien (le Territoire britannique de l'océan)", alpha2Code: "IO", alpha3Code: "IOT", numericCode: "086"},
	"IRL": {nameEn: "Ireland", nameFr: "Irlande (l')", alpha2Code: "IE", alpha3Code: "IRL", numericCode: "372"},
	"IRN": {nameEn: "Iran (Islamic Republic of)", nameFr: "Iran (République Islamique d')", alpha2Code: "IR", alpha3Code: "IRN", numericCode: "364"},
	"IRQ": {nameEn: "Iraq", nameFr: "Iraq (l')", alpha2Code: "IQ", alpha3Code: "IRQ", numericCode: "368"},
	"ISL": {nameEn: "Iceland", nameFr: "Islande (l')", alpha2Code: "IS", alpha3Code: "ISL", numericCode: "352"},
	"ISR": {nameEn: "Israel", nameFr: "Israël", alpha2Code: "IL", alpha3Code: "ISR", numericCode: "376"},
	"ITA": {nameEn: "Italy", nameFr: "Italie (l')", alpha2Code: "IT", alpha3Code: "ITA", numericCode: "380"},
	"JAM": {nameEn: "Jamaica", nameFr: "Jamaïque (la)", alpha2Code: "JM", alpha3Code: "JAM", numericCode: "388"},
	"JEY": {nameEn: "Jersey", nameFr: "Jersey", alpha2Code: "JE", alpha3Code: "JEY", numericCode: "832"},
	"JOR": {nameEn: "Jordan", nameFr: "Jordanie (la)", alpha2Code: "JO", alpha3Code: "JOR", numericCode: "400"},
	"JPN": {nameEn: "Japan", nameFr: "Japon (le)", alpha2Code: "JP", alpha3Code: "JPN", numericCode: "392"},
	"KAZ": {nameEn: "Kazakhstan", nameFr: "Kazakhstan (le)", alpha2Code: "KZ", alpha3Code: "KAZ", numericCode: "398"},
	"KEN": {nameEn: "Kenya", nameFr: "Kenya (le)", alpha2Code: "KE", alpha3Code: "KEN", numericCode: "404"},
	"KGZ": {nameEn: "Kyrgyzstan", nameFr: "Kirghizistan (le)", alpha2Code: "KG", alpha3Code: "KGZ", numericCode: "417"},
	"KHM": {nameEn: "Cambodia", nameFr: "Cambodge (le)", alpha2Code: "KH", alpha3Code: "KHM", numericCode: "116"},
	"KIR": {nameEn: "Kiribati", nameFr: "Kiribati", alpha2Code: "KI", alpha3Code: "KIR", numericCode: "296"},
	"KNA": {nameEn: "Saint Kitts and Nevis", nameFr: "Saint-Kitts-et-Nevis", alpha2Code: "KN", alpha3Code: "KNA", numericCode: "659"},
	"KOR": {nameEn: "Korea (the Republic of)", nameFr: "Corée (la République de)", alpha2Code: "KR", alpha3Code: "KOR", numericCode: "410"},
	"KWT": {nameEn: "Kuwait", nameFr: "Koweït (le)", alpha2Code: "KW", alpha3Code: "KWT", numericCode: "414"},
	"LAO": {nameEn: "Lao People's Democratic Republic (the)", nameFr: "Lao (la République démocratique populaire)", alpha2Code: "LA", alpha3Code: "LAO", numericCode: "418"},
	"LBN": {nameEn: "Lebanon", nameFr: "Liban (le)", alpha2Code: "LB", alpha3Code: "LBN", numericCode: "422"},
	"LBR": {nameEn: "Liberia", nameFr: "Libéria (le)", alpha2Code: "LR", alpha3Code: "LBR", numericCode: "430"},
	"LBY": {nameEn: "Libya", nameFr: "Libye (la)", alpha2Code: "LY", alpha3Code: "LBY", numericCode: "434"},
	"LCA": {nameEn: "Saint Lucia", nameFr: "Sainte-Lucie", alpha2Code: "LC", alpha3Code: "LCA", numericCode: "662"},
	"LIE": {nameEn: "Liechtenstein", nameFr: "Liechtenstein (le)", alpha2Code: "LI", alpha3Code: "LIE", numericCode: "438"},
	"LKA": {nameEn: "Sri Lanka", nameFr: "Sri Lanka", alpha2Code: "LK", alpha3Code: "LKA", numericCode: "144"},
	"LSO": {nameEn: "Lesotho", nameFr: "Lesotho (le)", alpha2Code: "LS", alpha3Code: "LSO", numericCode: "426"},
	"LTU": {nameEn: "Lithuania", nameFr: "Lituanie (la)", alpha2Code: "LT", alpha3Code: "LTU", numericCode: "440"},
	"LUX": {nameEn: "Luxembourg", nameFr: "Luxembourg (le)", alpha2Code: "LU", alpha3Code: "LUX", numericCode: "442"},
	"LVA": {nameEn: "Latvia", nameFr: "Lettonie (la)", alpha2Code: "LV", alpha3Code: "LVA", numericCode: "428"},
	"MAC": {nameEn: "Macao", nameFr: "Macao", alpha2Code: "MO", alpha3Code: "MAC", numericCode: "446"},
	"MAF": {nameEn: "Saint Martin (French part)", nameFr: "Saint-Martin (partie française)", alpha2Code: "MF", alpha3Code: "MAF", numericCode: "663"},
	"MAR": {nameEn: "Morocco", nameFr: "Maroc (le)", alpha2Code: "MA", alpha3Code: "MAR", numericCode: "504"},
	"MCO": {nameEn: "Monaco", nameFr: "Monaco", alpha2Code: "MC", alpha3Code: "MCO", numericCode: "492"},
	"MDA": {nameEn: "Moldova (the Republic of)", nameFr: "Moldova (la République de)", alpha2Code: "MD", alpha3Code: "MDA", numericCode: "498"},
	"MDG": {nameEn: "Madagascar", nameFr: "Madagascar", alpha2Code: "MG", alpha3Code: "MDG", numericCode: "450"},
	"MDV": {nameEn: "Maldives", nameFr: "Maldives (les)", alpha2Code: "MV", alpha3Code: "MDV", numericCode: "462"},
	"MEX": {nameEn: "Mexico", nameFr: "Mexique (le)", alpha2Code: "MX", alpha3Code: "MEX", numericCode: "484"},
	"MHL": {nameEn: "Marshall Islands (the)", nameFr: "Marshall (les Îles)", alpha2Code: "MH", alpha3Code: "MHL", numericCode: "584"},
	"MKD": {nameEn: "North Macedonia", nameFr: "Macédoine du Nord (la)", alpha2Code: "MK", alpha3Code: "MKD", numericCode: "807"},
	"MLI": {nameEn: "Mali", nameFr: "Mali (le)", alpha2Code: "ML", alpha3Code: "MLI", numericCode: "466"},
	"MLT": {nameEn: "Malta", nameFr: "Malte", alpha2Code: "MT", alpha3Code: "MLT", numericCode: "470"},
	"MMR": {nameEn: "Myanmar", nameFr: "Myanmar (le)", alpha2Code: "MM", alpha3Code: "MMR", numericCode: "104"},
	"MNE": {nameEn: "Montenegro", nameFr: "Monténégro (le)", alpha2Code: "ME", alpha3Code: "MNE", numericCode: "499"},
	"MNG": {nameEn: "Mongolia", nameFr: "Mongolie (la)", alpha2Code: "MN", alpha3Code: "MNG", numericCode: "496"},
	"MNP": {nameEn: "Northern Mariana Islands (the)", nameFr: "Mariannes du Nord (les Îles)", alpha2Code: "MP", alpha3Code: "MNP", numericCode: "580"},
	"MOZ": {nameEn: "Mozambique", nameFr: "Mozambique (le)", alpha2Code: "MZ", alpha3Code: "MOZ", numericCode: "508"},
	"MRT": {nameEn: "Mauritania", nameFr: "Mauritanie (la)", alpha2Code: "MR", alpha3Code: "MRT", numericCode: "478"},
	"MSR": {nameEn: "Montserrat", nameFr: "Montserrat", alpha2Code: "MS", alpha3Code: "MSR", numericCode: "500"},
	"MTQ": {nameEn: "Martinique", nameFr: "Martinique (la)", alpha2Code: "MQ", alpha3Code: "MTQ", numericCode: "474"},
	"MUS": {nameEn: "Mauritius", nameFr: "Maurice", alpha2Code: "MU", alpha3Code: "MUS", numericCode: "480"},
	"MWI": {nameEn: "Malawi", nameFr: "Malawi (le)", alpha2Code: "MW", alpha3Code: "MWI", numericCode: "454"},
	"MYS": {nameEn: "Malaysia", nameFr: "Malaisie (la)", alpha2Code: "MY", alpha3Code: "MYS", numericCode: "458"},
	"MYT": {nameEn: "Mayotte", nameFr: "Mayotte", alpha2Code: "YT", alpha3Code: "MYT", numericCode: "175"},
	"NAM": {nameEn: "Namibia", nameFr: "Namibie (la)", alpha2Code: "NA", alpha3Code: "NAM", numericCode: "516"},
	"NCL": {nameEn: "New Caledonia", nameFr: "Nouvelle-Calédonie (la)", alpha2Code: "NC", alpha3Code: "NCL", numericCode: "540"},
	"NER": {nameEn: "Niger (the)", nameFr: "Niger (le)", alpha2Code: "NE", alpha3Code: "NER", numericCode: "562"},
	"NFK": {nameEn: "Norfolk Island", nameFr: "Norfolk (l'Île)", alpha2Code: "NF", alpha3Code: "NFK", numericCode: "574"},
	"NGA": {nameEn: "Nigeria", nameFr: "Nigéria (le)", alpha2Code: "NG", alpha3Code: "NGA", numericCode: "566"},
	"NIC": {nameEn: "Nicaragua", nameFr: "Nicaragua (le)", alpha2Code: "NI", alpha3Code: "NIC", numericCode: "558"},
	"NIU": {nameEn: "Niue", nameFr: "Niue", alpha2Code: "NU", alpha3Code: "NIU", numericCode: "570"},
	"NLD": {nameEn: "Netherlands (the)", nameFr: "Pays-Bas (les)", alpha2Code: "NL", alpha3Code: "NLD", numericCode: "528"},
	"NOR": {nameEn: "Norway", nameFr: "Norvège (la)", alpha2Code: "NO", alpha3Code: "NOR", numericCode: "578"},
	"NPL": {nameEn: "Nepal", nameFr: "Népal (le)", alpha2Code: "NP", alpha3Code: "NPL", numericCode: "524"},
	"NRU": {nameEn: "Nauru", nameFr: "Nauru", alpha2Code: "NR", alpha3Code: "NRU", numericCode: "520"},
	"NZL": {nameEn: "New Zealand", nameFr: "Nouvelle-Zélande (la)", alpha2Code: "NZ", alpha3Code: "NZL", numericCode: "554"},
	"OMN": {nameEn: "Oman", nameFr: "Oman", alpha2Code: "OM", alpha3Code: "OMN", numericCode: "512"},
	"PAK": {nameEn: "Pakistan", nameFr: "Pakistan (le)", alpha2Code: "PK", alpha3Code: "PAK", numericCode: "586"},
	"PAN": {nameEn: "Panama", nameFr: "Panama (le)", alpha2Code: "PA", alpha3Code: "PAN", numericCode: "591"},
	"PCN": {nameEn: "Pitcairn", nameFr: "Pitcairn", alpha2Code: "PN", alpha3Code: "PCN", numericCode: "612"},
	"PER": {nameEn: "Peru", nameFr: "Pérou (le)", alpha2Code: "PE", alpha3Code: "PER", numericCode: "604"},
	"PHL": {nameEn: "Philippines (the)", nameFr: "Philippines (les)", alpha2Code: "PH", alpha3Code: "PHL", numericCode: "608"},
	"PLW": {nameEn: "Palau", nameFr: "Palaos (les)", alpha2Code: "PW", alpha3Code: "PLW", numericCode: "585"},
	"PNG": {nameEn: "Papua New Guinea", nameFr: "Papouasie-Nouvelle-Guinée (la)", alpha2Code: "PG", alpha3Code: "PNG", numericCode: "598"},
	"POL": {nameEn: "Poland", nameFr: "Pologne (la)", alpha2Code: "PL", alpha3Code: "POL", numericCode: "616"},
	"PRI": {nameEn: "Puerto Rico", nameFr: "Porto Rico", alpha2Code: "PR", alpha3Code: "PRI", numericCode: "630"},
	"PRK": {nameEn: "Korea (the Democratic People's Republic of)", nameFr: "Corée (la République populaire démocratique de)", alpha2Code: "KP", alpha3Code: "PRK", numericCode: "408"},
	"PRT": {nameEn: "Portugal", nameFr: "Portugal (le)", alpha2Code: "PT", alpha3Code: "PRT", numericCode: "620"},
	"PRY": {nameEn: "Paraguay", nameFr: "Paraguay (le)", alpha2Code: "PY", alpha3Code: "PRY", numericCode: "600"},
	"PSE": {nameEn: "Palestine, State of", nameFr: "Palestine, État de", alpha2Code: "PS", alpha3Code: "PSE", numericCode: "275"},
	"PYF": {nameEn: "French Polynesia", nameFr: "Polynésie française (la)", alpha2Code: "PF", alpha3Code: "PYF", numericCode: "258"},
	"QAT": {nameEn: "Qatar", nameFr: "Qatar (le)", alpha2Code: "QA", alpha3Code: "QAT", numericCode: "634"},
	"REU": {nameEn: "Réunion", nameFr: "Réunion (La)", alpha2Code: "RE", alpha3Code: "REU", numericCode: "638"},
	"ROU": {nameEn: "Romania", nameFr: "Roumanie (la)", alpha2Code: "RO", alpha3Code: "ROU", numericCode: "642"},
	"RUS": {nameEn: "Russian Federation (the)", nameFr: "Russie (la Fédération de)", alpha2Code: "RU", alpha3Code: "RUS", numericCode: "643"},
	"RWA": {nameEn: "Rwanda", nameFr: "Rwanda (le)", alpha2Code: "RW", alpha3Code: "RWA", numericCode: "646"},
	"SAU": {nameEn: "Saudi Arabia", nameFr: "Arabie saoudite (l')", alpha2Code: "SA", alpha3Code: "SAU", numericCode: "682"},
	"SDN": {nameEn: "Sudan (the)", nameFr: "Soudan (le)", alpha2Code: "SD", alpha3Code: "SDN", numericCode: "729"},
	"SEN": {nameEn: "Senegal", nameFr: "Sénégal (le)", alpha2Code: "SN", alpha3Code: "SEN", numericCode: "686"},
	"SGP": {nameEn: "Singapore", nameFr: "Singapour", alpha2Code: "SG", alpha3Code: "SGP", numericCode: "702"},
	"SGS": {nameEn: "South Georgia and the South Sandwich Islands", nameFr: "Géorgie du Sud-et-les Îles Sandwich du Sud (la)", alpha2Code: "GS", alpha3Code: "SGS", numericCode: "239"},
	"SHN": {nameEn: "Saint Helena, Ascension and Tristan da Cunha", nameFr: "Sainte-Hélène, Ascension et Tristan da Cunha", alpha2Code: "SH", alpha3Code: "SHN", numericCode: "654"},
	"SJM": {nameEn: "Svalbard and Jan Mayen", nameFr: "Svalbard et l'Île Jan Mayen (le)", alpha2Code: "SJ", alpha3Code: "SJM", numericCode: "744"},
	"SLB": {nameEn: "Solomon Islands", nameFr: "Salomon (les Îles)", alpha2Code: "SB", alpha3Code: "SLB", numericCode: "090"},
	"SLE": {nameEn: "Sierra Leone", nameFr: "Sierra Leone (la)", alpha2Code: "SL", alpha3Code: "SLE", numericCode: "694"},
	"SLV": {nameEn: "El Salvador", nameFr: "El Salvador", alpha2Code: "SV", alpha3Code: "SLV", numericCode: "222"},
	"SMR": {nameEn: "San Marino", nameFr: "Saint-Marin", alpha2Code: "SM", alpha3Code: "SMR", numericCode: "674"},
	"SOM": {nameEn: "Somalia", nameFr: "Somalie (la)", alpha2Code: "SO", alpha3Code: "SOM", numericCode: "706"},
	"SPM": {nameEn: "Saint Pierre and Miquelon", nameFr: "Saint-Pierre-et-Miquelon", alpha2Code: "PM", alpha3Code: "SPM", numericCode: "666"},
	"SRB": {nameEn: "Serbia", nameFr: "Serbie (la)", alpha2Code: "RS", alpha3Code: "SRB", numericCode: "688"},
	"SSD": {nameEn: "South Sudan", nameFr: "Soudan du Sud (le)", alpha2Code: "SS", alpha3Code: "SSD", numericCode: "728"},
	"STP": {nameEn: "Sao Tome and Principe", nameFr: "Sao Tomé-et-Principe", alpha2Code: "ST", alpha3Code: "STP", numericCode: "678"},
	"SUR": {nameEn: "Suriname", nameFr: "Suriname (le)", alpha2Code: "SR", alpha3Code: "SUR", numericCode: "740"},
	"SVK": {nameEn: "Slovakia", nameFr: "Slovaquie (la)", alpha2Code: "SK", alpha3Code: "SVK", numericCode: "703"},
	"SVN": {nameEn: "Slovenia", nameFr: "Slovénie (la)", alpha2Code: "SI", alpha3Code: "SVN", numericCode: "705"},
	"SWE": {nameEn: "Sweden", nameFr: "Suède (la)", alpha2Code: "SE", alpha3Code: "SWE", numericCode: "752"},
	"SWZ": {nameEn: "Eswatini", nameFr: "Eswatini (l')", alpha2Code: "SZ", alpha3Code: "SWZ", numericCode: "748"},
	"SXM": {nameEn: "Sint Maarten (Dutch part)", nameFr: "Saint-Martin (partie néerlandaise)", alpha2Code: "SX", alpha3Code: "SXM", numericCode: "534"},
	"SYC": {nameEn: "Seychelles", nameFr: "Seychelles (les)", alpha2Code: "SC", alpha3Code: "SYC", numericCode: "690"},
	"SYR": {nameEn: "Syrian Arab Republic (the)", nameFr: "République arabe syrienne (la)", alpha2Code: "SY", alpha3Code: "SYR", numericCode: "760"},
	"TCA": {nameEn: "Turks and Caicos Islands (the)", nameFr: "Turks-et-Caïcos (les Îles)", alpha2Code: "TC", alpha3Code: "TCA", numericCode: "796"},
	"TCD": {nameEn: "Chad", nameFr: "Tchad (le)", alpha2Code: "TD", alpha3Code: "TCD", numericCode: "148"},
	"TGO": {nameEn: "Togo", nameFr: "Togo (le)", alpha2Code: "TG", alpha3Code: "TGO", numericCode: "768"},
	"THA": {nameEn: "Thailand", nameFr: "Thaïlande (la)", alpha2Code: "TH", alpha3Code: "THA", numericCode: "764"},
	"TJK": {nameEn: "Tajikistan", nameFr: "Tadjikistan (le)", alpha2Code: "TJ", alpha3Code: "TJK", numericCode: "762"},
	"TKL": {nameEn: "Tokelau", nameFr: "Tokelau (les)", alpha2Code: "TK", alpha3Code: "TKL", numericCode: "772"},
	"TKM": {nameEn: "Turkmenistan", nameFr: "Turkménistan (le)", alpha2Code: "TM", alpha3Code: "TKM", numericCode: "795"},
	"TLS": {nameEn: "Timor-Leste", nameFr: "Timor-Leste (le)", alpha2Code: "TL", alpha3Code: "TLS", numericCode: "626"},
	"TON": {nameEn: "Tonga", nameFr: "Tonga (les)", alpha2Code: "TO", alpha3Code: "TON", numericCode: "776"},
	"TTO": {nameEn: "Trinidad and Tobago", nameFr: "Trinité-et-Tobago (la)", alpha2Code: "TT", alpha3Code: "TTO", numericCode: "780"},
	"TUN": {nameEn: "Tunisia", nameFr: "Tunisie (la)", alpha2Code: "TN", alpha3Code: "TUN", numericCode: "788"},
	"TUR": {nameEn: "Türkiye", nameFr: "Türkiye (la)", alpha2Code: "TR", alpha3Code: "TUR", numericCode: "792"},
	"TUV": {nameEn: "Tuvalu", nameFr: "Tuvalu (les)", alpha2Code: "TV", alpha3Code: "TUV", numericCode: "798"},
	"TWN": {nameEn: "Taiwan (Province of China)", nameFr: "Taïwan (Province de Chine)", alpha2Code: "TW", alpha3Code: "TWN", numericCode: "158"},
	"TZA": {nameEn: "Tanzania, the United Republic of", nameFr: "Tanzanie (la République-Unie de)", alpha2Code: "TZ", alpha3Code: "TZA", numericCode: "834"},
	"UGA": {nameEn: "Uganda", nameFr: "Ouganda (l')", alpha2Code: "UG", alpha3Code: "UGA", numericCode: "800"},
	"UKR": {nameEn: "Ukraine", nameFr: "Ukraine (l')", alpha2Code: "UA", alpha3Code: "UKR", numericCode: "804"},
	"UMI": {nameEn: "United States Minor Outlying Islands (the)", nameFr: "Îles mineures éloignées des États-Unis (les)", alpha2Code: "UM", alpha3Code: "UMI", numericCode: "581"},
	"URY": {nameEn: "Uruguay", nameFr: "Uruguay (l')", alpha2Code: "UY", alpha3Code: "URY", numericCode: "858"},
	"USA": {nameEn: "United States of America (the)", nameFr: "États-Unis d'Amérique (les)", alpha2Code: "US", alpha3Code: "USA", numericCode: "840"},
	"UZB": {nameEn: "Uzbekistan", nameFr: "Ouzbékistan (l')", alpha2Code: "UZ", alpha3Code: "UZB", numericCode: "860"},
	"VAT": {nameEn: "Holy See (the)", nameFr: "Saint-Siège (le)", alpha2Code: "VA", alpha3Code: "VAT", numericCode: "336"},
	"VCT": {nameEn: "Saint Vincent and the Grenadines", nameFr: "Saint-Vincent-et-les Grenadines", alpha2Code: "VC", alpha3Code: "VCT", numericCode: "670"},
	"VEN": {nameEn: "Venezuela (Bolivarian Republic of)", nameFr: "Venezuela (République bolivarienne du)", alpha2Code: "VE", alpha3Code: "VEN", numericCode: "862"},
	"VGB": {nameEn: "Virgin Islands (British)", nameFr: "Vierges britanniques (les Îles)", alpha2Code: "VG", alpha3Code: "VGB", numericCode: "092"},
	"VIR": {nameEn: "Virgin Islands (U.S.)", nameFr: "Vierges des États-Unis (les Îles)", alpha2Code: "VI", alpha3Code: "VIR", numericCode: "850"},
	"VNM": {nameEn: "Viet Nam", nameFr: "Viet Nam (le)", alpha2Code: "VN", alpha3Code: "VNM", numericCode: "704"},
	"VUT": {nameEn: "Vanuatu", nameFr: "Vanuatu (le)", alpha2Code: "VU", alpha3Code: "VUT", numericCode: "548"},
	"WLF": {nameEn: "Wallis and Futuna", nameFr: "Wallis-et-Futuna", alpha2Code: "WF", alpha3Code: "WLF", numericCode: "876"},
	"WSM": {nameEn: "Samoa", nameFr: "Samoa (le)", alpha2Code: "WS", alpha3Code: "WSM", numericCode: "882"},
	"YEM": {nameEn: "Yemen", nameFr: "Yémen (le)", alpha2Code: "YE", alpha3Code: "YEM", numericCode: "887"},
	"ZAF": {nameEn: "South Africa", nameFr: "Afrique du Sud (l')", alpha2Code: "ZA", alpha3Code: "ZAF", numericCode: "710"},
	"ZMB": {nameEn: "Zambia", nameFr: "Zambie (la)", alpha2Code: "ZM", alpha3Code: "ZMB", numericCode: "894"},
	"ZWE": {nameEn: "Zimbabwe", nameFr: "Zimbabwe (le)", alpha2Code: "ZW", alpha3Code: "ZWE", numericCode: "716"},
}
