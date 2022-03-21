package v2

func (c *Client) GetDatabaseSchema(dbName string) (allEntries []string, status DBStatusType, err error) {

	//	schemaSession, err := c.client.SessionOpen(c.ctx, &pb.Session_Open_Req{
	//		Database: dbName,
	//		Type:     pb.Session_SCHEMA,
	//		Options:  &pb.Options{},
	//	})
	//	if err != nil {
	//		return allEntries, ErrorOpenSession, err
	//	}
	//
	//	query := `
	//match
	//	$x sub thing;
	//get
	//	$x;
	//`
	//
	//	schemaTransaction, err := c.client.Transaction(c.ctx)
	//	if err != nil {
	//		return allEntries, ErrorCreateTransaction, err
	//	}
	//
	//	transactionId := ksuid.New().String()

	return allEntries, SchemaReadError, nil
}
