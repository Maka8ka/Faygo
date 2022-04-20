CREATE DATABASE IF NOT EXISTS faygo DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
USE faygo;
CREATE TABLE IF NOT EXISTS `faygo`.`client`(
   `id` INT NOT NULL AUTO_INCREMENT COMMENT 'client_id',
   `cname` VARCHAR(255) COMMENT 'client_cname',
   `lan` VARCHAR(255) COMMENT 'lan_ip',
   `wan` VARCHAR(255) COMMENT 'wan_ip',
   `os` VARCHAR(255) COMMENT 'os_type',
   `mac` VARCHAR(255) COMMENT 'client_mac',
   `latest_time` VARCHAR(255) COMMENT 'latest_time',
   `status` INT COMMENT 'client_status',
   `sleep` INT DEFAULT 10 COMMENT 'sleep_time',
   `task` VARCHAR(2000) COMMENT 'client_task',
   `result` TEXT(50000) COMMENT 'task_result',
   `filehex` LONGBLOB COMMENT 'file_hex',
   `filelo` VARCHAR(255) COMMENT 'file_location',
   `notes` VARCHAR(5000) COMMENT 'client_notes',
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;