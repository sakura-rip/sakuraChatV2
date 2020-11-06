from .api.chat_pb2_grpc import TalkServiceStub


class Talk:
    def __init__(self):
        self.talk = TalkServiceStub(self.channel)
