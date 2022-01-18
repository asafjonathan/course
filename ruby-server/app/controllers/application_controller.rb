class ApplicationController < ActionController::API
    rescue_from UserAuthenticator::AuthenticatorError, with: :authentication_error

    private 
    
    def authentication_error 
        error = {
            "status": "401",
            "source": {"pointer": "/data/attributes/code"},
            "title": "Authentication code is invalid",
            "detail": "You must supply valid code"
        }
        render  json:{"errors": error},status:401
    end

end
