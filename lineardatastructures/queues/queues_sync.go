// Структуры данных типа "Очереди" (Queues) или типа "Синхронизированная очередь" (Synchronized queue)
package lineardatastructures

import (
	"fmt"
	"math/rand"
	"time"
)

//Синхронизированная очередь состоит из элементов, которые необходимо обработать в определенной
//последовательности. Очередь пассажиров и очереди обработки билетов являются типами синхронизированных очередей.

const (
	Message_Pass_Start = iota
	Message_Ticket_Start
	Message_Pass_End
	Message_Ticket_End
)

// Queue class
type QueueSync struct {
	WaitPass, WaitTicket            int
	PlayPass, PlayTicket            bool
	QueuePass, QueueTicket, Message chan int
}

func (queue *QueueSync) New() {
	queue.Message = make(chan int)
	queue.QueuePass = make(chan int)
	queue.QueueTicket = make(chan int)
}

func (queue *QueueSync) SelectorTypeMessage() {
	for {
		switch <-queue.Message {
		case Message_Pass_Start:
			queue.WaitPass++
		case Message_Pass_End:
			queue.PlayPass = false
		case Message_Ticket_Start:
			queue.WaitTicket++
		case Message_Ticket_End:
			queue.PlayTicket = false
		}

		if queue.WaitPass > 0 && queue.WaitTicket > 0 && !queue.PlayPass && !queue.PlayTicket {
			queue.PlayPass = true
			queue.PlayTicket = true
			queue.WaitTicket--
			queue.WaitPass--
			queue.QueuePass <- 1
			queue.QueueTicket <- 1
		}
	}
}

// StartTicketIssue starts the ticket issue
func (queue *QueueSync) StartTicketIssue() {
	queue.Message <- Message_Ticket_Start
	<-queue.QueueTicket
}

// EndTicketIssue ends the ticket issue
func (Queue *QueueSync) EndTicketIssue() {
	Queue.Message <- Message_Ticket_End
}

// StartPass ends the Pass Queue
func (Queue *QueueSync) StartPass() {
	Queue.Message <- Message_Pass_Start
	<-Queue.QueuePass
}

// EndPass ends the Pass Queue
func (Queue *QueueSync) EndPass() {
	Queue.Message <- Message_Pass_End
}

// passenger method starts and ends the pass Queue
func passenger(Queue *QueueSync) {
	//fmt.Println("starting the passenger Queue")
	for {
		// fmt.Println("starting the processing")
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		Queue.StartPass()
		fmt.Println(" Passenger starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println(" Passenger ends")
		Queue.EndPass()
	}
}

// ticketIssue starts and ends the ticket issue
func ticketIssue(Queue *QueueSync) {
	for {
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		Queue.StartTicketIssue()
		fmt.Println("Ticket Issue starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket Issue ends")
		Queue.EndTicketIssue()
	}
}

func QueuesSyncExample() {
	queueSync := &QueueSync{}
	queueSync.New()
	fmt.Println(queueSync)

	for i := 0; i < 10; i++ {
		// fmt.Println(i, "passenger in the Queue")
		go passenger(queueSync)
	}

	//close(Queue.queuePass)
	for j := 0; j < 5; j++ {
		// fmt.Println(i, "ticket issued in the Queue")
		go ticketIssue(queueSync)
	}
}
