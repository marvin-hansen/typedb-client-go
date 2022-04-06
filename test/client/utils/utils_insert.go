package utils

func GetCompanyInsert() string {
	return `
	insert $company isa company, has name "Telecom";
`
}

func GetPersonInsert() string {
	return `
insert
	$_ isa person, has first-name "Melli", has last-name "Winchcum", has city "London", has age 55, has phone-number "+7 171 898 0853";

	$_ isa person, has first-name "Celinda", has last-name "Bonick", has city "London", has age 52, has phone-number "+370 351 224 5176";

	$_ isa person, has first-name "Chryste", has last-name "Lilywhite", has city "London", has age 66, has phone-number "+81 308 988 7153";

	$_ isa person, has first-name "D'arcy", has last-name "Byfford", has city "London", has age 19, has phone-number "+54 398 559 0423";

	$_ isa person, has first-name "Xylina", has last-name "D'Alesco", has city "Cambridge", has age 51, has phone-number "+7 690 597 4443";

	$_ isa person, has first-name "Roldan", has last-name "Cometti", has city "Oxford", has age 59, has phone-number "+263 498 495 0617";

	$_ isa person, has first-name "Cob", has last-name "Lafflin", has city "Cambridge", has age 56, has phone-number "+63 815 962 6097";

	$_ isa person, has first-name "Olag", has last-name "Heakey", has city "London", has age 45, has phone-number "+81 746 154 2598";

	$_ isa person, has first-name "Mandie", has last-name "Assender", has city "London", has age 18, has phone-number "+261 860 539 4754";

	$_ isa person, has first-name "Elenore", has last-name "Stokey", has city "Oxford", has age 35, has phone-number "+62 107 530 7500";

	$_ isa person, has phone-number "+86 921 547 9004";

	$_ isa person, has phone-number "+48 894 777 5173";

	$_ isa person, has phone-number "+86 922 760 0418";

	$_ isa person, has phone-number "+33 614 339 0298";

	$_ isa person, has phone-number "+30 419 575 7546";

	$_ isa person, has phone-number "+7 414 625 3019";

	$_ isa person, has phone-number "+57 629 420 5680";

	$_ isa person, has phone-number "+351 515 605 7915";

	$_ isa person, has phone-number "+36 318 105 5629";

	$_ isa person, has phone-number "+63 808 497 1769";

	$_ isa person, has phone-number "+62 533 266 3426";

	$_ isa person, has phone-number "+351 272 414 6570";

	$_ isa person, has phone-number "+86 825 153 5518";

	$_ isa person, has phone-number "+86 202 257 8619";

	$_ isa person, has phone-number "+27 117 258 4149";

	$_ isa person, has phone-number "+48 697 447 6933";

	$_ isa person, has phone-number "+48 195 624 2025";

	$_ isa person, has phone-number "+1 254 875 4647";

	$_ isa person, has phone-number "+7 552 196 4096";

	$_ isa person, has phone-number "+86 892 682 0628";
`
}

func GetContractInsert() string {
	return `
match 
	$company isa company, 
		has name "Telecom"; 
	$customer isa person, 
		has phone-number "+7 171 898 0853";
insert 
	(provider: $company, customer: $customer) isa contract;


match 
	$company isa company, 
		has name "Telecom"; 
	$customer isa person, 
		has phone-number "+370 351 224 5176";
insert 
	(provider: $company, customer: $customer) isa contract;


match 
	$company isa company, 
		has name "Telecom"; 
	$customer isa person, 
		has phone-number "+81 308 988 7153";
insert 
	(provider: $company, customer: $customer) isa contract;


match 
	$company isa company, 
		has name "Telecom"; 
	$customer isa person, 
		has phone-number "+54 398 559 0423";
insert 
	(provider: $company, customer: $customer) isa contract;


match 
	$company isa company, 
		has name "Telecom"; 
	$customer isa person, 
		has phone-number "+7 690 597 4443";
insert 
	(provider: $company, customer: $customer) isa contract;


match 
	$company isa company, 
		has name "Telecom"; 
	$customer isa person, 
		has phone-number "+263 498 495 0617";
insert 
	(provider: $company, customer: $customer) isa contract;


match 
	$company isa company, 
		has name "Telecom"; 
	$customer isa person, 
		has phone-number "+63 815 962 6097";
insert 
	(provider: $company, customer: $customer) isa contract;


match 
	$company isa company, 
		has name "Telecom"; 
	$customer isa person, 
		has phone-number "+81 746 154 2598";
insert 
	(provider: $company, customer: $customer) isa contract;


match 
	$company isa company, 
		has name "Telecom"; 
	$customer isa person, 
		has phone-number "+261 860 539 4754";
insert 
	(provider: $company, customer: $customer) isa contract;


match 
	$company isa company, 
		has name "Telecom"; 
	$customer isa person, 
		has phone-number "+62 107 530 7500";
insert 
	(provider: $company, customer: $customer) isa contract;

`
}
