package eventmanager

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/actors"
)

type Event struct {
	startTime    float64
	action       string
	moveX        float64
	moveY        float64
	moveDegree   int
	durationTime float64
	isFire       bool
	player       *actors.Player
	enemy        *actors.Enemy
	ws           WeaponSetup
	ss           ShotSetup
	status       string
}

type WeaponSetup struct {
	degree      int
	cooldown    int64
	rotation    int
	bulletSpeed float64
}

type ShotSetup struct {
	angles []int
	pos    [][]float64
}

func NewEvent() *Event {
	e := &Event{}

	return e
}

func (ev *Event) getInstance() *Event {
	event := ev
	return event
}

func (ev *Event) Move(x float64, y float64) (event *Event) {
	event = ev.getInstance()
	event.moveX = x
	event.moveY = y
	return event
}

func (ev *Event) MoveDegree(degree int) (event *Event) {
	event = ev.getInstance()
	event.moveDegree = degree
	return event
}

func (ev *Event) Duration(seconds float64) (event *Event) {
	event = ev.getInstance()
	event.durationTime = seconds
	return event
}

func (ev *Event) Actor(actor interface{}) (event *Event) {
	event = ev.getInstance()
	switch t := actor.(type) {
	case *actors.Player:
		event.player = t
	case *actors.Enemy:
		event.enemy = t
	}
	return event
}

func (ev *Event) Fire(doFire bool) (event *Event) {
	event = ev.getInstance()
	if doFire {
		ev.isFire = true
	} else {
		ev.isFire = false
	}
	return event
}

func (ev *Event) Weapon(degree int, cooldown int64, rotation int, bulletSpeed float64) (event *Event) {
	event = ev.getInstance()
	event.ws.degree = degree
	event.ws.cooldown = cooldown
	event.ws.rotation = rotation
	event.ws.bulletSpeed = bulletSpeed
	return event
}

func (ev *Event) SetupWeapon() {
	ev.enemy.SetWeaponCooldown(ev.ws.cooldown)
	ev.enemy.SetWeaponDegree(ev.enemy.GetWeaponDegree() + ev.ws.rotation)
	ev.enemy.SetBulletSpeed(ev.ws.bulletSpeed)
}

func (ev *Event) ShotAngles(angles []int) (event *Event) {
	event = ev.getInstance()
	event.ss.angles = angles
	return event
}

func (ev *Event) ShotPos(pos [][]float64) (event *Event) {
	event = ev.getInstance()
	event.ss.pos = pos
	return event
}

func (ev *Event) SetupShot() {
	ev.enemy.SetAdjustAngles(ev.ss.angles)
	ev.enemy.SetAdjustPos(ev.ss.pos)
}

func (ev *Event) Action(action string) (event *Event) {
	event = ev.getInstance()
	event.action = action
	return event
}

func (ev *Event) OnTime(seconds float64) (event *Event) {
	ev.startTime = seconds
	return ev
}

func (ev *Event) UpdateStatus(status string) (event *Event) {
	ev.status = status
	return ev
}
