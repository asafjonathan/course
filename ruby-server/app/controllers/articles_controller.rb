class ArticlesController < ActionController::API

    def index
        render  json: serializer.new(Article.all)
    end

    def serializer 
        ArticleSerializer
    end
end