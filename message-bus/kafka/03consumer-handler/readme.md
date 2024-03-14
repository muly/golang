
using consumer group
    - to process the messages
    - marking the processed messages
    - restarting processing the messages by changing the group name


very good kafka consumer group ref: 
    - https://ednsquare.com/story/how-to-create-consumer-group-consumer-in-golang-with-kafka------JrFWqc
    - https://stackoverflow.com/questions/41986674/how-to-create-a-kafka-consumer-group-in-golang




setup: 
    - kafka: for setup see instructions in /message-bus/kafka/hello-kafka/notes file
    - topic: after creating the kafka server, to create the topic, see the same notes file
    - producer: for producer use /message-bus/kafka/hello-kafka/producer.go code