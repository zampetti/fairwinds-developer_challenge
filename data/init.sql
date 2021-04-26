USE tic_tac_toe;
CREATE TABLE IF NOT EXISTS `playermoves` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `game` int(11) NOT NULL,
  `player` varchar(255) NOT NULL,
  `move` varchar(255) NOT NULL,
  `movecount` int(8) NOT NULL );

CREATE USER IF NOT EXISTS 'gameWriter'@'%' IDENTIFIED BY 'user_password';
GRANT INSERT, UPDATE, SELECT ON tic_tac_toe.PlayerMoves TO 'gameWriter'@'%';