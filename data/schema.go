package data

func GetPhoneCallsSchema() string {
	// https://github.com/vaticle/typedb-examples/blob/master/schemas/phone-calls-schema.gql
	return `define

    name sub attribute,
      value string;
    started-at sub attribute,
      value datetime;
    duration sub attribute,
      value long;
    first-name sub attribute,
      value string;
    last-name sub attribute,
      value string;
    phone-number sub attribute,
      value string;
    city sub attribute,
      value string;
    age sub attribute,
      value long;

    contract sub relation,
        relates provider,
        relates customer;

    call sub relation,
        relates caller,
        relates callee,
        owns started-at,
        owns duration;

    company sub entity,
        plays contract:provider,
        owns name;

    person sub entity,
        plays contract:customer,
        plays call:caller,
        plays call:callee,
        owns first-name,
        owns last-name,
        owns phone-number,
        owns city,
        owns age;
`
}
