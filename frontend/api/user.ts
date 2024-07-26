import {BaseResponse} from "./request";
import axios from "axios";


export interface UserListResponse {
    user_id: string;
    user_name: string;
    user_email: string;
    user_authority: number;
    user_sex: number;
}

export interface UserListRespnse extends BaseResponse {
    list: UserListResponse[];
}

export interface UserRegisterPoster {
    user_id: string;
    user_name: string;
    user_password: string;
    user_email: string;
    user_authority: number;
    user_sex: number;
}

export function getRes() {
    return axios.get<UserListRespnse>("/normal/userlist");
}

export function UpdateReg(form:UserRegisterPoster) {
    return axios.post<BaseResponse>("/admin/register", form);
}

export function ModifyUserInfo(form:UserRegisterPoster) {
    return axios.post<BaseResponse>("/normal/user", form);
}

export function DeleteRegister(Form:UserRegisterPoster) {
    return axios.delete<BaseResponse>("/admin/register", {data: Form});
}
