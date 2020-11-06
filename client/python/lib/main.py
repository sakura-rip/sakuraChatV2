import grpc
from .config import Config


class ChatClient:
    def __init__(self):
        self.channel = grpc.insecure_channel(Config.HOST_URL)
