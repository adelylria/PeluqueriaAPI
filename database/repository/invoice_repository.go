package repository

import (
	"github.com/adelylria/PeluqueriaAPI/database"
	"github.com/adelylria/PeluqueriaAPI/models"
)

func ObtenerFacturasCliente(clienteID int) ([]models.Factura, error) {
	var facturas []models.Factura
	db := database.GetDB()

	query := "SELECT id, cliente_nombre, cliente_email, fecha, total FROM factura WHERE cliente_id = ?"

	rows, err := db.Query(query, clienteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var factura models.Factura
		err := rows.Scan(&factura.ID, &factura.ClienteNombre, &factura.ClienteEmail, &factura.Fecha, &factura.Total)
		if err != nil {
			return nil, err
		}
		facturas = append(facturas, factura)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return facturas, nil
}
