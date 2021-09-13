package server

import (
	"context"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func listMagicTableColumnMetadataForjsonFile(tableTypeGuid string) ([]api.MagicTableColumnMetadata, error) {

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT magictable_metadata.\"ColumnHeaderName\", magictable_metadata.\"ColumnDataName\", magictable_metadata.\"ColumnDataType\", magictable_metadata.\"Sortable\", magictable_metadata.\"FormatPresentationType\", magictable_metadata.\"ShouldBeVisible\" "
	sqlToExecute = sqlToExecute + "FROM magictable_metadata, tabletoedit "
	sqlToExecute = sqlToExecute + "WHERE magictable_metadata.\"TableId\" = tabletoedit.\"id\" "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "tabletoedit.guid = '" + tableTypeGuid + "' "
	sqlToExecute = sqlToExecute + "ORDER BY magictable_metadata.\"PresentationOrder\""

	//fmt.Println(sqlToExecute)
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	var magicTableColumnMetadata api.MagicTableColumnMetadata
	var returnMessage []api.MagicTableColumnMetadata

	for rows.Next() {
		err := rows.Scan(&magicTableColumnMetadata.ColumnHeaderName, &magicTableColumnMetadata.ColumnDataName, &magicTableColumnMetadata.ColumnDataType, &magicTableColumnMetadata.Sortable, &magicTableColumnMetadata.FormatPresentationType, &magicTableColumnMetadata.ShouldBeVisible)
		if err != nil {
			return returnMessage, err
		}
		returnMessage = append(returnMessage, magicTableColumnMetadata)
		//fmt.Println(magicTableColumnMetadata)
	}

	return returnMessage, rows.Err()
}
