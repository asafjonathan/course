class AccessToken < ApplicationRecord
  belongs_to :user
  after_initialize :generate_token

  private 

  def generate_token
    loop do 
      pp AccessToken.exists?(token: token)
      break if token.present? && AccessToken.exists?(token: token)
      self.token = SecureRandom.hex(10)
    end


  end  
end
