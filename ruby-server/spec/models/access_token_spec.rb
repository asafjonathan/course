require 'rails_helper'

RSpec.describe AccessToken, type: :model do
  describe '#new' do
    it 'should have a token present after init' do
      expect(AccessToken.new.token).to be_present
    end
    
  end
end
