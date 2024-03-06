# Core Banking System Assesment, Author: Muhammad Hafidz Erdityo

## Getting Started

1. Clone this repository.

2. Inside the `account` folder, create the following `.env` files for configuration:

   - Create `service.env` (For Service config):

     ```plaintext
     SERVICE_PORT=8100
     SERVICE_HOST=0.0.0.0
     SERVICE_WORKERS=4 // unused
     ```

   - Create `config.env` (For Database config):

     ```plaintext
     POSTGRES_USER=postgres
     POSTGRES_DB=rest_template
     POSTGRES_PASSWORD=default15p
     POSTGRES_HOST=postgres_rest_template_go
     POSTGRES_PORT=8103
     ```

   - Create `kafka_config.env` (For Kafka Topic and Broker):

   ```plaintext
   TOPIC=jurnalku
   BROKER=kafka:9092
   ```

   Feel free to change the values in the `.env` files based on your needs.

3. Inside the `journal` folder, create the following `.env` files for configuration:

   - Create `config.env` (For Database config):

     ```plaintext
     POSTGRES_USER=postgres
     POSTGRES_DB=rest_template
     POSTGRES_PASSWORD=default15p
     POSTGRES_HOST=postgres_journal_go
     POSTGRES_PORT=8104
     ```

   - Create `kafka_config.env` (For Kafka Topic and Broker):

   ```plaintext
   TOPIC=jurnalku
   BROKER=kafka:9092
   GROUP=consumerku
   ```

   Feel free to change the values in the `.env` files based on your needs.

4. On the root folder, create the following `.env` files for configuration:

   - Create `config.env` (For docker-compose config):

   ```plaintext
   KAFKA_URL=kafka
   KAFKA_PORT=9092
   KAFKA_URL2=localhost
   KAFKA_PORT2=29092
   SERVICE_PORT_ACCOUNT=8100
   SERVICE_PORT_JOURNAL=8101
   DB_ACCOUNT_PORT=8103
   DB_JOURNAL_PORT=8104
   ```

   Feel free to change the values in the `.env` files based on your needs.

5. On the root directory, run the following command to start the development server along with the database. Make sure you have docker installed:

   ```bash
   docker-compose up
   ```

6. After the service is running, you can run the short_running_script simply using
   ```bash
   go run main.go
   ```