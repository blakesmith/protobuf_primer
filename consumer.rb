require 'rubygems'
require 'net/http'
require './messages/messages.pb'


uri = URI.parse("http://localhost:5555/users/joe")
response = Net::HTTP.get_response(uri)
user = Messages::User.new
user.parse_from_string(response.body)

puts "Got user: #{user.inspect}"


