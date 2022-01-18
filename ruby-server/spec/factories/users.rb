FactoryBot.define do
  factory :user do
    sequence(:login) { |n| "jsmith #{n}"}
    name {"Asaf Rofe"}
    url {"http:/example.com"}
    avatar_url {"http://example.com/avatar"}
    provider {"github"}
  end
end
