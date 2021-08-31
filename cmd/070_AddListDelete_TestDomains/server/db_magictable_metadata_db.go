package server

import (
	"context"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func listMagicTableColumnMetadataForjsonFile(tableType string) ([]api.MagicTableColumnMetadata, error) {
	sqlToExecute := "SELECT * FROM magictable_metadata WHERE \"TableId\" = " + tableType
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	var magicTableColumnMetadata api.MagicTableColumnMetadata
	var returnMessage []api.MagicTableColumnMetadata
	var tableId int

	for rows.Next() {
		err := rows.Scan(&magicTableColumnMetadata.ColumnHeaderName, &magicTableColumnMetadata.ColumnDataName, &magicTableColumnMetadata.ColumnDataType, &magicTableColumnMetadata.Sortable, &magicTableColumnMetadata.FormatPresentationType, &magicTableColumnMetadata.ShouldBeVisible, &tableId)
		if err != nil {
			return returnMessage, err
		}
		returnMessage = append(returnMessage, magicTableColumnMetadata)
		//fmt.Println(magicTableColumnMetadata)
	}

	return returnMessage, rows.Err()
}
