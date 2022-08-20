package shared

import "github.com/AhEhIOhYou/project2/prj2/internal/objectpool"

var (
	PlayerBullets *objectpool.Pool = objectpool.NewPool()
)
