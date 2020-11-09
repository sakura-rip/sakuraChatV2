import json

import requests

from .api.chat_pb2 import (
    VerifyIDTokenRequest,
    GetProfileRequest,
    GetSettingRequest,
    GetAllFriendIdsRequest,
    GetAllFriendRequestIdsRequest,
    GetALlFriendRequestedIdsRequest,
    GetAllBlockedIdsRequest,
    InitPrimaryAccountRequest
)
from .api.chat_pb2_grpc import AuthServiceStub, TalkServiceStub
from .config import Config


class Auth:
    def __init__(self):
        self.auth = AuthServiceStub(self.channel)

    def save_refresh_token(self, token: str):
        with open("refresh", "w") as f:
            f.write(token)

    def create_account(self):
        # Firebase sdk func
        firebase_user = self.signUpWithEmailAndPasswd(email=input(), password=input())
        self.save_refresh_token(firebase_user["refreshToken"])
        # SetHeader
        self.auth.VerifyIDToken(VerifyIDTokenRequest(firebase_user["idToken"]))
        self.auth.InitPrimaryAccount(InitPrimaryAccountRequest())
        self.init_account()

    def init_account(self):
        self.profile = self.talk.GetProfile(GetProfileRequest())
        self.setting = self.talk.GetSetting(GetSettingRequest())
        self.friends = self.talk.GetAllFriendIds(GetAllFriendIdsRequest())
        self.friend_requests = self.talk.GetAllFriendRequestIds(GetAllFriendRequestIdsRequest())
        self.firned_requesteds = self.talk.GetALlFriendRequestedIds(GetALlFriendRequestedIdsRequest())
        self.blocked = self.talk.GetAllBlockedIds(GetAllBlockedIdsRequest())

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
