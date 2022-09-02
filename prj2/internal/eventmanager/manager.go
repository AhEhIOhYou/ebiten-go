package eventmanager

import "time"

type EventManager struct {
	events []Event
	Time   time.Time
}

func NewEventManager() *EventManager {
	em := &EventManager{}
	return em
}

func (em *EventManager) AddEvent(ev *Event) {
	em.events = append(em.events, *ev)
}

func (em *EventManager) Execute() {
	for i := 0; i < len(em.events); i++ {
		em.execute(&em.events[i])
	}
}

func (em *EventManager) execute(event *Event) {
	timeStart := event.startTime
	timeEnd := event.startTime + event.durationTime

	if time.Since(em.Time).Seconds() > timeEnd {
		event.UpdateStatus("ended")
		return
	}

	if time.Since(em.Time).Seconds() >= timeStart {
		switch event.status {
		case "new":
			event.enemy.SetWeaponDegree(event.ws.degree)
			event.UpdateStatus("progress")
			fallthrough
		case "progress":
			if event.moveX != 0 || event.moveY != 0 {
				event.enemy.Action(event.moveX, event.moveY)
			}
			if event.isFire {
				event.SetupWeapon()
				event.SetupShot()
				event.enemy.FireWeapon(event.enemy.GetWeaponDegree(), event.enemy.GetBulletSpeed(), event.enemy.GetAdjustAngles())
			}
			return
		}
	}
}
