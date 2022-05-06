Consumer Service—Verify that the consumer can place an order and obtain their
payment information

Order Service sends a ValidateConsumerInfo message to Consumer Service.

Consumer Service receives a ValidateConsumerInfo message, verifies the consumer can place an order, and sends a ConsumerValidated message

even if Consumer Service is down, for example, Order Service still creates orders
and responds to its clients

The createOrder() operation accesses data in numerous services.
It reads data from Consumer Service and updates data in Order Service, Kitchen
Service, and Accounting Service


Consumer Service consumes the OrderCreated event, verifies that the consumer can place the order, and publishes a ConsumerVerified event.