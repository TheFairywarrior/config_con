# Config Connector
The concept of the Configurable Connector is to be as modular as possible. Each module will control a different segment of the connection and publishing of the data. 

## Setup
To setup the Configurable Connector, you will need to have [GoLang](https://golang.org/doc/install) version 1.18 or higher installed.

### Documentation
For full documentation on the configuraable connector look into the [docs](docs/README.md) folder.

To view the drawio files that contain the diagrams for the configurable connector, you will need to either install the drawio plugin for VSCode or go to [draw.io](https://app.diagrams.net/) and open the files from there.

## Architecture

### Queue Objects
`QueueObjects` are going to be an interface for different queue types that are going to be between each step of the connector pipe.

The `QueueObjects` can be seen as a simple `chan` for the most part, for the local runner it would simply be wrapping one. This would mean that multiple "workers" could be used to retrieve data and pass that data into the same "queue".

Using a `QueueObjects` object instead of just a channel means that horizontal scaling would be more possible because that logic can be held within the `QueueObjects` instead of having to be added into the pipe afterwards as its own step.

### Consumer
The consumer is responsible for retrieving the data from the specified data source. This is done by using an `interface` called `Consumer` that will use a method `Consume(context.Context, queue.Queue) error`. This method will run a service that can connect to the data source using the provided credentials.

The different types of consumers can be split into 3 types (though only one is made at the moment). Those types are "API", "Listener" and "Retriever".
* API: Is basically a GoFiber route that is going to be added to the main api server that is always run with the service. (This is the only consumer type currently setup)
* Listener: Is a service that is going to be listening for data to be pushed to it.
* Retriever: Is a service that is going to be retrieving data from a data source. Either consistently or on a schedule.

Each consumer type will can then be broken down into the specific consumers for specific data sources.

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
    - <consumer type>:
        <cosnumer type configuration>
transformers:
    transformers:
        - <transformer identifier>:
            steps: 
                - <step 1 name>
                - <step 2 name>
    steps:
        - <step type>:
            - <step configuration>
publishers:
    - <publisher type>:
        - <publisher type configurations>
        
pipelines:
    - <pipeline identifier>
      consumer: <consumer name>
      transformer: <transformer name>
      publisher: <publisher name>
      
```
