
import configparser
f_path = './config.ini'

ini = configparser.ConfigParser()
ini.read(f_path, encoding="utf-8")


class Config:
    HOST_URL = ini["URLS"]["HOST_DOMAIN"]
