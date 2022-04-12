package data

func GetQuerySinglePerson() string {
	return `
	match $p isa person, has phone-number "+54 398 559 0423";
`
}

func GetQueryDeletePerson() string {
	return `
	match $p isa person, has phone-number "+54 398 559 0423";
	delete $p isa person;
`
}

func GetQueryAllPeople() string {
	return `
	match
		$x isa person;
	get
		$x;
	`
}

func GetQueryAllPhoneNumbers() string {
	return `
	match
	  $p isa phone-number;
	get
		$p;
	`
}

func GetQueryPersonWithPhone() string {
	return `
	match
	  $p isa person,
	  has first-name $fName,	
	  has last-name $lName,
	  has phone-number $number;
	get
		$fName,
		$lName,
		$number;
	`
}

func GetQuerySept14() string {
	// Get me all the customers of company “Telecom” who called the target person with phone number +86 921 547 9004
	// from September 14th onwards.
	return `
	match
	  $customer isa person, has phone-number $phone-number;
	  $company isa company, has name "Telecom";
	  (customer: $customer, provider: $company) isa contract;
	  $target isa person, has phone-number "+86 921 547 9004";
	  (caller: $customer, callee: $target) isa call, has started-at $started-at;
	  $min-date 2018-09-14T17:18:49; $started-at > $min-date;
	get $phone-number;
`
}

func GetQueryLondon() string {
	// Who are the people who have received a call from a London customer aged over 50
	// who has previously called someone aged under 20?
	return `
	match
	  $suspect isa person, has city "London", has age > 50;
	  $company isa company, has name "Telecom";
	  (customer: $suspect, provider: $company) isa contract;
	  $pattern-callee isa person, has age < 20;
	  (caller: $suspect, callee: $pattern-callee) isa call, has started-at $pattern-call-date;
	  $target isa person, has phone-number $phone-number, has is-customer false;
	  (caller: $suspect, callee: $target) isa call, has started-at $target-call-date;
	  $target-call-date > $pattern-call-date;
	get $phone-number;
`
}

func GetQueryCommonContacts() string {
	// Who are the common contacts of customers X and Y?
	return `
	match
	  $common-contact isa person, has phone-number $phone-number;
	  $customer-a isa person, has phone-number "+7 171 898 0853";
	  $customer-b isa person, has phone-number "+370 351 224 5176";
	  (caller: $customer-a, callee: $common-contact) isa call;
	  (caller: $customer-b, callee: $common-contact) isa call;
	get $phone-number;
`
}

func GetQueryLead() string {
	// Who are the customers who
	// 1) have all called each other and
	// 2) have all called person X at least once?
	return `
	match
	  $target isa person, has phone-number "+48 894 777 5173";
	  $company isa company, has name "Telecom";
	  $customer-a isa person, has phone-number $phone-number-a;
	  $customer-b isa person, has phone-number $phone-number-b;
	  (customer: $customer-a, provider: $company) isa contract;
	  (customer: $customer-b, provider: $company) isa contract;
	  (caller: $customer-a, callee: $customer-b) isa call;
	  (caller: $customer-a, callee: $target) isa call;
	  (caller: $customer-b, callee: $target) isa call;
	get $phone-number-a, $phone-number-b;
`
}

func GetQueryCallDurationUnder20() string {
	// Get me the average call duration among customers who have a contract with company “Telecom” and are aged under 20.
	return `
	match
	  $customer isa person, has age < 20;
	  $company isa company, has name "Telecom";
	  (customer: $customer, provider: $company) isa contract;
	  (caller: $customer, callee: $anyone) isa call, has duration $duration;
	get $duration; mean $duration;
`
}

func GetQueryCallDurationOver40() string {
	// Get me the average call duration among customers who have a contract with company “Telecom” and are aged under 20.
	return `
	match
	  $customer isa person, has age > 40;
	  $company isa company, has name "Telecom";
	  (customer: $customer, provider: $company) isa contract;
	  (caller: $customer, callee: $anyone) isa call, has duration $duration;
	get $duration; mean $duration;
`
}
