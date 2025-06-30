package main

import (
	"envocc-service-go/config"
	"envocc-service-go/entities"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s", conf.Database.Host, conf.Database.User, conf.Database.Password, conf.Database.DBName, conf.Database.Port, conf.Database.TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.Migrator().DropTable(entities.GetAllModels()...)

	if err := db.AutoMigrate(entities.GetAllModels()...); err != nil {
		panic("failed to migrate database")
	}

	if err := seedTables(db); err != nil {
		log.Fatal("Failed to seed tables: ", err)
	}

	fmt.Println("Database reset + seed complete")
}

func seedTables(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	groups := []entities.Group{
		{Name: "ผู้บริหาร", ID: 1},
		{Name: "กลุ่มบริหารทั่วไป", ID: 2},
		{Name: "กลุ่มพัฒนาองค์กร", ID: 3},
		{Name: "กลุ่มยุทธศาสตร์", ID: 4},
		{Name: "กลุ่มอาชีวอนามัย", ID: 5},
		{Name: "กลุ่มกฎหมาย", ID: 6},
		{Name: "กลุ่มเฝ้าระวังและตอบโต้ภาวะฉุกเฉิน", ID: 7},
		{Name: "กลุ่มเวชศาสตร์สิ่งแวดล้อม", ID: 8},
		{Name: "กลุ่มสื่อสารความเสี่ยงและความรอบรู้สุขภาพ", ID: 9},
		{Name: "ศูนย์พัฒนาวิชาการอาชีวอนามัยและสิ่งแวดล้อม จ.ระยอง", ID: 10},
		{Name: "งานสุขศาสตร์อุตสาหกรรม", ID: 11},
	}

	for _, group := range groups {
		if err := tx.Where("name = ?", group.Name).FirstOrCreate(&group).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to seed group: %v", err)
		}
	}
	fmt.Println("group seeded")

	positions := []entities.Position{
		{ID: 1, Name: "ผู้อำนวยการกองโรคจากการประกอบอาชีพและสิ่งแวดล้อม"},
		{ID: 2, Name: "นายแพทย์"},
		{ID: 3, Name: "นักวิชาการด้านอาชีวอนามัยและเวชศาสตร์สิ่งแวดล้อม"},
		{ID: 4, Name: "นักวิชาการด้านงานวิชาการและงานต่างประเทศ"},
		{ID: 5, Name: "นักวิชาการพัสดุ"},
		{ID: 6, Name: "นักจัดการงานทั่วไป"},
		{ID: 7, Name: "นักวิชาการเงินและบัญชี"},
		{ID: 8, Name: "นักวิชาการเผยแพร่"},
		{ID: 9, Name: "นักวิเคราะห์นโยบายและแผน"},
		{ID: 10, Name: "นักวิชาการสาธารณสุข"},
		{ID: 11, Name: "นักวิทยาศาสตร์การแพทย์"},
		{ID: 12, Name: "นักทรัพยากรบุคคล"},
		{ID: 13, Name: "นักวิชาการคอมพิวเตอร์"},
		{ID: 14, Name: "เจ้าหน้าที่พัสดุ"},
		{ID: 15, Name: "เจ้าหน้าที่บัญชีและการเงิน"},
		{ID: 16, Name: "เจ้าพนักงานคอมพิวเตอร์"},
		{ID: 17, Name: "เจ้าพนักงานธุรการ"},
		{ID: 18, Name: "เจ้าพนักงานพัสดุ"},
		{ID: 19, Name: "พนักงานพัสดุ ส.3"},
		{ID: 20, Name: "พนักงานขับรถยนต์"},
		{ID: 21, Name: "นิติกร"},
		{ID: 22, Name: "นายช่างศิลป์"},
		{ID: 23, Name: "เด็กฝึกงาน"},
	}

	for _, position := range positions {
		if err := tx.Where("name = ?", position.Name).FirstOrCreate(&position).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to seed position: %v", err)
		}
	}
	fmt.Println("Positions seeded")

	// Seed Position Levels
	positionLevels := []entities.PositionLevel{
		{ID: 1, Name: "ปฏิบัติการ"},
		{ID: 2, Name: "ชำนาญการ"},
		{ID: 3, Name: "ชำนาญการพิเศษ"},
		{ID: 4, Name: "เชี่ยวชาญ"},
		{ID: 5, Name: "ทรงคุณวุฒิ"},
		{ID: 6, Name: "ไม่มี"}}

	for _, level := range positionLevels {
		if err := tx.Where("name = ?", level.Name).FirstOrCreate(&level).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to seed position level: %v", err)
		}
	}
	fmt.Println("Position levels seeded")

	// Seed Rooms
	rooms := []entities.Room{
		{
			ID:           1,
			Name:         "ห้องประชุม 1 กองโรคจากการประกอบอาชีพและสิ่งแวดล้อม",
			Capacity:     30,
			HasEquipment: true,
			ImageURL:     "assets/test.jpeg",
		},
		{
			ID:           2,
			Name:         "ห้องประชุม 2 (ห้องเล็ก)",
			Capacity:     5,
			HasEquipment: false,
			ImageURL:     "assets/test2.jpeg",
		},
		{
			ID:           3,
			Name:         "ห้องประชุม 3 (Free Setting)",
			Capacity:     10,
			HasEquipment: false,
			ImageURL:     "assets/test3.jpeg",
		},
	}

	for _, room := range rooms {
		if err := tx.Where("name = ?", room.Name).FirstOrCreate(&room).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to seed room: %v", err)
		}
	}
	fmt.Println("Rooms seeded")
	// Seed User
	user := entities.User{
		ID:              1,
		Username:        "test",
		Password:        "$2a$10$jH.LBbrJyelezFxLvKaEXuEKp.WEiYz/.h.VwabMrmiN3HKQffun2",
		Email:           "test@mail.com",
		Prefix:          "นาย",
		FnameTH:         "ทดสอบ",
		LnameTH:         "ทดสอบ",
		FnameEn:         "test",
		LnameEn:         "test",
		Phone:           "0000000000",
		Line:            "@test",
		GroupID:         1,
		PositionID:      1,
		PositionLevelID: 1,
		AvatarID:        "test.jpg",
	}

	if err := tx.Where("username = ?", user.Username).FirstOrCreate(&user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to seed user: %v", err)
	}
	fmt.Println("User seeded")

	// Commit transaction หากทุกอย่างสำเร็จ
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	fmt.Println("Seeding complete")

	return nil
}
