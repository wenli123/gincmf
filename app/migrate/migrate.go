package migrate

import "os"

type Migrate interface {
	AutoMigrate()
}

func AutoMigrate() {
	_, err := os.Stat("../conf/install.lock")
	if err != nil {
		new(Option).AutoMigrate()
		new(Slide).AutoMigrate()
		new(User).AutoMigrate()
		new(Asset).AutoMigrate()
		new(Log).AutoMigrate()
	}

}
