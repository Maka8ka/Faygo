CREATE DATABASE IF NOT EXISTS faygo DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
USE faygo;
CREATE TABLE IF NOT EXISTS `faygo`.`client`(
   `client_id` INT NOT NULL AUTO_INCREMENT COMMENT 'client_id',
   `note_name` VARCHAR(1000),
   `wan_ip` VARCHAR(20),
   `lan_ip` VARCHAR(20),
   `os_type` VARCHAR(10),
   `client_mac` VARCHAR(40),
   `latest_time` TIMESTAMP,
/*`latest_task_id` INT,*/
   `client_status` INT,
   PRIMARY KEY ( `client_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `faygo`.`task`(
   `task_id` INT,
   `client_id` INT,
   `comaand_line` VARCHAR(10000),
   `command_result` text(60000),
   `return_time` TIMESTAMP,
   `command_status` INT,
   PRIMARY KEY ( `task_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `faygo`.`file`(
   `file_id` INT,
   `client_id` INT,
   `file_client_location` VARCHAR(200),
   `file_name` VARCHAR(100),
   `file_hex` text(65000),
   `transfer_time` TIMESTAMP,
   `transfer_status` INT,
   PRIMARY KEY ( `file_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;