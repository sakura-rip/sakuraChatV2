import grpc
from .config import Config
from .talk import Talk


class ChatClient(Talk):
    def __init__(self):
        self.channel = grpc.insecure_channel(Config.HOST_URL)
        self.__init_all()

    def __init_all(self):
        Talk.__init__(self)
