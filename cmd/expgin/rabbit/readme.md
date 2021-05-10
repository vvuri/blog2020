## Test with RabbitMQ

### Install Windows:
1. Install Erlang
2. Install RabbitMQ
3. Add PATH for Erlang and Rabbit
4. Run Server RabbitMQ: 
   ``` rabbitmq-server.bat -detached```
5. Run plugin for visualization  
   ``` rabbitmq-plugins enable rabbitmq_management ```
6. Add Exchanges + Queue + Binding
    http://localhost:15672/ guest/guest 
7. Goland example https://www.rabbitmq.com/tutorials/tutorial-one-go.html    
``` go get github.com/streadway/amqp```

### Run:
- PS> rabbitmq-server.bat -detached
- PS> rabbitmq-plugins enable rabbitmq_management
