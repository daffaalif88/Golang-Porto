package controllers

import (
	"golang-porto/backend/pkg/models"
	"golang-porto/backend/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Function for logging in
func Login(c echo.Context) error {
	var user models.User
	var storedUser models.User

	// Bind JSON body ke struct user
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	// Ambil user dari database berdasarkan username
	if err := db.Where("Email = ?", user.Email).First(&storedUser).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Email or password"})
	}

	// Verifikasi password (dengan bcrypt atau metode hashing yang kamu gunakan)
	if err := VerifyPassword(storedUser.Password, user.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Email or password"})
	}

	// Generate JWT token
	token, err := utils.GenerateToken(storedUser.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	// Kirim token dalam respons
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
		"token":   token,
	})
}

// Function for registering a new user (for demonstration purposes)
func Register(c echo.Context) error {
	var user models.User

	// Bind JSON body ke struct user
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	// Hash password sebelum menyimpan
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}
	user.Password = hashedPassword // Simpan password yang sudah di-hash

	// Simpan data user ke database
	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	// Buat profil yang terkait dengan user yang baru dibuat
	profile := models.Profile{
		UserID: user.ID, // pastikan field ini ada di struct Profile
		// tambahkan field lain sesuai kebutuhan, seperti default avatar atau bio kosong
	}

	// Simpan profil ke database
	if err := db.Create(&profile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create profile"})
	}

	// Kirim response sukses
	return c.JSON(http.StatusCreated, map[string]string{"message": "User and profile registered successfully"})
}

func HashPassword(password string) (string, error) {
	// Menghasilkan hash dari password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword akan memverifikasi password dengan hash
func VerifyPassword(hashedPassword, password string) error {
	// Memeriksa password yang dimasukkan dengan hash
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Function for logging out
func Logout(c echo.Context) error {
	// Berikan pesan logout berhasil atau perbarui logika jika memerlukan tindakan tambahan di server
	return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
}
