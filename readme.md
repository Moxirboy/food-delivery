# Food Delivery

## Overview

This project integrates several technologies, including Firebase, Docker, Reflex, PostgreSQL, Casbin, Redis, and JWT, to build a scalable and efficient application. The setup allows for real-time updates, secure data management, and easy deployment across different environments.

## Technologies Used

### Firebase

- **Description**: Firebase is a platform developed by Google for creating mobile and web applications.
- **Use in Project**: Firebase is used for storing images and retrieving them as needed. It provides robust cloud storage capabilities that ensure fast and reliable access to image data.

### Docker

- **Description**: Docker is an open-source platform that automates the deployment of applications inside lightweight, portable containers.
- **Use in Project**: Docker is used to containerize the application and its dependencies, ensuring consistency across different environments and simplifying deployment processes. The project includes Dockerfiles and docker-compose configurations for managing multi-container setups.

### Reflex

- **Description**: Reflex is a tool that keeps Docker containers alive by automatically restarting them after code changes.
- **Use in Project**: Reflex is used to enable real-time updates. After making changes to the code, Reflex ensures that these changes are immediately reflected in the running Docker containers without manual intervention.

### PostgreSQL

- **Description**: PostgreSQL is a powerful, open-source object-relational database system with a strong reputation for reliability, feature robustness, and performance.
- **Use in Project**: PostgreSQL is used as the primary relational database for managing application data, providing robust data management capabilities.

### Casbin

- **Description**: Casbin is an authorization library that supports access control models like ACL, RBAC, and ABAC.
- **Use in Project**: Casbin is used for managing policies within the application, ensuring that user access is controlled and secure according to predefined rules.

### Redis

- **Description**: Redis is an open-source, in-memory data structure store used as a database, cache, and message broker.
- **Use in Project**: Redis is utilized for caching tokens, improving the performance of token validation by storing them in-memory for quick access.

### JWT (JSON Web Tokens)

- **Description**: JWT is an open standard for securely transmitting information between parties as a JSON object.
- **Use in Project**: JWT is used for handling access tokens, enabling secure and stateless user authentication within the application.

## Utils

- **Pagination**: Implements pagination for efficiently managing large sets of data.
- **Validator**: Ensures that JSON parameters are valid and clean before being processed by the application, helping to maintain data integrity and security.

## Setup and Installation

### Prerequisites

- Docker: Ensure Docker is installed and running on your machine.
- Docker Compose: Required for managing multi-container Docker applications.

### Getting Started
#### Create .env file from .env.example before running the project

### Run service
```
 make start
```

### Run migrations
```
 make migration-up
```

### Run create migration file
```
 make migration-generate name=foobar
```

### Run stop containers
```
make stop
```

### Run remove containers
```
make rm
```
## Troubleshooting

- **Common Issues**:
    - Ensure Docker and Docker Compose are running.
    - Verify that all services are up and running.
    - Check configuration files for any errors.
    - Review application logs for detailed error messages.


## Functional Requirements for Food Delivery Service Backend:

### 1. User and Courier Authorization:
- **User Authentication**:
    - Users and couriers can authenticate into the system using email/password
    - Implement secure authentication methods such as JWT tokens.
- **Authorization**:
    - Users are authorized to:
        - View product listings.
        - Place orders for products.
        - View their order history.
    - Couriers are authorized to:
        - Access recommended orders.
        - Update order status upon delivery.

### 2. Product Ordering:
- **Product Listings**:
    - Users can view a list of available products with details such as name, description, price, and image.
    - Implement pagination for product listings if there are many items.
- **Order Placement**:
    - Users can place orders for products by adding them to a cart.
    - Users can specify quantity and options (e.g., size, toppings).
    - Provide a checkout process for users to review their orders before placing them.
- **Order Recommendations**:
    - Automatically recommend orders to couriers based on location and delivery schedule.
    - Couriers receive notifications for new recommended orders.

### 3. Courier Delivery Status:
- **Order Assignment**:
    - Upon recommendation, assign orders to available couriers automatically.
    - Provide an interface for admins to manually assign orders if needed.
- **Order Status Updates**:
    - Couriers can update the status of assigned orders:
        - Order picked up.
        - En route to delivery.
        - Delivered.
        - Payment collected (if applicable).
    - Status updates trigger notifications to users and admins.
- **Real-time Tracking**:
    - Users can track the real-time status of their orders (confirmed, picked up, out for delivery, delivered).
    - Use GPS data and mapping services to display courier location and estimated delivery time.
- **Order History**:
    - Couriers have access to their delivery history, including completed orders and earnings.

### 4. Notifications and Alerts:
- **User Notifications**:
    - Users receive notifications for:
        - Order confirmation.
        - Order status updates (picked up, on the way, delivered).
    - Notify users of estimated delivery time.
- **Courier Notifications**:
    - Couriers receive notifications for:
        - New recommended orders.
        - Order assignment.
        - Order status updates.
- **Admin Alerts**:
    - Admins receive alerts for:
        - New orders.
        - Order status changes.
        - Courier activities.

These detailed functional requirements provide a comprehensive overview of the backend functionalities for a Food Delivery Service application. It covers user and courier authorization, product ordering, order management, courier delivery status, notifications, admin panel features, safety measures, and multi-language support. Adjustments can be made based on specific project needs and scope.


## Technical Requirements:
### 1. OS:
- The backend system should be developed and tested to run on Linux-based operating systems such as Ubuntu or CentOS.
- Docker containers should be compatible with Linux environments for deployment.

### 2. Networking:
- Implement secure communication protocols (SSL/TLS) for all network interactions.
- The backend should handle HTTP/HTTPS requests efficiently.
- Handle network timeouts and retries gracefully for reliable communication.

### 3. Version Control:
- Utilize Git as the version control system for code management.
- Follow Git best practices for branching, merging, and commit messages.
- Maintain a Git repository with clear commit history and meaningful comments.

### 4. DSA (Data Structures and Algorithms):
- Use appropriate data structures and algorithms for efficient order assignment and retrieval.
- Implement algorithms for order recommendations based on courier availability and proximity.

### 5. Programming Language:
- Use Golang (Go) as the primary programming language for backend development.
- Leverage Go's concurrency features for handling multiple requests efficiently.

### 6. Databases:
- Choose a scalable and performant database like PostgreSQL or MongoDB for storing user data, product information, orders, and courier details.
- Implement proper indexing for fast retrieval of data.
- Use database transactions for maintaining data integrity during order processing.

### 7. File Storage:
- Store user profile pictures, product images, and other media files in a scalable file storage system such as Amazon S3 or Google Cloud Storage.
- Implement secure file upload and retrieval mechanisms.
- Ensure encryption of sensitive files at rest.

### 8. Message Broker:
- Integrate a message broker system like RabbitMQ or Kafka for asynchronous communication.
- Use message queues for processing order updates, notifications, and system alerts.
- Implement pub/sub patterns for real-time updates to users and couriers.

### 9. Documentation:
- Maintain detailed technical documentation using Markdown or reStructuredText.
- Document API endpoints, data models, database schema, and system architecture.
- Include clear instructions for setting up the development environment and deploying the backend.

### 10. Architecture:
- Implement a microservices architecture to decouple functionalities like user management, order processing, and notifications.
- Use Docker containers for each microservice for scalability and portability.
- Ensure modularity and maintainability of the codebase.

### 11. System Design:
- Design scalable and fault-tolerant systems to handle high traffic and concurrent requests.
- Consider load balancing strategies for distributing traffic across multiple instances of microservices.
- Implement database sharding for horizontal scalability.

### 12. Design Patterns:
- Implement design patterns such as Singleton, Factory, Observer, and Strategy as appropriate.
- Use the repository pattern for data access to separate business logic from data storage.

### 13. DevOps:
- Use Docker for containerization of the application.
- Implement CI/CD pipelines using tools like Jenkins or GitLab CI for automated testing and deployment.
- Automate deployment processes for staging and production environments using configuration management tools like Ansible or Chef.

### 14. Cloud and Infrastructure:
- Deploy the backend on cloud platforms like AWS, Azure, or Google Cloud for scalability and reliability.
- Utilize infrastructure as code (IaC) tools like Terraform or AWS CloudFormation for provisioning and managing cloud resources.
- Set up virtual private clouds (VPCs), security groups, and network ACLs for security.

### 15. API Design:
- Design clear and consistent RESTful API endpoints for user interactions (e.g., placing orders, updating delivery status).
- Use standard HTTP methods (GET, POST, PUT, DELETE) for CRUD operations.
- Implement token-based authentication (JWT tokens) for API security.

### 16. Testing:
- Write unit tests for backend services using testing frameworks like Golang's built-in testing package or testify.
- Implement integration tests to ensure different components of the backend work together correctly.
- Use tools like Postman or Swagger for API testing and documentation.

### 17. SDLC (Software Development Lifecycle):
- Follow an agile development methodology with regular sprints.
- Use project management tools like Jira or Trello for task tracking and collaboration.
- Conduct regular code reviews to maintain code quality and consistency.

### 18. Mapping:
- Implement geolocation features for tracking courier locations and providing estimated delivery times.
- Use mapping libraries like Leaflet or Google Maps for displaying maps and route optimizations.

### 19. Caching:
- Utilize caching mechanisms like Redis for caching frequently accessed data such as product listings, user profiles, and order details.
- Implement cache invalidation strategies to keep data consistent and up to date.

### 20. Data Pipeline:
- Design a data pipeline for processing and analyzing order data, user interactions, and system events.
- Use tools like Apache Kafka for real-time data streaming and processing.

### 21. Performance:
- Optimize database queries with proper indexing and query optimization techniques.
- Use CDN (Content Delivery Network) for serving static assets to improve load times.
- Implement rate limiting and throttling to prevent abuse and ensure fair usage of the backend API.
- Write load tests using tools like k6 to ensure the backend can handle expected traffic loads.

### 22. Security:
- Implement secure password hashing (bcrypt) for user authentication and authorization.
- Sanitize user inputs to prevent SQL injection and XSS attacks.
- Use HTTPS for secure communication between clients and the backend.
- Implement role-based access control (RBAC) to restrict access to sensitive endpoints and data.

### 23. Cost Optimization:
- Optimize cloud infrastructure costs by right-sizing instances and storage.
- Use auto-scaling to adjust resources based on demand and reduce costs during off-peak hours.
- Monitor cloud usage and performance metrics to identify areas for optimization.


These technical requirements outline the necessary components and considerations for developing the backend of a Food Delivery Service application. They cover aspects such as the choice of programming language, database management, system architecture, security measures, API design, testing, and infrastructure setup. Adjustments can be made based on specific project requirements and constraints.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
