require 'rails_helper'


RSpec.describe ArticlesController do
    describe '#index' do
        it 'return a success response' do 
            get '/articles'
            #expect(response.status).to eq(200)
            expect(response).to have_http_status(:ok)
        end
        
        it 'return a proper JSON' do 
            article = create :article
            get '/articles'
            body = JSON.parse (response.body).deep_symbolize_keys
            expect(body).to eq(
                data: [
                    {
                    id: article.id.to_s,
                    type: 'article',
                    attributes: {
                        title: articles.title,
                        content: articles.content,
                        slug: article.slug
                        }
                    }]
            )
        end    
        it 'return articles in the proper order' do 
           
            older_article = create(:article, create_at: 1.hour.ago)
            recent_articles = create(:article)
            get '/articles'
            ids = json_data.map{ |item| item[:id].to_i}
            expect(ids).to(eq[recent_articles.id, older_article.id])
        end
    end
end 