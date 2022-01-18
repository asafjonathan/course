class UserAuthenticator
    class AuthenticatorError < StandardError; end
    attr_reader :user, :code, :access_token
    def  initialize(code)
        @code = code
    end

    def perform
        raise AuthenticatorError if code.blank?
        raise AuthenticatorError if token.try(:error).present?

        
        client = Octokit::Client.new(
            client_id: '2c7ee30e11ab9d992181',
            client_secret: 'e031121a691a50f3ae510d35b14d36eeee9cf268'
        )
       
        token = client.exchange_code_for_token(code)
        if token.try(:error).present?
            raise AuthenticatorError
        else
            user_client = Octokit::Client.new(access_token: token)
            user_data = user_client.user.to_h.slice(:login, :url, :avatar_url, :name)
            User.create(user_data.merge(provider: 'github'))
        end 
    end
end