package main

import "sync"

//func main() {
//	fmt.Println("Start")
//	go process()
//	time.Sleep(time.Millisecond * 10)
//	fmt.Println("Ready")
//}
//
//func process() {
//	fmt.Println("Processing")
//}

//func DoWork() int {
//	time.Sleep(time.Millisecond * 1000)
//	return rand.Intn(100)
//}
//
//func main() {
//	dataChan := make(chan int)
//
//	go func() {
//		wg := &sync.WaitGroup{}
//		for i := 0; i < 1000; i++ {
//			wg.Add(1)
//			go func() {
//				defer wg.Done()
//				result := DoWork()
//				dataChan <- result
//			}()
//		}
//
//		wg.Wait()
//
//		close(dataChan)
//	}()
//
//	for v := range dataChan {
//		fmt.Println(v)
//	}
//}
//
//func main() {
//	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
//	defer cancel()
//
//	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)
//	if err != nil {
//		panic(err)
//	}
//
//	res, err := http.DefaultClient.Do(req)
//	if err != nil {
//		panic(err)
//	}
//
//	defer res.Body.Close()
//
//	imageData, err := ioutil.ReadAll(res.Body)
//
//	fmt.Println("downloaded image", len(imageData))
//}
//
//func main() {
//	r := gin.Default()
//	fmt.Println(r)
//	//r.Get("/hello", func(c *gin.Context) {
//	//	fmt.Println(1)
//	//})
//}

//type BankAccount interface {
//	GetBalance() int
//	Deposit(amount int)
//	Withdraw(amount int) error
//}
//
//type PrivateBankAccount struct {
//	balance int
//}
//
//func NewPrivateBankAccount(balance int) *PrivateBankAccount {
//	return &PrivateBankAccount{balance: balance}
//}
//
//func (p *PrivateBankAccount) GetBalance() int {
//	return p.balance
//}
//
//func (p *PrivateBankAccount) Deposit(amount int) {
//	p.balance += amount
//}
//
//func (p *PrivateBankAccount) Withdraw(amount int) error {
//	if p.balance < amount {
//		return errors.New("insufficient funds")
//	}
//
//	p.balance -= amount
//	return nil
//}
//
//type MonoBankAccount struct {
//	balance int
//	fee     int
//}
//
//func NewMonoBankAccount(balance int) *MonoBankAccount {
//	return &MonoBankAccount{balance: balance, fee: 3}
//}
//
//func (m *MonoBankAccount) GetBalance() int {
//	return m.balance
//}
//
//func (m *MonoBankAccount) Deposit(amount int) {
//	m.balance += amount
//}
//
//func (m *MonoBankAccount) Withdraw(amount int) error {
//	if m.balance < amount {
//		return errors.New("insufficient funds")
//	}
//
//	m.balance -= amount + m.fee
//	return nil
//
//}
//
//func main() {
//	accounts := []BankAccount{
//		NewPrivateBankAccount(100),
//		NewMonoBankAccount(100),
//	}
//
//	for _, account := range accounts {
//		account.Deposit(100)
//		if err := account.Withdraw(50); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(account.GetBalance())
//	}
//}

type Collection struct {
	Mutex sync.Mutex
	Data  map[string]string
}

func NewCollection() *Collection {
	return &Collection{Data: make(map[string]string)}
}

func (c *Collection) Has(key string) bool {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	_, ok := c.Data[key]

	return ok
}

func (c *Collection) Add(key, value string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.Has(key) {
		return
	}

	c.Data[key] = value
}

func main() {
	collection := NewCollection()
	collection.Add("key", "value")
}
