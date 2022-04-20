- subir docker com mysql
  - `docker-compose -f docker/docker-compose.yaml up`
  - criar tabela de users:
    ```
    CREATE TABLE users (
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar(50) NOT NULL,
    email varchar(50) NOT NULL
    );
    ```
  - criar usuario para acessar o banco ip `172.22.0.1` do docker
  ```
  CREATE USER 'golang'@'172.22.0.1' IDENTIFIED BY 'golang';

  ```
  - dar permissao no usuario
  ```
  GRANT ALL PRIVILEGES ON books.* TO 'golang'@'172.22.0.1';
  ```