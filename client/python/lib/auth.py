from .api.chat_pb2_grpc import AuthServiceStub


class Auth:
    def __init__(self):
        self.auth = AuthServiceStub(self.channel)

    def create_account(self):
        sessionID = self.auth.CreateRegisterSession()
        self.auth.setPassword(sessionID, "passwd")
        self.auth.registerPrimary(sessionID) #return authToken
        pass
