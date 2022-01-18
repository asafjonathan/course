require 'rails_helper'


describe UserAuthenticator  do
    describe '#perform' do
        let (:authenticator) {described_class.new('sample_code')}
        subject { authenticator.perform } 

        context "when code is incorrect" do
            let(:error) {
                double("Sawyer::Resource", error: "bad_verfication_code")
            }
            it 'shold raise an error' do 
                authenticator = described_class.new('sample_code')
                expect{ subject }.to raise_error(
                    UserAuthenticator::AuthenticatorError
                )
                expect(authenticator.user).to be_nil

            end
        end
        
        context "when code is correct" do
            let(:user_data) do 
                {
                    login: 'asafrofe2',
                    url: 'http://exemple.com',
                    avatar_url: 'http://exemple.com/avatar',
                    name: 'Asaf Rofe'
                }
            end
            before do 
                allow_any_instance_of(Octokit::Client).to receive(:exchange_code_for_token).and_return('validaccesstoken')
                allow_any_instance_of(Octokit::Client).to receive(:user).and_return(user_data)
            end
            it 'should save the user when dose not exists' do
                expect{subject}.to change{User.count}.by(1)
            end
            it 'should reuse already register user' do
                user = create :user, user_data
                expect{subject }.not_to change{User.count}
                #expect(authenticator.user).to eq(user)
            end

            it 'should create and set users access token' do
                expect{ subject }.to change{ AccessToken.Count }.by(1)
                expect(authenticator.access_token).to be_present
            end

        end
    end
end