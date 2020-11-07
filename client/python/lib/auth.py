
import requests
import json

from .api.chat_pb2_grpc import AuthServiceStub
from .config import Config


class Auth:
    def __init__(self):
        self.auth = AuthServiceStub(self.channel)

    def create_account(self):
        sessionID = self.auth.CreateRegisterSession()
        firebase_user = self.signUpWithEmailAndPasswd(email=input(), password=input())
        self.auth.VerifyIDToken(firebase_user["idToken"])

    def signUpWithEmailAndPasswd(self, email: str, password: str):
        uri = f"https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key={Config.Firebase_api_key}"
        headers = {"Content-type": "application/json"}
        data = json.dumps({"email": email, "password": password, "returnSecureToken": True})
        result = requests.post(
            url=uri,
            headers=headers,
            data=data
        )
        return result.json()
