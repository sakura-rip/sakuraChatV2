
import requests
import json

from .api.chat_pb2_grpc import AuthServiceStub
from .api.chat_pb2 import (
    VerifyIDTokenRequest
)
from .config import Config


class Auth:
    def __init__(self):
        self.auth = AuthServiceStub(self.channel)

    def create_account(self):
        # Firebase sdk func
        firebase_user = self.signUpWithEmailAndPasswd(email=input(), password=input())
        # SetHeader
        ok = self.auth.VerifyIDToken(VerifyIDTokenRequest(firebase_user["idToken"]))
        if ok:
            self.profile = self.talk.GetProfile()
            self.setting = self.talk.GetSetting()
            self.talk,

    def signUpWithEmailAndPasswd(self, email: str, password: str):
        uri = f"https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key={Config.Firebase_api_key}"
        headers = {"Content-type": "application/json"}
        data = {"email": email, "password": password, "returnSecureToken": True}
        result = requests.post(
            url=uri,
            headers=headers,
            data=json.dumps(data)
        )
        return result.json()
