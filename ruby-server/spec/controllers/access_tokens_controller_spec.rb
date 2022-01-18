require 'rails_helper'

RSpec.describe AccessTokensController, type: :controller do

    describe '#create' do
        context "when invalid request" do
          
        let(:error) do 
            {
                "status": "401",
                "source": {"pointer": "/data/attributes/code"},
                "title": "Authentication code is invalid",
                "detail": "You must supply valid code"
            }
        end
            it 'should return 401 status code' do
               post :create
               expect(response).to have_http_status(401)           
            end

            it 'should return proper error body' do
                post :create
                json = JSON.parse(response.body)
                pp json['errors']
                expect(json['errors']).to include(error)  
            end
        end

        context "when success request" do
            
        end
        
        
    end
end
