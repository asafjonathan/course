require 'rails_helper'


RSpec.describe '/articles routes' do
    it 'routes to articles#index' do
        expect(get '/articles').to  route_to(controller: 'articles', action: 'index')
        #expect(get '/articles?page[number]=3').to  route_to(controller: 'articles#index',page: {number: 3})
    end
end
