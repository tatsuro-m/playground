CREATE TABLE IF NOT EXISTS Users(
                                    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
                                    Username VARCHAR (127) NOT NULL UNIQUE,
                                    Password VARCHAR (127) NOT NULL,
                                    PRIMARY KEY (ID)
)
