## To generate
rails new store --webpack
cd store
rails generate scaffold Item name:text quantity:integer description:text
bundle exec rails db:migrate
bundle exec rails s
