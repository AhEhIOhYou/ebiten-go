package scene

var events = []func(scene *Scene){Move, Attack}

func Move(scene *Scene) {

	scene.enemy.Action(0, 5)
}

func Pause(scene Scene) {

}

func Attack(scene *Scene) {
	//scene.enemy.SetWeaponDegree(scene.enemy.GetWeaponDegree() + 12)
	scene.enemy.FireWeapon(scene.enemy.GetWeaponDegree())
}
