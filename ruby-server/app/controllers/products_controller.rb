class ProductsController < ActionController::API

    def index
    

        render  json: {products: Product.all }
    end
end
