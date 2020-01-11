package models

// Configurations exported
type Configurations struct {
	Server		ServerConfigurations
	Database	DatabaseConfigurations
	JWT			JWTConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	host 	string
	Port 	int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBHost		string
	DBPort		int
	DBName     	string
	DBUser     	string
	DBPassword 	string
}

// JWTConfigurations exported
type JWTConfigurations struct {
	AdminSecret		string
	UserSecret		string
}