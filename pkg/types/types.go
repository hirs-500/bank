
package types
// Money представляет собой денежную сумму в минимальных единицах (дирамы, рубль, центы, и т.д ).
type Money int64

// Category представляет собой котегорию, который быль совершён платёж (авто, аптека, ресторан, т.д.).
type Category string

// Status представлет собой статус платжа.
type Status string
// Предопределённый статус платежей.
const(
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment представляет информацию о платеже.
type Payment struct {
	ID     int
	Amount  Money//использовали Money 
	Category  Category
	Status    Status
}
 