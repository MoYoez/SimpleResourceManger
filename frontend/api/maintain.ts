import {BaseResponse} from "./request";
import axios from "axios";

export interface InterfaceData {
    m_id: string;
    m_time: number;
    m_content: string;
    m_process: number;
    m_name: string;
    v_verify: boolean;
    v_time: number;
    v_user: string;
}

export interface MaintainResponse extends BaseResponse {
    list: InterfaceData[];
}

export interface MaintainUpdater {
    m_id: string;
    m_time: number;
    m_content: string;
    m_process: number;
    m_name: string;
    v_verify: boolean;
    v_time: number;
    v_user: string;
}

export function getRes() {
    return axios.get<MaintainResponse>("/normal/maintain");
}


export function UpdateRes(form:MaintainUpdater) {
    return axios.post<BaseResponse>("/root/maintain", form);
}

export function AddRes(Form:MaintainUpdater) {
    return axios.post<BaseResponse>("/root/maintain/add", Form);
}
