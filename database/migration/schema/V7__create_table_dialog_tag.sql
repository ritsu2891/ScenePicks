/**
 * @author Riku Nunokawa
 */

CREATE TABLE dialog_tag (
  tag_id INT UNSIGNED NOT NULL,
  dialog_id INT UNSIGNED NOT NULL,
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (`tag_id`) REFERENCES tag(`id`),
  FOREIGN KEY (`dialog_id`) REFERENCES dialog(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;