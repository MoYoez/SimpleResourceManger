import {BaseResponse} from "./request";
import axios from "axios";

export interface ReportMain {
    report_id: string;
    report_user:string;
    report_refer_res:string;
    report_content: string;
    report_time: number;
    report_emergency: boolean;
    report_process: number;
    report_feedback_time: number;
    report_feedback_rate: number;
    report_feedback_content: string;
    report_feedback_logs: string;
    report_feedback_user: string;
}

export interface ResResponse extends BaseResponse {
    list: ReportMain[];
}

export function getRes() {
    return axios.get<ResResponse>("/normal/report");
}

export function AddRes(Form:ReportMain) {
    return axios.post<BaseResponse>("/root/report/add", Form);
}

export function ModifyRes(Form:ReportMain) {
    return axios.post<BaseResponse>("/root/report", Form);
}
