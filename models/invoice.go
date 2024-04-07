package models

import "time"

type Factura struct {
	ID            int       `json:"id"`
	ClienteNombre string    `json:"cliente_nombre"`
	ClienteEmail  string    `json:"cliente_email"`
	Fecha         time.Time `json:"fecha"`
	Total         float64   `json:"total"`
	ClienteID     int       `json:"cliente_id"`
}

type DetalleFactura struct {
	ID             int     `json:"id"`
	FacturaID      int     `json:"factura_id"`
	ProductoID     int     `json:"producto_id,omitempty"`   // El omitempty evita que se muestre si es cero
	CortePeloID    int     `json:"corte_pelo_id,omitempty"` // Igual que en producto_id
	Cantidad       int     `json:"cantidad"`
	PrecioUnitario float64 `json:"precio_unitario"`
}
