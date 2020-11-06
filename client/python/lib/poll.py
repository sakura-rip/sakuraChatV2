from .api.chat_pb2_grpc import TalkServiceStub


class Poll:
    def __init__(self):
        self.poll = TalkServiceStub(self.channel)

