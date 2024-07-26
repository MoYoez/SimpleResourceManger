import {BaseResponse} from "./request";
import axios from "axios";

export interface ResData {
    res_id: string;
    res_number: number;
    res_type: string;
    res_name: string;
    res_model: string;
    res_date: number;
    res_place: string;
}

export interface ResResponse  {
    list: ResData[];
}

export interface ResUpdater {
    res_id: string;
    res_number: number;
    res_type: string;
    res_name: string;
    res_model: string;
    res_date: number;
    res_place: string;
}

export function getResResource() {
    return axios.get<ResResponse>("/normal/res");
}


export function UpdateRes(form:ResUpdater) {
    return axios.post<BaseResponse>("/root/res", form);
}

export function AddRes(Form:ResUpdater) {
    return axios.post<BaseResponse>("/root/res/add", Form);
}
