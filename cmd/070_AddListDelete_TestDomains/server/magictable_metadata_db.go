package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (server *Server) listMagicTableColumnMetadataForjsonFile(tableTypeGuid string) ([]api.MagicTableColumnMetadata, error) {

	server.logger.WithFields(logrus.Fields{
		"Id":    "4ff6f958-259d-4d58-91b4-2c79ceb871d3",
		"Trace": server.trace(false),
	}).Debug("Entering: listMagicTableColumnMetadataForjsonFile()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "d943c3d1-e22a-469e-b16a-b1e4b70f1dc0",
			"Trace": server.trace(false),
		}).Debug("Exiting: listMagicTableColumnMetadataForjsonFile()")
	}()

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
