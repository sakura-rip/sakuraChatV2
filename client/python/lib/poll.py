from .api.chat_pb2_grpc import AuthServiceStub


class Talk:
    def __init__(self):
        self.talk = AuthServiceStub(self.channel)

    def create_account(self):
        sessionID = self.talk.CreateRegisterSession()
        pass
