class RabbitmqJob < ApplicationJob
  self.queue_adapter = :sidekiq
  queue_as :default
  
  def perform(*args)
    conn = Bunny.new(ENV['RABBIT_CONNECTION_STRING'])
    conn.start
    ch = conn.create_channel
    queue = ch.queue("ProductQueue", :exclusive=>false, :auto_delete=> false)
    puts "in rabbit job"
    ch.prefetch(1)
    
    begin
      queue.subscribe(manual_ack: true, block: true) do | delivery_info, _properties, body|
        product = JSON.parse(body)
        sleep body.count('.').to_i
        Product.increment_counter(:likes, product["Id"], touch: true)
        ch.ack(delivery_info.delivery_tag)
      end
    rescue => exception
      puts "exception", exception
      conn.close
    end
  end
end
