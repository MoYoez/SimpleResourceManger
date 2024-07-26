import axios from "axios";
import {BaseResponse} from "./request";

export interface UserInfo {
    user_id: string;
    user_name: string;
    user_password: string;
    user_email: string;
    user_authority: number;
    user_sex: number;
}

export interface UserInfoWithoutPassword { // beware that password maybe broken.
    user_id: string;
    user_name: string;
    user_email: string;
    user_authority: number;
    user_sex: number;
}
export function GetUserInfo() {
   return axios.get<UserInfo>("/normal/user")
}

export function PostUserInfoUpdater(form:UserInfo) {
    return axios.post<BaseResponse>("/normal/user", form)
}

export function PostUserInfoUpdaterWithoutPassword(form:UserInfoWithoutPassword) {
    return axios.post<BaseResponse>("/normal/user", form)
}
