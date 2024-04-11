# make migrate_create name=create_table_first
migrate_create:
	migrate create -ext sql -dir db/migrations $(name)
migrate_up:
	migrate -database 'mysql://root@tcp(localhost:3306)/belajar_golang_database_migration' -path ./db/migrations up
migrate_down: migrate_down_1
migrate_down_1:
	echo 'y' | migrate -database 'mysql://root@tcp(localhost:3306)/belajar_golang_database_migration' -path ./db/migrations down 1
migrate_down_all:
	migrate -database 'mysql://root@tcp(localhost:3306)/belajar_golang_database_migration' -path ./db/migrations down
migrate_version:
	migrate -database 'mysql://root@tcp(localhost:3306)/belajar_golang_database_migration' -path ./db/migrations version
migrate_force:
	migrate -database 'mysql://root@tcp(localhost:3306)/belajar_golang_database_migration' -path ./db/migrations force $(version)