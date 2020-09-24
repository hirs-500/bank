
package types
// Money представляет собой денежную сумму в минимальных единицах (дирамы, рубль, центы, и т.д ).
type Money int64
// Category представляет собой котегорию, который быль совершён платёж (авто, аптека, ресторан, т.д.).
type Category string

// Payment представляет информацию о платеже.
type Payment struct {
	ID     int
	Amount  Money//использовали Money 
	Category  Category
	
} 