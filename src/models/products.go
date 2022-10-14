package models

import "time"

type Product struct {
	ID          				uint           `json:"id" gorm:"primaryKey"`
	Name        				string         `json:"name"`
	Ref  								string         `json:"ref"`
	Category  					string         `json:"category"`
	Brand   						string     		 `json:"brand"`
	Min_Quantity   			int      			 `json:"min_quantity"`
	Available_Quantity  int      			 `json:"available_quantity"`
	Purchase_Price   		float32      	 `json:"purchase_price"`
	Sale_Price   				float32      	 `json:"sale_price"`
	Benefit   					float32      	 `json:"benefit"`
	Description   			string         `json:"description"`
	CreatedAt   				time.Time      `json:"created"`
	UpdatedAt   				time.Time      `json:"updated"`
}