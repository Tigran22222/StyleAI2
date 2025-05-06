package models

type User struct { ID int json:"id" gorm:"primaryKey"Email string json:"email" gorm:"unique"Password string json:"password"}

type Clothing struct { ID int json:"id" gorm:"primaryKey"Category string json:"category"Name string json:"name"ImageURL string json:"image_url"}

type UserClothing struct { ID int json:"id" gorm:"primaryKey"UserID int json:"user_id"Category string json:"category"Color string json:"color"ImageURL string json:"image_url"}

type Outfit struct { ID int json:"id" gorm:"primaryKey"UserID int json:"user_id"TopID int json:"top_id"BottomID int json:"bottom_id"AccessoryID int json:"accessory_id"IsFromWardrobe bool json:"is_from_wardrobe"}

type StylistTip struct { ID int json:"id" gorm:"primaryKey"Content string json:"content"}

type Challenge struct { ID int json:"id" gorm:"primaryKey"Title string json:"title"Description string json:"description"}

type Subscription struct { ID int json:"id" gorm:"primaryKey"UserID int json:"user_id"Active bool json:"active"}

type AIStylistRequest struct { Category string json:"category"Color string json:"color"ImageURL string json:"image_url"}

type AIStylistResponse struct { Message string json:"message"Outfit Outfit json:"outfit"} 
