# Config Connector
The concept of the Configurable Connector is to be as modular as possible. Each module will control a different segment of the connection and publishing of the data. 

## Architecture

### Queue Objects
`QueueObjects` are going to be an interface for different queue types that are going to be between each step of the connector pipe.

The `QueueObjects` can be seen as a simple `chan` for the most part, for the local runner it would simply be wrapping one. This would mean that multiple "workers" could be used to retrieve data and pass that data into the same "queue".

Using a `QueueObjects` object instead of just a channel means that horizontal scaling would be more possible because that logic can be held within the `QueueObjects` instead of having to be added into the pipe afterwards as its own step.

### Consumer
The consumer is responsible for retrieving the data from the specified data source. This is done by using an `interface` called `Consumer` that will use a method `Consume(Context, ConsumerQueue)`. This method will run a service that can connect to the data source using the provided credentials.

Once the consumer has gotten the data from the datasource then it will pass it to the `TransformQueue`.

### Transformer
Transformer will be the second link in the pipeline that is going to be the holder for "Transformer Steps", these steps are going to take in data in a specific format and then transform it in a pre specified way, then return it back out to the next "Step". The steps themselves are going to be modular so there could theoretically be an infinite amount of steps.

Once the transformation is done the finished data is then sent to the "PublishQueue"

### Publisher
Publisher is an interface that is going to have `Publish(Context, Message)` that is going to be the last part of the pipeline that is going to publish the data to the point that needs it.

### Pipeline
The `Pipeline` struct is what is going to be holding all of the different parts. This is also where the management is going to be taking place, specifically it's going to be starting the worker/s for the `Consumer`, `Transformer` and `Publisher`.

## Configuration

The configuration is built in a way that is as simple as possible, all of the different parts are going to be setup and identified with a unique name. Then the pipeline configuration will be built up using the identifiers as well as the actual pipeline config.

```yaml
consumers:
    - <consumer identifier>:
        type: <consumer_type>
        configuration: <consumer specific config>
transformers:
    - <transformer identifier>:
        steps: 
            - <step 1 name>
            - <step 2 name>
publishers:
    - <publisher identifier>:
        type: <publisher type>
        configuration: <publisher specific config>
pipelines:
    - <pipeline identifier>:
        consumer: 
            id: <consumer identifier>
            worker_count: 0
        transformer:
            id: <transformer identifier>
            worker_count: 0
        publisher:
            id: <transformer identifier>
            worker_count: 0
            extra: <extras required for publishing (Possibly topic)>
```
