class Product < ApplicationRecord
    validates  :title, presence: true
    validates  :image, presence: true
    scope :recent, -> { order(id: :desc) }
end
